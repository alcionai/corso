package storage_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeebo/assert"

	"github.com/alcionai/corso/pkg/storage"
)

type CommonCfgSuite struct {
	suite.Suite
}

func TestCommonCfgSuite(t *testing.T) {
	suite.Run(t, new(CommonCfgSuite))
}

func (suite *CommonCfgSuite) TestCommonConfig_Config() {
	cfg := storage.CommonConfig{"passwd"}
	c, err := cfg.Config()
	assert.NoError(suite.T(), err)

	table := []struct {
		key    string
		expect string
	}{
		{"common_corsoPassword", cfg.CorsoPassword},
	}
	for _, test := range table {
		suite.T().Run(test.key, func(t *testing.T) {
			assert.Equal(t, test.expect, c[test.key])
		})
	}
}

func (suite *CommonCfgSuite) TestStorage_CommonConfig() {
	t := suite.T()

	in := storage.CommonConfig{"passwd"}
	s, err := storage.NewStorage(storage.ProviderUnknown, in)
	assert.NoError(t, err)
	out, err := s.CommonConfig()
	assert.NoError(t, err)

	assert.Equal(t, in.CorsoPassword, out.CorsoPassword)
}
