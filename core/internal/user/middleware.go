package user

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
)

const AuthHeader = "Authorization"
const CtxUserKey = "user"

type Interceptor struct {
	authService *Service
}

func NewInterceptor(authService *Service) *Interceptor {
	return &Interceptor{authService: authService}
}

func (i *Interceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		token := conn.RequestHeader().Get(AuthHeader)
		if token == "" {
			return connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("no auth header found"))
		}

		ctx, err := verifyAuthHeader(ctx, i.authService, token)
		if err != nil {
			return err
		}

		return next(ctx, conn)
	}
}

func (i *Interceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		clientToken := req.Header().Get(AuthHeader)

		ctx, err := verifyAuthHeader(ctx, i.authService, clientToken)
		if err != nil {
			return nil, err
		}

		return next(ctx, req)
	}
}

func (*Interceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		return next(ctx, spec)
	}
}

func verifyAuthHeader(ctx context.Context, authService *Service, clientToken string) (context.Context, error) {
	user, err := authService.VerifyToken(clientToken)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeUnauthenticated,
			err,
		)
	}
	// add user value to subsequent requests
	ctx = context.WithValue(ctx, CtxUserKey, user)
	return ctx, nil
}

// todo combine the 2 funcs

func GetUserContext(ctx context.Context) (*User, error) {
	userVal := ctx.Value(CtxUserKey)
	if userVal == nil {
		return nil, fmt.Errorf("could not find user in context")
	}
	user, ok := userVal.(*User)
	if !ok {
		return nil, fmt.Errorf("invalid user type in context")
	}

	return user, nil
}

func GetUserFromCtx(ctx context.Context) *User {
	return ctx.Value(CtxUserKey).(*User)
}
