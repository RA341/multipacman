package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/big"
	database "server/api"
	"server/entities"
	"strconv"
)

func updateUserAuthToken(db *sql.DB, authToken string, userId int) bool {
	stmt := "UPDATE users set auth_token=? where id=?"
	res, _ := database.RunStatements(db, stmt, false, authToken, strconv.Itoa(userId))
	affected, err := res.RowsAffected()
	if err != nil {
		return false
	}
	if affected != 0 {
		return true
	}
	return false
}

func GetUserAuthToken(db *sql.DB, token string) (string, int) {
	stmt := "SELECT id,username from users where auth_token=?"
	token = hashString(token)
	_, res := database.RunStatements(db, stmt, true, token)
	user := ""
	id := 0
	for res.Next() {
		err := res.Scan(&id, &user)
		if err != nil || user == "" {
			log.Print("Incorrect auth token" + err.Error())
			return "", 0
		}
	}

	return user, id
}

func retrieveUser(db *sql.DB, username string) entities.User {
	stmt := "SELECT * from users where username=?"
	_, res := database.RunStatements(db, stmt, true, username)

	id := 0
	user := ""
	password := ""
	token := ""

	for res.Next() {
		err := res.Scan(&id, &user, &password, &token)
		if err != nil {
			log.Print(err.Error())
			return entities.User{
				ID:       0,
				Username: "",
				Password: "",
				Token:    "",
			}
		}
	}

	// implies rows retrieved
	if id != 0 {
		return entities.User{
			ID:       id,
			Username: username,
			Password: password,
			Token:    token,
		}
	}

	// implies no rows retrieved
	return entities.User{
		ID:       0,
		Username: "",
		Password: "",
		Token:    "",
	}
}

func createUser(db *sql.DB, username string, password string, authToken string) bool {
	stmt := "INSERT INTO users (username, password, auth_token) values (?,?,?)"

	res, _ := database.RunStatements(db, stmt, false, username, encryptPassword([]byte(password)), authToken)

	affected, err := res.RowsAffected()
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	if affected != 0 {
		return true
	}
	return false
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
		log.Print("Failed Password check" + err.Error())
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
