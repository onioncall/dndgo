package main

import (
	// "flag"
	"fmt"
	"os"

	"github.com/onioncall/dndgo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
