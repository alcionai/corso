package storage

type S3Config struct {
	Bucket    string
	AccessKey string
	SecretKey string
}

const (
	keyS3Bucket    = "s3_bucket"
	keyS3AccessKey = "s3_accessKey"
	keyS3SecretKey = "s3_secretKey"
)

func (c S3Config) Config() config {
	return config{
		keyS3Bucket:    c.Bucket,
		keyS3AccessKey: c.AccessKey,
		keyS3SecretKey: c.SecretKey,
	}
}

// S3Config retrieves the S3Config details from the Storage config.
func (s Storage) S3Config() S3Config {
	c := S3Config{}
	if len(s.Config) > 0 {
		c.Bucket = s.Config[keyS3Bucket].(string)
		c.AccessKey = s.Config[keyS3AccessKey].(string)
		c.SecretKey = s.Config[keyS3SecretKey].(string)
	}
	return c
}
