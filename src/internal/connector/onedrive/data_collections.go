package onedrive

import (
	"context"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type odFolderMatcher struct {
	scope selectors.OneDriveScope
}

func (fm odFolderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.OneDriveFolder)
}

func (fm odFolderMatcher) Matches(dir string) bool {
	return fm.scope.Matches(selectors.OneDriveFolder, dir)
}

// OneDriveDataCollections returns a set of DataCollection which represents the OneDrive data
// for the specified user
func DataCollections(
	ctx context.Context,
	selector selectors.Selector,
	user idname.Provider,
	metadata []data.RestoreCollection,
	lastBackupVersion int,
	tenant string,
	itemClient graph.Requester,
	service graph.Servicer,
	su support.StatusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]map[string]struct{}, error) {
	odb, err := selector.ToOneDriveBackup()
	if err != nil {
		return nil, nil, clues.Wrap(err, "parsing selector").WithClues(ctx)
	}

	var (
		el          = errs.Local()
		categories  = map[path.CategoryType]struct{}{}
		collections = []data.BackupCollection{}
		allExcludes = map[string]map[string]struct{}{}
	)

	// for each scope that includes oneDrive items, get all
	for _, scope := range odb.Scopes() {
		if el.Failure() != nil {
			break
		}

		logger.Ctx(ctx).Debug("creating OneDrive collections")

		nc := NewCollections(
			itemClient,
			tenant,
			user.ID(),
			OneDriveSource,
			odFolderMatcher{scope},
			service,
			su,
			ctrlOpts)

		odcs, excludes, err := nc.Get(ctx, metadata, errs)
		if err != nil {
			el.AddRecoverable(clues.Stack(err).Label(fault.LabelForceNoBackupCreation))
		}

		categories[scope.Category().PathType()] = struct{}{}

		collections = append(collections, odcs...)

		for k, ex := range excludes {
			if _, ok := allExcludes[k]; !ok {
				allExcludes[k] = map[string]struct{}{}
			}

			maps.Copy(allExcludes[k], ex)
		}
	}

	mcs, err := migrationCollections(
		service,
		lastBackupVersion,
		tenant,
		user,
		su,
		ctrlOpts)
	if err != nil {
		return nil, nil, err
	}

	collections = append(collections, mcs...)

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			tenant,
			user.ID(),
			path.OneDriveService,
			categories,
			su,
			errs)
		if err != nil {
			return nil, nil, err
		}

		collections = append(collections, baseCols...)
	}

	return collections, allExcludes, el.Failure()
}

// adds data migrations to the collection set.
func migrationCollections(
	svc graph.Servicer,
	lastBackupVersion int,
	tenant string,
	user idname.Provider,
	su support.StatusUpdater,
	ctrlOpts control.Options,
) ([]data.BackupCollection, error) {
	// assume a version < 0 implies no prior backup, thus nothing to migrate.
	if lastBackupVersion < 0 {
		return nil, nil
	}

	if lastBackupVersion >= version.AllXMigrateUserPNToID {
		return nil, nil
	}

	// unlike exchange, which enumerates all folders on every
	// backup, onedrive needs to force the owner PN -> ID migration
	mc, err := path.ServicePrefix(
		tenant,
		user.ID(),
		path.OneDriveService,
		path.FilesCategory)
	if err != nil {
		return nil, clues.Wrap(err, "creating user id migration path")
	}

	mpc, err := path.ServicePrefix(
		tenant,
		user.Name(),
		path.OneDriveService,
		path.FilesCategory)
	if err != nil {
		return nil, clues.Wrap(err, "creating user name migration path")
	}

	mgn, err := graph.NewPrefixCollection(mpc, mc, su)
	if err != nil {
		return nil, clues.Wrap(err, "creating migration collection")
	}

	return []data.BackupCollection{mgn}, nil
}
