package connector

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	ctesting "github.com/alcionai/corso/internal/testing"
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
	// Verify Items() call returns an iterable channel(e.g. a channel that has been closed)
	channel := collectionList[0].Items()
	for object := range channel {
		buf := &bytes.Buffer{}
		_, err := buf.ReadFrom(object.ToReader())
		assert.Nil(suite.T(), err, "received a buf.Read error")
	}
	status := suite.connector.AwaitStatus()
	assert.NotNil(suite.T(), status, "status not blocking on async call")

	exchangeData := collectionList[0]
	suite.Greater(len(exchangeData.FullPath()), 2)
}

//TestGraphConnector_restoreMessages uses mock data to ensure GraphConnector
// is able to restore a messageable item to a Mailbox.
func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_restoreMessages() {
	user := "TEST_GRAPH_USER" // user.GetId()
	evs, err := ctesting.GetRequiredEnvVars(user)
	if err != nil {
		suite.T().Skipf("Environment not configured: %v\n", err)
	}
	mdc := mockconnector.NewMockExchangeCollection([]string{"tenant", evs[user], mailCategory, "Inbox"}, 1)
	err = suite.connector.RestoreMessages(context.Background(), []data.Collection{mdc})
	assert.NoError(suite.T(), err)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_CreateAndDeleteFolder() {
	user := "lidiah@8qzvrj.onmicrosoft.com"
	folderName := uuid.NewString() // todo - replace with danny's fix #391
	aFolder, err := exchange.CreateMailFolder(&suite.connector.graphService, user, folderName)
	assert.NoError(suite.T(), err, support.ConnectorStackErrorTrace(err))
	if aFolder != nil {
		err = exchange.DeleteMailFolder(&suite.connector.graphService, user, *aFolder.GetId())
		assert.NoError(suite.T(), err)
	}
}
