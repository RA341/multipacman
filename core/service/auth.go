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
	Db *gorm.DB
}

func (auth *AuthService) Register(username, password string) error {
	user := models.User{
		Username: username,
		Password: encryptPassword([]byte(password)),
		Token:    "",
	}

	res := auth.Db.Create(&user)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("Failed to create user")
		return res.Error
	}

	log.Info().Any("user", user).Msg("Created user")

	return nil
}

func (auth *AuthService) Login(username, inputPassword string) (*models.User, error) {
	var user models.User
	result := auth.Db.
		Where("username = ?", username).
		First(&user)

	if result.Error != nil || user.Username == "" {
		log.Error().Err(result.Error).Any("user", user).Msg("Failed to login")
		return &models.User{}, fmt.Errorf("failed retrive user info")
	}

	if !checkPassword(inputPassword, user.Password) {
		log.Error().Err(result.Error).Msg("invalid user/password")
		return &models.User{}, fmt.Errorf("invalid user/password")
	}

	finalUser, err := auth.updateUserAuthToken(user.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update user token")
		return &models.User{}, fmt.Errorf("failed update user token")
	}

	return finalUser, nil
}
func (auth *AuthService) updateUserAuthToken(userId uint) (*models.User, error) {
	token := CreateAuthToken(32)

	var user models.User
	result := auth.Db.
		Model(&user).
		Where("id = ?", userId).
		Update("token", hashString(token)).
		Find(&user)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update auth token")
		return &models.User{}, result.Error
	}
	// return un-hashed token
	user.Token = token
	return &user, nil
}

func (auth *AuthService) VerifyToken(token string) (*models.User, error) {
	user := models.User{}

	result := auth.Db.
		Where("token = ?", hashString(token)).
		Find(&user)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update auth token")
		return &models.User{}, errors.New("unable to update token")
	}

	if user.ID == 0 || user.Username == "" {
		log.Error().Msg("Invalid token")
		return &models.User{}, errors.New("invalid token")
	}

	// return un-hashed token
	user.Token = token
	return &user, nil
}

func (auth *AuthService) retrieveUser(username string) (models.User, error) {
	var user models.User
	result := auth.Db.Where("username = ?", username).First(&user)

	// Check if the query was successful
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to retrieve user")
		return models.User{}, fmt.Errorf("unable to retrive user info")
	}

	return user, nil
}

func CreateAuthToken(length int) string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var randomString []byte

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		randomString = append(randomString, characters[randomIndex.Int64()])
	}

	return string(randomString)
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
