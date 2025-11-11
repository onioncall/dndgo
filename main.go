package main

import (
	"fmt"
	"os"
	// "runtime/trace"

	"github.com/onioncall/dndgo/cmd"
	"github.com/onioncall/dndgo/logger"
)

var profile bool

func main() {
	// f, _ := os.Create("trace.out")
	// trace.Start(f)
	// defer trace.Stop()

	logger.NewLogger()
	defer logger.HandleLogs()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
