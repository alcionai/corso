package account_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
)

type M365CfgSuite struct {
	suite.Suite
}

func TestM365CfgSuite(t *testing.T) {
	suite.Run(t, new(M365CfgSuite))
}

var goodM365Config = account.M365Config{
	M365: credentials.M365{
		AzureClientID:     "cid",
		AzureClientSecret: "cs",
	},
	AzureTenantID: "tid",
}

func (suite *M365CfgSuite) TestM365Config_Config() {
	m365 := goodM365Config
	c, err := m365.StringConfig()
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		key    string
		expect string
	}{
		{"azure_clientid", m365.AzureClientID},
		{"azure_clientSecret", m365.AzureClientSecret},
		{"azure_tenantid", m365.AzureTenantID},
	}
	for _, test := range table {
		assert.Equal(suite.T(), test.expect, c[test.key])
	}
}

func (suite *M365CfgSuite) TestAccount_M365Config() {
	t := suite.T()

	in := goodM365Config
	a, err := account.NewAccount(account.ProviderM365, in)
	require.NoError(t, err, clues.ToCore(err))

	out, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, in.AzureClientID, out.AzureClientID)
	assert.Equal(t, in.AzureClientSecret, out.AzureClientSecret)
	assert.Equal(t, in.AzureTenantID, out.AzureTenantID)
}

func makeTestM365Cfg(cid, cs, tid string) account.M365Config {
	return account.M365Config{
		M365: credentials.M365{
			AzureClientID:     cid,
			AzureClientSecret: cs,
		},
		AzureTenantID: tid,
	}
}

func (suite *M365CfgSuite) TestAccount_M365Config_InvalidCases() {
	// missing required properties
	table := []struct {
		name string
		cfg  account.M365Config
	}{
		{"missing client ID", makeTestM365Cfg("", "cs", "tid")},
		{"missing client secret", makeTestM365Cfg("cid", "", "tid")},
		{"missing tenant ID", makeTestM365Cfg("cid", "cs", "")},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := account.NewAccount(account.ProviderUnknown, test.cfg)
			assert.Error(t, err)
		})
	}

	// required property not populated in account
	table2 := []struct {
		name  string
		amend func(account.Account)
	}{
		{
			"missing clientID",
			func(a account.Account) {
				a.Config["azure_clientid"] = ""
			},
		},
		{
			"missing client secret",
			func(a account.Account) {
				a.Config["azure_clientSecret"] = ""
			},
		},
		{
			"missing tenant id",
			func(a account.Account) {
				a.Config["azure_tenantid"] = ""
			},
		},
	}
	for _, test := range table2 {
		suite.T().Run(test.name, func(t *testing.T) {
			st, err := account.NewAccount(account.ProviderUnknown, goodM365Config)
			assert.NoError(t, err, clues.ToCore(err))
			test.amend(st)

			_, err = st.M365Config()
			assert.Error(t, err)
		})
	}
}
