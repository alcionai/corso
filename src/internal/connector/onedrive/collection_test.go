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

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/connector/onedrive/mock"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

type mockItemGetterAndAugmenter struct {
	gi      mock.GetsItem
	info    details.ItemInfo
	getCall int
	getResp []*http.Response
	getErr  []error
}

func (m mockItemGetterAndAugmenter) GetItem(context.Context, string, string) (models.DriveItemable, error) {
	return m.gi.GetItem(nil, "", "")
}

func (m mockItemGetterAndAugmenter) AugmentItemInfo(
	_ details.ItemInfo,
	_ models.DriveItemable,
	_ int64,
	_ *path.Builder,
) details.ItemInfo {
	return m.info
}

func (m *mockItemGetterAndAugmenter) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	c := m.getCall
	m.getCall++

	return m.getResp[c], m.getErr[c]
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

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
		testItemMeta = metadata.Metadata{
			Permissions: []metadata.Permission{
				{
					ID:         "testMetaID",
					Roles:      []string{"read", "write"},
					Email:      "email@provider.com",
					Expiration: &now,
				},
			},
		}
	)

	type nst struct {
		name string
		size int64
		time time.Time
	}

	table := []struct {
		name         string
		numInstances int
		service      path.ServiceType
		itemInfo     details.ItemInfo
		getBody      io.ReadCloser
		getErr       error
		itemDeets    nst
		infoFrom     func(*testing.T, details.ItemInfo) (string, string)
		expectErr    require.ErrorAssertionFunc
		expectLabels []string
	}{
		{
			name:         "oneDrive, no duplicates",
			numInstances: 1,
			service:      path.OneDriveService,
			itemDeets:    nst{testItemName, 42, now},
			itemInfo:     details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: testItemName, Modified: now}},
			getBody:      io.NopCloser(bytes.NewReader(testItemData)),
			getErr:       nil,
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
			expectErr: require.NoError,
		},
		{
			name:         "oneDrive, duplicates",
			numInstances: 3,
			service:      path.OneDriveService,
			itemDeets:    nst{testItemName, 42, now},
			getBody:      io.NopCloser(bytes.NewReader(testItemData)),
			getErr:       nil,
			itemInfo:     details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: testItemName, Modified: now}},
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.OneDrive)
				return dii.OneDrive.ItemName, dii.OneDrive.ParentPath
			},
			expectErr: require.NoError,
		},
		{
			name:         "oneDrive, malware",
			numInstances: 3,
			service:      path.OneDriveService,
			itemDeets:    nst{testItemName, 42, now},
			itemInfo:     details.ItemInfo{},
			getBody:      nil,
			getErr:       clues.New("test malware").Label(graph.LabelsMalware),
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
			service:      path.OneDriveService,
			itemDeets:    nst{testItemName, 42, now},
			itemInfo:     details.ItemInfo{},
			getBody:      nil,
			getErr:       clues.New("test not found").Label(graph.LabelStatus(http.StatusNotFound)),
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
			service:      path.SharePointService,
			itemDeets:    nst{testItemName, 42, now},
			itemInfo:     details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: testItemName, Modified: now}},
			getBody:      io.NopCloser(bytes.NewReader(testItemData)),
			getErr:       nil,
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.SharePoint)
				return dii.SharePoint.ItemName, dii.SharePoint.ParentPath
			},
			expectErr: require.NoError,
		},
		{
			name:         "sharePoint, duplicates",
			numInstances: 3,
			service:      path.SharePointService,
			itemDeets:    nst{testItemName, 42, now},
			itemInfo:     details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: testItemName, Modified: now}},
			getBody:      io.NopCloser(bytes.NewReader(testItemData)),
			getErr:       nil,
			infoFrom: func(t *testing.T, dii details.ItemInfo) (string, string) {
				require.NotNil(t, dii.SharePoint)
				return dii.SharePoint.ItemName, dii.SharePoint.ParentPath
			},
			expectErr: require.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				wg         = sync.WaitGroup{}
				collStatus = support.ConnectorOperationStatus{}
				readItems  = []data.Stream{}
			)

			pb := path.Builder{}.Append(path.Split("drive/driveID1/root:/dir1/dir2/dir3")...)

			folderPath, err := pb.ToDataLayerOneDrivePath("tenant", "owner", false)
			require.NoError(t, err, clues.ToCore(err))

			driveFolderPath, err := path.GetDriveFolderPath(folderPath)
			require.NoError(t, err, clues.ToCore(err))

			mbh := &mock.BackupHandler{
				ItemInfo: details.ItemInfo{},
				GetResps: []*http.Response{{Body: test.getBody}},
				GetErrs:  []error{test.getErr},
				GI:       mock.GetsItem{Err: assert.AnError},
				GIP:      mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()},
			}

			coll, err := NewCollection(
				mbh,
				folderPath,
				nil,
				"drive-id",
				suite.testStatusUpdater(&wg, &collStatus),
				control.Options{ToggleFeatures: control.Toggles{}},
				CollectionScopeFolder,
				true)
			require.NoError(t, err, clues.ToCore(err))
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

			assert.Equal(t, testItemID+metadata.DataFileSuffix, readItem.UUID())

			require.Implements(t, (*data.StreamModTime)(nil), readItem)
			mt := readItem.(data.StreamModTime)
			assert.Equal(t, now, mt.ModTime())

			readData, err := io.ReadAll(readItem.ToReader())
			test.expectErr(t, err)

			if err != nil {
				for _, label := range test.expectLabels {
					assert.Truef(t, clues.HasLabel(err, label), "has clues label: %s", label)
				}

				return
			}

			name, parentPath := test.infoFrom(t, readItemInfo.Info())

			assert.Equal(t, testItemData, readData)
			assert.Equal(t, testItemName, name)
			assert.Equal(t, driveFolderPath, parentPath)

			readItemMeta := readItems[1]

			assert.Equal(t, testItemID+metadata.MetaFileSuffix, readItemMeta.UUID())

			readMetaData, err := io.ReadAll(readItemMeta.ToReader())
			require.NoError(t, err, clues.ToCore(err))

			tm, err := json.Marshal(testItemMeta)
			if err != nil {
				t.Fatal("unable to marshall test permissions", err)
			}

			assert.Equal(t, tm, readMetaData)
		})
	}
}

func (suite *CollectionUnitTestSuite) TestCollectionReadError() {
	var (
		t                = suite.T()
		testItemID       = "fakeItemID"
		collStatus       = support.ConnectorOperationStatus{}
		wg               = sync.WaitGroup{}
		name             = "name"
		size       int64 = 42
		now              = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	wg.Add(1)

	pb := path.Builder{}.Append(path.Split("drive/driveID1/root:/folderPath")...)
	folderPath, err := pb.ToDataLayerOneDrivePath("a-tenant", "a-user", false)
	require.NoError(t, err, clues.ToCore(err))

	mbh := &mock.BackupHandler{
		ItemInfo: details.ItemInfo{},
		GetResps: []*http.Response{
			nil,
			{Body: io.NopCloser(strings.NewReader("test"))},
		},
		GetErrs: []error{
			clues.Stack(assert.AnError).Label(graph.LabelStatus(http.StatusUnauthorized)),
			nil,
		},
		GI:  mock.GetsItem{Err: assert.AnError},
		GIP: mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()},
	}

	coll, err := NewCollection(
		mbh,
		folderPath,
		nil,
		"fakeDriveID",
		suite.testStatusUpdater(&wg, &collStatus),
		control.Options{ToggleFeatures: control.Toggles{}},
		CollectionScopeFolder,
		true)
	require.NoError(t, err, clues.ToCore(err))

	mockItem := models.NewDriveItem()
	mockItem.SetId(&testItemID)
	mockItem.SetFile(models.NewFile())
	mockItem.SetName(&name)
	mockItem.SetSize(&size)
	mockItem.SetCreatedDateTime(&now)
	mockItem.SetLastModifiedDateTime(&now)
	coll.Add(mockItem)

	collItem, ok := <-coll.Items(ctx, fault.New(true))
	assert.True(t, ok)

	_, err = io.ReadAll(collItem.ToReader())
	assert.Error(t, err, clues.ToCore(err))

	wg.Wait()

	// Expect no items
	require.Equal(t, 1, collStatus.Metrics.Objects, "only one object should be counted")
	require.Equal(t, 1, collStatus.Metrics.Successes, "TODO: should be 0, but allowing 1 to reduce async management")
}

func (suite *CollectionUnitTestSuite) TestCollectionReadUnauthorizedErrorRetry() {
	var (
		t                = suite.T()
		testItemID       = "fakeItemID"
		collStatus       = support.ConnectorOperationStatus{}
		wg               = sync.WaitGroup{}
		name             = "name"
		size       int64 = 42
		now              = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	wg.Add(1)

	mockItem := models.NewDriveItem()
	mockItem.SetId(&testItemID)
	mockItem.SetFile(models.NewFile())
	mockItem.SetName(&name)
	mockItem.SetSize(&size)
	mockItem.SetCreatedDateTime(&now)
	mockItem.SetLastModifiedDateTime(&now)

	pb := path.Builder{}.Append(path.Split("drive/driveID1/root:/folderPath")...)
	folderPath, err := pb.ToDataLayerOneDrivePath("a-tenant", "a-user", false)
	require.NoError(t, err)

	mbh := &mock.BackupHandler{
		ItemInfo: details.ItemInfo{},
		GetResps: []*http.Response{
			nil,
			{Body: io.NopCloser(strings.NewReader("test"))},
		},
		GetErrs: []error{
			clues.Stack(assert.AnError).Label(graph.LabelStatus(http.StatusUnauthorized)),
			nil,
		},
		GI:  mock.GetsItem{Item: mockItem},
		GIP: mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()},
	}

	coll, err := NewCollection(
		mbh,
		folderPath,
		nil,
		"fakeDriveID",
		suite.testStatusUpdater(&wg, &collStatus),
		control.Options{ToggleFeatures: control.Toggles{}},
		CollectionScopeFolder,
		true)
	require.NoError(t, err, clues.ToCore(err))

	coll.Add(mockItem)

	count := 0

	collItem, ok := <-coll.Items(ctx, fault.New(true))
	assert.True(t, ok)

	_, err = io.ReadAll(collItem.ToReader())
	assert.NoError(t, err, clues.ToCore(err))

	wg.Wait()

	require.Equal(t, 1, collStatus.Metrics.Objects, "only one object should be counted")
	require.Equal(t, 1, collStatus.Metrics.Successes, "read object successfully")
	require.Equal(t, 1, count, "retry count")
}

// Ensure metadata file always uses latest time for mod time
func (suite *CollectionUnitTestSuite) TestCollectionPermissionBackupLatestModTime() {
	var (
		t            = suite.T()
		testItemID   = "fakeItemID"
		testItemName = "Fake Item"
		testItemSize = int64(10)

		collStatus = support.ConnectorOperationStatus{}
		wg         = sync.WaitGroup{}
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	wg.Add(1)

	pb := path.Builder{}.Append(path.Split("drive/driveID1/root:/folderPath")...)
	folderPath, err := pb.ToDataLayerOneDrivePath("a-tenant", "a-user", false)
	require.NoError(t, err, clues.ToCore(err))

	dii := details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: "fakeName", Modified: time.Now()}}

	mbh := &mock.BackupHandler{
		ItemInfo: dii,
		GetResps: []*http.Response{{Body: io.NopCloser(strings.NewReader("Fake Data!"))}},
		GetErrs:  []error{nil},
		GIP:      mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()},
	}

	coll, err := NewCollection(
		mbh,
		folderPath,
		nil,
		"drive-id",
		suite.testStatusUpdater(&wg, &collStatus),
		control.Options{ToggleFeatures: control.Toggles{}},
		CollectionScopeFolder,
		true)
	require.NoError(t, err, clues.ToCore(err))

	mtime := time.Now().AddDate(0, -1, 0)
	mockItem := models.NewDriveItem()
	mockItem.SetFile(models.NewFile())
	mockItem.SetId(&testItemID)
	mockItem.SetName(&testItemName)
	mockItem.SetSize(&testItemSize)
	mockItem.SetCreatedDateTime(&mtime)
	mockItem.SetLastModifiedDateTime(&mtime)
	coll.Add(mockItem)

	coll.handler = mbh

	readItems := []data.Stream{}
	for item := range coll.Items(ctx, fault.New(true)) {
		readItems = append(readItems, item)
	}

	wg.Wait()

	// Expect no items
	require.Equal(t, 1, collStatus.Metrics.Objects)
	require.Equal(t, 1, collStatus.Metrics.Successes)

	for _, i := range readItems {
		if strings.HasSuffix(i.UUID(), metadata.MetaFileSuffix) {
			content, err := io.ReadAll(i.ToReader())
			require.NoError(t, err, clues.ToCore(err))
			require.Equal(t, content, []byte("{}"))

			im, ok := i.(data.StreamModTime)
			require.Equal(t, ok, true, "modtime interface")
			require.Greater(t, im.ModTime(), mtime, "permissions time greater than mod time")
		}
	}
}

type GetDriveItemUnitTestSuite struct {
	tester.Suite
}

func TestGetDriveItemUnitTestSuite(t *testing.T) {
	suite.Run(t, &GetDriveItemUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GetDriveItemUnitTestSuite) TestGetDriveItem_error() {
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
			err:      clues.New("malware error").Label(graph.LabelsMalware),
			labels:   []string{graph.LabelsMalware, graph.LabelsSkippable},
		},
		{
			name:     "file not found error",
			colScope: CollectionScopeFolder,
			itemSize: 10,
			err:      clues.New("not found error").Label(graph.LabelStatus(http.StatusNotFound)),
			labels:   []string{graph.LabelStatus(http.StatusNotFound), graph.LabelsSkippable},
		},
		{
			// This should create an error that stops the backup
			name:     "small OneNote file",
			colScope: CollectionScopePackage,
			itemSize: 10,
			err:      clues.New("small onenote error").Label(graph.LabelStatus(http.StatusServiceUnavailable)),
			labels:   []string{graph.LabelStatus(http.StatusServiceUnavailable)},
		},
		{
			name:     "big OneNote file",
			colScope: CollectionScopePackage,
			itemSize: MaxOneNoteFileSize,
			err:      clues.New("big onenote error").Label(graph.LabelStatus(http.StatusServiceUnavailable)),
			labels:   []string{graph.LabelStatus(http.StatusServiceUnavailable), graph.LabelsSkippable},
		},
		{
			// This should block backup, only big OneNote files should be a problem
			name:     "big file",
			colScope: CollectionScopeFolder,
			itemSize: MaxOneNoteFileSize,
			err:      clues.New("big file error").Label(graph.LabelStatus(http.StatusServiceUnavailable)),
			labels:   []string{graph.LabelStatus(http.StatusServiceUnavailable)},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				errs = fault.New(false)
				item = models.NewDriveItem()
				col  = &Collection{scope: test.colScope}
			)

			item.SetId(&strval)
			item.SetName(&strval)
			item.SetSize(&test.itemSize)

			mbh := &mock.BackupHandler{
				ItemInfo: details.ItemInfo{},
				GetResps: []*http.Response{nil},
				GetErrs:  []error{test.err},
				GI:       mock.GetsItem{Item: item},
			}

			col.handler = mbh

			_, err := col.getDriveItemContent(ctx, "driveID", item, errs)
			if test.err == nil {
				assert.NoError(t, err, clues.ToCore(err))
				return
			}

			assert.ErrorIs(t, err, test.err, clues.ToCore(err))

			labelsMap := map[string]struct{}{}
			for _, l := range test.labels {
				labelsMap[l] = struct{}{}
			}

			assert.Equal(t, labelsMap, clues.Labels(err))
		})
	}
}

func (suite *GetDriveItemUnitTestSuite) TestDownloadContent() {
	var (
		driveID   string
		iorc      = io.NopCloser(bytes.NewReader([]byte("fnords")))
		item      = models.NewDriveItem()
		itemWID   = models.NewDriveItem()
		errUnauth = clues.Stack(assert.AnError).Label(graph.LabelStatus(http.StatusUnauthorized))
	)

	itemWID.SetId(ptr.To("brainhooldy"))

	table := []struct {
		name      string
		mgi       mock.GetsItem
		itemInfo  details.ItemInfo
		respBody  []io.ReadCloser
		getErr    []error
		expectErr require.ErrorAssertionFunc
		expect    require.ValueAssertionFunc
	}{
		{
			name:      "good",
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{iorc},
			getErr:    []error{nil},
			expectErr: require.NoError,
			expect:    require.NotNil,
		},
		{
			name:      "expired url redownloads",
			mgi:       mock.GetsItem{Item: itemWID, Err: nil},
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil, iorc},
			getErr:    []error{errUnauth, nil},
			expectErr: require.NoError,
			expect:    require.NotNil,
		},
		{
			name:      "immediate error",
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil},
			getErr:    []error{assert.AnError},
			expectErr: require.Error,
			expect:    require.Nil,
		},
		{
			name:      "re-fetching the item fails",
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil},
			getErr:    []error{errUnauth},
			mgi:       mock.GetsItem{Item: nil, Err: assert.AnError},
			expectErr: require.Error,
			expect:    require.Nil,
		},
		{
			name:      "expired url fails redownload",
			mgi:       mock.GetsItem{Item: itemWID, Err: nil},
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil, iorc},
			getErr:    []error{errUnauth, assert.AnError},
			expectErr: require.Error,
			expect:    require.Nil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			resps := make([]*http.Response, 0, len(test.respBody))

			for _, v := range test.respBody {
				resps = append(resps, &http.Response{Body: v})
			}

			igaa := &mockItemGetterAndAugmenter{
				gi:      test.mgi,
				info:    test.itemInfo,
				getResp: resps,
				getErr:  test.getErr,
			}

			r, err := downloadContent(ctx, igaa, item, driveID)

			test.expect(t, r)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
