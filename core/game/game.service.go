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
		IsPoweredUp:         false,
		CharactersList:      []SpriteType{Ghost1, Ghost2, Ghost3, Pacman},
		ConnectedPlayers:    map[string]*melody.Session{},
		PelletsCoordEaten:   []int{},
		PowerUpsCoordsEaten: []int{},
		GhostsIdsEaten:      []SpriteType{},
		mu:                  sync.Mutex{},
	}
	return lobby
}

type LobbyState struct {
	MatchStarted        bool
	IsPoweredUp         bool
	CharactersList      []SpriteType
	ConnectedPlayers    map[string]*melody.Session
	PelletsCoordEaten   []int
	PowerUpsCoordsEaten []int
	GhostsIdsEaten      []SpriteType
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
		l.GhostsIdsEaten = []SpriteType{}
		l.PelletsCoordEaten = []int{}
		l.PowerUpsCoordsEaten = []int{}
	}

	delete(l.ConnectedPlayers, id)
}

func (l *LobbyState) GetGameStateReport(secretToken, spriteId string) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	data := map[string]interface{}{
		"type":          "state",
		"ghostsEaten":   l.GhostsIdsEaten,
		"pelletsEaten":  l.PelletsCoordEaten,
		"powerUpsEaten": l.PowerUpsCoordsEaten,
		"secretToken":   secretToken,
		"spriteId":      spriteId,
	}

	// Convert map to JSON bytes
	return json.Marshal(data)
}

func (l *LobbyState) MovePlayer(player *PlayerEntity, x, y string) {
	player.X = x
	player.Y = y
}

func (l *LobbyState) IsLobbyFull() bool {
	return len(l.CharactersList) == 0
}

func (l *LobbyState) EatPellet(pelletId int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.PelletsCoordEaten = append(l.PelletsCoordEaten, pelletId)
}

func (l *LobbyState) EatPowerUp(powerUpId int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.IsPoweredUp {
		// already powered up do nothing
		return
	}

	l.IsPoweredUp = true
	l.PowerUpsCoordsEaten = append(l.PowerUpsCoordsEaten, powerUpId)
}

func (l *LobbyState) GhostEatenAction(ghostID SpriteType) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.GhostsIdsEaten = append(l.GhostsIdsEaten, ghostID)
}

//func main() {
//	// Example usage of Lobby
//	lobby := NewLobbyModel()
//	fmt.Println(lobby.Join("tmp1", "user1", "actual1", "lobby1"))
//	fmt.Println(lobby.Join("tmp2", "user2", "actual2", "lobby2"))
//	fmt.Println(lobby.GetGameStateReport())
//}
