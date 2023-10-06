package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ProduceBackupCollections returns a DataCollection which the caller can
// use to read mailbox data out for the specified user
func ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	tenantID string,
	su support.StatusUpdater,
	errs *fault.Bus,
) (inject.BackupProducerResults, error) {
	eb, err := bpc.Selector.ToExchangeBackup()
	if err != nil {
		return inject.BackupProducerResults{}, clues.Wrap(err, "exchange dataCollection selector").WithClues(ctx)
	}

	var (
		collections = []data.BackupCollection{}
		el          = errs.Local()
		categories  = map[path.CategoryType]struct{}{}
		mergedStats = inject.ExchangeStats{}
		handlers    = exchange.BackupHandlers(ac)
	)

	canMakeDeltaQueries, err := canMakeDeltaQueries(ctx, ac.Users(), bpc.ProtectedResource.ID())
	if err != nil {
		return inject.BackupProducerResults{}, clues.Stack(err)
	}

	if !canMakeDeltaQueries {
		logger.Ctx(ctx).Info("delta requests not available")

		bpc.Options.ToggleFeatures.DisableDelta = true
	}

	// Turn on concurrency limiter middleware for exchange backups
	// unless explicitly disabled through DisableConcurrencyLimiterFN cli flag
	graph.InitializeConcurrencyLimiter(
		ctx,
		bpc.Options.ToggleFeatures.DisableConcurrencyLimiter,
		bpc.Options.Parallelism.ItemFetch)

	cdps, canUsePreviousBackup, err := exchange.ParseMetadataCollections(ctx, bpc.MetadataCollections)
	if err != nil {
		return inject.BackupProducerResults{}, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePreviousBackup)

	for _, scope := range eb.Scopes() {
		if el.Failure() != nil {
			break
		}

		dcs, stats, err := exchange.CreateCollections(
			ctx,
			bpc,
			handlers,
			tenantID,
			scope,
			cdps[scope.Category().PathType()],
			su,
			errs)
		if err != nil {
			el.AddRecoverable(ctx, err)
			continue
		}

		categories[scope.Category().PathType()] = struct{}{}

		collections = append(collections, dcs...)

		mergedStats.ContactFolders += stats.ContactFolders
		mergedStats.ContactsAdded += stats.ContactsAdded
		mergedStats.ContactsDeleted += stats.ContactsDeleted
		mergedStats.EventFolders += stats.EventFolders
		mergedStats.EventsAdded += stats.EventsAdded
		mergedStats.EventsDeleted += stats.EventsDeleted
		mergedStats.EmailFolders += stats.EmailFolders
		mergedStats.EmailsAdded += stats.EmailsAdded
		mergedStats.EmailsDeleted += stats.EmailsDeleted
	}

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			tenantID,
			bpc.ProtectedResource.ID(),
			path.ExchangeService,
			categories,
			su,
			errs)
		if err != nil {
			return inject.BackupProducerResults{}, err
		}

		collections = append(collections, baseCols...)
	}

	return inject.BackupProducerResults{
			Collections:          collections,
			Excludes:             nil,
			CanUsePreviousBackup: canUsePreviousBackup,
			DiscoveredItems:      inject.Stats{Exchange: &mergedStats}},
		el.Failure()
}

func canMakeDeltaQueries(
	ctx context.Context,
	gmi getMailboxer,
	resourceOwner string,
) (bool, error) {
	mi, err := GetMailboxInfo(ctx, gmi, resourceOwner)
	if err != nil {
		return false, clues.Stack(err)
	}

	return !mi.QuotaExceeded, nil
}
