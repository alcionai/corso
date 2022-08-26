package onedrive

import (
	"bytes"
	"context"
	"errors"
	"io"
	"path/filepath"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/graph"
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

func (suite *OneDriveCollectionSuite) TestOneDriveCollection() {
	folderPath := "dir1/dir2/dir3"
	coll := NewCollection(folderPath, "fakeDriveID", suite, nil)
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
	readItems := []data.Stream{}
	for item := range coll.Items() {
		readItems = append(readItems, item)
	}

	// Expect only 1 item
	require.Len(suite.T(), readItems, 1)

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
	coll := NewCollection("folderPath", "fakeDriveID", suite, nil)
	coll.Add("testItemID")

	readError := errors.New("Test error")

	coll.itemReader = func(context.Context, graph.Service, string, string) (name string, data io.ReadCloser, err error) {
		return "", nil, readError
	}

	// Expect no items
	require.Len(suite.T(), coll.Items(), 0)
}
