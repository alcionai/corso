package kopia

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/storage"
)

func getTestStorage(bucket, prefix string) storage.Storage {
	return storage.NewStorage(
		storage.ProviderS3,
		storage.S3Config{
			AccessKey:    os.Getenv(storage.AWS_ACCESS_KEY_ID),
			Bucket:       bucket,
			Prefix:       prefix,
			SecretKey:    os.Getenv(storage.AWS_SECRET_ACCESS_KEY),
			SessionToken: os.Getenv(storage.AWS_SESSION_TOKEN),
		},
	)
}

// ---------------
// integration tests that use kopia
// ---------------
type KopiaIntegrationSuite struct {
	suite.Suite
}

func TestKopiaIntegrationSuite(t *testing.T) {
	runIntegrationTests := os.Getenv("INTEGRATION_TESTING")
	if runIntegrationTests != "true" {
		t.Skip()
	}

	suite.Run(t, new(KopiaIntegrationSuite))
}

func (suite *KopiaIntegrationSuite) SetupSuite() {
	s3Envs := []string{
		storage.AWS_ACCESS_KEY_ID,
		storage.AWS_SECRET_ACCESS_KEY,
		storage.AWS_SESSION_TOKEN,
	}
	for _, env := range s3Envs {
		require.NotZerof(
			suite.T(),
			os.Getenv(env),
			"env var [%s] must be populated for integration testing",
			env,
		)
	}
}

func (suite *KopiaIntegrationSuite) TestInitializeAndOpenRepo() {
	ctx := context.Background()
	timeOfTest := time.Now().UTC().Format("2016-01-02T15:04:05")
	suite.T().Logf("TestInitialize() run at %s", timeOfTest)

	storage := getTestStorage("test-corso-kopia-wrapper", "init-and-open-repo-s3"+timeOfTest)
	k := New(storage)
	err := k.Initialize(ctx)
	require.NoError(suite.T(), err)

	err = k.open(ctx)
	assert.NoError(suite.T(), err)
	k.Close(ctx)
}
