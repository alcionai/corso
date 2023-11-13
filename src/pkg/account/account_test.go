package account

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testConfig struct {
	expect string
	id     string
	err    error
}

func (c testConfig) providerID(ap accountProvider) string {
	return c.id
}

func (c testConfig) StringConfig() (map[string]string, error) {
	return map[string]string{"expect": c.expect}, c.err
}

func (c testConfig) configHash() (string, error) {
	return "hashed-config", c.err
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
		{"unknown no error", ProviderUnknown, testConfig{"configVal", "", nil}, assert.NoError},
		{"m365 no error", ProviderM365, testConfig{"configVal", "", nil}, assert.NoError},
		{"unknown w/ error", ProviderUnknown, testConfig{"configVal", "", assert.AnError}, assert.Error},
		{"m365 w/ error", ProviderM365, testConfig{"configVal", "", assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			s, err := NewAccount(test.p, test.c)
			test.errCheck(t, err, clues.ToCore(err))
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

func (suite *AccountSuite) TestGetAccountConfigHash() {
	tests := []struct {
		name     string
		provider accountProvider
		config   any
	}{
		{
			name:     "valid account",
			provider: ProviderM365,
			config:   getTestM365Config("1234", "5678"),
		},
		{
			name:     "invalid account",
			provider: ProviderUnknown,
			config:   testConfig{"configVal", "", nil},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			if test.provider == ProviderUnknown {
				s, err := NewAccount(test.provider, test.config.(testConfig))
				require.NoError(t, err)

				_, err = s.GetAccountConfigHash()
				require.Error(t, err)
			}

			if test.provider == ProviderM365 {
				_, ok := test.config.(providerIDer)
				require.True(t, ok)

				s, err := NewAccount(test.provider, test.config.(M365Config))
				require.NoError(t, err)

				hash, err := s.GetAccountConfigHash()
				require.NoError(t, err)
				assert.True(t, len(hash) > 0)
			}
		})
	}
}

func getTestM365Config(clientID, tenantID string) M365Config {
	c := M365Config{}
	c.AzureClientID = clientID
	c.AzureClientSecret = "super secret"
	c.AzureTenantID = tenantID

	return c
}
