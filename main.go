package main

import (
	"os"

	"github.com/onioncall/dndgo/cmd"
	"github.com/onioncall/dndgo/logger"
)

func main() {
	defer logger.RegisterPanicHandler()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
