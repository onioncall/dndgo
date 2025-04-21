package cmd

import (
	"github.com/onioncall/dndgo/handlers"
	"github.com/spf13/cobra"
)

var monster string

var monsterCmd = &cobra.Command{
	Use:   "monster",
	Short: "Get helpful data on DnD 5e Monsters",
	Long: `Get helpful data on DnD 5e Monsters, pass the monster name as an argument.\n
	If you want a list of spells, pass l or list as an argument"`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandleMonsterRequest(monster)
	},
}

func init() {
	// rootCmd.AddCommand(monsterCmd)
	monsterCmd.Flags().StringVarP(&monster, "monster", "m", "", "Name of the monster to look up")
	// monsterCmd.MarkFlagRequired("monster")
}
