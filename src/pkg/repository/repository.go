package repository

import (
	"context"
	"time"

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
}

// Account holds the user's M365 account details.
type Account struct {
	TenantID     string
	ClientID     string
	ClientSecret string
}

type (
	Connector interface {
		Connect(ctx context.Context) error
	}

	InitConnector interface {
		Connector
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
	cfg InitConnector,
) (Repository, error) {
	// todo: validate m365 creds
	if err := cfg.Initialize(ctx); err != nil {
		return Repository{}, errors.Wrapf(err, "initialializing %s repository", provider)
	}
	return Repository{
		ID:       uuid.New(),
		Version:  "v1",
		Provider: provider,
		Account:  acct,
	}, nil
}

func initializeWith(ctx context.Context, ic InitConnector) error {
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
	cfg Connector,
) (Repository, error) {
	// todo: validate m365 creds
	if err := cfg.Connect(ctx); err != nil {
		return Repository{}, errors.Wrapf(err, "connecting %s repository", provider)
	}
	// todo: ID and CreatedAt should get retrieved from a stored kopia config.
	return Repository{
		Version:  "v1",
		Provider: provider,
		Account:  acct,
	}, nil
}
