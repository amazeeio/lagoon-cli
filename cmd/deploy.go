package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"d", "dep"},
	Short:   "Actions for deploying or promoting branches or environments in lagoon",
}

var deployBranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Deploy a new branch",
	Long: `Deploy a new branch
This branch may or may not already exist in lagoon, if it already exists you may want to
use 'lagoon deploy latest' instead`,
	Aliases: []string{"b"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		branch, err := cmd.Flags().GetString("branch")
		if err != nil {
			return err
		}
		if cmdProjectName == "" || branch == "" {
			return fmt.Errorf("Missing arguments: Project name or branch name is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to deploy branch '%s' for project '%s', are you sure?", branch, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.DeployBranch(context.TODO(), &schema.DeployEnvironmentBranchInput{
				Branch:  branch,
				Project: cmdProjectName,
			}, lc)
			if err != nil {
				return err
			}
			response := result.DeployEnvironmentBranch
			if strings.HasPrefix(response, "Error: ") {
				return fmt.Errorf(strings.Trim(response, "Error: "))
			}
			fmt.Println(response)
		}
		return nil
	},
}

var deployPromoteCmd = &cobra.Command{
	Use:     "promote",
	Aliases: []string{"p"},
	Short:   "Promote an environment",
	Long:    "Promote one environment to another",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		sourceEnvironment, err := cmd.Flags().GetString("source")
		if err != nil {
			return err
		}
		destinationEnvironment, err := cmd.Flags().GetString("destination")
		if err != nil {
			return err
		}
		if cmdProjectName == "" || sourceEnvironment == "" || destinationEnvironment == "" {
			return fmt.Errorf("Missing arguments: Project name, source environment, or destination environment is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to promote environment '%s' to '%s' for project '%s', are you sure?", sourceEnvironment, destinationEnvironment, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.DeployPromote(context.TODO(), &schema.DeployEnvironmentPromoteInput{
				SourceEnvironment:      sourceEnvironment,
				DestinationEnvironment: destinationEnvironment,
				Project:                cmdProjectName,
			}, lc)
			if err != nil {
				return err
			}
			response := result.DeployEnvironmentPromote
			if strings.HasPrefix(response, "Error: ") {
				return fmt.Errorf(strings.Trim(response, "Error: "))
			}
			fmt.Println(response)
		}
		return nil
	},
}

var deployLatestCmd = &cobra.Command{
	Use:     "latest",
	Aliases: []string{"l"},
	Hidden:  false,
	Short:   "Deploy latest environment",
	Long: `Deploy latest environment
This environment should already exist in lagoon. It is analogous with the 'Deploy' button in the Lagoon UI`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			return fmt.Errorf("Missing arguments: Project name or environment name is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to deploy the latest environment '%s' for project '%s', are you sure?", cmdProjectEnvironment, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.DeployLatest(context.TODO(), &schema.DeployEnvironmentLatestInput{
				Environment: schema.EnvironmentInput{
					Name: cmdProjectEnvironment,
					Project: schema.ProjectInput{
						Name: cmdProjectName,
					},
				},
			}, lc)
			if err != nil {
				return err
			}
			response := result.DeployEnvironmentLatest
			if strings.HasPrefix(response, "Error: ") {
				return fmt.Errorf(strings.Trim(response, "Error: "))
			}
			fmt.Println(response)
		}
		return nil
	},
}

var deployPullrequestCmd = &cobra.Command{
	Use:     "pullrequest",
	Aliases: []string{"r"},
	Hidden:  false,
	Short:   "Deploy a pullrequest",
	Long: `Deploy a pullrequest
This pullrequest may not already exist as an environment in lagoon.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		prTitle, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}
		prNumber, err := cmd.Flags().GetUint("number")
		if err != nil {
			return err
		}
		baseBranchName, err := cmd.Flags().GetString("baseBranchName")
		if err != nil {
			return err
		}
		baseBranchRef, err := cmd.Flags().GetString("baseBranchRef")
		if err != nil {
			return err
		}
		headBranchName, err := cmd.Flags().GetString("headBranchName")
		if err != nil {
			return err
		}
		headBranchRef, err := cmd.Flags().GetString("headBranchRef")
		if err != nil {
			return err
		}
		if cmdProjectName == "" || prTitle == "" || prNumber == 0 || baseBranchName == "" ||
			baseBranchRef == "" || headBranchName == "" || headBranchRef == "" {
			return fmt.Errorf("Missing arguments: Project name, title, number, baseBranchName, baseBranchRef, headBranchName, or headBranchRef is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to deploy pull request '%v' for project '%s', are you sure?", prNumber, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)

			result, err := lagoon.DeployPullRequest(context.TODO(), &schema.DeployEnvironmentPullrequestInput{
				Project: schema.ProjectInput{
					Name: cmdProjectName,
				},
				Title:          prTitle,
				Number:         prNumber,
				BaseBranchName: baseBranchName,
				BaseBranchRef:  baseBranchRef,
				HeadBranchName: headBranchName,
				HeadBranchRef:  headBranchRef,
			}, lc)
			if err != nil {
				return err
			}
			response := result.DeployEnvironmentPullrequest
			if strings.HasPrefix(response, "Error: ") {
				return fmt.Errorf(strings.Trim(response, "Error: "))
			}
			fmt.Println(response)
		}
		return nil
	},
}

var (
	promoteSourceEnv string
	promoteDestEnv   string
)

func init() {
	deployCmd.AddCommand(deployBranchCmd)
	deployCmd.AddCommand(deployPromoteCmd)
	deployCmd.AddCommand(deployLatestCmd)
	deployCmd.AddCommand(deployPullrequestCmd)

	deployBranchCmd.Flags().StringP("branch", "b", "", "Branch name to deploy")

	deployPromoteCmd.Flags().StringP("destination", "d", "", "Destination environment name to create")
	deployPromoteCmd.Flags().StringP("source", "s", "", "Source environment name to use as the base to deploy from")

	deployPullrequestCmd.Flags().StringP("title", "t", "", "Pullrequest title")
	deployPullrequestCmd.Flags().UintP("number", "n", 0, "Pullrequest number")
	deployPullrequestCmd.Flags().StringP("baseBranchName", "N", "", "Pullrequest base branch name")
	deployPullrequestCmd.Flags().StringP("baseBranchRef", "R", "", "Pullrequest base branch reference hash")
	deployPullrequestCmd.Flags().StringP("headBranchName", "H", "", "Pullrequest head branch name")
	deployPullrequestCmd.Flags().StringP("headBranchRef", "M", "", "Pullrequest head branch reference hash")
}
