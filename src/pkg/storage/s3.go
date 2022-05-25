package storage

type S3Config struct {
	AccessKey    string
	Bucket       string
	SecretKey    string
	SessionToken string
}

const (
	keyS3AccessKey    = "s3_accessKey"
	keyS3Bucket       = "s3_bucket"
	keyS3SecretKey    = "s3_secretKey"
	keyS3SessionToken = "s3_sessionToken"
)

func (c S3Config) Config() config {
	return config{
		keyS3AccessKey:    c.AccessKey,
		keyS3Bucket:       c.Bucket,
		keyS3SecretKey:    c.SecretKey,
		keyS3SessionToken: c.SessionToken,
	}
}

// S3Config retrieves the S3Config details from the Storage config.
func (s Storage) S3Config() S3Config {
	c := S3Config{}
	if len(s.Config) > 0 {
		c.AccessKey = s.Config[keyS3AccessKey].(string)
		c.Bucket = s.Config[keyS3Bucket].(string)
		c.SecretKey = s.Config[keyS3SecretKey].(string)
		c.SessionToken = s.Config[keyS3SessionToken].(string)
	}
	return c
}
