package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/its"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/pkg/errs"
	"github.com/alcionai/canario/src/pkg/errs/core"
	"github.com/alcionai/canario/src/pkg/fault"
)

type GroupsIntgSuite struct {
	tester.Suite
	cli  client
	m365 its.M365IntgTestSetup
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

	suite.m365 = its.GetM365(t)

	// will init the concurrency limiter
	var err error

	suite.cli, err = NewM365Client(ctx, suite.m365.Acct)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *GroupsIntgSuite) TestGroupByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	gid := suite.m365.Group.ID

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

	gid := suite.m365.Group.ID

	group, err := suite.cli.GroupByID(ctx, gid)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, gid, group.ID, "must match expected id")
	assert.NotEmpty(t, group.DisplayName)

	gemail := suite.m365.Group.Email

	groupByEmail, err := suite.cli.GroupByID(ctx, gemail)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, groupByEmail, group, "must be the same group as the one gotten by id")
}

func (suite *GroupsIntgSuite) TestTeamByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	gid := suite.m365.Group.ID

	group, err := suite.cli.TeamByID(ctx, gid)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, group)

	assert.Equal(t, gid, group.ID, "must match expected id")
	assert.NotEmpty(t, group.DisplayName)
}

func (suite *GroupsIntgSuite) TestGroupByID_notFound() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	group, err := suite.cli.GroupByID(ctx, uuid.NewString())
	require.Nil(t, group)
	require.ErrorIs(t, err, core.ErrNotFound, clues.ToCore(err))
	require.True(t, errs.Is(err, core.ErrNotFound))
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
			if group.ID == suite.m365.Group.ID {
				assert.True(t, group.IsTeam)
			}
		})
	}
}

func (suite *GroupsIntgSuite) TestSitesInGroup() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	gid := suite.m365.Group.ID

	sites, err := suite.cli.SitesInGroup(ctx, gid, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)
}

// ---------------------------------------------------------------------------
// unit tests
// ---------------------------------------------------------------------------

type GroupsUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsUnitSuite) TestParseGroupFromTeamable() {
	id := uuid.NewString()
	name := uuid.NewString()

	table := []struct {
		name      string
		team      func() models.Teamable
		expectErr assert.ErrorAssertionFunc
		expect    Group
	}{
		{
			name: "good team",
			team: func() models.Teamable {
				team := models.NewTeam()
				team.SetId(ptr.To(id))
				team.SetDisplayName(ptr.To(name))

				return team
			},
			expectErr: assert.NoError,
			expect: Group{
				ID:          id,
				DisplayName: name,
				IsTeam:      true,
			},
		},
		{
			name: "no display name",
			team: func() models.Teamable {
				team := models.NewTeam()
				team.SetId(ptr.To(id))

				return team
			},
			expectErr: assert.NoError,
			expect: Group{
				ID:          id,
				DisplayName: "",
				IsTeam:      true,
			},
		},
		{
			name: "no id",
			team: func() models.Teamable {
				team := models.NewTeam()
				team.SetDisplayName(ptr.To(name))

				return team
			},
			expectErr: assert.Error,
			expect:    Group{},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := parseGroupFromTeamable(ctx, test.team())
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, test.expect, *result)
			}
		})
	}
}
