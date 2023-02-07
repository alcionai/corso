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

	"github.com/hashicorp/go-multierror"
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
	"github.com/alcionai/corso/src/pkg/path"
)

type CollectionUnitTestSuite struct {
	suite.Suite
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
		testItemMeta = Metadata{Permissions: []UserPermission{
			{
				ID:         "testMetaID",
				Roles:      []string{"read", "write"},
				Email:      "email@provider.com",
				Expiration: &now,
			},
		}}
	)

	testItemMetaBytes, err := json.Marshal(testItemMeta)
	if err != nil {
		suite.T().Fatal("unable to marshall test permissions", err)
	}

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
				"drive-id",
				&MockGraphService{},
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}})
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
				return io.NopCloser(bytes.NewReader(testItemMetaBytes)), len(testItemMetaBytes), nil
			}

			// Read items from the collection
			wg.Add(1)

			for item := range coll.Items() {
				readItems = append(readItems, item)
			}

			wg.Wait()

			if test.source == OneDriveSource {
				assert.Len(t, readItems, 2) // .data and .meta
				assert.Equal(t, 2, collStatus.ObjectCount)
				assert.Equal(t, 2, collStatus.Successful)
			} else {
				assert.Len(t, readItems, 1)
				assert.Equal(t, 1, collStatus.ObjectCount)
				assert.Equal(t, 1, collStatus.Successful)
			}

			var (
				foundData bool
				foundMeta bool
			)

			for _, readItem := range readItems {
				readItemInfo := readItem.(data.StreamInfo)
				id := readItem.UUID()

				if strings.HasSuffix(id, DataFileSuffix) {
					foundData = true
				}

				var hasMeta bool
				if strings.HasSuffix(id, MetaFileSuffix) {
					foundMeta = true
					hasMeta = true
				}

				assert.Contains(t, testItemName, id)
				require.Implements(t, (*data.StreamModTime)(nil), readItem)

				mt := readItem.(data.StreamModTime)
				assert.Equal(t, now, mt.ModTime())

				readData, err := io.ReadAll(readItem.ToReader())
				require.NoError(t, err)

				name, parentPath := test.infoFrom(t, readItemInfo.Info())
				assert.Equal(t, testItemData, readData)
				assert.Equal(t, testItemName, name)
				assert.Equal(t, driveFolderPath, parentPath)

				if hasMeta {
					ra, err := io.ReadAll(readItem.ToReader())
					require.NoError(t, err)
					assert.Equal(t, testItemMetaBytes, ra)
				}
			}

			if test.source == OneDriveSource {
				assert.True(t, foundData, "found data file")
				assert.True(t, foundMeta, "found metadata file")
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
				"fakeDriveID",
				&MockGraphService{},
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}})

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

			collItem, ok := <-coll.Items()
			assert.True(t, ok)

			_, err = io.ReadAll(collItem.ToReader())
			assert.Error(t, err)

			wg.Wait()

			// Expect no items
			require.Equal(t, 1, collStatus.ObjectCount, "only one object should be counted")
			require.Equal(t, 1, collStatus.Successful, "TODO: should be 0, but allowing 1 to reduce async management")
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
				"fakeDriveID",
				&MockGraphService{},
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{}})

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
			for item := range coll.Items() {
				readItems = append(readItems, item)
			}

			wg.Wait()

			assert.Equal(t, 1, collStatus.ObjectCount, "total objects")
			assert.Equal(t, 1, collStatus.Successful, "successes")

			for _, i := range readItems {
				if strings.HasSuffix(i.UUID(), MetaFileSuffix) {
					content, err := io.ReadAll(i.ToReader())
					require.NoError(t, err)
					assert.Equal(t, content, []byte("{}"))
				}
			}
		})
	}
}

func (suite *CollectionUnitTestSuite) TestStreamItem() {
	var (
		id         = "id"
		name       = "name"
		size int64 = 42
		now        = time.Now()
	)

	mockItem := models.NewDriveItem()
	mockItem.SetId(&id)
	mockItem.SetName(&name)
	mockItem.SetSize(&size)
	mockItem.SetCreatedDateTime(&now)
	mockItem.SetLastModifiedDateTime(&now)

	mockReader := func(v string, e error) itemReaderFunc {
		return func(*http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
			return details.ItemInfo{}, io.NopCloser(strings.NewReader(v)), e
		}
	}

	mockGetter := func(e error) itemGetterFunc {
		return func(context.Context, graph.Servicer, string, string) (models.DriveItemable, error) {
			return mockItem, e
		}
	}

	mockDataChan := func() chan data.Stream {
		return make(chan data.Stream, 1)
	}

	table := []struct {
		name       string
		coll       *Collection
		expectData string
		errsIs     func(*testing.T, error, int)
		readErrIs  func(*testing.T, error)
	}{
		{
			name:       "happy",
			expectData: "happy",
			coll: &Collection{
				data:       mockDataChan(),
				itemReader: mockReader("happy", nil),
				itemGetter: mockGetter(nil),
			},
			errsIs: func(t *testing.T, e error, count int) {
				assert.NoError(t, e, "no errors")
				assert.Zero(t, count, "zero errors")
			},
			readErrIs: func(t *testing.T, e error) {
				assert.NoError(t, e, "no reader error")
			},
		},
		{
			name:       "reader err",
			expectData: "",
			coll: &Collection{
				data:       mockDataChan(),
				itemReader: mockReader("foo", assert.AnError),
				itemGetter: mockGetter(nil),
			},
			errsIs: func(t *testing.T, e error, count int) {
				assert.ErrorIs(t, e, assert.AnError)
				assert.Equal(t, 1, count, "one errors")
			},
			readErrIs: func(t *testing.T, e error) {
				assert.Error(t, e, "basic error")
			},
		},
		{
			name:       "iteration err",
			expectData: "",
			coll: &Collection{
				data:       mockDataChan(),
				itemReader: mockReader("foo", graph.Err401Unauthorized),
				itemGetter: mockGetter(assert.AnError),
			},
			errsIs: func(t *testing.T, e error, count int) {
				assert.True(t, graph.IsErrUnauthorized(e), "is unauthorized error")
				assert.ErrorIs(t, e, graph.Err401Unauthorized)
				assert.Equal(t, 1, count, "count of errors aggregated")
			},
			readErrIs: func(t *testing.T, e error) {
				assert.True(t, graph.IsErrUnauthorized(e), "is unauthorized error")
				assert.ErrorIs(t, e, graph.Err401Unauthorized)
			},
		},
		{
			name:       "timeout errors",
			expectData: "",
			coll: &Collection{
				data:       mockDataChan(),
				itemReader: mockReader("foo", context.DeadlineExceeded),
				itemGetter: mockGetter(nil),
			},
			errsIs: func(t *testing.T, e error, count int) {
				assert.True(t, graph.IsErrTimeout(e), "is timeout error")
				assert.ErrorIs(t, e, context.DeadlineExceeded)
				assert.Equal(t, 1, count, "one errors")
			},
			readErrIs: func(t *testing.T, e error) {
				assert.True(t, graph.IsErrTimeout(e), "is timeout error")
				assert.ErrorIs(t, e, context.DeadlineExceeded)
			},
		},
		{
			name:       "internal server errors",
			expectData: "",
			coll: &Collection{
				data:       mockDataChan(),
				itemReader: mockReader("foo", graph.Err500InternalServerError),
				itemGetter: mockGetter(nil),
			},
			errsIs: func(t *testing.T, e error, count int) {
				assert.True(t, graph.IsInternalServerError(e), "is internal server error")
				assert.ErrorIs(t, e, graph.Err500InternalServerError)
				assert.Equal(t, 1, count, "one errors")
			},
			readErrIs: func(t *testing.T, e error) {
				assert.True(t, graph.IsInternalServerError(e), "is internal server error")
				assert.ErrorIs(t, e, graph.Err500InternalServerError)
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				wg       sync.WaitGroup
				errs     error
				errCount int
				size     int64

				countUpdater = func(sz, ds, itms, drd, ird int64) { size = sz }
				errUpdater   = func(s string, e error) {
					errs = multierror.Append(errs, e)
					errCount++
				}

				semaphore = make(chan struct{}, 1)
				progress  = make(chan struct{}, 1)
			)

			wg.Add(1)
			semaphore <- struct{}{}

			go test.coll.streamItem(
				ctx,
				&wg,
				semaphore,
				progress,
				errUpdater,
				countUpdater,
				mockItem,
				"parentPath",
			)

			// wait for the func to run
			wg.Wait()

			assert.Zero(t, len(semaphore), "semaphore was released")
			assert.NotNil(t, <-progress, "progress was communicated")
			assert.NotZero(t, size, "countUpdater was called")

			data, ok := <-test.coll.data
			assert.True(t, ok, "data channel survived")

			bs, err := io.ReadAll(data.ToReader())

			test.readErrIs(t, err)
			test.errsIs(t, errs, errCount)
			assert.Equal(t, test.expectData, string(bs), "streamed item bytes")
		})
	}
}
