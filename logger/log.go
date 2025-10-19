package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
)

var infoLogs []string

func NewLogger() {
	if infoLogs == nil {
		infoLogs = make([]string, 0)
	}

	// Considered persisting this longer, but for now we're
	// only going to keep log records from the last run
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "LOG ERROR- failed to get home directory: %v\n", err)
		return
	}

	logDir := filepath.Join(homeDir, ".config", "dndgo")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "LOG ERROR- failed to create log directory: %v\n", err)
		return
	}

	logPath := filepath.Join(logDir, "log")
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LOG ERROR- failed to clear log file: %v\n", err)
		return
	}
	f.Close()
}

func HandleInfo(info string) {
	if infoLogs == nil {
		return
	}

	infoLogs = append(infoLogs, info)
}

func HandleLogs() {
	// Info logs
	for _, log := range infoLogs {
		fmt.Printf("-> %s\n", log)
	}
	// Panic logs
	if r := recover(); r != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "LOG ERROR- failed to get home directory: %v\n", err)
			fmt.Fprintf(os.Stderr, "PANIC: %v\n", r)
			os.Exit(1)
		}
		logDir := filepath.Join(homeDir, ".config", "dndgo")
		if err := os.MkdirAll(logDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "LOG ERROR- failed to create log directory: %v\n", err)
			fmt.Fprintf(os.Stderr, "PANIC: %v\n", r)
			os.Exit(1)
		}
		logPath := filepath.Join(logDir, "log")
		f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "LOG ERROR- failed to open log file: %v\n", err)
			fmt.Fprintf(os.Stderr, "PANIC: %v\n", r)
			os.Exit(1)
		}
		defer f.Close()

		panicEntry := fmt.Sprintf("[PANIC] %v\n%s\n", r, debug.Stack())
		if _, err := f.WriteString(panicEntry); err != nil {
			fmt.Fprintf(os.Stderr, "LOG ERROR- failed to write panic to log file: %v\n", err)
		}

		fmt.Printf("Application had an unrecoverable error, check log file for more details: %s\n", logPath)
		os.Exit(1)
	}
}
