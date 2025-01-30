package game

import (
	"encoding/json"
	"fmt"
	"github.com/RA341/multipacman/models"
	"github.com/RA341/multipacman/service"
	"github.com/RA341/multipacman/utils"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"net/http"
	"strconv"
	"time"
)

const (
	userEntityKey  = "userEntity"
	userInfKey     = "userInf"
	lobbyEntityKey = "lobbyEntity"
	lobbyIdKey     = "lobbyIdKey"
	powerUpTime    = 8 * time.Second
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

	m.HandleConnect(manager.HandleConnect)
	m.HandleMessage(manager.HandleMessage)
	m.HandleDisconnect(manager.HandleDisconnect)
	// flutter sends close signal
	m.HandleClose(func(session *melody.Session, i int, s string) error {
		manager.HandleDisconnect(session)
		return nil
	})

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
	newPlayerSession.Set(lobbyEntityKey, lobbyEntity)
	newPlayerSession.Set(userInfKey, user)
	newPlayerSession.Set(lobbyIdKey, lobby.ID)

	// we now have the new player, the lobby joined

	// inform new player about current game state
	if manager.sendGameStateInfo(newPlayerSession, lobbyEntity) {
		return
	}

	lobbyEntity.mu.Lock()

	if manager.broadcastLobbyStatus(lobbyEntity, newPlayerJson) {
		return
	}

	broadCastSessions := maps.Values(lobbyEntity.ConnectedPlayers)
	for _, otherPlayerSession := range broadCastSessions {
		if otherPlayerSession == newPlayerSession {
			continue
		}
		otherPlayerEntity, err := getPlayerEntityFromSession(otherPlayerSession)
		if err != nil {
			continue
		}
		otherPlayerEntity.Type = "active"

		// inform other players about this player
		manager.informPlayerAboutOtherPlayer(newPlayerSession, otherPlayerEntity)
	}

	lobbyEntity.mu.Unlock()

	log.Info().Any("user", *user).Any("lobby", lobby).Msgf("New player joined lobby")

	// add new player count
	manager.lobbyService.UpdateLobbyPlayerCount(lobby.ID, len(lobbyEntity.ConnectedPlayers))
	manager.lobbyService.UpdateLobbies()
}

func (manager *Manager) HandleDisconnect(s *melody.Session) {
	exitingPlayer, err := getPlayerEntityFromSession(s)
	if err != nil {
		log.Warn().Msg("Player info not found in session, on disconnect")
		return
	}
	lobbyState, err := getLobbyEntityFromSession(s)
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
	manager.broadcastLobbyStatus(lobbyState, jsonData)

	log.Info().Any("player", *exitingPlayer).Msg("client disconnected")

	lobbyId, exist := s.Get(lobbyIdKey)
	if exist {
		manager.lobbyService.UpdateLobbyPlayerCount(lobbyId.(uint), len(lobbyState.ConnectedPlayers))
		manager.lobbyService.UpdateLobbies()
	}
}

func (manager *Manager) HandleMessage(s *melody.Session, msg []byte) {
	playerSession, err := getPlayerEntityFromSession(s)
	if err != nil {
		log.Error().Msg("Player info not found in session")
		return
	}

	msgInfo := map[string]interface{}{}
	err = json.Unmarshal(msg, &msgInfo)
	if err != nil {
		log.Error().Err(err).Msg("Unable to unmarshal msg")
		return
	}

	secretToken := msgInfo["secretToken"].(string)
	if secretToken != playerSession.secretToken {
		log.Error().Msg("unable to verify secret token")
		return
	}

	lobbyInfo, err := getLobbyEntityFromSession(s)
	if err != nil {
		log.Error().Err(err).Msg("Unable to find lobby info")
		return
	}

	msgType := msgInfo["type"].(string)

	switch msgType {
	case "mov":
		x, existsx := msgInfo["x"]
		y, existsy := msgInfo["y"]
		dir, existsdir := msgInfo["dir"]
		if !existsy || !existsx || !existsdir {
			log.Warn().Msg("no x,y, dir info found")
			return
		}

		lobbyInfo.MovePlayer(playerSession, x.(string), y.(string))
		playerSession.Type = "mov"
		playerSession.Dir = dir.(string)
		jsonData, err := playerSession.ToJSON()
		if err != nil {
			log.Warn().Err(err).Msg("Error while marshalling json")
			return
		}
		manager.broadcastLobbyStatus(lobbyInfo, jsonData)
		return
	case "pel":
		id, exists := msgInfo["id"]
		if !exists {
			log.Warn().Any("msg", msgInfo).Msg("no pellet id found")
			return
		}

		lobbyInfo.EatPellet(int(id.(float64)))
		encoded, err := json.Marshal(map[string]interface{}{
			"id":   id,
			"type": "pel",
		})
		if err != nil {
			log.Error().Err(err).Msg("Error marshalling json")
			return
		}

		manager.broadcastLobbyStatus(lobbyInfo, encoded)
		return
	case "pow":
		id, exists := msgInfo["id"]
		if !exists {
			log.Warn().Any("msg", msgInfo).Msg("no pellet id found")
			return
		}

		lobbyInfo.EatPowerUp(int(id.(float64)))
		time.AfterFunc(powerUpTime, func() {
			manager.sendEndPowerUpMessage(lobbyInfo)
		})

		encoded, err := json.Marshal(map[string]interface{}{
			"type": "pow",
			"id":   id,
		})
		if err != nil {
			log.Error().Err(err).Msg("Error marshalling json")
			return
		}

		manager.broadcastLobbyStatus(lobbyInfo, encoded)
		return
	case "gho":
		ghostId, exists := msgInfo["ghId"]
		if !exists {
			log.Warn().Any("msg", msgInfo).Msg("no pellet ghostId found")
			return
		}

		var err error
		var encoded []byte
		if lobbyInfo.IsPoweredUp {
			lobbyInfo.GhostEatenAction(SpriteType(ghostId.(string)))
			// ghost eaten
			encoded, err = json.Marshal(map[string]interface{}{
				"type":    "gho",
				"ghostId": ghostId,
			})
		} else {
			// pacman eaten
			// send pacman dead, game over
			encoded, err = json.Marshal(map[string]interface{}{
				"type": "pacd",
			})
		}
		if err != nil {
			log.Error().Err(err).Msg("Error marshalling json")
			return
		}

		manager.broadcastLobbyStatus(lobbyInfo, encoded)

		return
	default:
		log.Warn().Msgf("Unknown message type: %s", msgType)
	}
}

////////////////////////////
// Utility functions

func (manager *Manager) sendGameOverMessage(reason string, lobbyEntity *LobbyState) {

}

func (manager *Manager) sendEndPowerUpMessage(lobbyEntity *LobbyState) {
	lobbyEntity.IsPoweredUp = false
	encoded, err := json.Marshal(map[string]interface{}{
		"type": "nopow",
	})
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling json")
		return
	}

	manager.broadcastLobbyStatus(lobbyEntity, encoded)
}

func (manager *Manager) broadcastLobbyStatus(lobbyEntity *LobbyState, newPlayerJson []byte) bool {
	broadCastSessions := maps.Values(lobbyEntity.ConnectedPlayers)
	err := manager.mel.BroadcastMultiple(newPlayerJson, broadCastSessions)
	if err != nil {
		log.Error().Err(err).Msg("Unable to broadcast lobby info")
		return true
	}
	return false
}

func (manager *Manager) getLobbyIdFromSession(s *melody.Session) uint {
	lobbyId, exists := s.Get(lobbyEntityKey)
	if !exists {
		log.Warn().Msg("lobby id not found on disconnect")
		return 0
	}
	return lobbyId.(uint)
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
	player, err := getPlayerEntityFromSession(newPlayerSession)
	if err != nil {
		log.Error().Err(err).Msg("Unable to find player")
		return false
	}
	gameState, err := lobbyEntity.GetGameStateReport(player.secretToken, string(player.SpriteType))
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
