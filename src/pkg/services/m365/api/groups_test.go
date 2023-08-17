package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
		name           string
		args           models.Groupable
		errCheck       assert.ErrorAssertionFunc
		errIsSkippable bool
	}{
		{
			name: "Valid group ",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetId(ptr.To("id"))
				s.SetDisplayName(ptr.To("testgroup"))
				return s
			}(),
			errCheck: assert.NoError,
		},
		{
			name: "No name",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetId(ptr.To("id"))
				return s
			}(),
			errCheck: assert.Error,
		},
		{
			name: "No ID",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetDisplayName(ptr.To("testgroup"))
				return s
			}(),
			errCheck: assert.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := api.ValidateGroup(test.args)
			test.errCheck(t, err, clues.ToCore(err))

			if test.errIsSkippable {
				assert.ErrorIs(t, err, api.ErrKnownSkippableCase)
			}
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

func (suite *GroupsIntgSuite) TestGetAllGroups() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	groups, err := suite.its.ac.
		Groups().
		GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(groups), "must have at least one group")
}

func (suite *GroupsIntgSuite) TestGetAll() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	teams, err := suite.its.ac.
		Groups().
		GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(teams), "must have at least one teams")

	groups, err := suite.its.ac.
		Groups().
		GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(groups), "must have at least one group")

	var isTeam bool

	if len(groups) > len(teams) {
		isTeam = true
	}

	assert.True(t, isTeam, "must only return teams")
}

func (suite *GroupsIntgSuite) TestTeams_GetByID() {
	var (
		t      = suite.T()
		teamID = tconfig.M365TeamID(t)
	)

	teamsAPI := suite.its.ac.Groups()

	table := []struct {
		name      string
		id        string
		expectErr func(*testing.T, error)
	}{
		{
			name: "3 part id",
			id:   teamID,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "malformed id",
			id:   uuid.NewString(),
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "random id",
			id:   uuid.NewString() + "," + uuid.NewString(),
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},

		{
			name: "malformed url",
			id:   "barunihlda",
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := teamsAPI.GetByID(ctx, test.id)
			test.expectErr(t, err)
		})
	}
}

func (suite *GroupsIntgSuite) TestGroups_GetByID() {
	var (
		t       = suite.T()
		groupID = tconfig.M365GroupID(t)
	)

	groupsAPI := suite.its.ac.Groups()

	table := []struct {
		name      string
		id        string
		expectErr func(*testing.T, error)
	}{
		{
			name: "3 part id",
			id:   groupID,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "malformed id",
			id:   uuid.NewString(),
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "random id",
			id:   uuid.NewString() + "," + uuid.NewString(),
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},

		{
			name: "malformed url",
			id:   "barunihlda",
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := groupsAPI.GetByID(ctx, test.id)
			test.expectErr(t, err)
		})
	}
}
