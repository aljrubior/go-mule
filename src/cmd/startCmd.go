package cmd

import (
	"github.com/aljrubior/standalone-runtime/handlers"
	"github.com/spf13/cobra"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start",
	Run: func(cmd *cobra.Command, args []string) {

		serverHandler := handlers.NewDefaultServerHandler(ServerRegistrationManager)

		switch len(args) {
		case 1:
			serverId := args[0]
			err := serverHandler.StartServer(serverId)

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
}
