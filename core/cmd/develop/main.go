package main

import (
	"os"

	"github.com/RA341/multipacman/cmd"
	"github.com/RA341/multipacman/internal/config"
	"github.com/RA341/multipacman/pkg"
)

func main() {
	_ = os.Setenv("MP_DISABLE_AUTH", "true")

	pkg.ConsoleLogger()
	config.Load()
	pkg.FileConsoleLogger(config.Opts.LogFilePath)

	path := "../ui-web/dist"
	cmd.StartServer(path)
}
