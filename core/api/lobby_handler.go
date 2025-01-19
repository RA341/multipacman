package api

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/models"
	"github.com/RA341/multipacman/service"
	"github.com/rs/zerolog/log"
)

type LobbyHandler struct {
	lobbyService *service.LobbyService
	// indicates if the channel should send updates to connected clients
	lobbyChannel chan bool
}

func (l LobbyHandler) ListLobbies(_ context.Context, _ *connect.Request[v1.ListLobbiesRequest], stream *connect.ServerStream[v1.ListLobbiesResponse]) error {
	lobbies, err := l.lobbyService.RetrieveLobbies()
	if err != nil {
		return err // Return the error directly
	}

	// Send the initial lobbies list
	err = stream.Send(&v1.ListLobbiesResponse{
		Lobbies: lobbies,
	})
	if err != nil {
		return err
	}

	for msg := range l.lobbyChannel {
		if msg == true {
			lobbies, err := l.lobbyService.RetrieveLobbies()
			if err != nil {
				log.Warn().Err(err).Msgf("retreiving lobby list failed")
			}

			err = stream.Send(&v1.ListLobbiesResponse{Lobbies: lobbies})
			if err != nil {
				log.Warn().Err(err).Msgf("error sendign grpc stream")
			}
		}
	}

	return nil
}

func (l LobbyHandler) AddLobby(ctx context.Context, req *connect.Request[v1.AddLobbiesRequest]) (*connect.Response[v1.AddLobbiesResponse], error) {
	lobbyName := req.Msg.GetLobbyName()
	ctxValue := ctx.Value("user")
	if ctxValue == nil {
		log.Error().Msg("User not found in context")
		return nil, fmt.Errorf("internal server error")
	}
	user := ctxValue.(*models.User)

	err := l.lobbyService.CreateLobby(lobbyName, user.Username, user.ID)
	if err != nil {
		return nil, err
	}

	// indicate update
	l.lobbyChannel <- true

	return connect.NewResponse(&v1.AddLobbiesResponse{}), nil
}

func (l LobbyHandler) DeleteLobby(ctx context.Context, req *connect.Request[v1.DelLobbiesRequest]) (*connect.Response[v1.DelLobbiesResponse], error) {
	lobbyName, user := req.Msg.GetLobby(), ctx.Value("user").(*models.User)

	err := l.lobbyService.DeleteLobby(lobbyName.ID, user.ID)
	if err != nil {
		return nil, err
	}

	// indicate update
	l.lobbyChannel <- true

	return connect.NewResponse(&v1.DelLobbiesResponse{}), nil
}
