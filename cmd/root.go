package cmd

import (
	"fmt"

	"github.com/onioncall/dndgo/handlers"
	"github.com/onioncall/dndgo/logger"
	"github.com/spf13/cobra"
)

var (
	spellName   string
	monsterName string
	weaponName  string
	charAction  string
)

var rootCmd = &cobra.Command{
	Use:   "dndgo",
	Short: "A D&D helper CLI application",
	Long:  `A CLI application to help with D&D spells, monsters, and character management.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := handlers.LoadCharacter()
		if err != nil {
			errMsg := "Failed to save character json"
			logger.HandleInfo(errMsg)
			panic(fmt.Errorf("%s: %w", errMsg, err))
		}

		err = handlers.HandleCharacter(c)
		if err != nil {
			errMsg := "Failed to handle character"
			logger.HandleInfo(errMsg)
			panic(fmt.Errorf("%s: %w", errMsg, err))
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(characterCmd)
	rootCmd.AddCommand(searchCmd)
}
