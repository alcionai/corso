package groups

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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
	fooP, err := path.Build("t", "u", path.GroupsService, path.ChannelMessagesCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))
	barP, err := path.Build("t", "u", path.GroupsService, path.ChannelMessagesCategory, false, "bar")
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

			c := NewCollection[models.Channelable, models.ChatMessageable](
				data.NewBaseCollection(
					test.curr,
					test.prev,
					test.loc,
					control.DefaultOptions(),
					false,
					count.New()),
				nil,
				"g",
				nil,
				nil,
				container[models.Channelable]{},
				nil)
			assert.Equal(t, test.expect, c.State(), "collection state")
			assert.Equal(t, test.curr, c.FullPath(), "full path")
			assert.Equal(t, test.prev, c.PreviousPath(), "prev path")
			assert.Equal(t, test.loc, c.LocationPath(), "location path")
		})
	}
}

type getAndAugmentChannelMessage struct {
	Err error
}

//lint:ignore U1000 false linter issue due to generics
func (m getAndAugmentChannelMessage) getItem(
	_ context.Context,
	_ string,
	_ path.Elements,
	itemID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	msg := models.NewChatMessage()
	msg.SetId(ptr.To(itemID))

	return msg, &details.GroupsInfo{}, m.Err
}

//lint:ignore U1000 false linter issue due to generics
func (getAndAugmentChannelMessage) augmentItemInfo(*details.GroupsInfo, models.Channelable) {
	// no-op
}

func (suite *CollectionUnitSuite) TestCollection_streamItems() {
	var (
		t             = suite.T()
		start         = time.Now().Add(-1 * time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.Build("t", "pr", path.GroupsService, path.ChannelMessagesCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	locPath, err := path.Build("t", "pr", path.GroupsService, path.ChannelMessagesCategory, false, "fnords", "smarf")
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

			col := &Collection[models.Channelable, models.ChatMessageable]{
				BaseCollection: data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					control.DefaultOptions(),
					false,
					count.New()),
				added:         test.added,
				contains:      container[models.Channelable]{},
				removed:       test.removed,
				getAndAugment: getAndAugmentChannelMessage{},
				stream:        make(chan data.Item),
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
