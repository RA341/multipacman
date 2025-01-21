package api

import (
	"connectrpc.com/connect"
	auth "github.com/RA341/multipacman/generated/auth/v1/v1connect"
	lobby "github.com/RA341/multipacman/generated/lobby/v1/v1connect"
	"github.com/RA341/multipacman/service"
	"net/http"
)

func InitHandlers(as *service.AuthService, ls *service.LobbyService) *http.ServeMux {
	mux := http.NewServeMux()
	lobbyHandler := InitLobbyHandler(ls)

	authInterceptor := connect.WithInterceptors(&authInterceptor{authService: as})

	services := []func() (string, http.Handler){
		// auth
		func() (string, http.Handler) {
			return auth.NewAuthServiceHandler(InitAuthHandler(as))
		},
		// lobbies
		func() (string, http.Handler) {
			return lobby.NewLobbyServiceHandler(lobbyHandler, authInterceptor)
		},
	}

	for _, svc := range services {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}
