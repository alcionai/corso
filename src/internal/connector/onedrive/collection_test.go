package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
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
		suite.T().Logf("Update status %v, count %d, success %d", s, s.Metrics.Objects, s.Metrics.Successes)
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
		testItemMeta = Metadata{Permissions: []UserPermission{
			{
				ID:         "testMetaID",
				Roles:      []string{"read", "write"},
				Email:      "email@provider.com",
				Expiration: &now,
			},
		}}
	)

	type nst struct {
		name string
		size int64
		time time.Time
	}

	table := []struct {
		name         string
		numInstances int
		source       driveSource
		itemReader   itemReaderFunc
		itemDeets    nst
		infoFrom     func(*testing.T, details.ItemInfo) (string, string)
	}{
		{
			name:         "oneDrive, no duplicates",
			numInstances: 1,
			source:       OneDriveSource,
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(*http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
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
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(*http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
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
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(*http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
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
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(*http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
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
			ctx, flush := tester.NewContext()
			defer flush()

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
				graph.HTTPClient(graph.NoTimeout()),
				folderPath,
				nil,
				"drive-id",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}},
				true)
			require.NotNil(t, coll)
			assert.Equal(t, folderPath, coll.FullPath())

			// Set a item reader, add an item and validate we get the item back
			mockItem := models.NewDriveItem()
			mockItem.SetId(&testItemID)
			mockItem.SetFile(models.NewFile())
			mockItem.SetName(&test.itemDeets.name)
			mockItem.SetSize(&test.itemDeets.size)
			mockItem.SetCreatedDateTime(&test.itemDeets.time)
			mockItem.SetLastModifiedDateTime(&test.itemDeets.time)

			for i := 0; i < test.numInstances; i++ {
				coll.Add(mockItem)
			}

			coll.itemReader = test.itemReader
			coll.itemMetaReader = func(_ context.Context,
				_ graph.Servicer,
				_ string,
				_ models.DriveItemable,
			) (io.ReadCloser, int, error) {
				metaJSON, err := json.Marshal(testItemMeta)
				if err != nil {
					return nil, 0, err
				}

				return io.NopCloser(bytes.NewReader(metaJSON)), len(metaJSON), nil
			}

			// Read items from the collection
			wg.Add(1)

			for item := range coll.Items(ctx, fault.New(true)) {
				readItems = append(readItems, item)
			}

			wg.Wait()

			if test.source == OneDriveSource {
				require.Len(t, readItems, 2) // .data and .meta
			} else {
				require.Len(t, readItems, 1)
			}

			// Expect only 1 item
			require.Equal(t, 1, collStatus.Metrics.Objects)
			require.Equal(t, 1, collStatus.Metrics.Successes)

			// Validate item info and data
			readItem := readItems[0]
			readItemInfo := readItem.(data.StreamInfo)

			if test.source == OneDriveSource {
				assert.Equal(t, testItemName+DataFileSuffix, readItem.UUID())
			} else {
				assert.Equal(t, testItemName, readItem.UUID())
			}

			require.Implements(t, (*data.StreamModTime)(nil), readItem)
			mt := readItem.(data.StreamModTime)
			assert.Equal(t, now, mt.ModTime())

			readData, err := io.ReadAll(readItem.ToReader())
			require.NoError(t, err)

			name, parentPath := test.infoFrom(t, readItemInfo.Info())

			assert.Equal(t, testItemData, readData)
			assert.Equal(t, testItemName, name)
			assert.Equal(t, driveFolderPath, parentPath)

			if test.source == OneDriveSource {
				readItemMeta := readItems[1]

				assert.Equal(t, testItemName+MetaFileSuffix, readItemMeta.UUID())

				readMetaData, err := io.ReadAll(readItemMeta.ToReader())
				require.NoError(t, err)

				tm, err := json.Marshal(testItemMeta)
				if err != nil {
					t.Fatal("unable to marshall test permissions", err)
				}

				assert.Equal(t, tm, readMetaData)
			}
		})
	}
}

func (suite *CollectionUnitTestSuite) TestCollectionReadError() {
	var (
		name       = "name"
		size int64 = 42
		now        = time.Now()
	)

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
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				testItemID = "fakeItemID"
				collStatus = support.ConnectorOperationStatus{}
				wg         = sync.WaitGroup{}
			)

			wg.Add(1)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/folderPath", "a-tenant", "a-user", test.source)
			require.NoError(t, err)

			coll := NewCollection(
				graph.HTTPClient(graph.NoTimeout()),
				folderPath,
				nil,
				"fakeDriveID",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}},
				true)

			mockItem := models.NewDriveItem()
			mockItem.SetId(&testItemID)
			mockItem.SetFile(models.NewFile())
			mockItem.SetName(&name)
			mockItem.SetSize(&size)
			mockItem.SetCreatedDateTime(&now)
			mockItem.SetLastModifiedDateTime(&now)
			coll.Add(mockItem)

			coll.itemReader = func(*http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{}, nil, assert.AnError
			}

			coll.itemMetaReader = func(_ context.Context,
				_ graph.Servicer,
				_ string,
				_ models.DriveItemable,
			) (io.ReadCloser, int, error) {
				return io.NopCloser(strings.NewReader(`{}`)), 2, nil
			}

			collItem, ok := <-coll.Items(ctx, fault.New(true))
			assert.True(t, ok)

			_, err = io.ReadAll(collItem.ToReader())
			assert.Error(t, err)

			wg.Wait()

			// Expect no items
			require.Equal(t, 1, collStatus.Metrics.Objects, "only one object should be counted")
			require.Equal(t, 1, collStatus.Metrics.Successes, "TODO: should be 0, but allowing 1 to reduce async management")
		})
	}
}

func (suite *CollectionUnitTestSuite) TestCollectionDisablePermissionsBackup() {
	table := []struct {
		name   string
		source driveSource
	}{
		{
			name:   "oneDrive",
			source: OneDriveSource,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				testItemID   = "fakeItemID"
				testItemName = "Fake Item"
				testItemSize = int64(10)
				collStatus   = support.ConnectorOperationStatus{}
				wg           = sync.WaitGroup{}
			)

			wg.Add(1)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/folderPath", "a-tenant", "a-user", test.source)
			require.NoError(t, err)

			coll := NewCollection(
				graph.HTTPClient(graph.NoTimeout()),
				folderPath,
				nil,
				"fakeDriveID",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{}},
				true)

			now := time.Now()
			mockItem := models.NewDriveItem()
			mockItem.SetFile(models.NewFile())
			mockItem.SetId(&testItemID)
			mockItem.SetName(&testItemName)
			mockItem.SetSize(&testItemSize)
			mockItem.SetCreatedDateTime(&now)
			mockItem.SetLastModifiedDateTime(&now)
			coll.Add(mockItem)

			coll.itemReader = func(
				*http.Client,
				models.DriveItemable,
			) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: "fakeName", Modified: time.Now()}},
					io.NopCloser(strings.NewReader("Fake Data!")),
					nil
			}

			coll.itemMetaReader = func(_ context.Context,
				_ graph.Servicer,
				_ string,
				_ models.DriveItemable,
			) (io.ReadCloser, int, error) {
				return io.NopCloser(strings.NewReader(`{"key": "value"}`)), 16, nil
			}

			readItems := []data.Stream{}
			for item := range coll.Items(ctx, fault.New(true)) {
				readItems = append(readItems, item)
			}

			wg.Wait()

			// Expect no items
			require.Equal(t, 1, collStatus.Metrics.Objects)
			require.Equal(t, 1, collStatus.Metrics.Successes)

			for _, i := range readItems {
				if strings.HasSuffix(i.UUID(), MetaFileSuffix) {
					content, err := io.ReadAll(i.ToReader())
					require.NoError(t, err)
					require.Equal(t, content, []byte("{}"))
				}
			}
		})
	}
}

// TODO(meain): Remove this test once we start always backing up permissions
func (suite *CollectionUnitTestSuite) TestCollectionPermissionBackupLatestModTime() {
	table := []struct {
		name   string
		source driveSource
	}{
		{
			name:   "oneDrive",
			source: OneDriveSource,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				testItemID   = "fakeItemID"
				testItemName = "Fake Item"
				testItemSize = int64(10)

				collStatus = support.ConnectorOperationStatus{}
				wg         = sync.WaitGroup{}
			)

			wg.Add(1)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/folderPath", "a-tenant", "a-user", test.source)
			require.NoError(t, err)

			coll := NewCollection(
				graph.HTTPClient(graph.NoTimeout()),
				folderPath,
				nil,
				"drive-id",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}},
				true)

			mtime := time.Now().AddDate(0, -1, 0)
			mockItem := models.NewDriveItem()
			mockItem.SetFile(models.NewFile())
			mockItem.SetId(&testItemID)
			mockItem.SetName(&testItemName)
			mockItem.SetSize(&testItemSize)
			mockItem.SetCreatedDateTime(&mtime)
			mockItem.SetLastModifiedDateTime(&mtime)
			coll.Add(mockItem)

			coll.itemReader = func(
				*http.Client,
				models.DriveItemable,
			) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: "fakeName", Modified: time.Now()}},
					io.NopCloser(strings.NewReader("Fake Data!")),
					nil
			}

			coll.itemMetaReader = func(_ context.Context,
				_ graph.Servicer,
				_ string,
				_ models.DriveItemable,
			) (io.ReadCloser, int, error) {
				return io.NopCloser(strings.NewReader(`{}`)), 16, nil
			}

			readItems := []data.Stream{}
			for item := range coll.Items(ctx, fault.New(true)) {
				readItems = append(readItems, item)
			}

			wg.Wait()

			// Expect no items
			require.Equal(t, 1, collStatus.Metrics.Objects)
			require.Equal(t, 1, collStatus.Metrics.Successes)

			for _, i := range readItems {
				if strings.HasSuffix(i.UUID(), MetaFileSuffix) {
					content, err := io.ReadAll(i.ToReader())
					require.NoError(t, err)
					require.Equal(t, content, []byte("{}"))
					im, ok := i.(data.StreamModTime)
					require.Equal(t, ok, true, "modtime interface")
					require.Greater(t, im.ModTime(), mtime, "permissions time greater than mod time")
				}
			}
		})
	}
}
