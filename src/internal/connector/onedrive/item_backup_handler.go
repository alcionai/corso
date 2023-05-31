package onedrive

import (
	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ BackupHandler = &itemBackupHandler{}

type itemBackupHandler struct {
	ac api.Drives
}

func (h itemBackupHandler) PathPrefix(
	tenantID, resourceOwner, driveID string,
) (path.Path, error) {
	return path.Build(
		tenantID,
		resourceOwner,
		path.OneDriveService,
		path.FilesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h itemBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID, resourceOwner string,
) (path.Path, error) {
	return folders.ToDataLayerOneDrivePath(tenantID, resourceOwner, false)
}

func (h itemBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return path.OneDriveService, path.FilesCategory
}

func (h itemBackupHandler) DrivePager(
	resourceOwner string, fields []string,
) api.DrivePager {
	return api.NewUserDrivePager(nil, resourceOwner, fields)
}
