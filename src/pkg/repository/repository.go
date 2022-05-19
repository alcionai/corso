package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/blob"
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

	Provider repoProvider // must be repository.S3Provider
	Account  Account      // the user's m365 account connection details
	Config   Config       // provider-based configuration details
}

// Account holds the user's m365 account details.
type Account struct {
	TenantID     string
	ClientID     string
	ClientSecret string
}

type Config interface {
	KopiaStorage(ctx context.Context, create bool) (blob.Storage, error)
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
	provider repoProvider,
	acct Account,
	cfg Config,
) (Repository, error) {
	r := Repository{
		ID:       uuid.New(),
		Version:  "v1",
		Provider: provider,
		Account:  acct,
		Config:   cfg,
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
	provider repoProvider,
	acct Account,
	cfg Config,
) (Repository, error) {
	// todo: ID and CreatedAt should get retrieved from a stored kopia config.
	r := Repository{
		Version:  "v1",
		Provider: provider,
		Account:  acct,
		Config:   cfg,
	}
	return r, nil
}
