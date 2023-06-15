package stub

import (
	"encoding/json"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/google/uuid"

	odConsts "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	m365Stub "github.com/alcionai/corso/src/internal/m365/stub"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
)

// For any version post this(inclusive), we expect to be using IDs for
// permission instead of email
const versionPermissionSwitchedToID = version.OneDrive4DirIncludesPermissions

func getMetadata(fileName string, perm PermData, permUseID bool) metadata.Metadata {
	if len(perm.User) == 0 || len(perm.Roles) == 0 ||
		perm.SharingMode != metadata.SharingModeCustom {
		return metadata.Metadata{
			FileName:    fileName,
			SharingMode: perm.SharingMode,
		}
	}

	// In case of permissions, the id will usually be same for same
	// user/role combo unless deleted and readded, but we have to do
	// this as we only have two users of which one is already taken.
	id := uuid.NewString()
	uperm := metadata.Permission{ID: id, Roles: perm.Roles}

	if permUseID {
		uperm.EntityID = perm.EntityID
	} else {
		uperm.Email = perm.User
	}

	testMeta := metadata.Metadata{
		FileName:    fileName,
		Permissions: []metadata.Permission{uperm},
	}

	return testMeta
}

type PermData struct {
	User        string // user is only for older versions
	EntityID    string
	Roles       []string
	SharingMode metadata.SharingMode
}

type ItemData struct {
	Name  string
	Data  []byte
	Perms PermData
}

type ColInfo struct {
	PathElements []string
	Perms        PermData
	Files        []ItemData
	Folders      []ItemData
}

type collection struct {
	Service       path.ServiceType
	PathElements  []string
	Items         []m365Stub.ItemInfo
	Aux           []m365Stub.ItemInfo
	BackupVersion int
}

func (c collection) ColInfo() m365Stub.ColInfo {
	cat := path.FilesCategory
	if c.Service == path.SharePointService {
		cat = path.LibrariesCategory
	}

	return m365Stub.ColInfo{
		PathElements: c.PathElements,
		Category:     cat,
		Items:        c.Items,
		AuxItems:     c.Aux,
	}
}

func NewCollection(
	service path.ServiceType,
	PathElements []string,
	backupVersion int,
) *collection {
	return &collection{
		Service:       service,
		PathElements:  PathElements,
		BackupVersion: backupVersion,
	}
}

func DataForInfo(
	service path.ServiceType,
	cols []ColInfo,
	backupVersion int,
) ([]m365Stub.ColInfo, error) {
	var (
		res []m365Stub.ColInfo
		err error
	)

	for _, c := range cols {
		onedriveCol := NewCollection(service, c.PathElements, backupVersion)

		for _, f := range c.Files {
			_, err = onedriveCol.withFile(f.Name, f.Data, f.Perms)
			if err != nil {
				return res, err
			}
		}

		for _, d := range c.Folders {
			_, err = onedriveCol.withFolder(d.Name, d.Perms)
			if err != nil {
				return res, err
			}
		}

		_, err = onedriveCol.withPermissions(c.Perms)
		if err != nil {
			return res, err
		}

		res = append(res, onedriveCol.ColInfo())
	}

	return res, nil
}

func (c *collection) withFile(name string, fileData []byte, perm PermData) (*collection, error) {
	switch c.BackupVersion {
	case 0:
		// Lookups will occur using the most recent version of things so we need
		// the embedded file name to match that.
		item, err := FileWithData(
			name,
			name+metadata.DataFileSuffix,
			fileData)
		if err != nil {
			return c, err
		}

		c.Items = append(c.Items, item)

		// v1-5, early metadata design
	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker,
		version.OneDrive4DirIncludesPermissions, version.OneDrive5DirMetaNoName:
		items, err := FileWithData(
			name+metadata.DataFileSuffix,
			name+metadata.DataFileSuffix,
			fileData)
		if err != nil {
			return c, err
		}

		c.Items = append(c.Items, items)

		md, err := ItemWithMetadata(
			"",
			name+metadata.MetaFileSuffix,
			name+metadata.MetaFileSuffix,
			perm,
			c.BackupVersion >= versionPermissionSwitchedToID)
		if err != nil {
			return c, err
		}

		c.Items = append(c.Items, md)
		c.Aux = append(c.Aux, md)

		// v6+ current metadata design
	case version.OneDrive6NameInMeta, version.OneDrive7LocationRef, version.All8MigrateUserPNToID:
		item, err := FileWithData(
			name+metadata.DataFileSuffix,
			name+metadata.DataFileSuffix,
			fileData)
		if err != nil {
			return c, err
		}

		c.Items = append(c.Items, item)

		md, err := ItemWithMetadata(
			name,
			name+metadata.MetaFileSuffix,
			name,
			perm,
			c.BackupVersion >= versionPermissionSwitchedToID)
		if err != nil {
			return c, err
		}

		c.Items = append(c.Items, md)
		c.Aux = append(c.Aux, md)

	default:
		return c, clues.New(fmt.Sprintf("bad backup version. version %d", c.BackupVersion))
	}

	return c, nil
}

func (c *collection) withFolder(name string, perm PermData) (*collection, error) {
	switch c.BackupVersion {
	case 0, version.OneDrive4DirIncludesPermissions, version.OneDrive5DirMetaNoName,
		version.OneDrive6NameInMeta, version.OneDrive7LocationRef, version.All8MigrateUserPNToID:
		return c, nil

	case version.OneDrive1DataAndMetaFiles, 2, version.OneDrive3IsMetaMarker:
		item, err := ItemWithMetadata(
			"",
			name+metadata.DirMetaFileSuffix,
			name+metadata.DirMetaFileSuffix,
			perm,
			c.BackupVersion >= versionPermissionSwitchedToID)

		c.Items = append(c.Items, item)

		if err != nil {
			return c, err
		}

	default:
		return c, clues.New(fmt.Sprintf("bad backup version.version %d", c.BackupVersion))
	}

	return c, nil
}

// withPermissions adds permissions to the folder represented by this
// onedriveCollection.
func (c *collection) withPermissions(perm PermData) (*collection, error) {
	// These versions didn't store permissions for the folder or didn't store them
	// in the folder's collection.
	if c.BackupVersion < version.OneDrive4DirIncludesPermissions {
		return c, nil
	}

	name := c.PathElements[len(c.PathElements)-1]
	metaName := name

	if c.BackupVersion >= version.OneDrive5DirMetaNoName {
		// We switched to just .dirmeta for metadata file names.
		metaName = ""
	}

	if name == odConsts.RootPathDir {
		return c, nil
	}

	md, err := ItemWithMetadata(
		name,
		metaName+metadata.DirMetaFileSuffix,
		metaName+metadata.DirMetaFileSuffix,
		perm,
		c.BackupVersion >= versionPermissionSwitchedToID)
	if err != nil {
		return c, err
	}

	c.Items = append(c.Items, md)
	c.Aux = append(c.Aux, md)

	return c, err
}

type FileData struct {
	FileName string `json:"fileName,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

func FileWithData(
	name, lookupKey string,
	fileData []byte,
) (m365Stub.ItemInfo, error) {
	content := FileData{
		FileName: lookupKey,
		Data:     fileData,
	}

	serialized, err := json.Marshal(content)
	if err != nil {
		return m365Stub.ItemInfo{}, clues.Stack(err)
	}

	return m365Stub.ItemInfo{
		Name:      name,
		Data:      serialized,
		LookupKey: lookupKey,
	}, nil
}

func ItemWithMetadata(
	fileName, itemID, lookupKey string,
	perm PermData,
	permUseID bool,
) (m365Stub.ItemInfo, error) {
	testMeta := getMetadata(fileName, perm, permUseID)

	testMetaJSON, err := json.Marshal(testMeta)
	if err != nil {
		return m365Stub.ItemInfo{}, clues.Wrap(err, "marshalling metadata")
	}

	return m365Stub.ItemInfo{
		Name:      itemID,
		Data:      testMetaJSON,
		LookupKey: lookupKey,
	}, nil
}
