package groups

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/groups"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

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
	errs *fault.Bus,
) ([]export.Collectioner, error) {
	var (
		el = errs.Local()
		ec = make([]export.Collectioner, 0, len(dcs))
	)

	for _, restoreColl := range dcs {
		var (
			fp      = restoreColl.FullPath()
			cat     = fp.Category()
			folders = []string{cat.String()}
			coll    export.Collectioner
		)

		switch cat {
		case path.ChannelMessagesCategory:
			folders = append(folders, fp.Folders()...)

			coll = groups.NewExportCollection(
				path.Builder{}.Append(folders...).String(),
				[]data.RestoreCollection{restoreColl},
				backupVersion,
				exportCfg)
		case path.FilesCategory:
			drivePath, err := path.ToDrivePath(restoreColl.FullPath())
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
				Append("Libraries").
				Append(driveName).
				Append(drivePath.Folders...)

			coll = drive.NewExportCollection(
				baseDir.String(),
				[]data.RestoreCollection{restoreColl},
				backupVersion)
		default:
			el.AddRecoverable(
				ctx,
				clues.New("unsupported category for export").With("category", cat))
		}

		if coll != nil {
			ec = append(ec, coll)
		}
	}

	return ec, el.Failure()
}
