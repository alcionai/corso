package credentials

import (
	"os"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
)

// envvar consts
const (
	AWSAccessKeyID     = "AWS_ACCESS_KEY_ID"
	AWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	AWSSessionToken    = "AWS_SESSION_TOKEN"
)

// AWS aggregates aws credentials from flag and env_var values.
type AWS struct {
	AccessKey    string // required
	SecretKey    string // required
	SessionToken string // required
}

// GetAWS is a helper for aggregating aws secrets and credentials.
func GetAWS(override map[string]string) AWS {
	accessKey := os.Getenv(AWSAccessKeyID)
	if ovr, ok := override[AWSAccessKeyID]; ok && ovr != "" {
		accessKey = ovr
	}

	secretKey := os.Getenv(AWSSecretAccessKey)
	sessToken := os.Getenv(AWSSessionToken)

	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	return AWS{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		SessionToken: sessToken,
	}
}

func (c AWS) Validate() error {
	check := map[string]string{
		AWSAccessKeyID:     c.AccessKey,
		AWSSecretAccessKey: c.SecretKey,
		AWSSessionToken:    c.SessionToken,
	}

	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, errors.New(k))
		}
	}

	return nil
}
