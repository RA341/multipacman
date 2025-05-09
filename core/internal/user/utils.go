package user

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

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
