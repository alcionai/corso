package onedrive

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
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
	AugmentItemInfo(
		dii details.ItemInfo,
		item models.DriveItemable,
		size int64,
		parentPath *path.Builder,
	) details.ItemInfo
	FormatDisplayPath(driveName string, parentPath *path.Builder) string
	NewLocationIDer(driveID string, elems ...string) details.LocationIDer
}
