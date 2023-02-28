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

	"github.com/alcionai/corso/src/internal/connector/discovery"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

func getMetadata(fileName string, perm permData, permUseID bool) onedrive.Metadata {
	if len(perm.user) == 0 || len(perm.roles) == 0 {
		return onedrive.Metadata{FileName: fileName}
	}

	id := base64.StdEncoding.EncodeToString([]byte(perm.user + strings.Join(perm.roles, "+")))
	uperm := onedrive.UserPermission{ID: id, Roles: perm.roles}

	if permUseID {
		uperm.EntityID = perm.entityID
	} else {
		uperm.Email = perm.user
	}

	testMeta := onedrive.Metadata{
		FileName:    fileName,
		Permissions: []onedrive.UserPermission{uperm},
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
	perm permData,
	permUseID bool,
) itemInfo {
	t.Helper()

	testMeta := getMetadata(fileName, perm, permUseID)

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
	connector       *GraphConnector
	user            string
	userID          string
	secondaryUser   string
	secondaryUserID string
	acct            account.Account
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

	// Not using m365.Users to avoid cyclic imports
	users, err := discovery.Users(ctx, suite.connector.Owners.Users(), fault.New(true))
	require.NoError(suite.T(), err, "fetching users")

	for _, u := range users {
		if *u.GetUserPrincipalName() == suite.user {
			suite.userID = *u.GetId()
		} else if *u.GetUserPrincipalName() == suite.secondaryUser {
			suite.secondaryUserID = *u.GetId()
		}
	}

	if suite.userID == "" {
		require.FailNowf(suite.T(), "unable to find user id", "user %d", suite.user)
	}

	if suite.secondaryUserID == "" {
		require.FailNowf(suite.T(), "unable to find user id", "user %d", suite.secondaryUser)
	}

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
	perm permData,
	permUseID bool,
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

	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker, version.OneDrive4DirIncludesPermissions:
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name+onedrive.DataFileSuffix,
			name+onedrive.DataFileSuffix,
			fileData))

		metadata := onedriveMetadata(
			c.t,
			"",
			name+onedrive.MetaFileSuffix,
			perm,
			permUseID)
		c.items = append(c.items, metadata)
		c.aux = append(c.aux, metadata)

	default:
		assert.FailNowf(c.t, "bad backup version", "version %d", c.backupVersion)
	}

	return c
}

func (c *onedriveCollection) withFolder(
	name string,
	perm permData,
	permUseID bool,
) *onedriveCollection {
	switch c.backupVersion {
	case 0, version.OneDrive4DirIncludesPermissions:
		return c

	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker:
		c.items = append(
			c.items,
			onedriveMetadata(
				c.t,
				"",
				name+onedrive.DirMetaFileSuffix,
				perm,
				permUseID))

	default:
		assert.FailNowf(c.t, "bad backup version", "version %d", c.backupVersion)
	}

	return c
}

// withPermissions adds permissions to the folder represented by this
// onedriveCollection.
func (c *onedriveCollection) withPermissions(
	perm permData,
	permUseID bool,
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
		perm,
		permUseID)

	c.items = append(c.items, metadata)
	c.aux = append(c.aux, metadata)

	return c
}

type permData struct {
	user     string // user is only for older versions
	entityID string
	roles    []string
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
	// Version this test first be run for. Will run from
	// [startVersion, version.Backup] inclusive.
	startVersion int
	cols         []onedriveColInfo
}

func testDataForInfo(t *testing.T, cols []onedriveColInfo, backupVersion int, permUseID bool) []colInfo {
	var res []colInfo

	for _, c := range cols {
		onedriveCol := newOneDriveCollection(t, c.pathElements, backupVersion)

		for _, f := range c.files {
			onedriveCol.withFile(f.name, f.data, f.perms, permUseID)
		}

		for _, d := range c.folders {
			onedriveCol.withFolder(d.name, d.perms, permUseID)
		}

		onedriveCol.withPermissions(c.perms, permUseID)

		res = append(res, onedriveCol.collection())
	}

	return res
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
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

	test := onedriveTest{
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
	}

	expected := testDataForInfo(suite.T(), test.cols, version.Backup, true)

	for vn := test.startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input := testDataForInfo(t, test.cols, vn, vn > version.OneDrive3IsMetaMarker)

			testData := restoreBackupInfoMultiVersion{
				service:             path.OneDriveService,
				resource:            Users,
				backupVersion:       vn,
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

	fileName2 := "test-file2.txt"
	folderCName := "folder-c"

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
	folderBPath := []string{
		"drives",
		driveID,
		"root:",
		folderBName,
	}
	subfolderAPath := []string{
		"drives",
		driveID,
		"root:",
		folderBName,
		folderAName,
	}
	folderCPath := []string{
		"drives",
		driveID,
		"root:",
		folderCName,
	}

	startVersion := version.OneDrive1DataAndMetaFiles

	test := onedriveTest{
		startVersion: startVersion,
		cols: []onedriveColInfo{
			{
				pathElements: rootPath,
				files: []itemData{
					{
						// Test restoring a file that doesn't inherit permissions.
						name: fileName,
						data: fileAData,
						perms: permData{
							user:     suite.secondaryUser,
							entityID: suite.secondaryUserID,
							roles:    writePerm,
						},
					},
					{
						// Test restoring a file that doesn't inherit permissions and has
						// no permissions.
						name: fileName2,
						data: fileBData,
					},
				},
				folders: []itemData{
					{
						name: folderBName,
					},
					{
						name: folderAName,
						perms: permData{
							user:     suite.secondaryUser,
							entityID: suite.secondaryUserID,
							roles:    readPerm,
						},
					},
					{
						name: folderCName,
						perms: permData{
							user:     suite.secondaryUser,
							entityID: suite.secondaryUserID,
							roles:    readPerm,
						},
					},
				},
			},
			{
				pathElements: folderBPath,
				files: []itemData{
					{
						// Test restoring a file in a non-root folder that doesn't inherit
						// permissions.
						name: fileName,
						data: fileBData,
						perms: permData{
							user:     suite.secondaryUser,
							entityID: suite.secondaryUserID,
							roles:    readPerm,
						},
					},
				},
				folders: []itemData{
					{
						name: folderAName,
						perms: permData{
							user:     suite.secondaryUser,
							entityID: suite.secondaryUserID,
							roles:    readPerm,
						},
					},
				},
			},
			{
				// Tests a folder that has permissions with an item in the folder with
				// the same permissions.
				pathElements: subfolderAPath,
				files: []itemData{
					{
						name: fileName,
						data: fileDData,
						perms: permData{
							user:     suite.secondaryUser,
							entityID: suite.secondaryUserID,
							roles:    readPerm,
						},
					},
				},
				perms: permData{
					user:     suite.secondaryUser,
					entityID: suite.secondaryUserID,
					roles:    readPerm,
				},
			},
			{
				// Tests a folder that has permissions with an item in the folder with
				// the different permissions.
				pathElements: folderAPath,
				files: []itemData{
					{
						name: fileName,
						data: fileEData,
						perms: permData{
							user:     suite.secondaryUser,
							entityID: suite.secondaryUserID,
							roles:    writePerm,
						},
					},
				},
				perms: permData{
					user:     suite.secondaryUser,
					entityID: suite.secondaryUserID,
					roles:    readPerm,
				},
			},
			{
				// Tests a folder that has permissions with an item in the folder with
				// no permissions.
				pathElements: folderCPath,
				files: []itemData{
					{
						name: fileName,
						data: fileAData,
					},
				},
				perms: permData{
					user:     suite.secondaryUser,
					entityID: suite.secondaryUserID,
					roles:    readPerm,
				},
			},
		},
	}

	expected := testDataForInfo(suite.T(), test.cols, version.Backup, true)

	for vn := test.startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input := testDataForInfo(t, test.cols, vn, vn > version.OneDrive3IsMetaMarker)

			testData := restoreBackupInfoMultiVersion{
				service:             path.OneDriveService,
				resource:            Users,
				backupVersion:       vn,
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

	startVersion := version.OneDrive1DataAndMetaFiles

	inputCols := []onedriveColInfo{
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
						user:     suite.secondaryUser,
						entityID: suite.secondaryUserID,
						roles:    writePerm,
					},
				},
			},
		},
	}

	expectedCols := []onedriveColInfo{
		{
			pathElements: []string{
				"drives",
				driveID,
				"root:",
			},
			files: []itemData{
				{
					// No permissions on the output since they weren't restored.
					name: fileName,
					data: fileAData,
				},
			},
		},
	}

	expected := testDataForInfo(suite.T(), expectedCols, version.Backup, true)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input := testDataForInfo(t, inputCols, vn, vn > version.OneDrive3IsMetaMarker)

			testData := restoreBackupInfoMultiVersion{
				service:             path.OneDriveService,
				resource:            Users,
				backupVersion:       vn,
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
					RestorePermissions: false,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

// TestPermissionsRestoreAndNoBackup checks that even if permissions exist
// not setting EnablePermissionsBackup results in empty permissions. This test
// only needs to run on the current version.Backup because it's about backup
// behavior not restore behavior (restore behavior is checked in other tests).
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

	secondaryUserRead := permData{
		user:     suite.secondaryUser,
		entityID: suite.secondaryUserID,
		roles:    readPerm,
	}

	secondaryUserWrite := permData{
		user:     suite.secondaryUser,
		entityID: suite.secondaryUserID,
		roles:    writePerm,
	}

	test := restoreBackupInfoMultiVersion{
		service:       path.OneDriveService,
		resource:      Users,
		backupVersion: version.Backup,
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
					secondaryUserWrite,
					true,
				).
				withFolder(
					folderBName,
					secondaryUserRead,
					true,
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
					secondaryUserRead,
					true,
				).
				withPermissions(
					secondaryUserRead,
					true,
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
					permData{},
					true,
				).
				withFolder(
					folderBName,
					permData{},
					true,
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
					permData{},
					true,
				).
				// Call this to generate a meta file with the folder name that we can
				// check.
				withPermissions(
					permData{},
					true,
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
