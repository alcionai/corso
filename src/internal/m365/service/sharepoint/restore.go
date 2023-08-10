package sharepoint

import (
	"context"
	"errors"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/site"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ConsumeRestoreCollections will restore the specified data collections into OneDrive
func ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	ac api.Client,
	backupDriveIDNames idname.Cacher,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
	ctr *count.Bus,
) (*support.ControllerOperationStatus, error) {
	var (
		lrh            = drive.NewLibraryRestoreHandler(ac)
		restoreMetrics support.CollectionMetrics
		caches         = drive.NewRestoreCaches(backupDriveIDNames)
		el             = errs.Local()
	)

	err := caches.Populate(ctx, lrh, rcc.ProtectedResource.ID())
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
			err      error
			category = dc.FullPath().Category()
			metrics  support.CollectionMetrics
			ictx     = clues.Add(ctx,
				"category", category,
				"restore_location", clues.Hide(rcc.RestoreConfig.Location),
				"resource_owner", clues.Hide(dc.FullPath().ResourceOwner()),
				"full_path", dc.FullPath())
		)

		switch dc.FullPath().Category() {
		case path.LibrariesCategory:
			metrics, err = drive.RestoreCollection(
				ictx,
				lrh,
				rcc,
				dc,
				caches,
				deets,
				control.DefaultRestoreContainerName(dttm.HumanReadableDriveItem),
				errs,
				ctr)

		case path.ListsCategory:
			metrics, err = site.RestoreListCollection(
				ictx,
				ac.Stable,
				dc,
				rcc.RestoreConfig.Location,
				deets,
				errs)

		case path.PagesCategory:
			metrics, err = site.RestorePageCollection(
				ictx,
				ac.Stable,
				dc,
				rcc.RestoreConfig.Location,
				deets,
				errs)

		default:
			return nil, clues.Wrap(clues.New(category.String()), "category not supported").With("category", category)
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if err != nil {
			el.AddRecoverable(ctx, err)
		}

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
