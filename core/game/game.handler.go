package game

import (
	"github.com/RA341/multipacman/models"
	"github.com/RA341/multipacman/service"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
)

// game websocket handler

var (
	mu2           = sync.RWMutex{}
	connections2  = map[uint]*melody.Session{}
	updateChannel = make(chan bool)
)

func initLobbyWsHandler(mux *http.ServeMux, authService *service.AuthService, lobbyService *service.LobbyService) {
	m := melody.New()
	InitLobbyWs(m, lobbyService)

	wsHandler := func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			http.Error(w, "WebSocket connection failed", http.StatusInternalServerError)
			return
		}
	}

	mux.Handle("/api/lobbies", AuthMiddleware(authService, http.HandlerFunc(wsHandler)))
	go sendLobbyUpdates(lobbyService)
}

func InitLobbyWs(m *melody.Melody, lobbyService *service.LobbyService) {
	m.HandleConnect(func(session *melody.Session) {
		HandleConnect(session, m, lobbyService)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		HandleMessage(s, m, msg)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		HandleDisconnect(s, m)
	})
}

func HandleConnect(newPlayerSession *melody.Session, m *melody.Melody, lobbyService *service.LobbyService) {
	user := newPlayerSession.Request.Context().Value("user")
	if user == nil {
		log.Error().Msg("No user found in context")
		err := newPlayerSession.CloseWithMsg([]byte("Unable to determine user info"))
		if err != nil {
			log.Error().Err(err).Msg("Unable close connection")
		}
		return
	}

	verifiedUser := user.(*models.User)
	newPlayerSession.Set("user", verifiedUser)

	mu2.Lock()
	defer mu2.Unlock()
	connections2[verifiedUser.ID] = newPlayerSession

	log.Info().
		Str("username", verifiedUser.Username).
		Uint("id", verifiedUser.ID).
		Msg("client connected")

	// initial data
	jsonData, err := lobbyService.GetAndParseLobbies()
	if err != nil {
		return
	}

	err = newPlayerSession.Write(jsonData)
	if err != nil {
		log.Error().Err(err).Msg("Unable to send lobby data")
	}
}

func HandleDisconnect(s *melody.Session, m *melody.Melody) {
	user, status := s.Get("user")
	if status != true {
		log.Warn().Msg("no value associated with user key")
		return
	}

	id, username := user.(*models.User).ID, user.(*models.User).Username
	connection := connections2[id]
	if connection == nil {
		log.Warn().Uint("id", id).Msg("no connection found for id")
		return
	}
	// remove from connection
	delete(connections2, id)

	err := connection.Close()
	if err != nil {
		log.Error().Err(err).Msg("err while close connection")
		return
	}

	log.Info().Str("username", username).Msg("client disconnected")
}

func sendLobbyUpdates(lobby *service.LobbyService) {
	log.Info().Msg("starting channel watcher")

	for msg := range updateChannel {
		log.Info().Bool("message", msg).Msg("received message")

		jsonData, err := lobby.GetAndParseLobbies()
		if err != nil {
			continue
		}

		for _, session := range connections2 {
			err = session.Write(jsonData)
			if err != nil {
				log.Error().Err(err).Msg("Unable to write to session")
			}
		}
	}
}

// HandleMessage no need to handle message for lobbies
func HandleMessage(s *melody.Session, m *melody.Melody, msg []byte) {}

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/olahol/melody"
//	"log"
//	"server/entities"
//	"strconv"
//)
//
//func initMelody(m *melody.Melody) {
//	// Set up Melody event handlers
//	m.HandleConnect(func(session *melody.Session) {
//		HandleConnect(session, m)
//	})
//
//	m.HandleMessage(func(s *melody.Session, msg []byte) {
//		HandleMessage(s, m, msg)
//	})
//
//	m.HandleDisconnect(func(s *melody.Session) {
//		HandleDisconnect(s, m)
//	})
//}
//
//// HandleConnect websocket stuff
//func HandleConnect(newPlayerSession *melody.Session, m *melody.Melody) {
//	// Retrieve the http.Request associated with the WebSocket connection
//	queryParams := newPlayerSession.Request.URL.Query()
//
//	tml := newPlayerSession.Request.Context().Value("user")
//	if tml == nil || tml == "" {
//		tml = "Unknown username"
//	}
//	userName := tml.(string)
//
//	// get userid and lobbyid
//	userId := queryParams.Get("user")
//	tmp := queryParams.Get("lobby")
//	lobbyId, _ := strconv.Atoi(tmp)
//	log.Print("New player with id: " + userId + " joined lobby:")
//
//	// todo remove once lobby page is implemented
//	// add a tmp lobby
//	if LobbyList[lobbyId] == nil {
//		log.Print(newPlayerSession.Write([]byte(`{"redirect": "/lobby"}`)))
//		return
//	}
//
//	lobby := LobbyList[lobbyId]
//	// get lobby info if full throw error
//	if len(lobby.ConnectedPlayers) == 4 {
//		fmt.Println("more players are not allowed")
//		log.Print(newPlayerSession.Write([]byte(`{"redirect": "/static/game/full.html"}`)))
//		return
//	}
//
//	// get user info
//	// get current player
//	newPlayer := entities.NewPlayerEntity()
//
//	newPlayer.Username = userName
//	newPlayer.Type = "join"
//	newPlayer.PlayerId = userId
//	lobby.Join(newPlayer, newPlayerSession)
//
//	newPlayerJson, err := newPlayer.ToJSON()
//	if err != nil {
//		log.Fatal("Failed to send player info" + err.Error())
//		return
//	}
//
//	// sending new player info
//	allSessions, _ := m.Sessions()
//
//	// tell new player about current players
//	for _, otherPlayerSession := range allSessions {
//		pInfo, exists := otherPlayerSession.Get("info")
//
//		if !exists {
//			fmt.Println("PlayerEntity does not exist")
//			continue
//		}
//
//		otherPlayer := pInfo.(*entities.PlayerEntity)
//
//		// tell current player about other player
//		jsonData, err := otherPlayer.ToJSON()
//		if err != nil {
//			log.Print("Failed to convert PlayerEntity to JSON")
//			return
//		}
//
//		err = newPlayerSession.Write(jsonData)
//		if err != nil {
//			log.Fatal("Failed to send player info" + err.Error())
//			return
//		}
//	}
//
//	// store session info
//	newPlayerSession.Set("info", newPlayer)
//	newPlayerSession.Set("LobbyId", lobbyId)
//
//	// tell other players about joined player
//	err = m.BroadcastOthers(newPlayerJson, newPlayerSession)
//	if err != nil {
//		log.Fatal("Failed to send player info" + err.Error())
//		return
//	}
//	// sending player info to self
//	err = newPlayerSession.Write(newPlayerJson)
//	if err != nil {
//		log.Fatal("Failed to send player info" + err.Error())
//		return
//	}
//
//	gameState := LobbyList[lobbyId].GetGameStateReport()
//	if err != nil {
//		log.Fatal("Failed to marshal game state" + err.Error())
//		return
//	}
//	err = newPlayerSession.Write(gameState)
//	if err != nil {
//		log.Fatal("Failed to send game state" + err.Error())
//		return
//	}
//}
//
//func HandleDisconnect(s *melody.Session, m *melody.Melody) {
//	fmt.Println("PlayerEntity exiting")
//	value, exists := s.Get("info")
//	if !exists {
//		log.Print("Player info not found on disconnect")
//		return
//	}
//
//	lobbyId, exists := s.Get("LobbyId")
//	if !exists {
//		log.Print("lobby id not found on disconnect")
//		return
//	}
//
//	playerInfo := value.(*entities.PlayerEntity)
//	LobbyList[lobbyId.(int)].Leave(playerInfo)
//
//	playerInfo.Type = "dis"
//
//	jsonData, err := playerInfo.ToJSON()
//	if err != nil {
//		log.Print("Failed to convert PlayerEntity to JSON on exit")
//		return
//	}
//
//	err = m.BroadcastOthers(jsonData, s)
//	if err != nil {
//		log.Print("Failed to send player info on exit")
//		return
//	}
//}
//
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
