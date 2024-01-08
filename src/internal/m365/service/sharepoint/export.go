package sharepoint

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/site"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/metrics"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ inject.ServiceHandler = &sharepointHandler{}

func NewSharePointHandler(
	apiClient api.Client,
	resourceGetter idname.GetResourceIDAndNamer,
) *sharepointHandler {
	return &sharepointHandler{
		baseSharePointHandler: baseSharePointHandler{
			backupDriveIDNames: idname.NewCache(nil),
		},
		apiClient:      apiClient,
		resourceGetter: resourceGetter,
	}
}

// ========================================================================== //
//                          baseSharePointHandler
// ========================================================================== //

// baseSharePointHandler contains logic for tracking data and doing operations
// (e.x. export) that don't require contact with external M356 services.
type baseSharePointHandler struct {
	backupDriveIDNames idname.CacheBuilder
}

func (h *baseSharePointHandler) CacheItemInfo(v details.ItemInfo) {
	// Old versions would store SharePoint data as OneDrive.
	switch {
	case v.SharePoint != nil:
		h.backupDriveIDNames.Add(v.SharePoint.DriveID, v.SharePoint.DriveName)

	case v.OneDrive != nil:
		h.backupDriveIDNames.Add(v.OneDrive.DriveID, v.OneDrive.DriveName)
	}
}

// ProduceExportCollections will create the export collections for the
// given restore collections.
func (h *baseSharePointHandler) ProduceExportCollections(
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
		cat := dc.FullPath().Category()

		ictx := clues.Add(ctx, "fullpath_category", cat)

		switch cat {
		case path.LibrariesCategory:
			drivePath, err := path.ToDrivePath(dc.FullPath())
			if err != nil {
				return nil, clues.WrapWC(ictx, err, "transforming path to drive path")
			}

			driveName, ok := h.backupDriveIDNames.NameOf(drivePath.DriveID)
			if !ok {
				// This should not happen, but just in case
				logger.Ctx(ictx).With("drive_id", drivePath.DriveID).Info("drive name not found, using drive id")
				driveName = drivePath.DriveID
			}

			baseDir := path.Builder{}.
				Append(path.LibrariesCategory.HumanString()).
				Append(driveName).
				Append(drivePath.Folders...)

			coll := drive.NewExportCollection(
				baseDir.String(),
				[]data.RestoreCollection{dc},
				backupVersion,
				stats)

			ec = append(ec, coll)
		case path.ListsCategory:
			folders := dc.FullPath().Folders()
			pth := path.Builder{}.Append(path.ListsCategory.HumanString()).Append(folders...)

			ec = append(
				ec,
				site.NewExportCollection(
					pth.String(),
					[]data.RestoreCollection{dc},
					backupVersion,
					stats))
		default:
			return nil, clues.NewWC(ctx, "data category not supported").
				With("category", cat)
		}
	}

	return ec, el.Failure()
}

// ========================================================================== //
//                            sharepointHandler
// ========================================================================== //

// sharepointHandler contains logic for handling data and performing operations
// (e.x. restore) regardless of whether they require contact with external M365
// services or not.
type sharepointHandler struct {
	baseSharePointHandler
	apiClient      api.Client
	resourceGetter idname.GetResourceIDAndNamer
}

func (h *sharepointHandler) IsServiceEnabled(
	ctx context.Context,
	resourceID string,
) (bool, error) {
	// TODO(ashmrtn): Move free function implementation to this function.
	res, err := IsServiceEnabled(ctx, h.apiClient.Sites(), resourceID)
	return res, clues.Stack(err).OrNil()
}

func (h *sharepointHandler) PopulateProtectedResourceIDAndName(
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
