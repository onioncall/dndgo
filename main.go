package main

import (
	"os"

	"github.com/onioncall/dndgo/cmd"
	"github.com/onioncall/dndgo/di"
	"github.com/onioncall/dndgo/logger"
)

var profile bool

func main() {
	defer logger.RegisterPanicHandler()

	if err := di.Init(); err != nil {
		logger.Errorf("Failed to initialize services: %v", err.Error())
		os.Exit(1)
	}
	defer di.Deinit()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
