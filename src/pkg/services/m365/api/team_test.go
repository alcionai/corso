package api_test

import (
	"testing"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TeamsUnitSuite struct {
	tester.Suite
}

func TestTeamsUnitSuite(t *testing.T) {
	suite.Run(t, &TeamsUnitSuite{Suite: tester.NewUnitSuite(t)})
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

func (suite *TeamsIntgSuite) TestGetAll() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	teams, err := suite.its.ac.
		Groups().
		GetAll(ctx, true, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(teams), "must have at least one team")

	for _, team := range teams {
		assert.NotEmpty(t, ptr.Val(team.GetDisplayName()), "must not return onedrive teams")
	}
}
