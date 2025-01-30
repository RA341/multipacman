package models

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
}

func (l User) FromRPC(lobby *v1.UserResponse) *User {
	return &User{
		Model: gorm.Model{
			ID: l.ID,
		},
		Username: l.Username,
	}
}

func (l User) ToRPC() *v1.UserResponse {
	return &v1.UserResponse{
		ID:        uint64(l.ID),
		Username:  l.Username,
		AuthToken: l.Token, // make sure this is un-hashed
	}
}
