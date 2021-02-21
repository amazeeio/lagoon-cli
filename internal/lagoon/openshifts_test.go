//go:generate mockgen -source=openshifts.go -destination ../mock/mock_openshifts.go -package mock
package lagoon_test

import (
	"context"
	"testing"

	"github.com/amazeeio/lagoon-cli/internal/mock"
	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/golang/mock/gomock"
)

type openshiftCalls struct {
	Openshifts        []schema.Openshift
	DeletedOpenshifts []schema.DeleteOpenshift
	AddOpenshifts     []schema.AddOpenshiftInput
	UpdateOpenshifts  []schema.UpdateOpenshiftInput
	DeleteOpenshifts  []schema.DeleteOpenshiftInput
}

func TestGetAllOpenshifts(t *testing.T) {
	var testCases = map[string]struct {
		expect *openshiftCalls
	}{
		"simple": {expect: &openshiftCalls{
			Openshifts:        []schema.Openshift{},
			DeletedOpenshifts: []schema.DeleteOpenshift{},
			AddOpenshifts:     []schema.AddOpenshiftInput{},
			UpdateOpenshifts:  []schema.UpdateOpenshiftInput{},
			DeleteOpenshifts:  []schema.DeleteOpenshiftInput{},
		}},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			ctx := context.Background()
			// set up the mock openshifts
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()
			openshifts := mock.NewMockOpenshifts(ctrl)
			// use the provided openshiftCalls to set the expectations
			for i := range tc.expect.Openshifts {
				openshifts.EXPECT().AllOpenshifts(ctx, &tc.expect.Openshifts[i])
			}
			for i := range tc.expect.Openshifts {
				openshifts.EXPECT().AddOpenshift(ctx, &tc.expect.AddOpenshifts[i], &tc.expect.Openshifts[i])
			}
			for i := range tc.expect.Openshifts {
				openshifts.EXPECT().UpdateOpenshift(ctx, &tc.expect.UpdateOpenshifts[i], &tc.expect.Openshifts[i])
			}
			for i := range tc.expect.Openshifts {
				openshifts.EXPECT().DeleteOpenshift(ctx, &tc.expect.DeleteOpenshifts[i], &tc.expect.DeletedOpenshifts[i])
			}
		})
	}
}
