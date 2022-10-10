package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
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
			ctx, flush := tester.NewContext()
			defer flush()

			st, err := test.storage()
			assert.NoError(t, err)
			_, err = repository.Initialize(ctx, test.account, st, control.Options{})
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
			ctx, flush := tester.NewContext()
			defer flush()

			st, err := test.storage()
			assert.NoError(t, err)
			_, err = repository.Connect(ctx, test.account, st, control.Options{})
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
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoRepositoryTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(RepositoryIntegrationSuite))
}

// ensure all required env values are populated
func (suite *RepositoryIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(suite.T(), err)
}

func (suite *RepositoryIntegrationSuite) TestInitialize() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name     string
		account  account.Account
		storage  func(*testing.T) storage.Storage
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name:     "success",
			storage:  tester.NewPrefixedS3Storage,
			errCheck: assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			st := test.storage(t)
			r, err := repository.Initialize(ctx, test.account, st, control.Options{})
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
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	_, err := repository.Initialize(ctx, account.Account{}, st, control.Options{})
	require.NoError(t, err)

	// now re-connect
	_, err = repository.Connect(ctx, account.Account{}, st, control.Options{})
	assert.NoError(t, err)
}

func (suite *RepositoryIntegrationSuite) TestNewBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, acct, st, control.Options{})
	require.NoError(t, err)

	bo, err := r.NewBackup(ctx, selectors.Selector{})
	require.NoError(t, err)
	require.NotNil(t, bo)
}

func (suite *RepositoryIntegrationSuite) TestNewRestore() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)
	dest := tester.DefaultTestRestoreDestination()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, acct, st, control.Options{})
	require.NoError(t, err)

	ro, err := r.NewRestore(ctx, "backup-id", selectors.Selector{}, dest)
	require.NoError(t, err)
	require.NotNil(t, ro)
}
