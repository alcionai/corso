package onedrive

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ConsumeRestoreCollections will restore the specified data collections into OneDrive
func ConsumeRestoreCollections(
	ctx context.Context,
	rh drive.RestoreHandler,
	rcc inject.RestoreConsumerConfig,
	backupDriveIDNames idname.Cacher,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
	ctr *count.Bus,
) (*support.ControllerOperationStatus, error) {
	var (
		restoreMetrics    support.CollectionMetrics
		el                = errs.Local()
		caches            = drive.NewRestoreCaches(backupDriveIDNames)
		fallbackDriveName = rcc.RestoreConfig.Location
	)

	ctx = clues.Add(ctx, "backup_version", rcc.BackupVersion)

	err := caches.Populate(ctx, rh, rcc.ProtectedResource.ID())
	if err != nil {
		return nil, clues.Wrap(err, "initializing restore caches")
	}

	// Reorder collections so that the parents directories are created
	// before the child directories; a requirement for permissions.
	data.SortRestoreCollections(dcs)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			err     error
			metrics support.CollectionMetrics
			ictx    = clues.Add(
				ctx,
				"category", dc.FullPath().Category(),
				"full_path", dc.FullPath())
		)

		metrics, err = drive.RestoreCollection(
			ictx,
			rh,
			rcc,
			dc,
			caches,
			deets,
			fallbackDriveName,
			errs,
			ctr.Local())
		if err != nil {
			el.AddRecoverable(ctx, err)
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if errors.Is(err, context.Canceled) {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		restoreMetrics,
		rcc.RestoreConfig.Location)

	return status, el.Failure()
}
