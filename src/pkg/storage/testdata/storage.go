package testdata

import (
	"fmt"
	"os"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
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
	now := tester.LogTimeOfTest(t)

	cfg, err := tconfig.ReadTestConfig()
	require.NoError(t, err, "configuring storage from test file", clues.ToCore(err))

	prefix := testRepoRootPrefix + t.Name() + "-" + now
	t.Logf("testing at s3 bucket [%s] prefix [%s]", cfg[tconfig.TestCfgBucket], prefix)

	st, err := storage.NewStorage(
		storage.ProviderS3,
		storage.S3Config{
			Bucket: cfg[tconfig.TestCfgBucket],
			Prefix: prefix,
		},
		storage.CommonConfig{
			Corso:       GetAndInsertCorso(""),
			KopiaCfgDir: t.TempDir(),
		},
	)
	require.NoError(t, err, "creating storage", clues.ToCore(err))

	return st
}

// GetCorso is a helper for aggregating Corso secrets and credentials.
func GetAndInsertCorso(passphase string) credentials.Corso {
	fmt.Println(
		"flags.CorsoPassphraseFV: ",
		flags.CorsoPassphraseFV,
		"Length of flag value: ",
		len(flags.CorsoPassphraseFV))
	// fetch data from flag, env var or func param giving priority to func param
	// Func param generally will be value fetched from config file using viper.
	corsoPassph := str.First(flags.CorsoPassphraseFV, os.Getenv(credentials.CorsoPassphrase), passphase)

	return credentials.Corso{
		CorsoPassphrase: corsoPassph,
	}
}
