package m365

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type M365IntegrationSuite struct {
	suite.Suite
}

func TestM365IntegrationSuite(t *testing.T) {
	tester.RunOnAny(t, tester.CorsoCITests)
	suite.Run(t, new(M365IntegrationSuite))
}

func (suite *M365IntegrationSuite) SetupSuite() {
	tester.MustGetEnvSets(suite.T(), tester.M365AcctCredEnvs)
}

func (suite *M365IntegrationSuite) TestUsers() {
	ctx, flush := tester.NewContext()
	defer flush()

	acct := tester.NewM365Account(suite.T())

	users, err := Users(ctx, acct)
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), users)
	require.Greater(suite.T(), len(users), 0)

	for _, u := range users {
		suite.T().Log(u)
		assert.NotEmpty(suite.T(), u.ID)
		assert.NotEmpty(suite.T(), u.PrincipalName)
		assert.NotEmpty(suite.T(), u.Name)
	}
}
