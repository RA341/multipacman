package game

import (
	"context"
	"github.com/RA341/multipacman/internal/user"
	"github.com/rs/zerolog/log"
	"net/http"
)

func WSAuthMiddleware(authService *user.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		clientToken := getWSAuthToken(req)
		if clientToken == "" {
			log.Error().Msg("empty client token")
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		username, err := authService.VerifyToken(clientToken)
		if err != nil {
			log.Error().Err(err).Msg("Error verifying token")
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), user.CtxUserKey, username)
		next.ServeHTTP(writer, req.WithContext(ctx))
	})
}

func getWSAuthToken(r *http.Request) string {
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
