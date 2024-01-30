package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type exchangeBackup struct{}

// NewBackup provides a struct that matches standard apis
// across m365/service handlers.
func NewBackup() *exchangeBackup {
	return &exchangeBackup{}
}

// ProduceBackupCollections returns a DataCollection which the caller can
// use to read mailbox data out for the specified user
func (exchangeBackup) ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	creds account.M365Config,
	su support.StatusUpdater,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	eb, err := bpc.Selector.ToExchangeBackup()
	if err != nil {
		return nil, nil, false, clues.WrapWC(ctx, err, "exchange dataCollection selector")
	}

	var (
		collections = []data.BackupCollection{}
		el          = errs.Local()
		tenantID    = creds.AzureTenantID
		categories  = map[path.CategoryType]struct{}{}
		handlers    = exchange.BackupHandlers(ac)
	)

	canMakeDeltaQueries, err := canMakeDeltaQueries(ctx, ac.Users(), bpc.ProtectedResource.ID())
	if err != nil {
		return nil, nil, false, clues.Stack(err)
	}

	if !canMakeDeltaQueries {
		logger.Ctx(ctx).Info("delta requests not available")
		counter.Inc(count.NoDeltaQueries)

		bpc.Options.ToggleFeatures.DisableDelta = true
	}

	graph.InitializeConcurrencyLimiter(
		ctx,
		true,
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
			counter,
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
			counter,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	logger.Ctx(ctx).Infow("produced collections", "stats", counter.Values())

	return collections, nil, canUsePreviousBackup, el.Failure()
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
