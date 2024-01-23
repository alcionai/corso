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
)

type EnabledUnitSuite struct {
	tester.Suite
}

func TestEnabledUnitSuite(t *testing.T) {
	suite.Run(t, &EnabledUnitSuite{Suite: tester.NewUnitSuite(t)})
}

var _ api.GetByIDer[models.Userable] = mockGU{}

type mockGU struct {
	user models.Userable
	err  error
}

func (m mockGU) GetByID(
	ctx context.Context,
	identifier string,
	_ api.CallConfig,
) (models.Userable, error) {
	return m.user, m.err
}

func (suite *EnabledUnitSuite) TestIsServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) api.GetByIDer[models.Userable]
		expect    assert.BoolAssertionFunc
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) api.GetByIDer[models.Userable] {
				return mockGU{}
			},
			expect:    assert.True,
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gu := test.mock(ctx)

			ok, err := IsServiceEnabled(ctx, gu, "resource_id")
			test.expect(t, ok, "has mailbox flag")
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
