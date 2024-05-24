package entities

type Player struct {
	ID         string
	Username   string
	SpriteType string
	X          string
	Y          string
	ExtraInfo  string
}

// User storing user data for non-game uses
type User struct {
	ID       string
	Username string
	token    string
}
