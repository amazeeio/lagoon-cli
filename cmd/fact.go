package cmd

import (
	"context"
	"fmt"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
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
		if cmdProjectName == "" {
			return fmt.Errorf("Missing arguments - Project name is not defined")
		}

		if cmdProjectEnvironment == "" {
			return fmt.Errorf("Missing arguments - Environment name is not defined")
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

		factExists, factExistsErr := lagoon.FactExists(context.TODO(), projectDetails.ID, environment.Name, name, lc)

		if factExistsErr != nil {
			return factExistsErr
		}

		if factExists {
			return fmt.Errorf("Fact '%s' already exists", name)
		}

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
		if cmdProjectName == "" {
			return fmt.Errorf("Missing arguments - Project name is not defined")
		}

		if cmdProjectEnvironment == "" {
			return fmt.Errorf("Missing arguments - Environment name is not defined")
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

		factExists, factExistsErr := lagoon.FactExists(context.TODO(), projectDetails.ID, environment.Name, name, lc)
		if factExistsErr != nil {
			return factExistsErr
		}

		if !factExists {
			return fmt.Errorf("Fact '%s' does not exist for %s:%s", name, cmdProjectName, cmdProjectEnvironment)
		}

		_, errorval := lagoon.DeleteFact(context.TODO(), environment.ID, name, lc)
		if errorval != nil {
			return errorval
		}

		factExists, factExistsErr = lagoon.FactExists(context.TODO(), projectDetails.ID, environment.Name, name, lc)
		if factExistsErr != nil {
			return factExistsErr
		}

		if !factExists {
			resultData := output.Result{
				Result:     "success",
				ResultData: nil,
			}
			output.RenderResult(resultData, outputOptions)
		} else {
			return fmt.Errorf("Fact '%s' still exists for %s:%s", name, cmdProjectName, cmdProjectEnvironment)
		}

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
