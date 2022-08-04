package connector

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/internal/tester"
	"github.com/alcionai/corso/pkg/selectors"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
	user      string
}

func TestGraphConnectorIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(GraphConnectorIntegrationSuite))
}

func (suite *GraphConnectorIntegrationSuite) SetupSuite() {
	if err := tester.RunOnAny(tester.CorsoCITests); err != nil {
		suite.T().Skip(err)
	}

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)

	a, err := tester.NewM365Account()
	require.NoError(suite.T(), err)

	suite.connector, err = NewGraphConnector(a)
	suite.NoError(err)
	suite.user = "lidiah@8qzvrj.onmicrosoft.com"
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector() {
	tester.LogTimeOfTest(suite.T())
	suite.NotNil(suite.connector)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_setTenantUsers() {
	err := suite.connector.setTenantUsers()
	assert.NoError(suite.T(), err)
	suite.Greater(len(suite.connector.Users), 0)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_ExchangeDataCollection() {
	userID, err := tester.M365UserID()
	require.NoError(suite.T(), err)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{userID}))
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

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_MailRegressionTest() {
	t := suite.T()
	user, err := tester.M365UserID()
	require.NoError(t, err)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{user}))
	eb, err := sel.ToExchangeBackup()
	require.NoError(t, err)
	var mailScope selectors.ExchangeScope
	all := eb.Scopes()
	for _, scope := range all {
		fmt.Printf("%v\n", scope)
		if scope.IncludesCategory(selectors.ExchangeMail) {
			mailScope = scope
		}
	}

	collection, err := suite.connector.createCollections(context.Background(), mailScope)
	require.NoError(t, err)
	for _, edc := range collection {
		testName := strings.Join(edc.FullPath(), " ")
		suite.T().Run(testName, func(t *testing.T) {
			streamChannel := edc.Items()
			// Verify that each message can be restored
			for stream := range streamChannel {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				suite.NoError(err)
				suite.NotZero(read)
				message, err := support.CreateMessageFromBytes(buf.Bytes())
				suite.NotNil(message)
				suite.NoError(err)

			}
		})
	}
}

//TestGraphConnector_restoreMessages uses mock data to ensure GraphConnector
// is able to restore a messageable item to a Mailbox.
func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_restoreMessages() {
	user := "TEST_GRAPH_USER" // user.GetId()
	evs, err := tester.GetRequiredEnvVars(user)
	if err != nil {
		suite.T().Skipf("Environment not configured: %v\n", err)
	}
	mdc := mockconnector.NewMockExchangeCollection([]string{"tenant", evs[user], mailCategory, "Inbox"}, 1)
	err = suite.connector.RestoreMessages(context.Background(), []data.Collection{mdc})
	assert.NoError(suite.T(), err)
}

///------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------

//  TestGraphConnector_CreateAndDeleteFolder ensures msgraph application has the ability
//  to create and remove folders within the tenant
func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_CreateAndDeleteFolder() {
	userID, err := tester.M365UserID()
	require.NoError(suite.T(), err)
	now := time.Now()
	folderName := "TestFolder: " + common.FormatSimpleDateTime(now)
	aFolder, err := exchange.CreateMailFolder(&suite.connector.graphService, userID, folderName)
	assert.NoError(suite.T(), err, support.ConnectorStackErrorTrace(err))
	if aFolder != nil {
		err = exchange.DeleteMailFolder(suite.connector.Service(), userID, *aFolder.GetId())
		assert.NoError(suite.T(), err)
	}
}

// TestGraphConnector_GetMailFolderID verifies the ability to retrieve folder ID of folders
// at the top level of the file tree
func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_GetMailFolderID() {
	userID, err := tester.M365UserID()
	require.NoError(suite.T(), err)
	folderName := "Inbox"
	folderID, err := exchange.GetMailFolderID(&suite.connector.graphService, folderName, userID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), folderID)
}
