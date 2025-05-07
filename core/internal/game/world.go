package game

import (
	"encoding/json"
	"fmt"
	"github.com/RA341/multipacman/pkg"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"sync"
)

type World struct {
	cancelChan          chan struct{}
	MatchStarted        bool
	IsPoweredUp         bool
	CharactersList      []SpriteType
	GhostsIdsEaten      []SpriteType
	ConnectedPlayers    *pkg.Map[string, *melody.Session]
	PelletsCoordEaten   CoordList
	PowerUpsCoordsEaten CoordList
	worldLock           sync.Mutex
}

func NewWorldState() *World {
	return &World{
		MatchStarted:        false,
		IsPoweredUp:         false,
		CharactersList:      []SpriteType{Ghost1, Ghost2, Ghost3, Pacman},
		ConnectedPlayers:    &pkg.Map[string, *melody.Session]{},
		PelletsCoordEaten:   NewCordList(),
		PowerUpsCoordsEaten: NewCordList(),
		GhostsIdsEaten:      []SpriteType{},
		worldLock:           sync.Mutex{},
		cancelChan:          make(chan struct{}, 1),
	}
}

func (l *World) Join(player *PlayerEntity, session *melody.Session) error {
	if l.IsLobbyFull() {
		log.Error().Msg("lobby is full")
		return fmt.Errorf("lobby is full")
	}

	if len(l.CharactersList) == 0 {
		log.Error().Msg("No available sprites, this should never happen dumbass")
		return fmt.Errorf("no available sprites")
	}

	// assign the last available sprite in sprite list
	spriteId := l.CharactersList[len(l.CharactersList)-1]
	player.SpriteType = spriteId
	// pop this sprite
	l.CharactersList = l.CharactersList[:len(l.CharactersList)-1]

	// assign new player to world
	l.ConnectedPlayers.Store(player.PlayerId, session)

	return nil
}

func (l *World) Leave(player *PlayerEntity) {
	id := player.PlayerId

	_, exists := l.ConnectedPlayers.Load(id)
	if !exists {
		return
	}

	l.CharactersList = append(l.CharactersList, player.SpriteType)

	if len(l.CharactersList) == 4 {
		l.GhostsIdsEaten = []SpriteType{}
		l.PelletsCoordEaten = NewCordList()
		l.PowerUpsCoordsEaten = NewCordList()
	}

	l.ConnectedPlayers.Delete(id)
}

func (l *World) GetGameStateReport(secretToken, username, spriteId string) ([]byte, error) {
	data := map[string]interface{}{
		"type":          "state",
		"ghostsEaten":   l.GhostsIdsEaten,
		"pelletsEaten":  l.PelletsCoordEaten.GetList(),
		"powerUpsEaten": l.PowerUpsCoordsEaten.GetList(),
		"secretToken":   secretToken,
		"spriteId":      spriteId,
		"username":      username,
	}
	return json.Marshal(data)
}

func (l *World) MovePlayer(player *PlayerEntity, x, y float64) {
	player.X = x
	player.Y = y
}

func (l *World) IsLobbyFull() bool {
	return len(l.CharactersList) == 0
}

func (l *World) GameOver() {
	l.cancelChan <- struct{}{}
}

func (l *World) waitForGameOver() {
	<-l.cancelChan
}

type Consumable interface {
}

func (l *World) EatPellet(pelletX, PelletY float64) {
	l.PelletsCoordEaten.Add(pelletX, PelletY)
}

func (l *World) EatPowerUp(powerUpX, powerUpY float64) {
	if l.IsPoweredUp {
		// already powered up do nothing
		return
	}

	l.PowerUpsCoordsEaten.Add(powerUpX, powerUpY)
	l.IsPoweredUp = true
}

func (l *World) GhostEatenAction(ghostID SpriteType) {
	l.worldLock.Lock()
	defer l.worldLock.Unlock()

	l.GhostsIdsEaten = append(l.GhostsIdsEaten, ghostID)
}
