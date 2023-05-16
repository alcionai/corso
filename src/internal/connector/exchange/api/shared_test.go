package api

import (
	"context"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/tester"
)

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

type testPage struct{}

func (p testPage) GetOdataNextLink() *string {
	// no next, just one page
	return ptr.To("")
}

func (p testPage) GetOdataDeltaLink() *string {
	// delta is not tested here
	return ptr.To("")
}

var _ itemPager = &testPager{}

type testPager struct {
	t          *testing.T
	added      []string
	removed    []string
	errorCode  string
	needsReset bool
}

func (p *testPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	if p.errorCode != "" {
		ierr := odataerrors.NewMainError()
		ierr.SetCode(&p.errorCode)

		err := odataerrors.NewODataError()
		err.SetError(ierr)

		return nil, err
	}

	return testPage{}, nil
}
func (p *testPager) setNext(string) {}
func (p *testPager) reset(context.Context) {
	if !p.needsReset {
		require.Fail(p.t, "reset should not be called")
	}

	p.needsReset = false
	p.errorCode = ""
}

func (p *testPager) valuesIn(pl api.PageLinker) ([]getIDAndAddtler, error) {
	items := []getIDAndAddtler{}

	for _, id := range p.added {
		items = append(items, testPagerValue{id: id})
	}

	for _, id := range p.removed {
		items = append(items, testPagerValue{id: id, removed: true})
	}

	return items, nil
}

type SharedAPIUnitSuite struct {
	tester.Suite
}

func TestSharedAPIUnitSuite(t *testing.T) {
	suite.Run(t, &SharedAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SharedAPIUnitSuite) TestGetAddedAndRemovedItemIDs() {
	tests := []struct {
		name                string
		pagerGetter         func(context.Context, graph.Servicer, string, string, bool) (itemPager, error)
		deltaPagerGetter    func(context.Context, graph.Servicer, string, string, string, bool) (itemPager, error)
		added               []string
		removed             []string
		deltaUpdate         DeltaUpdate
		delta               string
		canMakeDeltaQueries bool
	}{
		{
			name: "no prev delta",
			pagerGetter: func(
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (itemPager, error) {
				// this should not be called
				return nil, assert.AnError
			},
			deltaPagerGetter: func(
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (itemPager, error) {
				return &testPager{
					t:       suite.T(),
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
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (itemPager, error) {
				// this should not be called
				return nil, assert.AnError
			},
			deltaPagerGetter: func(
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (itemPager, error) {
				return &testPager{
					t:       suite.T(),
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
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (itemPager, error) {
				// this should not be called
				return nil, assert.AnError
			},
			deltaPagerGetter: func(
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (itemPager, error) {
				return &testPager{
					t:          suite.T(),
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
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				immutableIDs bool,
			) (itemPager, error) {
				return &testPager{
					t:       suite.T(),
					added:   []string{"uno", "dos"},
					removed: []string{"tres", "quatro"},
				}, nil
			},
			deltaPagerGetter: func(
				ctx context.Context,
				gs graph.Servicer,
				user string,
				directory string,
				delta string,
				immutableIDs bool,
			) (itemPager, error) {
				return &testPager{errorCode: "ErrorQuotaExceeded"}, nil
			},
			added:               []string{"uno", "dos"},
			removed:             []string{"tres", "quatro"},
			deltaUpdate:         DeltaUpdate{Reset: true},
			canMakeDeltaQueries: false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			pager, _ := tt.pagerGetter(ctx, graph.Service{}, "user", "directory", false)
			deltaPager, _ := tt.deltaPagerGetter(ctx, graph.Service{}, "user", "directory", tt.delta, false)

			added, removed, deltaUpdate, err := getAddedAndRemovedItemIDs(
				ctx,
				graph.Service{},
				pager,
				deltaPager,
				tt.delta,
				tt.canMakeDeltaQueries,
			)

			require.NoError(suite.T(), err, "getting added and removed item IDs")
			require.EqualValues(suite.T(), tt.added, added, "added item IDs")
			require.EqualValues(suite.T(), tt.removed, removed, "removed item IDs")
			require.Equal(suite.T(), tt.deltaUpdate, deltaUpdate, "delta update")
		})
	}
}
