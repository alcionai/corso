package sharepoint

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

func NewSharePointHandler(
	opts control.Options,
) *sharepointHandler {
	return &sharepointHandler{
		opts:               opts,
		backupDriveIDNames: idname.NewCache(nil),
	}
}

type sharepointHandler struct {
	opts control.Options

	backupDriveIDNames idname.CacheBuilder
}

func (h *sharepointHandler) CacheItemInfo(v details.ItemInfo) {
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
func ProduceExportCollections(
	ctx context.Context,
	backupVersion int,
	exportCfg control.ExportConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	backupDriveIDNames idname.CacheBuilder,
	deets *details.Builder,
	stats *data.ExportStats,
	errs *fault.Bus,
) ([]export.Collectioner, error) {
	var (
		el = errs.Local()
		ec = make([]export.Collectioner, 0, len(dcs))
	)

	for _, dc := range dcs {
		drivePath, err := path.ToDrivePath(dc.FullPath())
		if err != nil {
			return nil, clues.Wrap(err, "transforming path to drive path").WithClues(ctx)
		}

		driveName, ok := backupDriveIDNames.NameOf(drivePath.DriveID)
		if !ok {
			// This should not happen, but just in case
			logger.Ctx(ctx).With("drive_id", drivePath.DriveID).Info("drive name not found, using drive id")
			driveName = drivePath.DriveID
		}

		baseDir := path.Builder{}.
			Append(path.LibrariesCategory.HumanString()).
			Append(driveName).
			Append(drivePath.Folders...)

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
