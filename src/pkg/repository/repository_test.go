package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/storage"
)

// ---------------
// unit tests
// ---------------

type RepositorySuite struct {
	suite.Suite
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (suite *RepositorySuite) TestInitialize() {
	table := []struct {
		name     string
		storage  func() (storage.Storage, error)
		account  account.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			func() (storage.Storage, error) {
				return storage.NewStorage(storage.ProviderUnknown)
			},
			account.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			st, err := test.storage()
			assert.NoError(t, err)
			_, err = repository.Initialize(context.Background(), test.account, st)
			test.errCheck(t, err, "")
		})
	}
}

// repository.Connect involves end-to-end communication with kopia, therefore this only
// tests expected error cases
func (suite *RepositorySuite) TestConnect() {
	table := []struct {
		name     string
		storage  func() (storage.Storage, error)
		account  account.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			func() (storage.Storage, error) {
				return storage.NewStorage(storage.ProviderUnknown)
			},
			account.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			st, err := test.storage()
			assert.NoError(t, err)
			_, err = repository.Connect(context.Background(), test.account, st)
			test.errCheck(t, err)
		})
	}
}

// ---------------
// integration tests
// ---------------

type RepositoryIntegrationSuite struct {
	suite.Suite
}

func TestRepositoryIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoRepositoryTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(RepositoryIntegrationSuite))
}

// ensure all required env values are populated
func (suite *RepositoryIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(
		append(
			ctesting.AWSStorageCredEnvs,
			ctesting.M365AcctCredEnvs...,
		)...,
	)
	require.NoError(suite.T(), err)
}

func (suite *RepositoryIntegrationSuite) TestInitialize() {
	ctx := context.Background()

	table := []struct {
		name     string
		account  account.Account
		storage  func(*testing.T) (storage.Storage, error)
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name:     "success",
			storage:  ctesting.NewPrefixedS3Storage,
			errCheck: assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			st, err := test.storage(t)
			assert.NoError(t, err)
			r, err := repository.Initialize(ctx, test.account, st)
			if err == nil {
				defer func() {
					assert.NoError(t, r.Close(ctx))
				}()
			}

			test.errCheck(t, err)
		})
	}
}

func (suite *RepositoryIntegrationSuite) TestConnect() {
	t := suite.T()
	ctx := context.Background()

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewPrefixedS3Storage(t)
	require.NoError(t, err)

	_, err = repository.Initialize(ctx, account.Account{}, st)
	require.NoError(t, err)

	// now re-connect
	_, err = repository.Connect(ctx, account.Account{}, st)
	assert.NoError(t, err)
}

func (suite *RepositoryIntegrationSuite) TestNewBackup() {
	t := suite.T()
	ctx := context.Background()

	acct, err := ctesting.NewM365Account()
	require.NoError(t, err)

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewPrefixedS3Storage(t)
	require.NoError(t, err)

	r, err := repository.Initialize(ctx, acct, st)
	require.NoError(t, err)

	bo, err := r.NewBackup(ctx, selectors.Selector{})
	require.NoError(t, err)
	require.NotNil(t, bo)
}

func (suite *RepositoryIntegrationSuite) TestNewRestore() {
	t := suite.T()
	ctx := context.Background()

	acct, err := ctesting.NewM365Account()
	require.NoError(t, err)

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewPrefixedS3Storage(t)
	require.NoError(t, err)

	r, err := repository.Initialize(ctx, acct, st)
	require.NoError(t, err)

	ro, err := r.NewRestore(ctx, "backup-id", selectors.Selector{})
	require.NoError(t, err)
	require.NotNil(t, ro)
}
