package game

import (
	"encoding/json"
	"github.com/RA341/multipacman/service"
	"strconv"
)

type PlayerMessageType string

const (
	Join       PlayerMessageType = "active"
	Disconnect PlayerMessageType = "disconnect"
	Move       PlayerMessageType = "Move"
	Pellet     PlayerMessageType = "inactive"
	Powerup    PlayerMessageType = "pending"
)

type SpriteType string

const (
	Ghost1 SpriteType = "gh0"
	Ghost2 SpriteType = "gh1"
	Ghost3 SpriteType = "gh2"
	Pacman SpriteType = "pacman"
)

func NewPlayerEntity(userId uint, username string) *PlayerEntity {
	return &PlayerEntity{
		Type:        string(Join),
		PlayerId:    strconv.Itoa(int(userId)),
		Username:    username,
		SpriteType:  "",
		X:           "0",
		Y:           "0",
		secretToken: service.CreateAuthToken(5),
	}
}

// PlayerEntity represents a player in the game
type PlayerEntity struct {
	Type        string     `json:"type"`
	PlayerId    string     `json:"playerid"`
	Username    string     `json:"user"`
	SpriteType  SpriteType `json:"spriteType"`
	X           string     `json:"x"`
	Y           string     `json:"y"`
	Dir         string     `json:"dir"`
	secretToken string
}

// ToJSON converts the PlayerEntity to a JSON string
func (p *PlayerEntity) ToJSON() ([]byte, error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

// FromJSON populates the PlayerEntity from a JSON string
func (p *PlayerEntity) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), p)
}
