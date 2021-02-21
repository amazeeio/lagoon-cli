// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Projects interface contains methods for getting info on projects.
type Projects interface {
	MinimalProjectByName(ctx context.Context, name string, project *schema.Project) error
	ProjectsByMetadata(ctx context.Context, key string, value string, project *[]schema.ProjectMetadata) error
	UpdateProjectMetadata(ctx context.Context, id int, key string, value string, project *schema.ProjectMetadata) error
	RemoveProjectMetadataByKey(ctx context.Context, id int, key string, project *schema.ProjectMetadata) error
}

// GetMinimalProjectByName gets info of projects in lagoon that have matching metadata.
func GetMinimalProjectByName(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.MinimalProjectByName(ctx, name, &project)
}

// GetProjectsByMetadata gets info of projects in lagoon that have matching metadata.
func GetProjectsByMetadata(ctx context.Context, key string, value string, p Projects) (*[]schema.ProjectMetadata, error) {
	project := []schema.ProjectMetadata{}
	return &project, p.ProjectsByMetadata(ctx, key, value, &project)
}

// UpdateProjectMetadata updates a project with provided metadata.
func UpdateProjectMetadata(ctx context.Context, id int, key string, value string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.UpdateProjectMetadata(ctx, id, key, value, &project)
}

// RemoveProjectMetadataByKey remove metadata from a project by key.
func RemoveProjectMetadataByKey(ctx context.Context, id int, key string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.RemoveProjectMetadataByKey(ctx, id, key, &project)
}
