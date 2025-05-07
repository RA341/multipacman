package game

import (
	"fmt"
	"github.com/RA341/multipacman/internal/auth"
	"github.com/RA341/multipacman/internal/lobby"
	"github.com/RA341/multipacman/pkg"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Manager struct {
	activeLobbies pkg.Map[uint, *World]
	lobbyService  *lobby.Service
	mel           *melody.Melody
}

func (manager *Manager) getLobbyIdFromSession(s *melody.Session) uint {
	lobbyId, exists := s.Get(worldKey)
	if !exists {
		log.Warn().Msg("lobby id not found on disconnect")
		return 0
	}
	return lobbyId.(uint)
}

// informNewPlayerAboutOtherPlayer sends otherPlayerEntity to playerSession
func (manager *Manager) informNewPlayerAboutOtherPlayer(playerSession *melody.Session, otherPlayerEntity *PlayerEntity) bool {
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

func (manager *Manager) broadcastPlayerChange(world *World, newPlayer []byte) error {
	broadCastSessions := world.ConnectedPlayers.GetValues()
	err := manager.mel.BroadcastMultiple(newPlayer, broadCastSessions)
	if err != nil {
		return err
	}
	return nil
}

func (manager *Manager) sendGameStateInfo(newPlayerSession *melody.Session, world *World) error {
	player, err := getPlayerEntityFromSession(newPlayerSession)
	if err != nil {
		return fmt.Errorf("unable to find player: %v", err)
	}
	gameState, err := world.GetGameStateReport(player.secretToken, player.Username, string(player.SpriteType))
	if err != nil {
		return fmt.Errorf("unable to marshal game state: %v", err)
	}
	err = newPlayerSession.Write(gameState)
	if err != nil {
		return fmt.Errorf("unable to send new player, game state: %v", err)
	}
	return nil
}

func (manager *Manager) getWorld(lobby *lobby.Lobby) (*World, error) {
	activeWorld, exists := manager.activeLobbies.Load(lobby.ID)
	if !exists {
		log.Info().Msgf("creating new lobby")

		newWorld := NewWorldState()
		manager.activeLobbies.Store(lobby.ID, newWorld)
		go func() {
			newWorld.waitForGameOver()
			log.Debug().Uint("id", lobby.ID).Msg("deleting lobby")
			manager.activeLobbies.Delete(lobby.ID)
		}()

		return newWorld, nil
	}

	if activeWorld.IsLobbyFull() {
		log.Warn().Any("lobby_state", activeWorld).Msgf("lobby is full, more players are not allowed")
		return nil, fmt.Errorf("lobby is full")
	}

	return activeWorld, nil
}

func (manager *Manager) getUserAndLobbyForNewConnection(newPlayerSession *melody.Session) (*auth.User, *lobby.Lobby, error) {
	// user info
	user, err := auth.GetUserContext(newPlayerSession.Request.Context())
	if err != nil {
		log.Error().Err(err).Msg("User context error")
		return nil, nil, fmt.Errorf("no user info found")
	}

	queryParams := newPlayerSession.Request.URL.Query()
	param := queryParams.Get("lobby")
	if param == "" {
		return nil, nil, fmt.Errorf("lobbyID query parameter not found")
	}

	lobbyId, err := strconv.Atoi(param)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert %s to int: %v", param, err)
	}

	lobbyInfo, err := manager.lobbyService.GetLobbyFromID(lobbyId)
	if err != nil {
		return nil, nil, err
	}

	return user, lobbyInfo, err
}

//func (manager *Manager) sendGameOverMessage(reason string, lobbyEntity *World) {
//
//}

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
