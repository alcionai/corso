package exchange

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
)

type CollectionUnitSuite struct {
	tester.Suite
}

func TestCollectionUnitSuite(t *testing.T) {
	suite.Run(t, &CollectionUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionUnitSuite) TestPrefetchedItem_Reader() {
	table := []struct {
		name     string
		readData []byte
	}{
		{
			name:     "HasData",
			readData: []byte("test message"),
		},
		{
			name:     "Empty",
			readData: []byte{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ed, err := data.NewPrefetchedItemWithInfo(
				io.NopCloser(bytes.NewReader(test.readData)),
				"itemID",
				details.ItemInfo{})
			require.NoError(t, err, clues.ToCore(err))

			r, err := readers.NewVersionedRestoreReader(ed.ToReader())
			require.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, readers.DefaultSerializationVersion, r.Format().Version)
			assert.False(t, r.Format().DelInFlight)

			buf := &bytes.Buffer{}
			_, err = buf.ReadFrom(r)
			assert.NoError(t, err, "reading data: %v", clues.ToCore(err))
			assert.Equal(t, test.readData, buf.Bytes(), "read data")
			assert.Equal(t, "itemID", ed.ID(), "item ID")
		})
	}
}

func (suite *CollectionUnitSuite) TestNewCollection_state() {
	type collectionTypes struct {
		name          string
		validModTimes bool
	}

	colTypes := []collectionTypes{
		{
			name: "prefetchCollection",
		},
		{
			name:          "lazyFetchCollection",
			validModTimes: true,
		},
	}

	fooP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))
	barP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "bar")
	require.NoError(suite.T(), err, clues.ToCore(err))

	locPB := path.Builder{}.Append("human-readable")

	table := []struct {
		name   string
		prev   path.Path
		curr   path.Path
		loc    *path.Builder
		expect data.CollectionState
	}{
		{
			name:   "new",
			curr:   fooP,
			loc:    locPB,
			expect: data.NewState,
		},
		{
			name:   "not moved",
			prev:   fooP,
			curr:   fooP,
			loc:    locPB,
			expect: data.NotMovedState,
		},
		{
			name:   "moved",
			prev:   fooP,
			curr:   barP,
			loc:    locPB,
			expect: data.MovedState,
		},
		{
			name:   "deleted",
			prev:   fooP,
			expect: data.DeletedState,
		},
	}

	for _, colType := range colTypes {
		suite.Run(colType.name, func() {
			for _, test := range table {
				suite.Run(test.name, func() {
					t := suite.T()

					c := NewCollection(
						data.NewBaseCollection(
							test.curr,
							test.prev,
							test.loc,
							control.DefaultOptions(),
							false,
							count.New()),
						"u",
						mock.DefaultItemGetSerialize(),
						mock.NeverCanSkipFailChecker(),
						nil,
						nil,
						colType.validModTimes,
						nil,
						count.New())
					assert.Equal(t, test.expect, c.State(), "collection state")
					assert.Equal(t, test.curr, c.FullPath(), "full path")
					assert.Equal(t, test.prev, c.PreviousPath(), "prev path")

					// TODO(ashmrtn): Add LocationPather as part of BackupCollection.
					require.Implements(t, (*data.LocationPather)(nil), c)
					assert.Equal(
						t,
						test.loc,
						c.(data.LocationPather).LocationPath(),
						"location path")
				})
			}
		})
	}
}

func (suite *CollectionUnitSuite) TestGetItemWithRetries() {
	table := []struct {
		name           string
		items          *mock.ItemGetSerialize
		expectErr      func(*testing.T, error)
		expectGetCalls int
	}{
		{
			name:  "happy",
			items: mock.DefaultItemGetSerialize(),
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectGetCalls: 1,
		},
		{
			name:  "an error",
			items: &mock.ItemGetSerialize{GetErr: assert.AnError},
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
			expectGetCalls: 3,
		},
		{
			name: "deleted in flight",
			items: &mock.ItemGetSerialize{
				GetErr: core.ErrNotFound,
			},
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, core.ErrNotFound, "is ErrItemNotFound")
			},
			expectGetCalls: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// itemer is mocked, so only the errors are configured atm.
			_, _, err := test.items.GetItem(ctx, "userID", "itemID", fault.New(true))
			test.expectErr(t, err)
		})
	}
}

func (suite *CollectionUnitSuite) TestPrefetchCollection_Items() {
	var (
		t             = suite.T()
		start         = time.Now().Add(-time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	locPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name            string
		added           map[string]time.Time
		removed         map[string]struct{}
		expectItemCount int
	}{
		{
			name: "no items",
		},
		{
			name: "only added items",
			added: map[string]time.Time{
				"fisher":    {},
				"flannigan": {},
				"fitzbog":   {},
			},
			expectItemCount: 3,
		},
		{
			name: "only removed items",
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount: 3,
		},
		{
			name: "added and removed items",
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t         = suite.T()
				errs      = fault.New(true)
				itemCount int
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			col := NewCollection(
				data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					control.DefaultOptions(),
					false,
					count.New()),
				"",
				&mock.ItemGetSerialize{},
				mock.NeverCanSkipFailChecker(),
				test.added,
				maps.Keys(test.removed),
				false,
				statusUpdater,
				count.New())

			for item := range col.Items(ctx, errs) {
				itemCount++

				_, rok := test.removed[item.ID()]
				if rok {
					assert.True(t, item.Deleted(), "removals should be marked as deleted")
					dimt, ok := item.(data.ItemModTime)
					require.True(t, ok, "item implements data.ItemModTime")
					assert.True(t, dimt.ModTime().After(start), "deleted items should set mod time to now()")
				}

				_, aok := test.added[item.ID()]
				if !rok && aok {
					assert.False(t, item.Deleted(), "additions should not be marked as deleted")
				}

				assert.True(t, aok || rok, "item must be either added or removed: %q", item.ID())
			}

			assert.NoError(t, errs.Failure())
			assert.Equal(
				t,
				test.expectItemCount,
				itemCount,
				"should see all expected items")
		})
	}
}

func (suite *CollectionUnitSuite) TestPrefetchCollection_Items_skipFailure() {
	var (
		start         = time.Now().Add(-time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	table := []struct {
		name               string
		category           path.CategoryType
		handler            backupHandler
		added              map[string]time.Time
		removed            map[string]struct{}
		expectItemCount    int
		expectSkippedCount int
		expectErr          assert.ErrorAssertionFunc
	}{
		{
			name:      "no items",
			category:  path.EventsCategory,
			handler:   newEventBackupHandler(api.Client{}),
			expectErr: assert.NoError,
		},
		{
			name:     "events only added items",
			category: path.EventsCategory,
			handler:  newEventBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"fisher":    {},
				"flannigan": {},
				"fitzbog":   {},
			},
			expectItemCount:    0,
			expectSkippedCount: 3,
			expectErr:          assert.NoError,
		},
		{
			name:     "events only removed items",
			category: path.EventsCategory,
			handler:  newEventBackupHandler(api.Client{}),
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectErr:          assert.NoError,
		},
		{
			name:     "events added and removed items",
			category: path.EventsCategory,
			handler:  newEventBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
			// not 1,  because general is removed from the added
			// map due to being in the removed map
			expectSkippedCount: 0,
			expectErr:          assert.NoError,
		},
		{
			name:     "contacts only added items",
			category: path.ContactsCategory,
			handler:  newContactBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"fisher":    {},
				"flannigan": {},
				"fitzbog":   {},
			},
			expectItemCount:    0,
			expectSkippedCount: 0,
			expectErr:          assert.Error,
		},
		{
			name:     "contacts only removed items",
			category: path.ContactsCategory,
			handler:  newContactBackupHandler(api.Client{}),
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectErr:          assert.NoError,
		},
		{
			name:     "contacts added and removed items",
			category: path.ContactsCategory,
			handler:  newContactBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
			// not 1,  because general is removed from the added
			// map due to being in the removed map
			expectSkippedCount: 0,
			expectErr:          assert.NoError,
		},
		{
			name:     "mail only added items",
			category: path.EmailCategory,
			handler:  newMailBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"fisher":    {},
				"flannigan": {},
				"fitzbog":   {},
			},
			expectItemCount:    0,
			expectSkippedCount: 0,
			expectErr:          assert.Error,
		},
		{
			name:     "mail only removed items",
			category: path.EmailCategory,
			handler:  newMailBackupHandler(api.Client{}),
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectErr:          assert.NoError,
		},
		{
			name:     "mail added and removed items",
			category: path.EmailCategory,
			handler:  newMailBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
			// not 1,  because general is removed from the added
			// map due to being in the removed map
			expectSkippedCount: 0,
			expectErr:          assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t         = suite.T()
				errs      = fault.New(true)
				itemCount int
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			fullPath, err := path.Build("t", "pr", path.ExchangeService, test.category, false, "fnords", "smarf")
			require.NoError(t, err, clues.ToCore(err))

			locPath, err := path.Build("t", "pr", path.ExchangeService, test.category, false, "fnords", "smarf")
			require.NoError(t, err, clues.ToCore(err))

			opts := control.DefaultOptions()
			opts.SkipEventsOnInstance503ForResources = map[string]struct{}{}
			opts.SkipEventsOnInstance503ForResources["pr"] = struct{}{}

			col := NewCollection(
				data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					opts,
					false,
					count.New()),
				"pr",
				&mock.ItemGetSerialize{
					SerializeErr: graph.ErrServiceUnavailableEmptyResp,
				},
				test.handler,
				test.added,
				maps.Keys(test.removed),
				false,
				statusUpdater,
				count.New())

			for item := range col.Items(ctx, errs) {
				itemCount++

				_, rok := test.removed[item.ID()]
				if rok {
					dimt, ok := item.(data.ItemModTime)
					require.True(t, ok, "item implements data.ItemModTime")
					assert.True(t, dimt.ModTime().After(start), "deleted items should set mod time to now()")
					assert.True(t, item.Deleted(), "removals should be marked as deleted")
				}

				_, aok := test.added[item.ID()]
				if !rok && aok {
					assert.False(t, item.Deleted(), "additions should not be marked as deleted")
				}

				assert.True(t, aok || rok, "item must be either added or removed: %q", item.ID())
			}

			test.expectErr(t, errs.Failure())
			assert.Equal(
				t,
				test.expectItemCount,
				itemCount,
				"should see all expected items")
			assert.Len(t, errs.Skipped(), test.expectSkippedCount)
		})
	}
}

// This test verifies skipped error cases are handled correctly by collection enumeration
func (suite *CollectionUnitSuite) TestCollection_SkippedErrors() {
	var (
		t             = suite.T()
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	locPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name              string
		added             map[string]time.Time
		expectItemCount   int
		itemGetter        itemGetterSerializer
		expectedSkipError *fault.Skipped
	}{
		{
			name: "ErrorInvalidRecipients",
			added: map[string]time.Time{
				"fisher": {},
			},
			expectItemCount: 0,
			itemGetter: &mock.ItemGetSerialize{
				GetErr: graphTD.ODataErr(string(graph.ErrorInvalidRecipients)),
			},
			expectedSkipError: fault.EmailSkip(fault.SkipInvalidRecipients, "", "fisher", nil),
		},
		{
			name: "ErrorCorruptData",
			added: map[string]time.Time{
				"fisher": {},
			},
			expectItemCount: 0,
			itemGetter: &mock.ItemGetSerialize{
				GetErr: graphTD.ODataErr(string(graph.ErrorCorruptData)),
			},
			expectedSkipError: fault.EmailSkip(fault.SkipCorruptData, "", "fisher", nil),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t         = suite.T()
				errs      = fault.New(true)
				itemCount int
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			col := NewCollection(
				data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					control.DefaultOptions(),
					false,
					count.New()),
				"",
				test.itemGetter,
				mock.NeverCanSkipFailChecker(),
				test.added,
				nil,
				false,
				statusUpdater,
				count.New())

			for range col.Items(ctx, errs) {
				itemCount++
			}

			assert.NoError(t, errs.Failure())
			if test.expectedSkipError != nil {
				assert.Len(t, errs.Skipped(), 1)
				skippedItem := errs.Skipped()[0].Item

				assert.Equal(t, skippedItem.Cause, test.expectedSkipError.Item.Cause)
				assert.Equal(t, skippedItem.ID, test.expectedSkipError.Item.ID)
			}

			assert.Equal(
				t,
				test.expectItemCount,
				itemCount,
				"should see all expected items")
		})
	}
}

type mockLazyItemGetterSerializer struct {
	*mock.ItemGetSerialize
	callIDs []string
}

func (mlg *mockLazyItemGetterSerializer) GetItem(
	ctx context.Context,
	user string,
	itemID string,
	errs *fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	mlg.callIDs = append(mlg.callIDs, itemID)
	return mlg.ItemGetSerialize.GetItem(ctx, user, itemID, errs)
}

func (mlg *mockLazyItemGetterSerializer) check(t *testing.T, expectIDs []string) {
	assert.ElementsMatch(t, expectIDs, mlg.callIDs)
}

func (suite *CollectionUnitSuite) TestLazyFetchCollection_Items_LazyFetch() {
	var (
		t             = suite.T()
		start         = time.Now().Add(-time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	locPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name            string
		added           map[string]time.Time
		removed         map[string]struct{}
		expectItemCount int
		expectReads     []string
	}{
		{
			name: "no items",
		},
		{
			name: "only added items",
			added: map[string]time.Time{
				"fisher":    start.Add(time.Minute),
				"flannigan": start.Add(2 * time.Minute),
				"fitzbog":   start.Add(3 * time.Minute),
			},
			expectItemCount: 3,
			expectReads: []string{
				"fisher",
				"flannigan",
				"fitzbog",
			},
		},
		{
			name: "only removed items",
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount: 3,
		},
		{
			name: "added and removed items",
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t         = suite.T()
				errs      = fault.New(true)
				itemCount int
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			mlg := &mockLazyItemGetterSerializer{
				ItemGetSerialize: &mock.ItemGetSerialize{},
			}
			defer mlg.check(t, test.expectReads)

			col := NewCollection(
				data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					control.DefaultOptions(),
					false,
					count.New()),
				"",
				mlg,
				mock.NeverCanSkipFailChecker(),
				test.added,
				maps.Keys(test.removed),
				true,
				statusUpdater,
				count.New())

			for item := range col.Items(ctx, errs) {
				itemCount++

				_, rok := test.removed[item.ID()]
				if rok {
					dimt, ok := item.(data.ItemModTime)
					require.True(t, ok, "item implements data.ItemModTime")
					assert.True(t, dimt.ModTime().After(start), "deleted items should set mod time to now()")
					assert.True(t, item.Deleted(), "removals should be marked as deleted")
				}

				modTime, aok := test.added[item.ID()]
				if !rok && aok {
					// Item's mod time should be what's passed into the collection
					// initializer.
					assert.Implements(t, (*data.ItemModTime)(nil), item)
					assert.Equal(t, modTime, item.(data.ItemModTime).ModTime(), "item mod time")
					assert.False(t, item.Deleted(), "additions should not be marked as deleted")

					// Check if the test want's us to read the item's data so the lazy
					// data fetch is executed.
					if slices.Contains(test.expectReads, item.ID()) {
						r := item.ToReader()

						_, err := io.ReadAll(r)
						assert.NoError(t, err, clues.ToCore(err))

						r.Close()

						assert.Implements(t, (*data.ItemInfo)(nil), item)
						info, err := item.(data.ItemInfo).Info()

						// ItemInfo's mod time should match what was passed into the
						// collection initializer.
						assert.NoError(t, err, clues.ToCore(err))
						assert.Equal(t, modTime, info.Modified(), "ItemInfo mod time")
					} else {
						assert.Fail(t, "unexpected read on item %s", item.ID())
					}
				}

				assert.True(t, aok || rok, "item must be either added or removed: %q", item.ID())
			}

			assert.NoError(t, errs.Failure())
			assert.Equal(
				t,
				test.expectItemCount,
				itemCount,
				"should see all expected items")
		})
	}
}

func (suite *CollectionUnitSuite) TestLazyFetchCollection_Items_skipFailure() {
	var (
		start         = time.Now().Add(-time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
		expectSkip    = func(t *testing.T, err error) {
			assert.Error(t, err, clues.ToCore(err))
			assert.ErrorContains(t, err, "skip")
			assert.True(t, clues.HasLabel(err, graph.LabelsSkippable), clues.ToCore(err))
		}
		expectNotSkipped = func(t *testing.T, err error) {
			assert.Error(t, err, clues.ToCore(err))
			assert.NotContains(t, err.Error(), "skip")
		}
	)

	table := []struct {
		name               string
		added              map[string]time.Time
		removed            map[string]struct{}
		category           path.CategoryType
		handler            backupHandler
		expectItemCount    int
		expectSkippedCount int
		expectReads        []string
		expectErr          func(t *testing.T, err error)
		expectFailure      assert.ErrorAssertionFunc
	}{
		{
			name:          "no items",
			category:      path.EventsCategory,
			handler:       newEventBackupHandler(api.Client{}),
			expectFailure: assert.NoError,
		},
		{
			name:     "events only added items",
			category: path.EventsCategory,
			handler:  newEventBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"fisher":    start.Add(time.Minute),
				"flannigan": start.Add(2 * time.Minute),
				"fitzbog":   start.Add(3 * time.Minute),
			},
			expectItemCount:    3,
			expectSkippedCount: 3,
			expectReads: []string{
				"fisher",
				"flannigan",
				"fitzbog",
			},
			expectErr:     expectSkip,
			expectFailure: assert.NoError,
		},
		{
			name:     "events only removed items",
			category: path.EventsCategory,
			handler:  newEventBackupHandler(api.Client{}),
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectErr:          expectSkip,
			expectFailure:      assert.NoError,
		},
		{
			name:     "events added and removed items",
			category: path.EventsCategory,
			handler:  newEventBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
			// not 1,  because general is removed from the added
			// map due to being in the removed map
			expectSkippedCount: 0,
			expectErr:          expectSkip,
			expectFailure:      assert.NoError,
		},
		{
			name:     "contacts only added items",
			category: path.ContactsCategory,
			handler:  newContactBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"fisher":    start.Add(time.Minute),
				"flannigan": start.Add(2 * time.Minute),
				"fitzbog":   start.Add(3 * time.Minute),
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectReads: []string{
				"fisher",
				"flannigan",
				"fitzbog",
			},
			expectErr:     expectNotSkipped,
			expectFailure: assert.Error,
		},
		{
			name:     "contacts only removed items",
			category: path.ContactsCategory,
			handler:  newContactBackupHandler(api.Client{}),
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectErr:          expectNotSkipped,
			expectFailure:      assert.NoError,
		},
		{
			name:     "contacts added and removed items",
			category: path.ContactsCategory,
			handler:  newContactBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
			// not 1,  because general is removed from the added
			// map due to being in the removed map
			expectSkippedCount: 0,
			expectErr:          expectNotSkipped,
			expectFailure:      assert.NoError,
		},
		{
			name:     "mail only added items",
			category: path.EmailCategory,
			handler:  newMailBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"fisher":    start.Add(time.Minute),
				"flannigan": start.Add(2 * time.Minute),
				"fitzbog":   start.Add(3 * time.Minute),
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectReads: []string{
				"fisher",
				"flannigan",
				"fitzbog",
			},
			expectErr:     expectNotSkipped,
			expectFailure: assert.Error,
		},
		{
			name:     "mail only removed items",
			category: path.EmailCategory,
			handler:  newMailBackupHandler(api.Client{}),
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
			expectItemCount:    3,
			expectSkippedCount: 0,
			expectErr:          expectNotSkipped,
			expectFailure:      assert.NoError,
		},
		{
			name:     "mail added and removed items",
			category: path.EmailCategory,
			handler:  newMailBackupHandler(api.Client{}),
			added: map[string]time.Time{
				"general": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 3,
			// not 1,  because general is removed from the added
			// map due to being in the removed map
			expectSkippedCount: 0,
			expectErr:          expectNotSkipped,
			expectFailure:      assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t         = suite.T()
				errs      = fault.New(false)
				itemCount int
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			fullPath, err := path.Build("t", "pr", path.ExchangeService, test.category, false, "fnords", "smarf")
			require.NoError(t, err, clues.ToCore(err))

			locPath, err := path.Build("t", "pr", path.ExchangeService, test.category, false, "fnords", "smarf")
			require.NoError(t, err, clues.ToCore(err))

			mlg := &mockLazyItemGetterSerializer{
				ItemGetSerialize: &mock.ItemGetSerialize{
					SerializeErr: graph.ErrServiceUnavailableEmptyResp,
				},
			}
			defer mlg.check(t, test.expectReads)

			opts := control.DefaultOptions()
			opts.SkipEventsOnInstance503ForResources = map[string]struct{}{}
			opts.SkipEventsOnInstance503ForResources["pr"] = struct{}{}

			col := NewCollection(
				data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					opts,
					false,
					count.New()),
				"pr",
				mlg,
				test.handler,
				test.added,
				maps.Keys(test.removed),
				true,
				statusUpdater,
				count.New())

			for item := range col.Items(ctx, errs) {
				itemCount++

				_, rok := test.removed[item.ID()]
				if rok {
					dimt, ok := item.(data.ItemModTime)
					require.True(t, ok, "item implements data.ItemModTime")
					assert.True(t, dimt.ModTime().After(start), "deleted items should set mod time to now()")
					assert.True(t, item.Deleted(), "removals should be marked as deleted")
				}

				modTime, aok := test.added[item.ID()]
				if !rok && aok {
					// Item's mod time should be what's passed into the collection
					// initializer.
					assert.Implements(t, (*data.ItemModTime)(nil), item)
					assert.Equal(t, modTime, item.(data.ItemModTime).ModTime(), "item mod time")
					assert.False(t, item.Deleted(), "additions should not be marked as deleted")

					// Check if the test want's us to read the item's data so the lazy
					// data fetch is executed.
					if slices.Contains(test.expectReads, item.ID()) {
						r := item.ToReader()

						_, err := io.ReadAll(r)
						test.expectErr(t, err)

						r.Close()
					} else {
						assert.Fail(t, "unexpected read on item %s", item.ID())
					}
				}

				assert.True(t, aok || rok, "item must be either added or removed: %q", item.ID())
			}

			failure := errs.Failure()
			if failure == nil && len(errs.Recovered()) > 0 {
				failure = errs.Recovered()[0]
			}

			test.expectFailure(t, failure, clues.ToCore(failure))
			assert.Equal(
				t,
				test.expectItemCount,
				itemCount,
				"should see all expected items")
			assert.Len(t, errs.Skipped(), test.expectSkippedCount)
		})
	}
}

func (suite *CollectionUnitSuite) TestLazyItem_NoRead_GetInfo_Errors() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	li := data.NewLazyItemWithInfo(
		ctx,
		nil,
		"itemID",
		time.Now(),
		count.New(),
		fault.New(true))

	_, err := li.Info()
	assert.Error(suite.T(), err, "Info without reading data should error")
}

func (suite *CollectionUnitSuite) TestLazyItem_GetDataErrors() {
	var (
		parentPath = "inbox/private/silly cats"
		now        = time.Now()
	)

	table := []struct {
		name              string
		getErr            error
		serializeErr      error
		expectReadErrType error
	}{
		{
			name:              "ReturnsErrorOnGenericGetError",
			getErr:            assert.AnError,
			expectReadErrType: assert.AnError,
		},
		{
			name:              "ReturnsErrorOnGenericSerializeError",
			serializeErr:      assert.AnError,
			expectReadErrType: assert.AnError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var testData serialization.Parsable

			if test.getErr == nil {
				// Exact data type doesn't really matter.
				item := models.NewMessage()
				item.SetSubject(ptr.To("hello world"))

				testData = item
			}

			getter := &mock.ItemGetSerialize{
				GetData:      testData,
				GetErr:       test.getErr,
				SerializeErr: test.serializeErr,
			}

			li := data.NewLazyItemWithInfo(
				ctx,
				&lazyItemGetter{
					userID:       "userID",
					itemID:       "itemID",
					getter:       getter,
					modTime:      now,
					immutableIDs: false,
					parentPath:   parentPath,
				},
				"itemID",
				now,
				count.New(),
				fault.New(true))

			assert.False(t, li.Deleted(), "item shouldn't be marked deleted")
			assert.Equal(t, now, li.ModTime(), "item mod time")

			_, err := readers.NewVersionedRestoreReader(li.ToReader())
			assert.ErrorIs(t, err, test.expectReadErrType)

			// Should get some form of error when trying to get info.
			_, err = li.Info()
			assert.Error(t, err, "Info()")
		})
	}
}

func (suite *CollectionUnitSuite) TestLazyItem_ReturnsEmptyReaderOnDeletedInFlight() {
	var (
		t = suite.T()

		parentPath = "inbox/private/silly cats"
		now        = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	getter := &mock.ItemGetSerialize{GetErr: core.ErrNotFound}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter{
			userID:       "userID",
			itemID:       "itemID",
			getter:       getter,
			modTime:      now,
			immutableIDs: false,
			parentPath:   parentPath,
		},
		"itemID",
		now,
		count.New(),
		fault.New(true))

	assert.False(t, li.Deleted(), "item shouldn't be marked deleted")
	assert.Equal(
		t,
		now,
		li.ModTime(),
		"item mod time")

	r, err := readers.NewVersionedRestoreReader(li.ToReader())
	require.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, readers.DefaultSerializationVersion, r.Format().Version)
	assert.True(t, r.Format().DelInFlight)

	readData, err := io.ReadAll(r)
	assert.NoError(t, err, "reading item data: %v", clues.ToCore(err))

	assert.Empty(t, readData, "read item data")

	_, err = li.Info()
	assert.ErrorIs(t, err, data.ErrNotFound, "Info() error")
}

func (suite *CollectionUnitSuite) TestLazyItem() {
	var (
		t = suite.T()

		parentPath = "inbox/private/silly cats"
		now        = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	// Exact data type doesn't really matter.
	testData := models.NewMessage()
	testData.SetSubject(ptr.To("hello world"))

	getter := &mock.ItemGetSerialize{GetData: testData}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter{
			userID:       "userID",
			itemID:       "itemID",
			getter:       getter,
			modTime:      now,
			immutableIDs: false,
			parentPath:   parentPath,
		},
		"itemID",
		now,
		count.New(),
		fault.New(true))

	assert.False(t, li.Deleted(), "item shouldn't be marked deleted")
	assert.Equal(
		t,
		now,
		li.ModTime(),
		"item mod time")

	r, err := readers.NewVersionedRestoreReader(li.ToReader())
	require.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, readers.DefaultSerializationVersion, r.Format().Version)
	assert.False(t, r.Format().DelInFlight)

	readData, err := io.ReadAll(r)
	assert.NoError(t, err, "reading item data: %v", clues.ToCore(err))

	assert.NotEmpty(t, readData, "read item data")

	info, err := li.Info()
	assert.NoError(t, err, "getting item info: %v", clues.ToCore(err))

	assert.Equal(t, parentPath, info.Exchange.ParentPath)
	assert.Equal(t, now, info.Modified())
}
