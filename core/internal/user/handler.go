package user

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	v1 "github.com/RA341/multipacman/generated/auth/v1"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	auth *Service
}

func NewAuthHandler(auth *Service) *Handler {
	return &Handler{auth: auth}
}

func (a *Handler) Register(_ context.Context, c *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	username, password, passwordVerify := c.Msg.Username, c.Msg.Password, c.Msg.PasswordVerify

	if username == "" || password == "" || passwordVerify == "" {
		log.Warn().Any("Msg", c.Msg).Msg("one or more fields are empty")
		return nil, fmt.Errorf("empty fields")
	}

	// Ensure that the password & passwordVerify match
	if password != passwordVerify {
		return nil, fmt.Errorf("password mismatch")
	}

	err := a.auth.Register(c.Msg.Username, c.Msg.Password, false)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.RegisterUserResponse{}), nil
}

func (a *Handler) Logout(ctx context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error) {
	clientToken := req.Header().Get(AuthHeader)
	ctx, err := verifyAuthHeader(ctx, a.auth, clientToken)
	if err != nil {
		log.Warn().Err(err).Msg("Logout failed, user info not found in request")
		return connect.NewResponse(&v1.Empty{}), nil
	}
	user, err := GetUserContext(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Logout failed, user info not found in context")
		return connect.NewResponse(&v1.Empty{}), nil
	}

	_, err = a.auth.Logout(user.ID)
	if err != nil {
		log.Warn().Err(err).Msg("Logout failed, error occurred while updating db")
	}

	return connect.NewResponse(&v1.Empty{}), nil
}

func (a *Handler) GuestLogin(_ context.Context, _ *connect.Request[v1.Empty]) (*connect.Response[v1.UserResponse], error) {
	username := randomdata.SillyName()
	password := randomdata.Alphanumeric(30)
	err := a.auth.Register(username, password, true)
	if err != nil {
		return nil, err
	}

	user, err := a.auth.Login(username, password)
	if err != nil {
		return nil, err
	}

	response := connect.NewResponse(user.ToRPC())
	setCookie(user, response)

	return response, nil
}

func (a *Handler) Login(_ context.Context, c *connect.Request[v1.AuthRequest]) (*connect.Response[v1.UserResponse], error) {
	username, password := c.Msg.Username, c.Msg.Password
	if username != c.Msg.Username || password != c.Msg.Password {
		return nil, fmt.Errorf("empty username or password")
	}

	user, err := a.auth.Login(username, password)
	if err != nil {
		return nil, err
	}

	response := connect.NewResponse(user.ToRPC())
	setCookie(user, response)

	return response, nil
}

func (a *Handler) Test(_ context.Context, c *connect.Request[v1.AuthResponse]) (*connect.Response[v1.UserResponse], error) {
	clientToken := c.Msg.GetAuthToken()

	user, err := a.auth.VerifyToken(clientToken)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(user.ToRPC()), nil
}
