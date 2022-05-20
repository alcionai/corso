package repository

import (
	"context"
	"time"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type repoProvider int

//go:generate stringer -type=repoProvider
const (
	ProviderUnknown repoProvider = iota // Unknown Provider
	ProviderS3                          // S3
)

// Repository represents backup storage for a specific M365 Account.
type Repository struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Version   string // in case of future breaking changes

	Provider repoProvider // must be repository.S3Provider
	Account  Account      // the user's m365 account connection details
	Config   Configurer   // provider-based configuration details
}

// Account holds the user's M365 account details.
type Account struct {
	TenantID     string
	ClientID     string
	ClientSecret string
}

type (
	Configurer interface {
		kopia.StorageMaker
	}

	connector interface {
		Connect(ctx context.Context) error
	}

	initConnector interface {
		connector
		Initialize(ctx context.Context) error
	}
)

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
	cfg Configurer,
) (Repository, error) {
	r := Repository{
		ID:       uuid.New(),
		Version:  "v1",
		Provider: provider,
		Account:  acct,
		Config:   cfg,
	}
	kInit, err := kopia.NewInitializer(ctx, cfg, nil, nil)
	if err != nil {
		return r, errors.Wrap(err, "preparing repository initialization")
	}
	return r, initializeWith(ctx, kInit)
}

func initializeWith(ctx context.Context, ic initConnector) error {
	return ic.Initialize(ctx)
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
	cfg Configurer,
) (Repository, error) {
	// todo: ID and CreatedAt should get retrieved from a stored kopia config.
	r := Repository{
		Version:  "v1",
		Provider: provider,
		Account:  acct,
		Config:   cfg,
	}
	kConn, err := kopia.NewConnector(ctx, cfg, nil)
	if err != nil {
		return r, errors.Wrap(err, "preparing repository connection")
	}
	return r, connectWith(ctx, kConn)
}

func connectWith(ctx context.Context, ic initConnector) error {
	return ic.Connect(ctx)
}
