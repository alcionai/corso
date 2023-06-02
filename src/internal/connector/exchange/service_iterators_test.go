package exchange

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

var _ backupHandler = &mockBackupHandler{}

type mockBackupHandler struct {
	mg       mockGetter
	category path.CategoryType
	ac       api.Client
	userID   string
}

func (bh mockBackupHandler) itemEnumerator() addedAndRemovedItemGetter { return bh.mg }
func (bh mockBackupHandler) itemHandler() itemGetterSerializer         { return nil }

func (bh mockBackupHandler) NewContainerCache(
	userID string,
) (string, graph.ContainerResolver) {
	return BackupHandlers(bh.ac)[bh.category].NewContainerCache(bh.userID)
}

var _ addedAndRemovedItemGetter = &mockGetter{}

type (
	mockGetter struct {
		noReturnDelta bool
		results       map[string]mockGetterResults
	}
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
	_ bool,
	_ bool,
) (
	[]string,
	[]string,
	api.DeltaUpdate,
	error,
) {
	results, ok := mg.results[cID]
	if !ok {
		return nil, nil, api.DeltaUpdate{}, clues.New("mock not found for " + cID)
	}

	delta := results.newDelta
	if mg.noReturnDelta {
		delta.URL = ""
	}

	return results.added, results.removed, delta, results.err
}

var _ graph.ContainerResolver = &mockResolver{}

type (
	mockResolver struct {
		items []graph.CachedContainer
		added map[string]string
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

func (m mockResolver) AddToCache(ctx context.Context, gc graph.Container) error {
	if len(m.added) == 0 {
		m.added = map[string]string{}
	}

	m.added[ptr.Val(gc.GetDisplayName())] = ptr.Val(gc.GetId())

	return nil
}
func (m mockResolver) DestinationNameToID(dest string) string { return m.added[dest] }
func (m mockResolver) IDToPath(context.Context, string) (*path.Builder, *path.Builder, error) {
	return nil, nil, nil
}
func (m mockResolver) PathInCache(string) (string, bool)                             { return "", false }
func (m mockResolver) LocationInCache(string) (string, bool)                         { return "", false }
func (m mockResolver) Populate(context.Context, *fault.Bus, string, ...string) error { return nil }

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type ServiceIteratorsSuite struct {
	tester.Suite
	creds account.M365Config
}

func TestServiceIteratorsUnitSuite(t *testing.T) {
	suite.Run(t, &ServiceIteratorsSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ServiceIteratorsSuite) SetupSuite() {
	a := tester.NewMockM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))
	suite.creds = m365
}

func (suite *ServiceIteratorsSuite) TestFilterContainersAndFillCollections() {
	var (
		qp = graph.QueryParams{
			Category:      path.EmailCategory, // doesn't matter which one we use.
			ResourceOwner: inMock.NewProvider("user_id", "user_name"),
			TenantID:      suite.creds.AzureTenantID,
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
			err:      graph.ErrDeletedInFlight,
		}
		container1 = mockContainer{
			id:          strPtr("1"),
			displayName: strPtr("display_name_1"),
			p:           path.Builder{}.Append("1"),
			l:           path.Builder{}.Append("display_name_1"),
		}
		container2 = mockContainer{
			id:          strPtr("2"),
			displayName: strPtr("display_name_2"),
			p:           path.Builder{}.Append("2"),
			l:           path.Builder{}.Append("display_name_2"),
		}
	)

	table := []struct {
		name                  string
		getter                mockGetter
		resolver              graph.ContainerResolver
		scope                 selectors.ExchangeScope
		failFast              control.FailureBehavior
		expectErr             assert.ErrorAssertionFunc
		expectNewColls        int
		expectMetadataColls   int
		expectDoNotMergeColls int
	}{
		{
			name: "happy path, one container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResult,
				},
			},
			resolver:            newMockResolver(container1),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, many containers",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      2,
			expectMetadataColls: 1,
		},
		{
			name: "no containers pass scope",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               selectors.NewExchangeBackup(nil).MailFolders(selectors.None())[0],
			expectErr:           assert.NoError,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "err: deleted in flight",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": deletedInFlightResult,
				},
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
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": errorResult,
				},
			},
			resolver:            newMockResolver(container1),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "half collections error: deleted in flight",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": deletedInFlightResult,
					"2": commonResult,
				},
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
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": errorResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "half collections error: deleted in flight, fail fast",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": deletedInFlightResult,
					"2": commonResult,
				},
			},
			resolver:              newMockResolver(container1, container2),
			scope:                 allScope,
			failFast:              control.FailFast,
			expectErr:             assert.NoError,
			expectNewColls:        2,
			expectMetadataColls:   1,
			expectDoNotMergeColls: 1,
		},
		{
			name: "half collections error: other error, fail fast",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": errorResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			failFast:            control.FailFast,
			expectErr:           assert.Error,
			expectNewColls:      0,
			expectMetadataColls: 0,
		},
	}
	for _, test := range table {
		for _, canMakeDeltaQueries := range []bool{true, false} {
			name := test.name

			if canMakeDeltaQueries {
				name += "-delta"
			} else {
				name += "-non-delta"
			}

			suite.Run(name, func() {
				t := suite.T()

				ctx, flush := tester.NewContext(t)
				defer flush()

				ctrlOpts := control.Options{FailureHandling: test.failFast}
				ctrlOpts.ToggleFeatures.DisableDelta = !canMakeDeltaQueries

				mbh := mockBackupHandler{
					mg:       test.getter,
					category: qp.Category,
				}

				collections, err := filterContainersAndFillCollections(
					ctx,
					qp,
					mbh,
					statusUpdater,
					test.resolver,
					test.scope,
					dps,
					ctrlOpts,
					fault.New(test.failFast == control.FailFast))
				test.expectErr(t, err, clues.ToCore(err))

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
				for k, expect := range test.getter.results {
					coll := collections[k]

					if coll == nil {
						continue
					}

					exColl, ok := coll.(*Collection)
					require.True(t, ok, "collection is an *exchange.Collection")

					ids := [][]string{
						make([]string, 0, len(exColl.added)),
						make([]string, 0, len(exColl.removed)),
					}

					for i, cIDs := range []map[string]struct{}{exColl.added, exColl.removed} {
						for id := range cIDs {
							ids[i] = append(ids[i], id)
						}
					}

					assert.ElementsMatch(t, expect.added, ids[0], "added items")
					assert.ElementsMatch(t, expect.removed, ids[1], "removed items")
				}
			})
		}
	}
}

func checkMetadata(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	cat path.CategoryType,
	expect DeltaPaths,
	c data.BackupCollection,
) {
	catPaths, err := parseMetadataCollections(
		ctx,
		[]data.RestoreCollection{data.NotFoundRestoreCollection{Collection: c}},
		fault.New(true))
	if !assert.NoError(t, err, "getting metadata", clues.ToCore(err)) {
		return
	}

	assert.Equal(t, expect, catPaths[cat])
}

func (suite *ServiceIteratorsSuite) TestFilterContainersAndFillCollections_DuplicateFolders() {
	type scopeCat struct {
		scope selectors.ExchangeScope
		cat   path.CategoryType
	}

	var (
		qp = graph.QueryParams{
			ResourceOwner: inMock.NewProvider("user_id", "user_name"),
			TenantID:      suite.creds.AzureTenantID,
		}

		statusUpdater = func(*support.ConnectorOperationStatus) {}

		dataTypes = []scopeCat{
			{
				scope: selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0],
				cat:   path.EmailCategory,
			},
			{
				scope: selectors.NewExchangeBackup(nil).ContactFolders(selectors.Any())[0],
				cat:   path.ContactsCategory,
			},
			{
				scope: selectors.NewExchangeBackup(nil).EventCalendars(selectors.Any())[0],
				cat:   path.EventsCategory,
			},
		}

		location = path.Builder{}.Append("foo", "bar")

		result1 = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: api.DeltaUpdate{URL: "delta_url"},
		}
		result2 = mockGetterResults{
			added:    []string{"a4", "a5", "a6"},
			removed:  []string{"r4", "r5", "r6"},
			newDelta: api.DeltaUpdate{URL: "delta_url2"},
		}

		container1 = mockContainer{
			id:          strPtr("1"),
			displayName: strPtr("bar"),
			p:           path.Builder{}.Append("1"),
			l:           location,
		}
		container2 = mockContainer{
			id:          strPtr("2"),
			displayName: strPtr("bar"),
			p:           path.Builder{}.Append("2"),
			l:           location,
		}
	)

	oldPath1 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := location.Append("1").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ResourceOwner.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	oldPath2 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := location.Append("2").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ResourceOwner.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	idPath1 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := path.Builder{}.Append("1").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ResourceOwner.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	idPath2 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := path.Builder{}.Append("2").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ResourceOwner.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	table := []struct {
		name           string
		getter         mockGetter
		resolver       graph.ContainerResolver
		inputMetadata  func(t *testing.T, cat path.CategoryType) DeltaPaths
		expectNewColls int
		expectDeleted  int
		expectMetadata func(t *testing.T, cat path.CategoryType) DeltaPaths
	}{
		{
			name: "1 moved to duplicate",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
					"2": result2,
				},
			},
			resolver: newMockResolver(container1, container2),
			inputMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{
					"1": DeltaPath{
						Delta: "old_delta",
						Path:  oldPath1(t, cat).String(),
					},
					"2": DeltaPath{
						Delta: "old_delta",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
			expectMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{
					"1": DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
					"2": DeltaPath{
						Delta: "delta_url2",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
		},
		{
			name: "both move to duplicate",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
					"2": result2,
				},
			},
			resolver: newMockResolver(container1, container2),
			inputMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{
					"1": DeltaPath{
						Delta: "old_delta",
						Path:  oldPath1(t, cat).String(),
					},
					"2": DeltaPath{
						Delta: "old_delta",
						Path:  oldPath2(t, cat).String(),
					},
				}
			},
			expectMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{
					"1": DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
					"2": DeltaPath{
						Delta: "delta_url2",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
		},
		{
			name: "both new",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
					"2": result2,
				},
			},
			resolver: newMockResolver(container1, container2),
			inputMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{}
			},
			expectNewColls: 2,
			expectMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{
					"1": DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
					"2": DeltaPath{
						Delta: "delta_url2",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
		},
		{
			name: "add 1 remove 2",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
				},
			},
			resolver: newMockResolver(container1),
			inputMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{
					"2": DeltaPath{
						Delta: "old_delta",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
			expectNewColls: 1,
			expectDeleted:  1,
			expectMetadata: func(t *testing.T, cat path.CategoryType) DeltaPaths {
				return DeltaPaths{
					"1": DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
				}
			},
		},
	}

	for _, sc := range dataTypes {
		suite.Run(sc.cat.String(), func() {
			qp.Category = sc.cat

			for _, test := range table {
				suite.Run(test.name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					mbh := mockBackupHandler{
						mg:       test.getter,
						category: qp.Category,
					}

					collections, err := filterContainersAndFillCollections(
						ctx,
						qp,
						mbh,
						statusUpdater,
						test.resolver,
						sc.scope,
						test.inputMetadata(t, qp.Category),
						control.Options{FailureHandling: control.FailFast},
						fault.New(true))
					require.NoError(t, err, "getting collections", clues.ToCore(err))

					// collection assertions

					deleteds, news, metadatas := 0, 0, 0
					for _, c := range collections {
						if c.State() == data.DeletedState {
							deleteds++
							continue
						}

						if c.FullPath().Service() == path.ExchangeMetadataService {
							metadatas++
							checkMetadata(t, ctx, qp.Category, test.expectMetadata(t, qp.Category), c)
							continue
						}

						if c.State() == data.NewState {
							news++
						}
					}

					assert.Equal(t, test.expectDeleted, deleteds, "deleted collections")
					assert.Equal(t, test.expectNewColls, news, "new collections")
					assert.Equal(t, 1, metadatas, "metadata collections")

					// items in collections assertions
					for k, expect := range test.getter.results {
						coll := collections[k]

						if coll == nil {
							continue
						}

						exColl, ok := coll.(*Collection)
						require.True(t, ok, "collection is an *exchange.Collection")

						ids := [][]string{
							make([]string, 0, len(exColl.added)),
							make([]string, 0, len(exColl.removed)),
						}

						for i, cIDs := range []map[string]struct{}{exColl.added, exColl.removed} {
							for id := range cIDs {
								ids[i] = append(ids[i], id)
							}
						}

						assert.ElementsMatch(t, expect.added, ids[0], "added items")
						assert.ElementsMatch(t, expect.removed, ids[1], "removed items")
					}
				})
			}
		})
	}
}

func (suite *ServiceIteratorsSuite) TestFilterContainersAndFillCollections_repeatedItems() {
	newDelta := api.DeltaUpdate{URL: "delta_url"}

	table := []struct {
		name          string
		getter        mockGetter
		expectAdded   map[string]struct{}
		expectRemoved map[string]struct{}
	}{
		{
			name: "repeated adds",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": {
						added:    []string{"a1", "a2", "a3", "a1"},
						newDelta: newDelta,
					},
				},
			},
			expectAdded: map[string]struct{}{
				"a1": {},
				"a2": {},
				"a3": {},
			},
			expectRemoved: map[string]struct{}{},
		},
		{
			name: "repeated removes",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": {
						removed:  []string{"r1", "r2", "r3", "r1"},
						newDelta: newDelta,
					},
				},
			},
			expectAdded: map[string]struct{}{},
			expectRemoved: map[string]struct{}{
				"r1": {},
				"r2": {},
				"r3": {},
			},
		},
		{
			name: "remove for same item wins",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": {
						added:    []string{"i1", "a2", "a3"},
						removed:  []string{"i1", "r2", "r3"},
						newDelta: newDelta,
					},
				},
			},
			expectAdded: map[string]struct{}{
				"a2": {},
				"a3": {},
			},
			expectRemoved: map[string]struct{}{
				"i1": {},
				"r2": {},
				"r3": {},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				qp = graph.QueryParams{
					Category:      path.EmailCategory, // doesn't matter which one we use.
					ResourceOwner: inMock.NewProvider("user_id", "user_name"),
					TenantID:      suite.creds.AzureTenantID,
				}
				statusUpdater = func(*support.ConnectorOperationStatus) {}
				allScope      = selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0]
				dps           = DeltaPaths{} // incrementals are tested separately
				container1    = mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("display_name_1"),
					p:           path.Builder{}.Append("1"),
					l:           path.Builder{}.Append("display_name_1"),
				}
				resolver = newMockResolver(container1)
				mbh      = mockBackupHandler{
					mg:       test.getter,
					category: qp.Category,
				}
			)

			require.Equal(t, "user_id", qp.ResourceOwner.ID(), qp.ResourceOwner)
			require.Equal(t, "user_name", qp.ResourceOwner.Name(), qp.ResourceOwner)

			collections, err := filterContainersAndFillCollections(
				ctx,
				qp,
				mbh,
				statusUpdater,
				resolver,
				allScope,
				dps,
				control.Options{FailureHandling: control.FailFast},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

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
			assert.Equal(t, 1, news, "new collections")
			assert.Equal(t, 1, metadatas, "metadata collections")
			assert.Zero(t, doNotMerges, "doNotMerge collections")

			// items in collections assertions
			for k := range test.getter.results {
				coll := collections[k]
				if !assert.NotNilf(t, coll, "missing collection for path %s", k) {
					continue
				}

				exColl, ok := coll.(*Collection)
				require.True(t, ok, "collection is an *exchange.Collection")

				assert.Equal(t, test.expectAdded, exColl.added, "added items")
				assert.Equal(t, test.expectRemoved, exColl.removed, "removed items")
			}
		})
	}
}

func (suite *ServiceIteratorsSuite) TestFilterContainersAndFillCollections_incrementals_nondelta() {
	var (
		userID   = "user_id"
		tenantID = suite.creds.AzureTenantID
		cat      = path.EmailCategory // doesn't matter which one we use,
		qp       = graph.QueryParams{
			Category:      cat,
			ResourceOwner: inMock.NewProvider("user_id", "user_name"),
			TenantID:      suite.creds.AzureTenantID,
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
		p, err := path.Build(tenantID, userID, path.ExchangeService, cat, false, at...)
		require.NoError(t, err, clues.ToCore(err))

		return p
	}

	type endState struct {
		state      data.CollectionState
		doNotMerge bool
	}

	table := []struct {
		name                  string
		getter                mockGetter
		resolver              graph.ContainerResolver
		dps                   DeltaPaths
		expect                map[string]endState
		skipWhenForcedNoDelta bool
	}{
		{
			name: "new container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("new"),
				p:           path.Builder{}.Append("1", "new"),
				l:           path.Builder{}.Append("1", "new"),
			}),
			dps: DeltaPaths{},
			expect: map[string]endState{
				"1": {data.NewState, false},
			},
		},
		{
			name: "not moved container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("not_moved"),
				p:           path.Builder{}.Append("1", "not_moved"),
				l:           path.Builder{}.Append("1", "not_moved"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "not_moved").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NotMovedState, false},
			},
		},
		{
			name: "moved container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("moved"),
				p:           path.Builder{}.Append("1", "moved"),
				l:           path.Builder{}.Append("1", "moved"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "prev").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.MovedState, false},
			},
		},
		{
			name: "deleted container",
			getter: mockGetter{
				results: map[string]mockGetterResults{},
			},
			resolver: newMockResolver(),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
			},
		},
		{
			name: "one deleted, one new",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"2": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("2"),
				displayName: strPtr("new"),
				p:           path.Builder{}.Append("2", "new"),
				l:           path.Builder{}.Append("2", "new"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "one deleted, one new, same path",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"2": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("2"),
				displayName: strPtr("same"),
				p:           path.Builder{}.Append("2", "same"),
				l:           path.Builder{}.Append("2", "same"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "same").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "one moved, one new, same path",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
					"2": commonResults,
				},
			},
			resolver: newMockResolver(
				mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("1", "moved"),
					l:           path.Builder{}.Append("1", "moved"),
				},
				mockContainer{
					id:          strPtr("2"),
					displayName: strPtr("prev"),
					p:           path.Builder{}.Append("2", "prev"),
					l:           path.Builder{}.Append("2", "prev"),
				},
			),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "prev").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.MovedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "bad previous path strings",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("not_moved"),
				p:           path.Builder{}.Append("1", "not_moved"),
				l:           path.Builder{}.Append("1", "not_moved"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  "1/fnords/mc/smarfs",
				},
				"2": DeltaPath{
					Delta: "old_delta_url",
					Path:  "2/fnords/mc/smarfs",
				},
			},
			expect: map[string]endState{
				"1": {data.NewState, false},
			},
		},
		{
			name: "delta expiration",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": expiredResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("same"),
				p:           path.Builder{}.Append("1", "same"),
				l:           path.Builder{}.Append("1", "same"),
			}),
			dps: DeltaPaths{
				"1": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "same").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NotMovedState, true},
			},
			skipWhenForcedNoDelta: true, // this is not a valid test for non-delta
		},
		{
			name: "a little bit of everything",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,  // new
					"2": commonResults,  // notMoved
					"3": commonResults,  // moved
					"4": expiredResults, // moved
					// "5" gets deleted
				},
			},
			resolver: newMockResolver(
				mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("new"),
					p:           path.Builder{}.Append("1", "new"),
					l:           path.Builder{}.Append("1", "new"),
				},
				mockContainer{
					id:          strPtr("2"),
					displayName: strPtr("not_moved"),
					p:           path.Builder{}.Append("2", "not_moved"),
					l:           path.Builder{}.Append("2", "not_moved"),
				},
				mockContainer{
					id:          strPtr("3"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("3", "moved"),
					l:           path.Builder{}.Append("3", "moved"),
				},
				mockContainer{
					id:          strPtr("4"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("4", "moved"),
					l:           path.Builder{}.Append("4", "moved"),
				},
			),
			dps: DeltaPaths{
				"2": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "2", "not_moved").String(),
				},
				"3": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "3", "prev").String(),
				},
				"4": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "4", "prev").String(),
				},
				"5": DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "5", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NewState, false},
				"2": {data.NotMovedState, false},
				"3": {data.MovedState, false},
				"4": {data.MovedState, true},
				"5": {data.DeletedState, false},
			},
			skipWhenForcedNoDelta: true,
		},
	}
	for _, test := range table {
		for _, deltaBefore := range []bool{true, false} {
			for _, deltaAfter := range []bool{true, false} {
				name := test.name

				if deltaAfter {
					name += "-delta"
				} else {
					if test.skipWhenForcedNoDelta {
						suite.T().Skip("intentionally skipped non-delta case")
					}
					name += "-non-delta"
				}

				suite.Run(name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					ctrlOpts := control.Defaults()
					ctrlOpts.ToggleFeatures.DisableDelta = !deltaAfter

					getter := test.getter
					if !deltaAfter {
						getter.noReturnDelta = false
					}

					mbh := mockBackupHandler{
						mg:       test.getter,
						category: qp.Category,
					}

					dps := test.dps
					if !deltaBefore {
						for k, dp := range dps {
							dp.Delta = ""
							dps[k] = dp
						}
					}

					collections, err := filterContainersAndFillCollections(
						ctx,
						qp,
						mbh,
						statusUpdater,
						test.resolver,
						allScope,
						test.dps,
						ctrlOpts,
						fault.New(true))
					assert.NoError(t, err, clues.ToCore(err))

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
	}
}
