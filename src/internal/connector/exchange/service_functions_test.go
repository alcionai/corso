package exchange

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
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
	gs := loadService(suite.T())
	ctx := context.Background()

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
			cals, err := GetAllCalendars(ctx, gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllContactFolders() {
	gs := loadService(suite.T())
	ctx := context.Background()

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
			cals, err := GetAllContactFolders(ctx, gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}

func (suite *ServiceFunctionsIntegrationSuite) TestGetAllMailFolders() {
	gs := loadService(suite.T())
	ctx := context.Background()

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
			cals, err := GetAllMailFolders(ctx, gs, test.user, test.contains)
			test.expectErr(t, err)
			test.expectCount(t, len(cals), 0)
		})
	}
}
