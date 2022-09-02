package onedrive

import (
	"bytes"
	"context"
	"errors"
	"io"
	"path/filepath"
	"sync"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
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
	wg := sync.WaitGroup{}
	collStatus := support.ConnectorOperationStatus{}
	folderPath := "dir1/dir2/dir3"
	coll := NewCollection(folderPath, "fakeDriveID", suite, suite.testStatusUpdater(&wg, &collStatus))
	require.NotNil(suite.T(), coll)
	assert.Equal(suite.T(), filepath.SplitList(folderPath), coll.FullPath())

	testItemID := "fakeItemID"
	testItemName := "itemName"
	testItemData := []byte("testdata")

	// Set a item reader, add an item and validate we get the item back
	coll.Add(testItemID)

	coll.itemReader = func(context.Context, graph.Service, string, string) (string, io.ReadCloser, error) {
		return testItemName, io.NopCloser(bytes.NewReader(testItemData)), nil
	}

	// Read items from the collection
	wg.Add(1)
	readItems := []data.Stream{}
	for item := range coll.Items() {
		readItems = append(readItems, item)
	}
	wg.Wait()
	// Expect only 1 item
	require.Len(suite.T(), readItems, 1)
	require.Equal(suite.T(), 1, collStatus.ObjectCount)
	require.Equal(suite.T(), 1, collStatus.Successful)

	// Validate item info and data
	readItem := readItems[0]
	readItemInfo := readItem.(data.StreamInfo)

	assert.Equal(suite.T(), testItemID, readItem.UUID())
	readData, err := io.ReadAll(readItem.ToReader())
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), testItemData, readData)
	require.NotNil(suite.T(), readItemInfo.Info())
	require.NotNil(suite.T(), readItemInfo.Info().OneDrive)
	assert.Equal(suite.T(), testItemName, readItemInfo.Info().OneDrive.ItemName)
	assert.Equal(suite.T(), folderPath, readItemInfo.Info().OneDrive.ParentPath)
}

func (suite *OneDriveCollectionSuite) TestOneDriveCollectionReadError() {
	wg := sync.WaitGroup{}
	collStatus := support.ConnectorOperationStatus{}
	wg.Add(1)

	coll := NewCollection("folderPath", "fakeDriveID", suite, suite.testStatusUpdater(&wg, &collStatus))
	coll.Add("testItemID")

	readError := errors.New("Test error")

	coll.itemReader = func(context.Context, graph.Service, string, string) (name string, data io.ReadCloser, err error) {
		return "", nil, readError
	}

	coll.Items()
	wg.Wait()
	// Expect no items
	require.Equal(suite.T(), 1, collStatus.ObjectCount)
	require.Equal(suite.T(), 0, collStatus.Successful)
}
