package cmd

import (
	"github.com/onioncall/dndgo/handlers"
	"github.com/spf13/cobra"
)

var (
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "Get details on elements of the DnD World",
		Run: func(cmd *cobra.Command, args []string) {
			s, _ := cmd.Flags().GetString("spell")
			e, _ := cmd.Flags().GetString("equipment")
			m, _ := cmd.Flags().GetString("monster")
			f, _ := cmd.Flags().GetString("feature")
			
			if s != "" {
				handlers.HandleSpellRequest(s)
			} else if e != "" {
				handlers.HandleEquipmentRequest(e)
			} else if m != "" {
				handlers.HandleMonsterRequest(m)
			} else if f != "" {
				handlers.HandleFeatureRequest(f)
			}
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List things",
		Run: func(cmd *cobra.Command, args []string) {
			s, _ := cmd.Flags().GetBool("spell")
			e, _ := cmd.Flags().GetBool("equipment")
			m, _ := cmd.Flags().GetBool("monster")
			f, _ := cmd.Flags().GetBool("feature")

			if s {
				handlers.HandleSpellListRequest()
			} else if e {
				handlers.HandleEquipmentListRequest()
			} else if m {
				handlers.HandleMonsterListRequest()
			} else if f {
				handlers.HandleFeatureListRequest()
			}
		},
	}
)

func init() {
	searchCmd.AddCommand(listCmd)

	searchCmd.Flags().StringP("monster", "m", "", "Name of the monster to look up")
	searchCmd.Flags().StringP("spell", "s", "", "Name of the spell to look up")
	searchCmd.Flags().StringP("equipment", "e", "", "Name of the equipment to look up")
	searchCmd.Flags().BoolP("feature", "f", false, "Name of the feature to look up")

	listCmd.Flags().BoolP("monster", "m", false, "List monsters")
	listCmd.Flags().BoolP("spell", "s", false, "List spells")
	listCmd.Flags().BoolP("equipment", "e", false, "List equipment")
	listCmd.Flags().BoolP("feature", "f", false, "List features")
}
