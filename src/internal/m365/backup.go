package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	kinject "github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/groups"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/m365/service/teamschats"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type backupHandler interface {
	produceBackupCollectionser
}

type produceBackupCollectionser interface {
	ProduceBackupCollections(
		ctx context.Context,
		bpc inject.BackupProducerConfig,
		ac api.Client,
		creds account.M365Config,
		su support.StatusUpdater,
		counter *count.Bus,
		errs *fault.Bus,
	) (
		collections []data.BackupCollection,
		excludeItems *prefixmatcher.StringSetMatcher,
		canUsePreviousBackup bool,
		err error,
	)
}

// ---------------------------------------------------------------------------
// Data Collections
// ---------------------------------------------------------------------------

// ProduceBackupCollections generates a slice of data.BackupCollections for the service
// specified in the selectors.
// The metadata field can include things like delta tokens or the previous backup's
// folder hierarchy. The absence of metadata causes the collection creation to ignore
// prior history (ie, incrementals) and run a full backup.
func (ctrl *Controller) ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error) {
	service := bpc.Selector.PathService()

	ctx, end := diagnostics.Span(
		ctx,
		"m365:produceBackupCollections",
		diagnostics.Index("service", bpc.Selector.PathService().String()))
	defer end()

	// Limit the max number of active requests to graph from this collection.
	bpc.Options.Parallelism.ItemFetch = graph.Parallelism(service).
		ItemOverride(ctx, bpc.Options.Parallelism.ItemFetch)

	err := verifyBackupInputs(bpc.Selector, ctrl.IDNameLookup.IDs())
	if err != nil {
		return nil, nil, false, clues.StackWC(ctx, err)
	}

	var (
		colls                []data.BackupCollection
		excludeItems         *prefixmatcher.StringSetMatcher
		canUsePreviousBackup bool
	)

	var handler backupHandler

	switch service {
	case path.ExchangeService:
		handler = exchange.NewBackup()

	case path.OneDriveService:
		handler = onedrive.NewBackup()

	case path.SharePointService:
		handler = sharepoint.NewBackup()

	case path.GroupsService:
		handler = groups.NewBackup()

	case path.TeamsChatsService:
		handler = teamschats.NewBackup()

	default:
		return nil, nil, false, clues.Wrap(clues.NewWC(ctx, service.String()), "service not supported")
	}

	colls, excludeItems, canUsePreviousBackup, err = handler.ProduceBackupCollections(
		ctx,
		bpc,
		ctrl.AC,
		ctrl.credentials,
		ctrl.UpdateStatus,
		counter,
		errs)
	if err != nil {
		return nil, nil, false, err
	}

	for _, c := range colls {
		// kopia doesn't stream Items() from deleted collections,
		// and so they never end up calling the UpdateStatus closer.
		// This is a brittle workaround, since changes in consumer
		// behavior (such as calling Items()) could inadvertently
		// break the process state, putting us into deadlock or
		// panics.
		if c.State() != data.DeletedState {
			ctrl.incrementAwaitingMessages()
		}
	}

	return colls, excludeItems, canUsePreviousBackup, nil
}

func (ctrl *Controller) IsServiceEnabled(
	ctx context.Context,
	service path.ServiceType,
	resourceOwner string,
) (bool, error) {
	switch service {
	case path.ExchangeService:
		return exchange.IsServiceEnabled(ctx, ctrl.AC.Users(), resourceOwner)
	case path.OneDriveService:
		return onedrive.IsServiceEnabled(ctx, ctrl.AC.Users(), resourceOwner)
	case path.SharePointService:
		return sharepoint.IsServiceEnabled(ctx, ctrl.AC.Sites(), resourceOwner)
	case path.GroupsService:
		return groups.IsServiceEnabled(ctx, ctrl.AC.Groups(), resourceOwner)
	case path.TeamsChatsService:
		return teamschats.IsServiceEnabled(ctx, ctrl.AC.Users(), resourceOwner)
	}

	return false, clues.Wrap(clues.NewWC(ctx, service.String()), "service not supported")
}

func verifyBackupInputs(sel selectors.Selector, cachedIDs []string) error {
	var ids []string

	switch sel.Service {
	case selectors.ServiceExchange, selectors.ServiceOneDrive:
		// Exchange and OneDrive user existence now checked in checkServiceEnabled.
		return nil

	case selectors.ServiceSharePoint, selectors.ServiceGroups, selectors.ServiceTeamsChats:
		ids = cachedIDs
	}

	if !filters.Contains(ids).Compare(sel.ID()) {
		return clues.Stack(core.ErrNotFound).With("selector_protected_resource", sel.ID())
	}

	return nil
}

func (ctrl *Controller) GetMetadataPaths(
	ctx context.Context,
	r kinject.RestoreProducer,
	base inject.ReasonAndSnapshotIDer,
	errs *fault.Bus,
) ([]path.RestorePaths, error) {
	var (
		paths = []path.RestorePaths{}
		err   error
	)

	for _, reason := range base.GetReasons() {
		filePaths := [][]string{}

		switch true {
		case reason.Service() == path.GroupsService && reason.Category() == path.LibrariesCategory:
			filePaths, err = groups.MetadataFiles(ctx, reason, r, base.GetSnapshotID(), errs)
			if err != nil {
				return nil, err
			}
		case reason.Service() == path.SharePointService && reason.Category() == path.ListsCategory:
			for _, fn := range sharepoint.ListsMetadataFileNames() {
				filePaths = append(filePaths, []string{fn})
			}
		default:
			for _, fn := range bupMD.AllMetadataFileNames() {
				filePaths = append(filePaths, []string{fn})
			}
		}

		for _, fp := range filePaths {
			pth, err := path.BuildMetadata(
				reason.Tenant(),
				reason.ProtectedResource(),
				reason.Service(),
				reason.Category(),
				true,
				fp...)
			if err != nil {
				return nil, err
			}

			dir, err := pth.Dir()
			if err != nil {
				return nil, clues.
					Wrap(err, "building metadata collection path").
					With("metadata_file", fp)
			}

			paths = append(paths, path.RestorePaths{StoragePath: pth, RestorePath: dir})
		}
	}

	return paths, nil
}

func (ctrl *Controller) SetRateLimiter(
	ctx context.Context,
	service path.ServiceType,
	options control.Options,
) context.Context {
	// Use sliding window limiter for Exchange if the feature is not explicitly
	// disabled. For other services we always use token bucket limiter.
	enableSlidingLim := false
	if service == path.ExchangeService &&
		!options.ToggleFeatures.DisableSlidingWindowLimiter {
		enableSlidingLim = true
	}

	ctx = graph.BindRateLimiterConfig(
		ctx,
		graph.LimiterCfg{
			Service:              service,
			EnableSlidingLimiter: enableSlidingLim,
		})

	return ctx
}
