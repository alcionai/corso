package repository

import (
	"time"

	"github.com/google/uuid"
)

type repoProvider int

const (
	UnknownProvider repoProvider = iota
	S3Provider
)

// the repository properies used by all providers.
type repo struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Version   string // in case of future breaking changes

	Account Account // the user's m365 account connection details
	Config  any     // non-secret, provider-agnostic configuration details
}

// Account holds the user's m365 account details.
type Account struct {
	TenantID     string
	ClientID     string
	ClientSecret string
}

func newRepo(tenantID, clientID, clientSecret string) repo {
	return repo{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		Version:   "v1",
		Account: Account{
			TenantID:     tenantID,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		Config: nil,
	}
}
