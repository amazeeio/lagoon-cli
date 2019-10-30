package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project or environment",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Using Lagoon:", cmdLagoon, "\n")
		// get a new token if the current one is invalid
		valid := graphql.VerifyTokenExpiry()
		if valid == false {
			loginErr := loginToken()
			if loginErr != nil {
				fmt.Println("Unable to refresh token, you may need to run `lagoon login` first")
				os.Exit(1)
			}
		}
	},
}