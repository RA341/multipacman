package lobby

import (
	"encoding/json"
	"fmt"
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	Db          *gorm.DB
	Mu          *sync.RWMutex
	Connections map[uint]chan bool
	PlayerCount sync.Map
}

func NewLobbyService(db *gorm.DB) *Service {
	return &Service{
		Db:          db,
		Connections: map[uint]chan bool{},
		Mu:          &sync.RWMutex{},
		PlayerCount: sync.Map{},
	}
}

func (lobbyService *Service) GetLobbyFromID(id int) (*Lobby, error) {
	var lobby *Lobby
	result := lobbyService.Db.Find(&lobby, id)
	if result.Error != nil {
		log.Error().Err(result.Error).Int("lobby-id", id).Msg("no lobby found with id")
		return nil, fmt.Errorf("unable to find lobby for %d", id)
	}

	return lobby, nil
}

func (lobbyService *Service) CreateLobby(lobbyName, username string, userId uint) error {
	err := lobbyService.countUserLobbies(userId)
	if err != nil {
		return err
	}

	lobby := &Lobby{
		LobbyName: lobbyName,
		UserID:    int64(userId),
		Username:  username,
	}

	result := lobbyService.Db.Create(lobby)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("unable to create lobby")
		return fmt.Errorf("unable to create lobby")
	}

	return nil
}

func (lobbyService *Service) DeleteLobby(lobbyId uint64, userId uint) error {
	res := lobbyService.Db.Where("user_id", userId).Delete(&Lobby{}, lobbyId)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("unable to delete lobby")
		return fmt.Errorf("unable to delete lobby")
	}

	return nil
}

func (lobbyService *Service) RetrieveLobbies() ([]Lobby, error) {
	var lobbies []Lobby

	res := lobbyService.Db.Find(&lobbies)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("unable to query lobbies")
		return []Lobby{}, fmt.Errorf("unable to find lobbies")
	}

	return lobbies, nil
}

func (lobbyService *Service) GetGrpcLobbies() ([]*v1.Lobby, error) {
	lobbies, err := lobbyService.RetrieveLobbies()
	if err != nil {
		log.Error().Err(err).Msg("unable to get lobbies")
		return nil, err
	}

	var grpcLobbies []*v1.Lobby
	for _, lobby := range lobbies {
		lobbyTmp := lobby.ToRPC()
		lobbyTmp.PlayerCount = lobbyService.GetLobbyPlayerCount(lobby.ID)
		grpcLobbies = append(grpcLobbies, lobbyTmp)
	}

	return grpcLobbies, nil
}

func (lobbyService *Service) GetAndParseLobbies() ([]byte, error) {
	lobbies, err := lobbyService.RetrieveLobbies()
	if err != nil {
		log.Error().Err(err).Msg("error retrieving lobbies")
	}

	jsonData, err := json.Marshal(lobbies)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling lobbies")
	}

	return jsonData, err
}

func (lobbyService *Service) countUserLobbies(uid uint) error {
	var count int64
	result := lobbyService.Db.
		Model(&Lobby{}).
		Where("user_id = ?", uid).
		Count(&count)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("unable to count user lobbies")
		return fmt.Errorf("unable to count user lobbies")
	}

	if count+1 <= int64(config.Opts.LobbyLimit) {
		return nil
	} else {
		return fmt.Errorf("user has reached max lobby limit of %d", config.Opts.LobbyLimit)
	}
}

func (lobbyService *Service) UpdateLobbies() {
	for _, chn := range lobbyService.Connections {
		chn <- true
	}
}

func (lobbyService *Service) NewUpdateChannel(channelId uint) chan bool {
	channel := make(chan bool)

	lobbyService.Mu.Lock()
	lobbyService.Connections[channelId] = channel
	lobbyService.Mu.Unlock()

	//log.Debug().Msg("Added to lobby list")
	return channel
}

func (lobbyService *Service) RemoveUpdateChannel(channelIndex uint) {
	lobbyService.Mu.Lock()
	delete(lobbyService.Connections, channelIndex)
	lobbyService.Mu.Unlock()
}

func (lobbyService *Service) GetLobbyPlayerCount(lobbyId uint) uint64 {
	val, exists := lobbyService.PlayerCount.Load(lobbyId)
	if !exists {
		return 0
	}

	//log.Debug().Msgf("Lobby %d has player count: %d", lobbyId, val)
	return uint64(val.(int))
}

func (lobbyService *Service) UpdateLobbyPlayerCount(lobbyId uint, count int) {
	//log.Debug().Msgf("Updating Lobby: %d, count: %d", lobbyId, count)
	lobbyService.PlayerCount.Store(lobbyId, count)
}
