package connector

import (
	"encoding/json"
	"testing"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

// For any version post this(inclusive), we expect to be using IDs for
// permission instead of email
const versionPermissionSwitchedToID = version.OneDrive4DirIncludesPermissions

var rootFolder = "root:"

type PermData struct {
	User        string // user is only for older versions
	EntityID    string
	Roles       []string
	SharingMode onedrive.SharingMode
}

type ItemData struct {
	Name  string
	Data  []byte
	Perms PermData
}

type OnedriveColInfo struct {
	PathElements []string
	Perms        PermData
	Files        []ItemData
	Folders      []ItemData
}

type ColInfo struct {
	// Elements (in order) for the path representing this collection. Should
	// only contain elements after the prefix that corso uses for the path. For
	// example, a collection for the Inbox folder in exchange mail would just be
	// "Inbox".
	PathElements []string
	Category     path.CategoryType
	Items        []itemInfo
	// auxItems are items that can be retrieved with Fetch but won't be returned
	// by Items(). These files do not directly participate in comparisosn at the
	// end of a test.
	AuxItems []itemInfo
}

type itemInfo struct {
	// lookupKey is a string that can be used to find this data from a set of
	// other data in the same collection. This key should be something that will
	// be the same before and after restoring the item in M365 and may not be
	// the M365 ID. When restoring items out of place, the item is assigned a
	// new ID making it unsuitable for a lookup key.
	lookupKey string
	name      string
	data      []byte
}

type onedriveCollection struct {
	service       path.ServiceType
	pathElements  []string
	items         []itemInfo
	aux           []itemInfo
	backupVersion int
	t             *testing.T
}

type testOneDriveData struct {
	FileName string `json:"fileName,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

func (c onedriveCollection) collection() ColInfo {
	cat := path.FilesCategory
	if c.service == path.SharePointService {
		cat = path.LibrariesCategory
	}

	return ColInfo{
		PathElements: c.pathElements,
		Category:     cat,
		Items:        c.items,
		AuxItems:     c.aux,
	}
}

func getMetadata(fileName string, perm PermData, permUseID bool) onedrive.Metadata {
	if len(perm.User) == 0 || len(perm.Roles) == 0 ||
		perm.SharingMode != onedrive.SharingModeCustom {
		return onedrive.Metadata{
			FileName:    fileName,
			SharingMode: perm.SharingMode,
		}
	}

	// In case of permissions, the id will usually be same for same
	// user/role combo unless deleted and readded, but we have to do
	// this as we only have two users of which one is already taken.
	id := uuid.NewString()
	uperm := onedrive.UserPermission{ID: id, Roles: perm.Roles}

	if permUseID {
		uperm.EntityID = perm.EntityID
	} else {
		uperm.Email = perm.User
	}

	testMeta := onedrive.Metadata{
		FileName:    fileName,
		Permissions: []onedrive.UserPermission{uperm},
	}

	return testMeta
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
	perm PermData,
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

func (c *onedriveCollection) withFile(name string, fileData []byte, perm PermData) *onedriveCollection {
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

func (c *onedriveCollection) withFolder(name string, perm PermData) *onedriveCollection {
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
func (c *onedriveCollection) withPermissions(perm PermData) *onedriveCollection {
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
	}
}

func DataForInfo(
	t *testing.T,
	service path.ServiceType,
	cols []OnedriveColInfo,
	backupVersion int,
) []ColInfo {
	var res []ColInfo

	for _, c := range cols {
		onedriveCol := newOneDriveCollection(t, service, c.PathElements, backupVersion)

		for _, f := range c.Files {
			onedriveCol.withFile(f.Name, f.Data, f.Perms)
		}

		for _, d := range c.Folders {
			onedriveCol.withFolder(d.Name, d.Perms)
		}

		onedriveCol.withPermissions(c.Perms)

		res = append(res, onedriveCol.collection())
	}

	return res
}

//-------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------------

func GetCollectionsAndExpected(
	t *testing.T,
	config ConfigInfo,
	testCollections []ColInfo,
	backupVersion int,
) (int, int, []data.RestoreCollection, map[string]map[string][]byte) {
	t.Helper()

	var (
		collections     []data.RestoreCollection
		expectedData    = map[string]map[string][]byte{}
		totalItems      = 0
		totalKopiaItems = 0
	)

	for _, owner := range config.ResourceOwners {
		numItems, kopiaItems, ownerCollections, userExpectedData := collectionsForInfo(
			t,
			config.Service,
			config.Tenant,
			owner,
			config.Dest,
			testCollections,
			backupVersion,
		)

		collections = append(collections, ownerCollections...)
		totalItems += numItems
		totalKopiaItems += kopiaItems

		maps.Copy(expectedData, userExpectedData)
	}

	return totalItems, totalKopiaItems, collections, expectedData
}
