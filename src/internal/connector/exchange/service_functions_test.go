package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type ServiceFunctionsIntegrationSuite struct {
	suite.Suite
	m365UserID string
	creds      account.M365Config
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
	suite.m365UserID = tester.M365UserID(suite.T())
	a := tester.NewM365Account(t)
	require.NoError(t, err)
	m365, err := a.M365Config()
	require.NoError(t, err)
	suite.creds = m365

}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllCalendars() {
	ctx, flush := tester.NewContext()
	defer flush()

	gs := loadService(suite.T())

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
			contains:    DefaultCalendar,
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
			cals, err := GetAllCalendars(ctx, gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllContactFolders() {
	ctx, flush := tester.NewContext()
	defer flush()

	gs := loadService(suite.T())
	user := tester.M365UserID(suite.T())

	table := []struct {
		name, contains, user string
		expectCount          assert.ComparisonAssertionFunc
		expectErr            assert.ErrorAssertionFunc
	}{
		{
			name:        "plain lookup",
			user:        user,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
		},
		{
			name:        "root folder",
			contains:    "Contact", // DefaultContactFolder doesn't work here?
			user:        user,
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
			user:        user,
			expectCount: assert.Equal,
			expectErr:   assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeBackup()
			sel.Includes()
			params := graph.QueryParams{
				User:        test.user,
				Scope:       nil,
				FailFast:    false,
				Credentials: suite.creds,
			}
			cals, err := GetAllContactFolders(ctx, params, gs, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllMailFolders() {
	ctx, flush := tester.NewContext()
	defer flush()

	gs := loadService(suite.T())

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
			contains:    DefaultMailFolder,
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
			cals, err := GetAllMailFolders(ctx, gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}
