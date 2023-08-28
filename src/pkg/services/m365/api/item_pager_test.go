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

// mock values

type testPagerValue struct {
	id      string
	removed bool
}

func (v testPagerValue) GetId() *string { return &v.id } //revive:disable-line:var-naming
func (v testPagerValue) GetAdditionalData() map[string]any {
	if v.removed {
		return map[string]any{graph.AddtlDataRemoved: true}
	}

	return map[string]any{}
}

// mock page

type testPage[T any] struct {
	values []T
}

func (p testPage[T]) GetOdataNextLink() *string {
	// no next, just one page
	return ptr.To("")
}

func (p testPage[T]) GetOdataDeltaLink() *string {
	// delta is not tested here
	return ptr.To("")
}

func (p testPage[T]) GetValue() []T {
	return p.values
}

// mock item pager

var _ Pager[any] = &testPager{}

type testPager struct {
	t       *testing.T
	pager   testPage[any]
	pageErr error
}

func (p *testPager) GetPage(ctx context.Context) (LinkValuer[any], error) {
	return p.pager, p.pageErr
}

func (p *testPager) SetNext(nextLink string) {}

// mock id pager

var _ DeltaPager[getIDAndAddtler] = &testIDsPager{}

type testIDsPager struct {
	t          *testing.T
	added      []string
	removed    []string
	errorCode  string
	needsReset bool
}

func (p *testIDsPager) GetPage(ctx context.Context) (DeltaLinkValuer[getIDAndAddtler], error) {
	if p.errorCode != "" {
		ierr := odataerrors.NewMainError()
		ierr.SetCode(&p.errorCode)

		err := odataerrors.NewODataError()
		err.SetErrorEscaped(ierr)

		return nil, err
	}

	return testPage[getIDAndAddtler]{}, nil
}
func (p *testIDsPager) SetNext(string) {}
func (p *testIDsPager) Reset() {
	if !p.needsReset {
		require.Fail(p.t, "reset should not be called")
	}

	p.needsReset = false
	p.errorCode = ""
}

func (p *testIDsPager) ValuesIn(pl LinkValuer[getIDAndAddtler]) ([]getIDAndAddtler, error) {
	items := []getIDAndAddtler{}

	for _, id := range p.added {
		items = append(items, testPagerValue{id: id})
	}

	for _, id := range p.removed {
		items = append(items, testPagerValue{id: id, removed: true})
	}

	return items, nil
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

type ItemPagerUnitSuite struct {
	tester.Suite
}

func TestItemPagerUnitSuite(t *testing.T) {
	suite.Run(t, &ItemPagerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemPagerUnitSuite) TestEnumerateItems() {
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
					pager: testPage[any]{[]any{"foo", "bar"}},
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

func (suite *ItemPagerUnitSuite) TestGetAddedAndRemovedItemIDs() {
	tests := []struct {
		name        string
		pagerGetter func(
			*testing.T,
			context.Context,
			graph.Servicer,
			string, string,
			bool,
		) (DeltaPager[getIDAndAddtler], error)
		deltaPagerGetter func(
			*testing.T,
			context.Context,
			graph.Servicer,
			string, string, string,
			bool,
		) (DeltaPager[getIDAndAddtler], error)
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
			) (DeltaPager[getIDAndAddtler], error) {
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
			) (DeltaPager[getIDAndAddtler], error) {
				return &testIDsPager{
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
			) (DeltaPager[getIDAndAddtler], error) {
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
			) (DeltaPager[getIDAndAddtler], error) {
				return &testIDsPager{
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
			) (DeltaPager[getIDAndAddtler], error) {
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
			) (DeltaPager[getIDAndAddtler], error) {
				return &testIDsPager{
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
			) (DeltaPager[getIDAndAddtler], error) {
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
			) (DeltaPager[getIDAndAddtler], error) {
				return &testIDsPager{errorCode: "ErrorQuotaExceeded"}, nil
			},
			added:               []string{"uno", "dos"},
			removed:             []string{"tres", "quatro"},
			deltaUpdate:         DeltaUpdate{Reset: true},
			canMakeDeltaQueries: false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			pager, _ := tt.pagerGetter(t, ctx, graph.Service{}, "user", "directory", false)
			deltaPager, _ := tt.deltaPagerGetter(t, ctx, graph.Service{}, "user", "directory", tt.delta, false)

			added, removed, deltaUpdate, err := getAddedAndRemovedItemIDs(
				ctx,
				graph.Service{},
				pager,
				deltaPager,
				tt.delta,
				tt.canMakeDeltaQueries)

			require.NoError(t, err, "getting added and removed item IDs")
			require.EqualValues(t, tt.added, added, "added item IDs")
			require.EqualValues(t, tt.removed, removed, "removed item IDs")
			require.Equal(t, tt.deltaUpdate, deltaUpdate, "delta update")
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

func (suite *ItemPagerUnitSuite) TestNextAndDeltaLink() {
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
func (suite *ItemPagerUnitSuite) TestIsLinkValid() {
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
