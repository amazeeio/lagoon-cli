package schema

import "github.com/amazeeio/lagoon-cli/pkg/api"

// AddEnvironmentInput is based on the input to
// addOrUpdateEnvironment.
type AddEnvironmentInput struct {
	ID                   uint           `json:"id,omitempty"`
	Name                 string         `json:"name"`
	ProjectID            uint           `json:"project"`
	DeployType           api.DeployType `json:"deployType"`
	DeployBaseRef        string         `json:"deployBaseRef"`
	DeployHeadRef        string         `json:"deployHeadRef,omitempty"`
	DeployTitle          string         `json:"deployTitle,omitempty"`
	EnvironmentType      api.EnvType    `json:"environmentType"`
	OpenshiftProjectName string         `json:"openshiftProjectName"`
}

// Environment is the Lagoon API Environment object.
type Environment struct {
	AddEnvironmentInput
	AutoIdle     uint          `json:"autoIdle"`
	EnvVariables []EnvKeyValue `json:"envVariables,omitempty"`
	Route        string        `json:"route,omitempty"`
	Routes       string        `json:"routes,omitempty"`
	// TODO use a unixtime type
	Updated string `json:"updated,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"`
	Facts   []Fact `json:"facts,omitempty"`
}

// EnvironmentConfig contains Environment configuration.
type EnvironmentConfig struct {
	Environment
	// override embedded AddEnvironmentInput.ProjectID to omitempty
	ProjectID uint `json:"project,omitempty"`
}

// EnvironmentInput  is based on the Lagoon API type.
type EnvironmentInput struct {
	ID      int          `json:"id,omitempty"`
	Name    string       `json:"name,omitempty"`
	Project ProjectInput `json:"project"`
}
