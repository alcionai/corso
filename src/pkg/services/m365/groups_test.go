package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/errs"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type GroupsIntgSuite struct {
	tester.Suite
	cli client
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

	acct := tconfig.NewM365Account(t)

	var err error

	// will init the concurrency limiter
	suite.cli, err = NewM365Client(ctx, acct)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *GroupsIntgSuite) TestGroupByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	gid := tconfig.M365TeamID(t)

	group, err := suite.cli.GroupByID(ctx, gid)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, gid, group.ID, "must match expected id")
	assert.NotEmpty(t, group.DisplayName)
}

func (suite *GroupsIntgSuite) TestGroupByID_ByEmail() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	gid := tconfig.M365TeamID(t)

	group, err := suite.cli.GroupByID(ctx, gid)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, gid, group.ID, "must match expected id")
	assert.NotEmpty(t, group.DisplayName)

	gemail := tconfig.M365TeamEmail(t)

	groupByEmail, err := suite.cli.GroupByID(ctx, gemail)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, groupByEmail, group, "must be the same group as the one gotten by id")
}

func (suite *GroupsIntgSuite) TestGroupByID_notFound() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	group, err := suite.cli.GroupByID(ctx, uuid.NewString())
	require.Nil(t, group)
	require.ErrorIs(t, err, graph.ErrResourceOwnerNotFound, clues.ToCore(err))
	require.True(t, errs.Is(err, errs.ResourceOwnerNotFound))
}

func (suite *GroupsIntgSuite) TestGroups() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	groups, err := suite.cli.Groups(ctx, fault.New(true))
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

	gid := tconfig.M365TeamID(t)

	sites, err := suite.cli.SitesInGroup(ctx, gid, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)
}

func (suite *GroupsIntgSuite) TestGroupsMap() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	gm, err := suite.cli.GroupsMap(ctx, fault.New(true))
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
