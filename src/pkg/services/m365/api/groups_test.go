package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
)

type GroupUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupUnitSuite) TestValidateGroup() {
	group := models.NewGroup()
	group.SetDisplayName(ptr.To("testgroup"))
	group.SetId(ptr.To("testID"))

	tests := []struct {
		name      string
		args      models.Groupable
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "Valid group ",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetId(ptr.To("id"))
				s.SetDisplayName(ptr.To("testgroup"))
				return s
			}(),
			expectErr: assert.NoError,
		},
		{
			name: "No name",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetId(ptr.To("id"))
				return s
			}(),
			expectErr: assert.Error,
		},
		{
			name: "No ID",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetDisplayName(ptr.To("testgroup"))
				return s
			}(),
			expectErr: assert.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := validateGroup(test.args)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}

type GroupsIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestGroupsIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *GroupsIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *GroupsIntgSuite) TestGetAll() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	groups, err := suite.its.ac.
		Groups().
		GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(groups), "must have at least one group")
}

func (suite *GroupsIntgSuite) TestGetTeamByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	team, err := suite.its.ac.Groups().GetTeamByID(ctx, suite.its.group.id, CallConfig{})
	require.NoError(t, err, "getting team by ID")
	require.NotNil(t, team, "must have valid team")
}

func (suite *GroupsIntgSuite) TestGetAllSites() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	channels, err := suite.its.ac.
		Channels().GetChannels(ctx, suite.its.group.id)
	require.NoError(t, err, "getting channels")
	require.NotZero(t, len(channels), "must have at least one channel")

	siteCount := 1

	for _, c := range channels {
		if ptr.Val(c.GetMembershipType()) != models.STANDARD_CHANNELMEMBERSHIPTYPE {
			siteCount++
		}
	}

	sites, err := suite.its.ac.
		Groups().
		GetAllSites(ctx, suite.its.group.id, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(sites), "must have at least one site")
	require.Equal(t, siteCount, len(sites), "incorrect number of sites")
}

// GetAllSites for Groups that are not Teams should return just the root site
func (suite *GroupsIntgSuite) TestGetAllSitesNonTeam() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	group, err := suite.its.ac.Groups().GetByID(ctx, suite.its.nonTeamGroup.id, CallConfig{})
	require.NoError(t, err)
	require.False(t, IsTeam(ctx, group), "group should not be a team for this test")

	sites, err := suite.its.ac.
		Groups().
		GetAllSites(ctx, suite.its.nonTeamGroup.id, fault.New(true))
	require.NoError(t, err)
	require.Equal(t, 1, len(sites), "incorrect number of sites")
}

func (suite *GroupsIntgSuite) TestGroups_GetByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		groupID     = suite.its.group.id
		groupsEmail = suite.its.group.email
		groupsAPI   = suite.its.ac.Groups()
	)

	grp, err := groupsAPI.GetByID(ctx, groupID, CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name      string
		id        string
		expectErr func(t *testing.T, err error)
	}{
		{
			name: "valid id",
			id:   groupID,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "valid email as identifier",
			id:   groupsEmail,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "invalid id",
			id:   uuid.NewString(),
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, core.ErrResourceOwnerNotFound, clues.ToCore(err))
			},
		},
		{
			name: "valid display name",
			id:   ptr.Val(grp.GetDisplayName()),
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "invalid displayName",
			id:   "jabberwocky",
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, core.ErrResourceOwnerNotFound, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := groupsAPI.GetByID(ctx, test.id, CallConfig{})
			test.expectErr(t, err)
		})
	}
}

func (suite *GroupsIntgSuite) TestGroups_GetAllIDsAndNames() {
	t := suite.T()
	groupsAPI := suite.its.ac.Groups()

	ctx, flush := tester.NewContext(t)
	defer flush()

	gm, err := groupsAPI.GetAllIDsAndNames(ctx, fault.New(true))
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

func (suite *GroupsIntgSuite) TestGroups_GetByID_mockResourceLockedErrs() {
	gID := uuid.NewString()

	table := []struct {
		name  string
		id    string
		setup func(t *testing.T)
	}{
		{
			name: "by name",
			id:   "g",
			setup: func(t *testing.T) {
				err := graphTD.ODataErr(string(graph.NotAllowed))

				interceptV1Path("groups").
					Reply(403).
					JSON(graphTD.ParseableToMap(t, err))
				interceptV1Path("teams").
					Reply(403).
					JSON(graphTD.ParseableToMap(t, err))
			},
		},
		{
			name: "by id",
			id:   gID,
			setup: func(t *testing.T) {
				err := graphTD.ODataErr(string(graph.NotAllowed))

				interceptV1Path("groups", gID).
					Reply(403).
					JSON(graphTD.ParseableToMap(t, err))
				interceptV1Path("teams", gID).
					Reply(403).
					JSON(graphTD.ParseableToMap(t, err))
			},
		},
		{
			name: "by id - matches error message",
			id:   gID,
			setup: func(t *testing.T) {
				err := graphTD.ODataErrWithMsg(
					string(graph.AuthenticationError),
					"AADSTS500014: The service principal for resource 'beefe6b7-f5df-413d-ac2d-abf1e3fd9c0b' "+
						"is disabled. This indicate that a subscription within the tenant has lapsed, or that the "+
						"administrator for this tenant has disabled the application, preventing tokens from being "+
						"issued for it. Trace ID: dead78e1-0830-4edf-bea7-f0a445620100 Correlation ID: "+
						"deadbeef-7f1e-4578-8215-36004a2c935c Timestamp: 2023-12-05 19:31:01Z")

				interceptV1Path("groups", gID).
					Reply(403).
					JSON(graphTD.ParseableToMap(t, err))
				interceptV1Path("teams", gID).
					Reply(403).
					JSON(graphTD.ParseableToMap(t, err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)

			defer flush()
			defer gock.Off()

			test.setup(t)

			// Verify both GetByID and GetTeamByID APIs handle this error
			_, err := suite.its.gockAC.
				Groups().
				GetByID(ctx, test.id, CallConfig{})
			require.ErrorIs(t, err, graph.ErrResourceLocked, clues.ToCore(err))

			_, err = suite.its.gockAC.
				Groups().
				GetTeamByID(ctx, test.id, CallConfig{})
			require.ErrorIs(t, err, graph.ErrResourceLocked, clues.ToCore(err))
		})
	}
}
