package storage

type S3Config struct {
	Bucket    string
	AccessKey string
	SecretKey string
}

func (c S3Config) Config() config {
	return config{
		"s3_bucket":    c.Bucket,
		"s3_accessKey": c.AccessKey,
		"s3_secretKey": c.SecretKey,
	}
}

// S3Config retrieves the S3Config details from the Storage config.
func (s Storage) S3Config() S3Config {
	c := S3Config{}
	if len(s.Config) > 0 {
		c.Bucket = s.Config["s3_bucket"].(string)
		c.AccessKey = s.Config["s3_accessKey"].(string)
		c.SecretKey = s.Config["s3_secretKey"].(string)
	}
	return c
}
