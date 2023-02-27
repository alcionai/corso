package m365

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type M365IntegrationSuite struct {
	tester.Suite
}

func TestM365IntegrationSuite(t *testing.T) {
	suite.Run(t, &M365IntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *M365IntegrationSuite) TestUsers() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(suite.T())
	)

	users, err := Users(ctx, acct, fault.New(true))
	require.NoError(t, err)
	require.NotNil(t, users)
	require.Greater(t, len(users), 0)

	for _, u := range users {
		suite.Run("user_"+u.ID, func() {
			t := suite.T()

			assert.NotEmpty(t, u.ID)
			assert.NotEmpty(t, u.PrincipalName)
			assert.NotEmpty(t, u.Name)
		})
	}
}
