package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/internal/operations"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/storage"
)

// Repository contains storage provider information.
type Repository struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Version   string // in case of future breaking changes

	Account    account.Account // the user's m365 account connection details
	Storage    storage.Storage // the storage provider details and configuration
	dataLayer  *kopia.Wrapper
	modelStore *kopia.ModelStore
}

// Initialize will:
//  * validate the m365 account & secrets
//  * connect to the m365 account to ensure communication capability
//  * validate the provider config & secrets
//  * initialize the kopia repo with the provider
//  * store the configuration details
//  * connect to the provider
//  * return the connected repository
func Initialize(
	ctx context.Context,
	acct account.Account,
	storage storage.Storage,
) (*Repository, error) {
	kopiaRef := kopia.NewConn(storage)
	if err := kopiaRef.Initialize(ctx); err != nil {
		return nil, err
	}
	// kopiaRef comes with a count of 1 and NewWrapper/NewModelStore bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	w, err := kopia.NewWrapper(kopiaRef)
	if err != nil {
		return nil, err
	}

	ms, err := kopia.NewModelStore(kopiaRef)
	if err != nil {
		return nil, err
	}

	r := Repository{
		ID:         uuid.New(),
		Version:    "v1",
		Account:    acct,
		Storage:    storage,
		dataLayer:  w,
		modelStore: ms,
	}
	return &r, nil
}

// Connect will:
//  * validate the m365 account details
//  * connect to the m365 account to ensure communication capability
//  * connect to the provider storage
//  * return the connected repository
func Connect(
	ctx context.Context,
	acct account.Account,
	storage storage.Storage,
) (*Repository, error) {
	kopiaRef := kopia.NewConn(storage)
	if err := kopiaRef.Connect(ctx); err != nil {
		return nil, err
	}
	// kopiaRef comes with a count of 1 and NewWrapper/NewModelStore bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	w, err := kopia.NewWrapper(kopiaRef)
	if err != nil {
		return nil, err
	}

	ms, err := kopia.NewModelStore(kopiaRef)
	if err != nil {
		return nil, err
	}

	// todo: ID and CreatedAt should get retrieved from a stored kopia config.
	r := Repository{
		Version:    "v1",
		Account:    acct,
		Storage:    storage,
		dataLayer:  w,
		modelStore: ms,
	}
	return &r, nil
}

func (r *Repository) Close(ctx context.Context) error {
	if r.dataLayer != nil {
		err := r.dataLayer.Close(ctx)
		r.dataLayer = nil
		if err != nil {
			return errors.Wrap(err, "closing corso DataLayer")
		}
	}

	if r.modelStore == nil {
		return nil
	}
	err := r.modelStore.Close(ctx)
	r.modelStore = nil
	return errors.Wrap(err, "closing corso ModelStore")
}

// NewBackup generates a BackupOperation runner.
func (r Repository) NewBackup(ctx context.Context, selector selectors.Selector) (operations.BackupOperation, error) {
	return operations.NewBackupOperation(
		ctx,
		operations.Options{},
		r.dataLayer,
		r.modelStore,
		r.Account,
		selector)
}

// NewRestore generates a restoreOperation runner.
func (r Repository) NewRestore(ctx context.Context, backupID string, sel selectors.Selector) (operations.RestoreOperation, error) {
	return operations.NewRestoreOperation(
		ctx,
		operations.Options{},
		r.dataLayer,
		r.modelStore,
		r.Account,
		model.ID(backupID),
		sel)
}

// backups lists backups in a respository
func (r Repository) Backups(ctx context.Context) ([]*backup.Backup, error) {
	bms, err := r.modelStore.GetIDsForType(ctx, kopia.BackupModel, nil)
	if err != nil {
		return nil, err
	}
	rps := make([]*backup.Backup, 0, len(bms))
	for _, bm := range bms {
		rp := backup.Backup{}
		err := r.modelStore.GetWithModelStoreID(ctx, kopia.BackupModel, bm.ModelStoreID, &rp)
		if err != nil {
			return nil, err
		}
		rps = append(rps, &rp)
	}
	return rps, nil
}

// BackupDetails returns the specified backup details object
func (r Repository) BackupDetails(ctx context.Context, rpDetailsID string) (*backup.Details, error) {
	rpd := backup.Details{}
	err := r.modelStore.GetWithModelStoreID(ctx, kopia.BackupDetailsModel, manifest.ID(rpDetailsID), &rpd)
	if err != nil {
		return nil, err
	}
	return &rpd, nil
}
