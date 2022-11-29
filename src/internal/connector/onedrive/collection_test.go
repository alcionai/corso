package onedrive

import (
	"bytes"
	"context"
	"errors"
	"io"
	"sync"
	"testing"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type OneDriveCollectionSuite struct {
	suite.Suite
}

// Allows `*OneDriveCollectionSuite` to be used as a graph.Service
// TODO: Implement these methods

func (suite *OneDriveCollectionSuite) Client() *msgraphsdk.GraphServiceClient {
	return nil
}

func (suite *OneDriveCollectionSuite) Adapter() *msgraphsdk.GraphRequestAdapter {
	return nil
}

func (suite *OneDriveCollectionSuite) ErrPolicy() bool {
	return false
}

func TestOneDriveCollectionSuite(t *testing.T) {
	suite.Run(t, new(OneDriveCollectionSuite))
}

// Returns a status update function that signals the specified WaitGroup when it is done
func (suite *OneDriveCollectionSuite) testStatusUpdater(
	wg *sync.WaitGroup,
	statusToUpdate *support.ConnectorOperationStatus,
) support.StatusUpdater {
	return func(s *support.ConnectorOperationStatus) {
		suite.T().Logf("Update status %v, count %d, success %d", s, s.ObjectCount, s.Successful)
		*statusToUpdate = *s

		wg.Done()
	}
}

func (suite *OneDriveCollectionSuite) TestOneDriveCollection() {
	t := suite.T()
	wg := sync.WaitGroup{}
	collStatus := support.ConnectorOperationStatus{}
	t1 := time.Now()

	folderPath, err := GetCanonicalPath("drive/driveID1/root:/dir1/dir2/dir3", "a-tenant", "a-user", OneDriveSource)
	require.NoError(t, err)
	driveFolderPath, err := getDriveFolderPath(folderPath)
	require.NoError(t, err)

	coll := NewCollection(folderPath, "fakeDriveID", suite, suite.testStatusUpdater(&wg, &collStatus))
	require.NotNil(t, coll)
	assert.Equal(t, folderPath, coll.FullPath())

	testItemID := "fakeItemID"
	testItemName := "itemName"
	testItemData := []byte("testdata")

	// Set a item reader, add an item and validate we get the item back
	coll.Add(testItemID)

	coll.itemReader = func(context.Context, graph.Service, string, string) (*details.OneDriveInfo, io.ReadCloser, error) {
		return &details.OneDriveInfo{
			ItemName: testItemName,
			Modified: t1,
		}, io.NopCloser(bytes.NewReader(testItemData)), nil
	}

	// Read items from the collection
	wg.Add(1)

	readItems := []data.Stream{}

	for item := range coll.Items() {
		readItems = append(readItems, item)
	}

	wg.Wait()

	// Expect only 1 item
	require.Len(t, readItems, 1)
	require.Equal(t, 1, collStatus.ObjectCount)
	require.Equal(t, 1, collStatus.Successful)

	// Validate item info and data
	readItem := readItems[0]
	readItemInfo := readItem.(data.StreamInfo)

	assert.Equal(t, testItemName, readItem.UUID())

	require.Implements(t, (*data.StreamModTime)(nil), readItem)
	mt := readItem.(data.StreamModTime)
	assert.Equal(t, t1, mt.ModTime())

	readData, err := io.ReadAll(readItem.ToReader())
	require.NoError(t, err)

	assert.Equal(t, testItemData, readData)
	require.NotNil(t, readItemInfo.Info())
	require.NotNil(t, readItemInfo.Info().OneDrive)
	assert.Equal(t, testItemName, readItemInfo.Info().OneDrive.ItemName)
	assert.Equal(t, driveFolderPath, readItemInfo.Info().OneDrive.ParentPath)
}

func (suite *OneDriveCollectionSuite) TestOneDriveCollectionReadError() {
	t := suite.T()
	collStatus := support.ConnectorOperationStatus{}
	wg := sync.WaitGroup{}
	wg.Add(1)

	folderPath, err := GetCanonicalPath("drive/driveID1/root:/folderPath", "a-tenant", "a-user", OneDriveSource)
	require.NoError(t, err)

	coll := NewCollection(folderPath, "fakeDriveID", suite, suite.testStatusUpdater(&wg, &collStatus))
	coll.Add("testItemID")

	readError := errors.New("Test error")

	coll.itemReader = func(context.Context, graph.Service, string, string) (*details.OneDriveInfo, io.ReadCloser, error) {
		return nil, nil, readError
	}

	coll.Items()
	wg.Wait()
	// Expect no items
	require.Equal(t, 1, collStatus.ObjectCount)
	require.Equal(t, 0, collStatus.Successful)
}
