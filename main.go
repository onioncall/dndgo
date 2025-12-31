package main

import (
	"os"

	"github.com/onioncall/dndgo/character-management/db"
	"github.com/onioncall/dndgo/cmd"
	"github.com/onioncall/dndgo/logger"
)

var version = "dev"

func main() {
	defer logger.RegisterPanicHandler()

	if err := db.Init(); err != nil {
		logger.Errorf("Failed to initialize services: %v", err.Error())
		os.Exit(1)
	}
	defer db.Repo.Deinit()

	if err := cmd.Execute(version); err != nil {
		os.Exit(1)
	}
}
