package sharepoint

import (
	"strings"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

var _ onedrive.BackupHandler = &libraryBackupHandler{}

type libraryBackupHandler struct {
	ac api.Drives
}

func (h libraryBackupHandler) PathPrefix(
	tenantID, resourceOwner, driveID string,
) (path.Path, error) {
	return path.Build(
		tenantID,
		resourceOwner,
		path.SharePointService,
		path.LibrariesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h libraryBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID, resourceOwner string,
) (path.Path, error) {
	return folders.ToDataLayerSharePointPath(tenantID, resourceOwner, path.LibrariesCategory, false)
}

func (h libraryBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return path.SharePointService, path.LibrariesCategory
}

func (h libraryBackupHandler) DrivePager(
	resourceOwner string, fields []string,
) api.DrivePager {
	return h.ac.NewSiteDrivePager(resourceOwner, fields)
}

// AugmentItemInfo will populate a details.SharePointInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func (h libraryBackupHandler) AugmentItemInfo(
	dii details.ItemInfo,
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

	gsi := item.GetSharepointIds()
	if gsi != nil {
		siteID = ptr.Val(gsi.GetSiteId())
		weburl = ptr.Val(gsi.GetSiteUrl())

		if len(weburl) == 0 {
			weburl = constructWebURL(item.GetAdditionalData())
		}
	}

	if item.GetParentReference() != nil {
		driveID = ptr.Val(item.GetParentReference().GetDriveId())
		driveName = strings.TrimSpace(ptr.Val(item.GetParentReference().GetName()))
	}

	dii.SharePoint = &details.SharePointInfo{
		Created:    ptr.Val(item.GetCreatedDateTime()),
		DriveID:    driveID,
		DriveName:  driveName,
		ItemName:   ptr.Val(item.GetName()),
		ItemType:   details.SharePointLibrary,
		Modified:   ptr.Val(item.GetLastModifiedDateTime()),
		ParentPath: parentPath.String(),
		Owner:      creatorEmail,
		SiteID:     siteID,
		Size:       size,
		WebURL:     weburl,
	}

	return dii
}

// constructWebURL helper function for recreating the webURL
// for the originating SharePoint site. Uses additional data map
// from a models.DriveItemable that possesses a downloadURL within the map.
// Returns "" if map nil or key is not present.
func constructWebURL(adtl map[string]any) string {
	var (
		desiredKey = "@microsoft.graph.downloadUrl"
		sep        = `/_layouts`
		url        string
	)

	if adtl == nil {
		return url
	}

	r := adtl[desiredKey]
	point, ok := r.(*string)

	if !ok {
		return url
	}

	value := ptr.Val(point)
	if len(value) == 0 {
		return url
	}

	temp := strings.Split(value, sep)
	url = temp[0]

	return url
}

func (h libraryBackupHandler) FormatDisplayPath(
	driveName string,
	pb *path.Builder,
) string {
	return "/" + driveName + "/" + pb.String()
}

func (h libraryBackupHandler) NewLocationIDer(
	driveID string,
	elems ...string,
) details.LocationIDer {
	return details.NewSharePointLocationIDer(driveID, elems...)
}
