package service

import (
	"gorm.io/gorm"
	"sync"
)

func InitSrv(db *gorm.DB) (*AuthService, *LobbyService) {
	// setup service structs
	authService := &AuthService{Db: db}
	lobSrv := &LobbyService{
		Db:          db,
		Connections: map[uint]chan bool{},
		Mu:          &sync.RWMutex{},
		playerCount: sync.Map{},
	}

	return authService, lobSrv
}
