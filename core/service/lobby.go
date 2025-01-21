package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"sync"
)

type LobbyService struct {
	Db          *gorm.DB
	Mu          *sync.RWMutex
	Connections map[uint]chan bool
}

func (lobbyService *LobbyService) CreateLobby(lobbyName, username string, userId uint) error {
	err := lobbyService.countUserLobbies(userId)
	if err != nil {
		return err
	}

	lobby := &models.Lobby{
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

func (lobbyService *LobbyService) DeleteLobby(lobbyId uint64, userId uint) error {
	res := lobbyService.Db.Where("user_id", userId).Delete(&models.Lobby{}, lobbyId)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("unable to delete lobby")
		return fmt.Errorf("unable to delete lobby")
	}

	return nil
}

func (lobbyService *LobbyService) RetrieveLobbies() ([]models.Lobby, error) {
	var lobbies []models.Lobby

	res := lobbyService.Db.Find(&lobbies)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("unable to query lobbies")
		return []models.Lobby{}, fmt.Errorf("unable to find lobbies")
	}

	return lobbies, nil
}

func (lobbyService *LobbyService) GetGrpcLobbies() ([]*v1.Lobby, error) {
	lobbies, err := lobbyService.RetrieveLobbies()
	if err != nil {
		log.Error().Err(err).Msg("unable to get lobbies")
		return nil, err
	}

	var grpcLobbies []*v1.Lobby
	for _, lobby := range lobbies {
		grpcLobbies = append(grpcLobbies, lobby.ToRPC())
	}

	return grpcLobbies, nil
}

func (lobbyService *LobbyService) GetAndParseLobbies() ([]byte, error) {
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

func (lobbyService *LobbyService) countUserLobbies(uid uint) error {
	var count int64
	result := lobbyService.Db.
		Model(&models.Lobby{}).
		Where("user_id = ?", uid).
		Count(&count)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("unable to count user lobbies")
		return fmt.Errorf("unable to count user lobbies")
	}

	// limit of 3 per user
	if count+1 <= 3 {
		return nil
	} else {
		return fmt.Errorf("user has 3 lobbies: %d", uid)
	}
}

func (lobbyService *LobbyService) UpdateLobbies() {
	for _, chn := range lobbyService.Connections {
		chn <- true
	}
}

func (lobbyService *LobbyService) NewUpdateChannel(channelId uint) chan bool {
	channel := make(chan bool)

	lobbyService.Mu.Lock()
	lobbyService.Connections[channelId] = channel
	lobbyService.Mu.Unlock()

	log.Debug().Msg("Added to lobby list")
	return channel
}

func (lobbyService *LobbyService) RemoveUpdateChannel(channelIndex uint) {
	lobbyService.Mu.Lock()
	delete(lobbyService.Connections, channelIndex)
	lobbyService.Mu.Unlock()
}
