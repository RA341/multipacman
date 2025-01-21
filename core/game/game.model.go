package game

import "encoding/json"

func NewPlayerEntity() *PlayerEntity {
	return &PlayerEntity{
		Type:       "",
		PlayerId:   "",
		Username:   "",
		SpriteType: "",
		X:          0,
		Y:          0,
		ExtraInfo:  "",
	}
}

// PlayerEntity represents a player in the game
type PlayerEntity struct {
	Type       string `json:"type"`
	PlayerId   string `json:"playerid"`
	Username   string `json:"user"`
	SpriteType string `json:"spriteType"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	ExtraInfo  string `json:"extraInfo"`
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
