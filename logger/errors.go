package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type errorEntry struct {
	err       error
	timestamp time.Time
}

var consoleErrors map[error][]errorEntry

func NewLogger() {
	if consoleErrors == nil {
		consoleErrors = make(map[error][]errorEntry)
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
	
	logPath := filepath.Join(logDir, "log.txt")
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LOG ERROR- failed to clear log file: %v\n", err)
		return
	}
	f.Close()
}

// Handles a friendly cli error, and a stacktrace-like error to a log file if required
func HandleError(cliErr error, logError error) {
	if cliErr == nil {
		return
	}
	
	// If no log error provided and this cli error hasn't been logged yet, mark it as "no stack trace needed"
	if logError == nil {
		if _, exists := consoleErrors[cliErr]; !exists {
			consoleErrors[cliErr] = nil
		}
		return
	}
	
	consoleErrors[cliErr] = append(consoleErrors[cliErr], errorEntry{
		err:       logError,  // Store the detailed error
		timestamp: time.Now(),
	})
}

func LogErrors() {
	if len(consoleErrors) < 1 {
		return
	}
	
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
	
	logPath := filepath.Join(logDir, "log.txt")
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LOG ERROR- failed to open log file: %v\n", err)
		return
	}
	defer f.Close()
	
	for cliErr, entries := range consoleErrors {
		fmt.Println(cliErr)
			
		consoleEntry := fmt.Sprintf("%v\n", cliErr)
		if _, err := f.WriteString(consoleEntry); err != nil {
			fmt.Fprintf(os.Stderr, "LOG ERROR- failed to write to log file: %v\n", err)
		}
		
		if entries != nil {
			for _, logError := range entries {
				timestamp := logError.timestamp.Format(time.RFC3339)
				logEntry := fmt.Sprintf("	[%s] %v\n", timestamp, logError.err)
				if _, err := f.WriteString(logEntry); err != nil {
					fmt.Fprintf(os.Stderr, "LOG ERROR- failed to write to log file: %v\n", err)
				}
			}
		} else {
			// No stack trace needed
			if _, err := f.WriteString("(no additional details)\n"); err != nil {
				fmt.Fprintf(os.Stderr, "LOG ERROR- failed to write to log file: %v\n", err)
			}
		}
	}
}
