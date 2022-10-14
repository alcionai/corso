package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/storage"
)

type S3CfgSuite struct {
	suite.Suite
}

func TestS3CfgSuite(t *testing.T) {
	suite.Run(t, new(S3CfgSuite))
}

var (
	goodS3Config = storage.S3Config{
		Bucket:   "bkt",
		Endpoint: "end",
		Prefix:   "pre/",
	}

	goodS3Map = map[string]string{
		"s3_bucket":   "bkt",
		"s3_endpoint": "end",
		"s3_prefix":   "pre/",
	}
)

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

func makeTestS3Cfg(bkt, end, pre string) storage.S3Config {
	return storage.S3Config{
		Bucket:   bkt,
		Endpoint: end,
		Prefix:   pre,
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config_invalidCases() {
	// missing required properties
	table := []struct {
		name string
		cfg  storage.S3Config
	}{
		{"missing bucket", makeTestS3Cfg("", "end", "pre/")},
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

func (suite *S3CfgSuite) TestStorage_S3Config_StringConfig() {
	table := []struct {
		name   string
		input  storage.S3Config
		expect map[string]string
	}{
		{
			name:   "standard",
			input:  goodS3Config,
			expect: goodS3Map,
		},
		{
			name:   "normalized bucket name",
			input:  makeTestS3Cfg("s3://"+goodS3Config.Bucket, goodS3Config.Endpoint, goodS3Config.Prefix),
			expect: goodS3Map,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := test.input.StringConfig()
			require.NoError(t, err)
			assert.Equal(t, test.expect, result)
		})
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config_Normalize() {
	const (
		prefixedBkt = "s3://bkt"
		normalBkt   = "bkt"
	)

	st := storage.S3Config{
		Bucket: prefixedBkt,
	}

	result := st.Normalize()
	assert.Equal(suite.T(), normalBkt, result.Bucket)
	assert.NotEqual(suite.T(), st.Bucket, result.Bucket)
}
