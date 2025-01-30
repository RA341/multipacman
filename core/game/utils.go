package game

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"math/rand"
)

func shuffleArray(array []SpriteType) []SpriteType {
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func getPlayerEntityFromSession(playerSession *melody.Session) (*PlayerEntity, error) {
	pInfo, exists := playerSession.Get(userEntityKey)
	if !exists {
		return nil, fmt.Errorf("player info not in session")
	}
	otherPlayerEntity := pInfo.(*PlayerEntity)
	return otherPlayerEntity, nil
}

func getLobbyEntityFromSession(s *melody.Session) (*LobbyState, error) {
	lobby, exists := s.Get(lobbyEntityKey)
	if !exists {
		log.Info().Msg("Lobby info not found in sessions")
		return nil, fmt.Errorf("lobby state not found in sessions")
	}

	return lobby.(*LobbyState), nil
}
