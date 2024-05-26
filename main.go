package main

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/olahol/melody"
	"log"
	"net/http"
	entities "server/entities"
	"time"
)

//go:embed client/*
var embeddedFs embed.FS

var LobbyList map[string]*entities.LobbyModel

func main() {
	router := chi.NewRouter()
	m := melody.New()
	// init variable
	LobbyList = make(map[string]*entities.LobbyModel)

	// middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(AuthMiddleware)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/lobby", http.StatusFound)
	})

	router.Get("/full", func(writer http.ResponseWriter, request *http.Request) {
		log.Print(writer.Write([]byte("Lobby is full, choose a new lobby")))
	})

	// ws handler
	router.Get("/ws/*", func(writer http.ResponseWriter, request *http.Request) {
		err := m.HandleRequest(writer, request)
		if err != nil {
			log.Printf("Something went wrong with ws handler")
			log.Printf(err.Error())
			return
		}
	})
	initMelody(m)

	router.Get("/static/*", func(writer http.ResponseWriter, request *http.Request) {
		assetPath := "client/" + request.URL.Path[len("/static/"):]

		fileContents, err := embeddedFs.ReadFile(assetPath)
		if err != nil {
			log.Printf("Failed to read" + assetPath)
		}

		contentType := DetectContentType(assetPath, fileContents)
		//log.Printf("Filetype for " + assetPath + " is " + contentType)

		writer.Header().Add("Content-Type", contentType)

		_, err = writer.Write(fileContents)
		if err != nil {
			log.Printf("Failed to write" + assetPath)
			return
		}
	})

	router.Get("/login", func(writer http.ResponseWriter, request *http.Request) {
		filepath := "client/auth/login.html"
		handleHtmlPath(writer, request, filepath)
	})

	router.Get("/register", func(writer http.ResponseWriter, request *http.Request) {
		filepath := "client/auth/register.html"
		handleHtmlPath(writer, request, filepath)
	})

	router.Get("/lobby", func(writer http.ResponseWriter, request *http.Request) {
		user := request.Context().Value("user")
		if user == nil {
			log.Printf("User not found in context")
			http.Redirect(writer, request, "/login", http.StatusFound)
			return
		}
		filepath := "client/lobby/lobby_page.html"
		handleHtmlPath(writer, request, filepath)
	})

	router.Get("/game", func(writer http.ResponseWriter, request *http.Request) {
		// TODO add auth once game is complete
		userID := request.URL.Query().Get("user")

		if userID == "" {
			http.Error(writer, "No user id found in query params check the url", http.StatusNotFound)
			return
		}

		lobbyId := "asdsad"

		fileContents := replaceIds(userID, lobbyId)
		writer.Header().Add("Content-Type", "text/html")

		_, err := writer.Write(fileContents)
		if err != nil {
			log.Printf("Failed to write game.html")
			return
		}
	})

	// router for login
	port := getServerPort()

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return
	}
}

func AddToLobbyList(name string, lobby *entities.LobbyModel) {
	// Add a new lobby to the lobbyList map
	LobbyList[name] = lobby
}
