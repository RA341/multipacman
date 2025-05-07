package game

import (
	"encoding/json"
	"github.com/RA341/multipacman/internal/auth"
	"github.com/RA341/multipacman/internal/lobby"
	"github.com/RA341/multipacman/pkg"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const (
	userInfoKey = "userEntity"
	userInfKey  = "userInf"
	worldKey    = "lobbyEntity"
	lobbyIdKey  = "lobbyIdKey"
	powerUpTime = 8 * time.Second
)

type WsHandler struct {
	lobbyService    *lobby.Service
	msgHandlerFuncs map[string]MessageHandlerFunc
	manager         *Manager
}

func RegisterGameWSHandler(mux *http.ServeMux, authService *auth.Service, lobbyService *lobby.Service) {
	mel := melody.New()
	manager := &Manager{
		lobbyService:  lobbyService,
		mel:           mel,
		activeLobbies: pkg.Map[uint, *World]{},
	}
	handler := WsHandler{
		lobbyService: lobbyService,
		manager:      manager,
		msgHandlerFuncs: registerMessageHandlers(
			MovMessage(),
			PowerUpMessage(manager),
			KillPlayer(),
			PelletMessage(),
		),
	}

	mel.HandleConnect(handler.HandleConnect)
	mel.HandleMessage(handler.HandleMessage)
	mel.HandleDisconnect(handler.HandleDisconnect)
	// flutter sends close signal
	mel.HandleClose(func(session *melody.Session, i int, s string) error {
		handler.HandleDisconnect(session)
		return nil
	})

	wsHandler := func(w http.ResponseWriter, r *http.Request) {
		err := mel.HandleRequest(w, r)
		if err != nil {
			http.Error(w, "WebSocket connection failed", http.StatusInternalServerError)
			return
		}
	}

	mux.Handle("/api/game", WsAuthMiddleware(authService, http.HandlerFunc(wsHandler)))
}

////////////////////////////
// main handlers

func (h *WsHandler) HandleConnect(newPlayerSession *melody.Session) {
	user, lobbyInfo, err := h.manager.getUserAndLobbyForNewConnection(newPlayerSession)
	if err != nil {
		log.Error().Err(err).Msg("Unable to find lobby or user info")
		return
	}

	world, err := h.manager.getWorld(lobbyInfo)
	if err != nil {
		sendMessage(newPlayerSession, wsError(err))
		return
	}

	player := NewPlayerEntity(user.ID, user.Username)

	err = world.Join(player, newPlayerSession)
	if err != nil {
		log.Error().Err(err).Msg("Unable to join lobby")
		sendMessage(newPlayerSession, wsError(err))
		return
	}

	newPlayerJson, err := player.ToJSON()
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert player to JSON")
		return
	}

	// store session info
	newPlayerSession.Set(userInfoKey, player)
	newPlayerSession.Set(worldKey, world)
	newPlayerSession.Set(userInfKey, user)
	newPlayerSession.Set(lobbyIdKey, lobbyInfo.ID)

	// we now have the new player, with lobby joined

	// inform new player about current game state
	if err := h.manager.sendGameStateInfo(newPlayerSession, world); err != nil {
		log.Error().Err(err).Msg("Unable to send game state info")
		return
	}

	// inform new player has joined to existing players
	if err := h.manager.broadcastPlayerChange(world, newPlayerJson); err != nil {
		log.Error().Err(err).Msg("Unable to broadcast status")
		return
	}

	broadCastSessions := world.ConnectedPlayers.GetValues()
	for _, otherPlayerSession := range broadCastSessions {
		if otherPlayerSession == newPlayerSession {
			continue
		}

		otherPlayerEntity, err := getPlayerEntityFromSession(otherPlayerSession)
		if err != nil {
			continue
		}
		otherPlayerEntity.Type = "active"

		// inform this player about other players
		h.manager.informNewPlayerAboutOtherPlayer(newPlayerSession, otherPlayerEntity)
	}

	log.Info().Any("user", *user).Any("lobby", lobbyInfo).Msgf("New player joined lobby")

	// add new player count
	h.lobbyService.UpdateLobbyPlayerCount(lobbyInfo.ID, len(broadCastSessions))
	h.lobbyService.UpdateLobbies()
}

func (h *WsHandler) HandleDisconnect(s *melody.Session) {
	exitingPlayer, err := getPlayerEntityFromSession(s)
	if err != nil {
		log.Warn().Msg("player not found in session on disconnect")
		return
	}
	lobbyState, err := getWorldFromSession(s)
	if err != nil {
		log.Warn().Msg("Lobby not found in active lobbies on disconnect")
		return
	}

	lobbyState.Leave(exitingPlayer)
	// set disconnect status
	exitingPlayer.Type = "dis"

	// inform other players
	jsonData, err := exitingPlayer.ToJSON()
	if err != nil {
		log.Error().Err(err).Any("other entity", exitingPlayer).Msg("Failed to convert PlayerEntity to JSON")
		return
	}
	// inform active players about player that left
	_ = h.manager.broadcastPlayerChange(lobbyState, jsonData)

	log.Info().Any("player", *exitingPlayer).Msg("client disconnected")

	lobbyId, exist := s.Get(lobbyIdKey)
	if exist {
		h.lobbyService.UpdateLobbyPlayerCount(lobbyId.(uint), len(lobbyState.ConnectedPlayers.GetValues()))
		h.lobbyService.UpdateLobbies()
	}
}

func (h *WsHandler) HandleMessage(s *melody.Session, msg []byte) {
	playerSession, err := getPlayerEntityFromSession(s)
	if err != nil {
		log.Error().Msg("Player info not found in session")
		return
	}

	msgInfo := map[string]interface{}{}
	if err = json.Unmarshal(msg, &msgInfo); err != nil {
		log.Error().Err(err).Msg("Unable to unmarshal msg")
		return
	}

	secretToken := msgInfo["secretToken"].(string)
	if secretToken != playerSession.secretToken {
		log.Error().Msg("unable to verify secret token")
		return
	}

	lobbyInfo, err := getWorldFromSession(s)
	if err != nil {
		log.Error().Err(err).Msg("Unable to find lobby info")
		return
	}
	msgType := msgInfo["type"].(string)

	msgHandler, ok := h.msgHandlerFuncs[msgType]
	if !ok {
		log.Warn().Msgf("Unknown message type: %s", msgType)
		return
	}

	data := msgHandler(MessageData{msgInfo, lobbyInfo, playerSession})
	if data == nil {
		log.Debug().Msg("null message, something went wrong while handling message")
		return
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Any("msg", data).Msg("Unable to marshal msg")
		return
	}

	_ = h.manager.broadcastPlayerChange(lobbyInfo, marshal)

}

func sendMessage(session *melody.Session, message []byte) {
	err := session.Write(message)
	if err != nil {
		log.Error().Err(err).Msg("Unable to send message")
	}
}

func wsError(err error) []byte {
	marshal, _ := json.Marshal(map[string]string{"error": err.Error()})
	return marshal
}
