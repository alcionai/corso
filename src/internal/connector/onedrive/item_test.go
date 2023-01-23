package onedrive

import (
	"bytes"
	"context"
	"io"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type ItemIntegrationSuite struct {
	suite.Suite
	// site        string
	// siteDriveID string
	user        string
	userDriveID string
	client      *msgraphsdk.GraphServiceClient
	adapter     *msgraphsdk.GraphRequestAdapter
}

func (suite *ItemIntegrationSuite) Client() *msgraphsdk.GraphServiceClient {
	return suite.client
}

func (suite *ItemIntegrationSuite) Adapter() *msgraphsdk.GraphRequestAdapter {
	return suite.adapter
}

func TestItemIntegrationSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorOneDriveTests)

	suite.Run(t, new(ItemIntegrationSuite))
}

func (suite *ItemIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	adapter, err := graph.CreateAdapter(m365.AzureTenantID, m365.AzureClientID, m365.AzureClientSecret)
	require.NoError(t, err)

	suite.client = msgraphsdk.NewGraphServiceClient(adapter)
	suite.adapter = adapter

	// TODO: fulfill file preconditions required for testing (expected files w/in drive
	// and guarateed drive read-write access)
	// suite.site = tester.M365SiteID(t)
	// spDrives, err := drives(ctx, suite, suite.site, SharePointSource)
	// require.NoError(t, err)
	// // Test Requirement 1: Need a drive
	// require.Greaterf(t, len(spDrives), 0, "site %s does not have a drive", suite.site)

	// // Pick the first drive
	// suite.siteDriveID = *spDrives[0].GetId()

	suite.user = tester.SecondaryM365UserID(t)

	odDrives, err := drives(ctx, suite, suite.user, OneDriveSource)
	require.NoError(t, err)
	// Test Requirement 1: Need a drive
	require.Greaterf(t, len(odDrives), 0, "user %s does not have a drive", suite.user)

	// Pick the first drive
	suite.userDriveID = *odDrives[0].GetId()
}

// TestItemReader is an integration test that makes a few assumptions
// about the test environment
// 1) It assumes the test user has a drive
// 2) It assumes the drive has a file it can use to test `driveItemReader`
// The test checks these in below
func (suite *ItemIntegrationSuite) TestItemReader_oneDrive() {
	ctx, flush := tester.NewContext()
	defer flush()

	var driveItem models.DriveItemable
	// This item collector tries to find "a" drive item that is a file to test the reader function
	itemCollector := func(
		ctx context.Context,
		driveID, driveName string,
		items []models.DriveItemable,
		paths map[string]string,
	) error {
		for _, item := range items {
			if item.GetFile() != nil {
				driveItem = item
				break
			}
		}

		return nil
	}
	_, _, err := collectItems(ctx, suite, suite.userDriveID, "General", itemCollector)
	require.NoError(suite.T(), err)

	// Test Requirement 2: Need a file
	require.NotEmpty(
		suite.T(),
		driveItem,
		"no file item found for user %s drive %s",
		suite.user,
		suite.userDriveID,
	)

	// Read data for the file

	itemInfo, itemData, err := oneDriveItemReader(ctx, driveItem)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), itemInfo.OneDrive)
	require.NotEmpty(suite.T(), itemInfo.OneDrive.ItemName)

	size, err := io.Copy(io.Discard, itemData)
	require.NoError(suite.T(), err)
	require.NotZero(suite.T(), size)
	require.Equal(suite.T(), size, itemInfo.OneDrive.Size)
	suite.T().Logf("Read %d bytes from file %s.", size, itemInfo.OneDrive.ItemName)
}

// TestItemWriter is an integration test for uploading data to OneDrive
// It creates a new `testfolder_<timestamp` folder with a new
// testitem_<timestamp> item and writes data to it
func (suite *ItemIntegrationSuite) TestItemWriter() {
	table := []struct {
		name    string
		driveID string
	}{
		{
			name:    "",
			driveID: suite.userDriveID,
		},
		// {
		// 	name:   "sharePoint",
		// 	driveID: suite.siteDriveID,
		// },
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			root, err := suite.Client().DrivesById(test.driveID).Root().Get(ctx, nil)
			require.NoError(suite.T(), err)

			// Test Requirement 2: "Test Folder" should exist
			folder, err := getFolder(ctx, suite, test.driveID, *root.GetId(), "Test Folder")
			require.NoError(suite.T(), err)

			newFolderName := "testfolder_" + common.FormatNow(common.SimpleTimeTesting)
			suite.T().Logf("Test will create folder %s", newFolderName)

			newFolder, err := createItem(ctx, suite, test.driveID, *folder.GetId(), newItem(newFolderName, true))
			require.NoError(suite.T(), err)

			require.NotNil(suite.T(), newFolder.GetId())

			newItemName := "testItem_" + common.FormatNow(common.SimpleTimeTesting)
			suite.T().Logf("Test will create item %s", newItemName)

			newItem, err := createItem(ctx, suite, test.driveID, *newFolder.GetId(), newItem(newItemName, false))
			require.NoError(suite.T(), err)

			require.NotNil(suite.T(), newItem.GetId())

			// HACK: Leveraging this to test getFolder behavior for a file. `getFolder()` on the
			// newly created item should fail because it's a file not a folder
			_, err = getFolder(ctx, suite, test.driveID, *newFolder.GetId(), newItemName)
			require.ErrorIs(suite.T(), err, errFolderNotFound)

			// Initialize a 100KB mockDataProvider
			td, writeSize := mockDataReader(int64(100 * 1024))

			w, err := driveItemWriter(ctx, suite, test.driveID, *newItem.GetId(), writeSize)
			require.NoError(suite.T(), err)

			// Using a 32 KB buffer for the copy allows us to validate the
			// multi-part upload. `io.CopyBuffer` will only write 32 KB at
			// a time
			copyBuffer := make([]byte, 32*1024)

			size, err := io.CopyBuffer(w, td, copyBuffer)
			require.NoError(suite.T(), err)

			require.Equal(suite.T(), writeSize, size)
		})
	}
}

func mockDataReader(size int64) (io.Reader, int64) {
	data := bytes.Repeat([]byte("D"), int(size))
	return bytes.NewReader(data), size
}

func (suite *ItemIntegrationSuite) TestDriveGetFolder() {
	table := []struct {
		name    string
		driveID string
	}{
		{
			name:    "oneDrive",
			driveID: suite.userDriveID,
		},
		// {
		// 	name:   "sharePoint",
		// 	driveID: suite.siteDriveID,
		// },
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			root, err := suite.Client().DrivesById(test.driveID).Root().Get(ctx, nil)
			require.NoError(suite.T(), err)

			// Lookup a folder that doesn't exist
			_, err = getFolder(ctx, suite, test.driveID, *root.GetId(), "FolderDoesNotExist")
			require.ErrorIs(suite.T(), err, errFolderNotFound)

			// Lookup a folder that does exist
			_, err = getFolder(ctx, suite, test.driveID, *root.GetId(), "")
			require.NoError(suite.T(), err)
		})
	}
}
