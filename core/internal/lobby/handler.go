package lobby

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/internal/user"
)

type Handler struct {
	lobbyService *Service
}

func NewLobbyHandler(ls *Service) *Handler {
	return &Handler{ls}
}

func (l Handler) ListLobbies(context.Context, *connect.Request[v1.ListLobbiesRequest]) (*connect.Response[v1.ListLobbiesResponse], error) {
	lobbies, err := l.lobbyService.GetGrpcLobbies()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ListLobbiesResponse{Lobbies: lobbies}), nil
}

func (l Handler) AddLobby(ctx context.Context, req *connect.Request[v1.AddLobbiesRequest]) (*connect.Response[v1.AddLobbiesResponse], error) {
	userInfo, err := user.UserDataFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if userInfo.Guest {
		return nil, fmt.Errorf("guest users cannot create lobbies")
	}

	lobbyName := req.Msg.GetLobbyName()
	err = l.lobbyService.CreateLobby(lobbyName, userInfo.Username, userInfo.ID)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.AddLobbiesResponse{}), nil
}

func (l Handler) DeleteLobby(ctx context.Context, req *connect.Request[v1.DelLobbiesRequest]) (*connect.Response[v1.DelLobbiesResponse], error) {
	lobbyName := req.Msg.GetLobby()
	userInfo, err := user.UserDataFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = l.lobbyService.DeleteLobby(lobbyName.ID, userInfo.ID)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DelLobbiesResponse{}), nil
}
