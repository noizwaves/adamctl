package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string = "v0.0.0-dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
