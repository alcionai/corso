package drive

import (
	"context"
	"net/http"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ BackupHandler = &groupBackupHandler{}

type groupBackupHandler struct {
	groupID string
	ac      api.Drives
	scope   selectors.GroupsScope
}

func NewGroupBackupHandler(groupID string, ac api.Drives, scope selectors.GroupsScope) groupBackupHandler {
	return groupBackupHandler{groupID, ac, scope}
}

func (h groupBackupHandler) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return h.ac.Get(ctx, url, headers)
}

func (h groupBackupHandler) PathPrefix(
	tenantID, resourceOwner, driveID string,
) (path.Path, error) {
	return path.Build(
		tenantID,
		resourceOwner,
		path.GroupsService,
		path.LibrariesCategory, // TODO(meain)
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h groupBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID, resourceOwner string,
) (path.Path, error) {
	// TODO(meain): path fixes: sharepoint site ids should be in the path
	return folders.ToDataLayerPath(
		tenantID,
		h.groupID,
		path.GroupsService,
		path.LibrariesCategory,
		false)
}

func (h groupBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return path.GroupsService, path.LibrariesCategory
}

func (h groupBackupHandler) NewDrivePager(
	resourceOwner string,
	fields []string,
) api.DrivePager {
	return h.ac.NewSiteDrivePager(resourceOwner, fields)
}

func (h groupBackupHandler) NewItemPager(
	driveID, link string,
	fields []string,
) api.DriveItemDeltaEnumerator {
	return h.ac.NewDriveItemDeltaPager(driveID, link, fields)
}

func (h groupBackupHandler) AugmentItemInfo(
	dii details.ItemInfo,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	return augmentGroupItemInfo(dii, item, size, parentPath)
}

func (h groupBackupHandler) FormatDisplayPath(
	driveName string,
	pb *path.Builder,
) string {
	return "/" + driveName + "/" + pb.String()
}

func (h groupBackupHandler) NewLocationIDer(
	driveID string,
	elems ...string,
) details.LocationIDer {
	return details.NewSharePointLocationIDer(driveID, elems...)
}

func (h groupBackupHandler) GetItemPermission(
	ctx context.Context,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	return h.ac.GetItemPermission(ctx, driveID, itemID)
}

func (h groupBackupHandler) GetItem(
	ctx context.Context,
	driveID, itemID string,
) (models.DriveItemable, error) {
	return h.ac.GetItem(ctx, driveID, itemID)
}

func (h groupBackupHandler) IsAllPass() bool {
	// TODO(meain)
	return true
}

func (h groupBackupHandler) IncludesDir(dir string) bool {
	// TODO(meain)
	// return h.scope.Matches(selectors.SharePointGroupFolder, dir)
	return true
}

// ---------------------------------------------------------------------------
// Common
// ---------------------------------------------------------------------------

func augmentGroupItemInfo(
	dii details.ItemInfo,
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

	// gsi := item.GetSharepointIds()
	// if gsi != nil {
	// 	siteID = ptr.Val(gsi.GetSiteId())
	// 	weburl = ptr.Val(gsi.GetSiteUrl())

	// 	if len(weburl) == 0 {
	// 		weburl = constructWebURL(item.GetAdditionalData())
	// 	}
	// }

	if item.GetParentReference() != nil {
		driveID = ptr.Val(item.GetParentReference().GetDriveId())
		driveName = strings.TrimSpace(ptr.Val(item.GetParentReference().GetName()))
	}

	var pps string
	if parentPath != nil {
		pps = parentPath.String()
	}

	dii.Groups = &details.GroupsInfo{
		Created:    ptr.Val(item.GetCreatedDateTime()),
		DriveID:    driveID,
		DriveName:  driveName,
		ItemName:   ptr.Val(item.GetName()),
		ItemType:   details.SharePointLibrary,
		Modified:   ptr.Val(item.GetLastModifiedDateTime()),
		Owner:      creatorEmail,
		ParentPath: pps,
		Size:       size,
	}

	dii.Extension = &details.ExtensionData{}

	return dii
}
