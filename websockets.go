package main

import (
	"github.com/olahol/melody"
	"log"
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
	//
	//if started {
	//	fmt.Println("Whoa there more players are not allowed")
	//	return
	//}
	//
	//allSessions, _ := m.Sessions()
	//// get current player
	//currentUserId := newPlayerSession.Request.URL.Query().Get("userId")
	//currentPlayer := lobby[currentUserId]
	//currentPlayer.Type = "join"
	//var spriteId string
	//availablePlayers, spriteId = popItem(availablePlayers)
	//currentPlayer.SpriteType = spriteId
	//
	//// tell new player 	about current players
	//for _, otherPlayerSession := range allSessions {
	//	pInfo, exists := otherPlayerSession.Get("info")
	//
	//	if !exists {
	//		fmt.Println("Player does not exist")
	//		continue
	//	}
	//
	//	otherPlayer := pInfo.(game.Player)
	//
	//	// tell current player about other player
	//	err := newPlayerSession.Write(otherPlayer.ToJson())
	//	if err != nil {
	//		log.Fatal("Failed to send player info" + err.Error())
	//		return
	//	}
	//}
	//
	//// store session info
	//newPlayerSession.Set("info", currentPlayer)
	//
	//// tell other players about joined player
	//err := m.BroadcastOthers(currentPlayer.ToJson(), newPlayerSession)
	//
	//if err != nil {
	//	log.Fatal("Failed to send player info" + err.Error())
	//	return
	//}
	//
	//err = newPlayerSession.Write(currentPlayer.ToJson())
	//
	//if err != nil {
	//	log.Fatal("Failed to send player info" + err.Error())
	//	return
	//}
}

func HandleDisconnect(s *melody.Session, m *melody.Melody) {
	//fmt.Println("Player exiting")
	//value, exists := s.Get("info")
	//
	//if !exists {
	//	return
	//}
	//
	//info := value.(game.Player)
	//availablePlayers = append(availablePlayers, info.SpriteType)
	//info.Type = "dis"
	//
	//err := m.BroadcastOthers(info.ToJson(), s)
	//if err != nil {
	//	return
	//}
}

func HandleMessage(s *melody.Session, m *melody.Melody, msg []byte) {
	err := m.BroadcastOthers(msg, s)
	if err != nil {
		log.Fatal("Failed to send data" + err.Error())
		return
	}
}
