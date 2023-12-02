// Disable revive linter since any structs in this file will expose the same
// funcs as the original structs in the msgraph-sdk-go package, which do not
// follow some of the golint rules.
//
//nolint:revive
package custom

import (
	"strings"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
)

// Replica of models.DriveItemable
type LiteDriveItemable interface {
	GetId() *string
	GetName() *string
	GetSize() *int64
	GetFolder() interface{}
	GetPackageEscaped() interface{}
	GetShared() interface{}
	GetMalware() malwareable
	GetDeleted() interface{}
	GetRoot() interface{}
	GetFile() fileItemable
	GetParentReference() itemReferenceable
	SetParentReference(itemReferenceable)
	GetCreatedBy() identitySetable
	GetCreatedByUser() userable
	GetLastModifiedByUser() userable
	GetCreatedDateTime() *time.Time
	GetLastModifiedDateTime() *time.Time
	GetAdditionalData() map[string]interface{}
}

var _ LiteDriveItemable = &driveItem{}

type driveItem struct {
	id                   string
	name                 string
	size                 int64
	folder               interface{}
	pkg                  interface{}
	shared               interface{}
	deleted              interface{}
	root                 interface{}
	malware              malwareable
	file                 fileItemable
	parentRef            itemReferenceable
	createdBy            identitySetable
	createdByUser        userable
	lastModifiedByUser   userable
	createdDateTime      time.Time
	lastModifiedDateTime time.Time
	additionalData       map[string]interface{}
}

func (c *driveItem) GetId() *string {
	return &c.id
}

func (c *driveItem) GetName() *string {
	return &c.name
}

func (c *driveItem) GetSize() *int64 {
	return &c.size
}

func (c *driveItem) GetFolder() interface{} {
	return c.folder
}

func (c *driveItem) GetPackageEscaped() interface{} {
	return c.pkg
}

func (c *driveItem) GetShared() interface{} {
	return c.shared
}

func (c *driveItem) GetMalware() malwareable {
	return c.malware
}

func (c *driveItem) GetDeleted() interface{} {
	return c.deleted
}

func (c *driveItem) GetRoot() interface{} {
	return c.root
}

func (c *driveItem) GetFile() fileItemable {
	return c.file
}

func (c *driveItem) GetParentReference() itemReferenceable {
	return c.parentRef
}

// TODO(pandeyabs): Should we only support GETs?
func (c *driveItem) SetParentReference(parent itemReferenceable) {
	c.parentRef = parent
}

func (c *driveItem) GetCreatedBy() identitySetable {
	return c.createdBy
}

func (c *driveItem) GetCreatedByUser() userable {
	return c.createdByUser
}

func (c *driveItem) GetLastModifiedByUser() userable {
	return c.lastModifiedByUser
}

func (c *driveItem) GetCreatedDateTime() *time.Time {
	return &c.createdDateTime
}

func (c *driveItem) GetLastModifiedDateTime() *time.Time {
	return &c.lastModifiedDateTime
}

func (c *driveItem) GetAdditionalData() map[string]interface{} {
	return c.additionalData
}

type (
	malwareable interface {
		GetDescription() *string
	}
	fileItemable interface {
		GetMimeType() *string
	}
	itemReferenceable interface {
		GetPath() *string
		GetId() *string
		GetName() *string
		GetDriveId() *string
	}
	identitySetable interface {
		GetUser() identityable
	}
	identityable interface {
		GetAdditionalData() map[string]interface{}
	}
	userable interface {
		GetId() *string
	}
)

// Concrete implementations

var _ malwareable = &malware{}

type malware struct {
	description string
}

func (m *malware) GetDescription() *string {
	return &m.description
}

var _ fileItemable = &fileItem{}

type fileItem struct {
	mimeType string
}

func (f *fileItem) GetMimeType() *string {
	return &f.mimeType
}

var _ itemReferenceable = &parentRef{}

type parentRef struct {
	path    string
	id      string
	name    string
	driveID string
}

func (pr *parentRef) GetPath() *string {
	return &pr.path
}

func (pr *parentRef) GetId() *string {
	return &pr.id
}

func (pr *parentRef) GetName() *string {
	return &pr.name
}

func (pr *parentRef) GetDriveId() *string {
	return &pr.driveID
}

var _ identitySetable = &identitySet{}

type identitySet struct {
	identity identityable
}

func (iis *identitySet) GetUser() identityable {
	return iis.identity
}

var _ identityable = &identity{}

type identity struct {
	additionalData map[string]interface{}
}

func (iu *identity) GetAdditionalData() map[string]interface{} {
	return iu.additionalData
}

var _ userable = &user{}

type user struct {
	id string
}

func (u *user) GetId() *string {
	return &u.id
}

// TODO(pandeyabs): This is duplicated from collection/drive package.
// Move to common pkg
var downloadURLKeys = []string{
	"@microsoft.graph.downloadUrl",
	"@content.downloadUrl",
}

func ToLiteDriveItemable(item models.DriveItemable) LiteDriveItemable {
	if item == nil {
		return nil
	}

	di := &driveItem{
		id:                   strings.Clone(ptr.Val(item.GetId())),
		name:                 strings.Clone(ptr.Val(item.GetName())),
		size:                 ptr.Val(item.GetSize()),
		createdDateTime:      ptr.Val(item.GetCreatedDateTime()),
		lastModifiedDateTime: ptr.Val(item.GetLastModifiedDateTime()),
	}

	if item.GetFolder() != nil {
		di.folder = &struct{}{}
	} else if item.GetFile() != nil {
		di.file = &fileItem{
			mimeType: strings.Clone(ptr.Val(item.GetFile().GetMimeType())),
		}
	} else if item.GetPackageEscaped() != nil {
		di.pkg = &struct{}{}
	}

	if item.GetParentReference() != nil {
		di.parentRef = &parentRef{
			id:      strings.Clone(ptr.Val(item.GetParentReference().GetId())),
			path:    strings.Clone(ptr.Val(item.GetParentReference().GetPath())),
			name:    strings.Clone(ptr.Val(item.GetParentReference().GetName())),
			driveID: strings.Clone(ptr.Val(item.GetParentReference().GetDriveId())),
		}
	}

	if item.GetShared() != nil {
		di.shared = &struct{}{}
	}

	if item.GetMalware() != nil {
		di.malware = &malware{
			description: strings.Clone(ptr.Val(item.GetMalware().GetDescription())),
		}
	}

	if item.GetDeleted() != nil {
		di.deleted = &struct{}{}
	}

	if item.GetRoot() != nil {
		di.root = &struct{}{}
	}

	if item.GetCreatedBy() != nil && item.GetCreatedBy().GetUser() != nil {
		additionalData := item.GetCreatedBy().GetUser().GetAdditionalData()
		ad := make(map[string]interface{})

		if v, err := str.AnyValueToString("email", additionalData); err == nil {
			email := strings.Clone(v)
			ad["email"] = &email
		}

		if v, err := str.AnyValueToString("displayName", additionalData); err == nil {
			displayName := strings.Clone(v)
			ad["displayName"] = &displayName
		}

		di.createdBy = &identitySet{
			identity: &identity{
				additionalData: ad,
			},
		}
	}

	if item.GetCreatedByUser() != nil {
		di.createdByUser = &user{
			id: strings.Clone(ptr.Val(item.GetCreatedByUser().GetId())),
		}
	}

	if item.GetLastModifiedByUser() != nil {
		di.lastModifiedByUser = &user{
			id: strings.Clone(ptr.Val(item.GetLastModifiedByUser().GetId())),
		}
	}

	// We only use the download URL from additional data
	ad := make(map[string]interface{})

	for _, key := range downloadURLKeys {
		if v, err := str.AnyValueToString(key, item.GetAdditionalData()); err == nil {
			downloadURL := strings.Clone(v)
			ad[key] = &downloadURL
		}
	}

	di.additionalData = ad

	return di
}

func SetParentName(orig parentReferenceable, driveName string) parentReferenceable {
	if orig == nil {
		return nil
	}

	pr := orig.(*parentRef)
	pr.name = driveName

	return pr
}
