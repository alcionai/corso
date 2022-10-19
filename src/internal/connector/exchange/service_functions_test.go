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

const (
	invalidUser       = "fnords_mc_snarfens"
	nonExistantLookup = "∂ç∂ç∂√≈∂ƒß∂ç√ßç√≈ç√ß∂ƒçß√ß≈∂ƒßç√"
)

func TestServiceFunctionsIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ServiceFunctionsIntegrationSuite))
}

func (suite *ServiceFunctionsIntegrationSuite) SetupSuite() {
	t := suite.T()
	suite.m365UserID = tester.M365UserID(t)
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllCalendars() {
	ctx, flush := tester.NewContext()
	defer flush()

	gs := loadService(suite.T())
	userID := tester.M365UserID(suite.T())

	table := []struct {
		name, user  string
		expectCount assert.ComparisonAssertionFunc
		getScope    func(t *testing.T) selectors.ExchangeScope
		expectErr   assert.ErrorAssertionFunc
	}{
		{
			name:        "plain lookup",
			user:        userID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.NewExchangeBackup().EventCalendars([]string{userID}, selectors.Any())[0]
			},
		},
		{
			name:        "Get Default Calendar",
			user:        userID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.NewExchangeBackup().EventCalendars([]string{userID}, []string{DefaultCalendar})[0]
			},
		},
		{
			name:        "non-root calendar",
			user:        userID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.NewExchangeBackup().EventCalendars([]string{userID}, []string{"Birthdays"})[0]
			},
		},
		{
			name:        "nonsense user",
			user:        invalidUser,
			expectCount: assert.Equal,
			expectErr:   assert.Error,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					EventCalendars([]string{invalidUser}, []string{DefaultContactFolder})[0]
			},
		},
		{
			name:        "nonsense matcher",
			user:        userID,
			expectCount: assert.Equal,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					EventCalendars([]string{userID}, []string{nonExistantLookup})[0]
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			params := graph.QueryParams{
				User:        test.user,
				Scope:       test.getScope(t),
				FailFast:    false,
				Credentials: suite.creds,
			}
			cals, err := GetAllCalendars(ctx, params, gs)
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
		name, user  string
		expectCount assert.ComparisonAssertionFunc
		getScope    func(t *testing.T) selectors.ExchangeScope
		expectErr   assert.ErrorAssertionFunc
	}{
		{
			name:        "plain lookup",
			user:        user,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					ContactFolders([]string{user}, selectors.Any())[0]
			},
		},
		{
			name:        "default contact folder",
			user:        user,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					ContactFolders([]string{user}, []string{DefaultContactFolder})[0]
			},
		},
		{
			name:        "Trial folder lookup",
			user:        user,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					ContactFolders([]string{user}, []string{"TrialFolder"})[0]
			},
		},
		{
			name:        "nonsense user",
			user:        invalidUser,
			expectCount: assert.Equal,
			expectErr:   assert.Error,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					ContactFolders([]string{invalidUser}, []string{DefaultContactFolder})[0]
			},
		},
		{
			name:        "nonsense matcher",
			user:        user,
			expectCount: assert.Equal,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					ContactFolders([]string{user}, []string{nonExistantLookup})[0]
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			params := graph.QueryParams{
				User:        test.user,
				Scope:       test.getScope(t),
				FailFast:    false,
				Credentials: suite.creds,
			}
			cals, err := GetAllContactFolders(ctx, params, gs)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllMailFolders() {
	ctx, flush := tester.NewContext()
	defer flush()

	gs := loadService(suite.T())
	userID := tester.M365UserID(suite.T())

	table := []struct {
		name, user  string
		expectCount assert.ComparisonAssertionFunc
		getScope    func(t *testing.T) selectors.ExchangeScope
		expectErr   assert.ErrorAssertionFunc
	}{
		{
			name:        "plain lookup",
			user:        userID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					MailFolders([]string{userID}, selectors.Any())[0]
			},
		},
		{
			name:        "root folder",
			user:        userID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					MailFolders([]string{userID}, []string{DefaultMailFolder})[0]
			},
		},
		{
			name:        "Trial folder lookup",
			user:        userID,
			expectCount: assert.Greater,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					MailFolders([]string{userID}, []string{"Drafts"})[0]
			},
		},
		{
			name:        "nonsense user",
			user:        invalidUser,
			expectCount: assert.Equal,
			expectErr:   assert.Error,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					MailFolders([]string{invalidUser}, []string{DefaultMailFolder})[0]
			},
		},
		{
			name:        "nonsense matcher",
			user:        userID,
			expectCount: assert.Equal,
			expectErr:   assert.NoError,
			getScope: func(t *testing.T) selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					MailFolders([]string{userID}, []string{nonExistantLookup})[0]
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			params := graph.QueryParams{
				User:        test.user,
				Scope:       test.getScope(t),
				FailFast:    false,
				Credentials: suite.creds,
			}
			cals, err := GetAllMailFolders(ctx, params, gs)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestCollectContainers() {
	ctx, flush := tester.NewContext()
	defer flush()

	failFast := false
	containerCount := 1
	t := suite.T()
	user := tester.M365UserID(t)
	a := tester.NewM365Account(t)
	service := loadService(t)
	credentials, err := a.M365Config()
	require.NoError(t, err)

	tests := []struct {
		name, contains string
		getScope       func() selectors.ExchangeScope
		expectedCount  assert.ComparisonAssertionFunc
	}{
		{
			name:          "All Events",
			contains:      "Birthdays",
			expectedCount: assert.Greater,
			getScope: func() selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					EventCalendars([]string{user}, selectors.Any())[0]
			},
		}, {
			name:          "Default Calendar",
			contains:      DefaultCalendar,
			expectedCount: assert.Equal,
			getScope: func() selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					EventCalendars([]string{user}, []string{DefaultCalendar})[0]
			},
		}, {
			name:          "Default Mail",
			contains:      DefaultMailFolder,
			expectedCount: assert.Equal,
			getScope: func() selectors.ExchangeScope {
				return selectors.
					NewExchangeBackup().
					MailFolders([]string{user}, []string{DefaultMailFolder})[0]
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			qp := graph.QueryParams{
				User:        user,
				Scope:       test.getScope(),
				FailFast:    failFast,
				Credentials: credentials,
			}
			collections, err := GetContainers(ctx, qp, service)
			assert.NoError(t, err)
			test.expectedCount(t, len(collections), containerCount)

			keys := make([]string, 0, len(collections))
			for _, k := range collections {
				keys = append(keys, *k.GetDisplayName())
			}
			t.Logf("Collections Made: %v\n", keys)
			assert.Contains(t, keys, test.contains)
		})
	}
}
