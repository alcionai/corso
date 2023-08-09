package onedrive

import (
	"context"
	"sort"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

// Augment restore path to add extra files(meta) needed for restore as
// well as do any other ordering operations on the paths
//
// Only accepts StoragePath/RestorePath pairs where the RestorePath is
// at least as long as the StoragePath. If the RestorePath is longer than the
// StoragePath then the first few (closest to the root) directories will use
// default permissions during restore.
func AugmentRestorePaths(
	backupVersion int,
	paths []path.RestorePaths,
) ([]path.RestorePaths, error) {
	// Keyed by each value's StoragePath.String() which corresponds to the RepoRef
	// of the directory.
	colPaths := map[string]path.RestorePaths{}

	for _, p := range paths {
		first := true

		for {
			sp, err := p.StoragePath.Dir()
			if err != nil {
				return nil, err
			}

			drivePath, err := path.ToDrivePath(sp)
			if err != nil {
				return nil, err
			}

			if len(drivePath.Folders) == 0 {
				break
			}

			if len(p.RestorePath.Elements()) < len(sp.Elements()) {
				return nil, clues.New("restorePath shorter than storagePath").
					With("restore_path", p.RestorePath, "storage_path", sp)
			}

			rp := p.RestorePath

			// Make sure the RestorePath always points to the level of the current
			// collection. We need to track if it's the first iteration because the
			// RestorePath starts out at the collection level to begin with.
			if !first {
				rp, err = p.RestorePath.Dir()
				if err != nil {
					return nil, err
				}
			}

			paths := path.RestorePaths{
				StoragePath: sp,
				RestorePath: rp,
			}

			colPaths[sp.String()] = paths
			p = paths
			first = false
		}
	}

	// Adds dirmeta files as we need to make sure collections for all
	// directories involved are created and not just the final one. No
	// need to add `.meta` files (metadata for files) as they will
	// anyways be looked up automatically.
	// TODO: Stop populating .dirmeta for newer versions once we can
	// get files from parent directory via `Fetch` in a collection.
	// As of now look up metadata for parent directories from a
	// collection.
	for _, p := range colPaths {
		el := p.StoragePath.Elements()

		if backupVersion >= version.OneDrive6NameInMeta {
			mPath, err := p.StoragePath.AppendItem(".dirmeta")
			if err != nil {
				return nil, err
			}

			paths = append(
				paths,
				path.RestorePaths{StoragePath: mPath, RestorePath: p.RestorePath})
		} else if backupVersion >= version.OneDrive4DirIncludesPermissions {
			mPath, err := p.StoragePath.AppendItem(el.Last() + ".dirmeta")
			if err != nil {
				return nil, err
			}

			paths = append(
				paths,
				path.RestorePaths{StoragePath: mPath, RestorePath: p.RestorePath})
		} else if backupVersion >= version.OneDrive1DataAndMetaFiles {
			pp, err := p.StoragePath.Dir()
			if err != nil {
				return nil, err
			}

			mPath, err := pp.AppendItem(el.Last() + ".dirmeta")
			if err != nil {
				return nil, err
			}

			prp, err := p.RestorePath.Dir()
			if err != nil {
				return nil, err
			}

			paths = append(
				paths,
				path.RestorePaths{StoragePath: mPath, RestorePath: prp})
		}
	}

	// This sort is done primarily to order `.meta` files after `.data`
	// files. This is only a necessity for OneDrive as we are storing
	// metadata for files/folders in separate meta files and we the
	// data to be restored before we can restore the metadata.
	//
	// This sorting assumes stuff in the same StoragePath directory end up in the
	// same RestorePath collection.
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].StoragePath.String() < paths[j].StoragePath.String()
	})

	return paths, nil
}
