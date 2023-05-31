package onedrive

import (
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// backup
// ---------------------------------------------------------------------------

type BackupHandler interface {
	PathPrefix(tenantID, resourceOwner, driveID string) (path.Path, error)
	CanonicalPath(
		folders *path.Builder,
		tenantID, resourceOwner string,
	) (path.Path, error)
	ServiceCat() (path.ServiceType, path.CategoryType)
	DrivePager(resourceOwner string, fields []string) api.DrivePager
}
