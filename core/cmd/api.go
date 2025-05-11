package cmd

import (
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"embed"
	"fmt"
	authrpc "github.com/RA341/multipacman/generated/auth/v1/v1connect"
	lobbyrpc "github.com/RA341/multipacman/generated/lobby/v1/v1connect"
	"github.com/RA341/multipacman/internal/config"
	"github.com/RA341/multipacman/internal/database"
	"github.com/RA341/multipacman/internal/game"
	"github.com/RA341/multipacman/internal/lobby"
	"github.com/RA341/multipacman/internal/user"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"io/fs"
	"net/http"
	"strconv"
)

//go:embed web
var frontendDir embed.FS

func StartServer() {
	baseUrl := fmt.Sprintf(":%s", strconv.Itoa(config.Opts.ServerPort))
	if err := setupServer(baseUrl); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func setupServer(baseUrl string) error {
	authSrv, lobSrv := initServices()

	router := http.NewServeMux()
	registerHandlers(router, authSrv, lobSrv)
	registerFrontend(router)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowPrivateNetwork: true,
		AllowedMethods:      connectcors.AllowedMethods(),
		AllowedHeaders:      append(connectcors.AllowedHeaders(), user.AuthHeader),
		ExposedHeaders:      connectcors.ExposedHeaders(),
	})

	log.Info().Str("Listening on:", baseUrl).Msg("")
	// Use h2c to serve HTTP/2 without TLS
	return http.ListenAndServe(
		baseUrl,
		corsHandler.Handler(h2c.NewHandler(router, &http2.Server{
			IdleTimeout: 0, // disable max timeout
		})),
	)
}

func initServices() (*user.Service, *lobby.Service) {
	config.Load()
	db := database.InitDB()

	// setup service structs
	authService := &user.Service{Db: db}
	lobSrv := lobby.NewLobbyService(db)

	return authService, lobSrv
}

func registerHandlers(mux *http.ServeMux, as *user.Service, ls *lobby.Service) {
	authInterceptor := connect.WithInterceptors(user.NewInterceptor(as))

	services := []func() (string, http.Handler){
		func() (string, http.Handler) {
			return authrpc.NewAuthServiceHandler(user.NewAuthHandler(as))
		},
		func() (string, http.Handler) {
			return lobbyrpc.NewLobbyServiceHandler(lobby.NewLobbyHandler(ls), authInterceptor)
		},
	}

	for _, svc := range services {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	game.RegisterGameWSHandler(mux, as, ls)
}

func registerFrontend(router *http.ServeMux) {
	subFS, err := fs.Sub(frontendDir, "web")
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to load frontend directory")
	}

	// serve frontend dir
	log.Info().Msgf("Setting up ui files")
	router.Handle("/", http.FileServer(http.FS(subFS)))
}
