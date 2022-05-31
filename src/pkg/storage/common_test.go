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
	c := cfg.Config()
	table := []struct {
		key    string
		expect string
	}{
		{"common_corsoPassword", cfg.CorsoPassword},
	}
	for _, test := range table {
		suite.T().Run(test.key, func(t *testing.T) {
			assert.Equal(t, c[test.key], test.expect)
		})
	}
}

func (suite *CommonCfgSuite) TestStorage_CommonConfig() {
	in := storage.CommonConfig{"passwd"}
	out := storage.NewStorage(storage.ProviderUnknown, in).CommonConfig()
	t := suite.T()
	assert.Equal(t, in.CorsoPassword, out.CorsoPassword)
}
