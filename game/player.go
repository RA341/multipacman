package game

import (
	json2 "encoding/json"
	"log"
)

type Player struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	User       string `json:"user"`
	SpriteType string `json:"spriteType"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
}

func (p *Player) ToJson() []byte {
	marshal, err := json2.Marshal(p)
	if err != nil {
		log.Fatal("Failed to convert to json" + err.Error())
		return nil
	}

	return marshal
}

func CreateUser(user string, id string) Player {
	x := 0
	y := 0

	// Todo add db connection and save it to DB
	return Player{
		Type:       "",
		ID:         id,
		User:       user,
		SpriteType: "",
		X:          x,
		Y:          y,
	}
}

func DeleteUser(uId string) {

}
