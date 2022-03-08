package cmd

import (
	"github.com/aljrubior/go-mule/handlers"
	"github.com/spf13/cobra"
	"log"
)

var hybridToken *string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create",
	Run: func(cmd *cobra.Command, args []string) {

		serverHandler := handlers.NewDefaultServerHandler(ServerRegistrationManager, ConfigManager)

		switch len(args) {
		case 1:
			serverName := args[0]
			muleVersion := "4.4.0"
			agentVersion := "2.4.27"
			environment := "qax"

			err := serverHandler.CreateServer(*hybridToken, serverName, muleVersion, agentVersion, environment)

			if err != nil {
				println("Error!")
				log.Fatalln(err.Error())
				return
			}

			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	hybridToken = createCmd.Flags().StringP("hybrid", "H", "", "-H <TOKEN>")
}
