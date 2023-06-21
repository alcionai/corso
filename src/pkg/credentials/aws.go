package credentials

import (
	"github.com/alcionai/clues"
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
	var accessKey, secretKey, sessToken string
	if ovr := override[AWSAccessKeyID]; ovr != "" {
		accessKey = ovr
	}

	if ovr := override[AWSSecretAccessKey]; ovr != "" {
		secretKey = ovr
	}

	if ovr := override[AWSSessionToken]; ovr != "" {
		sessToken = ovr
	}

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
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}
