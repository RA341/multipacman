package service

import (
	"database/sql"
	"fmt"
	"github.com/RA341/multipacman/models"
	database "github.com/RA341/multipacman/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

type LobbyService struct {
	db *gorm.DB
}

func (lobbyService *LobbyService) countUserLobbies(uid int) error {
	var count int64
	result := lobbyService.db.
		Model(&models.Lobby{}).
		Where("id = ?", uid).
		Count(&count)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("unable to count user lobbies")
		return fmt.Errorf("unable to count user lobbies")
	}

	// limit of 3 per user
	if count <= 3 {
		return nil
	} else {
		return fmt.Errorf("user has 3 lobbies: %d", uid)
	}
}

func (lobbyService *LobbyService) createLobby(name string, userId int) error {
	err := lobbyService.countUserLobbies(userId)
	if err != nil {
		return err
	}

	lobby := &models.Lobby{
		LobbyName: name,
		UserID:    int64(userId),
	}

	result := lobbyService.db.Create(lobby)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("unable to create lobby")
		return fmt.Errorf("unable to create lobby")
	}

	return nil
}

func (lobbyService *LobbyService) retrieveLobbies() []DbLobby {
	stmt := `
	SELECT l.id, l.lobby_name, u.username, l.created_at
	FROM lobbies AS l
	JOIN users AS u ON l.uid = u.id;
	`
	_, res := database.RunStatements(db, stmt, true)

	allLobbies := make([]DbLobby, 0)
	var lobbyId int
	var lobbyName string
	var userName string
	var createdA time.Time

	for res.Next() {
		err := res.Scan(&lobbyId, &lobbyName, &userName, &createdA)
		if err != nil {
			return []DbLobby{}
		}
		// insert data
		allLobbies = append(allLobbies, DbLobby{
			LobbyId:   lobbyId,
			LobbyName: lobbyName,
			UserName:  userName,
			CreatedAt: createdA,
		})
	}

	return allLobbies
}

func (lobbyService *LobbyService) RetrieveLobbyIds() []int {
	stmt := "SELECT id FROM lobbies"
	_, res := database.RunStatements(db, stmt, true)

	allLobbies := []int{}
	var uid int

	for res.Next() {
		err := res.Scan(&uid)
		if err != nil {
			return []int{}
		}
		// insert data
		allLobbies = append(allLobbies, uid)
	}

	return allLobbies
}

func (lobbyService *LobbyService) deleteLobby(lobbyId int, userId int) bool {
	stmt := "DELETE FROM lobbies WHERE uid = ? AND id = ?"
	res, _ := database.RunStatements(db, stmt, false, userId, lobbyId)
	affected, err := res.RowsAffected()

	if err != nil {
		log.Print(err)
		return false
	}
	if affected != 0 {
		return true
	}
	return false
}
