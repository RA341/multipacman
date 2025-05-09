package user

import (
	v1 "github.com/RA341/multipacman/generated/auth/v1"
	"gorm.io/gorm"
)

// User storing user data for non-game uses
type User struct {
	gorm.Model
	Username string
	Password string
	Token    string
	Guest    bool
}

func (user *User) FromRPC(userRpc *v1.UserResponse) {}

func (user *User) ToRPC() *v1.UserResponse {
	return &v1.UserResponse{
		ID:        uint64(user.ID),
		Username:  user.Username,
		AuthToken: user.Token, // make sure this is un-hashed
		IsGuest:   user.Guest,
	}
}
