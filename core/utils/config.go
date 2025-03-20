package utils

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"sync"
)

var (
	once             = sync.Once{}
	lobbyLimit int64 = 0
)

func InitConfig() {
	baseDir := getBaseDir()
	configDir := getConfigDir(baseDir)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)

	err := setupConfigOptions(configDir)
	if err != nil {
		log.Fatal().Err(err).Msg("could not setup config options")
	}

	if err := viper.SafeWriteConfig(); err != nil {
		var configFileAlreadyExistsError viper.ConfigFileAlreadyExistsError
		if !errors.As(err, &configFileAlreadyExistsError) {
			log.Fatal().Err(err).Msg("viper could not write to config file")
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("could not read config file")
	}

	log.Info().Msg("Connected and watching config file")
	viper.WatchConfig()

	// maybe create viper instance and return from this function
	// future setup in case https://github.com/spf13/viper/issues/1855 is accepted
}

func getConfigDir(baseDir string) string {
	configDir := fmt.Sprintf("%s/%s", baseDir, "config")
	err := os.MkdirAll(configDir, DefaultFilePerm)
	if err != nil {
		log.Fatal().Err(err).Str("Config dir", configDir).Msgf("could not create config directory")
	}
	return configDir
}

func GetLobbyLimit() int64 {
	once.Do(func() {
		lobbyLimit = int64(loadLobbyLimit())
		log.Info().Msgf("Lobby limit has been set to %v", lobbyLimit)
	})
	return lobbyLimit
}

func loadLobbyLimit() int {
	const defaultLobbyLimit = 3

	limit, ok := os.LookupEnv("LOBBY_LIMIT")
	if !ok {
		return defaultLobbyLimit
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return defaultLobbyLimit
	}

	return limitInt
}

func setupConfigOptions(configDir string) error {
	// misc application files
	// set database path
	viper.SetDefault("db_path", fmt.Sprintf("%s/gouda_database.db", configDir))
	// create log directory
	viper.SetDefault("log_file", fmt.Sprintf("%s/logs/gouda.log", configDir))
	// Set general settings
	viper.SetDefault("server.port", "11200")

	return nil
}

func getBaseDir() string {
	baseDir := "./appdata"
	if os.Getenv("IS_DOCKER") != "" {
		baseDir = "/appdata"
	}
	return baseDir
}

func getDebugMode() string {
	baseDir := "./appdata"
	if os.Getenv("IS_DOCKER") != "" {
		baseDir = "/appdata"
	}
	return baseDir
}
