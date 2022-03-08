package cmd

import (
	"github.com/aljrubior/standalone-runtime/handlers"
	"github.com/spf13/cobra"
	"os"
)

var flowsPerApplication *int

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start",
	Run: func(cmd *cobra.Command, args []string) {

		serverHandler := handlers.NewDefaultServerHandler(ServerRegistrationManager, ConfigManager)

		switch len(args) {
		case 1:
			serverId := args[0]
			err := serverHandler.StartServer(serverId, *flowsPerApplication)

			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	flowsPerApplication = startCmd.Flags().IntP("flows-per-app", "", 0, "Number of flows per application")
}
