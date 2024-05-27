package main

import (
	"database/sql"
	"embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olahol/melody"
	"log"
	"net/http"
	"server/api"
	"server/api/auth"
	entities "server/entities"
	"time"
)

//go:embed client/*
var embeddedFs embed.FS

var LobbyList map[string]*entities.LobbyModel

func main() {
	db, _ := sql.Open("sqlite3", "./multipacman.db")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error while closing database connection")
		}
	}(db)

	api.SetupDatabase(db, false)

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
	router.Use(func(handler http.Handler) http.Handler {
		return AuthMiddleware(db, handler)
	})

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

	// api routes
	router.Mount("/api/auth", auth.SetupAuthRouter(db))
	//router.Mount("/api/lobby", auth.SetupAuthRouter())

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
		if user == nil || user == "" {
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

		for lobbyKey := range LobbyList {
			lobby := LobbyList[lobbyKey]
			test := lobby.ConnectedPlayers[userID]
			if test != nil {
				// user tried to enter
				http.Redirect(writer, request, "/static/lobby/entered.html", http.StatusFound)
				return
			}
		}

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
