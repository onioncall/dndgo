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
				errMsg := "Failed to get terminal size"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			switch {
			case s != "":
				err = handlers.HandleSpellRequest(s, w)
			case e != "":
				err = handlers.HandleEquipmentRequest(e, w)
			case m != "":
				err = handlers.HandleMonsterRequest(m, w)
			case f != "":
				err = handlers.HandleFeatureRequest(f, w)
			}

			if err != nil {
				errMsg := "Failed to handle search request"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
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

			var err error

			switch {
			case s:
				err = handlers.HandleSpellListRequest()
			case e:
				err = handlers.HandleEquipmentListRequest()
			case m:
				err = handlers.HandleMonsterListRequest()
			case f:
				err = handlers.HandleFeatureListRequest()
			}

			if err != nil {
				errMsg := "Failed to handle search list request"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
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
