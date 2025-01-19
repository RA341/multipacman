package models

import "gorm.io/gorm"

// User storing user data for non-game uses
type User struct {
	gorm.Model
	Username string
	Password string
	Token    string
}
