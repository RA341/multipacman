package api

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
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
	authInterceptor := connect.WithInterceptors(NewAuthInterceptor(authService))

	services := []func() (string, http.Handler){
		// auth
		func() (string, http.Handler) {
			return auth.NewAuthServiceHandler(&AuthHandler{auth: authService})
		},
		// lobbies
		func() (string, http.Handler) {
			return lobby.NewLobbyServiceHandler(&LobbyHandler{lobbyService: lobbyService}, authInterceptor)
		},
	}

	for _, svc := range services {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}

func NewAuthInterceptor(authService *service.AuthService) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			clientToken := req.Header().Get("Authorization")
			user, err := authService.VerifyToken(clientToken)

			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					fmt.Errorf("invalid token %v", err),
				)
			}

			// add user value to subsequent requests
			ctx = context.WithValue(ctx, "user", &user)

			return next(ctx, req)
		}
	}
}
