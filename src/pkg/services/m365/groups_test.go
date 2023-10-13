package m365_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/errs"
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

func (suite *GroupsIntgSuite) TestGroupByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	gid := tconfig.M365TeamID(t)

	group, err := m365.GroupByID(ctx, suite.acct, gid, count.New())
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, gid, group.ID, "must match expected id")
	assert.NotEmpty(t, group.DisplayName)
}

func (suite *GroupsIntgSuite) TestGroupByID_ByEmail() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	gid := tconfig.M365TeamID(t)

	group, err := m365.GroupByID(ctx, suite.acct, gid, count.New())
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, gid, group.ID, "must match expected id")
	assert.NotEmpty(t, group.DisplayName)

	gemail := tconfig.M365TeamEmail(t)

	groupByEmail, err := m365.GroupByID(ctx, suite.acct, gemail, count.New())
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, groupByEmail, group, "must be the same group as the one gotten by id")
}

func (suite *GroupsIntgSuite) TestGroupByID_notFound() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	group, err := m365.GroupByID(ctx, suite.acct, uuid.NewString(), count.New())
	require.Nil(t, group)
	require.ErrorIs(t, err, graph.ErrResourceOwnerNotFound, clues.ToCore(err))
	require.True(t, errs.Is(err, errs.ResourceOwnerNotFound))
}

func (suite *GroupsIntgSuite) TestGroups() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	groups, err := m365.Groups(
		ctx,
		suite.acct,
		fault.New(true))
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

func (suite *GroupsIntgSuite) TestSitesInGroup() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	gid := tconfig.M365TeamID(t)

	sites, err := m365.SitesInGroup(
		ctx,
		suite.acct,
		gid,
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)
}

func (suite *GroupsIntgSuite) TestGroupsMap() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	gm, err := m365.GroupsMap(
		ctx,
		suite.acct,
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, gm)

	for _, gid := range gm.IDs() {
		suite.Run("group_"+gid, func() {
			t := suite.T()

			assert.NotEmpty(t, gid)

			name, ok := gm.NameOf(gid)
			assert.True(t, ok)
			assert.NotEmpty(t, name)
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
					})
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

			groups, err := m365.Groups(
				ctx,
				test.acct(t),
				fault.New(true))
			assert.Empty(t, groups, "returned no groups")
			assert.NotNil(t, err)
		})
	}
}
