package connector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/support"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/selectors"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
	user      string
}

func TestGraphConnectorIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(GraphConnectorIntegrationSuite))
}

func (suite *GraphConnectorIntegrationSuite) SetupSuite() {
	if err := ctesting.RunOnAny(ctesting.CorsoCITests); err != nil {
		suite.T().Skip(err)
	}

	_, err := ctesting.GetRequiredEnvVars(ctesting.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	a, err := ctesting.NewM365Account()
	require.NoError(suite.T(), err)

	suite.connector, err = NewGraphConnector(a)
	suite.NoError(err)
	suite.user = "lidiah@8qzvrj.onmicrosoft.com"
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector() {
	ctesting.LogTimeOfTest(suite.T())
	suite.NotNil(suite.connector)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_setTenantUsers() {
	err := suite.connector.setTenantUsers()
	assert.NoError(suite.T(), err)
	suite.Greater(len(suite.connector.Users), 0)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_ExchangeDataCollection() {
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{suite.user}))
	collectionList, err := suite.connector.ExchangeDataCollection(context.Background(), sel.Selector)
	assert.NotNil(suite.T(), collectionList, "collection list")
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), suite.connector.awaitingMessages > 0)
	assert.Nil(suite.T(), suite.connector.status)
	status := suite.connector.AwaitStatus()
	assert.NotNil(suite.T(), status, "status not blocking on async call")

	exchangeData := collectionList[0]
	suite.Greater(len(exchangeData.FullPath()), 2)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_Restore() {
	user := "TEST_GRAPH_USER" // user.GetId()
	file := "TEST_GRAPH_FILE" // Test file should be sent or received by the user
	evs, err := ctesting.GetRequiredEnvVars(user, file)
	if err != nil {
		suite.T().Skipf("Environment not configured: %v\n", err)
	}
	bytes, err := ctesting.LoadAFile(evs[file]) // TEST_GRAPH_FILE should have a single Message && not present in target inbox
	if err != nil {
		suite.T().Skipf("Support file not accessible: %v\n", err)
	}
	ds := ExchangeData{id: "test", message: bytes}
	edc := NewExchangeDataCollection("tenant", []string{"tenantId", evs[user], mailCategory, "Inbox"})
	edc.PopulateCollection(&ds)
	edc.FinishPopulation()
	err = suite.connector.Restore(context.Background(), []DataCollection{&edc})
	assert.NoError(suite.T(), err)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_HasFolder() {
	response, err := HasMailFolder("Inbox", suite.user, suite.connector.graphService)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), response)
	response, err = HasMailFolder("A_Wacky_World_Of_NonExistance", suite.user, suite.connector.graphService)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), response)

}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_createDeleteFolder() {
	t := suite.T()
	user := "TEST_GRAPH_USER"
	folderName := "createdForTest"
	_, err := createMailFolder(suite.connector.graphService, suite.user, folderName)
	require.NoError(t, err, support.ConnectorStackErrorTrace(err))
	response, err := HasMailFolder(folderName, user, suite.connector.graphService)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
	require.NotNil(t, response)
	err = deleteMailFolder(suite.connector.graphService, suite.user, *response)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
	response, err = HasMailFolder(folderName, suite.user, suite.connector.graphService)
	assert.NoError(t, err)
	assert.Nil(t, response)

}
