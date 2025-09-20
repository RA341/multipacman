package user

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
)

const AuthHeaderKey = "Authorization"
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
		ctx, err := i.verifyUser(ctx, conn.RequestHeader())
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
		ctx, err := i.verifyUser(ctx, req.Header())
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

func (i *Interceptor) verifyUser(ctx context.Context, headers http.Header) (context.Context, error) {
	user, err := i.authService.VerifyAuthHeader(headers)
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

func UserDataFromContext(ctx context.Context) (*User, error) {
	userVal := ctx.Value(CtxUserKey)
	if userVal == nil {
		return nil, fmt.Errorf("could not find user in context")
	}

	return userVal.(*User), nil
}
