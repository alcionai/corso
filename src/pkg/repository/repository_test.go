package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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
		storage  storage.Storage
		account  repository.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			storage.NewStorage(storage.ProviderUnknown),
			repository.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := repository.Initialize(context.Background(), test.account, test.storage)
			test.errCheck(suite.T(), err, "")
		})
	}
}

// repository.Connect involves end-to-end communication with kopia, therefore this only
// tests expected error cases
func (suite *RepositorySuite) TestConnect() {
	table := []struct {
		name     string
		storage  storage.Storage
		account  repository.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			storage.NewStorage(storage.ProviderUnknown),
			repository.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := repository.Connect(context.Background(), test.account, test.storage)
			test.errCheck(suite.T(), err)
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
	runIntegrationTests := os.Getenv("INTEGRATION_TESTING")
	if runIntegrationTests != "true" {
		t.Skip()
	}
	suite.Run(t, new(RepositoryIntegrationSuite))
}

// ensure all required env values are populated
func (suite *RepositoryIntegrationSuite) SetupSuite() {
	s3Envs := []string{
		storage.AWS_ACCESS_KEY_ID,
		storage.AWS_SECRET_ACCESS_KEY,
		storage.AWS_SESSION_TOKEN,
	}
	for _, env := range s3Envs {
		require.NotZerof(suite.T(), os.Getenv(env), "env var [%s] must be populated for integration testing", env)
	}
}

func (suite *RepositoryIntegrationSuite) TestInitialize() {
	ctx := context.Background()
	timeOfTest := time.Now().UTC().Format("2016-01-02T15:04:05")
	suite.T().Logf("TestInitialize() run at: %s", timeOfTest)

	table := []struct {
		prefix   string
		account  repository.Account
		storage  storage.Storage
		errCheck assert.ErrorAssertionFunc
	}{
		{
			prefix: "init-s3-" + timeOfTest,
			storage: storage.NewStorage(
				storage.ProviderS3,
				storage.S3Config{
					AccessKey:    os.Getenv(storage.AWS_ACCESS_KEY_ID),
					Bucket:       "test-corso-repo-init",
					Prefix:       "init-s3-" + timeOfTest,
					SecretKey:    os.Getenv(storage.AWS_SECRET_ACCESS_KEY),
					SessionToken: os.Getenv(storage.AWS_SESSION_TOKEN),
				},
			),
			errCheck: assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.prefix, func(t *testing.T) {
			_, err := repository.Initialize(ctx, test.account, test.storage)
			test.errCheck(suite.T(), err)
		})
	}
}
