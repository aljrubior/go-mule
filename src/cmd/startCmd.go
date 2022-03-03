package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start",
	Run: func(cmd *cobra.Command, args []string) {

		println("Standalone started")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
