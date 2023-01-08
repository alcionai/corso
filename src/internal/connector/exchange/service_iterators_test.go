package exchange

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

var _ addedAndRemovedItemIDsGetter = &mockGetter{}

type (
	mockGetter        map[string]mockGetterResults
	mockGetterResults struct {
		added    []string
		removed  []string
		newDelta api.DeltaUpdate
		err      error
	}
)

func (mg mockGetter) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, cID, prevDelta string,
) (
	[]string,
	[]string,
	api.DeltaUpdate,
	error,
) {
	results, ok := mg[cID]
	if !ok {
		return nil, nil, api.DeltaUpdate{}, errors.New("mock not found for " + cID)
	}

	return results.added, results.removed, results.newDelta, results.err
}

var _ graph.ContainerResolver = &mockResolver{}

type (
	mockResolver struct {
		items []graph.CachedContainer
	}
)

func newMockResolver(items ...mockContainer) mockResolver {
	is := make([]graph.CachedContainer, 0, len(items))

	for _, i := range items {
		is = append(is, i)
	}

	return mockResolver{items: is}
}

func (m mockResolver) Items() []graph.CachedContainer {
	return m.items
}

func (m mockResolver) AddToCache(context.Context, graph.Container) error       { return nil }
func (m mockResolver) IDToPath(context.Context, string) (*path.Builder, error) { return nil, nil }
func (m mockResolver) PathInCache(string) (string, bool)                       { return "", false }
func (m mockResolver) Populate(context.Context, string, ...string) error       { return nil }

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type ServiceIteratorsSuite struct {
	suite.Suite
	creds account.M365Config
}

func TestServiceIteratorsSuite(t *testing.T) {
	suite.Run(t, new(ServiceIteratorsSuite))
}

func (suite *ServiceIteratorsSuite) SetupSuite() {
	a := tester.NewM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err)
	suite.creds = m365
}

func (suite *ServiceIteratorsSuite) TestFilterContainersAndFillCollections() {
	var (
		userID = "user_id"
		qp     = graph.QueryParams{
			Category:      path.EmailCategory, // doesn't matter which one we use.
			ResourceOwner: userID,
			Credentials:   suite.creds,
		}
		statusUpdater = func(*support.ConnectorOperationStatus) {}
		allScope      = selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0]
		dps           = DeltaPaths{} // incrementals are tested separately
		commonResult  = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: api.DeltaUpdate{URL: "delta_url"},
		}
		errorResult = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: api.DeltaUpdate{URL: "delta_url"},
			err:      assert.AnError,
		}
		deletedInFlightResult = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: api.DeltaUpdate{URL: "delta_url"},
			err:      graph.ErrDeletedInFlight{Err: *common.EncapsulateError(assert.AnError)},
		}
		container1 = mockContainer{
			id:          strPtr("1"),
			displayName: strPtr("display_name_1"),
			p:           path.Builder{}.Append("display_name_1"),
		}
		container2 = mockContainer{
			id:          strPtr("2"),
			displayName: strPtr("display_name_2"),
			p:           path.Builder{}.Append("display_name_2"),
		}
	)

	table := []struct {
		name                  string
		getter                mockGetter
		resolver              graph.ContainerResolver
		scope                 selectors.ExchangeScope
		failFast              bool
		expectErr             assert.ErrorAssertionFunc
		expectNewColls        int
		expectMetadataColls   int
		expectDoNotMergeColls int
	}{
		{
			name: "happy path, one container",
			getter: map[string]mockGetterResults{
				"1": commonResult,
			},
			resolver:            newMockResolver(container1),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, many containers",
			getter: map[string]mockGetterResults{
				"1": commonResult,
				"2": commonResult,
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      2,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, many containers, same display name",
			getter: map[string]mockGetterResults{
				"1": commonResult,
				"2": commonResult,
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      2,
			expectMetadataColls: 1,
		},
		{
			name: "no containers pass scope",
			getter: map[string]mockGetterResults{
				"1": commonResult,
				"2": commonResult,
			},
			resolver:            newMockResolver(container1, container2),
			scope:               selectors.NewExchangeBackup(nil).MailFolders(selectors.None())[0],
			expectErr:           assert.NoError,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "err: deleted in flight",
			getter: map[string]mockGetterResults{
				"1": deletedInFlightResult,
			},
			resolver:              newMockResolver(container1),
			scope:                 allScope,
			expectErr:             assert.NoError,
			expectNewColls:        1,
			expectMetadataColls:   1,
			expectDoNotMergeColls: 1,
		},
		{
			name: "err: other error",
			getter: map[string]mockGetterResults{
				"1": errorResult,
			},
			resolver:            newMockResolver(container1),
			scope:               allScope,
			expectErr:           assert.Error,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "half collections error: deleted in flight",
			getter: map[string]mockGetterResults{
				"1": deletedInFlightResult,
				"2": commonResult,
			},
			resolver:              newMockResolver(container1, container2),
			scope:                 allScope,
			expectErr:             assert.NoError,
			expectNewColls:        2,
			expectMetadataColls:   1,
			expectDoNotMergeColls: 1,
		},
		{
			name: "half collections error: other error",
			getter: map[string]mockGetterResults{
				"1": errorResult,
				"2": commonResult,
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			expectErr:           assert.Error,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "half collections error: deleted in flight, fail fast",
			getter: map[string]mockGetterResults{
				"1": deletedInFlightResult,
				"2": commonResult,
			},
			resolver:              newMockResolver(container1, container2),
			scope:                 allScope,
			failFast:              true,
			expectErr:             assert.NoError,
			expectNewColls:        2,
			expectMetadataColls:   1,
			expectDoNotMergeColls: 1,
		},
		{
			name: "half collections error: other error, fail fast",
			getter: map[string]mockGetterResults{
				"1": errorResult,
				"2": commonResult,
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			failFast:            true,
			expectErr:           assert.Error,
			expectNewColls:      0,
			expectMetadataColls: 0,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			collections := map[string]data.Collection{}

			err := filterContainersAndFillCollections(
				ctx,
				qp,
				test.getter,
				collections,
				statusUpdater,
				test.resolver,
				test.scope,
				dps,
				control.Options{FailFast: test.failFast},
			)
			test.expectErr(t, err)

			// collection assertions

			deleteds, news, metadatas, doNotMerges := 0, 0, 0, 0
			for _, c := range collections {
				if c.FullPath().Service() == path.ExchangeMetadataService {
					metadatas++
					continue
				}

				if c.State() == data.DeletedState {
					deleteds++
				}

				if c.State() == data.NewState {
					news++
				}

				if c.DoNotMergeItems() {
					doNotMerges++
				}
			}

			assert.Zero(t, deleteds, "deleted collections")
			assert.Equal(t, test.expectNewColls, news, "new collections")
			assert.Equal(t, test.expectMetadataColls, metadatas, "metadata collections")
			assert.Equal(t, test.expectDoNotMergeColls, doNotMerges, "doNotMerge collections")

			// items in collections assertions
			for k, expect := range test.getter {
				coll := collections[k]

				if coll == nil {
					continue
				}

				exColl, ok := coll.(*Collection)
				require.True(t, ok, "collection is an *exchange.Collection")

				assert.ElementsMatch(t, expect.added, exColl.added, "added items")
				assert.ElementsMatch(t, expect.removed, exColl.removed, "removed items")
			}
		})
	}
}

func (suite *ServiceIteratorsSuite) TestFilterContainersAndFillCollections_incrementals() {
	var (
		userID   = "user_id"
		tenantID = tester.M365TenantID(suite.T())
		cat      = path.EmailCategory // doesn't matter which one we use,
		qp       = graph.QueryParams{
			Category:      cat,
			ResourceOwner: userID,
			Credentials:   suite.creds,
		}
		statusUpdater = func(*support.ConnectorOperationStatus) {}
		allScope      = selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0]
		commonResults = mockGetterResults{
			added:    []string{"added"},
			newDelta: api.DeltaUpdate{URL: "new_delta_url"},
		}
		expiredResults = mockGetterResults{
			added: []string{"added"},
			newDelta: api.DeltaUpdate{
				URL:   "new_delta_url",
				Reset: true,
			},
		}
	)

	prevPath := func(t *testing.T, at ...string) path.Path {
		p, err := path.Builder{}.
			Append(at...).
			ToDataLayerExchangePathForCategory(tenantID, userID, cat, false)
		require.NoError(t, err)

		return p
	}

	type endState struct {
		state      data.CollectionState
		doNotMerge bool
	}

	table := []struct {
		name     string
		getter   mockGetter
		resolver graph.ContainerResolver
		dps      DeltaPaths
		expect   map[string]endState
	}{
		{
			name: "new container",
			getter: map[string]mockGetterResults{
				"1": commonResults,
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("new"),
				p:           path.Builder{}.Append("1", "new"),
			}),
			dps: DeltaPaths{},
			expect: map[string]endState{
				"1": {data.NewState, false},
			},
		},
		{
			name: "not moved container",
			getter: map[string]mockGetterResults{
				"1": commonResults,
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("not_moved"),
				p:           path.Builder{}.Append("1", "not_moved"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "1", "not_moved").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NotMovedState, false},
			},
		},
		{
			name: "moved container",
			getter: map[string]mockGetterResults{
				"1": commonResults,
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("moved"),
				p:           path.Builder{}.Append("1", "moved"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "1", "prev").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.MovedState, false},
			},
		},
		{
			name:     "deleted container",
			getter:   map[string]mockGetterResults{},
			resolver: newMockResolver(),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "1", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
			},
		},
		{
			name: "one deleted, one new",
			getter: map[string]mockGetterResults{
				"2": commonResults,
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("2"),
				displayName: strPtr("new"),
				p:           path.Builder{}.Append("2", "new"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "1", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "one deleted, one new, same path",
			getter: map[string]mockGetterResults{
				"2": commonResults,
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("2"),
				displayName: strPtr("same"),
				p:           path.Builder{}.Append("2", "same"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "1", "same").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "one moved, one new, same path",
			getter: map[string]mockGetterResults{
				"1": commonResults,
				"2": commonResults,
			},
			resolver: newMockResolver(
				mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("1", "moved"),
				},
				mockContainer{
					id:          strPtr("2"),
					displayName: strPtr("prev"),
					p:           path.Builder{}.Append("2", "prev"),
				},
			),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "1", "prev").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.MovedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "bad previous path strings",
			getter: map[string]mockGetterResults{
				"1": commonResults,
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("not_moved"),
				p:           path.Builder{}.Append("1", "not_moved"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  "1/fnords/mc/smarfs",
				},
				"2": DeltaPath{
					delta: "old_delta_url",
					path:  "2/fnords/mc/smarfs",
				},
			},
			expect: map[string]endState{
				"1": {data.NewState, false},
			},
		},
		{
			name: "delta expiration",
			getter: map[string]mockGetterResults{
				"1": expiredResults,
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("same"),
				p:           path.Builder{}.Append("1", "same"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "1", "same").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NotMovedState, true},
			},
		},
		{
			name: "a little bit of everything",
			getter: map[string]mockGetterResults{
				"1": commonResults,  // new
				"2": commonResults,  // notMoved
				"3": commonResults,  // moved
				"4": expiredResults, // moved
				// "5" gets deleted
			},
			resolver: newMockResolver(
				mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("new"),
					p:           path.Builder{}.Append("1", "new"),
				},
				mockContainer{
					id:          strPtr("2"),
					displayName: strPtr("not_moved"),
					p:           path.Builder{}.Append("2", "not_moved"),
				},
				mockContainer{
					id:          strPtr("3"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("3", "moved"),
				},
				mockContainer{
					id:          strPtr("4"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("4", "moved"),
				},
			),
			dps: DeltaPaths{
				"2": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "2", "not_moved").String(),
				},
				"3": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "3", "prev").String(),
				},
				"4": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "4", "prev").String(),
				},
				"5": DeltaPath{
					delta: "old_delta_url",
					path:  prevPath(suite.T(), "5", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NewState, false},
				"2": {data.NotMovedState, false},
				"3": {data.MovedState, false},
				"4": {data.MovedState, true},
				"5": {data.DeletedState, false},
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			collections := map[string]data.Collection{}

			err := filterContainersAndFillCollections(
				ctx,
				qp,
				test.getter,
				collections,
				statusUpdater,
				test.resolver,
				allScope,
				test.dps,
				control.Options{},
			)
			assert.NoError(t, err)

			metadatas := 0
			for _, c := range collections {
				p := c.FullPath()
				if p == nil {
					p = c.PreviousPath()
				}

				require.NotNil(t, p)

				if p.Service() == path.ExchangeMetadataService {
					metadatas++
					continue
				}

				p0 := p.Folders()[0]

				expect, ok := test.expect[p0]
				assert.True(t, ok, "collection is expected in result")

				assert.Equalf(t, expect.state, c.State(), "collection %s state", p0)
				assert.Equalf(t, expect.doNotMerge, c.DoNotMergeItems(), "collection %s DoNotMergeItems", p0)
			}

			assert.Equal(t, 1, metadatas, "metadata collections")
		})
	}
}
