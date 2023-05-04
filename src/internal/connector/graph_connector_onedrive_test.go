package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

// For any version post this(inclusive), we expect to be using IDs for
// permission instead of email
const versionPermissionSwitchedToID = version.OneDrive4DirIncludesPermissions

func getMetadata(fileName string, perm permData, permUseID bool) metadata.Metadata {
	if len(perm.user) == 0 || len(perm.roles) == 0 ||
		perm.sharingMode != metadata.SharingModeCustom {
		return metadata.Metadata{
			FileName:    fileName,
			SharingMode: perm.sharingMode,
		}
	}

	// In case of permissions, the id will usually be same for same
	// user/role combo unless deleted and readded, but we have to do
	// this as we only have two users of which one is already taken.
	id := uuid.NewString()
	uperm := metadata.Permission{ID: id, Roles: perm.roles}

	if permUseID {
		uperm.EntityID = perm.entityID
	} else {
		uperm.Email = perm.user
	}

	testMeta := metadata.Metadata{
		FileName:    fileName,
		Permissions: []metadata.Permission{uperm},
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
	fileName          = "test-file.txt"
	folderAName       = "folder-a"
	folderBName       = "b"
	folderNamedFolder = "folder"
	rootFolder        = "root:"

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
	service path.ServiceType,
	pathElements []string,
	backupVersion int,
) *onedriveCollection {
	return &onedriveCollection{
		service:       service,
		pathElements:  pathElements,
		backupVersion: backupVersion,
		t:             t,
	}
}

type onedriveCollection struct {
	service       path.ServiceType
	pathElements  []string
	items         []itemInfo
	aux           []itemInfo
	backupVersion int
	t             *testing.T
}

func (c onedriveCollection) collection() colInfo {
	cat := path.FilesCategory
	if c.service == path.SharePointService {
		cat = path.LibrariesCategory
	}

	return colInfo{
		pathElements: c.pathElements,
		category:     cat,
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
			name+metadata.DataFileSuffix,
			fileData))

		// v1-5, early metadata design
	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker,
		version.OneDrive4DirIncludesPermissions, version.OneDrive5DirMetaNoName:
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name+metadata.DataFileSuffix,
			name+metadata.DataFileSuffix,
			fileData))

		md := onedriveMetadata(
			c.t,
			"",
			name+metadata.MetaFileSuffix,
			name+metadata.MetaFileSuffix,
			perm,
			c.backupVersion >= versionPermissionSwitchedToID)
		c.items = append(c.items, md)
		c.aux = append(c.aux, md)

		// v6+ current metadata design
	case version.OneDrive6NameInMeta, version.OneDrive7LocationRef, version.All8MigrateUserPNToID:
		c.items = append(c.items, onedriveItemWithData(
			c.t,
			name+metadata.DataFileSuffix,
			name+metadata.DataFileSuffix,
			fileData))

		md := onedriveMetadata(
			c.t,
			name,
			name+metadata.MetaFileSuffix,
			name,
			perm,
			c.backupVersion >= versionPermissionSwitchedToID)
		c.items = append(c.items, md)
		c.aux = append(c.aux, md)

	default:
		assert.FailNowf(c.t, "bad backup version", "version %d", c.backupVersion)
	}

	return c
}

func (c *onedriveCollection) withFolder(name string, perm permData) *onedriveCollection {
	switch c.backupVersion {
	case 0, version.OneDrive4DirIncludesPermissions, version.OneDrive5DirMetaNoName,
		version.OneDrive6NameInMeta, version.OneDrive7LocationRef, version.All8MigrateUserPNToID:
		return c

	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker:
		c.items = append(
			c.items,
			onedriveMetadata(
				c.t,
				"",
				name+metadata.DirMetaFileSuffix,
				name+metadata.DirMetaFileSuffix,
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

	if name == rootFolder {
		return c
	}

	md := onedriveMetadata(
		c.t,
		name,
		metaName+metadata.DirMetaFileSuffix,
		metaName+metadata.DirMetaFileSuffix,
		perm,
		c.backupVersion >= versionPermissionSwitchedToID)

	c.items = append(c.items, md)
	c.aux = append(c.aux, md)

	return c
}

type permData struct {
	user        string // user is only for older versions
	entityID    string
	roles       []string
	sharingMode metadata.SharingMode
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

func testDataForInfo(
	t *testing.T,
	service path.ServiceType,
	cols []onedriveColInfo,
	backupVersion int,
) []colInfo {
	var res []colInfo

	for _, c := range cols {
		onedriveCol := newOneDriveCollection(t, service, c.pathElements, backupVersion)

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
		d, err = service.Client().UsersById(resourceOwner).Drive().Get(ctx, nil)
	case path.SharePointService:
		d, err = service.Client().SitesById(resourceOwner).Drive().Get(ctx, nil)
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
//
// TODO(ashmrtn): SharePoint doesn't have permissions backup/restore enabled
// right now. Adjust the tests here when that is enabled so we have at least
// basic assurances that it's doing the right thing. We can leave the more
// extensive permissions tests to OneDrive as well.

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
	ctx, flush := tester.NewContext()
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
	require.NoError(suite.T(), err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(suite.T(), err, "fetching user", si.secondaryUser, clues.ToCore(err))
	si.secondaryUserID = ptr.Val(secondaryUser.GetId())

	tertiaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.tertiaryUser)
	require.NoError(suite.T(), err, "fetching user", si.tertiaryUser, clues.ToCore(err))
	si.tertiaryUserID = ptr.Val(tertiaryUser.GetId())

	suite.suiteInfo = si
}

func (suite *GraphConnectorSharePointIntegrationSuite) TestRestoreAndBackup_MultipleFilesAndFolders_NoPermissions() {
	testRestoreAndBackupMultipleFilesAndFoldersNoPermissions(suite, version.Backup)
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
	ctx, flush := tester.NewContext()
	defer flush()

	si := suiteInfoImpl{
		connector:     loadConnector(ctx, suite.T(), Users),
		user:          tester.M365UserID(suite.T()),
		secondaryUser: tester.SecondaryM365UserID(suite.T()),
		acct:          tester.NewM365Account(suite.T()),
		service:       path.OneDriveService,
		resourceType:  Users,
	}

	si.resourceOwner = si.user

	user, err := si.connector.Discovery.Users().GetByID(ctx, si.user)
	require.NoError(suite.T(), err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(suite.T(), err, "fetching user", si.secondaryUser, clues.ToCore(err))
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
	ctx, flush := tester.NewContext()
	defer flush()

	si := suiteInfoImpl{
		connector:     loadConnector(ctx, suite.T(), Users),
		user:          tester.M365UserID(suite.T()),
		secondaryUser: tester.SecondaryM365UserID(suite.T()),
		acct:          tester.NewM365Account(suite.T()),
		service:       path.OneDriveService,
		resourceType:  Users,
	}

	si.resourceOwner = si.user

	user, err := si.connector.Discovery.Users().GetByID(ctx, si.user)
	require.NoError(suite.T(), err, "fetching user", si.user, clues.ToCore(err))
	si.userID = ptr.Val(user.GetId())

	secondaryUser, err := si.connector.Discovery.Users().GetByID(ctx, si.secondaryUser)
	require.NoError(suite.T(), err, "fetching user", si.secondaryUser, clues.ToCore(err))
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
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

	rootPath := []string{
		"drives",
		driveID,
		rootFolder,
	}
	folderAPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
	}
	subfolderBPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
		folderBName,
	}
	subfolderAPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
		folderBName,
		folderAName,
	}
	folderBPath := []string{
		"drives",
		driveID,
		rootFolder,
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

	expected := testDataForInfo(suite.T(), suite.BackupService(), cols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input := testDataForInfo(t, suite.BackupService(), cols, vn)

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
	ctx, flush := tester.NewContext()
	defer flush()

	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

	fileName2 := "test-file2.txt"
	folderCName := "folder-c"

	rootPath := []string{
		"drives",
		driveID,
		rootFolder,
	}
	folderAPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
	}
	folderBPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderBName,
	}
	// For skipped test
	// subfolderAPath := []string{
	// 	"drives",
	// 	driveID,
	// 	rootFolder,
	// 	folderBName,
	// 	folderAName,
	// }
	folderCPath := []string{
		"drives",
		driveID,
		rootFolder,
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
						roles:    writePerm,
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
		// 	perms: permData{
		// 		user:     secondaryUserName,
		// 		entityID: secondaryUserID,
		// 		roles:    readPerm,
		// 	},
		// },
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

	expected := testDataForInfo(suite.T(), suite.BackupService(), cols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input := testDataForInfo(t, suite.BackupService(), cols, vn)

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
	ctx, flush := tester.NewContext()
	defer flush()

	secondaryUserName, secondaryUserID := suite.SecondaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

	inputCols := []onedriveColInfo{
		{
			pathElements: []string{
				"drives",
				driveID,
				rootFolder,
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
				rootFolder,
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

	expected := testDataForInfo(suite.T(), suite.BackupService(), expectedCols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input := testDataForInfo(t, suite.BackupService(), inputCols, vn)

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
	ctx, flush := tester.NewContext()
	defer flush()

	secondaryUserName, secondaryUserID := suite.SecondaryUser()
	tertiaryUserName, tertiaryUserID := suite.TertiaryUser()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

	folderAName := "custom"
	folderBName := "inherited"
	folderCName := "empty"

	rootPath := []string{
		"drives",
		driveID,
		rootFolder,
	}
	folderAPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
	}
	subfolderAAPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
		folderAName,
	}
	subfolderABPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
		folderBName,
	}
	subfolderACPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderAName,
		folderCName,
	}

	fileSet := []itemData{
		{
			name: "file-custom",
			data: fileAData,
			perms: permData{
				user:        secondaryUserName,
				entityID:    secondaryUserID,
				roles:       writePerm,
				sharingMode: metadata.SharingModeCustom,
			},
		},
		{
			name: "file-inherited",
			data: fileAData,
			perms: permData{
				sharingMode: metadata.SharingModeInherited,
			},
		},
		{
			name: "file-empty",
			data: fileAData,
			perms: permData{
				sharingMode: metadata.SharingModeCustom,
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

	cols := []onedriveColInfo{
		{
			pathElements: rootPath,
			files:        []itemData{},
			folders: []itemData{
				{name: folderAName},
			},
		},
		{
			pathElements: folderAPath,
			files:        fileSet,
			folders: []itemData{
				{name: folderAName},
				{name: folderBName},
				{name: folderCName},
			},
			perms: permData{
				user:     tertiaryUserName,
				entityID: tertiaryUserID,
				roles:    readPerm,
			},
		},
		{
			pathElements: subfolderAAPath,
			files:        fileSet,
			perms: permData{
				user:        tertiaryUserName,
				entityID:    tertiaryUserID,
				roles:       writePerm,
				sharingMode: metadata.SharingModeCustom,
			},
		},
		{
			pathElements: subfolderABPath,
			files:        fileSet,
			perms: permData{
				sharingMode: metadata.SharingModeInherited,
			},
		},
		{
			pathElements: subfolderACPath,
			files:        fileSet,
			perms: permData{
				sharingMode: metadata.SharingModeCustom,
			},
		},
	}

	expected := testDataForInfo(suite.T(), suite.BackupService(), cols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			// Ideally this can always be true or false and still
			// work, but limiting older versions to use emails so as
			// to validate that flow as well.
			input := testDataForInfo(t, suite.BackupService(), cols, vn)

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
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.BackupService(),
		suite.Service(),
		suite.BackupResourceOwner())

	rootPath := []string{
		"drives",
		driveID,
		rootFolder,
	}
	folderFolderPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderNamedFolder,
	}
	subfolderPath := []string{
		"drives",
		driveID,
		rootFolder,
		folderNamedFolder,
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
					name: folderNamedFolder,
				},
				{
					name: folderBName,
				},
			},
		},
		{
			pathElements: folderFolderPath,
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
			pathElements: subfolderPath,
			files: []itemData{
				{
					name: fileName,
					data: fileCData,
				},
			},
			folders: []itemData{
				{
					name: folderNamedFolder,
				},
			},
		},
	}

	expected := testDataForInfo(suite.T(), suite.BackupService(), cols, version.Backup)

	for vn := startVersion; vn <= version.Backup; vn++ {
		suite.Run(fmt.Sprintf("Version%d", vn), func() {
			t := suite.T()
			input := testDataForInfo(t, suite.BackupService(), cols, vn)

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
