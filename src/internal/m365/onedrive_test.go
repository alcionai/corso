package m365

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/stub"
	m365Stub "github.com/alcionai/corso/src/internal/m365/stub"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/path"
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

// type suiteInfoImpl struct {
// 	controller    *Controller
// 	resourceOwner string
// 	service       path.ServiceType
// }

// ---------------------------------------------------------------------------
// SharePoint Libraries
// ---------------------------------------------------------------------------
// SharePoint shares most of its libraries implementation with OneDrive so we
// only test simple things here and leave the more extensive testing to
// OneDrive.

type SharePointIntegrationSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
	rs   its.ResourceServicer
}

func TestSharePointIntegrationSuite(t *testing.T) {
	suite.Run(t, &SharePointIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SharePointIntegrationSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
	suite.rs = its.NewResourceService(suite.m365.Site, path.SharePointService)
}

func (suite *SharePointIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, suite.m365, suite.rs, version.Backup)
}

// TODO: Re-enable these tests (disabled as it currently acting up CI)
func (suite *SharePointIntegrationSuite) TestPermissionsRestoreAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testPermissionsRestoreAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestRestoreNoPermissionsAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testRestoreNoPermissionsAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestPermissionsInheritanceRestoreAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testPermissionsInheritanceRestoreAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestLinkSharesInheritanceRestoreAndBackup() {
	suite.T().Skip("Temporarily disabled due to CI issues")
	testLinkSharesInheritanceRestoreAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *SharePointIntegrationSuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, suite.m365, suite.rs, version.All8MigrateUserPNToID)
}

// ---------------------------------------------------------------------------
// OneDrive most recent backup version
// ---------------------------------------------------------------------------
type OneDriveIntegrationSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
	rs   its.ResourceServicer
}

func TestOneDriveIntegrationSuite(t *testing.T) {
	suite.Run(t, &OneDriveIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *OneDriveIntegrationSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
	suite.rs = its.NewResourceService(suite.m365.User, path.OneDriveService)
}

func (suite *OneDriveIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestPermissionsRestoreAndBackup() {
	testPermissionsRestoreAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestRestoreNoPermissionsAndBackup() {
	testRestoreNoPermissionsAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestPermissionsInheritanceRestoreAndBackup() {
	testPermissionsInheritanceRestoreAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestLinkSharesInheritanceRestoreAndBackup() {
	testLinkSharesInheritanceRestoreAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *OneDriveIntegrationSuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, suite.m365, suite.rs, version.All8MigrateUserPNToID)
}

// ---------------------------------------------------------------------------
// OneDrive regression
// ---------------------------------------------------------------------------
type OneDriveNightlySuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
	rs   its.ResourceServicer
}

func TestOneDriveNightlySuite(t *testing.T) {
	suite.Run(t, &OneDriveNightlySuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *OneDriveNightlySuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
	suite.rs = its.NewResourceService(suite.m365.User, path.OneDriveService)
}

func (suite *OneDriveNightlySuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, suite.m365, suite.rs, 0)
}

func (suite *OneDriveNightlySuite) TestPermissionsRestoreAndBackup() {
	testPermissionsRestoreAndBackup(suite, suite.m365, suite.rs, version.OneDrive1DataAndMetaFiles)
}

func (suite *OneDriveNightlySuite) TestRestoreNoPermissionsAndBackup() {
	testRestoreNoPermissionsAndBackup(suite, suite.m365, suite.rs, version.OneDrive1DataAndMetaFiles)
}

func (suite *OneDriveNightlySuite) TestPermissionsInheritanceRestoreAndBackup() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testPermissionsInheritanceRestoreAndBackup(suite, suite.m365, suite.rs, version.OneDrive4DirIncludesPermissions)
}

func (suite *OneDriveNightlySuite) TestLinkSharesInheritanceRestoreAndBackup() {
	testLinkSharesInheritanceRestoreAndBackup(suite, suite.m365, suite.rs, version.Backup)
}

func (suite *OneDriveNightlySuite) TestRestoreFolderNamedFolderRegression() {
	// No reason why it couldn't work with previous versions, but this is when it got introduced.
	testRestoreFolderNamedFolderRegression(suite, suite.m365, suite.rs, version.All8MigrateUserPNToID)
}

func testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(
	suite tester.Suite,
	m365 its.M365IntgTestSetup,
	irs its.ResourceServicer,
	startVersion int,
) {
	// Get the default drive ID for the test user.
	driveID := irs.Resource().DriveID

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

	expected, err := stub.DataForInfo(irs.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input, err := stub.DataForInfo(irs.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             irs.Service(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("od_restore_and_backup_multi")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			opts := control.DefaultOptions()
			opts.ToggleFeatures.UseDeltaTree = true

			cfg := m365Stub.ConfigInfo{
				Tenant:         m365.TenantID,
				ResourceOwners: []string{irs.Resource().ID},
				Service:        testData.service,
				Opts:           opts,
				RestoreCfg:     restoreCfg,
			}

			runRestoreBackupTestVersions(t, testData, cfg)
		})
	}
}

func testPermissionsRestoreAndBackup(
	suite tester.Suite,
	m365 its.M365IntgTestSetup,
	irs its.ResourceServicer,
	startVersion int,
) {
	// Get the default drive ID for the test user.
	driveID := irs.Resource().DriveID

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
							User:     m365.SecondaryUser.Email,
							EntityID: m365.SecondaryUser.ID,
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
							User:     m365.SecondaryUser.Email,
							EntityID: m365.SecondaryUser.ID,
							Roles:    readPerm,
						},
					},
				},
				{
					Name: folderCName,
					Meta: stub.MetaData{
						Perms: stub.PermData{
							User:     m365.SecondaryUser.Email,
							EntityID: m365.SecondaryUser.ID,
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
							User:     m365.SecondaryUser.Email,
							EntityID: m365.SecondaryUser.ID,
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
							User:     m365.SecondaryUser.Email,
							EntityID: m365.SecondaryUser.ID,
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
		// 				user:     m365.SecondaryUser.Email,
		// 				entityID: m365.SecondaryUser.ID,
		// 				roles:    readPerm,
		// 			},
		// 		},
		// 	},
		// 	Perms: stub.PermData{
		// 		User:     m365.SecondaryUser.Email,
		// 		EntityID: m365.SecondaryUser.ID,
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
							User:     m365.SecondaryUser.Email,
							EntityID: m365.SecondaryUser.ID,
							Roles:    writePerm,
						},
					},
				},
			},
			Meta: stub.MetaData{
				Perms: stub.PermData{
					User:     m365.SecondaryUser.Email,
					EntityID: m365.SecondaryUser.ID,
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
					User:     m365.SecondaryUser.Email,
					EntityID: m365.SecondaryUser.ID,
					Roles:    readPerm,
				},
			},
		},
	}

	expected, err := stub.DataForInfo(irs.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)

	bss := irs.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := stub.DataForInfo(irs.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             irs.Service(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("perms_restore_and_backup")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			opts := control.DefaultOptions()
			opts.ToggleFeatures.UseDeltaTree = true

			cfg := m365Stub.ConfigInfo{
				Tenant:         m365.TenantID,
				ResourceOwners: []string{irs.Resource().ID},
				Service:        testData.service,
				Opts:           opts,
				RestoreCfg:     restoreCfg,
			}

			runRestoreBackupTestVersions(t, testData, cfg)
		})
	}
}

func testRestoreNoPermissionsAndBackup(
	suite tester.Suite,
	m365 its.M365IntgTestSetup,
	irs its.ResourceServicer,
	startVersion int,
) {
	// Get the default drive ID for the test user.
	driveID := irs.Resource().DriveID

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
							User:     m365.SecondaryUser.Email,
							EntityID: m365.SecondaryUser.ID,
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

	expected, err := stub.DataForInfo(irs.Service(), expectedCols, version.Backup)
	require.NoError(suite.T(), err, clues.ToCore(err))

	bss := irs.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()

			input, err := stub.DataForInfo(irs.Service(), inputCols, vn)
			require.NoError(t, err, clues.ToCore(err))

			testData := restoreBackupInfoMultiVersion{
				service:             irs.Service(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("perms_backup_no_restore")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = false

			opts := control.DefaultOptions()
			opts.ToggleFeatures.UseDeltaTree = true

			cfg := m365Stub.ConfigInfo{
				Tenant:         m365.TenantID,
				ResourceOwners: []string{irs.Resource().ID},
				Service:        testData.service,
				Opts:           opts,
				RestoreCfg:     restoreCfg,
			}

			runRestoreBackupTestVersions(t, testData, cfg)
		})
	}
}

// This is similar to TestPermissionsRestoreAndBackup but tests purely
// for inheritance and that too only with newer versions
func testPermissionsInheritanceRestoreAndBackup(
	suite tester.Suite,
	m365 its.M365IntgTestSetup,
	irs its.ResourceServicer,
	startVersion int,
) {
	// Get the default drive ID for the test user.
	driveID := irs.Resource().DriveID

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
				User:     m365.SecondaryUser.Email,
				EntityID: m365.SecondaryUser.ID,
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
					User:     m365.TertiaryUser.Email,
					EntityID: m365.TertiaryUser.ID,
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
					User:     m365.TertiaryUser.Email,
					EntityID: m365.TertiaryUser.ID,
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

	expected, err := stub.DataForInfo(irs.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)

	bss := irs.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := stub.DataForInfo(irs.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             irs.Service(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("perms_inherit_restore_and_backup")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			opts := control.DefaultOptions()
			opts.ToggleFeatures.UseDeltaTree = true

			cfg := m365Stub.ConfigInfo{
				Tenant:         m365.TenantID,
				ResourceOwners: []string{irs.Resource().ID},
				Service:        testData.service,
				Opts:           opts,
				RestoreCfg:     restoreCfg,
			}

			runRestoreBackupTestVersions(t, testData, cfg)
		})
	}
}

func testLinkSharesInheritanceRestoreAndBackup(
	suite tester.Suite,
	m365 its.M365IntgTestSetup,
	irs its.ResourceServicer,
	startVersion int,
) {
	// Get the default drive ID for the test user.
	driveID := irs.Resource().DriveID

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
						EntityIDs: []string{m365.SecondaryUser.ID},
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
						EntityIDs: []string{m365.TertiaryUser.ID},
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
						EntityIDs: []string{m365.TertiaryUser.ID},
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

	expected, err := stub.DataForInfo(irs.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)

	bss := irs.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input, err := stub.DataForInfo(irs.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             irs.Service(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := testdata.DefaultRestoreConfig("linkshares_inherit_restore_and_backup")
			restoreCfg.OnCollision = control.Replace
			restoreCfg.IncludePermissions = true

			opts := control.DefaultOptions()
			opts.ToggleFeatures.UseDeltaTree = true

			cfg := m365Stub.ConfigInfo{
				Tenant:         m365.TenantID,
				ResourceOwners: []string{irs.Resource().ID},
				Service:        testData.service,
				Opts:           opts,
				RestoreCfg:     restoreCfg,
			}

			runRestoreBackupTestVersions(t, testData, cfg)
		})
	}
}

func testRestoreFolderNamedFolderRegression(
	suite tester.Suite,
	m365 its.M365IntgTestSetup,
	irs its.ResourceServicer,
	startVersion int,
) {
	// Get the default drive ID for the test user.
	driveID := irs.Resource().DriveID

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

	expected, err := stub.DataForInfo(irs.Service(), cols, version.Backup)
	require.NoError(suite.T(), err)

	bss := irs.Service().String()

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("%s-Version%d", bss, vn), func() {
			t := suite.T()
			input, err := stub.DataForInfo(irs.Service(), cols, vn)
			require.NoError(suite.T(), err)

			testData := restoreBackupInfoMultiVersion{
				service:             irs.Service(),
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			restoreCfg := control.DefaultRestoreConfig(dttm.SafeForTesting)
			restoreCfg.IncludePermissions = true

			opts := control.DefaultOptions()
			opts.ToggleFeatures.UseDeltaTree = true

			cfg := m365Stub.ConfigInfo{
				Tenant:         m365.TenantID,
				ResourceOwners: []string{irs.Resource().ID},
				Service:        testData.service,
				Opts:           opts,
				RestoreCfg:     restoreCfg,
			}

			runRestoreTestWithVersion(t, testData, cfg)
		})
	}
}
