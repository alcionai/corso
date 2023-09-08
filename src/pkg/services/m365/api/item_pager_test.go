package api

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

// ---------------------------------------------------------------------------
// mock impls & stubs
// ---------------------------------------------------------------------------

// next and delta links

type nextLink struct {
	nextLink *string
}

func (l nextLink) GetOdataNextLink() *string {
	return l.nextLink
}

type deltaNextLink struct {
	nextLink
	deltaLink *string
}

func (l deltaNextLink) GetOdataDeltaLink() *string {
	return l.deltaLink
}

// mock page

type testPage struct {
	values []any
}

func (p testPage) GetOdataNextLink() *string {
	// no next, just one page
	return ptr.To("")
}

func (p testPage) GetOdataDeltaLink() *string {
	// delta is not tested here
	return ptr.To("")
}

func (p testPage) GetValue() []any {
	return p.values
}

// mock item pager

var _ Pager[any] = &testPager{}

type testPager struct {
	t       *testing.T
	pager   testPage
	pageErr error
}

func (p *testPager) GetPage(ctx context.Context) (NextLinkValuer[any], error) {
	return p.pager, p.pageErr
}

func (p *testPager) SetNextLink(nextLink string) {}

// mock id pager

var _ Pager[any] = &testIDsPager{}

type testIDsPager struct {
	t          *testing.T
	added      []string
	removed    []string
	errorCode  string
	needsReset bool
}

func (p *testIDsPager) GetPage(
	ctx context.Context,
) (NextLinkValuer[any], error) {
	if p.errorCode != "" {
		ierr := odataerrors.NewMainError()
		ierr.SetCode(&p.errorCode)

		err := odataerrors.NewODataError()
		err.SetErrorEscaped(ierr)

		return nil, err
	}

	return testPage{}, nil
}

func (p *testIDsPager) SetNextLink(string) {}

func (p *testIDsPager) Reset(context.Context) {
	if !p.needsReset {
		require.Fail(p.t, "reset should not be called")
	}

	p.needsReset = false
	p.errorCode = ""
}

var _ DeltaPager[any] = &testIDsDeltaPager{}

type testIDsDeltaPager struct {
	t          *testing.T
	added      []string
	removed    []string
	errorCode  string
	needsReset bool
}

func (p *testIDsDeltaPager) GetPage(
	ctx context.Context,
) (DeltaLinkValuer[any], error) {
	if p.errorCode != "" {
		ierr := odataerrors.NewMainError()
		ierr.SetCode(&p.errorCode)

		err := odataerrors.NewODataError()
		err.SetErrorEscaped(ierr)

		return nil, err
	}

	return testPage{}, nil
}

func (p *testIDsDeltaPager) SetNextLink(string) {}

func (p *testIDsDeltaPager) Reset(context.Context) {
	if !p.needsReset {
		require.Fail(p.t, "reset should not be called")
	}

	p.needsReset = false
	p.errorCode = ""
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

type PagerUnitSuite struct {
	tester.Suite
}

func TestPagerUnitSuite(t *testing.T) {
	suite.Run(t, &PagerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PagerUnitSuite) TestEnumerateItems() {
	tests := []struct {
		name      string
		getPager  func(*testing.T, context.Context) Pager[any]
		expect    []any
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			getPager: func(
				t *testing.T,
				ctx context.Context,
			) Pager[any] {
				return &testPager{
					t:     t,
					pager: testPage{[]any{"foo", "bar"}},
				}
			},
			expect:    []any{"foo", "bar"},
			expectErr: require.NoError,
		},
		{
			name: "next page err",
			getPager: func(
				t *testing.T,
				ctx context.Context,
			) Pager[any] {
				return &testPager{
					t:       t,
					pageErr: assert.AnError,
				}
			},
			expect:    nil,
			expectErr: require.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := enumerateItems(ctx, test.getPager(t, ctx))
			test.expectErr(t, err, clues.ToCore(err))

			require.EqualValues(t, test.expect, result)
		})
	}
}

func (suite *PagerUnitSuite) TestGetAddedAndRemovedItemIDs() {
	tests := []struct {
		name        string
		pagerGetter func(
			*testing.T,
			context.Context,
			graph.Servicer,
			string, string,
			bool,
		) (Pager[any], error)
		deltaPagerGetter func(
			*testing.T,
			context.Context,
			graph.Servicer,
			string, string, string,
			bool,
		) (DeltaPager[any], error)
		added               []string
		removed             []string
		deltaUpdate         DeltaUpdate
		delta               string
		canMakeDeltaQueries bool
	}{
		{
			name: "no prev delta",
			pagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (Pager[any], error) {
				// this should not be called
				return nil, assert.AnError
			},
			deltaPagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (DeltaPager[any], error) {
				return &testIDsDeltaPager{
					t:       t,
					added:   []string{"uno", "dos"},
					removed: []string{"tres", "quatro"},
				}, nil
			},
			added:               []string{"uno", "dos"},
			removed:             []string{"tres", "quatro"},
			deltaUpdate:         DeltaUpdate{Reset: true},
			canMakeDeltaQueries: true,
		},
		{
			name: "with prev delta",
			pagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (Pager[any], error) {
				// this should not be called
				return nil, assert.AnError
			},
			deltaPagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (DeltaPager[any], error) {
				return &testIDsDeltaPager{
					t:       t,
					added:   []string{"uno", "dos"},
					removed: []string{"tres", "quatro"},
				}, nil
			},
			added:               []string{"uno", "dos"},
			removed:             []string{"tres", "quatro"},
			delta:               "delta",
			deltaUpdate:         DeltaUpdate{Reset: false},
			canMakeDeltaQueries: true,
		},
		{
			name: "delta expired",
			pagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (Pager[any], error) {
				// this should not be called
				return nil, assert.AnError
			},
			deltaPagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (DeltaPager[any], error) {
				return &testIDsDeltaPager{
					t:          t,
					added:      []string{"uno", "dos"},
					removed:    []string{"tres", "quatro"},
					errorCode:  "SyncStateNotFound",
					needsReset: true,
				}, nil
			},
			added:               []string{"uno", "dos"},
			removed:             []string{"tres", "quatro"},
			delta:               "delta",
			deltaUpdate:         DeltaUpdate{Reset: true},
			canMakeDeltaQueries: true,
		},
		{
			name: "quota exceeded",
			pagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (Pager[any], error) {
				return &testIDsPager{
					t:       t,
					added:   []string{"uno", "dos"},
					removed: []string{"tres", "quatro"},
				}, nil
			},
			deltaPagerGetter: func(
				t *testing.T,
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (DeltaPager[any], error) {
				return &testIDsDeltaPager{errorCode: "ErrorQuotaExceeded"}, nil
			},
			added:               []string{"uno", "dos"},
			removed:             []string{"tres", "quatro"},
			deltaUpdate:         DeltaUpdate{Reset: true},
			canMakeDeltaQueries: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			pager, _ := test.pagerGetter(t, ctx, graph.Service{}, "user", "directory", false)
			deltaPager, _ := test.deltaPagerGetter(t, ctx, graph.Service{}, "user", "directory", test.delta, false)

			added, removed, deltaUpdate, err := getAddedAndRemovedItemIDs[any](
				ctx,
				pager,
				deltaPager,
				test.delta,
				test.canMakeDeltaQueries)

			require.NoError(t, err, "getting added and removed item IDs")
			require.EqualValues(t, test.added, added, "added item IDs")
			require.EqualValues(t, test.removed, removed, "removed item IDs")
			require.Equal(t, test.deltaUpdate, deltaUpdate, "delta update")
		})
	}
}

type testInput struct {
	name         string
	inputLink    *string
	expectedLink string
}

// Needs to be var not const so we can take the address of it.
var (
	emptyLink = ""
	link      = "foo"
	link2     = "bar"

	nextLinkInputs = []testInput{
		{
			name:         "empty",
			inputLink:    &emptyLink,
			expectedLink: "",
		},
		{
			name:         "nil",
			inputLink:    nil,
			expectedLink: "",
		},
		{
			name:         "non_empty",
			inputLink:    &link,
			expectedLink: link,
		},
	}
)

func (suite *PagerUnitSuite) TestNextAndDeltaLink() {
	deltaTable := []testInput{
		{
			name:         "empty",
			inputLink:    &emptyLink,
			expectedLink: "",
		},
		{
			name:         "nil",
			inputLink:    nil,
			expectedLink: "",
		},
		{
			name: "non_empty",
			// Use a different link so we can see if the results get swapped or something.
			inputLink:    &link2,
			expectedLink: link2,
		},
	}

	for _, next := range nextLinkInputs {
		for _, delta := range deltaTable {
			name := strings.Join([]string{next.name, "next", delta.name, "delta"}, "_")

			suite.Run(name, func() {
				t := suite.T()

				l := deltaNextLink{
					nextLink:  nextLink{nextLink: next.inputLink},
					deltaLink: delta.inputLink,
				}
				gotNext, gotDelta := NextAndDeltaLink(l)

				assert.Equal(t, next.expectedLink, gotNext)
				assert.Equal(t, delta.expectedLink, gotDelta)
			})
		}
	}
}

// TestIsLinkValid check to verify is nextLink guard check for logging
// Related to: https://github.com/alcionai/corso/issues/2520
//
//nolint:lll
func (suite *PagerUnitSuite) TestIsLinkValid() {
	invalidString := `https://graph.microsoft.com/v1.0/users//mailFolders//messages/microsoft.graph.delta()?$select=id%2CisRead`
	tests := []struct {
		name        string
		inputString string
		isValid     assert.BoolAssertionFunc
	}{
		{
			name:        "Empty",
			inputString: emptyLink,
			isValid:     assert.True,
		},
		{
			name:        "Invalid",
			inputString: invalidString,
			isValid:     assert.False,
		},
		{
			name:        "Valid",
			inputString: `https://graph.microsoft.com/v1.0/users/aPerson/mailFolders/AMessage/messages/microsoft.graph.delta()?$select=id%2CisRead`,
			isValid:     assert.True,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			got := IsNextLinkValid(test.inputString)
			test.isValid(suite.T(), got)
		})
	}
}
