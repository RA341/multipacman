package game

import (
	"encoding/json"
	"fmt"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"sync"
)

func NewLobbyStateEntity() *LobbyState {
	lobby := &LobbyState{
		MatchStarted:        false,
		CharactersList:      []SpriteType{Ghost1, Ghost2, Ghost3, Pacman},
		ConnectedPlayers:    map[string]*melody.Session{},
		PelletsCoordEaten:   [][]float64{},
		PowerUpsCoordsEaten: [][]float64{},
		GhostsEaten:         []SpriteType{},
	}
	return lobby
}

type LobbyState struct {
	MatchStarted        bool
	CharactersList      []SpriteType
	ConnectedPlayers    map[string]*melody.Session
	PelletsCoordEaten   [][]float64
	PowerUpsCoordsEaten [][]float64
	GhostsEaten         []SpriteType
	mu                  sync.Mutex
}

func (l *LobbyState) Join(player *PlayerEntity, session *melody.Session) error {
	l.mu.Lock()
	defer l.mu.Unlock()

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

	// assign new player to lobby
	l.ConnectedPlayers[player.PlayerId] = session

	return nil
}

func (l *LobbyState) Leave(player *PlayerEntity) {
	l.mu.Lock()
	defer l.mu.Unlock()
	id := player.PlayerId

	_, exists := l.ConnectedPlayers[id]
	if !exists {
		return
	}

	l.CharactersList = append(l.CharactersList, player.SpriteType)

	if len(l.CharactersList) == 4 {
		l.GhostsEaten = []SpriteType{}
		l.PelletsCoordEaten = [][]float64{}
		l.PowerUpsCoordsEaten = [][]float64{}
	}

	delete(l.ConnectedPlayers, id)
}

func (l *LobbyState) GetGameStateReport() ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	data := map[string]interface{}{
		"type":          "state",
		"ghostsEaten":   l.GhostsEaten,
		"pelletsEaten":  l.PelletsCoordEaten,
		"powerUpsEaten": l.PowerUpsCoordsEaten,
	}

	// Convert map to JSON bytes
	return json.Marshal(data)
}

func (l *LobbyState) IsLobbyFull() bool {
	return len(l.CharactersList) == 0
}

func (l *LobbyState) PelletEatenAction(x, y float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.PelletsCoordEaten = append(l.PelletsCoordEaten, []float64{x, y})
}

func (l *LobbyState) PowerUpEatenAction(x, y float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.PowerUpsCoordsEaten = append(l.PowerUpsCoordsEaten, []float64{x, y})
}

func (l *LobbyState) GhostEatenAction(ghostID SpriteType) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.GhostsEaten = append(l.GhostsEaten, ghostID)
}

//func main() {
//	// Example usage of Lobby
//	lobby := NewLobbyModel()
//	fmt.Println(lobby.Join("tmp1", "user1", "actual1", "lobby1"))
//	fmt.Println(lobby.Join("tmp2", "user2", "actual2", "lobby2"))
//	fmt.Println(lobby.GetGameStateReport())
//}
