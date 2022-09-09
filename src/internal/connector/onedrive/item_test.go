package onedrive

import (
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
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(ItemIntegrationSuite))
}

func (suite *ItemIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	a := tester.NewM365Account(suite.T())

	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	adapter, err := graph.CreateAdapter(m365.TenantID, m365.ClientID, m365.ClientSecret)
	require.NoError(suite.T(), err)
	suite.client = msgraphsdk.NewGraphServiceClient(adapter)
	suite.adapter = adapter
}

// TestItemReader is an integration test that makes a few assumptions
// about the test environment
// 1) It assumes the test user has a drive
// 2) It assumes the drive has a file it can use to test `driveItemReader`
// The test checks these in below
func (suite *ItemIntegrationSuite) TestItemReader() {
	ctx := context.TODO()
	user := tester.M365UserID(suite.T())

	drives, err := drives(ctx, suite, user)
	require.NoError(suite.T(), err)
	// Test Requirement 1: Need a drive
	require.Greaterf(suite.T(), len(drives), 0, "user %s does not have a drive", user)

	// Pick the first drive
	driveID := *drives[0].GetId()

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
	err = collectItems(ctx, suite, driveID, itemCollector)
	require.NoError(suite.T(), err)

	// Test Requirement 2: Need a file
	require.NotEmpty(
		suite.T(),
		driveItemID,
		"no file item found for user %s drive %s",
		user,
		driveID,
	)

	// Read data for the file

	name, itemData, err := driveItemReader(ctx, suite, driveID, driveItemID)
	require.NoError(suite.T(), err)
	require.NotEmpty(suite.T(), name)
	size, err := io.Copy(io.Discard, itemData)
	require.NoError(suite.T(), err)
	require.NotZero(suite.T(), size)
	suite.T().Logf("Read %d bytes from file %s.", size, name)
}

// TestItemWriter is an integration test for uploading data to OneDrive
// It creates a new `testfolder_<timestamp` folder with a new
// testitem_<timestamp> item and writes data to it
func (suite *ItemIntegrationSuite) TestItemWriter() {
	ctx := context.TODO()
	user := tester.M365UserID(suite.T())

	drives, err := drives(ctx, suite, user)
	require.NoError(suite.T(), err)
	// Test Requirement 1: Need a drive
	require.Greaterf(suite.T(), len(drives), 0, "user %s does not have a drive", user)

	// Pick the first drive
	driveID := *drives[0].GetId()

	root, err := suite.Client().DrivesById(driveID).Root().Get()
	require.NoError(suite.T(), err)

	// Test Requirement 2: "Test Folder" should exist
	folder, err := getFolder(ctx, suite, driveID, *root.GetId(), "Test Folder")
	require.NoError(suite.T(), err)

	newFolderName := "testfolder_" + time.Now().Format("2006-01-02T15-04-05")
	suite.T().Logf("Test will create folder %s", newFolderName)

	newFolder, err := createItem(ctx, suite, driveID, *folder.GetId(), newItem(newFolderName, true))
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), newFolder.GetId())

	newItemName := "testItem_" + time.Now().Format("2006-01-02T15-04-05")
	suite.T().Logf("Test will create item %s", newItemName)

	newItem, err := createItem(ctx, suite, driveID, *newFolder.GetId(), newItem(newItemName, false))
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), newItem.GetId())

	// Initialize a 100KB mockDataProvider
	td := &mockDataProvider{size: 100 * 1024}

	writeSize := td.size

	w, err := driveItemWriter(ctx, suite, driveID, *newItem.GetId(), writeSize)
	require.NoError(suite.T(), err)

	// Using a 32 KB buffer for the copy allows us to validate the
	// multi-part upload. `io.CopyBuffer` will only write 32 KB at
	// a time
	copyBuffer := make([]byte, 32*1024)

	size, err := io.CopyBuffer(w, td, copyBuffer)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), writeSize, size)
}

// mockDataProvider implements an `io.Reader` that can be used to
// read the specified amount of mock data
type mockDataProvider struct {
	size              int64
	currentReadOffset int64
}

func (td *mockDataProvider) Read(p []byte) (n int, err error) {
	// If we've already read all the mock data, return EOF
	if td.currentReadOffset == td.size {
		return 0, io.EOF
	}

	// Check how much data is left to read
	toRead := td.size - td.currentReadOffset

	// If the remaining data doesn't fit in the buffer,
	// only read len(buffer)
	if toRead > int64(len(p)) {
		toRead = int64(len(p))
	}

	// Fill the buffer with mock data.
	// Currently we just write the character `D`
	for i := int64(0); i < toRead; i++ {
		p[i] = byte('D')
	}

	// Track how much data was read
	td.currentReadOffset += int64(toRead)

	return int(toRead), nil
}
