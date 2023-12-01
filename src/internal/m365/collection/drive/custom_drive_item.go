package drive

import (
	"strings"
	"time"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Replica of models.DriveItemable
type CorsoDriveItemable interface {
	GetId() *string
	GetName() *string
	GetSize() *int64
	GetFile() fileDriveItemable
	GetFolder() folderDriveItemable
	GetPackageEscaped() packageDriveItemable
	GetParentReference() parentReferenceable
	GetAdditionalData() map[string]interface{}
	SetParentReference(parentReferenceable)
	GetShared() itemSharedable
	GetCreatedBy() itemIdentitySetable
	GetCreatedDateTime() *time.Time
	GetLastModifiedDateTime() *time.Time
	GetMalware() malwareable
	GetDeleted() deletedable
	GetRoot() itemRootable
}

type fileDriveItemable interface {
	GetMimeType() *string
}
type folderDriveItemable interface{}
type packageDriveItemable interface{}
type parentReferenceable interface {
	GetPath() *string
	GetId() *string
	GetName() *string
	GetDriveId() *string
}
type itemSharedable interface{}
type malwareable interface{}
type deletedable interface{}
type itemRootable interface{}
type itemIdentitySetable interface {
	GetUser() itemUserable
}
type itemUserable interface {
	GetAdditionalData() map[string]interface{}
}

// Concrete implementations
type folderDriveItem struct {
	isFolder bool
}

type fileDriveItem struct {
	isFile   bool
	mimeType string
}

func (fdi *fileDriveItem) GetMimeType() *string {
	return &fdi.mimeType
}

type packageDriveItem struct {
	isPackage bool
}

type parentReference struct {
	path    string
	id      string
	name    string
	driveId string
}

func (pr *parentReference) GetPath() *string {
	return &pr.path
}

func (pr *parentReference) GetId() *string {
	return &pr.id
}

func (pr *parentReference) GetName() *string {
	return &pr.name
}

func (pr *parentReference) GetDriveId() *string {
	return &pr.driveId
}

type itemShared struct {
	isShared bool
}

type itemMalware struct {
	isMalware bool
}

type itemDeleted struct {
	isDeleted bool
}

type itemRoot struct {
	isRoot bool
}

type itemIdentitySet struct {
	user itemUserable
}

func (iis *itemIdentitySet) GetUser() itemUserable {
	return iis.user
}

type itemUser struct {
	additionalData map[string]interface{}
}

func (iu *itemUser) GetAdditionalData() map[string]interface{} {
	return iu.additionalData
}

type CorsoDriveItem struct {
	ID                   string
	Name                 string
	Size                 int64
	File                 fileDriveItemable
	Folder               folderDriveItemable
	Package              packageDriveItemable
	AdditionalData       map[string]interface{}
	ParentReference      parentReferenceable
	Shared               itemSharedable
	CreatedBy            itemIdentitySetable
	CreatedDateTime      time.Time
	LastModifiedDateTime time.Time
	Malware              malwareable
	Deleted              deletedable
	Root                 itemRootable
}

func (c *CorsoDriveItem) GetId() *string {
	return &c.ID
}

func (c *CorsoDriveItem) GetName() *string {
	return &c.Name
}

func (c *CorsoDriveItem) GetSize() *int64 {
	return &c.Size
}

func (c *CorsoDriveItem) GetFile() fileDriveItemable {
	return c.File
}

func (c *CorsoDriveItem) GetFolder() folderDriveItemable {
	return c.Folder
}

func (c *CorsoDriveItem) GetPackageEscaped() packageDriveItemable {
	return c.Package
}

func (c *CorsoDriveItem) GetParentReference() parentReferenceable {
	return c.ParentReference
}

// TODO: Should we only support GETs?
func (c *CorsoDriveItem) SetParentReference(parent parentReferenceable) {
	c.ParentReference = parent
}

func (c *CorsoDriveItem) GetAdditionalData() map[string]interface{} {
	return c.AdditionalData
}

func (c *CorsoDriveItem) GetShared() itemSharedable {
	return c.Shared
}

func (c *CorsoDriveItem) GetCreatedBy() itemIdentitySetable {
	return c.CreatedBy
}

func (c *CorsoDriveItem) GetCreatedDateTime() *time.Time {
	return &c.CreatedDateTime
}

func (c *CorsoDriveItem) GetLastModifiedDateTime() *time.Time {
	return &c.LastModifiedDateTime
}

func (c *CorsoDriveItem) GetMalware() malwareable {
	return c.Malware
}

func (c *CorsoDriveItem) GetDeleted() deletedable {
	return c.Deleted
}

func (c *CorsoDriveItem) GetRoot() itemRootable {
	return c.Root
}

// models.DriveItemable to CorsoDriveItemable
func ToCorsoDriveItemable(item models.DriveItemable) CorsoDriveItemable {
	cdi := &CorsoDriveItem{
		ID:                   strings.Clone(ptr.Val(item.GetId())),
		Name:                 strings.Clone(ptr.Val(item.GetName())),
		Size:                 ptr.Val(item.GetSize()),
		CreatedDateTime:      ptr.Val(item.GetCreatedDateTime()),
		LastModifiedDateTime: ptr.Val(item.GetLastModifiedDateTime()),
	}

	// Hacky way to cache the download url. Thats all we use from additional data
	// Otherwise, we'll hold a reference to the underlying store which will consume
	// lot more memory.
	if item.GetFile() != nil {
		ad := make(map[string]interface{})
		for _, key := range downloadURLKeys {
			if v, err := str.AnyValueToString(key, item.GetAdditionalData()); err == nil {
				ad[key] = strings.Clone(v)
				break
			}
		}

		cdi.AdditionalData = ad
	}

	if item.GetFolder() != nil {
		cdi.Folder = &folderDriveItem{
			isFolder: true,
		}
	}

	if item.GetFile() != nil {
		cdi.File = &fileDriveItem{
			isFile:   true,
			mimeType: strings.Clone(ptr.Val(item.GetFile().GetMimeType())),
		}
	}

	if item.GetPackageEscaped() != nil {
		cdi.Package = &packageDriveItem{
			isPackage: true,
		}
	}

	if item.GetParentReference() != nil {
		cdi.ParentReference = &parentReference{
			id:      strings.Clone(ptr.Val(item.GetParentReference().GetId())),
			path:    strings.Clone(ptr.Val(item.GetParentReference().GetPath())),
			name:    strings.Clone(ptr.Val(item.GetParentReference().GetName())),
			driveId: strings.Clone(ptr.Val(item.GetParentReference().GetDriveId())),
		}
	}

	if item.GetShared() != nil {
		cdi.Shared = &itemShared{
			isShared: true,
		}
	}

	if item.GetMalware() != nil {
		cdi.Malware = &itemMalware{
			isMalware: true,
		}
	}

	if item.GetDeleted() != nil {
		cdi.Deleted = &itemDeleted{
			isDeleted: true,
		}
	}

	if item.GetRoot() != nil {
		cdi.Root = &itemRoot{
			isRoot: true,
		}
	}

	if item.GetCreatedBy() != nil && item.GetCreatedBy().GetUser() != nil {
		additionalData := item.GetCreatedBy().GetUser().GetAdditionalData()
		ad := make(map[string]interface{})
		var str string

		ed, ok := additionalData["email"]
		if ok {
			str = strings.Clone(ptr.Val(ed.(*string)))
			ad["email"] = &str
		} else if ed, ok = additionalData["displayName"]; ok {
			str = strings.Clone(ptr.Val(ed.(*string)))
			ad["displayName"] = &str
		}

		cdi.CreatedBy = &itemIdentitySet{
			user: &itemUser{
				additionalData: ad,
			},
		}
	}

	return cdi
}