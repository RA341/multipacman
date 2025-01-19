package models

import (
	v1 "github.com/RA341/multipacman/generated/lobby/v1"
	"gorm.io/gorm"
)

type Lobby struct {
	gorm.Model
	LobbyName string `json:"lobbyName"`
	UserID    int64  `json:"user"`
	Joined    int    `json:"joined"`
}

func (l Lobby) FromRPC(lobby *v1.Lobby) *Lobby {
	return &Lobby{
		Model: gorm.Model{
			ID: uint(lobby.GetID()),
		},
		LobbyName: lobby.GetLobbyName(),
	}
}

func (l Lobby) ToRPC() *v1.Lobby {
	return &v1.Lobby{
		ID:        uint64(l.ID),
		LobbyName: l.LobbyName,
	}
}
