package game

type Lobby struct {
	Id        string   `json:"gameId"`
	Players   []Player `json:"players"`
	IsPlaying bool     `json:"isPlaying"`
}

func JoinLobby(player Player) {

}
