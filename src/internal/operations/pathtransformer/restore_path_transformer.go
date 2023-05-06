package pathtransformer

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

func locationRef(
	ent *details.Entry,
	repoRef path.Path,
	backupVersion int,
) (*path.Builder, error) {
	loc := ent.LocationRef

	// At this backup version all data types should populate LocationRef.
	if len(loc) > 0 || backupVersion >= version.OneDrive7LocationRef {
		return path.Builder{}.SplitUnescapeAppend(loc)
	}

	// We could get an empty LocationRef either because it wasn't populated or it
	// was in the root of the data type.
	elems := repoRef.Folders()

	if ent.OneDrive != nil || ent.SharePoint != nil {
		dp, err := path.ToDrivePath(repoRef)
		if err != nil {
			return nil, clues.Wrap(err, "fallback for LocationRef")
		}

		elems = append([]string{dp.Root}, dp.Folders...)
	}

	return path.Builder{}.Append(elems...), nil
}

func basicPathMerge(repoRef path.Path, locRef *path.Builder) (path.Path, error) {
	if len(locRef.Elements()) == 0 {
		res, err := path.ServicePrefix(
			repoRef.Tenant(),
			repoRef.ResourceOwner(),
			repoRef.Service(),
			repoRef.Category())
		if err != nil {
			return nil, clues.Wrap(err, "getting prefix for empty location")
		}

		return res, nil
	}

	return locRef.ToDataLayerPath(
		repoRef.Tenant(),
		repoRef.ResourceOwner(),
		repoRef.Service(),
		repoRef.Category(),
		false)
}

func drivePathMerge(
	ent *details.Entry,
	repoRef path.Path,
	locRef *path.Builder,
) (path.Path, error) {
	// Try getting the drive ID from the item. Not all details versions had it
	// though.
	var driveID string

	if ent.SharePoint != nil {
		driveID = ent.SharePoint.DriveID
	} else if ent.OneDrive != nil {
		driveID = ent.OneDrive.DriveID
	}

	// Fallback to trying to get from RepoRef.
	if len(driveID) == 0 {
		odp, err := path.ToDrivePath(repoRef)
		if err != nil {
			return nil, clues.Wrap(err, "fallback getting DriveID")
		}

		driveID = odp.DriveID
	}

	return basicPathMerge(
		repoRef,
		path.FormatDriveFolders(driveID, locRef.Elements()...))
}

func restorePathsForItem(
	ctx context.Context,
	backupVersion int,
	ent *details.Entry,
) (path.RestorePaths, error) {
	res := path.RestorePaths{}

	repoRef, err := path.FromDataLayerPath(ent.RepoRef, true)
	if err != nil {
		err = clues.Wrap(err, "parsing details path after reduction").
			WithClues(ctx).
			With("storage_path", clues.Hide(ent.RepoRef), "location", clues.Hide(ent.LocationRef))

		return res, err
	}

	res.StoragePath = repoRef
	ctx = clues.Add(ctx, "storage_path", repoRef)

	// Get the LocationRef so we can munge it onto our path.
	locRef, err := locationRef(ent, repoRef, backupVersion)
	if err != nil {
		err = clues.Wrap(err, "parsing LocationRef after reduction").
			WithClues(ctx).
			With("location", clues.Hide(ent.LocationRef))

		return res, err
	}

	ctx = clues.Add(ctx, "location", locRef)

	// Now figure out what type of ent it is and munge the path accordingly.
	// Eventually we're going to need munging for:
	//   * Exchange Calendars (different folder handling)
	//   * Exchange Email/Contacts
	//   * OneDrive/SharePoint (needs drive information)
	if ent.Exchange != nil {
		// TODO(ashmrtn): Eventually make Events have it's own function to handle
		// setting the restore destination properly.
		res.RestorePath, err = basicPathMerge(repoRef, locRef)
	} else if ent.OneDrive != nil || ent.SharePoint != nil {
		res.RestorePath, err = drivePathMerge(ent, repoRef, locRef)
	} else {
		return res, clues.New("unknown entry type").WithClues(ctx)
	}

	if err != nil {
		return res, clues.Wrap(err, "generating RestorePath").WithClues(ctx)
	}

	return res, nil
}

// GetPaths takes a set of filtered details entries and returns a set of
// RestorePaths for the entries.
func GetPaths(
	ctx context.Context,
	backupVersion int,
	items []*details.Entry,
	errs *fault.Bus,
) ([]path.RestorePaths, error) {
	var (
		paths     = make([]path.RestorePaths, len(items))
		shortRefs = make([]string, len(items))
		el        = errs.Local()
	)

	for i, ent := range items {
		if el.Failure() != nil {
			break
		}

		restorePaths, err := restorePathsForItem(ctx, backupVersion, ent)
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "getting restore paths"))
			continue
		}

		paths[i] = restorePaths
		shortRefs[i] = restorePaths.StoragePath.ShortRef()
	}

	logger.Ctx(ctx).With("short_refs", shortRefs).Infof("found %d details entries to restore", len(shortRefs))

	return paths, el.Failure()
}
