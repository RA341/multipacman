package user

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
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

func (a *Handler) Logout(_ context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error) {
	user, err := a.auth.VerifyAuthHeader(req.Header())
	if err != nil {
		log.Warn().Err(err).Msg("Logout failed, unauthenticated user")
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

	return a.loginWithCookie(username, password)
}

func (a *Handler) Login(_ context.Context, c *connect.Request[v1.AuthRequest]) (*connect.Response[v1.UserResponse], error) {
	username, password := c.Msg.Username, c.Msg.Password
	if username != c.Msg.Username || password != c.Msg.Password {
		return nil, fmt.Errorf("empty username or password")
	}

	return a.loginWithCookie(username, password)
}

func (a *Handler) loginWithCookie(user, pass string) (*connect.Response[v1.UserResponse], error) {
	userData, err := a.auth.Login(user, pass)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	response := connect.NewResponse(userData.ToRPC())
	setCookie(userData, response)

	return response, nil
}

func (a *Handler) Test(_ context.Context, c *connect.Request[v1.AuthResponse]) (*connect.Response[v1.UserResponse], error) {
	user, err := a.auth.VerifyAuthHeader(c.Header())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(user.ToRPC()), nil
}
