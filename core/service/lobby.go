package service

import (
	"fmt"
	rpc "github.com/RA341/multipacman/generated/lobby/v1"
	"github.com/RA341/multipacman/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type LobbyService struct {
	Db *gorm.DB
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

func (lobbyService *LobbyService) RetrieveLobbies() ([]*rpc.Lobby, error) {
	var lobbies []models.Lobby

	res := lobbyService.Db.Find(&lobbies)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("unable to query lobbies")
		return []*rpc.Lobby{}, fmt.Errorf("unable to find lobbies")
	}

	var result []*rpc.Lobby
	for _, lobby := range lobbies {
		result = append(result, lobby.ToRPC())
	}

	return result, nil
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
