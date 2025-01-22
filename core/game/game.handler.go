package game

import (
	"fmt"
	"github.com/RA341/multipacman/models"
	"github.com/RA341/multipacman/service"
	"github.com/RA341/multipacman/utils"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

const (
	userEntityKey = "userEntity"
	userInfKey    = "userInf"
	lobbyIdKey    = "lobby"
)

type Manager struct {
	lobbyService  *service.LobbyService
	activeLobbies map[uint]*LobbyState
	mel           *melody.Melody
}

func InitGameWsHandler(mux *http.ServeMux, authService *service.AuthService, lobbyService *service.LobbyService) {
	m := melody.New()
	manager := Manager{
		lobbyService:  lobbyService,
		activeLobbies: make(map[uint]*LobbyState),
		mel:           m,
	}

	manager.mel.HandleConnect(manager.HandleConnect)
	manager.mel.HandleMessage(manager.HandleMessage)
	manager.mel.HandleDisconnect(manager.HandleDisconnect)

	wsHandler := func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			http.Error(w, "WebSocket connection failed", http.StatusInternalServerError)
			return
		}
	}

	mux.Handle("/api/game", AuthMiddleware(authService, http.HandlerFunc(wsHandler)))
}

////////////////////////////
// main handlers

func (manager *Manager) HandleConnect(newPlayerSession *melody.Session) {
	user, lobby, err := manager.getUserAndLobbyInfo(newPlayerSession)
	if err != nil {
		log.Error().Err(err).Msg("Unable to find lobby or user info")
		return
	}

	lobbyEntity, err := manager.getLobbyEntityFromLobbyModel(lobby)
	if err != nil {
		err := newPlayerSession.Write([]byte(`{"error": "lobby is full"}`))
		if err != nil {
			log.Error().Err(err).Msg("Unable to write lobby info")
			return
		}
		return
	}

	playerEntity := NewPlayerEntity(user.ID, user.Username)

	err = lobbyEntity.Join(playerEntity, newPlayerSession)
	if err != nil {
		log.Error().Err(err).Msg("Unable to join lobby")

		err := newPlayerSession.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		if err != nil {
			log.Error().Err(err).Msg("Unable to write lobby info")
			return
		}

		return
	}

	newPlayerJson, err := playerEntity.ToJSON()
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert player to JSON")
		return
	}

	// store session info
	newPlayerSession.Set(userEntityKey, playerEntity)
	newPlayerSession.Set(lobbyIdKey, lobby.ID)
	newPlayerSession.Set(userInfKey, user)

	// we now have the new player, the lobby joined

	lobbyEntity.mu.Lock()
	for _, otherPlayerSession := range lobbyEntity.ConnectedPlayers {
		otherPlayerEntity := manager.getPlayerInfoFromSession(otherPlayerSession)
		if otherPlayerEntity == nil {
			continue
		}

		if otherPlayerEntity.PlayerId == playerEntity.PlayerId {
			// sending player info to self
			err = newPlayerSession.Write(newPlayerJson)
			if err != nil {
				log.Error().Err(err).Msg("Unable to send new player info")
				return
			}
			continue
		}

		// inform other players about this player
		if manager.informPlayerAboutOtherPlayer(otherPlayerSession, playerEntity) {
			return
		}

		// inform this player about other players
		if manager.informPlayerAboutOtherPlayer(newPlayerSession, otherPlayerEntity) {
			return
		}
	}
	lobbyEntity.mu.Unlock()

	log.Info().Any("user info", *user).Any("lobby", lobby).Msgf("New player joined lobby")

	// inform new player about current game state
	if manager.sendGameStateInfo(newPlayerSession, lobbyEntity) {
		return
	}

	log.Debug().Msg("Refreshing lobbies")
	manager.lobbyService.UpdateLobbies()
}

func (manager *Manager) HandleDisconnect(s *melody.Session) {
	exitingPlayer := manager.getPlayerInfoFromSession(s)
	if exitingPlayer == nil {
		log.Warn().Msg("Player info not found in session, on disconnect")
		return
	}

	lobbyId := manager.getLobbyIdFromSession(s)
	if lobbyId == 0 {
		return
	}

	lobbyState, exists := manager.activeLobbies[lobbyId]
	if !exists {
		log.Warn().Msg("Lobby not found in active lobbies")
		return
	}

	lobbyState.Leave(exitingPlayer)
	// set disconnect status
	exitingPlayer.Type = "dis"

	// inform other players
	for _, currentPlayers := range lobbyState.ConnectedPlayers {
		manager.informPlayerAboutOtherPlayer(currentPlayers, exitingPlayer)
	}

	log.Info().Any("player", *exitingPlayer).Msg("client disconnected")
}

func (manager *Manager) HandleMessage(s *melody.Session, msg []byte) {}

////////////////////////////
// Utility functions

func (manager *Manager) getLobbyIdFromSession(s *melody.Session) uint {
	lobbyId, exists := s.Get(lobbyIdKey)
	if !exists {
		log.Warn().Msg("lobby id not found on disconnect")
		return 0
	}
	return lobbyId.(uint)
}

func (manager *Manager) getPlayerInfoFromSession(playerSession *melody.Session) *PlayerEntity {
	pInfo, exists := playerSession.Get(userEntityKey)
	if !exists {
		log.Info().Msg("Player info not found in sessions")
		return nil
	}
	otherPlayerEntity := pInfo.(*PlayerEntity)
	return otherPlayerEntity
}

// informPlayerAboutOtherPlayer sends otherPlayerEntity to playerSession
func (manager *Manager) informPlayerAboutOtherPlayer(playerSession *melody.Session, otherPlayerEntity *PlayerEntity) bool {
	jsonData, err := otherPlayerEntity.ToJSON()
	if err != nil {
		log.Error().Err(err).Any("other entity", otherPlayerEntity).Msg("Failed to convert PlayerEntity to JSON")
		return true
	}
	err = playerSession.Write(jsonData)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send player info")
		return true
	}
	return false
}

func (manager *Manager) sendGameStateInfo(newPlayerSession *melody.Session, lobbyEntity *LobbyState) bool {
	gameState, err := lobbyEntity.GetGameStateReport()
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal game state")
		return true
	}
	err = newPlayerSession.Write(gameState)
	if err != nil {
		log.Error().Err(err).Msg("Unable to send new player, game state")
		return true
	}
	return false
}

func (manager *Manager) getLobbyEntityFromLobbyModel(lobby *models.Lobby) (*LobbyState, error) {
	activeLobby, exists := manager.activeLobbies[lobby.ID]
	if exists {
		if activeLobby.IsLobbyFull() {
			log.Warn().Any("lobby_state", activeLobby).Msgf("lobby is full, more players are not allowed")
			return nil, fmt.Errorf("lobby is full")
		}
	} else {
		log.Info().Msgf("creating new lobby")
		manager.activeLobbies[lobby.ID] = NewLobbyStateEntity()

		activeLobby = manager.activeLobbies[lobby.ID]
	}
	return activeLobby, nil
}

func (manager *Manager) getUserAndLobbyInfo(newPlayerSession *melody.Session) (*models.User, *models.Lobby, error) {
	// user info
	user, err := utils.GetUserContext(newPlayerSession.Request.Context())
	if err != nil {
		log.Error().Err(err).Msg("User context error")
		return nil, nil, fmt.Errorf("no user info found")
	}

	// get lobby id
	queryParams := newPlayerSession.Request.URL.Query()
	tmp := queryParams.Get("lobby")
	lobbyId, err := strconv.Atoi(tmp)
	if err != nil {
		log.Error().Err(err).Str("lobby id from the query", tmp).Msg("Failed to convert lobby id")
		return nil, nil, fmt.Errorf("invalid lobby id")
	}

	// get lobby info
	lobby, err := manager.lobbyService.GetLobbyFromID(lobbyId)
	if err != nil {
		return nil, nil, err
	}

	return user, lobby, err
}

//func HandleMessage(s *melody.Session, m *melody.Melody, msg []byte) {
//	var data map[string]interface{}
//
//	tmp, exists := s.Get("LobbyId")
//	if !exists {
//		log.Print("lobby id not found on message")
//		return
//	}
//	lobbyId := tmp.(int)
//	//LobbyList[lobbyId.(string)].
//
//	err := json.Unmarshal(msg, &data)
//	if err != nil {
//		log.Print("Failed to unmarshal ws msg")
//		return
//	}
//
//	// for these message types broadcast to all clients
//	// included the once who sent this ws request
//	// List of strings
//	messageType := data["type"].(string)
//
//	switch messageType {
//	//case "pos":
//	//	x, y := retrieveCoordinates(data)
//	//	log.Print(x)
//	//	log.Print(y)
//
//	case "pellet":
//		x, y := retrieveCoordinates(data)
//		LobbyList[lobbyId].PelletEatenAction(x, y)
//		//fmt.Println("Handling " + messageType)
//		broadCastAll(m, msg, lobbyId)
//	case "power":
//		x, y := retrieveCoordinates(data)
//		LobbyList[lobbyId].PowerUpEatenAction(x, y)
//		//fmt.Println("Handling " + messageType)
//		broadCastAll(m, msg, lobbyId)
//	case "pacded":
//		ghostId := data["id"].(string)
//		LobbyList[lobbyId].GhostEatenAction(ghostId)
//		//fmt.Println("Handling " + messageType)
//		broadCastAll(m, msg, lobbyId)
//	default:
//		//fmt.Println("Broadcasting others for type " + messageType)
//		broadCastOthers(m, msg, s, lobbyId)
//	}
//}
//
//func retrieveCoordinates(data map[string]interface{}) (float64, float64) {
//	x := data["x"].(float64)
//	y := data["y"].(float64)
//	return x, y
//}
//
//func broadCastAll(m *melody.Melody, msg []byte, lobbyId int) {
//	var sessionList []*melody.Session
//
//	for session := range LobbyList[lobbyId].ConnectedPlayers {
//		sessionList = append(sessionList, LobbyList[lobbyId].ConnectedPlayers[session])
//	}
//
//	err := m.BroadcastMultiple(msg, sessionList)
//	if err != nil {
//		log.Printf("Failed to send data" + err.Error())
//		return
//	}
//}
//
//func broadCastOthers(m *melody.Melody, msg []byte, session *melody.Session, lobbyId int) {
//	var sessionList []*melody.Session
//
//	for sessionKeys := range LobbyList[lobbyId].ConnectedPlayers {
//		tmpSession := LobbyList[lobbyId].ConnectedPlayers[sessionKeys]
//		// ignore the session calling the message
//		if session != tmpSession {
//			sessionList = append(sessionList, tmpSession)
//		}
//	}
//
//	err := m.BroadcastMultiple(msg, sessionList)
//	if err != nil {
//		log.Fatal("Failed to send data" + err.Error())
//		return
//	}
//}
