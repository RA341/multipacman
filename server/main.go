package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/olahol/melody"
	"log"
	"net/http"
	"os"
	"server/game"
)

var (
	started = false
)

const lobbySize = 2

func shuffleAndCreatePlayerList() []string { // Create a list (slice)
	myList := []string{"pacman", "ghost1"}
	return myList
}

func popItem(myList []string) ([]string, string) {
	poppedElement := myList[len(myList)-1]
	myList = myList[:len(myList)-1]
	return myList, poppedElement
}

func main() {

	lobby := make(map[string]game.Player)
	availablePlayers := shuffleAndCreatePlayerList()

	m := melody.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/index.html")
	})

	http.HandleFunc("/scripts/*", func(w http.ResponseWriter, r *http.Request) {
		assetPath := "./client/" + r.URL.Path[len("/scripts/"):]
		http.ServeFile(w, r, assetPath)
	})

	http.HandleFunc("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		// Extract the requested asset path from the URL
		assetPath := "./assets/" + r.URL.Path[len("/assets/"):]
		http.ServeFile(w, r, assetPath)
	})

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		// Extract the requested asset path from the URL
		id := uuid.New().String()

		lobby[id] =
			game.CreateUser(r.FormValue("username"), id)

		http.Redirect(w, r, fmt.Sprintf("/lobby?id=%s", id), http.StatusSeeOther)
	})

	http.HandleFunc("/lobby", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/lobby.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			return
		}
	})

	// websocket stuff
	m.HandleConnect(func(newPlayerSession *melody.Session) {

		if started {
			fmt.Println("Whoa there more players are not allowed")
			return
		}

		allSessions, _ := m.Sessions()
		// get current player
		currentUserId := newPlayerSession.Request.URL.Query().Get("userId")
		currentPlayer := lobby[currentUserId]
		currentPlayer.Type = "join"
		var spriteId string
		availablePlayers, spriteId = popItem(availablePlayers)
		currentPlayer.SpriteType = spriteId

		// tell new player 	about current players
		for _, otherPlayerSession := range allSessions {
			pInfo, exists := otherPlayerSession.Get("info")

			if !exists {
				fmt.Println("Player does not exist")
				continue
			}

			otherPlayer := pInfo.(game.Player)

			// tell current player about other player
			err := newPlayerSession.Write(otherPlayer.ToJson())
			if err != nil {
				log.Fatal("Failed to send player info" + err.Error())
				return
			}
		}

		// store session info
		newPlayerSession.Set("info", currentPlayer)

		// tell other players about joined player
		err := m.BroadcastOthers(currentPlayer.ToJson(), newPlayerSession)

		if err != nil {
			log.Fatal("Failed to send player info" + err.Error())
			return
		}

		err = newPlayerSession.Write(currentPlayer.ToJson())

		if err != nil {
			log.Fatal("Failed to send player info" + err.Error())
			return
		}
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("Player exiting")
		value, exists := s.Get("info")

		if !exists {
			return
		}

		info := value.(game.Player)
		availablePlayers = append(availablePlayers, info.SpriteType)
		info.Type = "dis"

		err := m.BroadcastOthers(info.ToJson(), s)
		if err != nil {
			return
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		err := m.BroadcastOthers(msg, s)
		if err != nil {
			log.Fatal("Failed to send data" + err.Error())
			return
		}
	})

	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
		fmt.Println(port)
	} else {
		port = "5000" // Default port if not specified
	}
	fmt.Println("Server started at " + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}

}
