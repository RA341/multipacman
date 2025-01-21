package game

import (
	"context"
	"github.com/RA341/multipacman/service"
	"github.com/rs/zerolog/log"
	"net/http"
)

func AuthMiddleware(authService *service.AuthService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("Sec-Websocket-Protocol")
		if clientToken == "" {
			log.Error().Msg("No auth cookie found, while connecting to websocket")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		username, err := authService.VerifyToken(clientToken)
		if err != nil {
			log.Error().Err(err).Msg("Error verifying token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
