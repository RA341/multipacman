package api

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/models"
	"github.com/RA341/multipacman/service"
	"github.com/rs/zerolog/log"
	"sync"
)

type LobbyHandler struct {
	mu           *sync.RWMutex
	connections  map[uint]chan bool
	lobbyService *service.LobbyService
}

func IniLobbyHandler(ls *service.LobbyService) *LobbyHandler {
	return &LobbyHandler{
		mu:           &sync.RWMutex{},
		connections:  map[uint]chan bool{},
		lobbyService: ls,
	}
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

	channel := l.newUpdateChannel(user.ID)

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

				l.removeUpdateChannel(user.ID)
				break
			}

			log.Error().Any("All connections", l.connections).Err(err).Msg("Error sending message to client")
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

func (l LobbyHandler) newUpdateChannel(channelId uint) chan bool {
	channel := make(chan bool)

	l.mu.Lock()
	l.connections[channelId] = channel
	l.mu.Unlock()

	log.Debug().Msg("Added to lobby list")
	return channel
}

func (l LobbyHandler) removeUpdateChannel(channelIndex uint) {
	l.mu.Lock()
	delete(l.connections, channelIndex)
	l.mu.Unlock()
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

	l.updateLobbies()

	return connect.NewResponse(&v1.AddLobbiesResponse{}), nil
}

func (l LobbyHandler) DeleteLobby(ctx context.Context, req *connect.Request[v1.DelLobbiesRequest]) (*connect.Response[v1.DelLobbiesResponse], error) {
	lobbyName, user := req.Msg.GetLobby(), ctx.Value("user").(*models.User)

	err := l.lobbyService.DeleteLobby(lobbyName.ID, user.ID)
	if err != nil {
		return nil, err
	}

	l.updateLobbies()

	return connect.NewResponse(&v1.DelLobbiesResponse{}), nil
}

func (l LobbyHandler) updateLobbies() {
	for _, chn := range l.connections {
		chn <- true
	}
}
