package betasdk

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/zeebo/assert"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type BetaClientSuite struct {
	suite.Suite
	credentials account.M365Config
}

func TestBetaClientSuite(t *testing.T) {
	suite.Run(t, new(BetaClientSuite))
}

func (suite *BetaClientSuite) SetupSuite() {
	t := suite.T()
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365
}

func (suite *BetaClientSuite) TestCreateSite() {
	t := suite.T()
	adpt, err := graph.CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
	)

	require.NoError(t, err)

	client := NewBetaClient(adpt)
	assert.NotNil(t, client)

}
