package game

import (
	"context"
	"github.com/RA341/multipacman/internal/user"
	"github.com/rs/zerolog/log"
	"net/http"
)

func WsAuthMiddleware(authService *user.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := getAuthToken(r)
		if clientToken == "" {
			log.Error().Msg("empty client token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		username, err := authService.VerifyToken(clientToken)
		if err != nil {
			log.Error().Err(err).Msg("Error verifying token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), user.CtxUserKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAuthToken(r *http.Request) string {
	cookie, err := r.Cookie(user.AuthHeader)
	if err == nil && cookie.Value != "" {
		return cookie.Value
	}
	log.Warn().Err(err).Msg("unable to find auth cookie")

	clientToken := r.Header.Get("Sec-Websocket-Protocol")
	if clientToken != "" {
		return clientToken
	}
	log.Warn().Msg("No auth token in websocket header found")

	log.Error().Msg("No auth token found in cookie or Sec-Websocket-Protocol header, unauthorized")
	return ""
}
