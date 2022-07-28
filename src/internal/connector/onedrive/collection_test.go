package onedrive

import (
	"context"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/tester"
)

type OnedriveIntegrationSuite struct {
	suite.Suite
	client  *msgraphsdk.GraphServiceClient
	adapter *msgraphsdk.GraphRequestAdapter
	graph.Service
}

func (suite *OnedriveIntegrationSuite) Client() *msgraphsdk.GraphServiceClient {
	return suite.client
}

func (suite *OnedriveIntegrationSuite) Adapter() *msgraphsdk.GraphRequestAdapter {
	return suite.adapter
}

func (suite *OnedriveIntegrationSuite) ErrPolicy() bool {
	return false
}

func TestGraphConnectorIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(OnedriveIntegrationSuite))
}

func (suite *OnedriveIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	a, err := tester.NewM365Account()
	require.NoError(suite.T(), err)

	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	adapter, err := graph.CreateAdapter(m365.TenantID, m365.ClientID, m365.ClientSecret)
	require.NoError(suite.T(), err)
	suite.client = msgraphsdk.NewGraphServiceClient(adapter)
	suite.adapter = adapter
}

func (suite *OnedriveIntegrationSuite) TestOnedriveEnumeration() {
	tester.LogTimeOfTest(suite.T())
	collections, err := NewCollections("george.martinez@8qzvrj.onmicrosoft.com", suite).Get(context.Background())
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), collections)
}
