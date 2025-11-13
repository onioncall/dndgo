package main

import (
	"fmt"
	"os"

	"github.com/onioncall/dndgo/cmd"
	"github.com/onioncall/dndgo/logger"
)

var profile bool

func main() {
	logger.NewLogger()
	defer logger.HandleLogs()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
