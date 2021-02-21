package lagoon

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/amazeeio/lagoon-cli/pkg/api"
)

// ErrExist indicates that an attempt was made to create an object that already
// exists.
var ErrExist = errors.New("object already exists")

// Importer interface contains methods for exporting data from Lagoon.
// TODO: compose this once simpler interfaces are defined.
type Importer interface {
	AddGroup(context.Context, *schema.AddGroupInput, *schema.Group) error
	AddUser(context.Context, *schema.AddUserInput, *schema.User) error
	AddSSHKey(context.Context, *schema.AddSSHKeyInput, *schema.SSHKey) error
	AddUserToGroup(
		context.Context, *schema.UserGroupRoleInput, *schema.Group) error
	AddNotificationSlack(context.Context,
		*schema.AddNotificationSlackInput,
		*schema.NotificationSlack) error
	AddNotificationRocketChat(context.Context,
		*schema.AddNotificationRocketChatInput,
		*schema.NotificationRocketChat) error
	AddNotificationEmail(context.Context,
		*schema.AddNotificationEmailInput,
		*schema.NotificationEmail) error
	AddNotificationMicrosoftTeams(context.Context,
		*schema.AddNotificationMicrosoftTeamsInput,
		*schema.NotificationMicrosoftTeams) error
	AddProject(context.Context, *schema.AddProjectInput, *schema.Project) error
	AddEnvVariable(
		context.Context, *schema.EnvVariableInput, *schema.EnvKeyValue) error
	ProjectByName(context.Context, string, *schema.Project) error
	AddOrUpdateEnvironment(
		context.Context, *schema.AddEnvironmentInput, *schema.Environment) error
	UpdateEnvironment(context.Context, *schema.UpdateEnvironmentInput, *schema.Environment) error
	EnvironmentByName(context.Context, string, uint, *schema.Environment) error
	AddGroupsToProject(
		context.Context, *schema.ProjectGroupsInput, *schema.Project) error
	AddNotificationToProject(context.Context,
		*schema.AddNotificationToProjectInput, *schema.Project) error
	AddBillingGroup(
		context.Context, *schema.AddBillingGroupInput, *schema.BillingGroup) error
	AddProjectToBillingGroup(context.Context, *schema.ProjectBillingGroupInput,
		*schema.Project) error
}

// Import creates objects in the Lagoon API based on a configuration object.
func Import(ctx context.Context, i Importer, r io.Reader, keepGoing bool,
	openshiftID uint) error {

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("couldn't read file: %w", err)
	}

	config := schema.Config{}
	if err = schema.UnmarshalConfigYAML(data, &config); err != nil {
		return fmt.Errorf("couldn't unmarshal config: %w", err)
	}

	// import the config
	l := log.New(os.Stderr, "import: ", 0)
	// add billing groups
	for _, bg := range config.BillingGroups {
		if err := i.AddBillingGroup(ctx, &bg, nil); err != nil {
			if !keepGoing {
				return fmt.Errorf("couldn't add billing group: %w", err)
			}
			l.Printf("couldn't add billing group: %v", err)
		}
	}
	// add groups
	for _, group := range config.Groups {
		if err := i.AddGroup(ctx, &group.AddGroupInput, nil); err != nil {
			if !keepGoing {
				return fmt.Errorf("couldn't add group: %w", err)
			}
			l.Printf("couldn't add group: %v", err)
		}
	}
	// add users
	for _, user := range config.Users {
		if err := i.AddUser(ctx, &user.AddUserInput, nil); err != nil {
			if !keepGoing {
				return fmt.Errorf("couldn't add user: %w", err)
			}
			l.Printf("couldn't add user: %v", err)
		}
	}
	// add ssh-keys to users
	for _, user := range config.Users {
		for _, sshKey := range user.SSHKeys {
			err := i.AddSSHKey(ctx, &schema.AddSSHKeyInput{
				SSHKey:    sshKey,
				UserEmail: user.Email,
			}, nil)
			if err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add SSH key: %w", err)
				}
				l.Printf("couldn't add SSH key: %v", err)
			}
		}
	}
	// add users to groups
	for _, group := range config.Groups {
		for _, userRole := range group.Users {
			err := i.AddUserToGroup(ctx, &schema.UserGroupRoleInput{
				UserEmail: userRole.Email,
				GroupName: group.Name,
				GroupRole: userRole.Role,
			}, nil)
			if err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add user to group: %w", err)
				}
				l.Printf("couldn't add user to group: %v", err)
			}
		}
	}
	if config.Notifications != nil {
		// add Slack notifications
		for _, n := range config.Notifications.Slack {
			if err := i.AddNotificationSlack(ctx, &n, nil); err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add Slack notification: %w", err)
				}
				l.Printf("couldn't add Slack notification: %v", err)
			}
		}
		// add RocketChat notifications
		for _, n := range config.Notifications.RocketChat {
			if err := i.AddNotificationRocketChat(ctx, &n, nil); err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add RocketChat notification: %w", err)
				}
				l.Printf("couldn't add RocketChat notification: %v", err)
			}
		}
		// add Email notifications
		for _, n := range config.Notifications.Email {
			if err := i.AddNotificationEmail(ctx, &n, nil); err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add Email notification: %w", err)
				}
				l.Printf("couldn't add Email notification: %v", err)
			}
		}
		// add MicrosoftTeams notifications
		for _, n := range config.Notifications.MicrosoftTeams {
			if err := i.AddNotificationMicrosoftTeams(ctx, &n, nil); err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add MicrosoftTeams notification: %w", err)
				}
				l.Printf("couldn't add MicrosoftTeams notification: %v", err)
			}
		}
	}
	// add projects
	newProj := schema.Project{}
	for _, p := range config.Projects {
		p.Openshift = openshiftID
		if err := i.AddProject(ctx, &p.AddProjectInput, &newProj); err != nil {
			if errors.Is(err, ErrExist) {
				// this project exists already
				if !keepGoing {
					return fmt.Errorf("project exists: %w", err)
				}
				if err = i.ProjectByName(ctx, p.Name, &newProj); err != nil {
					return fmt.Errorf(
						`couldn't get project "%s" by name: %w`, p.Name, err)
				}
				l.Printf(`project "%s" exists, using ID %d`, p.Name, newProj.ID)
			} else {
				if !keepGoing {
					return fmt.Errorf("couldn't add Project: %w", err)
				}
				l.Printf("couldn't add Project: %v", err)
				continue // next project
			}
		}
		// add project env-vars
		for _, ev := range p.EnvVariables {
			err := i.AddEnvVariable(ctx, &schema.EnvVariableInput{
				EnvKeyValue: ev,
				Type:        api.ProjectVar,
				TypeID:      newProj.ID,
			}, nil)
			if err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add Project EnvVariable: %w", err)
				}
				l.Printf("couldn't add Project EnvVariable: %v", err)
			}
		}
		// add project environments
		for _, env := range p.Environments {
			newEnv := schema.Environment{}
			// inject project ID
			env.Environment.AddEnvironmentInput.ProjectID = newProj.ID
			err := i.AddOrUpdateEnvironment(
				ctx, &env.Environment.AddEnvironmentInput, &newEnv)
			if errors.Is(err, ErrExist) {
				// this environment exists already
				if !keepGoing {
					return fmt.Errorf("environment exists: %w", err)
				}
				l.Printf(`environment "%s" (project "%s") exists, query by name for ID`,
					env.Name, p.Name)
				err = i.EnvironmentByName(ctx, env.Name, env.ProjectID, &newEnv)
				if err != nil {
					return fmt.Errorf("couldn't get environment by name: %w", err)
				}
			} else if err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add Environment: %w", err)
				}
				l.Printf("couldn't add Environment: %v", err)
				continue // next environment
			}
			// Set the correct created date
			envUpdateInput := schema.UpdateEnvironmentInput{
				ID: newEnv.ID,
				Patch: schema.UpdateEnvironmentPatch{
					Created: env.Created,
				},
			}
			err = i.UpdateEnvironment(ctx, &envUpdateInput, nil)
			if err != nil {
				return fmt.Errorf("couldn't update Environment: %w", err)
			}
			// add environment env-vars
			for _, ev := range env.EnvVariables {
				err = i.AddEnvVariable(ctx, &schema.EnvVariableInput{
					EnvKeyValue: ev,
					Type:        api.EnvironmentVar,
					TypeID:      newEnv.ID,
				}, nil)
				if err != nil {
					if !keepGoing {
						return fmt.Errorf("couldn't add Environment EnvVariable: %w", err)
					}
					l.Printf("couldn't add Environment EnvVariable: %v", err)
				}
			}
		}
		// add groups to project
		if len(p.Groups) > 0 {
			// convert group names to input type
			groupsInput := []schema.GroupInput{}
			for _, name := range p.Groups {
				groupsInput = append(groupsInput, schema.GroupInput{Name: name})
			}
			err = i.AddGroupsToProject(ctx, &schema.ProjectGroupsInput{
				Project: schema.ProjectInput{Name: p.Name},
				Groups:  groupsInput}, nil)
			if err != nil {
				if !keepGoing {
					return fmt.Errorf(
						`couldn't add Groups to Project "%s": %w`, p.Name, err)
				}
				l.Printf(`couldn't add Groups to Project "%s": %v`, p.Name, err)
			}
		}
		if len(p.BillingGroups) > 1 {
			return fmt.Errorf(
				`project can only have one billing group: %v`, p.BillingGroups)
		}
		// add project to billing group
		for _, bgName := range p.BillingGroups {
			err = i.AddProjectToBillingGroup(ctx, &schema.ProjectBillingGroupInput{
				Group:   schema.GroupInput{Name: bgName},
				Project: schema.ProjectInput{Name: p.Name},
			}, nil)
			if err != nil {
				if !keepGoing {
					return fmt.Errorf(
						`couldn't add Project "%s" to Billing Group "%s": %w`, p.Name,
						bgName, err)
				}
				l.Printf(`couldn't add Project "%s" to Billing Group "%s": %v`, p.Name,
					bgName, err)
			}
		}
		// add project users
		for _, u := range p.Users {
			err := i.AddUserToGroup(ctx, &schema.UserGroupRoleInput{
				UserEmail: u.Email,
				GroupName: fmt.Sprintf(`project-%s`, p.Name),
				GroupRole: u.Role,
			}, nil)
			if err != nil {
				if !keepGoing {
					return fmt.Errorf("couldn't add user to project group: %w", err)
				}
				l.Printf("couldn't add user to project group: %v", err)
			}
		}
		// add project notifications
		if p.Notifications != nil {
			for _, n := range p.Notifications.Slack {
				err := i.AddNotificationToProject(ctx,
					&schema.AddNotificationToProjectInput{
						Project:          p.Name,
						NotificationType: api.SlackNotification,
						NotificationName: n,
					}, nil)
				if err != nil {
					if !keepGoing {
						return fmt.Errorf(
							"couldn't add Slack Notification to project: %w", err)
					}
					l.Printf("couldn't add Slack Notification to project: %v", err)
				}
			}
			for _, n := range p.Notifications.RocketChat {
				err := i.AddNotificationToProject(ctx,
					&schema.AddNotificationToProjectInput{
						Project:          p.Name,
						NotificationType: api.RocketChatNotification,
						NotificationName: n,
					}, nil)
				if err != nil {
					if !keepGoing {
						return fmt.Errorf(
							"couldn't add RocketChat Notification to project: %w", err)
					}
					l.Printf("couldn't add RocketChat Notification to project: %v", err)
				}
			}
			for _, n := range p.Notifications.Email {
				err := i.AddNotificationToProject(ctx,
					&schema.AddNotificationToProjectInput{
						Project:          p.Name,
						NotificationType: api.EmailNotification,
						NotificationName: n,
					}, nil)
				if err != nil {
					if !keepGoing {
						return fmt.Errorf(
							"couldn't add Email Notification to project: %w", err)
					}
					l.Printf("couldn't add Email Notification to project: %v", err)
				}
			}
			for _, n := range p.Notifications.MicrosoftTeams {
				err := i.AddNotificationToProject(ctx,
					&schema.AddNotificationToProjectInput{
						Project:          p.Name,
						NotificationType: api.MicrosoftTeamsNotification,
						NotificationName: n,
					}, nil)
				if err != nil {
					if !keepGoing {
						return fmt.Errorf(
							"couldn't add MicrosoftTeams Notification to project: %w", err)
					}
					l.Printf("couldn't add MicrosoftTeams Notification to project: %v", err)
				}
			}
		}
	}

	return nil
}
