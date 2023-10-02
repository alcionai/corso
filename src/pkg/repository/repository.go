package repository

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlRepo "github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/storage"
	"github.com/alcionai/corso/src/pkg/store"
)

const NewRepoID = ""

var (
	ErrorRepoAlreadyExists = clues.New("a repository was already initialized with that configuration")
	ErrorBackupNotFound    = clues.New("no backup exists with that id")
)

type Repositoryer interface {
	Backuper
	BackupGetter
	Restorer
	Exporter
	DataProviderConnector

	Initialize(
		ctx context.Context,
		cfg InitConfig,
	) error
	Connect(
		ctx context.Context,
		cfg ConnConfig,
	) error
	GetID() string
	Close(context.Context) error

	NewMaintenance(
		ctx context.Context,
		mOpts ctrlRepo.Maintenance,
	) (operations.MaintenanceOperation, error)
	NewRetentionConfig(
		ctx context.Context,
		rcOpts ctrlRepo.Retention,
	) (operations.RetentionConfigOperation, error)
}

// Repository contains storage provider information.
type repository struct {
	ID        string
	CreatedAt time.Time
	Version   string // in case of future breaking changes

	Account  account.Account // the user's m365 account connection details
	Storage  storage.Storage // the storage provider details and configuration
	Opts     control.Options
	Provider DataProvider // the client controller used for external user data CRUD

	Bus        events.Eventer
	dataLayer  *kopia.Wrapper
	modelStore *kopia.ModelStore
}

func (r repository) GetID() string {
	return r.ID
}

// New constructs a repository that can be used to Initialize or Connect a repo instance.
func New(
	ctx context.Context,
	acct account.Account,
	st storage.Storage,
	opts control.Options,
	configFileRepoID string,
) (repo *repository, err error) {
	ctx = clues.Add(
		ctx,
		"acct_provider", acct.Provider.String(),
		"acct_id", clues.Hide(acct.ID()),
		"storage_provider", st.Provider.String())

	bus, err := events.NewBus(ctx, st, acct.ID(), opts)
	if err != nil {
		return nil, clues.Wrap(err, "constructing event bus").WithClues(ctx)
	}

	repoID := configFileRepoID
	if len(configFileRepoID) == 0 {
		repoID = newRepoID(st)
	}

	bus.SetRepoID(repoID)

	r := repository{
		ID:      repoID,
		Version: "v1",
		Account: acct,
		Storage: st,
		Bus:     bus,
		Opts:    opts,
	}

	if !r.Opts.DisableMetrics {
		bus.SetRepoID(r.ID)
	}

	return &r, nil
}

type InitConfig struct {
	// tells the data provider which service to
	// use for its connection pattern.  Optional.
	Service       path.ServiceType
	RetentionOpts ctrlRepo.Retention
}

// Initialize will:
//   - connect to the m365 account to ensure communication capability
//   - initialize the kopia repo with the provider and retention parameters
//   - update maintenance retention parameters as needed
//   - store the configuration details
//   - connect to the provider
func (r *repository) Initialize(
	ctx context.Context,
	cfg InitConfig,
) (err error) {
	ctx = clues.Add(
		ctx,
		"acct_provider", r.Account.Provider.String(),
		"acct_id", clues.Hide(r.Account.ID()),
		"storage_provider", r.Storage.Provider.String())

	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "repo init"); crErr != nil {
			err = crErr
		}
	}()

	if err := r.ConnectDataProvider(ctx, cfg.Service); err != nil {
		return clues.Stack(err)
	}

	observe.Message(ctx, "Initializing repository")

	kopiaRef := kopia.NewConn(r.Storage)
	if err := kopiaRef.Initialize(ctx, r.Opts.Repo, cfg.RetentionOpts); err != nil {
		// replace common internal errors so that sdk users can check results with errors.Is()
		if errors.Is(err, kopia.ErrorRepoAlreadyExists) {
			return clues.Stack(ErrorRepoAlreadyExists, err).WithClues(ctx)
		}

		return clues.Wrap(err, "initializing kopia")
	}
	// kopiaRef comes with a count of 1 and NewWrapper/NewModelStore bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	r.dataLayer, err = kopia.NewWrapper(kopiaRef)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	r.modelStore, err = kopia.NewModelStore(kopiaRef)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	if err := newRepoModel(ctx, r.modelStore, r.ID); err != nil {
		return clues.Wrap(err, "setting up repository").WithClues(ctx)
	}

	r.Bus.Event(ctx, events.RepoInit, nil)

	return nil
}

type ConnConfig struct {
	// tells the data provider which service to
	// use for its connection pattern.  Leave empty
	// to skip the provider connection.
	Service path.ServiceType
}

// Connect will:
//   - connect to the m365 account
//   - connect to the provider storage
//   - return the connected repository
func (r *repository) Connect(
	ctx context.Context,
	cfg ConnConfig,
) (err error) {
	ctx = clues.Add(
		ctx,
		"acct_provider", r.Account.Provider.String(),
		"acct_id", clues.Hide(r.Account.ID()),
		"storage_provider", r.Storage.Provider.String())

	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "repo connect"); crErr != nil {
			err = crErr
		}
	}()

	if err := r.ConnectDataProvider(ctx, cfg.Service); err != nil {
		return clues.Stack(err)
	}

	observe.Message(ctx, "Connecting to repository")

	kopiaRef := kopia.NewConn(r.Storage)
	if err := kopiaRef.Connect(ctx, r.Opts.Repo); err != nil {
		return clues.Wrap(err, "connecting kopia client")
	}
	// kopiaRef comes with a count of 1 and NewWrapper/NewModelStore bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	r.dataLayer, err = kopia.NewWrapper(kopiaRef)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	r.modelStore, err = kopia.NewModelStore(kopiaRef)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	if r.ID == events.RepoIDNotFound {
		rm, err := getRepoModel(ctx, r.modelStore)
		if err != nil {
			return clues.Wrap(err, "retrieving repo model info")
		}

		r.ID = string(rm.ID)
	}

	r.Bus.Event(ctx, events.RepoConnect, nil)

	return nil
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

func (r repository) NewMaintenance(
	ctx context.Context,
	mOpts ctrlRepo.Maintenance,
) (operations.MaintenanceOperation, error) {
	return operations.NewMaintenanceOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		mOpts,
		r.Bus)
}

func (r repository) NewRetentionConfig(
	ctx context.Context,
	rcOpts ctrlRepo.Retention,
) (operations.RetentionConfigOperation, error) {
	return operations.NewRetentionConfigOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		rcOpts,
		r.Bus)
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

func errWrapper(err error) error {
	if errors.Is(err, data.ErrNotFound) {
		return clues.Stack(ErrorBackupNotFound, err)
	}

	return err
}
