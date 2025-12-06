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
	spellName   string
	monsterName string
	weaponName  string
	charAction  string
	logOutput   string
)

var rootCmd = &cobra.Command{
	Use:   "dndgo",
	Short: "A D&D helper CLI application",
	Long:  `A CLI application to help with D&D spells, monsters, and character management.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := handlers.LoadCharacter()
		if err != nil {
			return fmt.Errorf("failed to load character data: %w", err)
		}

		if err := handlers.HandleCharacter(c); err != nil {
			return fmt.Errorf("failed to handle character: %w", err)
		}

		logger.Info("Character Update Successful")
		return nil
	},
}

// Main Entrypoint
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to get home directory: %w", err))
	}
	defaultLog := filepath.Join(homeDir, ".config", "dndgo", "log")

	rootCmd.PersistentFlags().StringVar(&logOutput, "log", defaultLog, "log output path, use ':stdout' for stdout")
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
