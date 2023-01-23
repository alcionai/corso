package onedrive

import (
	"bytes"
	"context"
	"errors"
	"io"
	"sync"
	"testing"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

type CollectionUnitTestSuite struct {
	suite.Suite
}

// Allows `*CollectionUnitTestSuite` to be used as a graph.Servicer
// TODO: Implement these methods

func (suite *CollectionUnitTestSuite) Client() *msgraphsdk.GraphServiceClient {
	return nil
}

func (suite *CollectionUnitTestSuite) Adapter() *msgraphsdk.GraphRequestAdapter {
	return nil
}

func TestCollectionUnitTestSuite(t *testing.T) {
	suite.Run(t, new(CollectionUnitTestSuite))
}

// Returns a status update function that signals the specified WaitGroup when it is done
func (suite *CollectionUnitTestSuite) testStatusUpdater(
	wg *sync.WaitGroup,
	statusToUpdate *support.ConnectorOperationStatus,
) support.StatusUpdater {
	return func(s *support.ConnectorOperationStatus) {
		suite.T().Logf("Update status %v, count %d, success %d", s, s.ObjectCount, s.Successful)
		*statusToUpdate = *s

		wg.Done()
	}
}

func (suite *CollectionUnitTestSuite) TestCollection() {
	var (
		testItemID   = "fakeItemID"
		testItemName = "itemName"
		testItemData = []byte("testdata")
		now          = time.Now()
	)

	table := []struct {
		name         string
		numInstances int
		source       driveSource
		itemReader   itemReaderFunc
		infoFrom     func(*testing.T, details.ItemInfo) (string, string)
	}{
		{
			name:         "oneDrive, no duplicates",
			numInstances: 1,
			source:       OneDriveSource,
			itemReader: func(context.Context, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
		},
		{
			name:         "oneDrive, duplicates",
			numInstances: 3,
			source:       OneDriveSource,
			itemReader: func(context.Context, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
		},
		{
			name:         "sharePoint, no duplicates",
			numInstances: 1,
			source:       SharePointSource,
			itemReader: func(context.Context, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.SharePoint)
				return dii.SharePoint.ItemName, dii.SharePoint.ParentPath
			},
		},
		{
			name:         "sharePoint, duplicates",
			numInstances: 3,
			source:       SharePointSource,
			itemReader: func(context.Context, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.SharePoint)
				return dii.SharePoint.ItemName, dii.SharePoint.ParentPath
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			var (
				wg         = sync.WaitGroup{}
				collStatus = support.ConnectorOperationStatus{}
				readItems  = []data.Stream{}
			)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/dir1/dir2/dir3", "tenant", "owner", test.source)
			require.NoError(t, err)
			driveFolderPath, err := path.GetDriveFolderPath(folderPath)
			require.NoError(t, err)

			coll := NewCollection(
				folderPath,
				"drive-id",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{})
			require.NotNil(t, coll)
			assert.Equal(t, folderPath, coll.FullPath())

			// Set a item reader, add an item and validate we get the item back
			mockItem := models.NewDriveItem()
			mockItem.SetId(&testItemID)

			for i := 0; i < test.numInstances; i++ {
				coll.Add(mockItem)
			}

			coll.itemReader = test.itemReader

			// Read items from the collection
			wg.Add(1)

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
			assert.Equal(t, now, mt.ModTime())

			readData, err := io.ReadAll(readItem.ToReader())
			require.NoError(t, err)

			name, parentPath := test.infoFrom(t, readItemInfo.Info())

			assert.Equal(t, testItemData, readData)
			assert.Equal(t, testItemName, name)
			assert.Equal(t, driveFolderPath, parentPath)
		})
	}
}

func (suite *CollectionUnitTestSuite) TestCollectionReadError() {
	table := []struct {
		name   string
		source driveSource
	}{
		{
			name:   "oneDrive",
			source: OneDriveSource,
		},
		{
			name:   "sharePoint",
			source: SharePointSource,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			var (
				testItemID = "fakeItemID"

				collStatus = support.ConnectorOperationStatus{}
				wg         = sync.WaitGroup{}
			)

			wg.Add(1)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/folderPath", "a-tenant", "a-user", test.source)
			require.NoError(t, err)

			coll := NewCollection(
				folderPath,
				"fakeDriveID",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{})

			mockItem := models.NewDriveItem()
			mockItem.SetId(&testItemID)
			coll.Add(mockItem)

			readError := errors.New("Test error")

			coll.itemReader = func(context.Context, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{}, nil, readError
			}

			coll.Items()
			wg.Wait()

			// Expect no items
			require.Equal(t, 1, collStatus.ObjectCount)
			require.Equal(t, 0, collStatus.Successful)
		})
	}
}
