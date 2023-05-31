package sharepoint

import (
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
	return api.NewSiteDrivePager(nil, resourceOwner, fields)
}
