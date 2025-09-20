package game

import (
	"context"
	"net/http"

	"github.com/RA341/multipacman/internal/user"
	"github.com/rs/zerolog/log"
)

func WSAuthMiddleware(authService *user.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		userData, err := authService.VerifyAuthHeader(req.Header)
		if err != nil {
			log.Error().Err(err).Msg("Error verifying token")
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), user.CtxUserKey, userData)
		next.ServeHTTP(writer, req.WithContext(ctx))
	})
}
