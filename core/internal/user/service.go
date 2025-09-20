package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Service struct {
	Db         *gorm.DB
	BypassAuth bool
}

func NewService(db *gorm.DB, bypassAuth bool) *Service {
	if bypassAuth {
		log.Warn().Msg("AUTH IS DISABLED DUMB DUMB, HOPE YOU KNOW WHAT YOU ARE DOING")
	}

	return &Service{
		Db:         db,
		BypassAuth: bypassAuth,
	}
}

func (auth *Service) Register(username, password string, isGuest bool) error {
	user := User{
		Username: username,
		Password: encryptPassword([]byte(password)),
		Token:    "",
		Guest:    isGuest,
	}

	res := auth.Db.Create(&user)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("Failed to create user")
		return res.Error
	}

	log.Info().Any("user", user.Username).Msg("Created user")

	return nil
}

func (auth *Service) Login(username, inputPassword string) (*User, error) {
	var user User
	result := auth.Db.
		Where("username = ?", username).
		First(&user)

	if result.Error != nil || user.Username == "" {
		log.Error().Err(result.Error).Any("user", user).Msg("Failed to login")
		return &User{}, fmt.Errorf("failed retrive user info")
	}

	if !checkPassword(inputPassword, user.Password) {
		log.Error().Err(result.Error).Msg("invalid user/password")
		return &User{}, fmt.Errorf("invalid user/password")
	}

	finalUser, err := auth.newUserAuthToken(user.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update user token")
		return &User{}, fmt.Errorf("failed update user token")
	}

	return finalUser, nil
}

func (auth *Service) Logout(userId uint) (*User, error) {
	return auth.updateUserAuthToken(userId, "")
}

func (auth *Service) newUserAuthToken(userId uint) (*User, error) {
	token := CreateAuthToken(32)
	user, err := auth.updateUserAuthToken(userId, hashString(token))
	if err != nil {
		return nil, err
	}
	// return un-hashed token
	user.Token = token
	return user, nil
}

func (auth *Service) updateUserAuthToken(userId uint, token string) (*User, error) {
	var user User
	result := auth.Db.
		Model(&user).
		Where("id = ?", userId).
		Update("token", token).
		Find(&user)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update auth token")
		return &User{}, result.Error
	}

	return &user, nil
}

// VerifyAuthHeader will check for direct auth header then cookie
func (auth *Service) VerifyAuthHeader(headers http.Header) (*User, error) {
	if auth.BypassAuth {
		// development test user
		return &User{
			Username: "test-user",
			Password: "",
			Token:    "",
			Guest:    false,
		}, nil
	}

	clientToken := headers.Get(AuthHeaderKey)
	if clientToken != "" {
		return auth.verifyToken(clientToken)
	}

	// for flutter clients
	clientToken = headers.Get("Sec-Websocket-Protocol")
	if clientToken != "" {
		return auth.verifyToken(clientToken)
	}

	req := http.Request{Header: headers}
	authCookie, err := req.Cookie(AuthHeaderKey)
	if err != nil {
		return nil, fmt.Errorf("cookie: %s not found", AuthHeaderKey)
	}

	return auth.verifyToken(authCookie.Value)
}

func (auth *Service) verifyToken(token string) (*User, error) {
	user := User{}

	result := auth.Db.
		Where("token = ?", hashString(token)).
		Find(&user)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update auth token")
		return &User{}, errors.New("unable to update token")
	}

	if user.ID == 0 || user.Username == "" {
		return &User{}, errors.New("invalid token")
	}

	// return un-hashed token
	user.Token = token
	return &user, nil
}

func (auth *Service) retrieveUser(username string) (User, error) {
	var user User
	result := auth.Db.Where("username = ?", username).First(&user)

	// Check if the query was successful
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to retrieve user")
		return User{}, fmt.Errorf("unable to retrive user info")
	}

	return user, nil
}
