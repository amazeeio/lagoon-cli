// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Deploy interface contains methods for deploying branches and environments in lagoon.
type Facts interface {
	ProjectByName(ctx context.Context, name string, project *schema.Project) error
	FactsforEnvironment(ctx context.Context, projectId uint, environmentName string, facts *[]schema.Fact) error
	AddFact(ctx context.Context, environmentId uint, name string, value string, fact *schema.Fact) error
	DeleteFact(ctx context.Context, environmentId uint, name string, ret *string) error
}

// GetProjectByNameForFacts gets project by name for the context of facts
func GetProjectByNameForFacts(ctx context.Context, projectName string, f Facts) (*schema.Project, error) {
	project := schema.Project{}
	return &project, f.ProjectByName(ctx, projectName, &project)
}

func GetEnvironmentFacts(ctx context.Context, projectId uint, environmentName string, f Facts) (*[]schema.Fact, error) {
	facts := []schema.Fact{}
	ret := f.FactsforEnvironment(ctx, projectId, environmentName, &facts)
	return &facts, ret
}

func AddFact(ctx context.Context, environmentId uint, name string, value string, f Facts) (*schema.Fact, error) {
	fact := schema.Fact{}
	err := f.AddFact(ctx, environmentId, name, value, &fact)
	return &fact, err
}

func DeleteFact(ctx context.Context, environmentId uint, name string, f Facts) (string, error) {
	var ret string
	err := f.DeleteFact(ctx, environmentId, name, &ret)
	return ret, err
}

func FactExists(ctx context.Context, projectId uint, environmentName string, name string, f Facts) (bool, error) {
	facts := []schema.Fact{}
	err := f.FactsforEnvironment(ctx, projectId, environmentName, &facts)
	if err != nil {
		return false, err
	}

	for _, fact := range facts {
		if fact.Name == name {
			return true, nil
		}
	}

	return false, nil
}
