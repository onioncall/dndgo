package main

import (
	"os"

	"github.com/onioncall/dndgo/cmd"
	"github.com/onioncall/dndgo/logger"
)

var version = "dev"

func main() {
	defer logger.RegisterPanicHandler()

	if err := cmd.Execute(version); err != nil {
		os.Exit(1)
	}
}
