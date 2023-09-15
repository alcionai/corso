package exchange

import (
	"bytes"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type CollectionUnitSuite struct {
	tester.Suite
}

func TestCollectionUnitSuite(t *testing.T) {
	suite.Run(t, &CollectionUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionUnitSuite) TestReader_Valid() {
	m := []byte("test message")
	description := "aFile"
	ed := &Item{id: description, message: m}

	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(ed.ToReader())
	assert.NoError(suite.T(), err, clues.ToCore(err))
	assert.Equal(suite.T(), buf.Bytes(), m)
	assert.Equal(suite.T(), description, ed.ID())
}

func (suite *CollectionUnitSuite) TestReader_Empty() {
	var (
		empty    []byte
		expected int64
		t        = suite.T()
	)

	ed := &Item{message: empty}
	buf := &bytes.Buffer{}
	received, err := buf.ReadFrom(ed.ToReader())

	assert.Equal(t, expected, received)
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *CollectionUnitSuite) TestCollection_NewCollection() {
	t := suite.T()
	tenant := "a-tenant"
	user := "a-user"
	folder := "a-folder"
	name := "User"

	fullPath, err := path.Build(
		tenant,
		user,
		path.ExchangeService,
		path.EmailCategory,
		false,
		folder)
	require.NoError(t, err, clues.ToCore(err))

	edc := Collection{
		user:     name,
		fullPath: fullPath,
	}
	assert.Equal(t, name, edc.user)
	assert.Equal(t, fullPath, edc.FullPath())
}

func (suite *CollectionUnitSuite) TestNewCollection_state() {
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
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			c := NewCollection(
				"u",
				test.curr, test.prev, test.loc,
				0,
				mock.DefaultItemGetSerialize(),
				nil,
				control.DefaultOptions(),
				false)
			assert.Equal(t, test.expect, c.State(), "collection state")
			assert.Equal(t, test.curr, c.fullPath, "full path")
			assert.Equal(t, test.prev, c.prevPath, "prev path")
			assert.Equal(t, test.loc, c.locationPath, "location path")
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
				GetErr: graph.ErrDeletedInFlight,
			},
			expectErr: func(t *testing.T, err error) {
				assert.True(t, graph.IsErrDeletedInFlight(err), "is ErrDeletedInFlight")
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
			_, _, err := test.items.GetItem(ctx, "userID", "itemID", false, fault.New(true))
			test.expectErr(t, err)
		})
	}
}

func (suite *CollectionUnitSuite) TestCollection_streamItems() {
	var (
		t             = suite.T()
		start         = time.Now().Add(-1 * time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	locPath, err := path.Build("t", "pr", path.ExchangeService, path.EmailCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name    string
		added   map[string]struct{}
		removed map[string]struct{}
	}{
		{
			name:    "no items",
			added:   map[string]struct{}{},
			removed: map[string]struct{}{},
		},
		{
			name: "only added items",
			added: map[string]struct{}{
				"fisher":    {},
				"flannigan": {},
				"fitzbog":   {},
			},
			removed: map[string]struct{}{},
		},
		{
			name:  "only removed items",
			added: map[string]struct{}{},
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
		},
		{
			name:  "added and removed items",
			added: map[string]struct{}{},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
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

			col := &Collection{
				added:         test.added,
				removed:       test.removed,
				ctrl:          control.DefaultOptions(),
				getter:        &mock.ItemGetSerialize{},
				stream:        make(chan data.Item),
				fullPath:      fullPath,
				locationPath:  locPath.ToBuilder(),
				statusUpdater: statusUpdater,
			}

			go col.streamItems(ctx, errs)

			for item := range col.stream {
				itemCount++

				_, aok := test.added[item.ID()]
				if aok {
					assert.False(t, item.Deleted(), "additions should not be marked as deleted")
				}

				_, rok := test.removed[item.ID()]
				if rok {
					assert.True(t, item.Deleted(), "removals should be marked as deleted")
					dimt, ok := item.(data.ItemModTime)
					require.True(t, ok, "item implements data.ItemModTime")
					assert.True(t, dimt.ModTime().After(start), "deleted items should set mod time to now()")
				}

				assert.True(t, aok || rok, "item must be either added or removed: %q", item.ID())
			}

			assert.NoError(t, errs.Failure())
			assert.Equal(
				t,
				len(test.added)+len(test.removed),
				itemCount,
				"should see all expected items")
		})
	}
}
