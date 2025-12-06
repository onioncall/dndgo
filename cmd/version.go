package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of dndgo",
	Long:  `All software has versions. This is dndgo's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("dndgo version %s (built %s)\n", Version, BuildDate)
	},
}

// Version information - can be overridden at build time with -ldflags
var (
	Version   = "dev"
	BuildDate = "unknown"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
