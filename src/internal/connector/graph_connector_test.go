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

	a := tester.NewM365Account(suite.T())

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

// TestExchangeDataCollection verifies interface between operation and
// GraphConnector remains stable to receive a non-zero amount of Collections
// for the Exchange Package. Enabled exchange applications:
// - mail
func (suite *GraphConnectorIntegrationSuite) TestExchangeDataCollection() {
	userID := tester.M365UserID(suite.T())
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

// TestMailSerializationRegression verifies that all mail data stored in the
// test account can be successfully downloaded into bytes and restored into
// M365 mail objects
func (suite *GraphConnectorIntegrationSuite) TestMailSerializationRegression() {
	t := suite.T()
	user := tester.M365UserID(suite.T())
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

// TestContactBackupSequence verifies ability to query contact items
// and to store contact within Collection. Downloaded contacts are run through
// a regression test to ensure that downloaded items can be uploaded.
func (suite *GraphConnectorIntegrationSuite) TestContactBackupSequence() {
	userID := tester.M365UserID(suite.T())
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{userID}))
	eb, err := sel.ToExchangeBackup()
	require.NoError(suite.T(), err)
	scopes := eb.Scopes()
	var contactsOnly selectors.ExchangeScope
	for _, scope := range scopes {
		if scope.IncludesCategory(selectors.ExchangeContactFolder) {
			contactsOnly = scope
		}
	}
	collections, err := suite.connector.createCollections(context.Background(), contactsOnly)
	assert.NoError(suite.T(), err)
	number := 0
	for _, edc := range collections {
		testName := fmt.Sprintf("%s_ContactFolder_%d", edc.FullPath()[1], number)
		suite.T().Run(testName, func(t *testing.T) {
			streamChannel := edc.Items()
			for stream := range streamChannel {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				suite.NoError(err)
				suite.NotZero(read)
				contact, err := support.CreateContactFromBytes(buf.Bytes())
				assert.NotNil(t, contact)
				assert.NoError(t, err)

			}
			number++
		})
	}
	suite.Greater(len(collections), 0)
}

// TestRestoreMessages uses mock data to ensure GraphConnector
// is able to restore a single messageable item to a Mailbox.
func (suite *GraphConnectorIntegrationSuite) TestRestoreMessages() {
	user := tester.M365UserID(suite.T())
	if len(user) == 0 {
		suite.T().Skip("Environment not configured: missing m365 test user")
	}
	mdc := mockconnector.NewMockExchangeCollection([]string{"tenant", user, mailCategory, "Inbox"}, 1)
	err := suite.connector.RestoreMessages(context.Background(), []data.Collection{mdc})
	assert.NoError(suite.T(), err)
}

// TestGraphConnector_SingleMailFolderCollectionQuery verifies single folder support
// for Backup operation
func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_SingleMailFolderCollectionQuery() {
	t := suite.T()
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.MailFolders([]string{suite.user}, []string{"Inbox"}))
	scopes := sel.Scopes()
	for _, scope := range scopes {
		collections, err := suite.connector.createCollections(context.Background(), scope)
		require.NoError(t, err)
		suite.Equal(len(collections), 1)
		for _, edc := range collections {
			number := 0
			streamChannel := edc.Items()
			// Verify that each message can be restored
			for stream := range streamChannel {
				testName := fmt.Sprintf("%s_InboxMessage_%d", edc.FullPath()[1], number)
				suite.T().Run(testName, func(t *testing.T) {
					buf := &bytes.Buffer{}
					read, err := buf.ReadFrom(stream.ToReader())
					suite.NoError(err)
					suite.NotZero(read)
					message, err := support.CreateMessageFromBytes(buf.Bytes())
					suite.NotNil(message)
					suite.NoError(err)
					number++
				})
			}
		}
	}
}

// TestEventsBackupSequence ensures functionality of createCollections
// to be able to successfully query, download and restore event objects
func (suite *GraphConnectorIntegrationSuite) TestEventsBackupSequence() {
	t := suite.T()
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Events([]string{suite.user}, []string{selectors.AnyTgt}))
	scopes := sel.Scopes()
	assert.Greater(t, len(scopes), 0)
	collections, err := suite.connector.createCollections(context.Background(), scopes[0])
	require.NoError(t, err)
	suite.Greater(len(collections), 0)
	for _, edc := range collections {
		streamChannel := edc.Items()
		number := 0
		for stream := range streamChannel {
			testName := fmt.Sprintf("%s_Event_%d", edc.FullPath()[1], number)
			suite.T().Run(testName, func(t *testing.T) {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				suite.NoError(err)
				suite.NotZero(read)
				event, err := support.CreateEventFromBytes(buf.Bytes())
				assert.NotNil(t, event)
				assert.NoError(t, err)
			})
		}
	}
}

///------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------

//  TestCreateAndDeleteFolder ensures GraphConnector has the ability
//  to create and remove folders within the tenant
func (suite *GraphConnectorIntegrationSuite) TestCreateAndDeleteFolder() {
	userID := tester.M365UserID(suite.T())
	now := time.Now()
	folderName := "TestFolder: " + common.FormatSimpleDateTime(now)
	aFolder, err := exchange.CreateMailFolder(&suite.connector.graphService, userID, folderName)
	assert.NoError(suite.T(), err, support.ConnectorStackErrorTrace(err))
	if aFolder != nil {
		err = exchange.DeleteMailFolder(suite.connector.Service(), userID, *aFolder.GetId())
		assert.NoError(suite.T(), err)
	}
}

// TestGetMailFolderID verifies the ability to retrieve folder ID of folders
// at the top level of the file tree
func (suite *GraphConnectorIntegrationSuite) TestGetMailFolderID() {
	userID := tester.M365UserID(suite.T())
	folderName := "Inbox"
	folderID, err := exchange.GetMailFolderID(&suite.connector.graphService, folderName, userID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), folderID)
}
