package drive

import (
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

func getItemCreator(item models.DriveItemable) string {
	if item.GetCreatedBy() == nil || item.GetCreatedBy().GetUser() == nil {
		return ""
	}

	// User is sometimes not available when created via some
	// external applications (like backup/restore solutions)
	additionalData := item.GetCreatedBy().GetUser().GetAdditionalData()

	ed, ok := additionalData["email"]
	if !ok {
		ed = additionalData["displayName"]
	}

	if ed == nil {
		return ""
	}

	return *ed.(*string)
}

func getItemDriveInfo(item models.DriveItemable) (string, string) {
	if item.GetParentReference() == nil {
		return "", ""
	}

	driveName := strings.TrimSpace(ptr.Val(item.GetParentReference().GetName()))
	driveID := ptr.Val(item.GetParentReference().GetDriveId())

	return driveName, driveID
}

func augmentItemInfo(
	dii details.ItemInfo,
	resource idname.Provider,
	service path.ServiceType,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	var driveName, driveID, creatorEmail string

	// TODO: we rely on this info for details/restore lookups,
	// so if it's nil we have an issue, and will need an alternative
	// way to source the data.

	if item.GetCreatedBy() != nil && item.GetCreatedBy().GetUser() != nil {
		// User is sometimes not available when created via some
		// external applications (like backup/restore solutions)
		additionalData := item.GetCreatedBy().GetUser().GetAdditionalData()

		ed, ok := additionalData["email"]
		if !ok {
			ed = additionalData["displayName"]
		}

		if ed != nil {
			creatorEmail = *ed.(*string)
		}
	}

	if item.GetParentReference() != nil {
		driveID = ptr.Val(item.GetParentReference().GetDriveId())
		driveName = strings.TrimSpace(ptr.Val(item.GetParentReference().GetName()))
	}

	var pps string
	if parentPath != nil {
		pps = parentPath.String()
	}

	switch service {
	case path.OneDriveService:
		dii.OneDrive = &details.OneDriveInfo{
			Created:    ptr.Val(item.GetCreatedDateTime()),
			DriveID:    driveID,
			DriveName:  driveName,
			ItemName:   ptr.Val(item.GetName()),
			ItemType:   details.OneDriveItem,
			Modified:   ptr.Val(item.GetLastModifiedDateTime()),
			Owner:      creatorEmail,
			ParentPath: pps,
			Size:       size,
		}
	case path.SharePointService:
		dii.SharePoint = &details.SharePointInfo{
			Created:    ptr.Val(item.GetCreatedDateTime()),
			DriveID:    driveID,
			DriveName:  driveName,
			ItemName:   ptr.Val(item.GetName()),
			ItemType:   details.SharePointLibrary,
			Modified:   ptr.Val(item.GetLastModifiedDateTime()),
			Owner:      creatorEmail,
			ParentPath: pps,
			SiteID:     resource.ID(),
			Size:       size,
			WebURL:     resource.Name(),
		}

	case path.GroupsService:
		dii.Groups = &details.GroupsInfo{
			Created:    ptr.Val(item.GetCreatedDateTime()),
			DriveID:    driveID,
			DriveName:  driveName,
			ItemName:   ptr.Val(item.GetName()),
			ItemType:   details.SharePointLibrary,
			Modified:   ptr.Val(item.GetLastModifiedDateTime()),
			Owner:      creatorEmail,
			ParentPath: pps,
			SiteID:     resource.ID(),
			Size:       size,
			WebURL:     resource.Name(),
		}
	}

	dii.Extension = &details.ExtensionData{}

	return dii
}
