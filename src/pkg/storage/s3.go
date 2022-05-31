package storage

type S3Config struct {
	AccessKey    string
	Bucket       string
	Endpoint     string
	Prefix       string
	SecretKey    string
	SessionToken string
}

// envvar consts
const (
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_SESSION_TOKEN     = "AWS_SESSION_TOKEN"
)

// config key consts
const (
	keyS3AccessKey    = "s3_accessKey"
	keyS3Bucket       = "s3_bucket"
	keyS3Endpoint     = "s3_endpoint"
	keyS3Prefix       = "s3_prefix"
	keyS3SecretKey    = "s3_secretKey"
	keyS3SessionToken = "s3_sessionToken"
)

func (c S3Config) Config() config {
	return config{
		keyS3AccessKey:    c.AccessKey,
		keyS3Bucket:       c.Bucket,
		keyS3Endpoint:     c.Endpoint,
		keyS3Prefix:       c.Prefix,
		keyS3SecretKey:    c.SecretKey,
		keyS3SessionToken: c.SessionToken,
	}
}

// S3Config retrieves the S3Config details from the Storage config.
func (s Storage) S3Config() S3Config {
	c := S3Config{}
	if len(s.Config) > 0 {
		c.AccessKey = orEmptyString(s.Config[keyS3AccessKey])
		c.Bucket = orEmptyString(s.Config[keyS3Bucket])
		c.Endpoint = orEmptyString(s.Config[keyS3Endpoint])
		c.Prefix = orEmptyString(s.Config[keyS3Prefix])
		c.SecretKey = orEmptyString(s.Config[keyS3SecretKey])
		c.SessionToken = orEmptyString(s.Config[keyS3SessionToken])
	}
	return c
}
