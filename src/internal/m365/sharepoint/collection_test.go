package sharepoint

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
	betaAPI "github.com/alcionai/corso/src/internal/m365/sharepoint/api"
	spMock "github.com/alcionai/corso/src/internal/m365/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

	ac, err := api.NewClient(m365)
	require.NoError(t, err, clues.ToCore(err))

	suite.ac = ac
}

func TestSharePointCollectionSuite(t *testing.T) {
	suite.Run(t, &SharePointCollectionSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *SharePointCollectionSuite) TestCollection_Item_Read() {
	t := suite.T()
	m := []byte("test message")
	name := "aFile"
	sc := &Item{
		id:   name,
		data: io.NopCloser(bytes.NewReader(m)),
	}
	readData, err := io.ReadAll(sc.ToReader())
	require.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, name, sc.id)
	assert.Equal(t, readData, m)
}

// TestListCollection tests basic functionality to create
// SharePoint collection and to use the data stream channel.
func (suite *SharePointCollectionSuite) TestCollection_Items() {
	var (
		tenant  = "some"
		user    = "user"
		dirRoot = "directory"
	)

	tables := []struct {
		name, itemName string
		category       DataCategory
		getDir         func(t *testing.T) path.Path
		getItem        func(t *testing.T, itemName string) *Item
	}{
		{
			name:     "List",
			itemName: "MockListing",
			category: List,
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
			getItem: func(t *testing.T, name string) *Item {
				ow := kioser.NewJsonSerializationWriter()
				listing := spMock.ListDefault(name)
				listing.SetDisplayName(&name)

				err := ow.WriteObjectValue("", listing)
				require.NoError(t, err, clues.ToCore(err))

				byteArray, err := ow.GetSerializedContent()
				require.NoError(t, err, clues.ToCore(err))

				data := &Item{
					id:   name,
					data: io.NopCloser(bytes.NewReader(byteArray)),
					info: listToSPInfo(listing, int64(len(byteArray))),
				}

				return data
			},
		},
		{
			name:     "Pages",
			itemName: "MockPages",
			category: Pages,
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
			getItem: func(t *testing.T, itemName string) *Item {
				byteArray := spMock.Page(itemName)
				page, err := betaAPI.CreatePageFromBytes(byteArray)
				require.NoError(t, err, clues.ToCore(err))

				data := &Item{
					id:   itemName,
					data: io.NopCloser(bytes.NewReader(byteArray)),
					info: betaAPI.PageInfo(page, int64(len(byteArray))),
				}

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
				test.getDir(t),
				suite.ac,
				test.category,
				nil,
				control.DefaultOptions())
			col.data <- test.getItem(t, test.itemName)

			readItems := []data.Stream{}

			for item := range col.Items(ctx, fault.New(true)) {
				readItems = append(readItems, item)
			}

			require.Equal(t, len(readItems), 1)
			item := readItems[0]
			shareInfo, ok := item.(data.StreamInfo)
			require.True(t, ok)
			require.NotNil(t, shareInfo.Info())
			require.NotNil(t, shareInfo.Info().SharePoint)
			assert.Equal(t, test.itemName, shareInfo.Info().SharePoint.ItemName)
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

	listData := &Item{
		id:   testName,
		data: io.NopCloser(bytes.NewReader(byteArray)),
		info: listToSPInfo(listing, int64(len(byteArray))),
	}

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
