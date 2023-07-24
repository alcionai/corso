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

type TeamsUnitSuite struct {
	tester.Suite
}

func TestTeamsUnitSuite(t *testing.T) {
	suite.Run(t, &TeamsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *TeamsUnitSuite) TestValidateTeams() {
	team := models.NewTeam()
	team.SetDisplayName(ptr.To("testteam"))
	team.SetId(ptr.To("testID"))

	tests := []struct {
		name           string
		args           models.Teamable
		errCheck       assert.ErrorAssertionFunc
		errIsSkippable bool
	}{
		{
			name: "Valid Team",
			args: func() *models.Team {
				s := models.NewTeam()
				s.SetId(ptr.To("id"))
				s.SetDisplayName(ptr.To("testTeam"))
				return s
			}(),
			errCheck: assert.NoError,
		},
		{
			name: "No name",
			args: func() *models.Team {
				s := models.NewTeam()
				s.SetId(ptr.To("id"))
				return s
			}(),
			errCheck: assert.Error,
		},
		{
			name: "No ID",
			args: func() *models.Team {
				s := models.NewTeam()
				s.SetDisplayName(ptr.To("testTeam"))
				return s
			}(),
			errCheck: assert.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := api.ValidateTeams(test.args)
			test.errCheck(t, err, clues.ToCore(err))

			if test.errIsSkippable {
				assert.ErrorIs(t, err, api.ErrKnownSkippableCase)
			}
		})
	}
}

type TeamsIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestTeamsIntgSuite(t *testing.T) {
	suite.Run(t, &TeamsIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *TeamsIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *TeamsIntgSuite) TestGetAllTeams() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	teams, err := suite.its.ac.
		Teams().
		GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(teams), "must have at least one team")
}

func (suite *TeamsIntgSuite) TestTeams_GetByID() {
	var (
		t      = suite.T()
		teamID = tconfig.M365TeamsID(t)
	)

	teamsAPI := suite.its.ac.Teams()

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
