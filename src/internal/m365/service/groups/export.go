package groups

import (
	"context"
	stdlibpath "path"

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
	backupDriveIDNames idname.Cacher,
	backupSiteIDWebURL idname.Cacher,
	deets *details.Builder,
	stats *data.ExportStats,
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
			folders = []string{cat.HumanString()}
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
		case path.LibrariesCategory:
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

			rfds := restoreColl.FullPath().Folders()
			siteName := rfds[1] // use siteID by default

			webURL, ok := backupSiteIDWebURL.NameOf(siteName)
			if !ok {
				// This should not happen, but just in case
				logger.Ctx(ctx).With("site_id", rfds[1]).Info("site weburl not found, using site id")
			}

			if len(webURL) != 0 {
				// We can't use the actual name anyways as it might
				// contain invalid characters. This should also avoid
				// possibility of name collisions.
				siteName = stdlibpath.Base(webURL)
			}

			baseDir := path.Builder{}.
				Append(folders...).
				Append(siteName).
				Append(driveName).
				Append(drivePath.Folders...)

			coll = drive.NewExportCollection(
				baseDir.String(),
				[]data.RestoreCollection{restoreColl},
				backupVersion,
				stats)
		default:
			el.AddRecoverable(
				ctx,
				clues.New("unsupported category for export").With("category", cat))

			continue
		}

		ec = append(ec, coll)
	}

	return ec, el.Failure()
}
