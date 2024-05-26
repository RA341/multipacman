package entities

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type LobbyModel struct {
	MatchStarted     bool
	CharactersList   []string
	ConnectedPlayers map[string]*PlayerEntity
	PelletsEaten     [][]float64
	PowerUpsEaten    [][]float64
	GhostsEaten      []string
	mu               sync.Mutex
}

func NewLobbyModel() *LobbyModel {
	// Create a new LobbyModel instance
	lobby := &LobbyModel{
		MatchStarted:     false,
		CharactersList:   []string{"gh1", "gh2", "gh3", "pcm"},
		ConnectedPlayers: make(map[string]*PlayerEntity),
		PelletsEaten:     [][]float64{},
		PowerUpsEaten:    [][]float64{},
		GhostsEaten:      []string{},
	}
	return lobby
}

func (l *LobbyModel) Join(player *PlayerEntity) bool {
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
	l.ConnectedPlayers[player.PlayerId] = player

	return true
}

func (l *LobbyModel) Leave(player *PlayerEntity) {
	l.mu.Lock()
	defer l.mu.Unlock()
	id := player.PlayerId

	player, exists := l.ConnectedPlayers[id]
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

func (l *LobbyModel) GetGameStateReport() []byte {
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

func (l *LobbyModel) checkIfLobbyIsFull() bool {
	return len(l.CharactersList) == 0
}

func (l *LobbyModel) StartMatchTimer(duration int, callBackFunc func(int), endFunc func()) {
	l.mu.Lock()
	l.MatchStarted = true
	l.mu.Unlock()

	go func() {
		timer := duration
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if timer <= 0 {
				endFunc()
				return
			}

			l.mu.Lock()
			callBackFunc(timer)
			l.mu.Unlock()

			timer--
		}
	}()
}

func (l *LobbyModel) PelletEatenAction(x, y float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.PelletsEaten = append(l.PelletsEaten, []float64{x, y})
}

func (l *LobbyModel) PowerUpEatenAction(x, y float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.PowerUpsEaten = append(l.PowerUpsEaten, []float64{x, y})
}

func (l *LobbyModel) GhostEatenAction(ghostID string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.GhostsEaten = append(l.GhostsEaten, ghostID)
}

func shuffleArray(array []string) []string {
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
	return array
}

//func main() {
//	// Example usage of LobbyModel
//	lobby := NewLobbyModel()
//	fmt.Println(lobby.Join("tmp1", "user1", "actual1", "lobby1"))
//	fmt.Println(lobby.Join("tmp2", "user2", "actual2", "lobby2"))
//	fmt.Println(lobby.GetGameStateReport())
//}
