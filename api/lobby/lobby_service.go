package lobby

import (
	"database/sql"
	"log"
	database "server/api"
	"time"
)

type DbLobby struct {
	LobbyName string    `json:"lobbyName"`
	UserName  string    `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	Joined    int       `json:"joined"`
	LobbyId   int       `json:"lobbyId"`
}

func countUserLobbies(db *sql.DB, uid int) bool {
	stmt := `SELECT COUNT(uid) FROM lobbies where uid=?`
	_, res := database.RunStatements(db, stmt, true, uid)

	count := 0
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			return false
		}
	}

	// lobby limit of 3 per user
	if count < 3 {
		return true
	} else {
		return false
	}
}

func createLobby(db *sql.DB, name string, userId int) int64 {
	if !countUserLobbies(db, userId) {
		// not allowed if user has more than 3 lobbies
		return -2 // error code of 2
	}

	stmt := "INSERT INTO lobbies (uid, lobby_name) values (?, ?)"
	res, _ := database.RunStatements(db, stmt, false, userId, name)
	affected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		return -1
	}
	if affected != 0 {
		id, err := res.LastInsertId()
		if err != nil {
			return -1
		}
		return id
	}
	return -1
}

func retrieveLobbies(db *sql.DB) []DbLobby {
	stmt := `
	SELECT l.id, l.lobby_name, u.username, l.created_at
	FROM lobbies AS l
	JOIN users AS u ON l.uid = u.id;
	`
	_, res := database.RunStatements(db, stmt, true)

	allLobbies := make([]DbLobby, 0)
	var lobbyId int
	var lobbyName string
	var userName string
	var createdA time.Time

	for res.Next() {
		err := res.Scan(&lobbyId, &lobbyName, &userName, &createdA)
		if err != nil {
			return []DbLobby{}
		}
		// insert data
		allLobbies = append(allLobbies, DbLobby{
			LobbyId:   lobbyId,
			LobbyName: lobbyName,
			UserName:  userName,
			CreatedAt: createdA,
		})
	}

	return allLobbies
}

func RetrieveLobbyIds(db *sql.DB) []int {
	stmt := "SELECT id FROM lobbies"
	_, res := database.RunStatements(db, stmt, true)

	allLobbies := []int{}
	var uid int

	for res.Next() {
		err := res.Scan(&uid)
		if err != nil {
			return []int{}
		}
		// insert data
		allLobbies = append(allLobbies, uid)
	}

	return allLobbies
}

func deleteLobby(db *sql.DB, lobbyId int, userId int) bool {
	stmt := "DELETE FROM lobbies WHERE uid = ? AND id = ?"
	res, _ := database.RunStatements(db, stmt, false, userId, lobbyId)
	affected, err := res.RowsAffected()

	if err != nil {
		log.Print(err)
		return false
	}
	if affected != 0 {
		return true
	}
	return false
}
