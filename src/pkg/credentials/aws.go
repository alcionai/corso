package credentials

import (
	"os"
)

// envvar consts
const (
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_SESSION_TOKEN     = "AWS_SESSION_TOKEN"
)

// AWS aggregates aws credentials from flag and env_var values.
type AWS struct {
	AccessKey    string // required
	SecretKey    string // required
	SessionToken string // required
}

// GetAWS is a helper for aggregating aws secrets and credentials.
func GetAWS(override map[string]string) AWS {
	accessKey := os.Getenv(AWS_ACCESS_KEY_ID)
	if ovr, ok := override[AWS_ACCESS_KEY_ID]; ok {
		accessKey = ovr
	}
	secretKey := os.Getenv(AWS_SECRET_ACCESS_KEY)
	sessToken := os.Getenv(AWS_SESSION_TOKEN)

	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	return AWS{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		SessionToken: sessToken,
	}
}
