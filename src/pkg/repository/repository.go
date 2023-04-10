package repository

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
	"github.com/alcionai/corso/src/pkg/store"
)

var (
	ErrorRepoAlreadyExists = clues.New("a repository was already initialized with that configuration")
	ErrorBackupNotFound    = clues.New("no backup exists with that id")
)

// BackupGetter deals with retrieving metadata about backups from the
// repository.
type BackupGetter interface {
	Backup(ctx context.Context, id string) (*backup.Backup, error)
	Backups(ctx context.Context, ids []string) ([]*backup.Backup, *fault.Bus)
	BackupsByTag(ctx context.Context, fs ...store.FilterOption) ([]*backup.Backup, error)
	GetBackupDetails(
		ctx context.Context,
		backupID string,
	) (*details.Details, *backup.Backup, *fault.Bus)
	GetBackupErrors(
		ctx context.Context,
		backupID string,
	) (*fault.Errors, *backup.Backup, *fault.Bus)
}

type Repository interface {
	GetID() string
	Close(context.Context) error
	NewBackup(
		ctx context.Context,
		self selectors.Selector,
	) (operations.BackupOperation, error)
	NewBackupWithLookup(
		ctx context.Context,
		self selectors.Selector,
		ins common.IDNameSwapper,
	) (operations.BackupOperation, error)
	NewRestore(
		ctx context.Context,
		backupID string,
		sel selectors.Selector,
		dest control.RestoreDestination,
	) (operations.RestoreOperation, error)
	DeleteBackup(ctx context.Context, id string) error
	BackupGetter
}

// Repository contains storage provider information.
type repository struct {
	ID        string
	CreatedAt time.Time
	Version   string // in case of future breaking changes

	Account account.Account // the user's m365 account connection details
	Storage storage.Storage // the storage provider details and configuration
	Opts    control.Options

	Bus        events.Eventer
	dataLayer  *kopia.Wrapper
	modelStore *kopia.ModelStore
}

func (r repository) GetID() string {
	return r.ID
}

// Initialize will:
//   - validate the m365 account & secrets
//   - connect to the m365 account to ensure communication capability
//   - validate the provider config & secrets
//   - initialize the kopia repo with the provider
//   - store the configuration details
//   - connect to the provider
//   - return the connected repository
func Initialize(
	ctx context.Context,
	acct account.Account,
	s storage.Storage,
	opts control.Options,
) (repo Repository, err error) {
	ctx = clues.Add(
		ctx,
		"acct_provider", acct.Provider.String(),
		"acct_id", clues.Hide(acct.ID()),
		"storage_provider", s.Provider.String())

	defer func() {
		if crErr := crash.Recovery(ctx, recover()); crErr != nil {
			err = crErr
		}
	}()

	kopiaRef := kopia.NewConn(s)
	if err := kopiaRef.Initialize(ctx); err != nil {
		// replace common internal errors so that sdk users can check results with errors.Is()
		if errors.Is(err, kopia.ErrorRepoAlreadyExists) {
			return nil, clues.Stack(ErrorRepoAlreadyExists, err).WithClues(ctx)
		}

		return nil, clues.Wrap(err, "initializing kopia")
	}
	// kopiaRef comes with a count of 1 and NewWrapper/NewModelStore bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	w, err := kopia.NewWrapper(kopiaRef)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	ms, err := kopia.NewModelStore(kopiaRef)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	bus, err := events.NewBus(ctx, s, acct.ID(), opts)
	if err != nil {
		return nil, clues.Wrap(err, "constructing event bus")
	}

	repoID := newRepoID(s)
	bus.SetRepoID(repoID)

	r := &repository{
		ID:         repoID,
		Version:    "v1",
		Account:    acct,
		Storage:    s,
		Bus:        bus,
		Opts:       opts,
		dataLayer:  w,
		modelStore: ms,
	}

	if err := newRepoModel(ctx, ms, r.ID); err != nil {
		return nil, clues.New("setting up repository").WithClues(ctx)
	}

	r.Bus.Event(ctx, events.RepoInit, nil)

	return r, nil
}

// Connect will:
//   - validate the m365 account details
//   - connect to the m365 account to ensure communication capability
//   - connect to the provider storage
//   - return the connected repository
func Connect(
	ctx context.Context,
	acct account.Account,
	s storage.Storage,
	opts control.Options,
) (r Repository, err error) {
	ctx = clues.Add(
		ctx,
		"acct_provider", acct.Provider.String(),
		"acct_id", clues.Hide(acct.ID()),
		"storage_provider", s.Provider.String())

	defer func() {
		if crErr := crash.Recovery(ctx, recover()); crErr != nil {
			err = crErr
		}
	}()

	// Close/Reset the progress bar. This ensures callers don't have to worry about
	// their output getting clobbered (#1720)
	defer observe.Complete()

	complete, closer := observe.MessageWithCompletion(ctx, "Connecting to repository")
	defer closer()
	defer close(complete)

	kopiaRef := kopia.NewConn(s)
	if err := kopiaRef.Connect(ctx); err != nil {
		return nil, clues.Wrap(err, "connecting kopia client")
	}
	// kopiaRef comes with a count of 1 and NewWrapper/NewModelStore bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	w, err := kopia.NewWrapper(kopiaRef)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	ms, err := kopia.NewModelStore(kopiaRef)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	bus, err := events.NewBus(ctx, s, acct.ID(), opts)
	if err != nil {
		return nil, clues.Wrap(err, "constructing event bus")
	}

	rm := &repositoryModel{}

	// Do not query repo ID if metrics are disabled
	if !opts.DisableMetrics {
		rm, err = getRepoModel(ctx, ms)
		if err != nil {
			return nil, clues.New("retrieving repo info")
		}

		bus.SetRepoID(string(rm.ID))
	}

	complete <- struct{}{}

	// todo: ID and CreatedAt should get retrieved from a stored kopia config.
	return &repository{
		ID:         string(rm.ID),
		Version:    "v1",
		Account:    acct,
		Storage:    s,
		Bus:        bus,
		Opts:       opts,
		dataLayer:  w,
		modelStore: ms,
	}, nil
}

func ConnectAndSendConnectEvent(ctx context.Context,
	acct account.Account,
	s storage.Storage,
	opts control.Options,
) (Repository, error) {
	repo, err := Connect(ctx, acct, s, opts)
	if err != nil {
		return nil, err
	}

	r := repo.(*repository)
	r.Bus.Event(ctx, events.RepoConnect, nil)

	return r, nil
}

func (r *repository) Close(ctx context.Context) error {
	if err := r.Bus.Close(); err != nil {
		logger.Ctx(ctx).With("err", err).Debugw("closing the event bus", clues.In(ctx).Slice()...)
	}

	if r.dataLayer != nil {
		if err := r.dataLayer.Close(ctx); err != nil {
			logger.Ctx(ctx).With("err", err).Debugw("closing Datalayer", clues.In(ctx).Slice()...)
		}

		r.dataLayer = nil
	}

	if r.modelStore != nil {
		if err := r.modelStore.Close(ctx); err != nil {
			logger.Ctx(ctx).With("err", err).Debugw("closing modelStore", clues.In(ctx).Slice()...)
		}

		r.modelStore = nil
	}

	return nil
}

// NewBackup generates a BackupOperation runner.
func (r repository) NewBackup(
	ctx context.Context,
	sel selectors.Selector,
) (operations.BackupOperation, error) {
	return r.NewBackupWithLookup(ctx, sel, nil)
}

// NewBackupWithLookup generates a BackupOperation runner.
// ownerIDToName and ownerNameToID are optional populations, in case the caller has
// already generated those values.
func (r repository) NewBackupWithLookup(
	ctx context.Context,
	sel selectors.Selector,
	ins common.IDNameSwapper,
) (operations.BackupOperation, error) {
	gc, err := connectToM365(ctx, sel, r.Account, fault.New(true))
	if err != nil {
		return operations.BackupOperation{}, errors.Wrap(err, "connecting to m365")
	}

	ownerID, ownerName, err := gc.PopulateOwnerIDAndNamesFrom(ctx, sel.DiscreteOwner, ins)
	if err != nil {
		return operations.BackupOperation{}, errors.Wrap(err, "resolving resource owner details")
	}

	// Exchange and OneDrive need to maintain the user PN as the ID until we're ready to migrate
	if sel.PathService() != path.SharePointService {
		ownerID = ownerName
	}

	// TODO: retrieve display name from gc
	sel = sel.SetDiscreteOwnerIDName(ownerID, ownerName)

	return operations.NewBackupOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewKopiaStore(r.modelStore),
		gc,
		r.Account,
		sel,
		sel,
		r.Bus)
}

// NewRestore generates a restoreOperation runner.
func (r repository) NewRestore(
	ctx context.Context,
	backupID string,
	sel selectors.Selector,
	dest control.RestoreDestination,
) (operations.RestoreOperation, error) {
	gc, err := connectToM365(ctx, sel, r.Account, fault.New(true))
	if err != nil {
		return operations.RestoreOperation{}, errors.Wrap(err, "connecting to m365")
	}

	return operations.NewRestoreOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewKopiaStore(r.modelStore),
		gc,
		r.Account,
		model.StableID(backupID),
		sel,
		dest,
		r.Bus)
}

// Backup retrieves a backup by id.
func (r repository) Backup(ctx context.Context, id string) (*backup.Backup, error) {
	return getBackup(ctx, id, store.NewKopiaStore(r.modelStore))
}

// getBackup handles the processing for Backup.
func getBackup(
	ctx context.Context,
	id string,
	sw store.BackupGetter,
) (*backup.Backup, error) {
	b, err := sw.GetBackup(ctx, model.StableID(id))
	if err != nil {
		return nil, errWrapper(err)
	}

	return b, nil
}

// BackupsByID lists backups by ID. Returns as many backups as possible with
// errors for the backups it was unable to retrieve.
func (r repository) Backups(ctx context.Context, ids []string) ([]*backup.Backup, *fault.Bus) {
	var (
		bups []*backup.Backup
		errs = fault.New(false)
		sw   = store.NewKopiaStore(r.modelStore)
	)

	for _, id := range ids {
		ictx := clues.Add(ctx, "backup_id", id)

		b, err := sw.GetBackup(ictx, model.StableID(id))
		if err != nil {
			errs.AddRecoverable(errWrapper(err))
		}

		bups = append(bups, b)
	}

	return bups, errs
}

// backups lists backups in a repository
func (r repository) BackupsByTag(ctx context.Context, fs ...store.FilterOption) ([]*backup.Backup, error) {
	sw := store.NewKopiaStore(r.modelStore)
	return sw.GetBackups(ctx, fs...)
}

// BackupDetails returns the specified backup.Details
func (r repository) GetBackupDetails(
	ctx context.Context,
	backupID string,
) (*details.Details, *backup.Backup, *fault.Bus) {
	errs := fault.New(false)

	deets, bup, err := getBackupDetails(
		ctx,
		backupID,
		r.Account.ID(),
		r.dataLayer,
		store.NewKopiaStore(r.modelStore),
		errs)

	return deets, bup, errs.Fail(err)
}

// getBackupDetails handles the processing for GetBackupDetails.
func getBackupDetails(
	ctx context.Context,
	backupID, tenantID string,
	kw *kopia.Wrapper,
	sw store.BackupGetter,
	errs *fault.Bus,
) (*details.Details, *backup.Backup, error) {
	b, err := sw.GetBackup(ctx, model.StableID(backupID))
	if err != nil {
		return nil, nil, errWrapper(err)
	}

	ssid := b.StreamStoreID
	if len(ssid) == 0 {
		ssid = b.DetailsID
	}

	if len(ssid) == 0 {
		return nil, b, clues.New("no streamstore id in backup").WithClues(ctx)
	}

	var (
		sstore = streamstore.NewStreamer(kw, tenantID, b.Selector.PathService())
		deets  details.Details
	)

	err = sstore.Read(
		ctx,
		ssid,
		streamstore.DetailsReader(details.UnmarshalTo(&deets)),
		errs)
	if err != nil {
		return nil, nil, err
	}

	// Retroactively fill in isMeta information for items in older
	// backup versions without that info
	// version.Restore2 introduces the IsMeta flag, so only v1 needs a check.
	if b.Version >= version.OneDrive1DataAndMetaFiles && b.Version < version.OneDrive3IsMetaMarker {
		for _, d := range deets.Entries {
			if d.OneDrive != nil {
				d.OneDrive.IsMeta = onedrive.IsMetaFile(d.RepoRef)
			}
		}
	}

	deets.DetailsModel = deets.FilterMetaFiles()

	return &deets, b, nil
}

// BackupErrors returns the specified backup's fault.Errors
func (r repository) GetBackupErrors(
	ctx context.Context,
	backupID string,
) (*fault.Errors, *backup.Backup, *fault.Bus) {
	errs := fault.New(false)

	fe, bup, err := getBackupErrors(
		ctx,
		backupID,
		r.Account.ID(),
		r.dataLayer,
		store.NewKopiaStore(r.modelStore),
		errs)

	return fe, bup, errs.Fail(err)
}

// getBackupErrors handles the processing for GetBackupErrors.
func getBackupErrors(
	ctx context.Context,
	backupID, tenantID string,
	kw *kopia.Wrapper,
	sw store.BackupGetter,
	errs *fault.Bus,
) (*fault.Errors, *backup.Backup, error) {
	b, err := sw.GetBackup(ctx, model.StableID(backupID))
	if err != nil {
		return nil, nil, errWrapper(err)
	}

	ssid := b.StreamStoreID
	if len(ssid) == 0 {
		return nil, b, clues.New("no errors in backup").WithClues(ctx)
	}

	var (
		sstore = streamstore.NewStreamer(kw, tenantID, b.Selector.PathService())
		fe     fault.Errors
	)

	err = sstore.Read(
		ctx,
		ssid,
		streamstore.FaultErrorsReader(fault.UnmarshalErrorsTo(&fe)),
		errs)
	if err != nil {
		return nil, nil, err
	}

	return &fe, b, nil
}

type snapshotDeleter interface {
	DeleteSnapshot(ctx context.Context, snapshotID string) error
}

// DeleteBackup removes the backup from both the model store and the backup storage.
func (r repository) DeleteBackup(ctx context.Context, id string) error {
	return deleteBackup(ctx, id, r.dataLayer, store.NewKopiaStore(r.modelStore))
}

// deleteBackup handles the processing for Backup.
func deleteBackup(
	ctx context.Context,
	id string,
	kw snapshotDeleter,
	sw store.BackupGetterDeleter,
) error {
	b, err := sw.GetBackup(ctx, model.StableID(id))
	if err != nil {
		return errWrapper(err)
	}

	if err := kw.DeleteSnapshot(ctx, b.SnapshotID); err != nil {
		return err
	}

	if len(b.SnapshotID) > 0 {
		if err := kw.DeleteSnapshot(ctx, b.SnapshotID); err != nil {
			return err
		}
	}

	if len(b.DetailsID) > 0 {
		if err := kw.DeleteSnapshot(ctx, b.DetailsID); err != nil {
			return err
		}
	}

	return sw.DeleteBackup(ctx, model.StableID(id))
}

// ---------------------------------------------------------------------------
// Repository ID Model
// ---------------------------------------------------------------------------

// repositoryModel identifies the current repository
type repositoryModel struct {
	model.BaseModel
}

// should only be called on init.
func newRepoModel(ctx context.Context, ms *kopia.ModelStore, repoID string) error {
	rm := repositoryModel{
		BaseModel: model.BaseModel{
			ID: model.StableID(repoID),
		},
	}

	return ms.Put(ctx, model.RepositorySchema, &rm)
}

// retrieves the repository info
func getRepoModel(ctx context.Context, ms *kopia.ModelStore) (*repositoryModel, error) {
	bms, err := ms.GetIDsForType(ctx, model.RepositorySchema, nil)
	if err != nil {
		return nil, err
	}

	rm := &repositoryModel{}
	if len(bms) == 0 {
		return rm, nil
	}

	rm.BaseModel = *bms[0]

	return rm, nil
}

// newRepoID generates a new unique repository id hash.
// Repo IDs should only be generated once per repository,
// and must be stored after that.
func newRepoID(s storage.Storage) string {
	return uuid.NewString()
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// produces a graph connector.
func connectToM365(
	ctx context.Context,
	sel selectors.Selector,
	acct account.Account,
	errs *fault.Bus,
) (*connector.GraphConnector, error) {
	complete, closer := observe.MessageWithCompletion(ctx, "Connecting to M365")
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	// retrieve data from the producer
	resource := connector.Users
	if sel.Service == selectors.ServiceSharePoint {
		resource = connector.Sites
	}

	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, resource, errs)
	if err != nil {
		return nil, err
	}

	return gc, nil
}

func errWrapper(err error) error {
	if errors.Is(err, data.ErrNotFound) {
		return clues.Stack(ErrorBackupNotFound, err)
	}

	return err
}
