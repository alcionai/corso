package connector

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

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
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
	backupService path.ServiceType,
	service graph.Servicer,
	resourceOwner string,
) string {
	var (
		err error
		d   models.Driveable
	)

	switch backupService {
	case path.OneDriveService:
		d, err = api.GetUsersDrive(ctx, service, resourceOwner)
	case path.SharePointService:
		d, err = api.GetSitesDefaultDrive(ctx, service, resourceOwner)
	default:
		assert.FailNowf(t, "unknown service type %s", backupService.String())
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
	Service() graph.Servicer
	Account() account.Account
	Tenant() string
	// Returns (username, user ID) for the user. These values are used for
	// permissions.
	PrimaryUser() (string, string)
	SecondaryUser() (string, string)
	TertiaryUser() (string, string)
	// BackupResourceOwner returns the resource owner to run the backup/restore
	// with. This can be different from the values used for permissions and it can
	// also be a site.
	BackupResourceOwner() string
	BackupService() path.ServiceType
	Resource() Resource
}

type oneDriveSuite interface {
	tester.Suite
	suiteInfo
}

type suiteInfoImpl struct {
	connector       *GraphConnector
	resourceOwner   string
	user            string
	userID          string
	secondaryUser   string
	secondaryUserID string
	tertiaryUser    string
	tertiaryUserID  string
	acct            account.Account
	service         path.ServiceType
	resourceType    Resource
}

func (si suiteInfoImpl) Service() graph.Servicer {
	return si.connector.Service
}

func (si suiteInfoImpl) Account() account.Account {
	return si.acct
}

func (si suiteInfoImpl) Tenant() string {
	return si.connector.tenant
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

func (si suiteInfoImpl) BackupResourceOwner() string {
	return si.resourceOwner
}

func (si suiteInfoImpl) BackupService() path.ServiceType {
	return si.service
}

func (si suiteInfoImpl) Resource() Resource {
	return si.resourceType
}

// ---------------------------------------------------------------------------
// SharePoint Libraries
// ---------------------------------------------------------------------------
// SharePoint shares most of its libraries implementation with OneDrive so we
// only test simple things here and leave the more extensive testing to
// OneDrive.

type GraphConnectorSharePointIntegrationSuite struct {
	tester.Suite
	suiteInfo
}

func TestGraphConnectorSharePointIntegrationSuite(t *testing.T) {
	suite.Run(t, &GraphConnectorSharePointIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *GraphConnectorSharePointIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	si := suiteInfoImpl{
		connector:     loadConnector(ctx, suite.T(), Sites),
		user:          tester.M365UserID(suite.T()),
		secondaryUser: tester.SecondaryM365UserID(suite.T()),
		tertiaryUser:  tester.TertiaryM365UserID(suite.T()),
		acct:          tester.NewM365Account(suite.T()),
		service:       path.SharePointService,
		resourceType:  Sites,
	}

	si.resourceOwner = tester.M365SiteID(suite.T())

	user, err := si.connector.Discovery.Users().GetByID(ctx, si.user)
	require.NoError(t, err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(t, err, "fetching user", si.secondaryUser, clues.ToCore(err))
	si.secondaryUserID = ptr.Val(secondaryUser.GetId())

	tertiaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.tertiaryUser)
	require.NoError(t, err, "fetching user", si.tertiaryUser, clues.ToCore(err))
	si.tertiaryUserID = ptr.Val(tertiaryUser.GetId())

	suite.suiteInfo = si
}

func (suite *GraphConnectorSharePointIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, version.Backup)
}

func (suite *GraphConnectorSharePointIntegrationSuite) TestPermissionsRestoreAndBackup() {
	testPermissionsRestoreAndBackup(suite, version.Backup)
}

func (suite *GraphConnectorSharePointIntegrationSuite) TestPermissionsBackupAndNoRestore() {
	testPermissionsBackupAndNoRestore(suite, version.Backup)
}

func (suite *GraphConnectorSharePointIntegrationSuite) TestPermissionsInheritanceRestoreAndBackup() {
	testPermissionsInheritanceRestoreAndBackup(suite, version.Backup)
}

func (suite *GraphConnectorSharePointIntegrationSuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, version.All8MigrateUserPNToID)
}

// ---------------------------------------------------------------------------
// OneDrive most recent backup version
// ---------------------------------------------------------------------------
type GraphConnectorOneDriveIntegrationSuite struct {
	tester.Suite
	suiteInfo
}

func TestGraphConnectorOneDriveIntegrationSuite(t *testing.T) {
	suite.Run(t, &GraphConnectorOneDriveIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *GraphConnectorOneDriveIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	si := suiteInfoImpl{
		connector:     loadConnector(ctx, t, Users),
		user:          tester.M365UserID(t),
		secondaryUser: tester.SecondaryM365UserID(t),
		acct:          tester.NewM365Account(t),
		service:       path.OneDriveService,
		resourceType:  Users,
	}

	si.resourceOwner = si.user

	user, err := si.connector.Discovery.Users().GetByID(ctx, si.user)
	require.NoError(t, err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(t, err, "fetching user", si.secondaryUser, clues.ToCore(err))
	si.secondaryUserID = ptr.Val(secondaryUser.GetId())

	suite.suiteInfo = si
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, version.Backup)
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsRestoreAndBackup() {
	testPermissionsRestoreAndBackup(suite, version.Backup)
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsBackupAndNoRestore() {
	testPermissionsBackupAndNoRestore(suite, version.Backup)
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsInheritanceRestoreAndBackup() {
	testPermissionsInheritanceRestoreAndBackup(suite, version.Backup)
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, version.All8MigrateUserPNToID)
}

// ---------------------------------------------------------------------------
// OneDrive regression
// ---------------------------------------------------------------------------
type GraphConnectorOneDriveNightlySuite struct {
	tester.Suite
	suiteInfo
}

func TestGraphConnectorOneDriveNightlySuite(t *testing.T) {
	suite.Run(t, &GraphConnectorOneDriveNightlySuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *GraphConnectorOneDriveNightlySuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	si := suiteInfoImpl{
		connector:     loadConnector(ctx, t, Users),
		user:          tester.M365UserID(t),
		secondaryUser: tester.SecondaryM365UserID(t),
		acct:          tester.NewM365Account(t),
		service:       path.OneDriveService,
		resourceType:  Users,
	}

	si.resourceOwner = si.user

	user, err := si.connector.Discovery.Users().GetByID(ctx, si.user)
	require.NoError(t, err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(t, err, "fetching user", si.secondaryUser, clues.ToCore(err))
	si.secondaryUserID = ptr.Val(secondaryUser.GetId())

	suite.suiteInfo = si
}

func (suite *GraphConnectorOneDriveNightlySuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, 0)
}

func (suite *GraphConnectorOneDriveNightlySuite) TestPermissionsRestoreAndBackup() {
	testPermissionsRestoreAndBackup(suite, version.OneDrive1DataAndMetaFiles)
}

func (suite *GraphConnectorOneDriveNightlySuite) TestPermissionsBackupAndNoRestore() {
	testPermissionsBackupAndNoRestore(suite, version.OneDrive1DataAndMetaFiles)
}

func (suite *GraphConnectorOneDriveNightlySuite) TestPermissionsInheritanceRestoreAndBackup() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testPermissionsInheritanceRestoreAndBackup(suite, version.OneDrive4DirIncludesPermissions)
}

func (suite *GraphConnectorOneDriveNightlySuite) TestRestoreFolderNamedFolderRegression() {
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
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

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

	cols := []OnedriveColInfo{
		{
			PathElements: rootPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileAData,
				},
			},
			Folders: []ItemData{
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
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileBData,
				},
			},
			Folders: []ItemData{
				{
					Name: folderBName,
				},
			},
		},
		{
			PathElements: subfolderBPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileCData,
				},
			},
			Folders: []ItemData{
				{
					Name: folderAName,
				},
			},
		},
		{
			PathElements: subfolderAPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileDData,
				},
			},
		},
		{
			PathElements: folderBPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileEData,
				},
			},
		},
	}

	expected, err := DataForInfo(suite.BackupService(), cols, version.Backup)
	require.NoError(suite.T(), err)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input, err := DataForInfo(suite.BackupService(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.BackupService(),
				resource:            suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{suite.BackupResourceOwner()},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{},
				})
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
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

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

	cols := []OnedriveColInfo{
		{
			PathElements: rootPath,
			Files: []ItemData{
				{
					// Test restoring a file that doesn't inherit permissions.
					Name: fileName,
					Data: fileAData,
					Perms: PermData{
						User:     secondaryUserName,
						EntityID: secondaryUserID,
						Roles:    writePerm,
					},
				},
				{
					// Test restoring a file that doesn't inherit permissions and has
					// no permissions.
					Name: fileName2,
					Data: fileBData,
				},
			},
			Folders: []ItemData{
				{
					Name: folderBName,
				},
				{
					Name: folderAName,
					Perms: PermData{
						User:     secondaryUserName,
						EntityID: secondaryUserID,
						Roles:    readPerm,
					},
				},
				{
					Name: folderCName,
					Perms: PermData{
						User:     secondaryUserName,
						EntityID: secondaryUserID,
						Roles:    readPerm,
					},
				},
			},
		},
		{
			PathElements: folderBPath,
			Files: []ItemData{
				{
					// Test restoring a file in a non-root folder that doesn't inherit
					// permissions.
					Name: fileName,
					Data: fileBData,
					Perms: PermData{
						User:     secondaryUserName,
						EntityID: secondaryUserID,
						Roles:    writePerm,
					},
				},
			},
			Folders: []ItemData{
				{
					Name: folderAName,
					Perms: PermData{
						User:     secondaryUserName,
						EntityID: secondaryUserID,
						Roles:    readPerm,
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
		// 	files: []itemData{
		// 		{
		// 			name: fileName,
		// 			data: fileDData,
		// 			perms: permData{
		// 				user:     secondaryUserName,
		// 				entityID: secondaryUserID,
		// 				roles:    readPerm,
		// 			},
		// 		},
		// 	},
		// 	Perms: PermData{
		// 		User:     secondaryUserName,
		// 		EntityID: secondaryUserID,
		// 		Roles:    readPerm,
		// 	},
		// },
		{
			// Tests a folder that has permissions with an item in the folder with
			// the different permissions.
			PathElements: folderAPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileEData,
					Perms: PermData{
						User:     secondaryUserName,
						EntityID: secondaryUserID,
						Roles:    writePerm,
					},
				},
			},
			Perms: PermData{
				User:     secondaryUserName,
				EntityID: secondaryUserID,
				Roles:    readPerm,
			},
		},
		{
			// Tests a folder that has permissions with an item in the folder with
			// no permissions.
			PathElements: folderCPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileAData,
				},
			},
			Perms: PermData{
				User:     secondaryUserName,
				EntityID: secondaryUserID,
				Roles:    readPerm,
			},
		},
	}

	expected, err := DataForInfo(suite.BackupService(), cols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.BackupService().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := DataForInfo(suite.BackupService(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.BackupService(),
				resource:            suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{suite.BackupResourceOwner()},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{},
				})
		})
	}
}

func testPermissionsBackupAndNoRestore(suite oneDriveSuite, startVersion int) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

	inputCols := []OnedriveColInfo{
		{
			PathElements: []string{
				odConsts.DrivesPathDir,
				driveID,
				odConsts.RootPathDir,
			},
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileAData,
					Perms: PermData{
						User:     secondaryUserName,
						EntityID: secondaryUserID,
						Roles:    writePerm,
					},
				},
			},
		},
	}

	expectedCols := []OnedriveColInfo{
		{
			PathElements: []string{
				odConsts.DrivesPathDir,
				driveID,
				odConsts.RootPathDir,
			},
			Files: []ItemData{
				{
					// No permissions on the output since they weren't restored.
					Name: fileName,
					Data: fileAData,
				},
			},
		},
	}

	expected, err := DataForInfo(suite.BackupService(), expectedCols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.BackupService().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			input, err := DataForInfo(suite.BackupService(), inputCols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.BackupService(),
				resource:            suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{suite.BackupResourceOwner()},
				control.Options{
					RestorePermissions: false,
					ToggleFeatures:     control.Toggles{},
				})
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
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

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

	fileSet := []ItemData{
		{
			Name: "file-custom",
			Data: fileAData,
			Perms: PermData{
				User:        secondaryUserName,
				EntityID:    secondaryUserID,
				Roles:       writePerm,
				SharingMode: metadata.SharingModeCustom,
			},
		},
		{
			Name: "file-inherited",
			Data: fileAData,
			Perms: PermData{
				SharingMode: metadata.SharingModeInherited,
			},
		},
		{
			Name: "file-empty",
			Data: fileAData,
			Perms: PermData{
				SharingMode: metadata.SharingModeCustom,
			},
		},
	}

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

	cols := []OnedriveColInfo{
		{
			PathElements: rootPath,
			Files:        []ItemData{},
			Folders: []ItemData{
				{Name: folderAName},
			},
		},
		{
			PathElements: folderAPath,
			Files:        fileSet,
			Folders: []ItemData{
				{Name: folderAName},
				{Name: folderBName},
				{Name: folderCName},
			},
			Perms: PermData{
				User:     tertiaryUserName,
				EntityID: tertiaryUserID,
				Roles:    readPerm,
			},
		},
		{
			PathElements: subfolderAAPath,
			Files:        fileSet,
			Perms: PermData{
				User:        tertiaryUserName,
				EntityID:    tertiaryUserID,
				Roles:       writePerm,
				SharingMode: metadata.SharingModeCustom,
			},
		},
		{
			PathElements: subfolderABPath,
			Files:        fileSet,
			Perms: PermData{
				SharingMode: metadata.SharingModeInherited,
			},
		},
		{
			PathElements: subfolderACPath,
			Files:        fileSet,
			Perms: PermData{
				SharingMode: metadata.SharingModeCustom,
			},
		},
	}

	expected, err := DataForInfo(suite.BackupService(), cols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.BackupService().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := DataForInfo(suite.BackupService(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.BackupService(),
				resource:            suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{suite.BackupResourceOwner()},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{},
				})
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
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

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

	cols := []OnedriveColInfo{
		{
			PathElements: rootPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileAData,
				},
			},
			Folders: []ItemData{
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
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileBData,
				},
			},
			Folders: []ItemData{
				{
					Name: folderBName,
				},
			},
		},
		{
			PathElements: subfolderPath,
			Files: []ItemData{
				{
					Name: fileName,
					Data: fileCData,
				},
			},
			Folders: []ItemData{
				{
					Name: folderNamedFolder,
				},
			},
		},
	}

	expected, err := DataForInfo(suite.BackupService(), cols, version.Backup)
	require.NoError(suite.T(), err)
	bss := suite.BackupService().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			input, err := DataForInfo(suite.BackupService(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             suite.BackupService(),
				resource:            suite.Resource(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreTestWithVerion(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{suite.BackupResourceOwner()},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{},
				},
			)
		})
	}
}
