package models

import (
	"gorm.io/gorm"
)

type Lobby struct {
	gorm.Model
	LobbyName string `json:"lobbyName"`
	UserID    int64  `json:"user"`
	Joined    int    `json:"joined"`
}
