package onedrive

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/onedrive/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

type RestoreUnitSuite struct {
	tester.Suite
}

func TestRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreUnitSuite) TestAugmentRestorePaths() {
	// Adding a simple test here so that we can be sure that this
	// function gets updated whenever we add a new version.
	require.LessOrEqual(suite.T(), version.Backup, version.All8MigrateUserPNToID, "unsupported backup version")

	table := []struct {
		name    string
		version int
		input   []string
		output  []string
	}{
		{
			name:    "no change v0",
			version: 0,
			input: []string{
				"file.txt.data",
				"file.txt", // v0 does not have `.data`
			},
			output: []string{
				"file.txt", // ordering artifact of sorting
				"file.txt.data",
			},
		},
		{
			name:    "one folder v0",
			version: 0,
			input: []string{
				"folder/file.txt.data",
				"folder/file.txt",
			},
			output: []string{
				"folder/file.txt",
				"folder/file.txt.data",
			},
		},
		{
			name:    "no change v1",
			version: version.OneDrive1DataAndMetaFiles,
			input: []string{
				"file.txt.data",
			},
			output: []string{
				"file.txt.data",
			},
		},
		{
			name:    "one folder v1",
			version: version.OneDrive1DataAndMetaFiles,
			input: []string{
				"folder/file.txt.data",
			},
			output: []string{
				"folder.dirmeta",
				"folder/file.txt.data",
			},
		},
		{
			name:    "nested folders v1",
			version: version.OneDrive1DataAndMetaFiles,
			input: []string{
				"folder/file.txt.data",
				"folder/folder2/file.txt.data",
			},
			output: []string{
				"folder.dirmeta",
				"folder/file.txt.data",
				"folder/folder2.dirmeta",
				"folder/folder2/file.txt.data",
			},
		},
		{
			name:    "no change v4",
			version: version.OneDrive4DirIncludesPermissions,
			input: []string{
				"file.txt.data",
			},
			output: []string{
				"file.txt.data",
			},
		},
		{
			name:    "one folder v4",
			version: version.OneDrive4DirIncludesPermissions,
			input: []string{
				"folder/file.txt.data",
			},
			output: []string{
				"folder/file.txt.data",
				"folder/folder.dirmeta",
			},
		},
		{
			name:    "nested folders v4",
			version: version.OneDrive4DirIncludesPermissions,
			input: []string{
				"folder/file.txt.data",
				"folder/folder2/file.txt.data",
			},
			output: []string{
				"folder/file.txt.data",
				"folder/folder.dirmeta",
				"folder/folder2/file.txt.data",
				"folder/folder2/folder2.dirmeta",
			},
		},
		{
			name:    "no change v6",
			version: version.OneDrive6NameInMeta,
			input: []string{
				"file.txt.data",
			},
			output: []string{
				"file.txt.data",
			},
		},
		{
			name:    "one folder v6",
			version: version.OneDrive6NameInMeta,
			input: []string{
				"folder/file.txt.data",
			},
			output: []string{
				"folder/.dirmeta",
				"folder/file.txt.data",
			},
		},
		{
			name:    "nested folders v6",
			version: version.OneDrive6NameInMeta,
			input: []string{
				"folder/file.txt.data",
				"folder/folder2/file.txt.data",
			},
			output: []string{
				"folder/.dirmeta",
				"folder/file.txt.data",
				"folder/folder2/.dirmeta",
				"folder/folder2/file.txt.data",
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			base := "id/onedrive/user/files/drives/driveID/root:/"

			inPaths := []path.RestorePaths{}
			for _, ps := range test.input {
				p, err := path.FromDataLayerPath(base+ps, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				pd, err := p.Dir()
				require.NoError(t, err, "creating collection path", clues.ToCore(err))

				inPaths = append(
					inPaths,
					path.RestorePaths{StoragePath: p, RestorePath: pd})
			}

			outPaths := []path.RestorePaths{}
			for _, ps := range test.output {
				p, err := path.FromDataLayerPath(base+ps, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				pd, err := p.Dir()
				require.NoError(t, err, "creating collection path", clues.ToCore(err))

				outPaths = append(
					outPaths,
					path.RestorePaths{StoragePath: p, RestorePath: pd})
			}

			actual, err := AugmentRestorePaths(test.version, inPaths)
			require.NoError(t, err, "augmenting paths", clues.ToCore(err))

			// Ordering of paths matter here as we need dirmeta files
			// to show up before file in dir
			assert.Equal(t, outPaths, actual, "augmented paths")
		})
	}
}

// TestAugmentRestorePaths_DifferentRestorePath tests that RestorePath
// substitution works properly. Since it's only possible for future backup
// versions to need restore path substitution (i.e. due to storing folders by
// ID instead of name) this is only tested against the most recent backup
// version at the moment.
func (suite *RestoreUnitSuite) TestAugmentRestorePaths_DifferentRestorePath() {
	// Adding a simple test here so that we can be sure that this
	// function gets updated whenever we add a new version.
	require.LessOrEqual(suite.T(), version.Backup, version.All8MigrateUserPNToID, "unsupported backup version")

	type pathPair struct {
		storage string
		restore string
	}

	table := []struct {
		name     string
		version  int
		input    []pathPair
		output   []pathPair
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name:    "nested folders",
			version: version.Backup,
			input: []pathPair{
				{storage: "folder-id/file.txt.data", restore: "folder"},
				{storage: "folder-id/folder2-id/file.txt.data", restore: "folder/folder2"},
			},
			output: []pathPair{
				{storage: "folder-id/.dirmeta", restore: "folder"},
				{storage: "folder-id/file.txt.data", restore: "folder"},
				{storage: "folder-id/folder2-id/.dirmeta", restore: "folder/folder2"},
				{storage: "folder-id/folder2-id/file.txt.data", restore: "folder/folder2"},
			},
			errCheck: assert.NoError,
		},
		{
			name:    "restore path longer one folder",
			version: version.Backup,
			input: []pathPair{
				{storage: "folder-id/file.txt.data", restore: "corso_restore/folder"},
			},
			output: []pathPair{
				{storage: "folder-id/.dirmeta", restore: "corso_restore/folder"},
				{storage: "folder-id/file.txt.data", restore: "corso_restore/folder"},
			},
			errCheck: assert.NoError,
		},
		{
			name:    "restore path shorter one folder",
			version: version.Backup,
			input: []pathPair{
				{storage: "folder-id/file.txt.data", restore: ""},
			},
			errCheck: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			base := "id/onedrive/user/files/drives/driveID/root:/"

			inPaths := []path.RestorePaths{}
			for _, ps := range test.input {
				p, err := path.FromDataLayerPath(base+ps.storage, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				r, err := path.FromDataLayerPath(base+ps.restore, false)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				inPaths = append(
					inPaths,
					path.RestorePaths{StoragePath: p, RestorePath: r})
			}

			outPaths := []path.RestorePaths{}
			for _, ps := range test.output {
				p, err := path.FromDataLayerPath(base+ps.storage, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				r, err := path.FromDataLayerPath(base+ps.restore, false)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				outPaths = append(
					outPaths,
					path.RestorePaths{StoragePath: p, RestorePath: r})
			}

			actual, err := AugmentRestorePaths(test.version, inPaths)
			test.errCheck(t, err, "augmenting paths", clues.ToCore(err))

			if err != nil {
				return
			}

			// Ordering of paths matter here as we need dirmeta files
			// to show up before file in dir
			assert.Equal(t, outPaths, actual, "augmented paths")
		})
	}
}

func (suite *RestoreUnitSuite) TestRestoreItem_collisionHandling() {
	const mndiID = "mndi-id"

	type counts struct {
		skip    int64
		replace int64
		new     int64
	}

	table := []struct {
		name          string
		collisionKeys map[string]api.DriveItemIDType
		onCollision   control.CollisionPolicy
		deleteErr     error
		expectSkipped assert.BoolAssertionFunc
		expectMock    func(*testing.T, *mock.RestoreHandler)
		expectCounts  counts
	}{
		{
			name:          "no collision, copy",
			collisionKeys: map[string]api.DriveItemIDType{},
			onCollision:   control.Copy,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:          "no collision, replace",
			collisionKeys: map[string]api.DriveItemIDType{},
			onCollision:   control.Replace,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:          "no collision, skip",
			collisionKeys: map[string]api.DriveItemIDType{},
			onCollision:   control.Skip,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name: "collision, copy",
			collisionKeys: map[string]api.DriveItemIDType{
				mock.DriveItemFileName: {ItemID: mndiID},
			},
			onCollision:   control.Copy,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name: "collision, replace",
			collisionKeys: map[string]api.DriveItemIDType{
				mock.DriveItemFileName: {ItemID: mndiID},
			},
			onCollision:   control.Replace,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.True(t, rh.CalledDeleteItem, "new item deleted")
				assert.Equal(t, mndiID, rh.CalledDeleteItemOn, "deleted the correct item")
			},
			expectCounts: counts{0, 1, 0},
		},
		{
			name: "collision, replace - err already deleted",
			collisionKeys: map[string]api.DriveItemIDType{
				mock.DriveItemFileName: {ItemID: "smarf"},
			},
			onCollision:   control.Replace,
			deleteErr:     graph.ErrDeletedInFlight,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.True(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{0, 1, 0},
		},
		{
			name: "collision, skip",
			collisionKeys: map[string]api.DriveItemIDType{
				mock.DriveItemFileName: {ItemID: mndiID},
			},
			onCollision:   control.Skip,
			expectSkipped: assert.True,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.False(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{1, 0, 0},
		},
		{
			name: "file-folder collision, copy",
			collisionKeys: map[string]api.DriveItemIDType{
				mock.DriveItemFileName: {
					ItemID:   mndiID,
					IsFolder: true,
				},
			},
			onCollision:   control.Copy,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name: "file-folder collision, replace",
			collisionKeys: map[string]api.DriveItemIDType{
				mock.DriveItemFileName: {
					ItemID:   mndiID,
					IsFolder: true,
				},
			},
			onCollision:   control.Replace,
			expectSkipped: assert.False,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.True(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name: "file-folder collision, skip",
			collisionKeys: map[string]api.DriveItemIDType{
				mock.DriveItemFileName: {
					ItemID:   mndiID,
					IsFolder: true,
				},
			},
			onCollision:   control.Skip,
			expectSkipped: assert.True,
			expectMock: func(t *testing.T, rh *mock.RestoreHandler) {
				assert.False(t, rh.CalledPostItem, "new item posted")
				assert.False(t, rh.CalledDeleteItem, "new item deleted")
			},
			expectCounts: counts{1, 0, 0},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mndi := models.NewDriveItem()
			mndi.SetId(ptr.To(mndiID))

			var (
				caches = NewRestoreCaches(nil)
				rh     = &mock.RestoreHandler{
					PostItemResp:  models.NewDriveItem(),
					DeleteItemErr: test.deleteErr,
				}
				restoreCfg = control.RestoreConfig{OnCollision: test.onCollision}
				dpb        = odConsts.DriveFolderPrefixBuilder("driveID1")
			)

			caches.collisionKeyToItemID = test.collisionKeys

			dpp, err := dpb.ToDataLayerOneDrivePath("t", "u", false)
			require.NoError(t, err)

			dp, err := path.ToDrivePath(dpp)
			require.NoError(t, err)

			ctr := count.New()

			_, skip, err := restoreItem(
				ctx,
				rh,
				restoreCfg,
				mock.FetchItemByName{
					Item: &mock.Data{
						Reader: mock.FileRespReadCloser(mock.DriveFileMetaData),
					},
				},
				version.Backup,
				dp,
				"",
				make([]byte, graph.CopyBufferSize),
				caches,
				false,
				&mock.Data{
					ID:     uuid.NewString(),
					Reader: mock.FileRespReadCloser(mock.DriveFilePayloadData),
				},
				nil,
				ctr)

			require.NoError(t, err, clues.ToCore(err))
			test.expectSkipped(t, skip)
			test.expectMock(t, rh)
			assert.Equal(t, test.expectCounts.skip, ctr.Get(count.CollisionSkip), "skips")
			assert.Equal(t, test.expectCounts.replace, ctr.Get(count.CollisionReplace), "replaces")
			assert.Equal(t, test.expectCounts.new, ctr.Get(count.NewItemCreated), "new items")
		})
	}
}

type mockPIIC struct {
	i     int
	errs  []error
	items []models.DriveItemable
}

func (m *mockPIIC) PostItemInContainer(
	context.Context,
	string, string,
	models.DriveItemable,
	control.CollisionPolicy,
) (models.DriveItemable, error) {
	j := m.i
	m.i++

	return m.items[j], m.errs[j]
}

func (suite *RestoreUnitSuite) TestCreateFolder() {
	table := []struct {
		name       string
		mock       *mockPIIC
		expectErr  assert.ErrorAssertionFunc
		expectItem assert.ValueAssertionFunc
	}{
		{
			name: "good",
			mock: &mockPIIC{
				errs:  []error{nil},
				items: []models.DriveItemable{models.NewDriveItem()},
			},
			expectErr:  assert.NoError,
			expectItem: assert.NotNil,
		},
		{
			name: "good with copy",
			mock: &mockPIIC{
				errs:  []error{graph.ErrItemAlreadyExistsConflict, nil},
				items: []models.DriveItemable{nil, models.NewDriveItem()},
			},
			expectErr:  assert.NoError,
			expectItem: assert.NotNil,
		},
		{
			name: "bad",
			mock: &mockPIIC{
				errs:  []error{assert.AnError},
				items: []models.DriveItemable{nil},
			},
			expectErr:  assert.Error,
			expectItem: assert.Nil,
		},
		{
			name: "bad with copy",
			mock: &mockPIIC{
				errs:  []error{graph.ErrItemAlreadyExistsConflict, assert.AnError},
				items: []models.DriveItemable{nil, nil},
			},
			expectErr:  assert.Error,
			expectItem: assert.Nil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := createFolder(ctx, test.mock, "d", "pf", "fn")
			test.expectErr(t, err, clues.ToCore(err))
			test.expectItem(t, result)
		})
	}
}

type mockGRF struct {
	err        error
	rootFolder models.DriveItemable
}

func (m *mockGRF) GetRootFolder(
	context.Context,
	string,
) (models.DriveItemable, error) {
	return m.rootFolder, m.err
}

func (suite *RestoreUnitSuite) TestRestoreCaches_AddDrive() {
	rfID := "this-is-id"
	driveID := "another-id"
	name := "name"

	rf := models.NewDriveItem()
	rf.SetId(&rfID)

	md := models.NewDrive()
	md.SetId(&driveID)
	md.SetName(&name)

	table := []struct {
		name        string
		mock        *mockGRF
		expectErr   require.ErrorAssertionFunc
		expectID    string
		checkValues bool
	}{
		{
			name:        "good",
			mock:        &mockGRF{rootFolder: rf},
			expectErr:   require.NoError,
			expectID:    rfID,
			checkValues: true,
		},
		{
			name:      "err",
			mock:      &mockGRF{err: assert.AnError},
			expectErr: require.Error,
			expectID:  "",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			rc := NewRestoreCaches(nil)
			err := rc.AddDrive(ctx, md, test.mock)
			test.expectErr(t, err, clues.ToCore(err))

			if test.checkValues {
				idResult := rc.DriveIDToDriveInfo[driveID]
				assert.Equal(t, driveID, idResult.id, "drive id")
				assert.Equal(t, name, idResult.name, "drive name")
				assert.Equal(t, test.expectID, idResult.rootFolderID, "root folder id")

				nameResult := rc.DriveNameToDriveInfo[name]
				assert.Equal(t, driveID, nameResult.id, "drive id")
				assert.Equal(t, name, nameResult.name, "drive name")
				assert.Equal(t, test.expectID, nameResult.rootFolderID, "root folder id")
			}
		})
	}
}

type mockGDPARF struct {
	err        error
	rootFolder models.DriveItemable
	pager      *apiMock.DrivePager
}

func (m *mockGDPARF) GetRootFolder(
	context.Context,
	string,
) (models.DriveItemable, error) {
	return m.rootFolder, m.err
}

func (m *mockGDPARF) NewDrivePager(
	string,
	[]string,
) api.DrivePager {
	return m.pager
}

func (suite *RestoreUnitSuite) TestRestoreCaches_Populate() {
	rfID := "this-is-id"
	driveID := "another-id"
	name := "name"

	rf := models.NewDriveItem()
	rf.SetId(&rfID)

	md := models.NewDrive()
	md.SetId(&driveID)
	md.SetName(&name)

	table := []struct {
		name        string
		mock        *apiMock.DrivePager
		expectErr   require.ErrorAssertionFunc
		expectLen   int
		checkValues bool
	}{
		{
			name: "no results",
			mock: &apiMock.DrivePager{
				ToReturn: []apiMock.PagerResult{
					{Drives: []models.Driveable{}},
				},
			},
			expectErr: require.NoError,
			expectLen: 0,
		},
		{
			name: "one result",
			mock: &apiMock.DrivePager{
				ToReturn: []apiMock.PagerResult{
					{Drives: []models.Driveable{md}},
				},
			},
			expectErr:   require.NoError,
			expectLen:   1,
			checkValues: true,
		},
		{
			name: "error",
			mock: &apiMock.DrivePager{
				ToReturn: []apiMock.PagerResult{
					{Err: assert.AnError},
				},
			},
			expectErr: require.Error,
			expectLen: 0,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gdparf := &mockGDPARF{
				rootFolder: rf,
				pager:      test.mock,
			}

			rc := NewRestoreCaches(nil)
			err := rc.Populate(ctx, gdparf, "shmoo")
			test.expectErr(t, err, clues.ToCore(err))

			assert.Len(t, rc.DriveIDToDriveInfo, test.expectLen)
			assert.Len(t, rc.DriveNameToDriveInfo, test.expectLen)

			if test.checkValues {
				idResult := rc.DriveIDToDriveInfo[driveID]
				assert.Equal(t, driveID, idResult.id, "drive id")
				assert.Equal(t, name, idResult.name, "drive name")
				assert.Equal(t, rfID, idResult.rootFolderID, "root folder id")

				nameResult := rc.DriveNameToDriveInfo[name]
				assert.Equal(t, driveID, nameResult.id, "drive id")
				assert.Equal(t, name, nameResult.name, "drive name")
				assert.Equal(t, rfID, nameResult.rootFolderID, "root folder id")
			}
		})
	}
}

type mockPDAGRF struct {
	i        int
	postResp []models.Driveable
	postErr  []error

	grf mockGRF
}

func (m *mockPDAGRF) PostDrive(
	ctx context.Context,
	protectedResourceID, driveName string,
) (models.Driveable, error) {
	defer func() { m.i++ }()

	md := m.postResp[m.i]
	if md != nil {
		md.SetName(&driveName)
	}

	return md, m.postErr[m.i]
}

func (m *mockPDAGRF) GetRootFolder(
	ctx context.Context,
	driveID string,
) (models.DriveItemable, error) {
	return m.grf.rootFolder, m.grf.err
}

func (suite *RestoreUnitSuite) TestEnsureDriveExists() {
	rfID := "this-is-id"
	driveID := "another-id"
	name := "name"

	rf := models.NewDriveItem()
	rf.SetId(&rfID)

	grf := mockGRF{rootFolder: rf}

	makeMD := func() models.Driveable {
		md := models.NewDrive()
		md.SetId(&driveID)
		md.SetName(&name)

		return md
	}

	dp := &path.DrivePath{
		DriveID: driveID,
		Root:    "root:",
		Folders: path.Elements{},
	}

	populatedCache := func(id string) *restoreCaches {
		rc := NewRestoreCaches(nil)
		di := driveInfo{
			id:   id,
			name: name,
		}
		rc.DriveIDToDriveInfo[id] = di
		rc.DriveNameToDriveInfo[name] = di

		return rc
	}

	table := []struct {
		name            string
		mock            *mockPDAGRF
		rc              *restoreCaches
		expectErr       require.ErrorAssertionFunc
		expectName      string
		skipValueChecks bool
	}{
		{
			name: "drive already in cache",
			mock: &mockPDAGRF{
				postResp: []models.Driveable{makeMD()},
				postErr:  []error{nil},
				grf:      grf,
			},
			rc:         populatedCache(driveID),
			expectErr:  require.NoError,
			expectName: name,
		},
		{
			name: "drive created",
			mock: &mockPDAGRF{
				postResp: []models.Driveable{makeMD()},
				postErr:  []error{nil},
				grf:      grf,
			},
			rc:         NewRestoreCaches(nil),
			expectErr:  require.NoError,
			expectName: name,
		},
		{
			name: "error creating drive",
			mock: &mockPDAGRF{
				postResp: []models.Driveable{nil},
				postErr:  []error{assert.AnError},
				grf:      grf,
			},
			rc:              NewRestoreCaches(nil),
			expectErr:       require.Error,
			expectName:      "",
			skipValueChecks: true,
		},
		{
			name: "drive name already exists",
			mock: &mockPDAGRF{
				postResp: []models.Driveable{makeMD()},
				postErr:  []error{nil},
				grf:      grf,
			},
			rc:         populatedCache("beaux"),
			expectErr:  require.NoError,
			expectName: name + " 1",
		},
		{
			name: "list with name already exists",
			mock: &mockPDAGRF{
				postResp: []models.Driveable{nil, makeMD()},
				postErr:  []error{graph.ErrItemAlreadyExistsConflict, nil},
				grf:      grf,
			},
			rc:         NewRestoreCaches(nil),
			expectErr:  require.NoError,
			expectName: name + " 1",
		},
		{
			name: "drive and list with name already exist",
			mock: &mockPDAGRF{
				postResp: []models.Driveable{nil, makeMD()},
				postErr:  []error{graph.ErrItemAlreadyExistsConflict, nil},
				grf:      grf,
			},
			rc:         populatedCache("regard"),
			expectErr:  require.NoError,
			expectName: name + " 2",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			rc := test.rc

			err := ensureDriveExists(
				ctx,
				test.mock,
				rc,
				dp,
				"prID",
				name)
			test.expectErr(t, err, clues.ToCore(err))

			if !test.skipValueChecks {
				nameResult := rc.DriveNameToDriveInfo[test.expectName]
				assert.Equal(t, test.expectName, nameResult.name, "found drive entry with expected name")
			}
		})
	}
}
