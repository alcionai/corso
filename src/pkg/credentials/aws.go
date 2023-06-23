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
	return AWS{
		AccessKey:    override[AWSAccessKeyID],
		SecretKey:    override[AWSSecretAccessKey],
		SessionToken: override[AWSSessionToken],
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
