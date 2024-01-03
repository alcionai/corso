package site

import (
	"bytes"
	"io"
	"slices"
	"testing"
	"time"

	"github.com/alcionai/clues"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/site/mock"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	spMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type SharePointCollectionSuite struct {
	tester.Suite
	siteID string
	creds  account.M365Config
	ac     api.Client
}

func (suite *SharePointCollectionSuite) SetupSuite() {
	t := suite.T()

	suite.siteID = tconfig.M365SiteID(t)
	a := tconfig.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = m365

	ac, err := api.NewClient(
		m365,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	suite.ac = ac
}

func TestSharePointCollectionSuite(t *testing.T) {
	suite.Run(t, &SharePointCollectionSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

// TestListCollection tests basic functionality to create
// SharePoint collection and to use the data stream channel.
func (suite *SharePointCollectionSuite) TestPrefetchCollection_Items() {
	var (
		tenant  = "some"
		user    = "user"
		dirRoot = "directory"
	)

	sel := selectors.NewSharePointBackup([]string{"site"})

	tables := []struct {
		name, itemName string
		cat            path.CategoryType
		scope          selectors.SharePointScope
		cat            path.CategoryType
		getter         getItemByIDer
		getDir         func(t *testing.T) path.Path
		getItem        func(t *testing.T, itemName string) data.Item
	}{
		{
			name:     "List",
			itemName: "MockListing",
			cat:      path.ListsCategory,
			scope:    sel.Lists(selectors.Any())[0],
			cat:      path.ListsCategory,
			getter:   &mock.ListHandler{},
			getDir: func(t *testing.T) path.Path {
				dir, err := path.Build(
					tenant,
					user,
					path.SharePointService,
					path.ListsCategory,
					false,
					dirRoot)
				require.NoError(t, err, clues.ToCore(err))

				return dir
			},
			getItem: func(t *testing.T, name string) data.Item {
				ow := kioser.NewJsonSerializationWriter()
				listing := spMock.ListDefault(name)
				listing.SetDisplayName(&name)

				err := ow.WriteObjectValue("", listing)
				require.NoError(t, err, clues.ToCore(err))

				byteArray, err := ow.GetSerializedContent()
				require.NoError(t, err, clues.ToCore(err))

				info := &details.SharePointInfo{
					List: &details.ListInfo{
						Name: name,
					},
				}

				data, err := data.NewPrefetchedItemWithInfo(
					io.NopCloser(bytes.NewReader(byteArray)),
					name,
					details.ItemInfo{SharePoint: info})
				require.NoError(t, err, clues.ToCore(err))

				return data
			},
		},
		{
			name:     "Pages",
			itemName: "MockPages",
			cat:      path.PagesCategory,
			scope:    sel.Pages(selectors.Any())[0],
			cat:      path.PagesCategory,
			getter:   nil,
			getDir: func(t *testing.T) path.Path {
				dir, err := path.Build(
					tenant,
					user,
					path.SharePointService,
					path.PagesCategory,
					false,
					dirRoot)
				require.NoError(t, err, clues.ToCore(err))

				return dir
			},
			getItem: func(t *testing.T, itemName string) data.Item {
				byteArray := spMock.Page(itemName)
				page, err := betaAPI.BytesToSitePageable(byteArray)
				require.NoError(t, err, clues.ToCore(err))

				data, err := data.NewPrefetchedItemWithInfo(
					io.NopCloser(bytes.NewReader(byteArray)),
					itemName,
					details.ItemInfo{SharePoint: betaAPI.PageInfo(page, int64(len(byteArray)))})
				require.NoError(t, err, clues.ToCore(err))

				return data
			},
		},
	}

	for _, test := range tables {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			col := NewPrefetchCollection(
				test.getter,
				test.getDir(t),
				suite.ac,
				test.scope,
				nil,
				control.DefaultOptions())
			col.stream[test.cat] = make(chan data.Item, collectionChannelBufferSize)
			col.stream[test.cat] <- test.getItem(t, test.itemName)

			readItems := []data.Item{}

			for item := range col.Items(ctx, fault.New(true)) {
				readItems = append(readItems, item)
			}

			require.Equal(t, len(readItems), 1)
			item := readItems[0]
			shareInfo, ok := item.(data.ItemInfo)
			require.True(t, ok)

			info, err := shareInfo.Info()
			require.NoError(t, err, clues.ToCore(err))

			assert.NotNil(t, info)
			assert.NotNil(t, info.SharePoint)

			if test.cat == path.ListsCategory {
				assert.Equal(t, test.itemName, info.SharePoint.List.Name)
			}
		})
	}
}

func (suite *SharePointCollectionSuite) TestLazyCollection_Items() {
	var (
		t             = suite.T()
		errs          = fault.New(true)
		start         = time.Now().Add(-time.Second)
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	fullPath, err := path.Build(
		"t", "pr", path.SharePointService, path.ListsCategory, false, "listid")
	require.NoError(t, err, clues.ToCore(err))

	tables := []struct {
		name            string
		items           map[string]time.Time
		expectItemCount int
		expectReads     []string
	}{
		{
			name: "no lists",
		},
		{
			name: "added lists",
			items: map[string]time.Time{
				"list1": start.Add(time.Minute),
				"list2": start.Add(2 * time.Minute),
				"list3": start.Add(3 * time.Minute),
			},
			expectItemCount: 3,
			expectReads: []string{
				"list1",
				"list2",
				"list3",
			},
		},
	}

	for _, test := range tables {
		suite.Run(test.name, func() {
			itemCount := 0

			ctx, flush := tester.NewContext(t)
			defer flush()

			getter := &mock.ListHandler{}
			defer getter.Check(t, test.expectReads)

			col := &lazyFetchCollection{
				stream:        make(chan data.Item),
				fullPath:      fullPath,
				items:         test.items,
				getter:        getter,
				statusUpdater: statusUpdater,
			}

			for item := range col.Items(ctx, errs) {
				itemCount++

				modTime, aok := test.items[item.ID()]
				require.True(t, aok, "item must have been added: %q", item.ID())
				assert.Implements(t, (*data.ItemModTime)(nil), item)
				assert.Equal(t, modTime, item.(data.ItemModTime).ModTime(), "item mod time")

				if slices.Contains(test.expectReads, item.ID()) {
					r := item.ToReader()

					_, err := io.ReadAll(r)
					assert.NoError(t, err, clues.ToCore(err))

					r.Close()

					assert.Implements(t, (*data.ItemInfo)(nil), item)
					info, err := item.(data.ItemInfo).Info()

					assert.NoError(t, err, clues.ToCore(err))
					assert.Equal(t, modTime, info.Modified(), "ItemInfo mod time")
				}
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

func (suite *SharePointCollectionSuite) TestLazyItem() {
	var (
		t   = suite.T()
		now = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	lh := mock.ListHandler{}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter{
			itemID:  "itemID",
			getter:  &lh,
			modTime: now,
		},
		"itemID",
		now,
		count.New(),
		fault.New(true))

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
	require.NoError(t, err, "reading item data: %v", clues.ToCore(err))
	assert.NotEmpty(t, readData, "read item data")

	info, err := li.Info()
	require.NoError(t, err, "getting item info: %v", clues.ToCore(err))
	assert.Equal(t, now, info.Modified())
}

func (suite *SharePointCollectionSuite) TestLazyItem_ReturnsEmptyReaderOnDeletedInFlight() {
	var (
		t   = suite.T()
		now = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	lh := mock.ListHandler{
		Err: graph.ErrDeletedInFlight,
	}

	li := data.NewLazyItemWithInfo(
		ctx,
		&lazyItemGetter{
			itemID:  "itemID",
			getter:  &lh,
			modTime: now,
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
	assert.ErrorIs(t, err, graph.ErrDeletedInFlight, "item should be marked deleted in flight")
	assert.Nil(t, r)
}
