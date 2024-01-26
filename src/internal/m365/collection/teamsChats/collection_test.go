package teamschats

import (
	"bytes"
	"context"
	"io"
	"slices"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/teamschats/testdata"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/errs/core"
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
	fooP, err := path.Build("t", "u", path.TeamsChatsService, path.ChatsCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))
	barP, err := path.Build("t", "u", path.TeamsChatsService, path.ChatsCategory, false, "bar")
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

			c := NewCollection[models.Chatable](
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
				container[models.Chatable]{},
				nil)

			assert.Equal(t, test.expect, c.State(), "collection state")
			assert.Equal(t, test.curr, c.FullPath(), "full path")
			assert.Equal(t, test.prev, c.PreviousPath(), "prev path")

			prefetch, ok := c.(*lazyFetchCollection[models.Chatable])
			require.True(t, ok, "collection type")

			assert.Equal(t, test.loc, prefetch.LocationPath(), "location path")
		})
	}
}

type getAndAugmentChat struct {
	err error
}

//lint:ignore U1000 false linter issue due to generics
func (m getAndAugmentChat) getItem(
	_ context.Context,
	chat models.Chatable,
) (models.Chatable, *details.TeamsChatsInfo, error) {
	chat.SetTopic(chat.GetId())

	return chat, &details.TeamsChatsInfo{}, m.err
}

func (suite *CollectionUnitSuite) TestLazyFetchCollection_Items_LazyFetch() {
	var (
		t             = suite.T()
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.BuildPrefix("t", "pr", path.TeamsChatsService, path.ChatsCategory)
	require.NoError(t, err, clues.ToCore(err))

	locPath, err := path.BuildPrefix("t", "pr", path.TeamsChatsService, path.ChatsCategory)
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name            string
		items           []models.Chatable
		expectItemCount int
		// Items we want to trigger lazy reader on.
		expectReads []string
	}{
		{
			name: "no items",
		},
		{
			name:            "items",
			items:           testdata.StubChats("fisher", "flannigan", "fitzbog"),
			expectItemCount: 3,
			expectReads: []string{
				"fisher",
				"flannigan",
				"fitzbog",
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

			getterAugmenter := &getAndAugmentChat{}

			col := &lazyFetchCollection[models.Chatable]{
				BaseCollection: data.NewBaseCollection(
					fullPath,
					nil,
					locPath.ToBuilder(),
					control.DefaultOptions(),
					false,
					count.New()),
				items:         test.items,
				contains:      container[models.Chatable]{},
				getter:        getterAugmenter,
				stream:        make(chan data.Item),
				statusUpdater: statusUpdater,
			}

			for item := range col.Items(ctx, errs) {
				itemCount++

				ok := slices.ContainsFunc(test.items, func(mc models.Chatable) bool {
					return ptr.Val(mc.GetId()) == item.ID()
				})

				require.True(t, ok, "item must be either added or removed: %q", item.ID())
				assert.False(t, item.Deleted(), "additions should not be marked as deleted")
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
		parentPath = ""
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

			chat := testdata.StubChats(uuid.NewString())[0]

			m := getAndAugmentChat{
				err: test.getErr,
			}

			li := data.NewLazyItemWithInfo(
				ctx,
				&lazyItemGetter[models.Chatable]{
					resourceID: "resourceID",
					item:       chat,
					getter:     &m,
					modTime:    now,
					parentPath: parentPath,
				},
				ptr.Val(chat.GetId()),
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

		parentPath = ""
		now        = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	chat := testdata.StubChats(uuid.NewString())[0]

	m := getAndAugmentChat{
		err: core.ErrNotFound,
	}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter[models.Chatable]{
			resourceID: "resourceID",
			item:       chat,
			getter:     &m,
			modTime:    now,
			parentPath: parentPath,
		},
		ptr.Val(chat.GetId()),
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
	assert.ErrorIs(t, err, core.ErrNotFound, "item should be marked deleted in flight")
}

func (suite *CollectionUnitSuite) TestLazyItem() {
	var (
		t = suite.T()

		parentPath = ""
		now        = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	chat := testdata.StubChats(uuid.NewString())[0]
	m := getAndAugmentChat{}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter[models.Chatable]{
			resourceID: "resourceID",
			item:       chat,
			getter:     &m,
			modTime:    now,
			parentPath: parentPath,
		},
		ptr.Val(chat.GetId()),
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

	assert.Empty(t, parentPath)
	assert.Equal(t, now, info.Modified())
}
