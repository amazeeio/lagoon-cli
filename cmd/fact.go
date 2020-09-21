package cmd

import (
	"context"
	// "encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
	// "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var factCmd = &cobra.Command{
	Use:   "fact",
	Short: "Add and update facts",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	},
}

var addFactCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a fact",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)

		projectDetails, err := lagoon.GetProjectByNameForFacts(
			context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		var environment schema.Environment

		lc.EnvironmentByName(context.TODO(), cmdProjectEnvironment, projectDetails.ID, &environment)

		retval, errorval := lagoon.AddFact(context.TODO(), environment.ID, name, value, lc)
		if errorval != nil {
			return errorval
		}

		data := []output.Data{}
		data = append(data, []string{
			returnNonEmptyString(fmt.Sprintf("%v", retval.ID)),
			returnNonEmptyString(fmt.Sprintf("%v", retval.Name)),
			returnNonEmptyString(fmt.Sprintf("%v", retval.Value)),
		})
		output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Name",
				"Value",
			},
			Data: data,
		}, outputOptions)

		return nil
	},
}

var deleteFactCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a fact",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)

		projectDetails, err := lagoon.GetProjectByNameForFacts(
			context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		var environment schema.Environment

		lc.EnvironmentByName(context.TODO(), cmdProjectEnvironment, projectDetails.ID, &environment)

		retval, errorval := lagoon.DeleteFact(context.TODO(), environment.ID, name, lc)
		if errorval != nil {
			return errorval
		}

		if errorval != nil {
			return errorval
		}

		fmt.Println(retval)
		return nil

		// data := []output.Data{}
		// data = append(data, []string{
		// 	returnNonEmptyString(fmt.Sprintf("%v", retval.ID)),
		// 	returnNonEmptyString(fmt.Sprintf("%v", retval.Name)),
		// 	returnNonEmptyString(fmt.Sprintf("%v", retval.Value)),
		// })
		// output.RenderOutput(output.Table{
		// 	Header: []string{
		// 		"ID",
		// 		"Name",
		// 		"Value",
		// 	},
		// 	Data: data,
		// }, outputOptions)

		return nil
	},
}

func init() {
	factCmd.AddCommand(addFactCommand)
	addFactCommand.Flags().StringP("name", "N", "", "The key name of the fact you are adding")
	addFactCommand.Flags().StringP("value", "V", "", "The value of the fact you are adding")
	factCmd.AddCommand(deleteFactCommand)
	deleteFactCommand.Flags().StringP("name", "N", "", "The key name of the fact you are deleting")
}
