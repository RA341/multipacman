package api

import (
	"connectrpc.com/connect"
	"context"
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/models"
	"github.com/RA341/multipacman/service"
)

type LobbyHandler struct {
	lobbyService *service.LobbyService
}

func (l LobbyHandler) ListLobbies(ctx context.Context, req *connect.Request[v1.ListLobbiesRequest]) (*connect.Response[v1.ListLobbiesResponse], error) {
	lobbies, err := l.lobbyService.RetrieveLobbies()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ListLobbiesResponse{Lobbies: lobbies}), nil
}

func (l LobbyHandler) AddLobby(ctx context.Context, req *connect.Request[v1.AddLobbiesRequest]) (*connect.Response[v1.AddLobbiesResponse], error) {
	lobbyName, user := req.Msg.GetLobbyName(), ctx.Value("user").(*models.User)

	err := l.lobbyService.CreateLobby(lobbyName, user.ID)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.AddLobbiesResponse{}), nil
}

func (l LobbyHandler) DeleteLobby(ctx context.Context, req *connect.Request[v1.DelLobbiesRequest]) (*connect.Response[v1.DelLobbiesResponse], error) {
	lobbyName, user := req.Msg.GetLobby(), ctx.Value("user").(*models.User)

	err := l.lobbyService.DeleteLobby(lobbyName.ID, user.ID)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DelLobbiesResponse{}), nil
}
