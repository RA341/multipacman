package game

import (
	"encoding/json"
	"fmt"
	"github.com/olahol/melody"
	"sync"
)

type Lobby struct {
	MatchStarted     bool
	CharactersList   []string
	ConnectedPlayers map[string]*melody.Session
	PelletsEaten     [][]float64
	PowerUpsEaten    [][]float64
	GhostsEaten      []string
	mu               sync.Mutex
}

func NewLobbyModel() *Lobby {
	// Create a new Lobby instance
	lobby := &Lobby{
		MatchStarted:     false,
		CharactersList:   []string{"gh1", "gh2", "gh3", "pcm"},
		ConnectedPlayers: make(map[string]*melody.Session),
		PelletsEaten:     [][]float64{},
		PowerUpsEaten:    [][]float64{},
		GhostsEaten:      []string{},
	}
	return lobby
}

func (l *Lobby) Join(player *PlayerEntity, session *melody.Session) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.checkIfLobbyIsFull() {
		fmt.Println("Lobby is full")
		return false
	}

	if len(l.CharactersList) == 0 {
		fmt.Println("No available sprites, this should never happen dumbass")
		return false
	}

	// assign the last available sprite in sprite list
	spriteId := l.CharactersList[len(l.CharactersList)-1]
	player.SpriteType = spriteId
	// pop this sprite
	l.CharactersList = l.CharactersList[:len(l.CharactersList)-1]

	// assign new player to lobby
	l.ConnectedPlayers[player.PlayerId] = session

	return true
}

func (l *Lobby) Leave(player *PlayerEntity) {
	l.mu.Lock()
	defer l.mu.Unlock()
	id := player.PlayerId

	_, exists := l.ConnectedPlayers[id]
	if !exists {
		return
	}

	l.CharactersList = append(l.CharactersList, player.SpriteType)

	if len(l.CharactersList) == 4 {
		l.GhostsEaten = []string{}
		l.PelletsEaten = [][]float64{}
		l.PowerUpsEaten = [][]float64{}
	}

	delete(l.ConnectedPlayers, id)
}

func (l *Lobby) GetGameStateReport() []byte {
	l.mu.Lock()
	defer l.mu.Unlock()

	data := map[string]interface{}{
		"type":          "state",
		"ghostsEaten":   l.GhostsEaten,
		"pelletsEaten":  l.PelletsEaten,
		"powerUpsEaten": l.PowerUpsEaten,
	}

	// Convert map to JSON bytes
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling game state json:", err)
		return nil
	}
	return jsonData
}

func (l *Lobby) checkIfLobbyIsFull() bool {
	return len(l.CharactersList) == 0
}

func (l *Lobby) CountPLayers() int {
	return len(l.ConnectedPlayers)
}

func (l *Lobby) PelletEatenAction(x, y float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.PelletsEaten = append(l.PelletsEaten, []float64{x, y})
}

func (l *Lobby) PowerUpEatenAction(x, y float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.PowerUpsEaten = append(l.PowerUpsEaten, []float64{x, y})
}

func (l *Lobby) GhostEatenAction(ghostID string) {
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
