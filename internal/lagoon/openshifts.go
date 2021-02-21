// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Openshifts interface contains methods for interacting with openshifts in lagoon.
type Openshifts interface {
	AllOpenshifts(ctx context.Context, result *[]schema.Openshift) error
	AddOpenshift(ctx context.Context, openshift *schema.AddOpenshiftInput, result *schema.Openshift) error
	UpdateOpenshift(ctx context.Context, openshift *schema.UpdateOpenshiftInput, result *schema.Openshift) error
	DeleteOpenshift(ctx context.Context, openshift *schema.DeleteOpenshiftInput, result *schema.DeleteOpenshift) error
}

// GetAllOpenshifts lists all openshifts in a lagoon.
func GetAllOpenshifts(ctx context.Context, o Openshifts) (*[]schema.Openshift, error) {
	result := []schema.Openshift{}
	return &result, o.AllOpenshifts(ctx, &result)
}

// AddOpenshift add an openshift to a lagoon.
func AddOpenshift(ctx context.Context, openshift *schema.AddOpenshiftInput, o Openshifts) (*schema.Openshift, error) {
	result := schema.Openshift{}
	return &result, o.AddOpenshift(ctx, openshift, &result)
}

// UpdateOpenshift add an openshift to a lagoon.
func UpdateOpenshift(ctx context.Context, openshift *schema.UpdateOpenshiftInput, o Openshifts) (*schema.Openshift, error) {
	result := schema.Openshift{}
	return &result, o.UpdateOpenshift(ctx, openshift, &result)
}

// DeleteOpenshift delete an openshift from a lagoon.
func DeleteOpenshift(ctx context.Context, openshift *schema.DeleteOpenshiftInput, o Openshifts) (*schema.DeleteOpenshift, error) {
	result := schema.DeleteOpenshift{}
	return &result, o.DeleteOpenshift(ctx, openshift, &result)
}
