package m365_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365"
)

type GroupsIntgSuite struct {
	tester.Suite
	acct account.Account
}

func TestGroupsIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *GroupsIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.acct = tconfig.NewM365Account(t)
}

func (suite *GroupsIntgSuite) TestGroups() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	groups, err := m365.Groups(ctx, suite.acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, groups)

	for _, group := range groups {
		suite.Run("group_"+group.ID, func() {
			t := suite.T()

			assert.NotEmpty(t, group.ID)
			assert.NotEmpty(t, group.DisplayName)

			// at least one known group should be a team
			if group.ID == tconfig.M365TeamID(t) {
				assert.True(t, group.IsTeam)
			}
		})
	}
}

func (suite *GroupsIntgSuite) TestGroups_InvalidCredentials() {
	table := []struct {
		name string
		acct func(t *testing.T) account.Account
	}{
		{
			name: "Invalid Credentials",
			acct: func(t *testing.T) account.Account {
				a, err := account.NewAccount(
					account.ProviderM365,
					account.M365Config{
						M365: credentials.M365{
							AzureClientID:     "Test",
							AzureClientSecret: "without",
						},
						AzureTenantID: "data",
					},
				)
				require.NoError(t, err, clues.ToCore(err))

				return a
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			groups, err := m365.Groups(ctx, test.acct(t), fault.New(true))
			assert.Empty(t, groups, "returned no groups")
			assert.NotNil(t, err)
		})
	}
}
