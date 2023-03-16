package storage

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testConfig struct {
	expect string
	err    error
}

func (c testConfig) StringConfig() (map[string]string, error) {
	return map[string]string{"expect": c.expect}, c.err
}

type StorageSuite struct {
	suite.Suite
}

func TestStorageSuite(t *testing.T) {
	suite.Run(t, new(StorageSuite))
}

func (suite *StorageSuite) TestNewStorage() {
	table := []struct {
		name     string
		p        storageProvider
		c        testConfig
		errCheck assert.ErrorAssertionFunc
	}{
		{"unknown no error", ProviderUnknown, testConfig{"configVal", nil}, assert.NoError},
		{"s3 no error", ProviderS3, testConfig{"configVal", nil}, assert.NoError},
		{"unknown w/ error", ProviderUnknown, testConfig{"configVal", assert.AnError}, assert.Error},
		{"s3 w/ error", ProviderS3, testConfig{"configVal", assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			s, err := NewStorage(test.p, test.c)
			test.errCheck(t, err, clues.ToCore(err))

			// remaining tests are dependent upon error-free state
			if test.c.err != nil {
				return
			}

			assert.Equalf(t,
				test.p,
				s.Provider,
				"expected storage provider [%s], got [%s]", test.p, s.Provider)
			assert.Equalf(t,
				test.c.expect,
				s.Config["expect"],
				"expected storage config [%s], got [%s]", test.c.expect, s.Config["expect"])
		})
	}
}
