package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/credentials"
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
		ctesting.CorsoCITests,
		ctesting.CorsoRepositoryTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(RepositoryIntegrationSuite))
}

// ensure all required env values are populated
func (suite *RepositoryIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(ctesting.AWSCredentialEnvs...)
	require.NoError(suite.T(), err)
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
			prefix: "repository-init-s3-" + timeOfTest,
			storage: func() (storage.Storage, error) {
				return ctesting.NewS3Storage("repository-init-s3-" + timeOfTest)
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

func (suite *RepositoryIntegrationSuite) TestConnect() {
	t := suite.T()
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(t)
	prefix := "repository-conn-s3-" + timeOfTest

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewS3Storage(prefix)
	require.NoError(t, err)

	_, err = repository.Initialize(ctx, repository.Account{}, st)
	require.NoError(t, err)

	// now re-connect
	_, err = repository.Connect(ctx, repository.Account{}, st)
	assert.NoError(t, err)
}

func (suite *RepositoryIntegrationSuite) TestNewBackup() {
	t := suite.T()
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(t)
	prefix := "repository-new-backup-" + timeOfTest

	m365 := credentials.GetM365()
	acct := repository.Account{
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
		TenantID:     m365.TenantID,
	}

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewS3Storage(prefix)
	require.NoError(t, err)

	r, err := repository.Initialize(ctx, acct, st)
	require.NoError(t, err)

	bo, err := r.NewBackup(ctx, []string{})
	require.NoError(t, err)
	require.NotNil(t, bo)
}

func (suite *RepositoryIntegrationSuite) TestNewRestore() {
	t := suite.T()
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(t)
	prefix := "repository-new-restore-" + timeOfTest

	m365 := credentials.GetM365()
	acct := repository.Account{
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
		TenantID:     m365.TenantID,
	}

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewS3Storage(prefix)
	require.NoError(t, err)

	r, err := repository.Initialize(ctx, acct, st)
	require.NoError(t, err)

	ro, err := r.NewRestore(ctx, []string{})
	require.NoError(t, err)
	require.NotNil(t, ro)
}
