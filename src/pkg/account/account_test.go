package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testConfig struct {
	expect string
	err    error
}

func (c testConfig) Config() (config, error) {
	return config{"expect": c.expect}, c.err
}

type AccountSuite struct {
	suite.Suite
}

func TestAccountSuite(t *testing.T) {
	suite.Run(t, new(AccountSuite))
}

func (suite *AccountSuite) TestNewAccount() {
	table := []struct {
		name     string
		p        accountProvider
		c        testConfig
		errCheck assert.ErrorAssertionFunc
	}{
		{"unknown no error", ProviderUnknown, testConfig{"configVal", nil}, assert.NoError},
		{"m365 no error", ProviderM365, testConfig{"configVal", nil}, assert.NoError},
		{"unknown w/ error", ProviderUnknown, testConfig{"configVal", assert.AnError}, assert.Error},
		{"m365 w/ error", ProviderM365, testConfig{"configVal", assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			s, err := NewAccount(test.p, test.c)
			test.errCheck(t, err)
			// remaining tests are dependent upon error-free state
			if test.c.err != nil {
				return
			}
			assert.Equalf(t,
				test.p,
				s.Provider,
				"expected account provider [%s], got [%s]", test.p, s.Provider)
			assert.Equalf(t,
				test.c.expect,
				s.Config["expect"],
				"expected account config [%s], got [%s]", test.c.expect, s.Config["expect"])
		})
	}
}

type fooConfig struct {
	foo string
	err error
}

func (c fooConfig) Config() (config, error) {
	return config{"foo": c.foo}, c.err
}

func (suite *AccountSuite) TestUnionConfigs() {
	table := []struct {
		name     string
		tc       testConfig
		fc       fooConfig
		errCheck assert.ErrorAssertionFunc
	}{
		{"no error", testConfig{"test", nil}, fooConfig{"foo", nil}, assert.NoError},
		{"tc error", testConfig{"test", assert.AnError}, fooConfig{"foo", nil}, assert.Error},
		{"fc error", testConfig{"test", nil}, fooConfig{"foo", assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cs, err := unionConfigs(test.tc, test.fc)
			test.errCheck(t, err)
			// remaining tests depend on error-free state
			if test.tc.err != nil || test.fc.err != nil {
				return
			}
			assert.Equalf(t,
				test.tc.expect,
				cs["expect"],
				"expected unioned config to have value [%s] at key [expect], got [%s]", test.tc.expect, cs["expect"])
			assert.Equalf(t,
				test.fc.foo,
				cs["foo"],
				"expected unioned config to have value [%s] at key [foo], got [%s]", test.fc.foo, cs["foo"])
		})
	}
}
