package drive

import (
	"bytes"
	"context"
	"encoding/json"
	"hash/crc32"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	metaTD "github.com/alcionai/corso/src/internal/m365/collection/drive/metadata/testdata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	odTD "github.com/alcionai/corso/src/internal/m365/service/onedrive/testdata"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/extensions"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type CollectionUnitTestSuite struct {
	tester.Suite
}

func TestCollectionUnitTestSuite(t *testing.T) {
	suite.Run(t, &CollectionUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

// Returns a status update function that signals the specified WaitGroup when it is done
func (suite *CollectionUnitTestSuite) testStatusUpdater(
	wg *sync.WaitGroup,
	statusToUpdate *support.ControllerOperationStatus,
) support.StatusUpdater {
	return func(s *support.ControllerOperationStatus) {
		suite.T().Logf("Update status %v, count %d, success %d", s, s.Metrics.Objects, s.Metrics.Successes)
		*statusToUpdate = *s

		wg.Done()
	}
}

func (suite *CollectionUnitTestSuite) TestCollection() {
	var (
		now = time.Now()

		stubItemID      = "fakeItemID"
		stubItemName    = "itemName"
		stubItemContent = []byte("stub_content")

		stubMetaID       = "testMetaID"
		stubMetaEntityID = "email@provider.com"
		stubMetaRoles    = []string{"read", "write"}
		stubMeta         = metadata.Metadata{
			FileName: stubItemName,
			Permissions: []metadata.Permission{
				{
					ID:         stubMetaID,
					EntityID:   stubMetaEntityID,
					EntityType: metadata.GV2User,
					Roles:      stubMetaRoles,
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
		expectErr    require.ErrorAssertionFunc
		expectLabels []string
	}{
		{
			name:         "oneDrive, no duplicates",
			numInstances: 1,
			service:      path.OneDriveService,
			itemDeets:    nst{stubItemName, 42, now},
			itemInfo:     details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: stubItemName, Modified: now}},
			getBody:      io.NopCloser(bytes.NewReader(stubItemContent)),
			getErr:       nil,
			expectErr:    require.NoError,
		},
		{
			name:         "oneDrive, duplicates",
			numInstances: 3,
			service:      path.OneDriveService,
			itemDeets:    nst{stubItemName, 42, now},
			getBody:      io.NopCloser(bytes.NewReader(stubItemContent)),
			getErr:       nil,
			itemInfo:     details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: stubItemName, Modified: now}},
			expectErr:    require.NoError,
		},
		{
			name:         "oneDrive, malware",
			numInstances: 3,
			service:      path.OneDriveService,
			itemDeets:    nst{stubItemName, 42, now},
			itemInfo:     details.ItemInfo{},
			getBody:      nil,
			getErr:       clues.New("test malware").Label(graph.LabelsMalware),
			expectErr:    require.Error,
			expectLabels: []string{graph.LabelsMalware, graph.LabelsSkippable},
		},
		{
			name:         "oneDrive, not found",
			numInstances: 3,
			service:      path.OneDriveService,
			itemDeets:    nst{stubItemName, 42, now},
			itemInfo:     details.ItemInfo{},
			getBody:      nil,
			getErr:       clues.New("test not found").Label(graph.LabelStatus(http.StatusNotFound)),
			expectErr:    require.Error,
			expectLabels: []string{graph.LabelStatus(http.StatusNotFound), graph.LabelsSkippable},
		},
		{
			name:         "sharePoint, no duplicates",
			numInstances: 1,
			service:      path.SharePointService,
			itemDeets:    nst{stubItemName, 42, now},
			itemInfo:     details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: stubItemName, Modified: now}},
			getBody:      io.NopCloser(bytes.NewReader(stubItemContent)),
			getErr:       nil,
			expectErr:    require.NoError,
		},
		{
			name:         "sharePoint, duplicates",
			numInstances: 3,
			service:      path.SharePointService,
			itemDeets:    nst{stubItemName, 42, now},
			itemInfo:     details.ItemInfo{SharePoint: &details.SharePointInfo{ItemName: stubItemName, Modified: now}},
			getBody:      io.NopCloser(bytes.NewReader(stubItemContent)),
			getErr:       nil,
			expectErr:    require.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				wg         = sync.WaitGroup{}
				collStatus = support.ControllerOperationStatus{}
				readItems  = []data.Item{}
			)

			pb := path.Builder{}.Append(path.Split("drive/driveID1/root:/dir1/dir2/dir3")...)

			folderPath, err := pb.ToDataLayerOneDrivePath("tenant", "owner", false)
			require.NoError(t, err, clues.ToCore(err))

			mbh := mock.DefaultOneDriveBH("a-user")
			if test.service == path.SharePointService {
				mbh = mock.DefaultSharePointBH("a-site")
				mbh.ItemInfo.SharePoint.Modified = now
				mbh.ItemInfo.SharePoint.ItemName = stubItemName
			} else {
				mbh.ItemInfo.OneDrive.Modified = now
				mbh.ItemInfo.OneDrive.ItemName = stubItemName
			}

			mbh.GetResps = []*http.Response{
				{
					StatusCode: http.StatusOK,
					Body:       test.getBody,
				},
			}
			mbh.GetErrs = []error{test.getErr}
			mbh.GI = mock.GetsItem{Err: assert.AnError}

			pcr := metaTD.NewStubPermissionResponse(metadata.GV2User, stubMetaID, stubMetaEntityID, stubMetaRoles)
			mbh.GIP = mock.GetsItemPermission{Perm: pcr}

			coll, err := NewCollection(
				mbh,
				folderPath,
				nil,
				"drive-id",
				suite.testStatusUpdater(&wg, &collStatus),
				control.Options{ToggleFeatures: control.Toggles{}},
				CollectionScopeFolder,
				true,
				nil)
			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, coll)
			assert.Equal(t, folderPath, coll.FullPath())

			stubItem := odTD.NewStubDriveItem(
				stubItemID,
				test.itemDeets.name,
				test.itemDeets.size,
				test.itemDeets.time,
				test.itemDeets.time,
				true,
				true)

			for i := 0; i < test.numInstances; i++ {
				coll.Add(stubItem)
			}

			// Read items from the collection
			// only needs 1 because multiple items should get deduped.
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

			assert.Equal(t, stubItemID+metadata.DataFileSuffix, readItem.ID())
			require.Implements(t, (*data.ItemModTime)(nil), readItem)

			mt := readItem.(data.ItemModTime)
			assert.Equal(t, now, mt.ModTime())

			readData, err := io.ReadAll(readItem.ToReader())
			test.expectErr(t, err)

			if err != nil {
				for _, label := range test.expectLabels {
					assert.Truef(t, clues.HasLabel(err, label), "has clues label: %s", label)
				}

				return
			}

			assert.Equal(t, stubItemContent, readData)

			readItemMeta := readItems[1]
			assert.Equal(t, stubItemID+metadata.MetaFileSuffix, readItemMeta.ID())

			readMeta := metadata.Metadata{}
			err = json.NewDecoder(readItemMeta.ToReader()).Decode(&readMeta)
			require.NoError(t, err, clues.ToCore(err))

			metaTD.AssertMetadataEqual(t, stubMeta, readMeta)
		})
	}
}

func (suite *CollectionUnitTestSuite) TestCollectionReadError() {
	var (
		t                = suite.T()
		stubItemID       = "fakeItemID"
		collStatus       = support.ControllerOperationStatus{}
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

	mbh := mock.DefaultOneDriveBH("a-user")
	mbh.GI = mock.GetsItem{Err: assert.AnError}
	mbh.GIP = mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()}
	mbh.GetResps = []*http.Response{
		nil,
		{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("test"))},
	}
	mbh.GetErrs = []error{
		clues.Stack(assert.AnError).Label(graph.LabelStatus(http.StatusUnauthorized)),
		nil,
	}

	coll, err := NewCollection(
		mbh,
		folderPath,
		nil,
		"fakeDriveID",
		suite.testStatusUpdater(&wg, &collStatus),
		control.Options{ToggleFeatures: control.Toggles{}},
		CollectionScopeFolder,
		true,
		nil)
	require.NoError(t, err, clues.ToCore(err))

	stubItem := odTD.NewStubDriveItem(
		stubItemID,
		name,
		size,
		now,
		now,
		true,
		false)

	coll.Add(stubItem)

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
		stubItemID       = "fakeItemID"
		collStatus       = support.ControllerOperationStatus{}
		wg               = sync.WaitGroup{}
		name             = "name"
		size       int64 = 42
		now              = time.Now()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	wg.Add(1)

	stubItem := odTD.NewStubDriveItem(
		stubItemID,
		name,
		size,
		now,
		now,
		true,
		false)

	pb := path.Builder{}.Append(path.Split("drive/driveID1/root:/folderPath")...)
	folderPath, err := pb.ToDataLayerOneDrivePath("a-tenant", "a-user", false)
	require.NoError(t, err)

	mbh := mock.DefaultOneDriveBH("a-user")
	mbh.GI = mock.GetsItem{Item: stubItem}
	mbh.GIP = mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()}
	mbh.GetResps = []*http.Response{
		nil,
		{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("test"))},
	}
	mbh.GetErrs = []error{
		clues.Stack(assert.AnError).Label(graph.LabelStatus(http.StatusUnauthorized)),
		nil,
	}

	coll, err := NewCollection(
		mbh,
		folderPath,
		nil,
		"fakeDriveID",
		suite.testStatusUpdater(&wg, &collStatus),
		control.Options{ToggleFeatures: control.Toggles{}},
		CollectionScopeFolder,
		true,
		nil)
	require.NoError(t, err, clues.ToCore(err))

	coll.Add(stubItem)

	collItem, ok := <-coll.Items(ctx, fault.New(true))
	assert.True(t, ok)

	_, err = io.ReadAll(collItem.ToReader())
	assert.NoError(t, err, clues.ToCore(err))

	wg.Wait()

	require.Equal(t, collStatus.Metrics.Objects, 1, "only one object should be counted")
	require.Equal(t, collStatus.Metrics.Successes, 1, "read object successfully")
}

// Ensure metadata file always uses latest time for mod time
func (suite *CollectionUnitTestSuite) TestCollectionPermissionBackupLatestModTime() {
	var (
		t            = suite.T()
		stubItemID   = "fakeItemID"
		stubItemName = "Fake Item"
		stubItemSize = int64(10)
		collStatus   = support.ControllerOperationStatus{}
		wg           = sync.WaitGroup{}
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	wg.Add(1)

	pb := path.Builder{}.Append(path.Split("drive/driveID1/root:/folderPath")...)
	folderPath, err := pb.ToDataLayerOneDrivePath("a-tenant", "a-user", false)
	require.NoError(t, err, clues.ToCore(err))

	mbh := mock.DefaultOneDriveBH("a-user")
	mbh.ItemInfo = details.ItemInfo{OneDrive: &details.OneDriveInfo{ItemName: "fakeName", Modified: time.Now()}}
	mbh.GIP = mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()}
	mbh.GetResps = []*http.Response{{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("Fake Data!")),
	}}
	mbh.GetErrs = []error{nil}

	coll, err := NewCollection(
		mbh,
		folderPath,
		nil,
		"drive-id",
		suite.testStatusUpdater(&wg, &collStatus),
		control.Options{ToggleFeatures: control.Toggles{}},
		CollectionScopeFolder,
		true,
		nil)
	require.NoError(t, err, clues.ToCore(err))

	mtime := time.Now().AddDate(0, -1, 0)

	stubItem := odTD.NewStubDriveItem(
		stubItemID,
		stubItemName,
		stubItemSize,
		mtime,
		mtime,
		true,
		false)

	coll.Add(stubItem)

	coll.handler = mbh

	readItems := []data.Item{}
	for item := range coll.Items(ctx, fault.New(true)) {
		readItems = append(readItems, item)
	}

	wg.Wait()

	// Expect no items
	require.Equal(t, 1, collStatus.Metrics.Objects)
	require.Equal(t, 1, collStatus.Metrics.Successes)

	for _, i := range readItems {
		if strings.HasSuffix(i.ID(), metadata.MetaFileSuffix) {
			content, err := io.ReadAll(i.ToReader())
			require.NoError(t, err, clues.ToCore(err))
			require.Equal(t, `{"filename":"Fake Item","permissionMode":1}`, string(content))

			im, ok := i.(data.ItemModTime)
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
				col  = &Collection{scope: test.colScope}
				now  = time.Now()
			)

			stubItem := odTD.NewStubDriveItem(
				strval,
				strval,
				test.itemSize,
				now,
				now,
				true,
				false)

			mbh := mock.DefaultOneDriveBH("a-user")
			mbh.GI = mock.GetsItem{Item: stubItem}
			mbh.GetResps = []*http.Response{{StatusCode: http.StatusOK}}
			mbh.GetErrs = []error{test.err}

			col.handler = mbh

			_, err := col.getDriveItemContent(ctx, "driveID", stubItem, errs)
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

var _ getItemPropertyer = &mockURLCache{}

type mockURLCache struct {
	Get func(ctx context.Context, itemID string) (itemProps, error)
}

func (muc *mockURLCache) getItemProperties(
	ctx context.Context,
	itemID string,
) (itemProps, error) {
	return muc.Get(ctx, itemID)
}

func (suite *GetDriveItemUnitTestSuite) TestDownloadContent() {
	var (
		driveID   string
		iorc      = io.NopCloser(bytes.NewReader([]byte("fnords")))
		item      = odTD.NewStubDriveItem("id", "n", 1, time.Now(), time.Now(), true, false)
		itemWID   = odTD.NewStubDriveItem("id", "n", 1, time.Now(), time.Now(), true, false)
		errUnauth = clues.Stack(assert.AnError).Label(graph.LabelStatus(http.StatusUnauthorized))
	)

	itemWID.SetId(ptr.To("brainhooldy"))

	m := &mockURLCache{
		Get: func(ctx context.Context, itemID string) (itemProps, error) {
			return itemProps{}, clues.Stack(assert.AnError)
		},
	}

	table := []struct {
		name      string
		mgi       mock.GetsItem
		itemInfo  details.ItemInfo
		respBody  []io.ReadCloser
		getErr    []error
		expectErr require.ErrorAssertionFunc
		expect    require.ValueAssertionFunc
		muc       *mockURLCache
	}{
		{
			name:      "good",
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{iorc},
			getErr:    []error{nil},
			expectErr: require.NoError,
			expect:    require.NotNil,
			muc:       m,
		},
		{
			name:      "expired url redownloads",
			mgi:       mock.GetsItem{Item: itemWID, Err: nil},
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil, iorc},
			getErr:    []error{errUnauth, nil},
			expectErr: require.NoError,
			expect:    require.NotNil,
			muc:       m,
		},
		{
			name:      "immediate error",
			itemInfo:  details.ItemInfo{},
			getErr:    []error{assert.AnError},
			expectErr: require.Error,
			expect:    require.Nil,
			muc:       m,
		},
		{
			name:      "re-fetching the item fails",
			itemInfo:  details.ItemInfo{},
			getErr:    []error{errUnauth},
			mgi:       mock.GetsItem{Item: nil, Err: assert.AnError},
			expectErr: require.Error,
			expect:    require.Nil,
			muc:       m,
		},
		{
			name:      "expired url fails redownload",
			mgi:       mock.GetsItem{Item: itemWID, Err: nil},
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil, nil},
			getErr:    []error{errUnauth, assert.AnError},
			expectErr: require.Error,
			expect:    require.Nil,
			muc:       m,
		},
		{
			name:      "url refreshed from cache",
			mgi:       mock.GetsItem{Item: itemWID, Err: nil},
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil, iorc},
			getErr:    []error{errUnauth, nil},
			expectErr: require.NoError,
			expect:    require.NotNil,
			muc: &mockURLCache{
				Get: func(ctx context.Context, itemID string) (itemProps, error) {
					return itemProps{
							downloadURL: "http://example.com",
							isDeleted:   false,
						},
						nil
				},
			},
		},
		{
			name:      "url refreshed from cache but item deleted",
			mgi:       mock.GetsItem{Item: itemWID, Err: graph.ErrDeletedInFlight},
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil, nil, nil},
			getErr:    []error{errUnauth, graph.ErrDeletedInFlight, graph.ErrDeletedInFlight},
			expectErr: require.Error,
			expect:    require.Nil,
			muc: &mockURLCache{
				Get: func(ctx context.Context, itemID string) (itemProps, error) {
					return itemProps{
							downloadURL: "http://example.com",
							isDeleted:   true,
						},
						nil
				},
			},
		},
		{
			name:      "fallback to item fetch on any cache error",
			mgi:       mock.GetsItem{Item: itemWID, Err: nil},
			itemInfo:  details.ItemInfo{},
			respBody:  []io.ReadCloser{nil, iorc},
			getErr:    []error{errUnauth, nil},
			expectErr: require.NoError,
			expect:    require.NotNil,
			muc: &mockURLCache{
				Get: func(ctx context.Context, itemID string) (itemProps, error) {
					return itemProps{}, assert.AnError
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			resps := make([]*http.Response, 0, len(test.respBody))

			for _, v := range test.respBody {
				if v == nil {
					resps = append(resps, nil)
				} else {
					resps = append(resps, &http.Response{StatusCode: http.StatusOK, Body: v})
				}
			}

			mbh := mock.DefaultOneDriveBH("a-user")
			mbh.GI = test.mgi
			mbh.ItemInfo = test.itemInfo
			mbh.GetResps = resps
			mbh.GetErrs = test.getErr

			r, err := downloadContent(ctx, mbh, test.muc, item, driveID)
			test.expect(t, r)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}

func (suite *CollectionUnitTestSuite) TestItemExtensions() {
	type verifyExtensionOutput func(
		t *testing.T,
		info details.ItemInfo,
		payload []byte,
	)

	var (
		t            = suite.T()
		stubItemID   = "itemID"
		stubItemName = "name"
		driveID      = "driveID"
		collStatus   = support.ControllerOperationStatus{}
		wg           = sync.WaitGroup{}
		now          = time.Now()
		readData     = []byte("hello world!")
		pb           = path.Builder{}.Append(path.Split("drive/driveID1/root:/folderPath")...)
	)

	folderPath, err := pb.ToDataLayerOneDrivePath("a-tenant", "a-user", false)
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name           string
		factories      []extensions.CreateItemExtensioner
		payload        []byte
		expectReadErr  require.ErrorAssertionFunc
		expectCloseErr require.ErrorAssertionFunc
		rc             io.ReadCloser
		expect         verifyExtensionOutput
	}{
		{
			name:           "nil extensions",
			factories:      nil,
			payload:        readData,
			expectReadErr:  require.NoError,
			expectCloseErr: require.NoError,
			rc:             io.NopCloser(bytes.NewReader(readData)),
			expect: func(
				t *testing.T,
				info details.ItemInfo,
				payload []byte,
			) {
				require.Nil(t, info.Extension.Data)
			},
		},
		{
			name:           "no extensions",
			factories:      []extensions.CreateItemExtensioner{},
			payload:        readData,
			expectReadErr:  require.NoError,
			expectCloseErr: require.NoError,
			rc:             io.NopCloser(bytes.NewReader(readData)),
			expect: func(
				t *testing.T,
				info details.ItemInfo,
				payload []byte,
			) {
				require.Nil(t, info.Extension.Data)
			},
		},
		{
			name: "with extension",
			factories: []extensions.CreateItemExtensioner{
				&extensions.MockItemExtensionFactory{},
			},
			payload:        readData,
			expectReadErr:  require.NoError,
			expectCloseErr: require.NoError,
			rc:             io.NopCloser(bytes.NewReader(readData)),
			expect: func(
				t *testing.T,
				info details.ItemInfo,
				payload []byte,
			) {
				verifyExtensionData(
					t,
					info.Extension,
					int64(len(payload)),
					crc32.ChecksumIEEE(payload))
			},
		},
		{
			name: "zero length payload",
			factories: []extensions.CreateItemExtensioner{
				&extensions.MockItemExtensionFactory{},
			},
			payload:        []byte{},
			expectReadErr:  require.NoError,
			expectCloseErr: require.NoError,
			rc:             io.NopCloser(bytes.NewReader([]byte{})),
			expect: func(
				t *testing.T,
				info details.ItemInfo,
				payload []byte,
			) {
				verifyExtensionData(
					t,
					info.Extension,
					int64(len(payload)),
					crc32.ChecksumIEEE(payload))
			},
		},
		{
			name: "extension fails on read",
			factories: []extensions.CreateItemExtensioner{
				&extensions.MockItemExtensionFactory{
					FailOnRead: true,
				},
			},
			payload:        readData,
			expectReadErr:  require.Error,
			expectCloseErr: require.NoError,
			rc:             io.NopCloser(bytes.NewReader(readData)),
			expect: func(
				t *testing.T,
				info details.ItemInfo,
				payload []byte,
			) {
				// The extension may have dirty data in this case, hence skipping
				// verification of extension info
			},
		},
		{
			name: "extension fails on close",
			factories: []extensions.CreateItemExtensioner{
				&extensions.MockItemExtensionFactory{
					FailOnClose: true,
				},
			},
			payload:        readData,
			expectReadErr:  require.NoError,
			expectCloseErr: require.Error,
			rc:             io.NopCloser(bytes.NewReader(readData)),
			expect: func(
				t *testing.T,
				info details.ItemInfo,
				payload []byte,
			) {
				// The extension may have dirty data in this case, hence skipping
				// verification of extension info
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)
			defer flush()

			wg.Add(1)

			mbh := mock.DefaultOneDriveBH("a-user")
			mbh.GI = mock.GetsItem{Err: assert.AnError}
			mbh.GIP = mock.GetsItemPermission{Perm: models.NewPermissionCollectionResponse()}
			mbh.GetResps = []*http.Response{
				{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(string(test.payload))),
				},
			}
			mbh.GetErrs = []error{
				nil,
			}

			opts := control.DefaultOptions()
			opts.ItemExtensionFactory = append(
				opts.ItemExtensionFactory,
				test.factories...)

			coll, err := NewCollection(
				mbh,
				folderPath,
				nil,
				driveID,
				suite.testStatusUpdater(&wg, &collStatus),
				opts,
				CollectionScopeFolder,
				true,
				nil)
			require.NoError(t, err, clues.ToCore(err))

			stubItem := odTD.NewStubDriveItem(
				stubItemID,
				stubItemName,
				int64(len(test.payload)),
				now,
				now,
				true,
				false)

			coll.Add(stubItem)

			collItem, ok := <-coll.Items(ctx, fault.New(true))
			assert.True(t, ok)

			wg.Wait()

			ei, ok := collItem.(data.ItemInfo)
			assert.True(t, ok)
			itemInfo := ei.Info()

			_, err = io.ReadAll(collItem.ToReader())
			test.expectReadErr(t, err, clues.ToCore(err))

			err = collItem.ToReader().Close()
			test.expectCloseErr(t, err, clues.ToCore(err))

			// Verify extension data
			test.expect(t, itemInfo, test.payload)
		})
	}
}

func verifyExtensionData(
	t *testing.T,
	extensionData *details.ExtensionData,
	expectedBytes int64,
	expectedCrc uint32,
) {
	require.NotNil(t, extensionData, "nil extension")
	assert.NotNil(t, extensionData.Data[extensions.KNumBytes], "key not found")
	assert.NotNil(t, extensionData.Data[extensions.KCrc32], "key not found")

	eSize := extensionData.Data[extensions.KNumBytes].(int64)
	assert.Equal(t, expectedBytes, eSize, "incorrect num bytes")

	c := extensionData.Data[extensions.KCrc32].(uint32)
	require.Equal(t, expectedCrc, c, "incorrect crc")
}
