package cmd

import (
	"github.com/spf13/cobra"
)

var (
	spellName   string
	monsterName string
	weaponName 	string
	charAction  string
)

var rootCmd = &cobra.Command{
	Use:   "dnd-cli",
	Short: "A D&D helper CLI application",
	Long:  `A CLI application to help with D&D spells, monsters, and character management.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(characterCmd)
	rootCmd.AddCommand(searchCmd)
}
