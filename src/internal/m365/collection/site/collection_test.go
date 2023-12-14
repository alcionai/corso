package site

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
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
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
func (suite *SharePointCollectionSuite) TestCollection_Items() {
	var (
		tenant  = "some"
		user    = "user"
		dirRoot = "directory"
	)

	sel := selectors.NewSharePointBackup([]string{"site"})

	tables := []struct {
		name, itemName string
		scope          selectors.SharePointScope
		getter         getItemByIDer
		getDir         func(t *testing.T) path.Path
		getItem        func(t *testing.T, itemName string) data.Item
	}{
		{
			name:     "List",
			itemName: "MockListing",
			scope:    sel.Lists(selectors.Any())[0],
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

				data, err := data.NewPrefetchedItemWithInfo(
					io.NopCloser(bytes.NewReader(byteArray)),
					name,
					details.ItemInfo{SharePoint: ListToSPInfo(listing)})
				require.NoError(t, err, clues.ToCore(err))

				return data
			},
		},
		{
			name:     "Pages",
			itemName: "MockPages",
			scope:    sel.Pages(selectors.Any())[0],
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
				page, err := betaAPI.CreatePageFromBytes(byteArray)
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

			col := NewCollection(
				test.getter,
				test.getDir(t),
				suite.ac,
				test.scope,
				nil,
				control.DefaultOptions())
			col.stream <- test.getItem(t, test.itemName)

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
			assert.Equal(t, test.itemName, info.SharePoint.ItemName)
		})
	}
}

func (suite *SharePointCollectionSuite) TestCollection_streamItems() {
	var (
		t             = suite.T()
		statusUpdater = func(*support.ControllerOperationStatus) {}
		tenant        = "some"
		resource      = "siteid"
		list          = "list"
	)

	table := []struct {
		name     string
		category path.CategoryType
		items    []string
		getDir   func(t *testing.T) path.Path
	}{
		{
			name:     "no items",
			items:    []string{},
			category: path.ListsCategory,
			getDir: func(t *testing.T) path.Path {
				dir, err := path.Build(
					tenant,
					resource,
					path.SharePointService,
					path.ListsCategory,
					false,
					list)
				require.NoError(t, err, clues.ToCore(err))

				return dir
			},
		},
		{
			name:     "with items",
			items:    []string{"list1", "list2", "list3"},
			category: path.ListsCategory,
			getDir: func(t *testing.T) path.Path {
				dir, err := path.Build(
					tenant,
					resource,
					path.SharePointService,
					path.ListsCategory,
					false,
					list)
				require.NoError(t, err, clues.ToCore(err))

				return dir
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t.Log("running test", test)

			var (
				errs      = fault.New(true)
				itemCount int
			)

			ctx, flush := tester.NewContext(t)
			defer flush()

			col := &Collection{
				fullPath:      test.getDir(t),
				category:      test.category,
				items:         test.items,
				getter:        &mock.ListHandler{},
				stream:        make(chan data.Item),
				statusUpdater: statusUpdater,
			}

			itemMap := func(js []string) map[string]struct{} {
				m := make(map[string]struct{})
				for _, j := range js {
					m[j] = struct{}{}
				}
				return m
			}(test.items)

			go col.streamItems(ctx, errs)

			for item := range col.stream {
				itemCount++
				_, ok := itemMap[item.ID()]
				assert.True(t, ok, "should fetch item")
			}

			assert.NoError(t, errs.Failure())
			assert.Equal(t, len(test.items), itemCount, "should see all expected items")
		})
	}
}

// TestRestoreListCollection verifies Graph Restore API for the List Collection
func (suite *SharePointCollectionSuite) TestListCollection_Restore() {
	t := suite.T()
	// https://github.com/microsoftgraph/msgraph-sdk-go/issues/490
	t.Skip("disabled until upstream issue with list restore is fixed.")

	ctx, flush := tester.NewContext(t)
	defer flush()

	service := createTestService(t, suite.creds)
	listing := spMock.ListDefault("Mock List")
	testName := "MockListing"
	listing.SetDisplayName(&testName)
	byteArray, err := service.Serialize(listing)
	require.NoError(t, err, clues.ToCore(err))

	listData, err := data.NewPrefetchedItemWithInfo(
		io.NopCloser(bytes.NewReader(byteArray)),
		testName,
		details.ItemInfo{SharePoint: ListToSPInfo(listing)})
	require.NoError(t, err, clues.ToCore(err))

	destName := testdata.DefaultRestoreConfig("").Location

	deets, err := restoreListItem(ctx, service, listData, suite.siteID, destName)
	assert.NoError(t, err, clues.ToCore(err))
	t.Logf("List created: %s\n", deets.SharePoint.ItemName)

	// Clean-Up
	var (
		builder  = service.Client().Sites().BySiteId(suite.siteID).Lists()
		isFound  bool
		deleteID string
	)

	for {
		resp, err := builder.Get(ctx, nil)
		assert.NoError(t, err, "getting site lists", clues.ToCore(err))

		for _, temp := range resp.GetValue() {
			if ptr.Val(temp.GetDisplayName()) == deets.SharePoint.ItemName {
				isFound = true
				deleteID = ptr.Val(temp.GetId())

				break
			}
		}
		// Get Next Link
		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = sites.NewItemListsRequestBuilder(link, service.Adapter())
	}

	if isFound {
		err := DeleteList(ctx, service, suite.siteID, deleteID)
		assert.NoError(t, err, clues.ToCore(err))
	}
}
