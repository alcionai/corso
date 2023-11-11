package onedrive

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

func ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	creds account.M365Config,
	su support.StatusUpdater,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	odb, err := bpc.Selector.ToOneDriveBackup()
	if err != nil {
		return nil, nil, false, clues.WrapWC(ctx, err, "parsing selector")
	}

	var (
		el                   = errs.Local()
		tenantID             = creds.AzureTenantID
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
			drive.NewUserDriveBackupHandler(ac.Drives(), bpc.ProtectedResource.ID(), scope),
			tenantID,
			bpc.ProtectedResource,
			su,
			bpc.Options,
			counter)

		pcfg := observe.ProgressCfg{
			Indent:            1,
			CompletionMessage: func() string { return fmt.Sprintf("(found %d files)", nc.NumFiles) },
		}
		progressBar := observe.MessageWithCompletion(ctx, pcfg, path.FilesCategory.HumanString())

		defer close(progressBar)

		odcs, canUsePreviousBackup, err = nc.Get(ctx, bpc.MetadataCollections, ssmb, errs)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err).Label(fault.LabelForceNoBackupCreation))
		}

		categories[scope.Category().PathType()] = struct{}{}

		collections = append(collections, odcs...)
	}

	mcs, err := migrationCollections(bpc, tenantID, su, counter)
	if err != nil {
		return nil, nil, false, err
	}

	collections = append(collections, mcs...)

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			tenantID,
			bpc.ProtectedResource.ID(),
			path.OneDriveService,
			categories,
			su,
			counter,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	logger.Ctx(ctx).Infow("produced collections", "stats", counter.Values())

	return collections, ssmb.ToReader(), canUsePreviousBackup, el.Failure()
}

// adds data migrations to the collection set.
func migrationCollections(
	bpc inject.BackupProducerConfig,
	tenant string,
	su support.StatusUpdater,
	counter *count.Bus,
) ([]data.BackupCollection, error) {
	// assume a version < 0 implies no prior backup, thus nothing to migrate.
	if version.IsNoBackup(bpc.LastBackupVersion) {
		return nil, nil
	}

	if bpc.LastBackupVersion >= version.All8MigrateUserPNToID {
		return nil, nil
	}

	counter.Inc("requires_migrate_user_pn_to_id")

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

	mgn, err := graph.NewPrefixCollection(mpc, mc, su, counter)
	if err != nil {
		return nil, clues.Wrap(err, "creating migration collection")
	}

	return []data.BackupCollection{mgn}, nil
}
