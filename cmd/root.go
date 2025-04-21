package cmd

import (
	"github.com/onioncall/dndgo/handlers"
	"github.com/spf13/cobra"
)

var (
	spellName   string
	monsterName string
	charAction  string
)

var rootCmd = &cobra.Command{
	Use:   "dnd-cli",
	Short: "A D&D helper CLI application",
	Long:  `A CLI application to help with D&D spells, monsters, and character management.`,
	Run: func(cmd *cobra.Command, args []string) {
		if spellName != "" {
			handlers.HandleSpellRequest(spellName)
		} else if monsterName != "" {
			handlers.HandleMonsterRequest(monsterName)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add direct flags to root command
	rootCmd.Flags().StringVarP(&spellName, "spell", "s", "", "Name of the spell to look up")
	rootCmd.Flags().StringVarP(&monsterName, "monster", "m", "", "Name of the monster to look up")
	
	// Add subcommands to the root command
	rootCmd.AddCommand(characterCmd)
}
