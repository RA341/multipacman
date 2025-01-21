package api

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/RA341/multipacman/service"
	"github.com/rs/zerolog/log"
	"net/http"
)

const AuthHeader = "Authorization"

type authInterceptor struct {
	authService *service.AuthService
}

func (i *authInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		token := conn.RequestHeader().Get(AuthHeader)
		if token == "" {
			return connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("no auth header found"))
		}

		ctx, _, err := verifyAuthHeader(ctx, i.authService, token)
		if err != nil {
			return err
		}

		return next(ctx, conn)
	}
}

func (i *authInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		clientToken := req.Header().Get(AuthHeader)

		ctx, response, err := verifyAuthHeader(ctx, i.authService, clientToken)
		if err != nil {
			return response, err
		}

		return next(ctx, req)
	}
}

func (*authInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		return next(ctx, spec)
	}
}

func verifyAuthHeader(ctx context.Context, authService *service.AuthService, clientToken string) (context.Context, connect.AnyResponse, error) {
	user, err := authService.VerifyToken(clientToken)
	if err != nil {
		return nil, nil, connect.NewError(
			connect.CodeUnauthenticated,
			fmt.Errorf("invalid token %v", err),
		)
	}
	// add user value to subsequent requests
	ctx = context.WithValue(ctx, "user", user)
	return ctx, nil, nil
}

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
