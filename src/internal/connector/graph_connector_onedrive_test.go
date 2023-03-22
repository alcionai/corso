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

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

// For any version post this(inclusive), we expect to be using IDs for
// permission instead of email
const versionPermissionSwitchedToID = version.OneDrive4DirIncludesPermissions

func getMetadata(fileName string, perm permData, permUseID bool) onedrive.Metadata {
	if len(perm.user) == 0 || len(perm.roles) == 0 ||
		perm.sharingMode != onedrive.SharingModeCustom {
		return onedrive.Metadata{
			FileName:    fileName,
			SharingMode: perm.sharingMode,
		}
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
	require.NoError(t, err, clues.ToCore(err))

	return itemInfo{
		name:      name,
		data:      serialized,
		lookupKey: lookupKey,
	}
}

func onedriveMetadata(
	t *testing.T,
	fileName, itemID, lookupKey string,
	perm permData,
	permUseID bool,
) itemInfo {
	t.Helper()

	testMeta := getMetadata(fileName, perm, permUseID)

	testMetaJSON, err := json.Marshal(testMeta)
	require.NoError(t, err, "marshalling metadata", clues.ToCore(err))

	return itemInfo{
		name:      itemID,
		data:      testMetaJSON,
		lookupKey: lookupKey,
	}
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

	// Cannot restore owner or empty permissions and so not testing them
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

func (c *onedriveCollection) withFile(name string, fileData []byte, perm permData) *onedriveCollection {
	switch c.backupVersion {
	case 0:
		// Lookups will occur using the most recent version of things so we need
		// the embedded file name to match that.
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name,
			name+onedrive.DataFileSuffix,
			fileData))

	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker,
		version.OneDrive4DirIncludesPermissions, version.OneDrive5DirMetaNoName:
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name+onedrive.DataFileSuffix,
			name+onedrive.DataFileSuffix,
			fileData))

		metadata := onedriveMetadata(
			c.t,
			"",
			name+onedrive.MetaFileSuffix,
			name+onedrive.MetaFileSuffix,
			perm,
			c.backupVersion >= versionPermissionSwitchedToID)
		c.items = append(c.items, metadata)
		c.aux = append(c.aux, metadata)

	case version.OneDrive6NameInMeta:
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name+onedrive.DataFileSuffix,
			name+onedrive.DataFileSuffix,
			fileData))

		metadata := onedriveMetadata(
			c.t,
			name,
			name+onedrive.MetaFileSuffix,
			name,
			perm,
			c.backupVersion >= versionPermissionSwitchedToID)
		c.items = append(c.items, metadata)
		c.aux = append(c.aux, metadata)

	default:
		assert.FailNowf(c.t, "bad backup version", "version %d", c.backupVersion)
	}

	return c
}

func (c *onedriveCollection) withFolder(name string, perm permData) *onedriveCollection {
	switch c.backupVersion {
	case 0, version.OneDrive4DirIncludesPermissions, version.OneDrive5DirMetaNoName,
		version.OneDrive6NameInMeta:
		return c

	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker:
		c.items = append(
			c.items,
			onedriveMetadata(
				c.t,
				"",
				name+onedrive.DirMetaFileSuffix,
				name+onedrive.DirMetaFileSuffix,
				perm,
				c.backupVersion >= versionPermissionSwitchedToID))

	default:
		assert.FailNowf(c.t, "bad backup version", "version %d", c.backupVersion)
	}

	return c
}

// withPermissions adds permissions to the folder represented by this
// onedriveCollection.
func (c *onedriveCollection) withPermissions(perm permData) *onedriveCollection {
	// These versions didn't store permissions for the folder or didn't store them
	// in the folder's collection.
	if c.backupVersion < version.OneDrive4DirIncludesPermissions {
		return c
	}

	name := c.pathElements[len(c.pathElements)-1]
	metaName := name

	if c.backupVersion >= version.OneDrive5DirMetaNoName {
		// We switched to just .dirmeta for metadata file names.
		metaName = ""
	}

	if name == "root:" {
		return c
	}

	metadata := onedriveMetadata(
		c.t,
		name,
		metaName+onedrive.DirMetaFileSuffix,
		metaName+onedrive.DirMetaFileSuffix,
		perm,
		c.backupVersion >= versionPermissionSwitchedToID)

	c.items = append(c.items, metadata)
	c.aux = append(c.aux, metadata)

	return c
}

type permData struct {
	user        string // user is only for older versions
	entityID    string
	roles       []string
	sharingMode onedrive.SharingMode
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

func testDataForInfo(t *testing.T, cols []onedriveColInfo, backupVersion int) []colInfo {
	var res []colInfo

	for _, c := range cols {
		onedriveCol := newOneDriveCollection(t, c.pathElements, backupVersion)

		for _, f := range c.files {
			onedriveCol.withFile(f.name, f.data, f.perms)
		}

		for _, d := range c.folders {
			onedriveCol.withFolder(d.name, d.perms)
		}

		onedriveCol.withPermissions(c.perms)

		res = append(res, onedriveCol.collection())
	}

	return res
}

type oneDriveSuite interface {
	tester.Suite
	Service() graph.Servicer
	Account() account.Account
	Tenant() string
	// Returns (username, user ID) for the user.
	PrimaryUser() (string, string)
	SecondaryUser() (string, string)
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
		),
	})
}

func (suite *GraphConnectorOneDriveIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	suite.connector = loadConnector(ctx, suite.T(), graph.HTTPClient(graph.NoTimeout()), Users)
	suite.user = tester.M365UserID(suite.T())
	suite.secondaryUser = tester.SecondaryM365UserID(suite.T())
	suite.acct = tester.NewM365Account(suite.T())

	user, err := suite.connector.Owners.Users().GetByID(ctx, suite.user)
	require.NoError(suite.T(), err, "fetching user", suite.user, clues.ToCore(err))
	suite.userID = ptr.Val(user.GetId())

	secondaryUser, err := suite.connector.Owners.Users().GetByID(ctx, suite.secondaryUser)
	require.NoError(suite.T(), err, "fetching user", suite.secondaryUser, clues.ToCore(err))
	suite.secondaryUserID = ptr.Val(secondaryUser.GetId())
}

func (suite *GraphConnectorOneDriveIntegrationSuite) Service() graph.Servicer {
	return suite.connector.Service
}

func (suite *GraphConnectorOneDriveIntegrationSuite) Account() account.Account {
	return suite.acct
}

func (suite *GraphConnectorOneDriveIntegrationSuite) Tenant() string {
	return suite.connector.tenant
}

func (suite *GraphConnectorOneDriveIntegrationSuite) PrimaryUser() (string, string) {
	return suite.user, suite.userID
}

func (suite *GraphConnectorOneDriveIntegrationSuite) SecondaryUser() (string, string) {
	return suite.secondaryUser, suite.secondaryUserID
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

// TestPermissionsRestoreAndNoBackup checks that even if permissions exist
// not setting EnablePermissionsBackup results in empty permissions. This test
// only needs to run on the current version.Backup because it's about backup
// behavior not restore behavior (restore behavior is checked in other tests).
func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsRestoreAndNoBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	userName, _ := suite.PrimaryUser()
	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	driveID := mustGetDefaultDriveID(
		t,
		ctx,
		suite.Service(),
		userName,
	)

	secondaryUserRead := permData{
		user:     secondaryUserName,
		entityID: secondaryUserID,
		roles:    readPerm,
	}

	secondaryUserWrite := permData{
		user:     secondaryUserName,
		entityID: secondaryUserID,
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
				).
				withFolder(
					folderBName,
					secondaryUserRead,
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
				).
				withPermissions(
					secondaryUserRead,
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
				).
				withFolder(
					folderBName,
					permData{},
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
				).
				// Call this to generate a meta file with the folder name that we can
				// check.
				withPermissions(
					permData{},
				).
				collection(),
		},
	}

	runRestoreBackupTestVersions(
		t,
		suite.Account(),
		test,
		suite.Tenant(),
		[]string{userName},
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{EnablePermissionsBackup: false},
		},
	)
}

type GraphConnectorOneDriveNightlySuite struct {
	tester.Suite
	connector       *GraphConnector
	user            string
	userID          string
	secondaryUser   string
	secondaryUserID string
	acct            account.Account
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
	ctx, flush := tester.NewContext()
	defer flush()

	suite.connector = loadConnector(ctx, suite.T(), graph.HTTPClient(graph.NoTimeout()), Users)
	suite.user = tester.M365UserID(suite.T())
	suite.secondaryUser = tester.SecondaryM365UserID(suite.T())
	suite.acct = tester.NewM365Account(suite.T())

	user, err := suite.connector.Owners.Users().GetByID(ctx, suite.user)
	require.NoError(suite.T(), err, "fetching user", suite.user, clues.ToCore(err))
	suite.userID = ptr.Val(user.GetId())

	secondaryUser, err := suite.connector.Owners.Users().GetByID(ctx, suite.secondaryUser)
	require.NoError(suite.T(), err, "fetching user", suite.secondaryUser, clues.ToCore(err))
	suite.secondaryUserID = ptr.Val(secondaryUser.GetId())
}

func (suite *GraphConnectorOneDriveNightlySuite) Service() graph.Servicer {
	return suite.connector.Service
}

func (suite *GraphConnectorOneDriveNightlySuite) Account() account.Account {
	return suite.acct
}

func (suite *GraphConnectorOneDriveNightlySuite) Tenant() string {
	return suite.connector.tenant
}

func (suite *GraphConnectorOneDriveNightlySuite) PrimaryUser() (string, string) {
	return suite.user, suite.userID
}

func (suite *GraphConnectorOneDriveNightlySuite) SecondaryUser() (string, string) {
	return suite.secondaryUser, suite.secondaryUserID
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
	// No reason why it couldn't work with previous versions, but this is when it
	// got introduced.
	testPermissionsInheritanceRestoreAndBackup(suite, version.OneDrive4DirIncludesPermissions)
}

func testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(
	suite oneDriveSuite,
	startVersion int,
) {
	ctx, flush := tester.NewContext()
	defer flush()

	userName, _ := suite.PrimaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.Service(),
		userName,
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

	cols := []onedriveColInfo{
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
	}

	expected := testDataForInfo(suite.T(), cols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input := testDataForInfo(t, cols, vn)

			testData := restoreBackupInfoMultiVersion{
				service:             path.OneDriveService,
				resource:            Users,
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{userName},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func testPermissionsRestoreAndBackup(suite oneDriveSuite, startVersion int) {
	ctx, flush := tester.NewContext()
	defer flush()

	userName, _ := suite.PrimaryUser()
	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.Service(),
		userName,
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

	cols := []onedriveColInfo{
		{
			pathElements: rootPath,
			files: []itemData{
				{
					// Test restoring a file that doesn't inherit permissions.
					name: fileName,
					data: fileAData,
					perms: permData{
						user:     secondaryUserName,
						entityID: secondaryUserID,
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
						user:     secondaryUserName,
						entityID: secondaryUserID,
						roles:    readPerm,
					},
				},
				{
					name: folderCName,
					perms: permData{
						user:     secondaryUserName,
						entityID: secondaryUserID,
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
						user:     secondaryUserName,
						entityID: secondaryUserID,
						roles:    readPerm,
					},
				},
			},
			folders: []itemData{
				{
					name: folderAName,
					perms: permData{
						user:     secondaryUserName,
						entityID: secondaryUserID,
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
						user:     secondaryUserName,
						entityID: secondaryUserID,
						roles:    readPerm,
					},
				},
			},
			perms: permData{
				user:     secondaryUserName,
				entityID: secondaryUserID,
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
						user:     secondaryUserName,
						entityID: secondaryUserID,
						roles:    writePerm,
					},
				},
			},
			perms: permData{
				user:     secondaryUserName,
				entityID: secondaryUserID,
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
				user:     secondaryUserName,
				entityID: secondaryUserID,
				roles:    readPerm,
			},
		},
	}

	expected := testDataForInfo(suite.T(), cols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input := testDataForInfo(t, cols, vn)

			testData := restoreBackupInfoMultiVersion{
				service:             path.OneDriveService,
				resource:            Users,
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{userName},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func testPermissionsBackupAndNoRestore(suite oneDriveSuite, startVersion int) {
	ctx, flush := tester.NewContext()
	defer flush()

	userName, _ := suite.PrimaryUser()
	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.Service(),
		userName,
	)

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
						user:     secondaryUserName,
						entityID: secondaryUserID,
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

	expected := testDataForInfo(suite.T(), expectedCols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input := testDataForInfo(t, inputCols, vn)

			testData := restoreBackupInfoMultiVersion{
				service:             path.OneDriveService,
				resource:            Users,
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{userName},
				control.Options{
					RestorePermissions: false,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

// This is similar to TestPermissionsRestoreAndBackup but tests purely
// for inheritance and that too only with newer versions
func testPermissionsInheritanceRestoreAndBackup(suite oneDriveSuite, startVersion int) {
	ctx, flush := tester.NewContext()
	defer flush()

	userName, _ := suite.PrimaryUser()
	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.Service(),
		userName,
	)

	folderAName := "custom"
	folderBName := "inherited"

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
	subfolderAPath := []string{
		"drives",
		driveID,
		"root:",
		folderAName,
		folderAName,
	}
	subfolderBPath := []string{
		"drives",
		driveID,
		"root:",
		folderAName,
		folderBName,
	}

	fileSet := []itemData{
		{
			name: "file-custom",
			data: fileAData,
			perms: permData{
				user:        secondaryUserName,
				entityID:    secondaryUserID,
				roles:       writePerm,
				sharingMode: onedrive.SharingModeCustom,
			},
		},
		{
			name: "file-inherited",
			data: fileAData,
			perms: permData{
				sharingMode: onedrive.SharingModeInherited,
			},
		},
	}

	// Here is what this test is testing
	// - custom-permission-folder
	//   - custom-permission-file
	//   - inherted-permission-file
	//   - custom-permission-folder
	// 	   - custom-permission-file
	// 	   - inherted-permission-file
	//   - inherted-permission-folder
	// 	   - custom-permission-file
	// 	   - inherted-permission-file

	cols := []onedriveColInfo{
		{
			pathElements: rootPath,
			files:        []itemData{},
			folders: []itemData{
				{
					name: folderAName,
				},
			},
		},
		{
			pathElements: folderAPath,
			files:        fileSet,
			folders: []itemData{
				{name: folderAName},
				{name: folderBName},
			},
			perms: permData{
				user:     secondaryUserName,
				entityID: secondaryUserID,
				roles:    readPerm,
			},
		},
		{
			pathElements: subfolderAPath,
			files:        fileSet,
			perms: permData{
				user:        secondaryUserName,
				entityID:    secondaryUserID,
				roles:       writePerm,
				sharingMode: onedrive.SharingModeCustom,
			},
		},
		{
			pathElements: subfolderBPath,
			files:        fileSet,
			perms: permData{
				sharingMode: onedrive.SharingModeInherited,
			},
		},
	}

	expected := testDataForInfo(suite.T(), cols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input := testDataForInfo(t, cols, vn)

			testData := restoreBackupInfoMultiVersion{
				service:             path.OneDriveService,
				resource:            Users,
				backupVersion:       vn,
				collectionsPrevious: input,
				collectionsLatest:   expected,
			}

			runRestoreBackupTestVersions(
				t,
				suite.Account(),
				testData,
				suite.Tenant(),
				[]string{userName},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}
