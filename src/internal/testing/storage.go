package testing

import (
	"os"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/storage"
)

// CheckS3EnvVars returns as error if any of the environment variables required for
// integration tests using S3 is empty. It does not check the validity of the
// variables with S3.
func CheckS3EnvVars() error {
	s3Envs := []string{
		credentials.AWS_ACCESS_KEY_ID,
		credentials.AWS_SECRET_ACCESS_KEY,
		credentials.AWS_SESSION_TOKEN,
	}
	for _, env := range s3Envs {
		if os.Getenv(env) == "" {
			return errors.Errorf("env var [%s] must be populated for integration testing", env)
		}
	}

	return nil
}

// NewS3Storage returns a storage.Storage object initialized with environment
// variables used for integration tests that use S3.
func NewS3Storage(prefix string) (storage.Storage, error) {
	return storage.NewStorage(
		storage.ProviderS3,
		storage.S3Config{
			AWS:    credentials.GetAWS(nil),
			Bucket: "test-corso-repo-init",
			Prefix: prefix,
		},
		storage.CommonConfig{
			Corso: credentials.GetCorso(),
		},
	)
}
