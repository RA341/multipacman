package main

import (
	"database/sql"
	"log"
)

func SetupDatabase(db *sql.DB) {
	query := "CREATE TABLE if not exists users ( id INT AUTO_INCREMENT PRIMARY KEY,username VARCHAR(255) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL, auth_token VARCHAR(255), profile_image_url VARCHAR(255))"
	runStatements(db, query)

	query = `CREATE TABLE if not exists lobbies (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				uid INTEGER NOT NULL,
				lobby_name TEXT NOT NULL,
				lobby_id TEXT UNIQUE NOT NULL,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (uid) REFERENCES users(id));`
	runStatements(db, query)

	query = `
		CREATE TRIGGER update_lobbies_updated_at
		AFTER UPDATE ON lobbies
		FOR EACH ROW
		BEGIN
			UPDATE lobbies
			SET updated_at = CURRENT_TIMESTAMP
			WHERE id = OLD.id;
		END;`
	runStatements(db, query)
}

func runStatements(db *sql.DB, query string) {
	prepare, err := db.Prepare(query)
	if err != nil {
		log.Fatal("Db Error" + err.Error())
		return
	}
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			log.Fatal("Db Error" + err.Error())
		}
	}(prepare)
	_, err = prepare.Exec()
	if err != nil {
		log.Fatal("Db Error" + err.Error())
		return
	}
}
