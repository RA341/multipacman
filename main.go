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
	lobby "server/api/lobby"
	user "server/api/user"
	entities "server/entities"
	"time"
)

//go:embed client/*
var embeddedFs embed.FS

var LobbyList map[int]*entities.LobbyModel

func main() {
	db, _ := sql.Open("sqlite3", "./multipacman.db")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error while closing database connection")
		}
	}(db)

	api.SetupDatabase(db, false)
	// routers and websocket
	router := chi.NewRouter()
	m := melody.New()
	// init lobby list
	ids := lobby.RetrieveLobbyIds(db)
	LobbyList = make(map[int]*entities.LobbyModel)
	for _, data := range ids {
		LobbyList[data] = entities.NewLobbyModel()
	}

	// middlewares
	initMiddlewares(router)
	// ws handler
	initWSHandler(router, db, m)
	// api routes
	router.Mount("/api/auth", AuthMiddleware(db, auth.SetupAuthRouter(db)))
	router.Mount("/api/user", AuthMiddleware(db, user.SetupUserRouter()))
	router.Mount("/api/lobby", AuthMiddleware(db, lobby.SetupLobbyRouter(db, LobbyList)))
	// main paths
	initMainPaths(router, db)

	// router for login
	port := getServerPort()

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return
	}
}

func initMainPaths(router *chi.Mux, db *sql.DB) {
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/lobby", http.StatusFound)
	})

	router.Get("/full", func(writer http.ResponseWriter, request *http.Request) {
		log.Print(writer.Write([]byte("Lobby is full, choose a new lobby")))
	})

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

	router.Route("/lobby", func(r chi.Router) {
		r.Use(func(handler http.Handler) http.Handler {
			return AuthMiddleware(db, handler)
		})

		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			username := request.Context().Value("user")
			if username == nil || username == "" {
				log.Printf("User not found in context")
				http.Redirect(writer, request, "/login", http.StatusFound)
				return
			}
			filepath := "client/lobby/lobby_page.html"
			handleHtmlPath(writer, request, filepath)
		})
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
			l := LobbyList[lobbyKey]
			test := l.ConnectedPlayers[userID]
			if test != nil {
				// user tried to enter
				http.Redirect(writer, request, "/static/l/entered.html", http.StatusFound)
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
}

func initWSHandler(router *chi.Mux, db *sql.DB, m *melody.Melody) {
	router.Route("/ws", func(r chi.Router) {
		r.Use(func(handler http.Handler) http.Handler {
			return AuthMiddleware(db, handler)
		})

		r.Get("/*", func(writer http.ResponseWriter, request *http.Request) {
			username := request.Context().Value("user")
			if username == nil || username == "" {
				log.Printf("User not found in context")
				http.Redirect(writer, request, "/login", http.StatusFound)
				return
			}
			err := m.HandleRequest(writer, request)
			if err != nil {
				log.Printf("Something went wrong with ws handler")
				log.Printf(err.Error())
				return
			}
		})
	})

	initMelody(m)
}

func initMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
}
