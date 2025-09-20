package game

import (
	"encoding/json"
	"fmt"
	"github.com/RA341/multipacman/internal/lobby"
	"github.com/RA341/multipacman/internal/user"
	"github.com/RA341/multipacman/pkg"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Manager struct {
	activeLobbies pkg.Map[uint, *World]
	lobbyService  *lobby.Service
	mel           *melody.Melody
}

func (manager *Manager) getLobbyIdFromSession(s *melody.Session) uint {
	lobbyId, exists := s.Get(worldKey)
	if !exists {
		log.Warn().Msg("lobby id not found on disconnect")
		return 0
	}
	return lobbyId.(uint)
}

func (manager *Manager) broadcastAll(world *World, message []byte) error {
	broadCastSessions := world.ConnectedPlayers.GetValues()
	err := manager.mel.BroadcastMultiple(message, broadCastSessions)
	if err != nil {
		return err
	}
	return nil
}

func (manager *Manager) broadcastExceptPlayer(player *melody.Session, message []byte) error {
	err := manager.mel.BroadcastOthers(message, player)
	if err != nil {
		return err
	}
	return nil
}

func (manager *Manager) sendGameStateInfo(newPlayerSession *melody.Session, world *World) error {
	player, err := getPlayerEntityFromSession(newPlayerSession)
	if err != nil {
		return fmt.Errorf("unable to find player: %v", err)
	}

	gameState, err := world.GetGameStateReport(player.secretToken, player.Username, string(player.SpriteType), newPlayerSession)
	if err != nil {
		return fmt.Errorf("unable to marshal game state: %v", err)
	}

	err = newPlayerSession.Write(gameState)
	if err != nil {
		return fmt.Errorf("unable to send game state: %v", err)
	}

	return nil
}

func (manager *Manager) getWorld(lobby *lobby.Lobby) (*World, error) {
	activeWorld, exists := manager.activeLobbies.Load(lobby.ID)
	if !exists {
		log.Info().Msgf("creating new lobby")

		newWorld := NewWorldState()
		manager.activeLobbies.Store(lobby.ID, newWorld)
		go func() {
			reason := newWorld.waitForGameOver()
			// endgame does not need any info
			msg := EndGameMessage(reason).handler(MessageData{})
			marshal, err := json.Marshal(msg)
			if err != nil {
				log.Error().Err(err).Msg("Unable to marshal msg")
			} else {
				pkg.Elog(manager.broadcastAll(newWorld, marshal))
			}

			log.Debug().Uint("id", lobby.ID).Str("reason", reason).Msg("game end deleting lobby")
			manager.activeLobbies.Delete(lobby.ID)
		}()

		return newWorld, nil
	}

	if activeWorld.IsLobbyFull() {
		log.Warn().Any("lobby_state", activeWorld).Msgf("lobby is full, more players are not allowed")
		return nil, fmt.Errorf("lobby is full")
	}

	return activeWorld, nil
}

func (manager *Manager) getUserAndLobbyInfo(newPlayerSession *melody.Session) (*user.User, *lobby.Lobby, error) {
	userInfo, err := user.UserDataFromContext(newPlayerSession.Request.Context())
	if err != nil {
		log.Error().Err(err).Msg("User context error")
		return nil, nil, fmt.Errorf("no userInfo info found")
	}

	queryParams := newPlayerSession.Request.URL.Query()
	param := queryParams.Get("lobby")
	if param == "" {
		return nil, nil, fmt.Errorf("lobby query parameter not found")
	}

	lobbyId, err := strconv.Atoi(param)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert %s to int: %v", param, err)
	}

	lobbyInfo, err := manager.lobbyService.GetLobbyFromID(lobbyId)
	if err != nil {
		return nil, nil, err
	}

	return userInfo, lobbyInfo, err
}
