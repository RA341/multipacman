package main

import (
	"github.com/RA341/multipacman/cmd"
	"github.com/RA341/multipacman/internal/config"
	"github.com/RA341/multipacman/pkg"
)

func main() {
	pkg.ConsoleLogger()
	config.Load()
	pkg.FileConsoleLogger(config.Opts.LogFilePath)

	path := "dist"
	cmd.StartServer(path)
}
