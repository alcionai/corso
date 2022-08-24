package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/internal/operations"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/control"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/storage"
	"github.com/alcionai/corso/pkg/store"
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
	s storage.Storage,
) (*Repository, error) {
	kopiaRef := kopia.NewConn(s)
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
		Storage:    s,
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
	s storage.Storage,
) (*Repository, error) {
	kopiaRef := kopia.NewConn(s)
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
		Storage:    s,
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
func (r Repository) NewBackup(
	ctx context.Context,
	selector selectors.Selector,
	opts control.Options,
) (operations.BackupOperation, error) {
	return operations.NewBackupOperation(
		ctx,
		opts,
		r.dataLayer,
		store.NewKopiaStore(r.modelStore),
		r.Account,
		selector)
}

// NewRestore generates a restoreOperation runner.
func (r Repository) NewRestore(
	ctx context.Context,
	backupID string,
	sel selectors.Selector,
	opts control.Options,
) (operations.RestoreOperation, error) {
	return operations.NewRestoreOperation(
		ctx,
		opts,
		r.dataLayer,
		store.NewKopiaStore(r.modelStore),
		r.Account,
		model.StableID(backupID),
		sel)
}

// backups lists a backup by id
func (r Repository) Backup(ctx context.Context, id model.StableID) (*backup.Backup, error) {
	sw := store.NewKopiaStore(r.modelStore)
	return sw.GetBackup(ctx, id)
}

// backups lists backups in a repository
func (r Repository) Backups(ctx context.Context) ([]backup.Backup, error) {
	sw := store.NewKopiaStore(r.modelStore)
	return sw.GetBackups(ctx)
}

// BackupDetails returns the specified backup details object
func (r Repository) BackupDetails(ctx context.Context, backupID string) (*details.Details, *backup.Backup, error) {
	sw := store.NewKopiaStore(r.modelStore)
	return sw.GetDetailsFromBackupID(ctx, model.StableID(backupID))
}

// DeleteBackup removes the backup from both the model store and the backup storage.
func (r Repository) DeleteBackup(ctx context.Context, id model.StableID) error {
	bu, err := r.Backup(ctx, id)
	if err != nil {
		return err
	}

	if err := r.dataLayer.DeleteSnapshot(ctx, bu.SnapshotID); err != nil {
		return err
	}

	sw := store.NewKopiaStore(r.modelStore)
	return sw.DeleteBackup(ctx, id)
}
