package game

import (
	"encoding/json"
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
	Ghost1 SpriteType = "gh1"
	Ghost2 SpriteType = "gh2"
	Ghost3 SpriteType = "gh3"
	Pacman SpriteType = "pacman"
)

func NewPlayerEntity(userId uint, username string) *PlayerEntity {
	return &PlayerEntity{
		Type:       string(Join),
		PlayerId:   strconv.Itoa(int(userId)),
		Username:   username,
		SpriteType: "",
		X:          0,
		Y:          0,
		ExtraInfo:  "",
	}
}

// PlayerEntity represents a player in the game
type PlayerEntity struct {
	Type       string     `json:"type"`
	PlayerId   string     `json:"playerid"`
	Username   string     `json:"user"`
	SpriteType SpriteType `json:"spriteType"`
	X          int        `json:"x"`
	Y          int        `json:"y"`
	ExtraInfo  string     `json:"extraInfo"`
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
