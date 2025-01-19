package utils

import (
	"github.com/RA341/multipacman/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbPath := viper.GetString("db_path")
	if dbPath == "" {
		log.Fatal().Msgf("db_path is empty")
	}

	connectionStr := sqlite.Open(dbPath + "?_journal_mode=WAL&_busy_timeout=5000")
	config := &gorm.Config{
		PrepareStmt: true,
	}

	db, err := gorm.Open(connectionStr, config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(models.User{}, models.Lobby{})
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to migrate database")
	}

	log.Info().Msgf("successfully connected to database: %s", dbPath)

	return db
}
