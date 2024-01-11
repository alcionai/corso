package groups

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
)

type EnabledUnitSuite struct {
	tester.Suite
}

func TestEnabledUnitSuite(t *testing.T) {
	suite.Run(t, &EnabledUnitSuite{Suite: tester.NewUnitSuite(t)})
}

var _ api.GetByIDer[models.Groupable] = mockGBI{}

type mockGBI struct {
	group models.Groupable
	err   error
}

func (m mockGBI) GetByID(
	ctx context.Context,
	identifier string,
	_ api.CallConfig,
) (models.Groupable, error) {
	return m.group, m.err
}

func (suite *EnabledUnitSuite) TestIsServiceEnabled() {
	var (
		unified    = models.NewGroup()
		nonUnified = models.NewGroup()
	)

	unified.SetGroupTypes([]string{"unified"})

	table := []struct {
		name      string
		mock      func(context.Context) api.GetByIDer[models.Groupable]
		expect    assert.BoolAssertionFunc
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) api.GetByIDer[models.Groupable] {
				return mockGBI{
					group: unified,
				}
			},
			expect:    assert.True,
			expectErr: assert.NoError,
		},
		{
			name: "non-unified group",
			mock: func(ctx context.Context) api.GetByIDer[models.Groupable] {
				return mockGBI{
					group: nonUnified,
				}
			},
			expect:    assert.False,
			expectErr: assert.NoError,
		},
		{
			name: "group not found",
			mock: func(ctx context.Context) api.GetByIDer[models.Groupable] {
				return mockGBI{
					err: clues.StackWC(ctx, graphTD.ODataErrWithMsg(string(graph.RequestResourceNotFound), "message")),
				}
			},
			expect:    assert.False,
			expectErr: assert.Error,
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) api.GetByIDer[models.Groupable] {
				return mockGBI{
					err: assert.AnError,
				}
			},
			expect:    assert.False,
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gmi := test.mock(ctx)

			ok, err := IsServiceEnabled(ctx, gmi, "resource_id")
			test.expect(t, ok, "has mailbox flag")
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
