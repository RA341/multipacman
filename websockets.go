package main

import (
	"fmt"
	"github.com/olahol/melody"
	"log"
	"server/entities"
)

func initMelody(m *melody.Melody) {
	// Set up Melody event handlers
	m.HandleConnect(func(session *melody.Session) {
		HandleConnect(session, m)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		HandleMessage(s, m, msg)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		HandleDisconnect(s, m)
	})
}

// HandleConnect websocket stuff
func HandleConnect(newPlayerSession *melody.Session, m *melody.Melody) {
	// Retrieve the http.Request associated with the WebSocket connection
	queryParams := newPlayerSession.Request.URL.Query()

	// get userid and lobbyid
	userId := queryParams.Get("user")
	lobbyId := queryParams.Get("lobby")

	log.Print("New player with id: " + userId + " joined lobby: " + lobbyId)

	// todo remove once lobby page is implemented
	// add a tmp lobby
	if LobbyList[lobbyId] == nil {
		LobbyList[lobbyId] = entities.NewLobbyModel()
	}

	// todo check if lobby id and user id exist
	if lobbyId == "" || userId == "" {
		// throw error if it doesn't exist
		log.Print(newPlayerSession.Write([]byte(`{"redirect": "/redirected_page"}`)))
		return
	}

	lobby := LobbyList[lobbyId]
	// get lobby info if full throw error
	if len(lobby.ConnectedPlayers) == 4 {
		fmt.Println("Whoa there more players are not allowed")

		log.Print(newPlayerSession.Write([]byte(`{"redirect": "/static/game/full.html"}`)))
		return
	}

	// get user info
	// get current player
	newPlayer := entities.NewPlayerEntity()

	newPlayer.Username = "TODO username"
	newPlayer.Type = "join"
	newPlayer.PlayerId = userId
	lobby.Join(newPlayer)

	newPlayerJson, err := newPlayer.ToJSON()
	if err != nil {
		log.Fatal("Failed to send player info" + err.Error())
		return
	}

	// sending new player info
	allSessions, _ := m.Sessions()

	// tell new player about current players
	for _, otherPlayerSession := range allSessions {
		pInfo, exists := otherPlayerSession.Get("info")

		if !exists {
			fmt.Println("PlayerEntity does not exist")
			continue
		}

		otherPlayer := pInfo.(*entities.PlayerEntity)

		// tell current player about other player
		json, err := otherPlayer.ToJSON()
		if err != nil {
			log.Print("Failed to convert PlayerEntity to JSON")
			return
		}

		err = newPlayerSession.Write(json)
		if err != nil {
			log.Fatal("Failed to send player info" + err.Error())
			return
		}
	}

	// store session info
	newPlayerSession.Set("info", newPlayer)
	newPlayerSession.Set("LobbyId", lobbyId)

	// tell other players about joined player
	err = m.BroadcastOthers(newPlayerJson, newPlayerSession)
	if err != nil {
		log.Fatal("Failed to send player info" + err.Error())
		return
	}
	// sending player info to self
	err = newPlayerSession.Write(newPlayerJson)
	if err != nil {
		log.Fatal("Failed to send player info" + err.Error())
		return
	}

	gameState := LobbyList[lobbyId].GetGameStateReport()
	if err != nil {
		log.Fatal("Failed to marshal game state" + err.Error())
		return
	}
	err = newPlayerSession.Write(gameState)
	if err != nil {
		log.Fatal("Failed to send game state" + err.Error())
		return
	}
}

func HandleDisconnect(s *melody.Session, m *melody.Melody) {
	fmt.Println("PlayerEntity exiting")
	value, exists := s.Get("info")
	if !exists {
		log.Print("Player info not found on disconnect")
		return
	}

	lobbyId, exists := s.Get("LobbyId")
	if !exists {
		log.Print("lobby id not found on disconnect")
		return
	}

	playerInfo := value.(*entities.PlayerEntity)
	LobbyList[lobbyId.(string)].Leave(playerInfo)

	playerInfo.Type = "dis"

	json, err := playerInfo.ToJSON()
	if err != nil {
		log.Print("Failed to convert PlayerEntity to JSON on exit")
		return
	}

	err = m.BroadcastOthers(json, s)
	if err != nil {
		log.Print("Failed to send player info on exit")
		return
	}
}

func HandleMessage(s *melody.Session, m *melody.Melody, msg []byte) {
	err := m.BroadcastOthers(msg, s)
	if err != nil {
		log.Fatal("Failed to send data" + err.Error())
		return
	}
}
