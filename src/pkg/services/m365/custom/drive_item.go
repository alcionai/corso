package custom

import (
	"strings"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
)

// ---------------------------------------------------------------------------
// DriveItem
// ---------------------------------------------------------------------------
type DriveItem struct {
	id                   *string
	name                 *string
	size                 *int64
	createdDateTime      *time.Time
	lastModifiedDateTime *time.Time
	folder               *struct{}
	pkg                  *struct{}
	shared               *struct{}
	deleted              *struct{}
	root                 *struct{}
	malware              *malware
	file                 *fileItem
	parentRef            *itemReference
	createdBy            *identitySet
	createdByUser        *user
	lastModifiedByUser   *user
	additionalData       map[string]any
}

func NewDriveItem(
	id, name string,
) *DriveItem {
	return &DriveItem{
		id:   ptr.To(id),
		name: ptr.To(name),
	}
}

// Disable revive linter since we want to follow naming scheme used by graph SDK here.
// nolint: revive
func (c *DriveItem) GetId() *string {
	return c.id
}

func (c *DriveItem) GetName() *string {
	return c.name
}

func (c *DriveItem) GetSize() *int64 {
	return c.size
}

func (c *DriveItem) GetCreatedDateTime() *time.Time {
	return c.createdDateTime
}

func (c *DriveItem) GetLastModifiedDateTime() *time.Time {
	return c.lastModifiedDateTime
}

func (c *DriveItem) GetFolder() *struct{} {
	return c.folder
}

func (c *DriveItem) GetPackageEscaped() *struct{} {
	return c.pkg
}

func (c *DriveItem) GetShared() *struct{} {
	return c.shared
}

func (c *DriveItem) GetDeleted() *struct{} {
	return c.deleted
}

func (c *DriveItem) GetRoot() *struct{} {
	return c.root
}

func (c *DriveItem) GetMalware() *malware {
	return c.malware
}

func (c *DriveItem) GetFile() *fileItem {
	return c.file
}

func (c *DriveItem) GetParentReference() *itemReference {
	return c.parentRef
}

func (c *DriveItem) SetParentReference(parent *itemReference) {
	c.parentRef = parent
}

func (c *DriveItem) GetCreatedBy() *identitySet {
	return c.createdBy
}

func (c *DriveItem) GetCreatedByUser() *user {
	return c.createdByUser
}

func (c *DriveItem) GetLastModifiedByUser() *user {
	return c.lastModifiedByUser
}

func (c *DriveItem) GetAdditionalData() map[string]any {
	return c.additionalData
}

// ---------------------------------------------------------------------------
// malware
// ---------------------------------------------------------------------------
type malware struct {
	description *string
}

func (m *malware) GetDescription() *string {
	return m.description
}

// ---------------------------------------------------------------------------
// fileItem
// ---------------------------------------------------------------------------
type fileItem struct {
	mimeType *string
}

func (f *fileItem) GetMimeType() *string {
	return f.mimeType
}

// ---------------------------------------------------------------------------
// itemReference
// ---------------------------------------------------------------------------
type itemReference struct {
	path    *string
	id      *string
	name    *string
	driveID *string
}

func (ir *itemReference) GetPath() *string {
	return ir.path
}

// nolint: revive
func (ir *itemReference) GetId() *string {
	return ir.id
}

func (ir *itemReference) GetName() *string {
	return ir.name
}

// nolint: revive
func (ir *itemReference) GetDriveId() *string {
	return ir.driveID
}

// ---------------------------------------------------------------------------
// identitySet
// ---------------------------------------------------------------------------
type identitySet struct {
	identity *identity
}

func (iis *identitySet) GetUser() *identity {
	return iis.identity
}

// ---------------------------------------------------------------------------
// identity
// ---------------------------------------------------------------------------
type identity struct {
	additionalData map[string]any
}

func (i *identity) GetAdditionalData() map[string]any {
	return i.additionalData
}

// ---------------------------------------------------------------------------
// user
// ---------------------------------------------------------------------------
type user struct {
	id *string
}

// nolint: revive
func (u *user) GetId() *string {
	return u.id
}

// TODO(pandeyabs): This is duplicated from collection/drive package.
// Move to api/graph.
var downloadURLKeys = []string{
	"@microsoft.graph.downloadUrl",
	"@content.downloadUrl",
}

// ToCustomDriveItem converts a DriveItemable to a flattened DriveItem struct
// that stores only the properties we care about during the backup operation.
func ToCustomDriveItem(item models.DriveItemable) *DriveItem {
	if item == nil {
		return nil
	}

	di := &DriveItem{}

	if item.GetId() != nil {
		itemID := strings.Clone(ptr.Val(item.GetId()))
		di.id = &itemID
	}

	if item.GetName() != nil {
		itemName := strings.Clone(ptr.Val(item.GetName()))
		di.name = &itemName
	}

	if item.GetSize() != nil {
		itemSize := ptr.Val(item.GetSize())
		di.size = &itemSize
	}

	if item.GetCreatedDateTime() != nil {
		createdTime := ptr.Val(item.GetCreatedDateTime())
		di.createdDateTime = &createdTime
	}

	if item.GetLastModifiedDateTime() != nil {
		lastModifiedTime := ptr.Val(item.GetLastModifiedDateTime())
		di.lastModifiedDateTime = &lastModifiedTime
	}

	if item.GetFolder() != nil {
		di.folder = &struct{}{}
	}

	if item.GetPackageEscaped() != nil {
		di.pkg = &struct{}{}
	}

	if item.GetMalware() != nil {
		mw := &malware{}

		if item.GetMalware().GetDescription() != nil {
			desc := strings.Clone(ptr.Val(item.GetMalware().GetDescription()))
			mw.description = &desc
		}

		di.malware = mw
	}

	if item.GetFile() != nil {
		fi := &fileItem{}

		if item.GetFile().GetMimeType() != nil {
			mimeType := strings.Clone(ptr.Val(item.GetFile().GetMimeType()))
			fi.mimeType = &mimeType
		}

		di.file = fi
	}

	if item.GetParentReference() != nil {
		iRef := &itemReference{}

		if item.GetParentReference().GetId() != nil {
			parentID := strings.Clone(ptr.Val(item.GetParentReference().GetId()))
			iRef.id = &parentID
		}

		if item.GetParentReference().GetPath() != nil {
			parentPath := strings.Clone(ptr.Val(item.GetParentReference().GetPath()))
			iRef.path = &parentPath
		}

		if item.GetParentReference().GetName() != nil {
			parentName := strings.Clone(ptr.Val(item.GetParentReference().GetName()))
			iRef.name = &parentName
		}

		if item.GetParentReference().GetDriveId() != nil {
			parentDriveID := strings.Clone(ptr.Val(item.GetParentReference().GetDriveId()))
			iRef.driveID = &parentDriveID
		}

		di.parentRef = iRef
	}

	if item.GetShared() != nil {
		di.shared = &struct{}{}
	}

	if item.GetDeleted() != nil {
		di.deleted = &struct{}{}
	}

	if item.GetRoot() != nil {
		di.root = &struct{}{}
	}

	if item.GetCreatedBy() != nil {
		createdBy := &identitySet{}

		if item.GetCreatedBy().GetUser() != nil {
			additionalData := item.GetCreatedBy().GetUser().GetAdditionalData()
			ad := make(map[string]any)

			if v, err := str.AnyValueToString("email", additionalData); err == nil {
				email := strings.Clone(v)
				ad["email"] = &email
			}

			if v, err := str.AnyValueToString("displayName", additionalData); err == nil {
				displayName := strings.Clone(v)
				ad["displayName"] = &displayName
			}

			createdBy.identity = &identity{
				additionalData: ad,
			}
		}

		di.createdBy = createdBy
	}

	if item.GetCreatedByUser() != nil {
		createdByUser := &user{}

		if item.GetCreatedByUser().GetId() != nil {
			userID := strings.Clone(ptr.Val(item.GetCreatedByUser().GetId()))
			createdByUser.id = &userID
		}

		di.createdByUser = createdByUser
	}

	if item.GetLastModifiedByUser() != nil {
		lastModifiedByUser := &user{}

		if item.GetLastModifiedByUser().GetId() != nil {
			userID := strings.Clone(ptr.Val(item.GetLastModifiedByUser().GetId()))
			lastModifiedByUser.id = &userID
		}

		di.lastModifiedByUser = lastModifiedByUser
	}

	// We only use the download URL from additional data
	aData := make(map[string]any)

	for _, key := range downloadURLKeys {
		if v, err := str.AnyValueToString(key, item.GetAdditionalData()); err == nil {
			downloadURL := strings.Clone(v)
			aData[key] = &downloadURL
		}
	}

	di.additionalData = aData

	return di
}

func SetParentName(orig *itemReference, driveName string) *itemReference {
	if orig == nil {
		return nil
	}

	orig.name = &driveName

	return orig
}
