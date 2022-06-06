package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/repository"
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
		account  repository.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			func() (storage.Storage, error) {
				return storage.NewStorage(storage.ProviderUnknown)
			},
			repository.Account{},
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
		account  repository.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			func() (storage.Storage, error) {
				return storage.NewStorage(storage.ProviderUnknown)
			},
			repository.Account{},
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
		ctesting.CORSO_CI_TESTS,
		ctesting.CORSO_REPOSITORY_TESTS,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(RepositoryIntegrationSuite))
}

// ensure all required env values are populated
func (suite *RepositoryIntegrationSuite) SetupSuite() {
	require.NoError(suite.T(), ctesting.CheckS3EnvVars())
}

func (suite *RepositoryIntegrationSuite) TestInitialize() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())

	table := []struct {
		prefix   string
		account  repository.Account
		storage  func() (storage.Storage, error)
		errCheck assert.ErrorAssertionFunc
	}{
		{
			prefix: "init-s3-" + timeOfTest,
			storage: func() (storage.Storage, error) {
				return ctesting.NewS3Storage("init-s3-" + timeOfTest)
			},
			errCheck: assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.prefix, func(t *testing.T) {
			st, err := test.storage()
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
