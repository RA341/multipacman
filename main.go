package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed client/*
var embeddedFs embed.FS

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/login", http.StatusFound)
	})

	router.Get("/static/*", func(writer http.ResponseWriter, request *http.Request) {
		assetPath := "client/" + request.URL.Path[len("/static/"):]

		fileContents, err := embeddedFs.ReadFile(assetPath)
		if err != nil {
			log.Printf("Failed to read" + assetPath)
		}

		contentType := DetectContentType(assetPath, fileContents)
		log.Printf("Filetype for " + assetPath + " is " + contentType)

		writer.Header().Add("Content-Type", contentType)

		_, err = writer.Write(fileContents)
		if err != nil {
			log.Printf("Failed to write" + assetPath)
			return
		}
	})

	router.Get("/login", func(writer http.ResponseWriter, request *http.Request) {
		filepath := "./client/auth/login.html"
		handleHtmlPath(writer, request, filepath)
	})
	router.Get("/register", func(writer http.ResponseWriter, request *http.Request) {
		filepath := "./client/auth/register.html"
		handleHtmlPath(writer, request, filepath)
	})
	router.Get("/lobby", func(writer http.ResponseWriter, request *http.Request) {
		filepath := "./client/lobby/lobby_page.html"
		handleHtmlPath(writer, request, filepath)
	})

	// router for login
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
		fmt.Println(port)
	} else {
		port = "5000" // Default port if not specified
	}
	fmt.Println("Server started at " + port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return
	}
}

func handleHtmlPath(writer http.ResponseWriter, request *http.Request, filepath string) {
	if !FileExists(filepath) {
		http.NotFound(writer, request)
		return
	}
	http.ServeFile(writer, request, filepath)
}

func loginRouter() http.Handler {
	r := chi.NewRouter()
	return r
}

func loginPage(w http.ResponseWriter, r *http.Request) {

}

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		log.Printf("File does not exist")
		return false
	}
	return true
}

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

//http.HandleFunc("/scripts/*", func(w http.ResponseWriter, r *http.Request) {
//	assetPath := "./ref/" + r.URL.Path[len("/scripts/"):]
//	http.ServeFile(w, r, assetPath)
//})
//
//http.HandleFunc("/assets/*", func(w http.ResponseWriter, r *http.Request) {
//	// Extract the requested asset path from the URL
//	assetPath := "./assets/" + r.URL.Path[len("/assets/"):]
//	http.ServeFile(w, r, assetPath)
//})
//
//http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
//	// Extract the requested asset path from the URL
//	id := uuid.New().String()
//
//	lobby[id] =
//		game.CreateUser(r.FormValue("username"), id)
//
//	http.Redirect(w, r, fmt.Sprintf("/lobby?id=%s", id), http.StatusSeeOther)
//})
//
//http.HandleFunc("/lobby", func(w http.ResponseWriter, r *http.Request) {
//	http.ServeFile(w, r, "./ref/lobby.html")
//})
//
//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
//	err := m.HandleRequest(w, r)
//	if err != nil {
//		return
//	}
//})
//
//// websocket stuff
//m.HandleConnect(func(newPlayerSession *melody.Session) {
//
//	if started {
//		fmt.Println("Whoa there more players are not allowed")
//		return
//	}
//
//	allSessions, _ := m.Sessions()
//	// get current player
//	currentUserId := newPlayerSession.Request.URL.Query().Get("userId")
//	currentPlayer := lobby[currentUserId]
//	currentPlayer.Type = "join"
//	var spriteId string
//	availablePlayers, spriteId = popItem(availablePlayers)
//	currentPlayer.SpriteType = spriteId
//
//	// tell new player 	about current players
//	for _, otherPlayerSession := range allSessions {
//		pInfo, exists := otherPlayerSession.Get("info")
//
//		if !exists {
//			fmt.Println("Player does not exist")
//			continue
//		}
//
//		otherPlayer := pInfo.(game.Player)
//
//		// tell current player about other player
//		err := newPlayerSession.Write(otherPlayer.ToJson())
//		if err != nil {
//			log.Fatal("Failed to send player info" + err.Error())
//			return
//		}
//	}
//
//	// store session info
//	newPlayerSession.Set("info", currentPlayer)
//
//	// tell other players about joined player
//	err := m.BroadcastOthers(currentPlayer.ToJson(), newPlayerSession)
//
//	if err != nil {
//		log.Fatal("Failed to send player info" + err.Error())
//		return
//	}
//
//	err = newPlayerSession.Write(currentPlayer.ToJson())
//
//	if err != nil {
//		log.Fatal("Failed to send player info" + err.Error())
//		return
//	}
//})
//
//m.HandleDisconnect(func(s *melody.Session) {
//	fmt.Println("Player exiting")
//	value, exists := s.Get("info")
//
//	if !exists {
//		return
//	}
//
//	info := value.(game.Player)
//	availablePlayers = append(availablePlayers, info.SpriteType)
//	info.Type = "dis"
//
//	err := m.BroadcastOthers(info.ToJson(), s)
//	if err != nil {
//		return
//	}
//})
//
//m.HandleMessage(func(s *melody.Session, msg []byte) {
//	err := m.BroadcastOthers(msg, s)
//	if err != nil {
//		log.Fatal("Failed to send data" + err.Error())
//		return
//	}
//})
