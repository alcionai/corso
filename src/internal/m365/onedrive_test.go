package m365

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/resource"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/stub"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var (
	fileName          = "test-file.txt"
	folderAName       = "folder-a"
	folderBName       = "b"
	folderNamedFolder = "folder"

	fileAData = []byte(strings.Repeat("a", 33))
	fileBData = []byte(strings.Repeat("b", 65))
	fileCData = []byte(strings.Repeat("c", 129))
	fileDData = []byte(strings.Repeat("d", 257))
	fileEData = []byte(strings.Repeat("e", 257))

	// Cannot restore owner or empty permissions and so not testing them
	writePerm = []string{"write"}
	readPerm  = []string{"read"}
)

func mustGetDefaultDriveID(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	ac api.Client,
	service path.ServiceType,
	resourceOwner string,
) string {
	var (
		err error
		d   models.Driveable
	)

	switch service {
	case path.OneDriveService:
		d, err = ac.Users().GetDefaultDrive(ctx, resourceOwner)
	case path.SharePointService:
		d, err = ac.Sites().GetDefaultDrive(ctx, resourceOwner)
	default:
		assert.FailNowf(t, "unknown service type %s", service.String())
	}

	if err != nil {
		err = graph.Wrap(ctx, err, "retrieving drive")
	}

	require.NoError(t, err, clues.ToCore(err))

	id := ptr.Val(d.GetId())
	require.NotEmpty(t, id)

	return id
}

type suiteInfo interface {
	APIClient() api.Client
	Tenant() string
	// Returns (username, user ID) for the user. These values are used for
	// permissions.
	PrimaryUser() (string, string)
	SecondaryUser() (string, string)
	TertiaryUser() (string, string)
	// ResourceOwner returns the resource owner to run the backup/restore
	// with. This can be different from the values used for permissions and it can
	// also be a site.
	ResourceOwner() string
	Service() path.ServiceType
	Resource() resource.Category
}

type oneDriveSuite interface {
	tester.Suite
	suiteInfo
}

type suiteInfoImpl struct {
	ac               api.Client
	controller       *Controller
	resourceOwner    string
	resourceCategory resource.Category
	secondaryUser    string
	secondaryUserID  string
	service          path.ServiceType
	tertiaryUser     string
	tertiaryUserID   string
	user             string
	userID           string
}

func NewSuiteInfoImpl(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	resourceOwner string,
	service path.ServiceType,
) suiteInfoImpl {
	rsc := resource.Users
	if service == path.SharePointService {
		rsc = resource.Sites
	}

	ctrl := newController(ctx, t, rsc, path.OneDriveService)

	return suiteInfoImpl{
		ac:               ctrl.AC,
		controller:       ctrl,
		resourceOwner:    resourceOwner,
		resourceCategory: rsc,
		secondaryUser:    tconfig.SecondaryM365UserID(t),
		service:          service,
		tertiaryUser:     tconfig.TertiaryM365UserID(t),
		user:             tconfig.M365UserID(t),
	}
}

func (si suiteInfoImpl) APIClient() api.Client {
	return si.ac
}

func (si suiteInfoImpl) Tenant() string {
	return si.controller.tenant
}

func (si suiteInfoImpl) PrimaryUser() (string, string) {
	return si.user, si.userID
}

func (si suiteInfoImpl) SecondaryUser() (string, string) {
	return si.secondaryUser, si.secondaryUserID
}

func (si suiteInfoImpl) TertiaryUser() (string, string) {
	return si.tertiaryUser, si.tertiaryUserID
}

func (si suiteInfoImpl) ResourceOwner() string {
	return si.resourceOwner
}

func (si suiteInfoImpl) Service() path.ServiceType {
	return si.service
}

func (si suiteInfoImpl) Resource() resource.Category {
	return si.resourceCategory
}

// ---------------------------------------------------------------------------
// SharePoint Libraries
// ---------------------------------------------------------------------------
// SharePoint shares most of its libraries implementation with OneDrive so we
// only test simple things here and leave the more extensive testing to
// OneDrive.

type SharePointIntegrationSuite struct {
	tester.Suite
	suiteInfo
}

func TestSharePointIntegrationSuite(t *testing.T) {
	suite.Run(t, &SharePointIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SharePointIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	si := NewSuiteInfoImpl(suite.T(), ctx, tconfig.M365SiteID(suite.T()), path.SharePointService)

	// users needed for permissions
	user, err := si.controller.AC.Users().GetByID(ctx, si.user)
	require.NoError(t, err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.controller.AC.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(t, err, "fetching user", si.secondaryUser, clues.ToCore(err))
	si.secondaryUserID = ptr.Val(secondaryUser.GetId())

	tertiaryUser, err := si.controller.AC.Users().GetByID(ctx, si.tertiaryUser)
	require.NoError(t, err, "fetching user", si.tertiaryUser, clues.ToCore(err))
	si.tertiaryUserID = ptr.Val(tertiaryUser.GetId())

	suite.suiteInfo = si
}

func (suite *SharePointIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, version.Backup)
}

// TODO: Re-enable these tests (disabled as it currently acting up CI)
func (suite *SharePointIntegrationSuite) TestPermissionsRestoreAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testPermissionsRestoreAndBackup(suite, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestRestoreNoPermissionsAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testRestoreNoPermissionsAndBackup(suite, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestPermissionsInheritanceRestoreAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testPermissionsInheritanceRestoreAndBackup(suite, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestLinkSharesInheritanceRestoreAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testLinkSharesInheritanceRestoreAndBackup(suite, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, version.All8MigrateUserPNToID)
}

// ---------------------------------------------------------------------------
// OneDrive most recent backup version
// ---------------------------------------------------------------------------
type OneDriveIntegrationSuite struct {
	tester.Suite
	suiteInfo
}

func TestOneDriveIntegrationSuite(t *testing.T) {
	suite.Run(t, &OneDriveIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *OneDriveIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	si := NewSuiteInfoImpl(t, ctx, tconfig.M365UserID(t), path.OneDriveService)

	user, err := si.controller.AC.Users().GetByID(ctx, si.user)
	require.NoError(t, err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.controller.AC.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(t, err, "fetching user", si.secondaryUser, clues.ToCore(err))
	si.secondaryUserID = ptr.Val(secondaryUser.GetId())

	tertiaryUser, err := si.controller.AC.Users().GetByID(ctx, si.tertiaryUser)
	require.NoError(t, err, "fetching user", si.tertiaryUser, clues.ToCore(err))
	si.tertiaryUserID = ptr.Val(tertiaryUser.GetId())

	suite.suiteInfo = si
}

func (suite *OneDriveIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestPermissionsRestoreAndBackup() {
	testPermissionsRestoreAndBackup(suite, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestRestoreNoPermissionsAndBackup() {
	testRestoreNoPermissionsAndBackup(suite, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestPermissionsInheritanceRestoreAndBackup() {
	testPermissionsInheritanceRestoreAndBackup(suite, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestLinkSharesInheritanceRestoreAndBackup() {
	testLinkSharesInheritanceRestoreAndBackup(suite, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, version.All8MigrateUserPNToID)
}

// ---------------------------------------------------------------------------
// OneDrive regression
// ---------------------------------------------------------------------------
type OneDriveNightlySuite struct {
	tester.Suite
	suiteInfo
}

func TestOneDriveNightlySuite(t *testing.T) {
	suite.Run(t, &OneDriveNightlySuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *OneDriveNightlySuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	si := NewSuiteInfoImpl(t, ctx, tconfig.M365UserID(t), path.OneDriveService)

	user, err := si.controller.AC.Users().GetByID(ctx, si.user)
	require.NoError(t, err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.controller.AC.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(t, err, "fetching user", si.secondaryUser, clues.ToCore(err))
	si.secondaryUserID = ptr.Val(secondaryUser.GetId())

	tertiaryUser, err := si.controller.AC.Users().GetByID(ctx, si.tertiaryUser)
	require.NoError(t, err, "fetching user", si.tertiaryUser, clues.ToCore(err))
	si.tertiaryUserID = ptr.Val(tertiaryUser.GetId())

	suite.suiteInfo = si
}

func (suite *OneDriveNightlySuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, 0)
}

func (suite *OneDriveNightlySuite) TestPermissionsRestoreAndBackup() {
	testPermissionsRestoreAndBackup(suite, version.OneDrive1DataAndMetaFiles)
}

func (suite *OneDriveNightlySuite) TestRestoreNoPermissionsAndBackup() {
	testRestoreNoPermissionsAndBackup(suite, version.OneDrive1DataAndMetaFiles)
}

func (suite *OneDriveNightlySuite) TestPermissionsInheritanceRestoreAndBackup() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testPermissionsInheritanceRestoreAndBackup(suite, version.OneDrive4DirIncludesPermissions)
}

func (suite *OneDriveNightlySuite) TestLinkSharesInheritanceRestoreAndBackup() {
	testLinkSharesInheritanceRestoreAndBackup(suite, version.Backup)
}

func (suite *OneDriveNightlySuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, version.All8MigrateUserPNToID)
}

func testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(
	suite oneDriveSuite,
	startVersion int,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.APIClient(),
		suite.Service(),
		suite.ResourceOwner())

	rootPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
	}
	folderAPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
	}
	subfolderBPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderBName,
	}
	subfolderAPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderBName,
		folderAName,
	}
	folderBPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderBName,
	}

	defaultMetadata := stub.MetaData{SharingMode: metadata.SharingModeInherited}

	cols := []stub.ColInfo{
		{
			PathElements: rootPath,
			Meta:         defaultMetadata,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileAData,
					Meta: defaultMetadata,
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderAName,
				},
				{
					Name: folderBName,
				},
			},
		},
		{
			PathElements: folderAPath,
			Meta:         defaultMetadata,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileBData,
					Meta: defaultMetadata,
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderBName,
				},
			},
		},
		{
			PathElements: subfolderBPath,
			Meta:         defaultMetadata,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileCData,
					Meta: defaultMetadata,
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderAName,
				},
			},
		},
		{
			PathElements: subfolderAPath,
			Meta:         defaultMetadata,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileDData,
					Meta: defaultMetadata,
				},
			},
		},
		{
			PathElements: folderBPath,
			Meta:         defaultMetadata,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileEData,
					Meta: defaultMetadata,
				},
			},
		},
	}

	expected, err := stub.DataForInfo(suite.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input, err := stub.DataForInfo(suite.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.Service(),
				resourceCat:         suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("od_restore_and_backup_multi")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			runRestoreBackupTestVersions(
				t,
				testData,
				suite.Tenant(),
				[]string{suite.ResourceOwner()},
				control.DefaultOptions(),
				restoreCfg)
		})
	}
}

func testPermissionsRestoreAndBackup(suite oneDriveSuite, startVersion int) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.APIClient(),
		suite.Service(),
		suite.ResourceOwner())

	fileName2 := "test-file2.txt"
	folderCName := "folder-c"

	rootPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
	}
	folderAPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
	}
	folderBPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderBName,
	}
	// For skipped test
	// subfolderAPath := []string{
	// 	odConsts.DrivesPathDir,
	// 	driveID,
	// 	odConsts.RootPathDir,
	// 	folderBName,
	// 	folderAName,
	// }
	folderCPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderCName,
	}

	cols := []stub.ColInfo{
		{
			PathElements: rootPath,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeInherited,
			},
			Files: []stub.ItemData{
				{
					// Test restoring a file that doesn't inherit permissions.
					Name: fileName,
					Data: fileAData,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
				},
				{
					// Test restoring a file that doesn't inherit permissions and has
					// no permissions.
					Name: fileName2,
					Data: fileBData,
					Meta: stub.MetaData{
						SharingMode: metadata.SharingModeInherited,
					},
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderBName,
					Meta: stub.MetaData{
						SharingMode: metadata.SharingModeInherited,
					},
				},
				{
					Name: folderAName,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    readPerm,
						},
					},
				},
				{
					Name: folderCName,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    readPerm,
						},
					},
				},
			},
		},
		{
			PathElements: folderBPath,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeInherited,
			},
			Files: []stub.ItemData{
				{
					// Test restoring a file in a non-root folder that doesn't inherit
					// permissions.
					Name: fileName,
					Data: fileBData,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderAName,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    readPerm,
						},
					},
				},
			},
		},
		// TODO: We can't currently support having custom permissions
		// with the same set of permissions internally
		// {
		// 	// Tests a folder that has permissions with an item in the folder with
		// 	// the same permissions.
		// 	pathElements: subfolderAPath,
		// 	files: []stub.ItemData{
		// 		{
		// 			name: fileName,
		// 			data: fileDData,
		// 			perms: stub.PermData{
		// 				user:     secondaryUserName,
		// 				entityID: secondaryUserID,
		// 				roles:    readPerm,
		// 			},
		// 		},
		// 	},
		// 	Perms: stub.PermData{
		// 		User:     secondaryUserName,
		// 		EntityID: secondaryUserID,
		// 		Roles:    readPerm,
		// 	},
		// },
		{
			// Tests a folder that has permissions with an item in the folder with
			// the different permissions.
			PathElements: folderAPath,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileEData,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
				},
			},
			Meta: stub.MetaData{
				Perms: stub.PermData{
					User:     secondaryUserName,
					EntityID: secondaryUserID,
					Roles:    readPerm,
				},
			},
		},
		{
			// Tests a folder that has permissions with an item in the folder with
			// no permissions.
			PathElements: folderCPath,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileAData,
					Meta: stub.MetaData{
						SharingMode: metadata.SharingModeInherited,
					},
				},
			},
			Meta: stub.MetaData{
				Perms: stub.PermData{
					User:     secondaryUserName,
					EntityID: secondaryUserID,
					Roles:    readPerm,
				},
			},
		},
	}

	expected, err := stub.DataForInfo(suite.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := stub.DataForInfo(suite.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.Service(),
				resourceCat:         suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("perms_restore_and_backup")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			runRestoreBackupTestVersions(
				t,
				testData,
				suite.Tenant(),
				[]string{suite.ResourceOwner()},
				control.DefaultOptions(),
				restoreCfg)
		})
	}
}

func testRestoreNoPermissionsAndBackup(suite oneDriveSuite, startVersion int) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.APIClient(),
		suite.Service(),
		suite.ResourceOwner())

	inputCols := []stub.ColInfo{
		{
			PathElements: []string{
				odConsts.DrivesPathDir,
				driveID,
				odConsts.RootPathDir,
			},
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileAData,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
						SharingMode: metadata.SharingModeCustom,
					},
				},
			},
		},
	}

	expectedCols := []stub.ColInfo{
		{
			PathElements: []string{
				odConsts.DrivesPathDir,
				driveID,
				odConsts.RootPathDir,
			},
			Files: []stub.ItemData{
				{
					// No permissions on the output since they weren't restored.
					Name: fileName,
					Data: fileAData,
				},
			},
		},
	}

	expected, err := stub.DataForInfo(suite.Service(), expectedCols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			input, err := stub.DataForInfo(suite.Service(), inputCols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.Service(),
				resourceCat:         suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("perms_backup_no_restore")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = false

			runRestoreBackupTestVersions(
				t,
				testData,
				suite.Tenant(),
				[]string{suite.ResourceOwner()},
				control.DefaultOptions(),
				restoreCfg)
		})
	}
}

// This is similar to TestPermissionsRestoreAndBackup but tests purely
// for inheritance and that too only with newer versions
func testPermissionsInheritanceRestoreAndBackup(suite oneDriveSuite, startVersion int) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	secondaryUserName, secondaryUserID := suite.SecondaryUser()
	tertiaryUserName, tertiaryUserID := suite.TertiaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.APIClient(),
		suite.Service(),
		suite.ResourceOwner())

	folderAName := "custom"
	folderBName := "inherited"
	folderCName := "empty"

	rootPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
	}
	folderAPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
	}
	subfolderAAPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderAName,
	}
	subfolderABPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderBName,
	}
	subfolderACPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderCName,
	}

	fileCustom := stub.ItemData{
		Name: "file-custom",
		Data: fileAData,
		Meta: stub.MetaData{
			Perms: stub.PermData{
				User:     secondaryUserName,
				EntityID: secondaryUserID,
				Roles:    writePerm,
			},
			SharingMode: metadata.SharingModeCustom,
		},
	}
	fileInherited := stub.ItemData{
		Name: "file-inherited",
		Data: fileAData,
		Meta: stub.MetaData{
			SharingMode: metadata.SharingModeInherited,
		},
	}
	fileEmpty := stub.ItemData{
		Name: "file-empty",
		Data: fileAData,
		Meta: stub.MetaData{
			SharingMode: metadata.SharingModeCustom,
		},
	}

	// If parent is empty, then empty permissions would be inherited
	fileEmptyInherited := stub.ItemData{
		Name: "file-empty",
		Data: fileAData,
		Meta: stub.MetaData{
			SharingMode: metadata.SharingModeInherited,
		},
	}

	fileSet := []stub.ItemData{fileCustom, fileInherited, fileEmpty}
	fileSetEmpty := []stub.ItemData{fileCustom, fileInherited, fileEmptyInherited}

	// Here is what this test is testing
	// - custom-permission-folder
	//   - custom-permission-file
	//   - inherted-permission-file
	//   - empty-permission-file
	//   - custom-permission-folder
	// 	   - custom-permission-file
	// 	   - inherted-permission-file
	//     - empty-permission-file
	//   - inherted-permission-folder
	// 	   - custom-permission-file
	// 	   - inherted-permission-file
	//     - empty-permission-file
	//   - empty-permission-folder
	// 	   - custom-permission-file
	// 	   - inherted-permission-file
	//     - empty-permission-file (empty/empty might have interesting behavior)

	cols := []stub.ColInfo{
		{
			PathElements: rootPath,
			Files:        []stub.ItemData{},
			Folders: []stub.ItemData{
				{Name: folderAName},
			},
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeInherited,
			},
		},
		{
			PathElements: folderAPath,
			Files:        fileSet,
			Folders: []stub.ItemData{
				{Name: folderAName},
				{Name: folderBName},
				{Name: folderCName},
			},
			Meta: stub.MetaData{
				Perms: stub.PermData{
					User:     tertiaryUserName,
					EntityID: tertiaryUserID,
					Roles:    readPerm,
				},
				SharingMode: metadata.SharingModeCustom,
			},
		},
		{
			PathElements: subfolderAAPath,
			Files:        fileSet,
			Meta: stub.MetaData{
				Perms: stub.PermData{
					User:     tertiaryUserName,
					EntityID: tertiaryUserID,
					Roles:    writePerm,
				},
				SharingMode: metadata.SharingModeCustom,
			},
		},
		{
			PathElements: subfolderABPath,
			Files:        fileSet,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeInherited,
			},
		},
		{
			PathElements: subfolderACPath,
			Files:        fileSetEmpty,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeCustom,
			},
		},
	}

	expected, err := stub.DataForInfo(suite.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := stub.DataForInfo(suite.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.Service(),
				resourceCat:         suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("perms_inherit_restore_and_backup")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			runRestoreBackupTestVersions(
				t,
				testData,
				suite.Tenant(),
				[]string{suite.ResourceOwner()},
				control.DefaultOptions(),
				restoreCfg)
		})
	}
}

func testLinkSharesInheritanceRestoreAndBackup(suite oneDriveSuite, startVersion int) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	_, secondaryUserID := suite.SecondaryUser()
	_, tertiaryUserID := suite.TertiaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.APIClient(),
		suite.Service(),
		suite.ResourceOwner())

	folderAName := "custom"
	folderBName := "inherited"
	folderCName := "empty"

	rootPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
	}
	folderAPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
	}
	subfolderAAPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderAName,
	}
	subfolderABPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderBName,
	}
	subfolderACPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderAName,
		folderCName,
	}

	fileSet := []stub.ItemData{
		{
			Name: "file-custom",
			Data: fileAData,
			Meta: stub.MetaData{
				LinkShares: []stub.LinkShareData{
					{
						EntityIDs: []string{secondaryUserID},
						Scope:     "users",
						Type:      "edit",
					},
				},
				SharingMode: metadata.SharingModeCustom,
			},
		},
		{
			Name: "file-inherited",
			Data: fileBData,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeInherited,
			},
		},
		{
			Name: "file-empty",
			Data: fileCData,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeCustom,
			},
		},
	}

	// Here is what this test is testing
	// - custom-link-share-folder
	//   - custom-link-share-file
	//   - inherted-link-share-file
	//   - empty-link-share-file
	//   - custom-link-share-folder
	// 	   - custom-link-share-file
	// 	   - inherted-link-share-file
	//     - empty-link-share-file
	//   - inherted-link-share-folder
	// 	   - custom-link-share-file
	// 	   - inherted-link-share-file
	//     - empty-link-share-file
	//   - empty-link-share-folder
	// 	   - custom-link-share-file
	// 	   - inherted-link-share-file
	//     - empty-link-share-file

	cols := []stub.ColInfo{
		{
			PathElements: rootPath,
			Files:        []stub.ItemData{},
			Folders: []stub.ItemData{
				{Name: folderAName},
			},
		},
		{
			PathElements: folderAPath,
			Files:        fileSet,
			Folders: []stub.ItemData{
				{Name: folderAName},
				{Name: folderBName},
				{Name: folderCName},
			},
			Meta: stub.MetaData{
				LinkShares: []stub.LinkShareData{
					{
						EntityIDs: []string{tertiaryUserID},
						Scope:     "anonymous",
						Type:      "edit",
					},
				},
			},
		},
		{
			PathElements: subfolderAAPath,
			Files:        fileSet,
			Meta: stub.MetaData{
				LinkShares: []stub.LinkShareData{
					{
						EntityIDs: []string{tertiaryUserID},
						Scope:     "users",
						Type:      "edit",
					},
				},
				SharingMode: metadata.SharingModeCustom,
			},
		},
		{
			PathElements: subfolderABPath,
			Files:        fileSet,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeInherited,
			},
		},
		{
			PathElements: subfolderACPath,
			Files:        fileSet,
			Meta: stub.MetaData{
				SharingMode: metadata.SharingModeCustom,
			},
		},
	}

	expected, err := stub.DataForInfo(suite.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := stub.DataForInfo(suite.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.Service(),
				resourceCat:         suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("linkshares_inherit_restore_and_backup")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			runRestoreBackupTestVersions(
				t,
				testData,
				suite.Tenant(),
				[]string{suite.ResourceOwner()},
				control.DefaultOptions(),
				restoreCfg)
		})
	}
}

func testRestoreFolderNamedFolderRegression(
	suite oneDriveSuite,
	startVersion int,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.APIClient(),
		suite.Service(),
		suite.ResourceOwner())

	rootPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
	}
	folderFolderPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderNamedFolder,
	}
	subfolderPath := []string{
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir,
		folderNamedFolder,
		folderBName,
	}

	cols := []stub.ColInfo{
		{
			PathElements: rootPath,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileAData,
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderNamedFolder,
				},
				{
					Name: folderBName,
				},
			},
		},
		{
			PathElements: folderFolderPath,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileBData,
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderBName,
				},
			},
		},
		{
			PathElements: subfolderPath,
			Files: []stub.ItemData{
				{
					Name: fileName,
					Data: fileCData,
				},
			},
			Folders: []stub.ItemData{
				{
					Name: folderNamedFolder,
				},
			},
		},
	}

	expected, err := stub.DataForInfo(suite.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			input, err := stub.DataForInfo(suite.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.Service(),
				resourceCat:         suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := control.DefaultRestoreConfig(dttm.HumanReadableDriveItem)
			restoreCfg.IncludePermissions = true

			runRestoreTestWithVersion(
				t,
				testData,
				suite.Tenant(),
				[]string{suite.ResourceOwner()},
				control.DefaultOptions(),
				restoreCfg)
		})
	}
}
