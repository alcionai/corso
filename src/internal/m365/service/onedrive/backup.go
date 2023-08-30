package onedrive

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	tenant string,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	odb, err := bpc.Selector.ToOneDriveBackup()
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "parsing selector").WithClues(ctx)
	}

	var (
		el                   = errs.Local()
		categories           = map[path.CategoryType]struct{}{}
		collections          = []data.BackupCollection{}
		ssmb                 = prefixmatcher.NewStringSetBuilder()
		odcs                 []data.BackupCollection
		canUsePreviousBackup bool
	)

	// for each scope that includes oneDrive items, get all
	for _, scope := range odb.Scopes() {
		if el.Failure() != nil {
			break
		}

		logger.Ctx(ctx).Debug("creating OneDrive collections")

		nc := drive.NewCollections(
			drive.NewItemBackupHandler(ac.Drives(), bpc.ProtectedResource.ID(), scope),
			tenant,
			bpc.ProtectedResource.ID(),
			su,
			bpc.Options)

		odcs, canUsePreviousBackup, err = nc.Get(ctx, bpc.MetadataCollections, ssmb, errs)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err).Label(fault.LabelForceNoBackupCreation))
		}

		categories[scope.Category().PathType()] = struct{}{}

		collections = append(collections, odcs...)
	}

	mcs, err := migrationCollections(bpc, tenant, su)
	if err != nil {
		return nil, nil, false, err
	}

	collections = append(collections, mcs...)

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			tenant,
			bpc.ProtectedResource.ID(),
			path.OneDriveService,
			categories,
			su,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	return collections, ssmb.ToReader(), canUsePreviousBackup, el.Failure()
}

// adds data migrations to the collection set.
func migrationCollections(
	bpc inject.BackupProducerConfig,
	tenant string,
	su support.StatusUpdater,
) ([]data.BackupCollection, error) {
	// assume a version < 0 implies no prior backup, thus nothing to migrate.
	if version.IsNoBackup(bpc.LastBackupVersion) {
		return nil, nil
	}

	if bpc.LastBackupVersion >= version.All8MigrateUserPNToID {
		return nil, nil
	}

	// unlike exchange, which enumerates all folders on every
	// backup, onedrive needs to force the owner PN -> ID migration
	mc, err := path.BuildPrefix(
		tenant,
		bpc.ProtectedResource.ID(),
		path.OneDriveService,
		path.FilesCategory)
	if err != nil {
		return nil, clues.Wrap(err, "creating user id migration path")
	}

	mpc, err := path.BuildPrefix(
		tenant,
		bpc.ProtectedResource.Name(),
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
