/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/onioncall/dndgo/handlers"
	"github.com/spf13/cobra"
)

var spell string

// spellCmd represents the spell command
var spellCmd = &cobra.Command{
	Use:   "spell",
	Short: "Get helpful data on DnD 5e Spells",
	Long: `Get helpful data on DnD 5e Spells, pass the spell name as an argument.\n
	If you want a list of spells, pass l or list as an argument"`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandleMonsterRequest(spell)
	},
}

func init() {
	// rootCmd.AddCommand(spellCmd)
	spellCmd.Flags().StringVarP(&spell, "spell", "s", "", "Name of the spell to look up")
	// spellCmd.MarkFlagRequired("spell")
}
