package game

import (
	"encoding/json"
	"github.com/RA341/multipacman/internal/auth"
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
		X:           0,
		Y:           0,
		secretToken: auth.CreateAuthToken(5),
	}
}

// PlayerEntity represents a player in the game
type PlayerEntity struct {
	Type        string     `json:"type"`
	PlayerId    string     `json:"playerid"`
	Username    string     `json:"user"`
	SpriteType  SpriteType `json:"spriteType"`
	X           float64    `json:"x"`
	Y           float64    `json:"y"`
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

func (p *PlayerEntity) ToMap() map[string]interface{} {
	playerMap := map[string]interface{}{}
	playerMap["type"] = p.Type
	playerMap["playerid"] = p.PlayerId
	playerMap["user"] = p.Username
	playerMap["spriteType"] = p.SpriteType
	playerMap["x"] = p.X
	playerMap["y"] = p.Y
	playerMap["dir"] = p.Dir
	return playerMap
}

// FromJSON populates the PlayerEntity from a JSON string
func (p *PlayerEntity) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), p)
}
