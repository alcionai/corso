package drive

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

var _ BackupHandler = &groupBackupHandler{}

type groupBackupHandler struct {
	siteBackupHandler
	groupQP graph.QueryParams
	scope   selectors.GroupsScope
}

func NewGroupBackupHandler(
	groupQP, siteQP graph.QueryParams,
	ac api.Drives,
	scope selectors.GroupsScope,
) groupBackupHandler {
	return groupBackupHandler{
		siteBackupHandler: siteBackupHandler{
			baseSiteHandler: baseSiteHandler{
				qp: siteQP,
				ac: ac,
			},
			// Not adding scope here. Anything that needs scope has to
			// be from group handler
			service: path.GroupsService,
		},
		groupQP: groupQP,
		scope:   scope,
	}
}

func (h groupBackupHandler) PathPrefix(
	driveID string,
) (path.Path, error) {
	// TODO: move tenantID to struct
	return path.Build(
		h.groupQP.TenantID,
		h.groupQP.ProtectedResource.ID(),
		h.service,
		path.LibrariesCategory,
		false,
		odConsts.SitesPathDir,
		h.siteBackupHandler.qp.ProtectedResource.ID(),
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h groupBackupHandler) MetadataPathPrefix() (path.Path, error) {
	p, err := path.BuildMetadata(
		h.groupQP.TenantID,
		h.groupQP.ProtectedResource.ID(),
		h.service,
		path.LibrariesCategory,
		false)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata path")
	}

	p, err = p.Append(false, odConsts.SitesPathDir, h.siteBackupHandler.qp.ProtectedResource.ID())
	if err != nil {
		return nil, clues.Wrap(err, "appending site id to metadata path")
	}

	return p, nil
}

func (h groupBackupHandler) CanonicalPath(
	folders *path.Builder,
) (path.Path, error) {
	return folders.ToDataLayerPath(
		h.groupQP.TenantID,
		h.groupQP.ProtectedResource.ID(),
		h.service,
		path.LibrariesCategory,
		false,
		odConsts.SitesPathDir,
		h.siteBackupHandler.qp.ProtectedResource.ID())
}

func (h groupBackupHandler) SitePathPrefix() (path.Path, error) {
	return path.Build(
		h.groupQP.TenantID,
		h.groupQP.ProtectedResource.ID(),
		h.service,
		path.LibrariesCategory,
		false,
		odConsts.SitesPathDir,
		h.siteBackupHandler.qp.ProtectedResource.ID())
}

func (h groupBackupHandler) AugmentItemInfo(
	dii details.ItemInfo,
	resource idname.Provider,
	item *custom.DriveItem,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	var pps string

	if parentPath != nil {
		pps = parentPath.String()
	}

	driveName, driveID := getItemDriveInfo(item)

	dii.Extension = &details.ExtensionData{}
	dii.Groups = &details.GroupsInfo{
		Created:    ptr.Val(item.GetCreatedDateTime()),
		DriveID:    driveID,
		DriveName:  driveName,
		ItemName:   ptr.Val(item.GetName()),
		ItemType:   details.SharePointLibrary,
		Modified:   ptr.Val(item.GetLastModifiedDateTime()),
		Owner:      getItemCreator(item),
		ParentPath: pps,
		SiteID:     resource.ID(),
		Size:       size,
		WebURL:     resource.Name(),
	}

	return dii
}

func (h groupBackupHandler) IsAllPass() bool {
	return h.scope.IsAny(selectors.GroupsLibraryFolder)
}

func (h groupBackupHandler) IncludesDir(dir string) bool {
	return h.scope.Matches(selectors.GroupsLibraryFolder, dir)
}
