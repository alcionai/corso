package drive

import (
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

func augmentItemInfo(
	dii details.ItemInfo,
	service path.ServiceType,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	var driveName, siteID, driveID, weburl, creatorEmail string

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

	if service == path.SharePointService ||
		service == path.GroupsService {
		gsi := item.GetSharepointIds()
		if gsi != nil {
			siteID = ptr.Val(gsi.GetSiteId())
			weburl = ptr.Val(gsi.GetSiteUrl())

			if len(weburl) == 0 {
				weburl = constructWebURL(item.GetAdditionalData())
			}
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
			SiteID:     siteID,
			Size:       size,
			WebURL:     weburl,
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
			SiteID:     siteID,
			Size:       size,
			WebURL:     weburl,
		}
	}

	dii.Extension = &details.ExtensionData{}

	return dii
}
