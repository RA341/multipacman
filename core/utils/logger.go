package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// cache os.getenv('debug') value for perf
var (
	consoleWriter = zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}
	baseLogger = log.With().Caller().Logger()
)

func FileConsoleLogger() zerolog.Logger {
	logFile := viper.GetString("log_file")
	return baseLogger.Output(zerolog.MultiLevelWriter(GetFileLogger(logFile), consoleWriter))
}

func ConsoleLogger() zerolog.Logger {
	return baseLogger.Output(consoleWriter)
}

func GetFileLogger(logFile string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // MB
		MaxBackups: 5,  // number of backups
		MaxAge:     30, // days
		Compress:   true,
	}
}
