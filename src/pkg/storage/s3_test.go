package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/storage"
)

type S3CfgSuite struct {
	suite.Suite
}

func TestS3CfgSuite(t *testing.T) {
	suite.Run(t, new(S3CfgSuite))
}

var goodS3Config = storage.S3Config{
	Bucket:   "bkt",
	Endpoint: "end",
	Prefix:   "pre",
}

func (suite *S3CfgSuite) TestS3Config_Config() {
	s3 := goodS3Config
	c, err := s3.StringConfig()
	assert.NoError(suite.T(), err)

	table := []struct {
		key    string
		expect string
	}{
		{"s3_bucket", s3.Bucket},
		{"s3_endpoint", s3.Endpoint},
		{"s3_prefix", s3.Prefix},
	}
	for _, test := range table {
		assert.Equal(suite.T(), test.expect, c[test.key])
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config() {
	t := suite.T()

	in := goodS3Config
	s, err := storage.NewStorage(storage.ProviderS3, in)
	assert.NoError(t, err)
	out, err := s.S3Config()
	assert.NoError(t, err)

	assert.Equal(t, in.Bucket, out.Bucket)
	assert.Equal(t, in.Endpoint, out.Endpoint)
	assert.Equal(t, in.Prefix, out.Prefix)
}

func makeTestS3Cfg(ak, bkt, end, pre, sk, tkn string) storage.S3Config {
	return storage.S3Config{
		Bucket:   bkt,
		Endpoint: end,
		Prefix:   pre,
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config_InvalidCases() {
	// missing required properties
	table := []struct {
		name string
		cfg  storage.S3Config
	}{
		{"missing bucket", makeTestS3Cfg("ak", "", "end", "pre", "sk", "tkn")},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := storage.NewStorage(storage.ProviderUnknown, test.cfg)
			assert.Error(t, err)
		})
	}

	// required property not populated in storage
	table2 := []struct {
		name  string
		amend func(storage.Storage)
	}{
		{
			"missing bucket",
			func(s storage.Storage) {
				s.Config["s3_bucket"] = ""
			},
		},
	}
	for _, test := range table2 {
		suite.T().Run(test.name, func(t *testing.T) {
			st, err := storage.NewStorage(storage.ProviderUnknown, goodS3Config)
			assert.NoError(t, err)
			test.amend(st)
			_, err = st.S3Config()
			assert.Error(t, err)
		})
	}
}
