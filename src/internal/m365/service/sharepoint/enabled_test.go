package sharepoint

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/canario/src/pkg/services/m365/api/graph/testdata"
)

type EnabledUnitSuite struct {
	tester.Suite
}

func TestEnabledUnitSuite(t *testing.T) {
	suite.Run(t, &EnabledUnitSuite{Suite: tester.NewUnitSuite(t)})
}

var _ getSiteRooter = mockGSR{}

type mockGSR struct {
	response models.Siteable
	err      error
}

func (m mockGSR) GetRoot(
	context.Context,
	api.CallConfig,
) (models.Siteable, error) {
	return m.response, m.err
}

func (suite *EnabledUnitSuite) TestIsServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getSiteRooter
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getSiteRooter {
				return mockGSR{models.NewSite(), nil}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "no sharepoint license",
			mock: func(ctx context.Context) getSiteRooter {
				odErr := graphTD.ODataErrWithMsg("code", string(graph.NoSPLicense))
				// needs graph.Stack, not clues.StackWC
				return mockGSR{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getSiteRooter {
				odErr := graphTD.ODataErrWithMsg("code", "message")
				return mockGSR{nil, clues.StackWC(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gsr := test.mock(ctx)

			ok, err := IsServiceEnabled(ctx, gsr, "resource_id")
			test.expect(t, ok, "has sites flag")
			test.expectErr(t, err)
		})
	}
}
