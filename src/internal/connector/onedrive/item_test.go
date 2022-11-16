package onedrive

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type ItemIntegrationSuite struct {
	suite.Suite
	user    string
	driveID string
	client  *msgraphsdk.GraphServiceClient
	adapter *msgraphsdk.GraphRequestAdapter
}

func (suite *ItemIntegrationSuite) Client() *msgraphsdk.GraphServiceClient {
	return suite.client
}

func (suite *ItemIntegrationSuite) Adapter() *msgraphsdk.GraphRequestAdapter {
	return suite.adapter
}

func (suite *ItemIntegrationSuite) ErrPolicy() bool {
	return false
}

func TestItemIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorOneDriveTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ItemIntegrationSuite))
}

func (suite *ItemIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	a := tester.NewM365Account(suite.T())

	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	adapter, err := graph.CreateAdapter(m365.AzureTenantID, m365.AzureClientID, m365.AzureClientSecret)
	require.NoError(suite.T(), err)
	suite.client = msgraphsdk.NewGraphServiceClient(adapter)
	suite.adapter = adapter

	suite.user = tester.SecondaryM365UserID(suite.T())

	drives, err := drives(ctx, suite, suite.user, OneDriveSource)
	require.NoError(suite.T(), err)
	// Test Requirement 1: Need a drive
	require.Greaterf(suite.T(), len(drives), 0, "user %s does not have a drive", suite.user)

	// Pick the first drive
	suite.driveID = *drives[0].GetId()
}

// TestItemReader is an integration test that makes a few assumptions
// about the test environment
// 1) It assumes the test user has a drive
// 2) It assumes the drive has a file it can use to test `driveItemReader`
// The test checks these in below
func (suite *ItemIntegrationSuite) TestItemReader() {
	ctx, flush := tester.NewContext()
	defer flush()

	var driveItemID string
	// This item collector tries to find "a" drive item that is a file to test the reader function
	itemCollector := func(ctx context.Context, driveID string, items []models.DriveItemable) error {
		for _, item := range items {
			if item.GetFile() != nil {
				driveItemID = *item.GetId()
				break
			}
		}

		return nil
	}
	err := collectItems(ctx, suite, suite.driveID, itemCollector)
	require.NoError(suite.T(), err)

	// Test Requirement 2: Need a file
	require.NotEmpty(
		suite.T(),
		driveItemID,
		"no file item found for user %s drive %s",
		suite.user,
		suite.driveID,
	)

	// Read data for the file

	itemInfo, itemData, err := driveItemReader(ctx, suite, suite.driveID, driveItemID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), itemInfo)
	require.NotEmpty(suite.T(), itemInfo.ItemName)

	size, err := io.Copy(io.Discard, itemData)
	require.NoError(suite.T(), err)
	require.NotZero(suite.T(), size)
	require.Equal(suite.T(), size, itemInfo.Size)
	suite.T().Logf("Read %d bytes from file %s.", size, itemInfo.ItemName)
}

// TestItemWriter is an integration test for uploading data to OneDrive
// It creates a new `testfolder_<timestamp` folder with a new
// testitem_<timestamp> item and writes data to it
func (suite *ItemIntegrationSuite) TestItemWriter() {
	ctx, flush := tester.NewContext()
	defer flush()

	root, err := suite.Client().DrivesById(suite.driveID).Root().Get(ctx, nil)
	require.NoError(suite.T(), err)

	// Test Requirement 2: "Test Folder" should exist
	folder, err := getFolder(ctx, suite, suite.driveID, *root.GetId(), "Test Folder")
	require.NoError(suite.T(), err)

	newFolderName := "testfolder_" + time.Now().Format("2006-01-02T15-04-05")
	suite.T().Logf("Test will create folder %s", newFolderName)

	newFolder, err := createItem(ctx, suite, suite.driveID, *folder.GetId(), newItem(newFolderName, true))
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), newFolder.GetId())

	newItemName := "testItem_" + time.Now().Format("2006-01-02T15-04-05")
	suite.T().Logf("Test will create item %s", newItemName)

	newItem, err := createItem(ctx, suite, suite.driveID, *newFolder.GetId(), newItem(newItemName, false))
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), newItem.GetId())

	// HACK: Leveraging this to test getFolder behavior for a file. `getFolder()` on the
	// newly created item should fail because it's a file not a folder
	_, err = getFolder(ctx, suite, suite.driveID, *newFolder.GetId(), newItemName)
	require.ErrorIs(suite.T(), err, errFolderNotFound)

	// Initialize a 100KB mockDataProvider
	td, writeSize := mockDataReader(int64(100 * 1024))

	w, err := driveItemWriter(ctx, suite, suite.driveID, *newItem.GetId(), writeSize)
	require.NoError(suite.T(), err)

	// Using a 32 KB buffer for the copy allows us to validate the
	// multi-part upload. `io.CopyBuffer` will only write 32 KB at
	// a time
	copyBuffer := make([]byte, 32*1024)

	size, err := io.CopyBuffer(w, td, copyBuffer)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), writeSize, size)
}

func mockDataReader(size int64) (io.Reader, int64) {
	data := bytes.Repeat([]byte("D"), int(size))
	return bytes.NewReader(data), size
}

func (suite *ItemIntegrationSuite) TestDriveGetFolder() {
	ctx, flush := tester.NewContext()
	defer flush()

	root, err := suite.Client().DrivesById(suite.driveID).Root().Get(ctx, nil)
	require.NoError(suite.T(), err)

	// Lookup a folder that doesn't exist
	_, err = getFolder(ctx, suite, suite.driveID, *root.GetId(), "FolderDoesNotExist")
	require.ErrorIs(suite.T(), err, errFolderNotFound)

	// Lookup a folder that does exist
	_, err = getFolder(ctx, suite, suite.driveID, *root.GetId(), "")
	require.NoError(suite.T(), err)
}
