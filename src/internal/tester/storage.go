package tester

import (
	"testing"

	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/storage"
	"github.com/pkg/errors"
)

var AWSStorageCredEnvs = []string{
	credentials.AWSAccessKeyID,
	credentials.AWSSecretAccessKey,
	credentials.AWSSessionToken,
}

// NewPrefixedS3Storage returns a storage.Storage object initialized with environment
// variables used for integration tests that use S3. The prefix for the storage
// path will be unique.
func NewPrefixedS3Storage(t *testing.T) (storage.Storage, error) {
	// now := LogTimeOfTest(t)
	cfg, err := readTestConfig()
	if err != nil {
		return storage.Storage{}, errors.Wrap(err, "configuring storage from test file")
	}

	return storage.NewStorage(
		storage.ProviderS3,
		storage.S3Config{
			AWS:    credentials.GetAWS(nil),
			Bucket: cfg[testCfgBucket],
			// Prefix: t.Name() + "-" + now,
		},
		storage.CommonConfig{
			Corso: credentials.GetCorso(),
		},
	)
}
