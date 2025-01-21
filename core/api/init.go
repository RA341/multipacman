package api

import (
	"connectrpc.com/connect"
	auth "github.com/RA341/multipacman/generated/auth/v1/v1connect"
	lobby "github.com/RA341/multipacman/generated/lobby/v1/v1connect"
	"github.com/RA341/multipacman/service"
	"gorm.io/gorm"
	"net/http"
)

func InitHandlers(database *gorm.DB) *http.ServeMux {
	// setup service structs
	authService := &service.AuthService{Db: database}
	lobbyService := &service.LobbyService{Db: database}

	mux := http.NewServeMux()

	authInterceptor := connect.WithInterceptors(&authInterceptor{authService: authService})

	services := []func() (string, http.Handler){
		// auth
		func() (string, http.Handler) {
			return auth.NewAuthServiceHandler(&AuthHandler{auth: authService})
		},
		// lobbies
		func() (string, http.Handler) {
			return lobby.NewLobbyServiceHandler(IniLobbyHandler(lobbyService), authInterceptor)
		},
	}

	for _, svc := range services {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}
