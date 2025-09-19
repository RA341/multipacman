package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
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
)

func StartServer(frontendPath string) {
	baseUrl := fmt.Sprintf(":%s", strconv.Itoa(config.Opts.ServerPort))
	if err := setupServer(baseUrl, frontendPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func setupServer(baseUrl, frontendPath string) error {
	authSrv, lobSrv := initServices()

	router := http.NewServeMux()
	registerHandlers(router, authSrv, lobSrv)
	registerFrontend(router, frontendPath)

	cor := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowPrivateNetwork: true,
		AllowedMethods:      connectcors.AllowedMethods(),
		AllowedHeaders:      append(connectcors.AllowedHeaders(), user.AuthHeader),
		ExposedHeaders:      connectcors.ExposedHeaders(),
	})

	log.Info().Str("port", baseUrl).Msg("listening on:")
	return http.ListenAndServe(
		baseUrl,
		cor.Handler(h2c.NewHandler(router,
			&http2.Server{},
		)),
	)
}

func initServices() (*user.Service, *lobby.Service) {
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

func registerFrontend(router *http.ServeMux, frontEndPath string) {
	if frontEndPath == "" {
		log.Warn().Msg("Empty frontend frontend path, no UI will be served")
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte("No Ui configured")); err != nil {
				log.Warn().Err(err).Msg("Failed to write response")
				return
			}
		})
		return
	}

	root, err := os.OpenRoot(frontEndPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to open frontend dir")
	}
	defer func(root *os.Root) {
		err := root.Close()
		if err != nil {
			log.Warn().Err(err).Msg("Failed to close frontend dir")
		}
	}(root)

	router.Handle("/", http.FileServer(http.FS(root.FS())))

}
