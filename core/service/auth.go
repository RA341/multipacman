package service

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/RA341/multipacman/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/big"
)

type AuthService struct {
	db *gorm.DB
}

func (auth *AuthService) createUser(username string, password string, authToken string) error {
	user := models.User{
		Username: username,
		Password: hashString(password),
		Token:    authToken,
	}

	res := auth.db.Create(&user)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("Failed to create user")
		return res.Error
	}

	return nil
}

func (auth *AuthService) updateUserAuthToken(authToken string, userId int) error {
	result := auth.db.Model(&models.User{}).Where("id = ?", userId).Update("auth_token", authToken)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update auth token")
		return result.Error
	}
	return nil
}

func (auth *AuthService) GetUserAuthToken(token string) (models.User, error) {
	user := models.User{}

	result := auth.db.Select("id, username").Where("auth_token = ?", token).Find(&user)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update auth token")
		return models.User{}, errors.New("unable to update token")
	}

	if user.ID == 0 || user.Username == "" {
		log.Error().Msg("Invalid token")
		return models.User{}, errors.New("invalid token")
	}

	return user, nil
}

func (auth *AuthService) retrieveUser(username string) (models.User, error) {
	var user models.User
	result := auth.db.Where("username = ?", username).First(&user)

	// Check if the query was successful
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to retrieve user")
		return models.User{}, fmt.Errorf("unable to retrive user info")
	}

	return user, nil
}

func createAuthToken(length int) string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var randomString []byte

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		randomString = append(randomString, characters[randomIndex.Int64()])
	}

	return encryptPassword(randomString)
}

func checkPassword(inputPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		log.Error().Err(err).Msg("Failed Password check")
		return false
	}
	return true
}
func encryptPassword(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func hashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	// Get the resulting hash sum
	hashedBytes := hash.Sum(nil)
	// Convert the hashed bytes to a hexadecimal string
	hashedString := fmt.Sprintf("%x", hashedBytes)
	return hashedString
}
