package cmd

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listOpenshifts = &cobra.Command{
	Use:     "openshifts",
	Aliases: []string{"o", "os"},
	Short:   "List all Openshifts Lagoon knows about (platform admin user permissions only)",
	Long: `List all Openshifts Lagoon knows about (platform admin user permissions only)
Only platform admin role can list openshifts.
NOTE: only openshifts that are used by projects can be listed`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		fields, err := cmd.Flags().GetStringSlice("fields")
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
		openshifts, err := lagoon.GetAllOpenshifts(context.TODO(), lc)
		if err != nil {
			return err
		}
		header := []string{}
		sortedFields := []string{}
		sortFields(listOpenshiftFields, &sortedFields)
		sort.Strings(fields) // sort alphabetically before we move id,name to first positions
		// set id to position 0 and name to position 1
		fields = alwaysShowFields(fields, map[int]string{
			0: "id",
			1: "name",
		})
		// loop through the supplied flag fields and check if they exist in the returned schema
		// add append to the header
		for _, opt := range fields {
			for _, k := range sortedFields {
				if strings.ToLower(opt) == strings.ToLower(k) {
					header = append(header, fmt.Sprintf("%v", listOpenshiftFields[strings.ToLower(opt)]))
				}
			}
		}
		data := []output.Data{}
		// loop through the supplied flag fields and check if they exist in the returned schema
		// add append to the data
		for _, openshift := range *openshifts {
			mapData := []string{}
			for _, opt := range fields {
				for _, k := range sortedFields {
					if strings.ToLower(opt) == strings.ToLower(k) {
						fieldVal := reflect.ValueOf(&openshift).Elem().FieldByName(listOpenshiftFields[strings.ToLower(opt)])
						mapData = append(mapData, returnNonEmptyString(fmt.Sprintf("%v", fieldVal)))
					}
				}
			}
			data = append(data, mapData)
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i][0] < data[j][0]
		})
		output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		return nil
	},
}

var addOpenshift = &cobra.Command{
	Use:     "openshift",
	Short:   "Add a new openshift (platform admin user permissions only)",
	Long:    `Add a new openshift (platform admin user permissions only)`,
	Aliases: []string{"o", "os"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		consoleURL, err := cmd.Flags().GetString("consoleUrl")
		if err != nil {
			return err
		}
		ID, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}
		osToken, err := cmd.Flags().GetString("token")
		if err != nil {
			return err
		}
		projectUser, err := cmd.Flags().GetString("projectUser")
		if err != nil {
			return err
		}
		routerPattern, err := cmd.Flags().GetString("routerPattern")
		if err != nil {
			return err
		}
		sshHost, err := cmd.Flags().GetString("sshHost")
		if err != nil {
			return err
		}
		sshPort, err := cmd.Flags().GetString("sshPort")
		if err != nil {
			return err
		}
		monitoringConfig, err := cmd.Flags().GetString("monitoringConfig")
		if err != nil {
			return err
		}
		if name == "" {
			return fmt.Errorf("Missing arguments: name is not defined")
		}
		if consoleURL == "" {
			return fmt.Errorf("Missing arguments: consoleUrl is not defined")
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)
		openshift := &schema.AddOpenshiftInput{
			ID: ID,
			UpdateOpenshiftPatchInput: schema.UpdateOpenshiftPatchInput{
				Name:          name,
				ConsoleURL:    consoleURL,
				Token:         osToken,
				ProjectUser:   projectUser,
				RouterPattern: routerPattern,
				SSHHost:       sshHost,
				SSHPort:       sshPort,
			},
		}
		// @TODO: supporting versioned schemas might become a bit tricky
		if greaterThanOrEqualVersion(viper.GetString("lagoons."+current+".version"), "v1.7.0") {
			openshift.UpdateOpenshiftPatchInput.MonitoringConfig = monitoringConfig
		}
		result, err := lagoon.AddOpenshift(context.TODO(), openshift, lc)
		if err != nil {
			return err
		}
		output.RenderResult(output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"OpenshiftID":   result.ID,
				"OpenshiftName": result.Name,
			},
		}, outputOptions)
		return nil
	},
}

var updateOpenshift = &cobra.Command{
	Use:     "openshift",
	Short:   "Update an openshift (platform admin user permissions only)",
	Long:    `Update an openshift (platform admin user permissions only)`,
	Aliases: []string{"o", "os"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		ID, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		consoleURL, err := cmd.Flags().GetString("consoleUrl")
		if err != nil {
			return err
		}
		osToken, err := cmd.Flags().GetString("token")
		if err != nil {
			return err
		}
		projectUser, err := cmd.Flags().GetString("projectUser")
		if err != nil {
			return err
		}
		routerPattern, err := cmd.Flags().GetString("routerPattern")
		if err != nil {
			return err
		}
		sshHost, err := cmd.Flags().GetString("sshHost")
		if err != nil {
			return err
		}
		sshPort, err := cmd.Flags().GetString("sshPort")
		if err != nil {
			return err
		}
		monitoringConfig, err := cmd.Flags().GetString("monitoringConfig")
		if err != nil {
			return err
		}
		if ID == 0 {
			return fmt.Errorf("Missing arguments: ID is not defined")
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)
		openshift := &schema.UpdateOpenshiftInput{
			ID: ID,
			Patch: schema.UpdateOpenshiftPatchInput{
				Name:          name,
				ConsoleURL:    consoleURL,
				Token:         osToken,
				ProjectUser:   projectUser,
				RouterPattern: routerPattern,
				SSHHost:       sshHost,
				SSHPort:       sshPort,
			},
		}
		// @TODO: supporting versioned schemas might become a bit tricky
		if greaterThanOrEqualVersion(viper.GetString("lagoons."+current+".version"), "v1.7.0") {
			openshift.Patch.MonitoringConfig = monitoringConfig
		}
		result, err := lagoon.UpdateOpenshift(context.TODO(), openshift, lc)
		if err != nil {
			return err
		}
		output.RenderResult(output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"OpenshiftID":   result.ID,
				"OpenshiftName": result.Name,
			},
		}, outputOptions)
		return nil
	},
}

var deleteOpenshift = &cobra.Command{
	Use:     "openshift",
	Short:   "Delete an openshift (platform admin user permissions only)",
	Long:    `Delete an openshift (platform admin user permissions only)`,
	Aliases: []string{"o", "os"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if name == "" {
			return fmt.Errorf("Missing arguments: name is not defined")
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)
		result, err := lagoon.DeleteOpenshift(context.TODO(), &schema.DeleteOpenshiftInput{
			Name: name,
		}, lc)
		if err != nil {
			return err
		}
		fmt.Println(result.DeleteOpenshift)
		return nil
	},
}

// this map is to reference a `--field` flag slice name to a field in returned schema
// since the returned structs in the lagoon-cli can use embedded structs sometimes
// we can define the fields we want to pick from a possible embedded struct to a specific named field
var listOpenshiftFields = map[string]string{
	"id":               "ID",
	"name":             "Name",
	"consoleurl":       "ConsoleURL",
	"routerpattern":    "RouterPattern",
	"projectuser":      "ProjectUser",
	"sshhost":          "SSHHost",
	"sshport":          "SSHPort",
	"created":          "Created",
	"token":            "Token",
	"monitoringconfig": "MonitoringConfig",
}

func init() {
	listCmd.AddCommand(listOpenshifts)

	sortedFields := []string{}
	sortFields(listOpenshiftFields, &sortedFields)
	listOpenshifts.Flags().StringSlice("fields", []string{"id", "name"},
		`Select which fields to display when listing Openshifts. Valid options (others are ignored): `+strings.Join(sortedFields, ","))

	addCmd.AddCommand(addOpenshift)
	addOpenshift.Flags().StringP("name", "N", "",
		"Name of the Openshift to be used in Lagoon")
	addOpenshift.Flags().StringP("consoleUrl", "C", "",
		"ConsoleURL of the openshift")
	addOpenshift.Flags().StringP("token", "T", "",
		"Openshift token used to access the cluster")
	addOpenshift.Flags().IntP("id", "I", 0,
		"ID to assign the Openshift in Lagoon")
	addOpenshift.Flags().StringP("routerPattern", "R", "",
		"Router pattern to use")
	addOpenshift.Flags().StringP("projectUser", "P", "",
		"Project user to use")
	addOpenshift.Flags().StringP("sshHost", "S", "",
		"SSH host address")
	addOpenshift.Flags().StringP("sshPort", "s", "",
		"SSH port number")
	addOpenshift.Flags().StringP("monitoringConfig", "m", "",
		"Configuration for monitoring (JSON)")

	updateCmd.AddCommand(updateOpenshift)
	updateOpenshift.Flags().StringP("name", "N", "",
		"Name of the Openshift in Lagoon")
	updateOpenshift.Flags().StringP("consoleUrl", "C", "",
		"ConsoleURL of the openshift")
	updateOpenshift.Flags().StringP("token", "T", "",
		"Openshift token used to access the cluster")
	updateOpenshift.Flags().IntP("id", "I", 0,
		"ID to assign the Openshift in Lagoon")
	updateOpenshift.Flags().StringP("routerPattern", "R", "",
		"Router pattern to use")
	updateOpenshift.Flags().StringP("projectUser", "P", "",
		"Project user to use")
	updateOpenshift.Flags().StringP("sshHost", "S", "",
		"SSH host address")
	updateOpenshift.Flags().StringP("sshPort", "s", "",
		"SSH port number")
	updateOpenshift.Flags().StringP("monitoringConfig", "m", "",
		"Configuration for monitoring (JSON)")

	deleteCmd.AddCommand(deleteOpenshift)
	deleteOpenshift.Flags().StringP("name", "N", "",
		"Name of the Openshift in Lagoon")
}
