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

	"github.com/alcionai/clues"
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
	tester.Suite
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
	suite.Run(t, &CollectionUnitTestSuite{Suite: tester.NewUnitSuite(t)})
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
		expectErr    require.ErrorAssertionFunc
		expectLabels []string
	}{
		{
			name:         "oneDrive, no duplicates",
			numInstances: 1,
			source:       OneDriveSource,
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(context.Context, *http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
			expectErr: require.NoError,
		},
		{
			name:         "oneDrive, duplicates",
			numInstances: 3,
			source:       OneDriveSource,
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(context.Context, *http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
			expectErr: require.NoError,
		},
		{
			name:         "oneDrive, malware",
			numInstances: 3,
			source:       OneDriveSource,
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(context.Context, *http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{}, nil, clues.New("test malware").Label(graph.LabelsMalware)
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
			expectErr:    require.Error,
			expectLabels: []string{graph.LabelsMalware, graph.LabelsSkippable},
		},
		{
			name:         "oneDrive, not found",
			numInstances: 3,
			source:       OneDriveSource,
			itemDeets:    nst{testItemName, 42, now},
			// Usually `Not Found` is returned from itemGetter and not itemReader
			itemReader: func(context.Context, *http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{}, nil, clues.New("test not found").Label(graph.LabelStatus(http.StatusNotFound))
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
			expectErr:    require.Error,
			expectLabels: []string{graph.LabelStatus(http.StatusNotFound), graph.LabelsSkippable},
		},
		{
			name:         "sharePoint, no duplicates",
			numInstances: 1,
			source:       SharePointSource,
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(context.Context, *http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.SharePoint)
				return dii.SharePoint.ItemName, dii.SharePoint.ParentPath
			},
			expectErr: require.NoError,
		},
		{
			name:         "sharePoint, duplicates",
			numInstances: 3,
			source:       SharePointSource,
			itemDeets:    nst{testItemName, 42, now},
			itemReader: func(context.Context, *http.Client, models.DriveItemable) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: testItemName, Modified: now}},
					io.NopCloser(bytes.NewReader(testItemData)),
					nil
			},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.SharePoint)
				return dii.SharePoint.ItemName, dii.SharePoint.ParentPath
			},
			expectErr: require.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t          = suite.T()
				wg         = sync.WaitGroup{}
				collStatus = support.ConnectorOperationStatus{}
				readItems  = []data.Stream{}
			)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/dir1/dir2/dir3", "tenant", "owner", test.source)
			require.NoError(t, err, clues.ToCore(err))
			driveFolderPath, err := path.GetDriveFolderPath(folderPath)
			require.NoError(t, err, clues.ToCore(err))

			coll := NewCollection(
				graph.HTTPClient(graph.NoTimeout()),
				folderPath,
				nil,
				"drive-id",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{}},
				CollectionScopeFolder,
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

			require.Len(t, readItems, 2) // .data and .meta

			// Expect only 1 item
			require.Equal(t, 1, collStatus.Metrics.Objects)
			require.Equal(t, 1, collStatus.Metrics.Successes)

			// Validate item info and data
			readItem := readItems[0]
			readItemInfo := readItem.(data.StreamInfo)

			assert.Equal(t, testItemID+DataFileSuffix, readItem.UUID())

			require.Implements(t, (*data.StreamModTime)(nil), readItem)
			mt := readItem.(data.StreamModTime)
			assert.Equal(t, now, mt.ModTime())

			readData, err := io.ReadAll(readItem.ToReader())
			test.expectErr(t, err)

			if err != nil {
				for _, label := range test.expectLabels {
					assert.True(t, clues.HasLabel(err, label), "has clues label:", label)
				}

				return
			}

			name, parentPath := test.infoFrom(t, readItemInfo.Info())

			assert.Equal(t, testItemData, readData)
			assert.Equal(t, testItemName, name)
			assert.Equal(t, driveFolderPath, parentPath)

			if test.source == OneDriveSource {
				readItemMeta := readItems[1]

				assert.Equal(t, testItemID+MetaFileSuffix, readItemMeta.UUID())

				readMetaData, err := io.ReadAll(readItemMeta.ToReader())
				require.NoError(t, err, clues.ToCore(err))

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
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t          = suite.T()
				testItemID = "fakeItemID"
				collStatus = support.ConnectorOperationStatus{}
				wg         = sync.WaitGroup{}
			)

			wg.Add(1)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/folderPath", "a-tenant", "a-user", test.source)
			require.NoError(t, err, clues.ToCore(err))

			coll := NewCollection(
				graph.HTTPClient(graph.NoTimeout()),
				folderPath,
				nil,
				"fakeDriveID",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{}},
				CollectionScopeFolder,
				true)

			mockItem := models.NewDriveItem()
			mockItem.SetId(&testItemID)
			mockItem.SetFile(models.NewFile())
			mockItem.SetName(&name)
			mockItem.SetSize(&size)
			mockItem.SetCreatedDateTime(&now)
			mockItem.SetLastModifiedDateTime(&now)
			coll.Add(mockItem)

			coll.itemReader = func(
				context.Context,
				*http.Client,
				models.DriveItemable,
			) (details.ItemInfo, io.ReadCloser, error) {
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
			assert.Error(t, err, clues.ToCore(err))

			wg.Wait()

			// Expect no items
			require.Equal(t, 1, collStatus.Metrics.Objects, "only one object should be counted")
			require.Equal(t, 1, collStatus.Metrics.Successes, "TODO: should be 0, but allowing 1 to reduce async management")
		})
	}
}

func (suite *CollectionUnitTestSuite) TestCollectionReadUnauthorizedErrorRetry() {
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
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t          = suite.T()
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
				control.Options{ToggleFeatures: control.Toggles{}},
				CollectionScopeFolder,
				true)

			mockItem := models.NewDriveItem()
			mockItem.SetId(&testItemID)
			mockItem.SetFile(models.NewFile())
			mockItem.SetName(&name)
			mockItem.SetSize(&size)
			mockItem.SetCreatedDateTime(&now)
			mockItem.SetLastModifiedDateTime(&now)
			coll.Add(mockItem)

			count := 0

			coll.itemGetter = func(
				ctx context.Context,
				srv graph.Servicer,
				driveID, itemID string,
			) (models.DriveItemable, error) {
				return mockItem, nil
			}

			coll.itemReader = func(
				context.Context,
				*http.Client,
				models.DriveItemable,
			) (details.ItemInfo, io.ReadCloser, error) {
				if count < 2 {
					count++
					return details.ItemInfo{}, nil, clues.Stack(assert.AnError).
						Label(graph.LabelStatus(http.StatusUnauthorized))
				}

				return details.ItemInfo{}, io.NopCloser(strings.NewReader("test")), nil
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
			assert.NoError(t, err)

			wg.Wait()

			require.Equal(t, 1, collStatus.Metrics.Objects, "only one object should be counted")
			require.Equal(t, 1, collStatus.Metrics.Successes, "read object successfully")
			require.Equal(t, 2, count, "retry count")
		})
	}
}

// Ensure metadata file always uses latest time for mod time
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
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t            = suite.T()
				testItemID   = "fakeItemID"
				testItemName = "Fake Item"
				testItemSize = int64(10)

				collStatus = support.ConnectorOperationStatus{}
				wg         = sync.WaitGroup{}
			)

			wg.Add(1)

			folderPath, err := GetCanonicalPath("drive/driveID1/root:/folderPath", "a-tenant", "a-user", test.source)
			require.NoError(t, err, clues.ToCore(err))

			coll := NewCollection(
				graph.HTTPClient(graph.NoTimeout()),
				folderPath,
				nil,
				"drive-id",
				suite,
				suite.testStatusUpdater(&wg, &collStatus),
				test.source,
				control.Options{ToggleFeatures: control.Toggles{}},
				CollectionScopeFolder,
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
				context.Context,
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
					require.NoError(t, err, clues.ToCore(err))
					require.Equal(t, content, []byte("{}"))

					im, ok := i.(data.StreamModTime)
					require.Equal(t, ok, true, "modtime interface")
					require.Greater(t, im.ModTime(), mtime, "permissions time greater than mod time")
				}
			}
		})
	}
}

type GetDriveItemUnitTestSuite struct {
	tester.Suite
}

func TestGetDriveItemUnitTestSuite(t *testing.T) {
	suite.Run(t, &GetDriveItemUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GetDriveItemUnitTestSuite) TestGetDriveItemError() {
	strval := "not-important"

	table := []struct {
		name     string
		colScope collectionScope
		itemSize int64
		labels   []string
		err      error
	}{
		{
			name:     "Simple item fetch no error",
			colScope: CollectionScopeFolder,
			itemSize: 10,
			err:      nil,
		},
		{
			name:     "Simple item fetch error",
			colScope: CollectionScopeFolder,
			itemSize: 10,
			err:      assert.AnError,
		},
		{
			name:     "malware error",
			colScope: CollectionScopeFolder,
			itemSize: 10,
			err:      clues.New("test error").Label(graph.LabelsMalware),
			labels:   []string{graph.LabelsMalware, graph.LabelsSkippable},
		},
		{
			name:     "file not found error",
			colScope: CollectionScopeFolder,
			itemSize: 10,
			err:      clues.New("test error").Label(graph.LabelStatus(http.StatusNotFound)),
			labels:   []string{graph.LabelStatus(http.StatusNotFound), graph.LabelsSkippable},
		},
		{
			// This should create an error that stops the backup
			name:     "small OneNote file",
			colScope: CollectionScopePackage,
			itemSize: 10,
			err:      clues.New("test error").Label(graph.LabelStatus(http.StatusServiceUnavailable)),
			labels:   []string{graph.LabelStatus(http.StatusServiceUnavailable)},
		},
		{
			name:     "big OneNote file",
			colScope: CollectionScopePackage,
			itemSize: MaxOneNoteFileSize,
			err:      clues.New("test error").Label(graph.LabelStatus(http.StatusServiceUnavailable)),
			labels:   []string{graph.LabelStatus(http.StatusServiceUnavailable), graph.LabelsSkippable},
		},
		{
			// This should block backup, only big OneNote files should be a problem
			name:     "big file",
			colScope: CollectionScopeFolder,
			itemSize: MaxOneNoteFileSize,
			err:      clues.New("test error").Label(graph.LabelStatus(http.StatusServiceUnavailable)),
			labels:   []string{graph.LabelStatus(http.StatusServiceUnavailable)},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t    = suite.T()
				errs = fault.New(false)
				item = models.NewDriveItem()
				col  = &Collection{scope: test.colScope}
			)

			item.SetId(&strval)
			item.SetName(&strval)
			item.SetSize(&test.itemSize)

			col.itemReader = func(
				ctx context.Context,
				hc *http.Client,
				item models.DriveItemable,
			) (details.ItemInfo, io.ReadCloser, error) {
				return details.ItemInfo{}, nil, test.err
			}

			col.itemGetter = func(
				ctx context.Context,
				srv graph.Servicer,
				driveID, itemID string,
			) (models.DriveItemable, error) {
				// We are not testing this err here
				return item, nil
			}

			_, err := col.getDriveItemContent(ctx, item, errs)
			if test.err == nil {
				assert.NoError(t, err, "no error")
				return
			}

			assert.EqualError(t, err, clues.Wrap(test.err, "downloading item").Error(), "error")

			labelsMap := map[string]struct{}{}
			for _, l := range test.labels {
				labelsMap[l] = struct{}{}
			}

			assert.Equal(t, labelsMap, clues.Labels(err))
		})
	}
}
