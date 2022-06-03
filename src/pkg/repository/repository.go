package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/storage"
)

type repoProvider int

//go:generate stringer -type=repoProvider
const (
	ProviderUnknown repoProvider = iota // Unknown Provider
	ProviderS3                          // S3
)

// Repository contains storage provider information.
type Repository struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Version   string // in case of future breaking changes

	Account   Account         // the user's m365 account connection details
	Storage   storage.Storage // the storage provider details and configuration
	dataLayer *kopia.KopiaWrapper
}

// Account holds the user's m365 account details.
type Account struct {
	TenantID     string
	ClientID     string
	ClientSecret string
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
	acct Account,
	storage storage.Storage,
) (Repository, error) {
	k := kopia.New(storage)
	if err := k.Initialize(ctx); err != nil {
		return Repository{}, err
	}
	r := Repository{
		ID:        uuid.New(),
		Version:   "v1",
		Account:   acct,
		Storage:   storage,
		dataLayer: &k,
	}
	return r, nil
}

// Connect will:
//  * validate the m365 account details
//  * connect to the m365 account to ensure communication capability
//  * connect to the provider storage
//  * return the connected repository
func Connect(
	ctx context.Context,
	acct Account,
	storage storage.Storage,
) (Repository, error) {
	k := kopia.New(storage)
	if err := k.Connect(ctx); err != nil {
		return Repository{}, err
	}
	// todo: ID and CreatedAt should get retrieved from a stored kopia config.
	r := Repository{
		Version:   "v1",
		Account:   acct,
		Storage:   storage,
		dataLayer: &k,
	}
	return r, nil
}

func (r Repository) Close(ctx context.Context) error {
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
