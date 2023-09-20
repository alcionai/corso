package groups

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type EnabledUnitSuite struct {
	tester.Suite
}

func TestEnabledUnitSuite(t *testing.T) {
	suite.Run(t, &EnabledUnitSuite{Suite: tester.NewUnitSuite(t)})
}

var _ getByIDer = mockGBI{}

type mockGBI struct {
	group models.Groupable
	err   error
}

func (m mockGBI) GetByID(ctx context.Context, identifier string) (models.Groupable, error) {
	return m.group, m.err
}

// TODO(pandeyabs): Duplicate of graph/errors_test.go. Remove
// this and identical funcs in od/sp and use the one in graph/errors_test.go
// instead.
func odErrMsg(code, message string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	merr.SetMessage(&message)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func (suite *EnabledUnitSuite) TestIsServiceEnabled() {
	var (
		unified    = models.NewGroup()
		nonUnified = models.NewGroup()
	)

	unified.SetGroupTypes([]string{"unified"})

	table := []struct {
		name      string
		mock      func(context.Context) getByIDer
		expect    assert.BoolAssertionFunc
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getByIDer {
				return mockGBI{
					group: unified,
				}
			},
			expect:    assert.True,
			expectErr: assert.NoError,
		},
		{
			name: "non-unified group",
			mock: func(ctx context.Context) getByIDer {
				return mockGBI{
					group: nonUnified,
				}
			},
			expect:    assert.False,
			expectErr: assert.NoError,
		},
		{
			name: "group not found",
			mock: func(ctx context.Context) getByIDer {
				return mockGBI{
					err: graph.Stack(ctx, odErrMsg(string(graph.RequestResourceNotFound), "message")),
				}
			},
			expect:    assert.False,
			expectErr: assert.Error,
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getByIDer {
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
