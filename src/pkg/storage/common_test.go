package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

type CommonCfgSuite struct {
	suite.Suite
}

func TestCommonCfgSuite(t *testing.T) {
	suite.Run(t, new(CommonCfgSuite))
}

var goodCommonConfig = storage.CommonConfig{
	Corso: credentials.Corso{
		CorsoPassphrase: "passph",
	},
}

func (suite *CommonCfgSuite) TestCommonConfig_Config() {
	cfg := goodCommonConfig
	c, err := cfg.StringConfig()
	assert.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		key    string
		expect string
	}{
		{"common_corsoPassphrase", cfg.CorsoPassphrase},
	}
	for _, test := range table {
		suite.T().Run(test.key, func(t *testing.T) {
			assert.Equal(t, test.expect, c[test.key])
		})
	}
}

func (suite *CommonCfgSuite) TestStorage_CommonConfig() {
	t := suite.T()

	in := goodCommonConfig
	s, err := storage.NewStorage(storage.ProviderUnknown, in)
	assert.NoError(t, err, clues.ToCore(err))
	out, err := s.CommonConfig()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, in.CorsoPassphrase, out.CorsoPassphrase)
}

func (suite *CommonCfgSuite) TestStorage_CommonConfig_InvalidCases() {
	// missing required properties
	table := []struct {
		name string
		cfg  storage.CommonConfig
	}{
		{"missing passphrase", storage.CommonConfig{}},
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
			"missing passphrase",
			func(s storage.Storage) {
				s.Config["common_corsoPassphrase"] = ""
			},
		},
	}
	for _, test := range table2 {
		suite.T().Run(test.name, func(t *testing.T) {
			st, err := storage.NewStorage(storage.ProviderUnknown, goodCommonConfig)
			assert.NoError(t, err, clues.ToCore(err))
			test.amend(st)
			_, err = st.CommonConfig()
			assert.Error(t, err)
		})
	}
}
