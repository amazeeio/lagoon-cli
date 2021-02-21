// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"fmt"
	"strings"

	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/hashicorp/go-version"
)

// APIVersion interface contains methods for getting info on the current version of lagoon.
type APIVersion interface {
	LagoonAPIVersion(ctx context.Context, apiVersion *schema.LagoonVersion) error
	LagoonSchema(ctx context.Context, lagoonSchema *schema.LagoonSchema) error
}

// GetLagoonAPIVersion gets info on the current API version of lagoon, supported in lagoon v1.4.1+
func GetLagoonAPIVersion(ctx context.Context, l APIVersion) (*schema.LagoonVersion, error) {
	// always start at v1.0.0, this is when rbac was introduced and CLI only supports v1.0.0+
	lagoonVersion := "v1.0.0"
	apiVersion := schema.LagoonVersion{}
	err := l.LagoonAPIVersion(ctx, &apiVersion)
	if err != nil {
		if !strings.Contains(err.Error(), `Cannot query field "lagoonVersion" on type "Query"`) {
			return &apiVersion, err
		}
	}
	if apiVersion.LagoonVersion != "" {
		_, err := version.NewSemver(apiVersion.LagoonVersion)
		if err == nil {
			return &apiVersion, nil
		}
	}
	lagoonSchema := schema.LagoonSchema{}
	err = l.LagoonSchema(ctx, &lagoonSchema)
	if err != nil {
		return &apiVersion, fmt.Errorf("couldn't perform request: %w", err)
	}
	// otherwise lets try to determine from schema changes
	lagoonVersion, err = determineLagoonVersion(lagoonVersion, lagoonSchema)
	if err != nil {
		return &apiVersion, fmt.Errorf("unable to determine version: %w", err)
	}
	apiVersion.LagoonVersion = lagoonVersion
	return &apiVersion, nil
}

// determine the version of the API based on the schema
func determineLagoonVersion(lagoonVersion string, lagoonSchema schema.LagoonSchema) (string, error) {
	var err error
	for _, schemaType := range lagoonSchema.Types {
		if schemaType.Name == "Mutation" {
			for _, field := range schemaType.Fields {
				if field.Name == "switchActiveStandby" {
					lagoonVersion, err = greaterThanOrEqualVersion(lagoonVersion, "v1.4.0")
					if err != nil {
						return lagoonVersion, err
					}
				}
			}
		}
		if schemaType.Name == "Query" {
			for _, field := range schemaType.Fields {
				if field.Name == "allGroups" {
					lagoonVersion, err = greaterThanOrEqualVersion(lagoonVersion, "v1.1.0")
					if err != nil {
						return lagoonVersion, err
					}
				}
				if field.Name == "me" {
					lagoonVersion, err = greaterThanOrEqualVersion(lagoonVersion, "v1.3.0")
					if err != nil {
						return lagoonVersion, err
					}
				}
				if field.Name == "projectsByMetadata" {
					lagoonVersion, err = greaterThanOrEqualVersion(lagoonVersion, "v1.6.0")
					if err != nil {
						return lagoonVersion, err
					}
				}
			}
		}
		if schemaType.Name == "NotificationMicrosoftTeams" {
			lagoonVersion, err = greaterThanOrEqualVersion(lagoonVersion, "v1.2.0")
			if err != nil {
				return lagoonVersion, err
			}
		}
	}
	return lagoonVersion, nil
}

// return the given or greater than version
func greaterThanOrEqualVersion(a string, b string) (string, error) {
	aVer, err := version.NewSemver(a)
	if err != nil {
		return a, err
	}
	bVer, err := version.NewSemver(b)
	if err != nil {
		return b, err
	}
	if aVer.GreaterThanOrEqual(bVer) {
		return a, nil
	}
	return b, nil
}
