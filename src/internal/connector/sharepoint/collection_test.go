package sharepoint

import (
	"bytes"
	"io"
	"testing"

	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint/api"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type SharePointCollectionSuite struct {
	suite.Suite
	siteID string
	creds  account.M365Config
}

func (suite *SharePointCollectionSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	suite.siteID = tester.M365SiteID(t)
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
}

func TestSharePointCollectionSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorSharePointTests)

	suite.Run(t, new(SharePointCollectionSuite))
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
	require.NoError(t, err)

	assert.Equal(t, name, sc.id)
	assert.Equal(t, readData, m)
}

// TestListCollection tests basic functionality to create
// SharePoint collection and to use the data stream channel.
func (suite *SharePointCollectionSuite) TestCollection_Items() {
	t := suite.T()
	tenant := "some"
	user := "user"
	dirRoot := "directory"
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
				dir, err := path.Builder{}.Append(dirRoot).
					ToDataLayerSharePointPath(
						tenant,
						user,
						path.ListsCategory,
						false)
				require.NoError(t, err)

				return dir
			},
			getItem: func(t *testing.T, name string) *Item {
				ow := kioser.NewJsonSerializationWriter()
				listing := mockconnector.GetMockListDefault(name)
				listing.SetDisplayName(&name)

				err := ow.WriteObjectValue("", listing)
				require.NoError(t, err)

				byteArray, err := ow.GetSerializedContent()
				require.NoError(t, err)

				data := &Item{
					id:   name,
					data: io.NopCloser(bytes.NewReader(byteArray)),
					info: sharePointListInfo(listing, int64(len(byteArray))),
				}

				return data
			},
		},
		{
			name:     "Pages",
			itemName: "MockPages",
			category: Pages,
			getDir: func(t *testing.T) path.Path {
				dir, err := path.Builder{}.Append(dirRoot).
					ToDataLayerSharePointPath(
						tenant,
						user,
						path.PagesCategory,
						false)
				require.NoError(t, err)

				return dir
			},
			getItem: func(t *testing.T, itemName string) *Item {
				byteArray := mockconnector.GetMockPage(itemName)
				page, err := support.CreatePageFromBytes(byteArray)
				require.NoError(t, err)

				data := &Item{
					id:   itemName,
					data: io.NopCloser(bytes.NewReader(byteArray)),
					info: api.PageInfo(page, int64(len(byteArray))),
				}

				return data
			},
		},
	}

	for _, test := range tables {
		t.Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			col := NewCollection(test.getDir(t), nil, test.category, nil, control.Defaults())
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
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	service := createTestService(t, suite.creds)
	listing := mockconnector.GetMockListDefault("Mock List")
	testName := "MockListing"
	listing.SetDisplayName(&testName)
	byteArray, err := service.Serialize(listing)
	require.NoError(t, err)

	listData := &Item{
		id:   testName,
		data: io.NopCloser(bytes.NewReader(byteArray)),
		info: sharePointListInfo(listing, int64(len(byteArray))),
	}

	destName := "Corso_Restore_" + common.FormatNow(common.SimpleTimeTesting)

	deets, err := restoreListItem(ctx, service, listData, suite.siteID, destName)
	assert.NoError(t, err)
	t.Logf("List created: %s\n", deets.SharePoint.ItemName)

	// Clean-Up
	var (
		builder  = service.Client().SitesById(suite.siteID).Lists()
		isFound  bool
		deleteID string
	)

	for {
		resp, err := builder.Get(ctx, nil)
		assert.NoError(t, err, "experienced query error during clean up. Details:  "+support.ConnectorStackErrorTrace(err))

		for _, temp := range resp.GetValue() {
			if *temp.GetDisplayName() == deets.SharePoint.ItemName {
				isFound = true
				deleteID = *temp.GetId()

				break
			}
		}
		// Get Next Link
		link := resp.GetOdataNextLink()
		if link == nil {
			break
		}

		builder = sites.NewItemListsRequestBuilder(*link, service.Adapter())
	}

	if isFound {
		err := DeleteList(ctx, service, suite.siteID, deleteID)
		assert.NoError(t, err)
	}
}

// TestRestoreLocation temporary test for greater restore operation
// TODO delete after full functionality tested in GraphConnector
func (suite *SharePointCollectionSuite) TestRestoreLocation() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	service := createTestService(t, suite.creds)
	rootFolder := "General_" + common.FormatNow(common.SimpleTimeTesting)
	folderID, err := createRestoreFolders(ctx, service, suite.siteID, []string{rootFolder})
	assert.NoError(t, err)
	t.Log("FolderID: " + folderID)

	_, err = createRestoreFolders(ctx, service, suite.siteID, []string{rootFolder, "Tsao"})
	assert.NoError(t, err)

	// CleanUp
	siteDrive, err := service.Client().SitesById(suite.siteID).Drive().Get(ctx, nil)
	require.NoError(t, err)

	driveID := *siteDrive.GetId()
	err = onedrive.DeleteItem(ctx, service, driveID, folderID)
	assert.NoError(t, err)
}
