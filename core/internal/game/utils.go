package game

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"maps"
	"math/rand"
	"slices"
	"sync"
)

type Point struct {
	X, Y float64
}
type CoordList struct {
	mu        sync.Mutex
	coordList map[Point]struct{}
}

func NewCordList() CoordList {
	return CoordList{
		mu:        sync.Mutex{},
		coordList: map[Point]struct{}{},
	}
}

func (c *CoordList) Add(x, y float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	point := Point{X: x, Y: y}
	c.coordList[point] = struct{}{}
}

func (c *CoordList) GetList() []Point {
	c.mu.Lock()
	defer c.mu.Unlock()
	v := slices.Collect(maps.Keys(c.coordList))
	if v == nil {
		return []Point{}
	}
	return v
}

func (c *CoordList) Len() int {
	return len(c.GetList())
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
