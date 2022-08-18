package tester

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/storage"
)

var AWSStorageCredEnvs = []string{
	credentials.AWSAccessKeyID,
	credentials.AWSSecretAccessKey,
	credentials.AWSSessionToken,
}

// NewPrefixedS3Storage returns a storage.Storage object initialized with environment
// variables used for integration tests that use S3. The prefix for the storage
// path will be unique.
// If kopiaCfgDir is populated, kopia will use that value for its config storage
// and caching directory location.  Integration test should set this value to
// a t.TempDir() to avoid caching collision with other integration tests.
func NewPrefixedS3Storage(t *testing.T, kopiaCfgDir string) storage.Storage {
	now := LogTimeOfTest(t)
	cfg, err := readTestConfig()
	require.NoError(t, err, "configuring storage from test file")

	st, err := storage.NewStorage(
		storage.ProviderS3,
		storage.S3Config{
			AWS:    credentials.GetAWS(nil),
			Bucket: cfg[TestCfgBucket],
			Prefix: t.Name() + "-" + now,
		},
		storage.CommonConfig{
			Corso:       credentials.GetCorso(),
			KopiaCfgDir: kopiaCfgDir,
		},
	)
	require.NoError(t, err, "creating storage")
	return st
}
