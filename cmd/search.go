package cmd

import (
	"fmt"
	"os"

	"github.com/onioncall/dndgo/handlers"
	"github.com/onioncall/dndgo/logger"
	"github.com/spf13/cobra"
	"golang.org/x/term"
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

			w, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				errLog := fmt.Errorf("Failed to get terminal size: %s", err)
				logger.HandleError(err, errLog)
			}

			if s != "" {
				handlers.HandleSpellRequest(s, w)
			} else if e != "" {
				handlers.HandleEquipmentRequest(e, w)
			} else if m != "" {
				handlers.HandleMonsterRequest(m, w)
			} else if f != "" {
				handlers.HandleFeatureRequest(f, w)
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
