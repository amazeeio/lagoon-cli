package client

import (
	"context"
	"fmt"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// ProjectByName queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectByName(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newVersionedRequest("_lgraphql/projectByName.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// Me queries the Lagoon API for me, and
// unmarshals the response into project.
func (c *Client) Me(
	ctx context.Context, user *schema.User) error {

	req, err := c.newRequest("_lgraphql/me.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"me"`
	}{
		Response: user,
	})
}

func (c *Client) FactsforEnvironment(ctx context.Context, projectId uint, environmentName string, facts *[]schema.Fact) error {
	req, err := c.newRequest("_lgraphql/getFacts.graphql", map[string]interface{}{
		"name":      environmentName,
		"projectID": projectId,
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var environment schema.Environment

	ret := c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: &environment,
	})

	if ret != nil {
		fmt.Println("ERROR")
	}

	*facts = environment.Facts
	return ret

}

// EnvironmentByName queries the Lagoon API for an environment by its name and
// parent projectID, and unmarshals the response into environment.
func (c *Client) EnvironmentByName(ctx context.Context, name string,
	projectID uint, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environmentByName.graphql",
		map[string]interface{}{
			"name":    name,
			"project": projectID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// LagoonAPIVersion queries the Lagoon API for its version, and
// unmarshals the response.
func (c *Client) LagoonAPIVersion(
	ctx context.Context, lagoonAPIVersion *schema.LagoonVersion) error {

	req, err := c.newRequest("_lgraphql/lagoonVersion.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &lagoonAPIVersion)
}

// LagoonSchema queries the Lagoon API for its schema, and
// unmarshals the response.
func (c *Client) LagoonSchema(
	ctx context.Context, lagoonSchema *schema.LagoonSchema) error {

	req, err := c.newRequest("_lgraphql/lagoonSchema.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.LagoonSchema `json:"__schema"`
	}{
		Response: lagoonSchema,
	})
}

// GetTaskByID queries the Lagoon API for a task by its ID, and
// unmarshals the response.
func (c *Client) GetTaskByID(
	ctx context.Context, id int, task *schema.Task) error {

	req, err := c.newVersionedRequest("_lgraphql/taskByID.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"taskById"`
	}{
		Response: task,
	})
}
