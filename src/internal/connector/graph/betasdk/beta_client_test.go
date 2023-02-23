package betasdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type BetaClientSuite struct {
	tester.Suite
	credentials account.M365Config
}

func TestBetaClientSuite(t *testing.T) {
	suite.Run(t, &BetaClientSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
			tester.CorsoGraphConnectorTests),
	})
}

func (suite *BetaClientSuite) SetupSuite() {
	t := suite.T()
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365
}

func (suite *BetaClientSuite) TestCreateBetaClient() {
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

// TestBasicClientGetFunctionality. Tests that adapter is able
// to parse retrieved Site Page. Additional tests should
// be handled within the /internal/connector/sharepoint when
// additional features are added.
func (suite *BetaClientSuite) TestBasicClientGetFunctionality() {
	ctx, flush := tester.NewContext()
	defer flush()
	t := suite.T()
	adpt, err := graph.CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
	)

	require.NoError(t, err)
	client := NewBetaClient(adpt)
	require.NotNil(t, client)

	siteID := tester.M365SiteID(t)
	// TODO(dadams39) document allowable calls in main
	collection, err := client.SitesById(siteID).Pages().Get(ctx, nil)
	// Ensures that the client is able to receive data from beta
	// Not Registered Error: content type application/json does not have a factory registered to be parsed
	require.NoError(t, err)

	for _, page := range collection.GetValue() {
		assert.NotNil(t, page, "betasdk call for page does not return value.")

		if page != nil {
			t.Logf("Page :%s ", *page.GetName())
			assert.NotNil(t, page.GetId())
		}
	}
}
