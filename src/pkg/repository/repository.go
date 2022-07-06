package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/operations"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/storage"
)

// Repository contains storage provider information.
type Repository struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Version   string // in case of future breaking changes

	Account   account.Account // the user's m365 account connection details
	Storage   storage.Storage // the storage provider details and configuration
	dataLayer *kopia.Wrapper
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
	// kopiaRef comes with a count of 1 and NewWrapper bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	w, err := kopia.NewWrapper(kopiaRef)
	if err != nil {
		return nil, err
	}

	r := Repository{
		ID:        uuid.New(),
		Version:   "v1",
		Account:   acct,
		Storage:   storage,
		dataLayer: w,
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
	// kopiaRef comes with a count of 1 and NewWrapper bumps it again so safe
	// to close here.
	defer kopiaRef.Close(ctx)

	w, err := kopia.NewWrapper(kopiaRef)
	if err != nil {
		return nil, err
	}

	// todo: ID and CreatedAt should get retrieved from a stored kopia config.
	r := Repository{
		Version:   "v1",
		Account:   acct,
		Storage:   storage,
		dataLayer: w,
	}
	return &r, nil
}

func (r *Repository) Close(ctx context.Context) error {
	if r.dataLayer == nil {
		return nil
	}

	err := r.dataLayer.Close(ctx)
	r.dataLayer = nil

	if err != nil {
		return errors.Wrap(err, "closing corso Repository")
	}

	return nil
}

// NewBackup generates a backupOperation runner.
func (r Repository) NewBackup(ctx context.Context, targets []string) (operations.BackupOperation, error) {
	return operations.NewBackupOperation(
		ctx,
		operations.Options{},
		r.dataLayer,
		r.Account,
		targets)
}

// NewRestore generates a restoreOperation runner.
func (r Repository) NewRestore(ctx context.Context, restorePointID string, targets []string) (operations.RestoreOperation, error) {
	return operations.NewRestoreOperation(
		ctx,
		operations.Options{},
		r.dataLayer,
		r.Account,
		restorePointID,
		targets)
}
