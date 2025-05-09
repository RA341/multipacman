package game

import (
	"context"
	"github.com/RA341/multipacman/internal/user"
	"github.com/rs/zerolog/log"
	"net/http"
)

func WsAuthMiddleware(authService *user.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := getAuthToken(w, r)
		if clientToken == "" {
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

func getAuthToken(w http.ResponseWriter, r *http.Request) string {
	clientToken := r.Header.Get("Sec-Websocket-Protocol")
	if clientToken == "" {
		log.Warn().Msg("No websocket header found, checking cookie")
		// check cookies
		cookie, err := r.Cookie("auth")
		if err != nil {
			log.Error().Err(err).Msg("unable to find auth cookie")
			w.WriteHeader(http.StatusBadRequest)
			return ""
		}

		if cookie.Value == "" {
			log.Error().Msg("No auth cookie found, while connecting to websocket")
			w.WriteHeader(http.StatusBadRequest)
			return ""
		}
		clientToken = cookie.Value
	}

	return clientToken
}
