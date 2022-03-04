package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var hybridToken *string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create",
	Run: func(cmd *cobra.Command, args []string) {

		switch len(args) {
		case 1:
			serverName := args[0]
			muleVersion := "4.4.0"
			agentVersion := "2.4.27"
			environment := "qax"

			_, err := ServerRegistrationManager.Register(*hybridToken, serverName, muleVersion, agentVersion, environment)

			if err != nil {
				log.Fatalln(err.Error())
				return
			}
			println("Standalone created")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	hybridToken = createCmd.Flags().StringP("hybrid", "H", "", "-H <TOKEN>")
}
