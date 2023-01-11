package tester

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

const testRepoRootPrefix = "corso_integration_test/"

var AWSStorageCredEnvs = []string{
	credentials.AWSAccessKeyID,
	credentials.AWSSecretAccessKey,
	credentials.AWSSessionToken,
}

// NewPrefixedS3Storage returns a storage.Storage object initialized with environment
// variables used for integration tests that use S3. The prefix for the storage
// path will be unique.
// Uses t.TempDir() to generate a unique config storage and caching directory for this
// test.  Suites that need to identify this value can retrieve it again from the common
// configs.
func NewPrefixedS3Storage(t *testing.T) storage.Storage {
	now := LogTimeOfTest(t)

	cfg, err := readTestConfig()
	require.NoError(t, err, "configuring storage from test file")

	st, err := storage.NewStorage(
		storage.ProviderS3,
		storage.S3Config{
			Bucket: cfg[TestCfgBucket],
			Prefix: testRepoRootPrefix + t.Name() + "-" + now,
		},
		storage.CommonConfig{
			Corso:       credentials.GetCorso(),
			KopiaCfgDir: t.TempDir(),
		},
	)
	require.NoError(t, err, "creating storage")

	return st
}
