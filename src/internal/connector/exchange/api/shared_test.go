package api

import (
	"context"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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
	next := "" // no next, just one page
	return &next
}

var _ itemPager = &testPager{}

type testPager struct {
	errorCode string
	added     []string
	removed   []string
}

func (p *testPager) getPage(ctx context.Context) (api.PageLinker, error) {
	if p.errorCode != "" {
		ierr := odataerrors.NewMainError()
		ierr.SetCode(&p.errorCode)

		err := odataerrors.NewODataError()
		err.SetError(ierr)

		return nil, err
	}

	return testPage{}, nil
}
func (p *testPager) setNext(nextLink string) {}
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
		name             string
		pagerGetter      func(context.Context, graph.Servicer, string, string, bool) (itemPager, error)
		deltaPagerGetter func(context.Context, graph.Servicer, string, string, string, bool) (itemPager, error)
		added            []string
		removed          []string
		deltaUpdate      DeltaUpdate
		delta            string
	}{
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
				if len(delta) == 0 {
					return &testPager{
						added:   []string{"uno", "dos"},
						removed: []string{"tres", "quatro"},
					}, nil
				}

				return nil, assert.AnError
			},
			added:   []string{"uno", "dos"},
			removed: []string{"tres", "quatro"},
		},
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
				if len(delta) != 0 {
					return &testPager{
						added:   []string{"uno", "dos"},
						removed: []string{"tres", "quatro"},
					}, nil
				}

				return nil, assert.AnError
			},
			added:       []string{"uno", "dos"},
			removed:     []string{"tres", "quatro"},
			delta:       "delta",
			deltaUpdate: DeltaUpdate{Reset: true},
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
				if len(delta) == 0 {
					return &testPager{
						added:   []string{"uno", "dos"},
						removed: []string{"tres", "quatro"},
					}, nil
				}

				return &testPager{errorCode: "SyncStateNotFound"}, nil
			},
			added:       []string{"uno", "dos"},
			removed:     []string{"tres", "quatro"},
			delta:       "delta",
			deltaUpdate: DeltaUpdate{Reset: true},
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
			added:       []string{"uno", "dos"},
			removed:     []string{"tres", "quatro"},
			deltaUpdate: DeltaUpdate{Reset: true},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			added, removed, deltaUpdate, err := getAddedAndRemovedItemIDs(
				ctx,
				graph.Service{},
				"user",
				"directory",
				tt.delta,
				tt.pagerGetter,
				tt.deltaPagerGetter,
				false,
			)

			require.NoError(suite.T(), err, "getting added and removed item IDs")
			require.EqualValues(suite.T(), tt.added, added, "added item IDs")
			require.EqualValues(suite.T(), tt.removed, removed, "removed item IDs")
			require.Equal(suite.T(), tt.deltaUpdate, deltaUpdate, "delta update")
		})
	}
}
