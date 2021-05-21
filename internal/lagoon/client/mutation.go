package client

import (
	"context"
	"fmt"
	"regexp"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/schema"
)

var duplicate = regexp.MustCompile("^graphql: Duplicate entry ")

// wrapErr wraps a response error with a lagoon.ErrExist type if the response
// is due to an object already existing
func wrapErr(err error) error {
	if err != nil && duplicate.MatchString(err.Error()) {
		return fmt.Errorf("couldn't create object: %w: %v", lagoon.ErrExist, err)
	}
	return err
}

// AddGroup adds a group.
func (c *Client) AddGroup(
	ctx context.Context, in *schema.AddGroupInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/addGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addGroup"`
	}{
		Response: out,
	})
}

// AddUser adds a user.
func (c *Client) AddUser(
	ctx context.Context, in *schema.AddUserInput, out *schema.User) error {
	req, err := c.newRequest("_lgraphql/addUser.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"addUser"`
	}{
		Response: out,
	})
}

// AddSSHKey adds an SSH key to a user.
func (c *Client) AddSSHKey(
	ctx context.Context, in *schema.AddSSHKeyInput, out *schema.SSHKey) error {
	req, err := c.newRequest("_lgraphql/addSshKey.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.SSHKey `json:"addSshKey"`
	}{
		Response: out,
	})
}

// AddUserToGroup adds a user to a group.
func (c *Client) AddUserToGroup(
	ctx context.Context, in *schema.UserGroupRoleInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/addUserToGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addUserToGroup"`
	}{
		Response: out,
	})
}

// AddNotificationSlack defines a Slack notification.
func (c *Client) AddNotificationSlack(ctx context.Context,
	in *schema.AddNotificationSlackInput, out *schema.NotificationSlack) error {
	req, err := c.newRequest("_lgraphql/addNotificationSlack.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationSlack `json:"addNotificationSlack"`
	}{
		Response: out,
	})
}

// AddNotificationRocketChat defines a RocketChat notification.
func (c *Client) AddNotificationRocketChat(ctx context.Context,
	in *schema.AddNotificationRocketChatInput,
	out *schema.NotificationRocketChat) error {
	req, err := c.newRequest("_lgraphql/addNotificationRocketChat.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationRocketChat `json:"addNotificationRocketChat"`
	}{
		Response: out,
	})
}

// AddNotificationEmail defines an Email notification.
func (c *Client) AddNotificationEmail(ctx context.Context,
	in *schema.AddNotificationEmailInput,
	out *schema.NotificationEmail) error {
	req, err := c.newRequest("_lgraphql/addNotificationEmail.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationEmail `json:"addNotificationEmail"`
	}{
		Response: out,
	})
}

// AddNotificationMicrosoftTeams defines a MicrosoftTeams notification.
func (c *Client) AddNotificationMicrosoftTeams(ctx context.Context,
	in *schema.AddNotificationMicrosoftTeamsInput,
	out *schema.NotificationMicrosoftTeams) error {
	req, err := c.newRequest("_lgraphql/addNotificationMicrosoftTeams.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationMicrosoftTeams `json:"addNotificationMicrosoftTeams"`
	}{
		Response: out,
	})
}

// AddProject adds a project.
func (c *Client) AddProject(
	ctx context.Context, in *schema.AddProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addProject.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addProject"`
	}{
		Response: out,
	}))
}

// AddEnvVariable adds an EnvVariable to an Environment or Project.
func (c *Client) AddEnvVariable(ctx context.Context,
	in *schema.EnvVariableInput, out *schema.EnvKeyValue) error {
	req, err := c.newRequest("_lgraphql/addEnvVariable.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.EnvKeyValue `json:"addEnvVariable"`
	}{
		Response: out,
	})
}

// AddOrUpdateEnvironment adds or updates a Project Environment.
func (c *Client) AddOrUpdateEnvironment(ctx context.Context,
	in *schema.AddEnvironmentInput, out *schema.Environment) error {
	req, err := c.newRequest("_lgraphql/addOrUpdateEnvironment.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"addOrUpdateEnvironment"`
	}{
		Response: out,
	}))
}

// AddGroupsToProject adds Groups to a Project.
func (c *Client) AddGroupsToProject(ctx context.Context,
	in *schema.ProjectGroupsInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addGroupsToProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addGroupsToProject"`
	}{
		Response: out,
	})
}

// AddNotificationToProject adds a Notification to a Project.
func (c *Client) AddNotificationToProject(ctx context.Context,
	in *schema.AddNotificationToProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addNotificationToProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addNotificationToProject"`
	}{
		Response: out,
	})
}

// AddBillingGroup adds a Billing Group.
func (c *Client) AddBillingGroup(ctx context.Context,
	in *schema.AddBillingGroupInput, out *schema.BillingGroup) error {
	req, err := c.newRequest("_lgraphql/addBillingGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.BillingGroup `json:"addBillingGroup"`
	}{
		Response: out,
	})
}

// AddProjectToBillingGroup adds a Project to a Billing Group.
func (c *Client) AddProjectToBillingGroup(ctx context.Context,
	in *schema.ProjectBillingGroupInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addProjectToBillingGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addProjectToBillingGroup"`
	}{
		Response: out,
	})
}

// AddFact adds an Fact to an Environment.
func (c *Client) AddFact(
	ctx context.Context, environmentId uint, name string, value string, fact *schema.Fact) error {

	req, err := c.newVersionedRequest("_lgraphql/addFact.graphql",
		map[string]interface{}{
			"environment": environmentId,
			"name":        name,
			"value":       value,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Fact `json:"addFact"`
	}{
		Response: fact,
	})
}

// DeleteFact deletes a Fact from an Environment.
func (c *Client) DeleteFact(
	ctx context.Context, environmentId uint, name string, ret *string) error {

	req, err := c.newVersionedRequest("_lgraphql/deleteFact.graphql",
		map[string]interface{}{
			"environment": environmentId,
			"name":        name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response string `json:"data"`
	}{
		Response: *ret,
	})
}

// DeployEnvironmentLatest deploys a latest environment.
func (c *Client) DeployEnvironmentLatest(ctx context.Context,
	in *schema.DeployEnvironmentLatestInput, out *schema.DeployEnvironmentLatest) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentLatest.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPullrequest deploys a pullreguest.
func (c *Client) DeployEnvironmentPullrequest(ctx context.Context,
	in *schema.DeployEnvironmentPullrequestInput, out *schema.DeployEnvironmentPullrequest) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentPullrequest.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPromote promotes one environment into a new environment.
func (c *Client) DeployEnvironmentPromote(ctx context.Context,
	in *schema.DeployEnvironmentPromoteInput, out *schema.DeployEnvironmentPromote) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentPromote.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentBranch deploys a branch.
func (c *Client) DeployEnvironmentBranch(ctx context.Context,
	in *schema.DeployEnvironmentBranchInput, out *schema.DeployEnvironmentBranch) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentBranch.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// RunActiveStandbySwitch deploys a branch.
func (c *Client) RunActiveStandbySwitch(ctx context.Context,
	project string, out *schema.Task) error {
	req, err := c.newVersionedRequest("_lgraphql/switchActiveStandby.graphql", map[string]interface{}{
		"project": project,
	})
	if err != nil {
		return err
	}

	// return c.client.Run(ctx, req, &out)
	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"switchActiveStandby"`
	}{
		Response: out,
	})
}
