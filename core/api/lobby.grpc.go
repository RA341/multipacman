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
}

func InitLobbyHandler(ls *service.LobbyService) *LobbyHandler {
	return &LobbyHandler{ls}
}

func (l LobbyHandler) ListLobbies(ctx context.Context, _ *connect.Request[v1.ListLobbiesRequest], stream *connect.ServerStream[v1.ListLobbiesResponse]) error {
	user, err := getUserContext(ctx)
	if err != nil {
		return err
	}

	lobbies, err := l.lobbyService.GetGrpcLobbies()
	if err != nil {
		return err // Return the error directly
	}

	// Send the initial lobbies list
	err = stream.Send(&v1.ListLobbiesResponse{Lobbies: lobbies})
	if err != nil {
		return err
	}

	channel := l.lobbyService.NewUpdateChannel(user.ID)

	for range channel {
		//log.Debug().Bool("Msg", msg).Msg("Received update message")

		lobbies, err := l.lobbyService.GetGrpcLobbies()
		if err != nil {
			log.Error().Err(err).Msg("Error getting lobbies")
			continue
		}

		err = stream.Send(&v1.ListLobbiesResponse{Lobbies: lobbies})
		if err != nil {
			if err.Error() == "canceled: client disconnected" {
				log.Debug().
					Uint("user id", user.ID).
					Str("username", user.Username).
					Msg("Client disconnected, removing channel at ind")

				l.lobbyService.RemoveUpdateChannel(user.ID)
				break
			}

			log.Error().Any("All connections", l.lobbyService.Connections).Err(err).Msg("Error sending message to client")
		}
	}

	return nil
}

func getUserContext(ctx context.Context) (*models.User, error) {
	userVal := ctx.Value("user")
	if userVal == nil {
		return nil, fmt.Errorf("could not find user in context")
	}
	user, ok := userVal.(*models.User)
	if !ok {
		return nil, fmt.Errorf("invalid user type in context")
	}

	return user, nil
}

func (l LobbyHandler) AddLobby(ctx context.Context, req *connect.Request[v1.AddLobbiesRequest]) (*connect.Response[v1.AddLobbiesResponse], error) {
	user, err := getUserContext(ctx)
	if err != nil {
		return nil, err
	}

	lobbyName := req.Msg.GetLobbyName()
	err = l.lobbyService.CreateLobby(lobbyName, user.Username, user.ID)
	if err != nil {
		return nil, err
	}

	l.lobbyService.UpdateLobbies()

	return connect.NewResponse(&v1.AddLobbiesResponse{}), nil
}

func (l LobbyHandler) DeleteLobby(ctx context.Context, req *connect.Request[v1.DelLobbiesRequest]) (*connect.Response[v1.DelLobbiesResponse], error) {
	lobbyName, user := req.Msg.GetLobby(), ctx.Value("user").(*models.User)

	err := l.lobbyService.DeleteLobby(lobbyName.ID, user.ID)
	if err != nil {
		return nil, err
	}

	l.lobbyService.UpdateLobbies()

	return connect.NewResponse(&v1.DelLobbiesResponse{}), nil
}
