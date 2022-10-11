package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type ServiceFunctionsIntegrationSuite struct {
	suite.Suite
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
	suite.m365UserID = tester.M365UserID(suite.T())
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
			params := graph.QueryParams{
				User:        test.user,
				Scope:       nil,
				FailFast:    false,
				Credentials: suite.creds,
			}
			cals, err := GetAllContactFolders(ctx, params, gs, test.user, test.contains)
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

func (suite *ServiceFunctionsIntegrationSuite) TestCollectContainers() {
	ctx, flush := tester.NewContext()
	defer flush()

	failFast := false
	containerCount := 1
	t := suite.T()
	user := tester.M365UserID(t)
	a := tester.NewM365Account(t)
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
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{user}, selectors.Any()))

				scopes := sel.Scopes()
				assert.Equal(t, len(scopes), 1)

				return scopes[0]
			},
		}, {
			name:          "Default Calendar",
			contains:      DefaultCalendar,
			expectedCount: assert.Equal,
			getScope: func() selectors.ExchangeScope {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{user}, []string{DefaultCalendar}))

				scopes := sel.Scopes()
				assert.Equal(t, len(scopes), 1)

				return scopes[0]
			},
		}, {
			name:          "Default Mail",
			contains:      DefaultMailFolder,
			expectedCount: assert.Equal,
			getScope: func() selectors.ExchangeScope {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.MailFolders([]string{user}, []string{DefaultMailFolder}))

				scopes := sel.Scopes()
				assert.Equal(t, len(scopes), 1)

				return scopes[0]
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
			collections := make(map[string]*Collection)
			err := CollectFolders(ctx, qp, collections, nil, nil)
			assert.NoError(t, err)
			test.expectedCount(t, len(collections), containerCount)

			keys := make([]string, 0, len(collections))
			for k := range collections {
				keys = append(keys, k)
			}
			t.Logf("Collections Made: %v\n", keys)
			assert.Contains(t, keys, test.contains)
		})
	}
}
