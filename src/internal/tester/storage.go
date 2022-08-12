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
func NewPrefixedS3Storage(t *testing.T) storage.Storage {
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
			Corso: credentials.GetCorso(),
		},
	)
	require.NoError(t, err, "creating storage")

	return st
}
