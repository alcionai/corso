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

// TODO(pandeyabs): Remove this duplicate
// downloadUrlKeys is used to find the download URL in a DriveItem response.
var downloadURLKeys = []string{
	"@microsoft.graph.downloadUrl",
	"@content.downloadUrl",
}

// Replica of models.DriveItemable
type LiteDriveItemable interface {
	GetId() *string
	GetName() *string
	GetSize() *int64
	// TODO(pandeyabs): replace with any
	GetFolder() interface{}
	GetPackageEscaped() interface{}
	GetShared() interface{}
	GetMalware() interface{}
	GetDeleted() interface{}
	GetRoot() interface{}
	GetFile() fileItemable
	GetParentReference() parentReferenceable
	SetParentReference(parentReferenceable)
	GetCreatedBy() itemIdentitySetable
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
	malware              interface{}
	deleted              interface{}
	root                 interface{}
	file                 fileItemable
	parentRef            parentReferenceable
	createdBy            itemIdentitySetable
	createdDateTime      time.Time
	lastModifiedDateTime time.Time
	additionalData       map[string]interface{}
}

// nolint
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

func (c *driveItem) GetMalware() interface{} {
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

func (c *driveItem) GetParentReference() parentReferenceable {
	return c.parentRef
}

// TODO(pandeyabs): Should we only support GETs?
func (c *driveItem) SetParentReference(parent parentReferenceable) {
	c.parentRef = parent
}

func (c *driveItem) GetCreatedBy() itemIdentitySetable {
	return c.createdBy
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
	fileItemable interface {
		GetMimeType() *string
	}
	parentReferenceable interface {
		GetPath() *string
		GetId() *string
		GetName() *string
		GetDriveId() *string
	}
	itemIdentitySetable interface {
		GetUser() itemUserable
	}
	itemUserable interface {
		GetAdditionalData() map[string]interface{}
	}
)

// Concrete implementations

var _ fileItemable = &fileItem{}

type fileItem struct {
	mimeType string
}

func (f *fileItem) GetMimeType() *string {
	return &f.mimeType
}

var _ parentReferenceable = &parentRef{}

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

var _ itemIdentitySetable = &itemIdentitySet{}

type itemIdentitySet struct {
	user itemUserable
}

func (iis *itemIdentitySet) GetUser() itemUserable {
	return iis.user
}

var _ itemUserable = &itemUser{}

type itemUser struct {
	additionalData map[string]interface{}
}

func (iu *itemUser) GetAdditionalData() map[string]interface{} {
	return iu.additionalData
}

func ToLiteDriveItemable(item models.DriveItemable) LiteDriveItemable {
	cdi := &driveItem{
		id:                   strings.Clone(ptr.Val(item.GetId())),
		name:                 strings.Clone(ptr.Val(item.GetName())),
		size:                 ptr.Val(item.GetSize()),
		createdDateTime:      ptr.Val(item.GetCreatedDateTime()),
		lastModifiedDateTime: ptr.Val(item.GetLastModifiedDateTime()),
	}

	if item.GetFolder() != nil {
		cdi.folder = &struct{}{}
	} else if item.GetFile() != nil {
		cdi.file = &fileItem{
			mimeType: strings.Clone(ptr.Val(item.GetFile().GetMimeType())),
		}
	} else if item.GetPackageEscaped() != nil {
		cdi.pkg = &struct{}{}
	}

	if item.GetParentReference() != nil {
		cdi.parentRef = &parentRef{
			id:      strings.Clone(ptr.Val(item.GetParentReference().GetId())),
			path:    strings.Clone(ptr.Val(item.GetParentReference().GetPath())),
			name:    strings.Clone(ptr.Val(item.GetParentReference().GetName())),
			driveID: strings.Clone(ptr.Val(item.GetParentReference().GetDriveId())),
		}
	}

	if item.GetShared() != nil {
		cdi.shared = &struct{}{}
	}

	if item.GetMalware() != nil {
		cdi.malware = &struct{}{}
	}

	if item.GetDeleted() != nil {
		cdi.deleted = &struct{}{}
	}

	if item.GetRoot() != nil {
		cdi.root = &struct{}{}
	}

	if item.GetCreatedBy() != nil && item.GetCreatedBy().GetUser() != nil {
		additionalData := item.GetCreatedBy().GetUser().GetAdditionalData()
		ad := make(map[string]interface{})

		var s string

		ed, ok := additionalData["email"]
		if ok {
			s = strings.Clone(ptr.Val(ed.(*string)))
			ad["email"] = &s
		} else if ed, ok = additionalData["displayName"]; ok {
			s = strings.Clone(ptr.Val(ed.(*string)))
			ad["displayName"] = &s
		}

		cdi.createdBy = &itemIdentitySet{
			user: &itemUser{
				additionalData: ad,
			},
		}
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

		cdi.additionalData = ad
	}

	return cdi
}
