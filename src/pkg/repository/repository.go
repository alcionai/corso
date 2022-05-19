package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/blob"
)

type repoProvider int

const (
	UnknownProvider repoProvider = iota
	S3Provider
)

func (rp repoProvider) Name() string {
	switch rp {
	case S3Provider:
		return "S3"
	default:
		return "Default Provider"
	}
}

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

func newRepo(provider repoProvider, acct Account, cfg Config) Repository {
	return Repository{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		Version:   "v1",
		Provider:  provider,
		Account:   acct,
		Config:    cfg,
	}
}

// Initialize will:
//  * validate the m365 account & secrets
//  * connect to the m365 account to ensure communication capability
//  * validate the provider config & secrets
//  * initialize the kopia repo with the provider
//  * store the configuration details
//  * connect to the provider
func (r Repository) Initialize(ctx context.Context) error {
	return nil
}

// Connect will:
//  * validate the m365 account details
//  * connect to the m365 account to ensure communication capability
//  * connect to the provider storage
func (r Repository) Connect(ctx context.Context) error {
	return nil
}
