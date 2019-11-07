package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a project, or add notifications and variables to projects or environments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	},
}

func init() {
	addCmd.AddCommand(addVariableCmd)
	addCmd.AddCommand(addSlackNotificationCmd)
	addCmd.AddCommand(addProjectSlackNotificationCmd)
	addCmd.AddCommand(addRocketChatNotificationCmd)
	addCmd.AddCommand(addProjectRocketChatNotificationCmd)
}
