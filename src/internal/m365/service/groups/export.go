package groups

import (
	"context"
	stdlibpath "path"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/groups"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ inject.ServiceHandler = &baseGroupsHandler{}

func NewGroupsHandler(
	opts control.Options,
) *baseGroupsHandler {
	return &baseGroupsHandler{
		opts:               opts,
		backupDriveIDNames: idname.NewCache(nil),
		backupSiteIDWebURL: idname.NewCache(nil),
	}
}

type baseGroupsHandler struct {
	opts control.Options

	backupDriveIDNames idname.CacheBuilder
	backupSiteIDWebURL idname.CacheBuilder
}

func (h *baseGroupsHandler) CacheItemInfo(v details.ItemInfo) {
	if v.Groups == nil {
		return
	}

	h.backupDriveIDNames.Add(v.Groups.DriveID, v.Groups.DriveName)
	h.backupSiteIDWebURL.Add(v.Groups.SiteID, v.Groups.WebURL)
}

// ProduceExportCollections will create the export collections for the
// given restore collections.
func (h *baseGroupsHandler) ProduceExportCollections(
	ctx context.Context,
	backupVersion int,
	exportCfg control.ExportConfig,
	dcs []data.RestoreCollection,
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
				exportCfg,
				stats)

		case path.LibrariesCategory:
			drivePath, err := path.ToDrivePath(restoreColl.FullPath())
			if err != nil {
				return nil, clues.Wrap(err, "transforming path to drive path").WithClues(ctx)
			}

			driveName, ok := h.backupDriveIDNames.NameOf(drivePath.DriveID)
			if !ok {
				// This should not happen, but just in case
				logger.Ctx(ctx).With("drive_id", drivePath.DriveID).Info("drive name not found, using drive id")
				driveName = drivePath.DriveID
			}

			rfds := restoreColl.FullPath().Folders()
			siteName := rfds[1] // use siteID by default

			webURL, ok := h.backupSiteIDWebURL.NameOf(siteName)
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
