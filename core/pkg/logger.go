package pkg

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func getConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}
}

func getBaseLogger() zerolog.Logger {
	return log.With().Caller().Logger()
}

func FileConsoleLogger(logFilePath string) {
	log.Logger = getBaseLogger().Output(zerolog.MultiLevelWriter(GetFileLogger(logFilePath), getConsoleWriter()))
}

func ConsoleLogger() {
	log.Logger = getBaseLogger().Output(getConsoleWriter())
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
