package connector

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
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
	tester.Suite
	connector     *GraphConnector
	user          string
	secondaryUser string
	acct          account.Account
}

func TestGraphConnectorOneDriveIntegrationSuite(t *testing.T) {
	suite.Run(t, &GraphConnectorOneDriveIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
			tester.CorsoGraphConnectorTests,
			tester.CorsoGraphConnectorOneDriveTests),
	})
}

func (suite *GraphConnectorOneDriveIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

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

	writePerm = []string{"write"}
	readPerm  = []string{"read"}
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

	case 1, 2, 3, 4:
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
		assert.FailNowf(c.t, "bad backup version", "version %d", c.backupVersion)
	}

	return c
}

func (c *onedriveCollection) withFolder(
	name string,
	user string,
	roles []string,
) *onedriveCollection {
	switch c.backupVersion {
	case 0, 4:
		return c

	case 1, 2, 3:
		c.items = append(
			c.items,
			onedriveMetadata(
				c.t,
				"",
				name+onedrive.DirMetaFileSuffix,
				user,
				roles))

	default:
		assert.FailNowf(c.t, "bad backup version", "version %d", c.backupVersion)
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
	if c.backupVersion < version.OneDrive4DirIncludesPermissions {
		return c
	}

	name := c.pathElements[len(c.pathElements)-1]

	if name == "root:" {
		return c
	}

	metadata := onedriveMetadata(
		c.t,
		name,
		name+onedrive.DirMetaFileSuffix,
		user,
		roles)

	c.items = append(c.items, metadata)
	c.aux = append(c.aux, metadata)

	return c
}

type permData struct {
	user  string
	roles []string
}

type itemData struct {
	name  string
	data  []byte
	perms permData
}

type onedriveColInfo struct {
	pathElements []string
	perms        permData
	files        []itemData
	folders      []itemData
}

type onedriveTest struct {
	name string
	// Version this test first be run for. Will run from
	// [startVersion, version.Backup] inclusive.
	startVersion int
	cols         []onedriveColInfo
}

func testDataForInfo(t *testing.T, cols []onedriveColInfo, backupVersion int) []colInfo {
	var res []colInfo

	for _, c := range cols {
		onedriveCol := newOneDriveCollection(t, c.pathElements, backupVersion)

		for _, f := range c.files {
			onedriveCol.withFile(f.name, f.data, f.perms.user, f.perms.roles)
		}

		for _, d := range c.folders {
			onedriveCol.withFolder(d.name, d.perms.user, d.perms.roles)
		}

		onedriveCol.withPermissions(c.perms.user, c.perms.roles)

		res = append(res, onedriveCol.collection())
	}

	return res
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders() {
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
		folderAName,
	}
	subfolderBPath := []string{
		"drives",
		driveID,
		"root:",
		folderAName,
		folderBName,
	}
	subfolderAPath := []string{
		"drives",
		driveID,
		"root:",
		folderAName,
		folderBName,
		folderAName,
	}
	folderBPath := []string{
		"drives",
		driveID,
		"root:",
		folderBName,
	}

	table := []onedriveTest{
		{
			name:         "WithMetadata",
			startVersion: 1,
			cols: []onedriveColInfo{
				{
					pathElements: rootPath,
					files: []itemData{
						{
							name: fileName,
							data: fileAData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: writePerm,
							},
						},
					},
					folders: []itemData{
						{
							name: folderBName,
							perms: permData{
								user:  suite.secondaryUser,
								roles: readPerm,
							},
						},
					},
				},
				{
					pathElements: folderBPath,
					perms: permData{
						user:  suite.secondaryUser,
						roles: readPerm,
					},
					files: []itemData{
						{
							name: fileName,
							data: fileEData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: readPerm,
							},
						},
					},
				},
			},
		},
		{
			name:         "NoMetadata",
			startVersion: 0,
			cols: []onedriveColInfo{
				{
					pathElements: rootPath,
					files: []itemData{
						{
							name: fileName,
							data: fileAData,
						},
					},
					folders: []itemData{
						{
							name: folderAName,
						},
						{
							name: folderBName,
						},
					},
				},
				{
					pathElements: folderAPath,
					files: []itemData{
						{
							name: fileName,
							data: fileBData,
						},
					},
					folders: []itemData{
						{
							name: folderBName,
						},
					},
				},
				{
					pathElements: subfolderBPath,
					files: []itemData{
						{
							name: fileName,
							data: fileCData,
						},
					},
					folders: []itemData{
						{
							name: folderAName,
						},
					},
				},
				{
					pathElements: subfolderAPath,
					files: []itemData{
						{
							name: fileName,
							data: fileDData,
						},
					},
				},
				{
					pathElements: folderBPath,
					files: []itemData{
						{
							name: fileName,
							data: fileEData,
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		expected := testDataForInfo(suite.T(), test.cols, version.Backup)

		for vn := test.startVersion; vn <= version.Backup; vn++ {
			suite.Run(fmt.Sprintf("%s_Version%d", test.name, vn), func() {
				t := suite.T()
				input := testDataForInfo(t, test.cols, vn)

				testData := restoreBackupInfoMultiVersion{
					service:             path.OneDriveService,
					resource:            Users,
					backupVersion:       vn,
					countMeta:           vn == 0,
					collectionsPrevious: input,
					collectionsLatest:   expected,
				}

				runRestoreBackupTestVersions(
					t,
					suite.acct,
					testData,
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
		folderBName,
	}

	startVersion := 1

	table := []onedriveTest{
		{
			name:         "FilePermissionsRestore",
			startVersion: startVersion,
			cols: []onedriveColInfo{
				{
					pathElements: rootPath,
					files: []itemData{
						{
							name: fileName,
							data: fileAData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: writePerm,
							},
						},
					},
				},
			},
		},
		{
			name:         "FileInsideFolderPermissionsRestore",
			startVersion: startVersion,
			cols: []onedriveColInfo{
				{
					pathElements: rootPath,
					files: []itemData{
						{
							name: fileName,
							data: fileAData,
						},
					},
					folders: []itemData{
						{
							name: folderBName,
						},
					},
				},
				{
					pathElements: folderPath,
					files: []itemData{
						{
							name: fileName,
							data: fileEData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: readPerm,
							},
						},
					},
				},
			},
		},
		{
			name:         "FilesAndFolderPermissionsRestore",
			startVersion: startVersion,
			cols: []onedriveColInfo{
				{
					pathElements: rootPath,
					files: []itemData{
						{
							name: fileName,
							data: fileAData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: writePerm,
							},
						},
					},
					folders: []itemData{
						{
							name: folderBName,
							perms: permData{
								user:  suite.secondaryUser,
								roles: readPerm,
							},
						},
					},
				},
				{
					pathElements: folderPath,
					files: []itemData{
						{
							name: fileName,
							data: fileEData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: readPerm,
							},
						},
					},
					perms: permData{
						user:  suite.secondaryUser,
						roles: readPerm,
					},
				},
			},
		},
		{
			name:         "FilesAndFolderSeparatePermissionsRestore",
			startVersion: startVersion,
			cols: []onedriveColInfo{
				{
					pathElements: rootPath,
					folders: []itemData{
						{
							name: folderBName,
							perms: permData{
								user:  suite.secondaryUser,
								roles: readPerm,
							},
						},
					},
				},
				{
					pathElements: folderPath,
					files: []itemData{
						{
							name: fileName,
							data: fileEData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: writePerm,
							},
						},
					},
					perms: permData{
						user:  suite.secondaryUser,
						roles: readPerm,
					},
				},
			},
		},
		{
			name:         "FolderAndNoChildPermissionsRestore",
			startVersion: startVersion,
			cols: []onedriveColInfo{
				{
					pathElements: rootPath,
					folders: []itemData{
						{
							name: folderBName,
							perms: permData{
								user:  suite.secondaryUser,
								roles: readPerm,
							},
						},
					},
				},
				{
					pathElements: folderPath,
					files: []itemData{
						{
							name: fileName,
							data: fileEData,
						},
					},
					perms: permData{
						user:  suite.secondaryUser,
						roles: readPerm,
					},
				},
			},
		},
	}

	for _, test := range table {
		expected := testDataForInfo(suite.T(), test.cols, version.Backup)

		for vn := test.startVersion; vn <= version.Backup; vn++ {
			suite.Run(fmt.Sprintf("%s_Version%d", test.name, vn), func() {
				t := suite.T()
				input := testDataForInfo(t, test.cols, vn)

				testData := restoreBackupInfoMultiVersion{
					service:             path.OneDriveService,
					resource:            Users,
					backupVersion:       vn,
					countMeta:           vn == 0,
					collectionsPrevious: input,
					collectionsLatest:   expected,
				}

				runRestoreBackupTestVersions(
					t,
					suite.acct,
					testData,
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
}

// TODO(ashmrtn): What this test is supposed to do needs investigated. It
// doesn't seem to do what it says.
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

	startVersion := 1

	table := []onedriveTest{
		{
			startVersion: startVersion,
			cols: []onedriveColInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					files: []itemData{
						{
							name: fileName,
							data: fileAData,
							perms: permData{
								user:  suite.secondaryUser,
								roles: writePerm,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		expected := testDataForInfo(suite.T(), test.cols, version.Backup)

		for vn := test.startVersion; vn <= version.Backup; vn++ {
			suite.Run(fmt.Sprintf("Version%d", vn), func() {
				t := suite.T()
				input := testDataForInfo(t, test.cols, vn)

				testData := restoreBackupInfoMultiVersion{
					service:             path.OneDriveService,
					resource:            Users,
					backupVersion:       vn,
					countMeta:           vn == 0,
					collectionsPrevious: input,
					collectionsLatest:   expected,
				}

				runRestoreBackupTestVersions(
					t,
					suite.acct,
					testData,
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
		backupVersion: version.Backup,
		countMeta:     false,
		collectionsPrevious: []colInfo{
			newOneDriveCollection(
				suite.T(),
				[]string{
					"drives",
					driveID,
					"root:",
				},
				version.Backup,
			).
				withFile(
					fileName,
					fileAData,
					suite.secondaryUser,
					writePerm,
				).
				withFolder(
					folderBName,
					suite.secondaryUser,
					readPerm,
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
				version.Backup,
			).
				withFile(
					fileName,
					fileEData,
					suite.secondaryUser,
					readPerm,
				).
				withPermissions(
					suite.secondaryUser,
					readPerm,
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
				version.Backup,
			).
				withFile(
					fileName,
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
				version.Backup,
			).
				withFile(
					fileName,
					fileEData,
					"",
					nil,
				).
				// Call this to generate a meta file with the folder name that we can
				// check.
				withPermissions(
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
