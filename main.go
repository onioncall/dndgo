package main

import (
	"os"

	"github.com/onioncall/dndgo/cmd"
	"github.com/onioncall/dndgo/logger"
)

func main() {
	defer logger.RegisterPanicHandler()

	// Set version info from root package
	cmd.Version = Version
	cmd.BuildDate = BuildDate

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
