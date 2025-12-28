package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/tui/menu"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the dndgo tui",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(
			menu.New(rootCmd.Version),
			tea.WithAltScreen(),
		)

		if _, err := p.Run(); err != nil {
			panic(err)
		}
	},
}
