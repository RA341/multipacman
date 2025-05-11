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
	gameOverChan        chan string
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
		gameOverChan:        make(chan string, 1),
	}
}

func (w *World) Join(player *PlayerEntity, session *melody.Session) error {
	if w.IsLobbyFull() {
		log.Error().Msg("lobby is full")
		return fmt.Errorf("lobby is full")
	}

	if len(w.CharactersList) == 0 {
		log.Error().Msg("No available sprites, this should never happen dumbass")
		return fmt.Errorf("no available sprites")
	}

	// assign the last available sprite in sprite list
	spriteId := w.CharactersList[len(w.CharactersList)-1]
	player.SpriteType = spriteId
	// pop this sprite
	w.CharactersList = w.CharactersList[:len(w.CharactersList)-1]

	// assign new player to world
	w.ConnectedPlayers.Store(player.PlayerId, session)

	return nil
}

func (w *World) Leave(player *PlayerEntity) {
	id := player.PlayerId

	_, exists := w.ConnectedPlayers.Load(id)
	if !exists {
		return
	}

	w.CharactersList = append(w.CharactersList, player.SpriteType)
	w.ConnectedPlayers.Delete(id)

	if len(w.CharactersList) == 4 {
		w.GameOver("all players left, lobby is now empty")
	}
}

func (w *World) GetGameStateReport(secretToken, username, spriteId string, newPlayer *melody.Session) ([]byte, error) {
	connectedMap := map[string]interface{}{}

	for _, otherPlayerSession := range w.ConnectedPlayers.GetValues() {
		if otherPlayerSession == newPlayer {
			continue
		}

		otherPlayerEntity, err := getPlayerEntityFromSession(otherPlayerSession)
		if err != nil {
			continue
		}

		connectedMap[string(otherPlayerEntity.SpriteType)] = map[string]interface{}{
			"username": otherPlayerEntity.Username,
			"x":        otherPlayerEntity.X,
			"y":        otherPlayerEntity.Y,
		}
	}

	data := map[string]interface{}{
		"type":          "state",
		"ghostsEaten":   w.GhostsIdsEaten,
		"activePlayers": connectedMap,
		"pelletsEaten":  w.PelletsCoordEaten.GetList(),
		"powerUpsEaten": w.PowerUpsCoordsEaten.GetList(),
		"secretToken":   secretToken,
		"spriteId":      spriteId,
		"username":      username,
	}
	return json.Marshal(data)
}

func (w *World) MovePlayer(player *PlayerEntity, x, y float64) {
	player.X = x
	player.Y = y
}

func (w *World) IsLobbyFull() bool {
	return len(w.CharactersList) == 0
}

func (w *World) checkGameOver() (reason string) {
	if len(w.GhostsIdsEaten) == 3 {
		return "all ghosts eaten"
	}

	return ""
}

func (w *World) GameOver(reason string) {
	w.gameOverChan <- reason
}

func (w *World) waitForGameOver() string {
	return <-w.gameOverChan
}

type Consumable interface {
}

func (w *World) EatPellet(pelletX, PelletY float64) {
	w.PelletsCoordEaten.Add(pelletX, PelletY)
}

func (w *World) EatPowerUp(powerUpX, powerUpY float64) {
	if w.IsPoweredUp {
		// already powered up do nothing
		return
	}

	w.PowerUpsCoordsEaten.Add(powerUpX, powerUpY)
	w.IsPoweredUp = true
}

func (w *World) GhostEatenAction(ghostID SpriteType) {
	w.worldLock.Lock()
	defer w.worldLock.Unlock()

	w.GhostsIdsEaten = append(w.GhostsIdsEaten, ghostID)
}
