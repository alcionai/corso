package exchange_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/tester"
)

type ServiceFunctionsIntegrationSuite struct {
	suite.Suite
	gc         *connector.GraphConnector
	m365UserID string
}

func TestServiceFunctionsIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ServiceFunctionsIntegrationSuite))
}

func (suite *ServiceFunctionsIntegrationSuite) SetupSuite() {
	t := suite.T()

	_, err := tester.GetRequiredEnvSls(tester.AWSStorageCredEnvs)
	require.NoError(t, err)

	acct := tester.NewM365Account(t)
	gc, err := connector.NewGraphConnector(acct)
	require.NoError(t, err)

	suite.gc = gc
	suite.m365UserID = tester.M365UserID(t)
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllCalendars() {
	gs := suite.gc.Service()

	table := []struct {
		name, contains, user string
		expectCount          assert.ComparisonAssertionFunc
		expectErr            assert.ErrorAssertionFunc
	}{
		{
			name:        "plain lookup",
			user:        suite.m365UserID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
		},
		{
			name:        "root calendar",
			contains:    "Calendar",
			user:        suite.m365UserID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
		},
		{
			name:        "nonsense user",
			user:        "fnords_mc_snarfens",
			expectCount: assert.Equal,
			expectErr:   assert.Error,
		},
		{
			name:        "nonsense matcher",
			contains:    "∂ç∂ç∂√≈∂ƒß∂ç√ßç√≈ç√ß∂ƒçß√ß≈∂ƒßç√",
			user:        suite.m365UserID,
			expectCount: assert.Equal,
			expectErr:   assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cals, err := exchange.GetAllCalendars(gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllContactFolders() {
	gs := suite.gc.Service()

	table := []struct {
		name, contains, user string
		expectCount          assert.ComparisonAssertionFunc
		expectErr            assert.ErrorAssertionFunc
	}{
		{
			name:        "plain lookup",
			user:        suite.m365UserID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
		},
		{
			name:        "root folder",
			contains:    "Contact",
			user:        suite.m365UserID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
		},
		{
			name:        "nonsense user",
			user:        "fnords_mc_snarfens",
			expectCount: assert.Equal,
			expectErr:   assert.Error,
		},
		{
			name:        "nonsense matcher",
			contains:    "∂ç∂ç∂√≈∂ƒß∂ç√ßç√≈ç√ß∂ƒçß√ß≈∂ƒßç√",
			user:        suite.m365UserID,
			expectCount: assert.Equal,
			expectErr:   assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cals, err := exchange.GetAllContactFolders(gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllMailFolders() {
	gs := suite.gc.Service()

	table := []struct {
		name, contains, user string
		expectCount          assert.ComparisonAssertionFunc
		expectErr            assert.ErrorAssertionFunc
	}{
		{
			name:        "plain lookup",
			user:        suite.m365UserID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
		},
		{
			name:        "Root folder",
			contains:    "Inbox",
			user:        suite.m365UserID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
		},
		{
			name:        "nonsense user",
			user:        "fnords_mc_snarfens",
			expectCount: assert.Equal,
			expectErr:   assert.Error,
		},
		{
			name:        "nonsense matcher",
			contains:    "∂ç∂ç∂√≈∂ƒß∂ç√ßç√≈ç√ß∂ƒçß√ß≈∂ƒßç√",
			user:        suite.m365UserID,
			expectCount: assert.Equal,
			expectErr:   assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cals, err := exchange.GetAllMailFolders(gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}
