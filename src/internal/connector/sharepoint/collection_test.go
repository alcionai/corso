package sharepoint

import (
	"bytes"
	"io"
	"testing"

	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type SharePointCollectionSuite struct {
	suite.Suite
}

func TestSharePointCollectionSuite(t *testing.T) {
	suite.Run(t, new(SharePointCollectionSuite))
}

func (suite *SharePointCollectionSuite) TestSharePointDataReader_Valid() {
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

// TestSharePointListCollection tests basic functionality to create
// SharePoint collection and to use the data stream channel.
func (suite *SharePointCollectionSuite) TestSharePointListCollection() {
	t := suite.T()

	ow := kw.NewJsonSerializationWriter()
	listing := mockconnector.GetMockList("Mock List")
	testName := "MockListing"
	listing.SetDisplayName(&testName)

	err := ow.WriteObjectValue("", listing)
	require.NoError(t, err)

	byteArray, err := ow.GetSerializedContent()
	require.NoError(t, err)

	dir, err := path.Builder{}.Append("directory").
		ToDataLayerSharePointPath(
			"some",
			"user",
			path.ListsCategory,
			false)
	require.NoError(t, err)

	col := NewCollection(dir, nil, nil)
	col.data <- &Item{
		id:   testName,
		data: io.NopCloser(bytes.NewReader(byteArray)),
		info: sharePointListInfo(listing, int64(len(byteArray))),
	}

	readItems := []data.Stream{}

	for item := range col.Items() {
		readItems = append(readItems, item)
	}

	require.Equal(t, len(readItems), 1)
	item := readItems[0]
	shareInfo, ok := item.(data.StreamInfo)
	require.True(t, ok)
	require.NotNil(t, shareInfo.Info())
	require.NotNil(t, shareInfo.Info().SharePoint)
	assert.Equal(t, testName, shareInfo.Info().SharePoint.ItemName)
}

func (suite *SharePointCollectionSuite) TestRestoreListCollection() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t            = suite.T()
		tenant       = "Mocked Tenant"
		a            = tester.NewM365Account(t)
		siteID       = tester.M365SiteID(t)
		name         = "MockRestoreList"
		mockList     = mockconnector.GetMockList(name)
		account, err = a.M365Config()
	)

	require.NoError(t, err)

	dir, err := path.Builder{}.Append(name).
		ToDataLayerSharePointPath(
			tenant,
			siteID,
			path.ListsCategory,
			false)
	collection := NewCollection(dir, nil, nil)

	service, err := createTestService(account)
	require.NoError(t, err)

	driveID := *siteDrive.GetId()
	err = onedrive.DeleteItem(ctx, service, driveID, folderID)
	assert.NoError(t, err)

	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	a := tester.NewM365Account(t)
	account, err := a.M365Config()
	require.NoError(t, err)

	require.NoError(t, err)
	siteID := tester.M365SiteID(t)

	ow := kw.NewJsonSerializationWriter()
	listing := mockconnector.GetMockList("Mock List")
	testName := "MockListing"
	listing.SetDisplayName(&testName)

	err = ow.WriteObjectValue("", listing)
	require.NoError(t, err)

	byteArray, err := ow.GetSerializedContent()
	require.NoError(t, err)

	require.NoError(t, err)

	listData := &Item{
		id:   testName,
		data: io.NopCloser(bytes.NewReader(byteArray)),
		info: sharePointListInfo(listing, int64(len(byteArray))),
	}

	destName := "Corso_Restore_" + common.FormatNow(common.SimpleTimeTesting)

	_, canceled := RestoreCollection(ctx, service, collection, destName, deets, updater)
	assert.False(t, canceled)

}

// TestRestoreLocation temporary test for greater restore operation
// TODO delete after full functionality tested in GraphConnector
func (suite *SharePointCollectionSuite) TestRestoreLocation() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	a := tester.NewM365Account(t)
	account, err := a.M365Config()
	require.NoError(t, err)

	service, err := createTestService(account)
	require.NoError(t, err)

	rootFolder := "General_" + common.FormatNow(common.SimpleTimeTesting)
	siteID := tester.M365SiteID(t)

	folderID, err := createRestoreFolders(ctx, service, siteID, []string{rootFolder})
	assert.NoError(t, err)
	t.Log("FolderID: " + folderID)

	_, err = createRestoreFolders(ctx, service, siteID, []string{rootFolder, "Tsao"})
	assert.NoError(t, err)

	// CleanUp
	siteDrive, err := service.Client().SitesById(siteID).Drive().Get(ctx, nil)
	require.NoError(t, err)

	driveID := *siteDrive.GetId()
	err = onedrive.DeleteItem(ctx, service, driveID, folderID)
	assert.NoError(t, err)
}
