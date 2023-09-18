package storage

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/credentials"
)

type S3CfgSuite struct {
	suite.Suite
}

func TestS3CfgSuite(t *testing.T) {
	suite.Run(t, new(S3CfgSuite))
}

var (
	goodS3Config = S3Config{
		Bucket:         "bkt",
		Endpoint:       "end",
		Prefix:         "pre/",
		DoNotUseTLS:    false,
		DoNotVerifyTLS: false,
		AWS:            credentials.AWS{AccessKey: "access", SecretKey: "secret", SessionToken: "token"},
	}

	goodS3Map = map[string]string{
		keyS3Bucket:         "bkt",
		keyS3Endpoint:       "end",
		keyS3Prefix:         "pre/",
		keyS3DoNotUseTLS:    "false",
		keyS3DoNotVerifyTLS: "false",
		keyS3AccessKey:      "access",
		keyS3SecretKey:      "secret",
		keyS3SessionToken:   "token",
	}
)

func (suite *S3CfgSuite) TestS3Config_Config() {
	s3 := goodS3Config

	c, err := s3.StringConfig()
	assert.NoError(suite.T(), err, clues.ToCore(err))

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
	s, err := NewStorage(ProviderS3, &in)
	assert.NoError(t, err, clues.ToCore(err))
	sc, err := s.StorageConfig()
	assert.NoError(t, err, clues.ToCore(err))

	out := sc.(*S3Config)
	assert.Equal(t, in.Bucket, out.Bucket)
	assert.Equal(t, in.Endpoint, out.Endpoint)
	assert.Equal(t, in.Prefix, out.Prefix)
}

func makeTestS3Cfg(bkt, end, pre, access, secret, session string) S3Config {
	return S3Config{
		Bucket:   bkt,
		Endpoint: end,
		Prefix:   pre,
		AWS:      credentials.AWS{AccessKey: access, SecretKey: secret, SessionToken: session},
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config_invalidCases() {
	// missing required properties
	table := []struct {
		name string
		cfg  S3Config
	}{
		{"missing bucket", makeTestS3Cfg("", "end", "pre/", "", "", "")},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			_, err := NewStorage(ProviderUnknown, &test.cfg)
			assert.Error(suite.T(), err)
		})
	}

	// required property not populated in storage
	table2 := []struct {
		name  string
		amend func(Storage)
	}{
		{
			"missing bucket",
			func(s Storage) {
				s.Config["s3_bucket"] = ""
			},
		},
	}
	for _, test := range table2 {
		suite.Run(test.name, func() {
			t := suite.T()

			st, err := NewStorage(ProviderUnknown, &goodS3Config)
			assert.NoError(t, err, clues.ToCore(err))
			test.amend(st)
			_, err = st.StorageConfig()
			assert.Error(t, err)
		})
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config_StringConfig() {
	table := []struct {
		name   string
		input  S3Config
		expect map[string]string
	}{
		{
			name:   "standard",
			input:  goodS3Config,
			expect: goodS3Map,
		},
		{
			name: "normalized bucket name",
			input: makeTestS3Cfg(
				"s3://"+goodS3Config.Bucket,
				goodS3Config.Endpoint,
				goodS3Config.Prefix,
				goodS3Config.AccessKey,
				goodS3Config.SecretKey,
				goodS3Config.SessionToken),
			expect: goodS3Map,
		},
		{
			name: "disabletls",
			input: S3Config{
				Bucket:         "bkt",
				Endpoint:       "end",
				Prefix:         "pre/",
				DoNotUseTLS:    true,
				DoNotVerifyTLS: true,
			},
			expect: map[string]string{
				keyS3Bucket:         "bkt",
				keyS3Endpoint:       "end",
				keyS3Prefix:         "pre/",
				keyS3DoNotUseTLS:    "true",
				keyS3DoNotVerifyTLS: "true",
				keyS3AccessKey:      "",
				keyS3SecretKey:      "",
				keyS3SessionToken:   "",
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := test.input.StringConfig()
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, result)
		})
	}
}

func (suite *S3CfgSuite) TestStorage_S3Config_Normalize() {
	const (
		prefixedBkt = "s3://bkt"
		normalBkt   = "bkt"
	)

	st := S3Config{
		Bucket: prefixedBkt,
	}

	result := st.normalize()
	assert.Equal(suite.T(), normalBkt, result.Bucket)
	assert.NotEqual(suite.T(), st.Bucket, result.Bucket)
}
