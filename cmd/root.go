package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/amazeeio/lagoon-cli/app"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/manifoldco/promptui"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

var cmdProject app.LagoonProject
var cmdLagoon = ""
var forceAction bool
var cmdSSHKey = ""
var inputScanner = bufio.NewScanner(os.Stdin)
var versionFlag bool
var docsFlag bool

var rootCmd = &cobra.Command{
	Use:               "lagoon",
	Short:             "Command line integration for Lagoon",
	Long:              `Lagoon CLI. Manage your Lagoon hosted projects.`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		if docsFlag {
			err := doc.GenMarkdownTree(cmd, "docs/commands")
			if err != nil {
				log.Fatal(err)
			}
		}
		if versionFlag {
			displayVersionInfo()
		}
		cmd.Help()
		os.Exit(1)
	},
}

// version/build information
var (
	version string
	build   string
)

// Execute the root command.
func Execute() {
	viper.AutomaticEnv()
	if err := rootCmd.Execute(); err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.EnableCommandSorting = false

	rootCmd.PersistentFlags().StringVarP(&cmdProjectName, "project", "p", "", "Specify a project to use")
	rootCmd.PersistentFlags().StringVarP(&cmdProjectEnvironment, "environment", "e", "", "Specify an environment to use")

	rootCmd.PersistentFlags().StringVarP(&cmdLagoon, "lagoon", "l", "", "The Lagoon instance to interact with")
	rootCmd.PersistentFlags().BoolVarP(&forceAction, "force", "", false, "Force (if supported)")
	rootCmd.PersistentFlags().StringVarP(&cmdSSHKey, "ssh-key", "i", "", "Specify a specific SSH key to use")

	rootCmd.PersistentFlags().BoolVarP(&listAllProjects, "all-projects", "", false, "All projects (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.Header, "no-header", "", false, "No header on table (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.CSV, "output-csv", "", false, "Output as CSV (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.JSON, "output-json", "", false, "Output as JSON (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.Pretty, "pretty", "", false, "Make JSON pretty (if supported)")

	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "", false, "Version information")
	rootCmd.Flags().BoolVarP(&docsFlag, "docs", "", false, "Generate docs")

	rootCmd.Flags().MarkHidden("docs")

	rootCmd.SetUsageTemplate(`Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}

{{- if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
{{- $current_command:= . }}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}

{{if .HasAvailableLocalFlags -}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(devCmd)
	rootCmd.AddCommand(sshEnvCmd)
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(kibanaCmd)
	rootCmd.AddCommand(versionCmd)
}

// version/build information command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		displayVersionInfo()
	},
}

func displayVersionInfo() {
	fmt.Println("Version:", version)
	fmt.Println("Build:", build)
	os.Exit(0)
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}

	var configName = ".lagoon"

	// Search config in home directory with name ".lagoon" (without extension).
	// @todo see if we can grok the proper info from the cwd .lagoon.yml
	viper.AddConfigPath(home)
	viper.SetConfigName(configName)
	viper.SetDefault("lagoons.amazeeio.hostname", "ssh.lagoon.amazeeio.cloud")
	viper.SetDefault("lagoons.amazeeio.port", 32222)
	viper.SetDefault("lagoons.amazeeio.token", "")
	viper.SetDefault("lagoons.amazeeio.graphql", "https://api.lagoon.amazeeio.cloud/graphql")
	viper.SetDefault("lagoons.amazeeio.ui", "https://ui-lagoon-master.ch.amazee.io")
	viper.SetDefault("lagoons.amazeeio.kibana", "https://logs-db-ui-lagoon-master.ch.amazee.io/")
	viper.SetDefault("default", "amazeeio")
	err = viper.ReadInConfig()
	if err != nil {
		err = viper.WriteConfigAs(filepath.Join(home, configName+".yml"))
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
	}
	if cmdLagoon == "" {
		if viper.GetString("default") == "" {
			cmdLagoon = "amazeeio"
		} else {
			cmdLagoon = viper.GetString("default")
		}
	}
	viper.Set("current", strings.TrimSpace(string(cmdLagoon))) // set the current lagoon to whatever we defined from config or as override in a flag

	err = viper.WriteConfig()
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}

	// if the directory or repository you're in has a valid .lagoon.yml and docker-compose.yml with x-lagoon-project in it
	// we can use that inplaces where projects already exist so you don't have to type it out
	// and environments too
	cmdProject, _ = app.GetLocalProject()
	if cmdProject.Name != "" && cmdProjectName == "" {
		cmdProjectName = cmdProject.Name
	}
	if cmdProject.Environment != "" && cmdProjectEnvironment == "" {
		cmdProjectEnvironment = cmdProject.Environment
	}

	// if !outputOptions.CSV && !outputOptions.JSON {
	// 	fmt.Println("Using Lagoon:", cmdLagoon)
	// }

}

func yesNo() bool {
	if forceAction != true {
		prompt := promptui.Select{
			Label: "Select[Yes/No]",
			Items: []string{"No", "Yes"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		return result == "Yes"
	}
	return true
}

func selectList(listItems []string) string {
	if forceAction != true {
		prompt := promptui.Select{
			Label: "Select item",
			Items: listItems,
		}
		_, result, err := prompt.Run()
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		return result
	}
	return ""
}

// GetInput reads input from an input buffer and returns the result as a string.
func GetInput() string {
	inputScanner.Scan()
	return strings.TrimSpace(inputScanner.Text())
}

// Prompt gets input with a prompt and returns the input
func Prompt(prompt string) string {
	fullPrompt := fmt.Sprintf("%s", prompt)
	fmt.Print(fullPrompt + ": ")
	return GetInput()
}

func unset(key string) error {
	delete(viper.Get("lagoons").(map[string]interface{}), key)
	err := viper.WriteConfig()
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	return nil
}

// FormatType .
type FormatType string

// . .
const (
	JSON   FormatType = "JSON"
	YAML   FormatType = "YAML"
	STDOUT FormatType = "STDOUT"
)

func validateToken(lagoon string) {
	valid := graphql.VerifyTokenExpiry(lagoon)
	if valid == false {
		loginErr := loginToken()
		if loginErr != nil {
			fmt.Println("Unable to refresh token, you may need to run `lagoon login` first, error was", loginErr.Error())
			os.Exit(1)
		}
	}
}
