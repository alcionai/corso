package groups

import (
	"bytes"
	"context"
	"io"
	"slices"
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
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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
				nil,
				false)

			assert.Equal(t, test.expect, c.State(), "collection state")
			assert.Equal(t, test.curr, c.FullPath(), "full path")
			assert.Equal(t, test.prev, c.PreviousPath(), "prev path")

			prefetch, ok := c.(*prefetchCollection[models.Channelable, models.ChatMessageable])
			require.True(t, ok, "collection type")

			assert.Equal(t, test.loc, prefetch.LocationPath(), "location path")
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

func (suite *CollectionUnitSuite) TestPrefetchCollection_streamItems() {
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
		added   map[string]time.Time
		removed map[string]struct{}
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
		},
		{
			name: "only removed items",
			removed: map[string]struct{}{
				"princess": {},
				"poppy":    {},
				"petunia":  {},
			},
		},
		{
			name: "added and removed items",
			added: map[string]time.Time{
				"goblin": {},
			},
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

			col := &prefetchCollection[models.Channelable, models.ChatMessageable]{
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

type getAndAugmentConversation struct {
	GetItemErr error
	CallIDs    []string
}

//lint:ignore U1000 false linter issue due to generics
func (m *getAndAugmentConversation) getItem(
	_ context.Context,
	_ string,
	_ path.Elements,
	postID string,
) (models.Postable, *details.GroupsInfo, error) {
	m.CallIDs = append(m.CallIDs, postID)

	p := models.NewPost()
	p.SetId(ptr.To(postID))

	return p, &details.GroupsInfo{}, m.GetItemErr
}

//
//lint:ignore U1000 false linter issue due to generics
func (m *getAndAugmentConversation) augmentItemInfo(*details.GroupsInfo, models.Conversationable) {
	// no-op
}

func (m *getAndAugmentConversation) check(t *testing.T, expected []string) {
	// Sort before comparing. We could use a set, but that would prevent us from
	// detecting duplicates.
	slices.Sort(m.CallIDs)
	slices.Sort(expected)

	assert.Equal(t, expected, m.CallIDs, "expected calls")
}

func (suite *CollectionUnitSuite) TestLazyFetchCollection_Items_LazyFetch() {
	var (
		t             = suite.T()
		start         = time.Now().Add(-time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.Build(
		"t", "pr", path.GroupsService, path.ConversationPostsCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	locPath, err := path.Build(
		"t", "pr", path.GroupsService, path.ConversationPostsCategory, false, "fnords", "smarf")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name            string
		added           map[string]time.Time
		removed         map[string]struct{}
		expectItemCount int
		// Items we want to trigger lazy reader on.
		expectReads []string
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
			// TODO(pandeyabs): Overlaps between added and removed are deleted
			// by NewCollection caller code. This is a slight deviation from how
			// exchange does it. It's harmless but should be fixed for consistency.
			//
			// Since we are calling NewCollection here directly, we are not testing
			// with overlaps, else those tests with fail. Same behavior exists for
			// prefetch collections.
			name: "added and removed items",
			added: map[string]time.Time{
				"goblin": {},
			},
			removed: map[string]struct{}{
				"general":  {},
				"goose":    {},
				"grumbles": {},
			},
			expectItemCount: 4,
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

			getterAugmenter := &getAndAugmentConversation{}
			defer getterAugmenter.check(t, test.expectReads)

			col := &lazyFetchCollection[models.Conversationable, models.Postable]{
				BaseCollection: data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					control.DefaultOptions(),
					false,
					count.New()),
				added:         test.added,
				contains:      container[models.Conversationable]{},
				removed:       test.removed,
				getAndAugment: getterAugmenter,
				stream:        make(chan data.Item),
				statusUpdater: statusUpdater,
			}

			for item := range col.Items(ctx, errs) {
				itemCount++

				_, rok := test.removed[item.ID()]
				if rok {
					assert.True(t, item.Deleted(), "removals should be marked as deleted")
					dimt, ok := item.(data.ItemModTime)
					require.True(t, ok, "item implements data.ItemModTime")
					assert.True(t, dimt.ModTime().After(start), "deleted items should set mod time to now()")
				}

				modTime, aok := test.added[item.ID()]
				if !rok && aok {
					// Item's mod time should be what's passed into the collection
					// initializer.
					assert.Implements(t, (*data.ItemModTime)(nil), item)
					assert.Equal(t, modTime, item.(data.ItemModTime).ModTime(), "item mod time")

					assert.False(t, item.Deleted(), "additions should not be marked as deleted")

					// Check if the test wants us to read the item's data so the lazy
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

func (suite *CollectionUnitSuite) TestLazyItem_GetDataErrors() {
	var (
		parentPath = "thread/private/silly cats"
		now        = time.Now()
	)

	table := []struct {
		name              string
		getErr            error
		expectReadErrType error
	}{
		{
			name:              "ReturnsErrorOnGenericGetError",
			getErr:            assert.AnError,
			expectReadErrType: assert.AnError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			m := getAndAugmentConversation{
				GetItemErr: test.getErr,
			}

			li := data.NewLazyItemWithInfo(
				ctx,
				&lazyItemGetter[models.Conversationable, models.Postable]{
					resourceID:    "resourceID",
					itemID:        "itemID",
					getAndAugment: &m,
					modTime:       now,
					parentPath:    parentPath,
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

		parentPath = "thread/private/silly cats"
		now        = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	m := getAndAugmentConversation{
		GetItemErr: graph.ErrDeletedInFlight,
	}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter[models.Conversationable, models.Postable]{
			resourceID:    "resourceID",
			itemID:        "itemID",
			getAndAugment: &m,
			modTime:       now,
			parentPath:    parentPath,
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

	_, err := readers.NewVersionedRestoreReader(li.ToReader())
	assert.ErrorIs(t, err, graph.ErrDeletedInFlight, "item should be marked deleted in flight")
}

func (suite *CollectionUnitSuite) TestLazyItem() {
	var (
		t = suite.T()

		parentPath = "thread/private/silly cats"
		now        = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	m := getAndAugmentConversation{}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter[models.Conversationable, models.Postable]{
			resourceID:    "resourceID",
			itemID:        "itemID",
			getAndAugment: &m,
			modTime:       now,
			parentPath:    parentPath,
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

	assert.Equal(t, parentPath, info.Groups.ParentPath)
	assert.Equal(t, now, info.Modified())
}
