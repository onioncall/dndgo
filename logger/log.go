package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"time"
)

// Log is the global logger. genrally this should be configured during initialization.
var Log = New(LevelInfo, os.Stdout)

// NewFileLogger creates a logger that writes to the specified file path.
// Use ":stdout" to write to stdout.
func NewFileLogger(level Level, path string) (*Logger, error) {
	var out io.Writer
	if path == ":stdout" {
		out = os.Stdout
	} else {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		out = f
	}
	return New(level, out), nil
}

// tells [log.Output] how many stack frames to skip when
// determining the file and line number to report
const calldepth = 3

// Logger is a leveled logger wrapping the standard library logger
type Logger struct {
	minLevel Level
	out      io.Writer
}

// New creates a new Logger with the specified minimum level and output
func New(level Level, out io.Writer) *Logger {
	return &Logger{
		minLevel: level,
		out:      out,
	}
}

func (l *Logger) log(calldepth int, level Level, msg string) {
	if level < l.minLevel {
		return
	}

	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	file = filepath.Base(file)

	timestamp := time.Now().UTC().Format(time.RFC3339)
	_, err := fmt.Fprintf(l.out, "%s %s:%d: [%s] %s\n", timestamp, file, line, level, msg)
	if err != nil {
		panic(fmt.Errorf("failed to write to log writer, unrecoverable error: %w", err))
	}
}

// Debug logs a debug message
func Debug(v ...any) {
	Log.log(calldepth, LevelDebug, fmt.Sprint(v...))
}

// Debugf logs a formatted debug message
func Debugf(format string, v ...any) {
	Log.log(calldepth, LevelDebug, fmt.Sprintf(format, v...))
}

// Info logs an info message
func Info(v ...any) {
	Log.log(calldepth, LevelInfo, fmt.Sprint(v...))
}

// Infof logs a formatted info message
func Infof(format string, v ...any) {
	Log.log(calldepth, LevelInfo, fmt.Sprintf(format, v...))
}

// Warn logs a warning message
func Warn(v ...any) {
	Log.log(calldepth, LevelWarn, fmt.Sprint(v...))
}

// Warnf logs a formatted warning message
func Warnf(format string, v ...any) {
	Log.log(calldepth, LevelWarn, fmt.Sprintf(format, v...))
}

// Error logs an error message
func Error(v ...any) {
	Log.log(calldepth, LevelError, fmt.Sprint(v...))
}

// Errorf logs a formatted error message
func Errorf(format string, v ...any) {
	Log.log(calldepth, LevelError, fmt.Sprintf(format, v...))
}

// RegisterPanicHandler uses recover to catch and attempt to log panic to the log file
// and print the panic and a user message to the console.
func RegisterPanicHandler() {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("%v\n%s", r, debug.Stack())
		Log.log(calldepth, LevelPanic, msg)
		fmt.Fprintf(os.Stderr, "Application had an unrecoverable error: %v\n", r)
		os.Exit(1)
	}
}

func ConsoleError(text string) {
	fmt.Printf("-> \033[31m%s\033[0m\n", text)
}

func ConsoleSuccess(text string) {
	fmt.Printf("-> \033[32m%s\033[0m\n", text)
}
