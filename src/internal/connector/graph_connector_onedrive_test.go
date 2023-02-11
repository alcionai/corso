package connector

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

func getTestMetaJSON(t *testing.T, user string, roles []string) []byte {
	id := base64.StdEncoding.EncodeToString([]byte(user + strings.Join(roles, "+")))
	testMeta := onedrive.Metadata{Permissions: []onedrive.UserPermission{
		{ID: id, Roles: roles, Email: user},
	}}

	testMetaJSON, err := json.Marshal(testMeta)
	if err != nil {
		t.Fatal("unable to marshall test permissions", err)
	}

	return testMetaJSON
}

func itemWithData(name string, itemData []byte) itemInfo {
	return itemInfo{
		name:      name,
		data:      itemData,
		lookupKey: name,
	}
}

func fileWithMetadata(baseName string, fileData, metadata []byte) []itemInfo {
	return []itemInfo{
		itemWithData(baseName+onedrive.DataFileSuffix, fileData),
		itemWithData(baseName+onedrive.MetaFileSuffix, metadata),
	}
}

type GraphConnectorOneDriveIntegrationSuite struct {
	suite.Suite
	connector     *GraphConnector
	user          string
	secondaryUser string
	acct          account.Account
}

func TestGraphConnectorOneDriveIntegrationSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

	suite.Run(t, new(GraphConnectorOneDriveIntegrationSuite))
}

func (suite *GraphConnectorOneDriveIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	tester.MustGetEnvSets(suite.T(), tester.M365AcctCredEnvs)

	suite.connector = loadConnector(ctx, suite.T(), graph.HTTPClient(graph.NoTimeout()), Users)
	suite.user = tester.M365UserID(suite.T())
	suite.secondaryUser = tester.SecondaryM365UserID(suite.T())
	suite.acct = tester.NewM365Account(suite.T())

	tester.LogTimeOfTest(suite.T())
}

var (
	fileEmptyPerms = itemWithData(
		"test-file.txt"+onedrive.MetaFileSuffix,
		[]byte("{}"),
	)

	fileAEmptyPerms = []itemInfo{
		itemWithData(
			"test-file.txt"+onedrive.DataFileSuffix,
			[]byte(strings.Repeat("a", 33)),
		),
		fileEmptyPerms,
	}

	fileBEmptyPerms = []itemInfo{
		itemWithData(
			"test-file.txt"+onedrive.DataFileSuffix,
			[]byte(strings.Repeat("b", 65)),
		),
		fileEmptyPerms,
	}

	fileCEmptyPerms = []itemInfo{
		itemWithData(
			"test-file.txt"+onedrive.DataFileSuffix,
			[]byte(strings.Repeat("c", 129)),
		),
		fileEmptyPerms,
	}

	fileDEmptyPerms = []itemInfo{
		itemWithData(
			"test-file.txt"+onedrive.DataFileSuffix,
			[]byte(strings.Repeat("d", 257)),
		),
		fileEmptyPerms,
	}

	fileEEmptyPerms = []itemInfo{
		itemWithData(
			"test-file.txt"+onedrive.DataFileSuffix,
			[]byte(strings.Repeat("e", 257)),
		),
		fileEmptyPerms,
	}

	folderAEmptyPerms = []itemInfo{
		itemWithData("folder-a"+onedrive.DirMetaFileSuffix, []byte("{}")),
	}

	folderBEmptyPerms = []itemInfo{
		itemWithData("b"+onedrive.DirMetaFileSuffix, []byte("{}")),
	}
)

func withItems(items ...[]itemInfo) []itemInfo {
	res := []itemInfo{}
	for _, i := range items {
		res = append(res, i...)
	}

	return res
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	table := []restoreBackupInfo{
		{
			name:     "OneDriveMultipleFoldersAndFiles",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: withItems(
						fileAEmptyPerms,
						folderAEmptyPerms,
						folderBEmptyPerms,
					),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
					},
					category: path.FilesCategory,
					items:    withItems(fileBEmptyPerms, folderBEmptyPerms),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
					},
					category: path.FilesCategory,
					items:    withItems(fileCEmptyPerms, folderAEmptyPerms),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
						"folder-a",
					},
					category: path.FilesCategory,
					items:    withItems(fileDEmptyPerms),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items:    withItems(fileEEmptyPerms),
					auxItems: []itemInfo{fileEmptyPerms},
				},
			},
		},
		{
			name:     "OneDriveFoldersAndFilesWithMetadata",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: withItems(
						fileWithMetadata(
							"test-file.txt",
							[]byte(strings.Repeat("a", 33)),
							getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
						),
						[]itemInfo{itemWithData(
							"b"+onedrive.DirMetaFileSuffix,
							getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
						)},
					),
					auxItems: []itemInfo{
						itemWithData(
							"test-file.txt"+onedrive.MetaFileSuffix,
							getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
						),
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: fileWithMetadata(
						"test-file.txt",
						[]byte(strings.Repeat("e", 66)),
						getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
					),
					auxItems: []itemInfo{
						itemWithData(
							"test-file.txt"+onedrive.MetaFileSuffix,
							getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
						),
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(
				t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackupVersion0() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	table := []restoreBackupInfoMultiVersion{
		{
			name:     "OneDriveMultipleFoldersAndFiles",
			service:  path.OneDriveService,
			resource: Users,

			collectionsPrevious: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						itemWithData(
							"test-file.txt",
							[]byte(strings.Repeat("a", 33)),
						),
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						itemWithData(
							"test-file.txt",
							[]byte(strings.Repeat("b", 65)),
						),
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						itemWithData(
							"test-file.txt",
							[]byte(strings.Repeat("c", 129)),
						),
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						itemWithData(
							"test-file.txt",
							[]byte(strings.Repeat("d", 257)),
						),
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						itemWithData(
							"test-file.txt",
							[]byte(strings.Repeat("e", 257)),
						),
					},
				},
			},

			collectionsLatest: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: withItems(
						fileAEmptyPerms,
						folderAEmptyPerms,
						folderBEmptyPerms,
					),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
					},
					category: path.FilesCategory,
					items:    withItems(fileBEmptyPerms, folderBEmptyPerms),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
					},
					category: path.FilesCategory,
					items:    withItems(fileCEmptyPerms, folderAEmptyPerms),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
						"folder-a",
					},
					category: path.FilesCategory,
					items:    fileDEmptyPerms,
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items:    fileEEmptyPerms,
					auxItems: []itemInfo{fileEmptyPerms},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTestVersions(
				t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsRestoreAndBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	var (
		fileAWritePerms = fileWithMetadata(
			"test-file.txt",
			[]byte(strings.Repeat("a", 33)),
			getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
		)

		fileEReadPerms = fileWithMetadata(
			"test-file.txt",
			[]byte(strings.Repeat("e", 66)),
			getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
		)

		folderBReadPerms = []itemInfo{itemWithData(
			"b"+onedrive.DirMetaFileSuffix,
			getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
		)}

		fileWritePerms = itemWithData(
			"test-file.txt"+onedrive.MetaFileSuffix,
			getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
		)

		fileReadPerms = itemWithData(
			"test-file.txt"+onedrive.MetaFileSuffix,
			getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
		)
	)

	table := []restoreBackupInfo{
		{
			name:     "FilePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items:    fileAWritePerms,
					auxItems: []itemInfo{fileWritePerms},
				},
			},
		},

		{
			name:     "FileInsideFolderPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items:    withItems(fileAEmptyPerms, folderBEmptyPerms),
					auxItems: []itemInfo{fileEmptyPerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items:    fileEReadPerms,
					auxItems: []itemInfo{fileReadPerms},
				},
			},
		},

		{
			name:     "FileAndFolderPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items:    withItems(fileAWritePerms, folderBReadPerms),
					auxItems: []itemInfo{fileWritePerms},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items:    fileEReadPerms,
					auxItems: []itemInfo{fileReadPerms},
				},
			},
		},

		{
			name:     "FileAndFolderSeparatePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items:    folderBReadPerms,
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items:    fileAWritePerms,
					auxItems: []itemInfo{fileWritePerms},
				},
			},
		},

		{
			name:     "FolderAndNoChildPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items:    folderBReadPerms,
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items:    fileEEmptyPerms,
					auxItems: []itemInfo{fileEmptyPerms},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsBackupAndNoRestore() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	table := []restoreBackupInfo{
		{
			name:     "FilePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: fileWithMetadata(
						"test-file.txt",
						[]byte(strings.Repeat("a", 33)),
						getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
					),
					auxItems: []itemInfo{
						itemWithData(
							"test-file.txt"+onedrive.MetaFileSuffix,
							getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
						),
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(
				t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}
