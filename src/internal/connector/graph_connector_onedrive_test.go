package connector

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

func getMetadata(fileName, user string, roles []string) onedrive.Metadata {
	if len(user) == 0 || len(roles) == 0 {
		return onedrive.Metadata{FileName: fileName}
	}

	id := base64.StdEncoding.EncodeToString([]byte(user + strings.Join(roles, "+")))
	testMeta := onedrive.Metadata{
		FileName: fileName,
		Permissions: []onedrive.UserPermission{
			{ID: id, Roles: roles, Email: user},
		},
	}

	return testMeta
}

func getTestMetaJSON(t *testing.T, fileName, user string, roles []string) []byte {
	testMeta := getMetadata(fileName, user, roles)

	testMetaJSON, err := json.Marshal(testMeta)
	if err != nil {
		t.Fatal("unable to marshall test permissions", err)
	}

	return testMetaJSON
}

type testOneDriveData struct {
	FileName string `json:"fileName,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

func onedriveItemWithData(
	t *testing.T,
	name, lookupKey string,
	fileData []byte,
) itemInfo {
	t.Helper()

	content := testOneDriveData{
		FileName: lookupKey,
		Data:     fileData,
	}

	serialized, err := json.Marshal(content)
	require.NoError(t, err)

	return itemInfo{
		name:      name,
		data:      serialized,
		lookupKey: lookupKey,
	}
}

func onedriveMetadata(
	t *testing.T,
	fileName, itemID string,
	user string,
	roles []string,
) itemInfo {
	t.Helper()

	testMeta := getMetadata(fileName, user, roles)

	testMetaJSON, err := json.Marshal(testMeta)
	require.NoError(t, err, "marshalling metadata")

	return itemInfo{
		name:      itemID,
		data:      testMetaJSON,
		lookupKey: itemID,
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
	fileName    = "test-file.txt"
	folderAName = "folder-a"
	folderBName = "b"

	fileAData = []byte(strings.Repeat("a", 33))
	fileBData = []byte(strings.Repeat("b", 65))
	fileCData = []byte(strings.Repeat("c", 129))
	fileDData = []byte(strings.Repeat("d", 257))
	fileEData = []byte(strings.Repeat("e", 257))
)

func newOneDriveCollection(
	t *testing.T,
	pathElements []string,
	backupVersion int,
) *onedriveCollection {
	return &onedriveCollection{
		pathElements:  pathElements,
		backupVersion: backupVersion,
		t:             t,
	}
}

type onedriveCollection struct {
	pathElements  []string
	items         []itemInfo
	aux           []itemInfo
	backupVersion int
	t             *testing.T
}

func (c onedriveCollection) collection() colInfo {
	return colInfo{
		pathElements: c.pathElements,
		category:     path.FilesCategory,
		items:        c.items,
		auxItems:     c.aux,
	}
}

func (c *onedriveCollection) withFile(
	name string,
	fileData []byte,
	user string,
	roles []string,
) *onedriveCollection {
	switch c.backupVersion {
	case 0:
		// Lookups will occur using the most recent version of things so we need
		// the embedded file name to match that.
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name,
			name+onedrive.DataFileSuffix,
			fileData))

	case 1:
		fallthrough
	case 2:
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name+onedrive.DataFileSuffix,
			name+onedrive.DataFileSuffix,
			fileData))

		metadata := onedriveMetadata(
			c.t,
			"",
			name+onedrive.MetaFileSuffix,
			user,
			roles)
		c.items = append(c.items, metadata)
		c.aux = append(c.aux, metadata)

	default:
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name+"-id"+onedrive.DataFileSuffix,
			name+onedrive.DataFileSuffix,
			fileData))

		metadata := onedriveMetadata(
			c.t,
			name,
			name+"-id"+onedrive.MetaFileSuffix,
			user,
			roles)
		c.items = append(c.items, metadata)
		c.aux = append(c.aux, metadata)
	}

	return c
}

func (c *onedriveCollection) withFolder(
	name string,
	user string,
	roles []string,
) *onedriveCollection {
	if c.backupVersion < 1 {
		return c
	}

	switch c.backupVersion {
	case 1:
		fallthrough
	case 2:
		c.items = append(
			c.items,
			onedriveMetadata(
				c.t,
				"",
				name+onedrive.DirMetaFileSuffix,
				user,
				roles),
		)

	default:
		c.items = append(
			c.items,
			onedriveMetadata(
				c.t,
				name,
				name+onedrive.DirMetaFileSuffix,
				user,
				roles),
		)
	}

	return c
}

// withPermissions adds permissions to the folder represented by this
// onedriveCollection.
func (c *onedriveCollection) withPermissions(
	user string,
	roles []string,
) *onedriveCollection {
	// These versions didn't store permissions for the folder or didn't store them
	// in the folder's collection.
	if c.backupVersion < 3 {
		return c
	}

	name := c.pathElements[len(c.pathElements)-1]

	c.items = append(
		c.items,
		onedriveMetadata(
			c.t,
			name,
			name+onedrive.DirMetaFileSuffix,
			user,
			roles),
	)
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

	collectionsLatest := []colInfo{
		newOneDriveCollection(
			suite.T(),
			[]string{
				"drives",
				driveID,
				"root:",
			},
			3,
		).
			withFile(
				"test-file.txt",
				fileAData,
				suite.secondaryUser,
				[]string{"write"},
			).
			withFolder(
				"b",
				suite.secondaryUser,
				[]string{"read"},
			).
			collection(),

		newOneDriveCollection(
			suite.T(),
			[]string{
				"drives",
				driveID,
				"root:",
				"b",
			},
			3,
		).
			withFile(
				"test-file.txt",
				fileEData,
				suite.secondaryUser,
				[]string{"read"},
			).
			collection(),
	}

	table := []restoreBackupInfoMultiVersion{
		{
			name:          "OneDriveMultipleFoldersAndFilesWithMetadata_Version1",
			service:       path.OneDriveService,
			resource:      Users,
			backupVersion: 1, // The OG version ;)
			countMeta:     false,
			collectionsPrevious: []colInfo{
				newOneDriveCollection(
					suite.T(),
					[]string{
						"drives",
						driveID,
						"root:",
					},
					1,
				).
					withFile(
						"test-file.txt",
						fileAData,
						suite.secondaryUser,
						[]string{"write"},
					).
					withFolder(
						"b",
						suite.secondaryUser,
						[]string{"read"},
					).
					collection(),

				newOneDriveCollection(
					suite.T(),
					[]string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					1,
				).
					withFile(
						"test-file.txt",
						fileEData,
						suite.secondaryUser,
						[]string{"read"},
					).
					collection(),
			},
			collectionsLatest: collectionsLatest,
		},
		{
			name:                "OneDriveMultipleFoldersAndFilesWithMetadata_Version3",
			service:             path.OneDriveService,
			resource:            Users,
			backupVersion:       3, // The OG version ;)
			countMeta:           false,
			collectionsPrevious: collectionsLatest,
			collectionsLatest:   collectionsLatest,
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

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackup_Versions() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	rootPath := []string{
		"drives",
		driveID,
		"root:",
	}
	folderAPath := []string{
		"drives",
		driveID,
		"root:",
		"folder-a",
	}
	subfolderBPath := []string{
		"drives",
		driveID,
		"root:",
		"folder-a",
		"b",
	}
	subfolderAPath := []string{
		"drives",
		driveID,
		"root:",
		"folder-a",
		"b",
		"folder-a",
	}
	folderBPath := []string{
		"drives",
		driveID,
		"root:",
		"b",
	}

	collectionsLatest := []colInfo{
		newOneDriveCollection(suite.T(), rootPath, 1).
			withFile(
				fileName,
				fileAData,
				"",
				nil,
			).
			withFolder(
				"folder-a",
				"",
				nil,
			).
			withFolder(
				"b",
				"",
				nil,
			).
			collection(),
		newOneDriveCollection(suite.T(), folderAPath, 1).
			withFile(
				fileName,
				fileBData,
				"",
				nil,
			).
			withFolder(
				"b",
				"",
				nil,
			).
			collection(),
		newOneDriveCollection(suite.T(), subfolderBPath, 1).
			withFile(
				fileName,
				fileCData,
				"",
				nil,
			).
			withFolder(
				"folder-a",
				"",
				nil,
			).
			collection(),
		newOneDriveCollection(suite.T(), subfolderAPath, 1).
			withFile(
				fileName,
				fileDData,
				"",
				nil,
			).
			collection(),
		newOneDriveCollection(suite.T(), folderBPath, 1).
			withFile(
				fileName,
				fileEData,
				"",
				nil,
			).
			collection(),
	}

	table := []restoreBackupInfoMultiVersion{
		{
			name:          "OneDriveMultipleFoldersAndFiles_Version0",
			service:       path.OneDriveService,
			resource:      Users,
			backupVersion: 0, // The OG version ;)
			countMeta:     true,

			collectionsPrevious: []colInfo{
				newOneDriveCollection(suite.T(), rootPath, 0).
					withFile(
						fileName,
						fileAData,
						"",
						nil,
					).
					collection(),
				newOneDriveCollection(suite.T(), folderAPath, 0).
					withFile(
						fileName,
						fileBData,
						"",
						nil,
					).
					collection(),
				newOneDriveCollection(suite.T(), subfolderBPath, 0).
					withFile(
						fileName,
						fileCData,
						"",
						nil,
					).
					collection(),
				newOneDriveCollection(suite.T(), subfolderAPath, 0).
					withFile(
						fileName,
						fileDData,
						"",
						nil,
					).
					collection(),
				newOneDriveCollection(suite.T(), folderBPath, 0).
					withFile(
						fileName,
						fileEData,
						"",
						nil,
					).
					collection(),
			},

			collectionsLatest: collectionsLatest,
		},

		{
			name:                "OneDriveMultipleFoldersAndFiles_Version1",
			service:             path.OneDriveService,
			resource:            Users,
			backupVersion:       1,
			countMeta:           false,
			collectionsPrevious: collectionsLatest,
			collectionsLatest:   collectionsLatest,
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

	rootPath := []string{
		"drives",
		driveID,
		"root:",
	}
	folderPath := []string{
		"drives",
		driveID,
		"root:",
		"b",
	}

	table := []restoreBackupInfo{
		{
			name:     "FilePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				newOneDriveCollection(suite.T(), rootPath, 1).
					withFile(
						fileName,
						fileAData,
						suite.secondaryUser,
						[]string{"write"},
					).
					collection(),
			},
		},

		{
			name:     "FileInsideFolderPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				newOneDriveCollection(suite.T(), rootPath, 1).
					withFile(
						fileName,
						fileAData,
						"",
						nil,
					).
					withFolder(
						"b",
						"",
						nil,
					).
					collection(),
				newOneDriveCollection(suite.T(), folderPath, 1).
					withFile(
						fileName,
						fileEData,
						suite.secondaryUser,
						[]string{"read"},
					).
					collection(),
			},
		},

		{
			name:     "FileAndFolderPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				newOneDriveCollection(suite.T(), rootPath, 1).
					withFile(
						fileName,
						fileAData,
						suite.secondaryUser,
						[]string{"write"},
					).
					withFolder(
						"b",
						suite.secondaryUser,
						[]string{"read"},
					).
					collection(),
				newOneDriveCollection(suite.T(), folderPath, 1).
					withFile(
						fileName,
						fileEData,
						suite.secondaryUser,
						[]string{"read"},
					).
					collection(),
			},
		},

		{
			name:     "FileAndFolderSeparatePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				newOneDriveCollection(suite.T(), rootPath, 1).
					withFolder(
						"b",
						suite.secondaryUser,
						[]string{"read"},
					).
					collection(),
				newOneDriveCollection(suite.T(), folderPath, 1).
					withFile(
						fileName,
						fileAData,
						suite.secondaryUser,
						[]string{"write"},
					).
					collection(),
			},
		},

		{
			name:     "FolderAndNoChildPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				newOneDriveCollection(suite.T(), rootPath, 1).
					withFolder(
						"b",
						suite.secondaryUser,
						[]string{"read"},
					).
					collection(),
				newOneDriveCollection(suite.T(), folderPath, 1).
					withFile(
						"test-file.txt",
						fileEData,
						"",
						nil,
					).
					collection(),
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
				newOneDriveCollection(
					suite.T(),
					[]string{
						"drives",
						driveID,
						"root:",
					},
					1,
				).
					withFile(
						"test-file.txt",
						fileAData,
						suite.secondaryUser,
						[]string{"write"},
					).
					collection(),
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

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsRestoreAndNoBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.connector.Service,
		suite.user,
	)

	test := restoreBackupInfoMultiVersion{
		service:       path.OneDriveService,
		resource:      Users,
		backupVersion: backup.Version,
		countMeta:     false,
		collectionsPrevious: []colInfo{
			newOneDriveCollection(
				suite.T(),
				[]string{
					"drives",
					driveID,
					"root:",
				},
				backup.Version,
			).
				withFile(
					"test-file.txt",
					fileAData,
					suite.secondaryUser,
					[]string{"write"},
				).
				withFolder(
					folderBName,
					suite.secondaryUser,
					[]string{"read"},
				).
				collection(),
			newOneDriveCollection(
				suite.T(),
				[]string{
					"drives",
					driveID,
					"root:",
					folderBName,
				},
				backup.Version,
			).
				withFile(
					"test-file.txt",
					fileEData,
					suite.secondaryUser,
					[]string{"read"},
				).
				withPermissions(
					suite.secondaryUser,
					[]string{"read"},
				).
				collection(),
		},
		collectionsLatest: []colInfo{
			newOneDriveCollection(
				suite.T(),
				[]string{
					"drives",
					driveID,
					"root:",
				},
				backup.Version,
			).
				withFile(
					"test-file.txt",
					fileAData,
					"",
					nil,
				).
				withFolder(
					folderBName,
					"",
					nil,
				).
				collection(),
			newOneDriveCollection(
				suite.T(),
				[]string{
					"drives",
					driveID,
					"root:",
					folderBName,
				},
				backup.Version,
			).
				withFile(
					"test-file.txt",
					fileEData,
					"",
					nil,
				).
				collection(),
		},
	}

	runRestoreBackupTestVersions(
		t,
		suite.acct,
		test,
		suite.connector.tenant,
		[]string{suite.user},
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{EnablePermissionsBackup: false},
		},
	)
}
