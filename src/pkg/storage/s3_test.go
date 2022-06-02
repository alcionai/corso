package storage_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeebo/assert"

	"github.com/alcionai/corso/pkg/storage"
)

type S3CfgSuite struct {
	suite.Suite
}

func TestS3CfgSuite(t *testing.T) {
	suite.Run(t, new(S3CfgSuite))
}

func (suite *S3CfgSuite) TestS3Config_Config() {
	s3 := storage.S3Config{"ak", "bkt", "end", "pre", "sk", "tkn"}
	c, err := s3.Config()
	assert.NoError(suite.T(), err)

	table := []struct {
		key    string
		expect string
	}{
		{"s3_bucket", s3.Bucket},
		{"s3_accessKey", s3.AccessKey},
		{"s3_endpoint", s3.Endpoint},
		{"s3_prefix", s3.Prefix},
		{"s3_secretKey", s3.SecretKey},
		{"s3_sessionToken", s3.SessionToken},
	}
	for _, test := range table {
		assert.Equal(suite.T(), test.expect, c[test.key])
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config() {
	t := suite.T()

	in := storage.S3Config{"ak", "bkt", "end", "pre", "sk", "tkn"}
	s, err := storage.NewStorage(storage.ProviderS3, in)
	assert.NoError(t, err)
	out, err := s.S3Config()
	assert.NoError(t, err)

	assert.Equal(t, in.Bucket, out.Bucket)
	assert.Equal(t, in.AccessKey, out.AccessKey)
	assert.Equal(t, in.Endpoint, out.Endpoint)
	assert.Equal(t, in.Prefix, out.Prefix)
	assert.Equal(t, in.SecretKey, out.SecretKey)
	assert.Equal(t, in.SessionToken, out.SessionToken)
}
