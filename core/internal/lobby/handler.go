package lobby

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/internal/auth"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	lobbyService *Service
}

func NewLobbyHandler(ls *Service) *Handler {
	return &Handler{ls}
}

func (l Handler) ListLobbies(ctx context.Context, _ *connect.Request[v1.ListLobbiesRequest], stream *connect.ServerStream[v1.ListLobbiesResponse]) error {
	user, err := auth.GetUserContext(ctx)
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

func (l Handler) AddLobby(ctx context.Context, req *connect.Request[v1.AddLobbiesRequest]) (*connect.Response[v1.AddLobbiesResponse], error) {
	user, err := auth.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	if user.Guest {
		return nil, fmt.Errorf("guest users cannot create lobbies")
	}

	lobbyName := req.Msg.GetLobbyName()
	err = l.lobbyService.CreateLobby(lobbyName, user.Username, user.ID)
	if err != nil {
		return nil, err
	}

	l.lobbyService.UpdateLobbies()

	return connect.NewResponse(&v1.AddLobbiesResponse{}), nil
}

func (l Handler) DeleteLobby(ctx context.Context, req *connect.Request[v1.DelLobbiesRequest]) (*connect.Response[v1.DelLobbiesResponse], error) {
	lobbyName, user := req.Msg.GetLobby(), ctx.Value("user").(*auth.User)

	err := l.lobbyService.DeleteLobby(lobbyName.ID, user.ID)
	if err != nil {
		return nil, err
	}

	l.lobbyService.UpdateLobbies()

	return connect.NewResponse(&v1.DelLobbiesResponse{}), nil
}
