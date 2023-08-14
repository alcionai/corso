package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/fault"
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
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	eb, err := bpc.Selector.ToExchangeBackup()
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "exchange dataCollection selector").WithClues(ctx)
	}

	var (
		collections = []data.BackupCollection{}
		el          = errs.Local()
		categories  = map[path.CategoryType]struct{}{}
		handlers    = exchange.BackupHandlers(ac)
	)

	// Turn on concurrency limiter middleware for exchange backups
	// unless explicitly disabled through DisableConcurrencyLimiterFN cli flag
	graph.InitializeConcurrencyLimiter(
		ctx,
		bpc.Options.ToggleFeatures.DisableConcurrencyLimiter,
		bpc.Options.Parallelism.ItemFetch)

	cdps, canUsePreviousBackup, err := exchange.ParseMetadataCollections(ctx, bpc.MetadataCollections)
	if err != nil {
		return nil, nil, false, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePreviousBackup)

	for _, scope := range eb.Scopes() {
		if el.Failure() != nil {
			break
		}

		dcs, err := exchange.CreateCollections(
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
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	return collections, nil, canUsePreviousBackup, el.Failure()
}
