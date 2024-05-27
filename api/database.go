package api

import (
	"database/sql"
	"log"
)

func SetupDatabase(db *sql.DB, dropTables bool) {
	if dropTables {
		query := `DROP TABLE IF EXISTS users`
		RunStatements(db, query, false)
		query = `DROP TABLE IF EXISTS lobbies`
		RunStatements(db, query, false)
	}

	query := `CREATE TABLE IF NOT EXISTS users ( 
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				username VARCHAR(255) UNIQUE NOT NULL, 
				password VARCHAR(255) NOT NULL, 
				auth_token VARCHAR(255))`
	RunStatements(db, query, false)

	query = `CREATE TABLE IF NOT EXISTS lobbies (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				uid INTEGER NOT NULL,
				lobby_name TEXT NOT NULL,
				lobby_id TEXT UNIQUE NOT NULL,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (uid) REFERENCES users(id));`
	RunStatements(db, query, false)

	query = `
		CREATE TRIGGER IF NOT EXISTS update_lobbies_updated_at
		AFTER UPDATE ON lobbies
		FOR EACH ROW
		BEGIN
			UPDATE lobbies
			SET updated_at = CURRENT_TIMESTAMP
			WHERE id = OLD.id;
		END;`
	RunStatements(db, query, false)
}

func RunStatements(db *sql.DB, query string, isQuery bool, args ...any) (sql.Result, *sql.Rows) {
	prepare, err := db.Prepare(query)
	if err != nil {
		log.Fatal("Db Error" + err.Error())
		return nil, nil
	}
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			log.Fatal("Error while closing " + err.Error())
		}
	}(prepare)

	if isQuery {
		result, err := prepare.Query(args...)
		if err != nil {
			log.Fatal("Db Error " + err.Error())
			return nil, nil
		}
		return nil, result
	} else {
		result, err := prepare.Exec(args...)
		if err != nil {
			log.Fatal("Db Error " + err.Error())
			return nil, nil
		}
		return result, nil
	}
}
