package onedrive

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/metrics"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ inject.ServiceHandler = &onedriveHandler{}

func NewOneDriveHandler(
	opts control.Options,
	apiClient api.Client,
	resourceGetter idname.GetResourceIDAndNamer,
) *onedriveHandler {
	return &onedriveHandler{
		baseOneDriveHandler: baseOneDriveHandler{
			opts:               opts,
			backupDriveIDNames: idname.NewCache(nil),
		},
		apiClient:      apiClient,
		resourceGetter: resourceGetter,
	}
}

// ========================================================================== //
//                            baseOneDriveHandler
// ========================================================================== //

// baseOneDriveHandler contains logic for tracking data and doing operations
// (e.x. export) that don't require contact with external M356 services.
type baseOneDriveHandler struct {
	opts               control.Options
	backupDriveIDNames idname.CacheBuilder
}

func (h *baseOneDriveHandler) CacheItemInfo(v details.ItemInfo) {
	if v.OneDrive == nil {
		return
	}

	h.backupDriveIDNames.Add(v.OneDrive.DriveID, v.OneDrive.DriveName)
}

// ProduceExportCollections will create the export collections for the
// given restore collections.
func (h *baseOneDriveHandler) ProduceExportCollections(
	ctx context.Context,
	backupVersion int,
	exportCfg control.ExportConfig,
	dcs []data.RestoreCollection,
	stats *metrics.ExportStats,
	errs *fault.Bus,
) ([]export.Collectioner, error) {
	var (
		el = errs.Local()
		ec = make([]export.Collectioner, 0, len(dcs))
	)

	for _, dc := range dcs {
		drivePath, err := path.ToDrivePath(dc.FullPath())
		if err != nil {
			return nil, clues.WrapWC(ctx, err, "transforming path to drive path")
		}

		baseDir := path.Builder{}.Append(drivePath.Folders...)

		ec = append(
			ec,
			drive.NewExportCollection(
				baseDir.String(),
				[]data.RestoreCollection{dc},
				backupVersion,
				stats))
	}

	return ec, el.Failure()
}

// ========================================================================== //
//                              onedriveHandler
// ========================================================================== //

// onedriveHandler contains logic for handling data and performing operations
// (e.x. restore) regardless of whether they require contact with external M365
// services or not.
type onedriveHandler struct {
	baseOneDriveHandler
	apiClient      api.Client
	resourceGetter idname.GetResourceIDAndNamer
}

func (h *onedriveHandler) IsServiceEnabled(
	ctx context.Context,
	resourceID string,
) (bool, error) {
	// TODO(ashmrtn): Move free function implementation to this function.
	res, err := IsServiceEnabled(ctx, h.apiClient.Users(), resourceID)
	return res, clues.Stack(err).OrNil()
}

func (h *onedriveHandler) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	resourceID string, // Can be either ID or name.
	ins idname.Cacher,
) (idname.Provider, error) {
	if h.resourceGetter == nil {
		return nil, clues.StackWC(ctx, resource.ErrNoResourceLookup)
	}

	pr, err := h.resourceGetter.GetResourceIDAndNameFrom(ctx, resourceID, ins)

	return pr, clues.Wrap(err, "identifying resource owner").OrNil()
}
