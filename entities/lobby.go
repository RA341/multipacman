package entities

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type LobbyModel struct {
	matchStarted     bool
	charactersList   []string
	connectedPlayers map[string]*Player
	playerActualIds  map[string][]string
	pelletsEaten     [][]string
	powerUpsEaten    [][]string
	ghostsEaten      []string
	mu               sync.Mutex
}

func NewLobbyModel() *LobbyModel {
	return &LobbyModel{
		matchStarted:     false,
		charactersList:   []string{"gh1", "gh2", "gh3", "pcm"},
		connectedPlayers: make(map[string]*Player),
		playerActualIds:  make(map[string][]string),
		pelletsEaten:     [][]string{},
		powerUpsEaten:    [][]string{},
		ghostsEaten:      []string{},
	}
}

func (l *LobbyModel) Join(playerTmpId, username, playerActualID, lobbyActualId string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.checkIfLobbyIsFull() {
		fmt.Println("Lobby is full")
		return false
	}

	for playerUUID, ids := range l.playerActualIds {
		if ids[0] == playerActualID {
			l.Leave(playerUUID)
			fmt.Println("Stale information found, Resetting player info")
		}
	}

	l.playerActualIds[playerTmpId] = []string{playerActualID, lobbyActualId}

	if len(l.charactersList) == 0 {
		fmt.Println("No available sprites, this should never happen dumbass")
		return false
	}

	spriteId := l.charactersList[len(l.charactersList)-1]
	l.charactersList = l.charactersList[:len(l.charactersList)-1]

	l.connectedPlayers[playerTmpId] = &Player{
		ID:         playerTmpId,
		Username:   username,
		SpriteType: spriteId,
		X:          "0",
		Y:          "0",
	}
	return true
}

func (l *LobbyModel) Leave(playerTmpId string) []string {
	l.mu.Lock()
	defer l.mu.Unlock()

	player, exists := l.connectedPlayers[playerTmpId]
	if !exists {
		return nil
	}

	l.charactersList = append(l.charactersList, player.SpriteType)

	if len(l.charactersList) == 4 {
		l.ghostsEaten = []string{}
		l.pelletsEaten = [][]string{}
		l.powerUpsEaten = [][]string{}
	}

	delete(l.connectedPlayers, playerTmpId)
	returnData := l.playerActualIds[playerTmpId]
	delete(l.playerActualIds, playerTmpId)
	return returnData
}

func (l *LobbyModel) GetGameStateReport() map[string]interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	return map[string]interface{}{
		"ghostsEaten":   l.ghostsEaten,
		"pelletsEaten":  l.pelletsEaten,
		"powerUpsEaten": l.powerUpsEaten,
	}
}

func (l *LobbyModel) checkIfLobbyIsFull() bool {
	return len(l.charactersList) == 0
}

func (l *LobbyModel) StartMatchTimer(duration int, callBackFunc func(int), endFunc func()) {
	l.mu.Lock()
	l.matchStarted = true
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

func (l *LobbyModel) PelletEatenAction(x, y string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.pelletsEaten = append(l.pelletsEaten, []string{x, y})
}

func (l *LobbyModel) PowerUpEatenAction(x, y string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.powerUpsEaten = append(l.powerUpsEaten, []string{x, y})
}

func (l *LobbyModel) GhostEatenAction(ghostID string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.ghostsEaten = append(l.ghostsEaten, ghostID)
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
