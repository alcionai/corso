package m365

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type M365IntegrationSuite struct {
	suite.Suite
}

func TestM365IntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(M365IntegrationSuite))
}

func (suite *M365IntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.M365AcctCredEnvs)
	require.NoError(suite.T(), err)
}

func (suite *M365IntegrationSuite) TestUsers() {
	acct := tester.NewM365Account(suite.T())

	users, err := Users(acct)
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), users)
	require.Greater(suite.T(), len(users), 0)
}
