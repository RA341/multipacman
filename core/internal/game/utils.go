package game

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"math/rand"
	"sync"
)

type CoordList struct {
	mu   sync.Mutex
	list []struct{ X, Y float64 }
}

func NewCordList() CoordList {
	return CoordList{
		mu:   sync.Mutex{},
		list: []struct{ X, Y float64 }{},
	}
}

func (c *CoordList) Add(X, Y float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.list = append(c.list, struct{ X, Y float64 }{X: X, Y: Y})
}

func (c *CoordList) GetList() []struct{ X, Y float64 } {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.list
}

func shuffleArray(array []SpriteType) []SpriteType {
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func getPlayerEntityFromSession(playerSession *melody.Session) (*PlayerEntity, error) {
	pInfo, exists := playerSession.Get(userInfoKey)
	if !exists {
		return nil, fmt.Errorf("player info not in session")
	}
	otherPlayerEntity := pInfo.(*PlayerEntity)
	return otherPlayerEntity, nil
}

func getWorldFromSession(s *melody.Session) (*World, error) {
	lobby, exists := s.Get(worldKey)
	if !exists {
		log.Info().Msg("Lobby info not found in sessions")
		return nil, fmt.Errorf("lobby state not found in sessions")
	}

	return lobby.(*World), nil
}
