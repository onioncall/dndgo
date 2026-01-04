package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/logger"
	"github.com/spf13/cobra"
)

var (
	logOutput string
	clearLog  bool
)

var rootCmd = &cobra.Command{
	Use:   "dndgo",
	Short: "A D&D helper CLI application",
	Long:  `A CLI application to help with D&D spells, monsters, and character management.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if clearLog {
			err := logger.ClearLog(logOutput)
			if err != nil {
				logger.PrintError(fmt.Sprintf("Failed to clear log file: %v", err))
			}
			return nil
		}

		c, err := handlers.LoadCharacter()
		if err != nil {
			logger.Error(err)
			fmt.Println("Failed to load character data")
			return nil
		}

		if err := handlers.HandleCharacter(c); err != nil {
			logger.Error(err)
			fmt.Println("Failed to handle character data")
		}

		logger.PrintSuccess("Character Update Successful")
		return nil
	},
}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func init() {
	xdgState := os.Getenv("XDG_STATE_HOME")
	if xdgState == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Errorf("failed to get home directory: %w", err))
		}
		xdgState = filepath.Join(homeDir, ".local", "state")
	}

	dbDir := filepath.Join(xdgState, "dndgo", "dndgo.log")

	rootCmd.PersistentFlags().StringVar(&logOutput, "log", dbDir, "log output path, use ':stdout' for stdout")
	rootCmd.Flags().BoolVar(&clearLog, "clear-log", false, "clear the log file")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		l, err := logger.NewFileLogger(logger.LevelInfo, logOutput)
		if err != nil {
			return fmt.Errorf("failed to configure logger: %w", err)
		}
		logger.Log = l
		return nil
	}

	rootCmd.AddCommand(characterCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(tuiCmd)
}
