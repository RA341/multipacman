package database

import (
	"github.com/RA341/multipacman/internal/config"
	"github.com/RA341/multipacman/internal/lobby"
	"github.com/RA341/multipacman/internal/user"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbPath := config.Opts.DbPath
	connectionStr := sqlite.Open(dbPath + "?_journal_mode=WAL&_busy_timeout=5000")
	dbConf := &gorm.Config{
		PrepareStmt: true,
	}

	db, err := gorm.Open(connectionStr, dbConf)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(user.User{}, lobby.Lobby{})
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to migrate database")
	}

	log.Info().Msgf("successfully connected to database: %s", dbPath)

	return db
}
