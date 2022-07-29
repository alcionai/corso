package connector

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/selectors"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
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
	sel.Include(sel.Users([]string{"lidiah@8qzvrj.onmicrosoft.com"}))
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

//TODO add mockdata for this test... Will not pass as is
func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_restoreMessages() {
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

	ds := exchange.NewStream("test", bytes, details.ExchangeInfo{})
	edc := exchange.NewCollection("tenant", []string{"tenantId", evs[user], mailCategory, "Inbox"})

	edc.PopulateCollection(&ds)
	//edc.FinishPopulation()
	err = suite.connector.RestoreMessages(context.Background(), []data.Collection{&edc})
	assert.NoError(suite.T(), err)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_createDeleteFolder() {
	user := "lidiah@8qzvrj.onmicrosoft.com"
	folderName := uuid.NewString() // todo - replace with danny's fix #391
	aFolder, err := createMailFolder(suite.connector.graphService, user, folderName)
	assert.NoError(suite.T(), err, support.ConnectorStackErrorTrace(err))
	if aFolder != nil {
		err = deleteMailFolder(suite.connector.graphService, user, *aFolder.GetId())
		assert.NoError(suite.T(), err)
	}
}
