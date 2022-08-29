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

func loadConnector(t *testing.T) *GraphConnector {
	a := tester.NewM365Account(t)
	connector, err := NewGraphConnector(a)
	require.NoError(t, err)
	return connector
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
	suite.connector = loadConnector(suite.T())
	suite.user = tester.M365UserID(suite.T())
	tester.LogTimeOfTest(suite.T())
}

// TestSetTenantUsers verifies GraphConnector's ability to query
// the users associated with the credentials
func (suite *GraphConnectorIntegrationSuite) TestSetTenantUsers() {
	newConnector := GraphConnector{
		tenant:      "test_tenant",
		Users:       make(map[string]string, 0),
		status:      nil,
		statusCh:    make(chan *support.ConnectorOperationStatus),
		credentials: suite.connector.credentials,
	}
	service, err := newConnector.createService(false)
	require.NoError(suite.T(), err)
	newConnector.graphService = *service

	suite.Equal(len(newConnector.Users), 0)
	err = newConnector.setTenantUsers()
	assert.NoError(suite.T(), err)
	suite.Greater(len(newConnector.Users), 0)
}

// TestExchangeDataCollection verifies interface between operation and
// GraphConnector remains stable to receive a non-zero amount of Collections
// for the Exchange Package. Enabled exchange applications:
// - mail
// - contacts
// - events
func (suite *GraphConnectorIntegrationSuite) TestExchangeDataCollection() {
	t := suite.T()
	connector := loadConnector(t)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{suite.user}))
	collectionList, err := connector.ExchangeDataCollection(context.Background(), sel.Selector)
	assert.NotNil(t, collectionList, "collection list")
	assert.NoError(t, err)
	assert.True(t, connector.awaitingMessages > 0)
	assert.Nil(t, connector.status)
	streams := make(map[string]<-chan data.Stream)
	// Verify Items() call returns an iterable channel(e.g. a channel that has been closed)
	for _, collection := range collectionList {
		temp := collection.Items()
		testName := collection.FullPath()[2]
		streams[testName] = temp
	}

	for i := 0; i < int(connector.awaitingMessages); i++ {
		status := connector.AwaitStatus()
		assert.NotNil(t, status)
	}

	for name, channel := range streams {
		suite.T().Run(name, func(t *testing.T) {
			t.Logf("Test: %s\t Items: %d", name, len(channel))
			for object := range channel {
				buf := &bytes.Buffer{}
				_, err := buf.ReadFrom(object.ToReader())
				assert.NoError(t, err, "received a buf.Read error")
			}
		})
	}
	exchangeData := collectionList[0]
	suite.Greater(len(exchangeData.FullPath()), 2)
}

// TestMailSerializationRegression verifies that all mail data stored in the
// test account can be successfully downloaded into bytes and restored into
// M365 mail objects
func (suite *GraphConnectorIntegrationSuite) TestMailSerializationRegression() {
	t := suite.T()
	connector := loadConnector(t)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.MailFolders([]string{suite.user}, selectors.Any()))
	eb, err := sel.ToExchangeBackup()
	require.NoError(t, err)
	scopes := eb.Scopes()
	suite.Len(scopes, 1)
	mailScope := scopes[0]
	collection, err := connector.createCollections(context.Background(), mailScope)
	require.NoError(t, err)
	for _, edc := range collection {
		testName := strings.Join(edc.FullPath(), " ")
		suite.T().Run(testName, func(t *testing.T) {
			streamChannel := edc.Items()
			// Verify that each message can be restored
			for stream := range streamChannel {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				assert.NoError(t, err)
				assert.NotZero(t, read)
				message, err := support.CreateMessageFromBytes(buf.Bytes())
				assert.NotNil(t, message)
				assert.NoError(t, err)
			}
		})
	}
	status := connector.AwaitStatus()
	suite.NotNil(status)
	suite.Equal(status.ObjectCount, status.Successful)
}

// TestContactSerializationRegression verifies ability to query contact items
// and to store contact within Collection. Downloaded contacts are run through
// a regression test to ensure that downloaded items can be uploaded.
func (suite *GraphConnectorIntegrationSuite) TestContactSerializationRegression() {
	t := suite.T()
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.ContactFolders([]string{suite.user}, selectors.Any()))
	eb, err := sel.ToExchangeBackup()
	require.NoError(t, err)
	scopes := eb.Scopes()
	connector := loadConnector(t)
	suite.Len(scopes, 1)
	contactsOnly := scopes[0]
	collections, err := connector.createCollections(context.Background(), contactsOnly)
	assert.NoError(t, err)
	number := 0
	for _, edc := range collections {
		testName := fmt.Sprintf("%s_ContactFolder_%d", edc.FullPath()[1], number)
		suite.T().Run(testName, func(t *testing.T) {
			streamChannel := edc.Items()
			for stream := range streamChannel {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				assert.NoError(t, err)
				assert.NotZero(t, read)
				contact, err := support.CreateContactFromBytes(buf.Bytes())
				assert.NotNil(t, contact)
				assert.NoError(t, err)

			}
			number++
		})
	}
	status := connector.AwaitStatus()
	suite.NotNil(status)
	suite.Equal(status.ObjectCount, status.Successful)
}

// TestEventsSerializationRegression ensures functionality of createCollections
// to be able to successfully query, download and restore event objects
func (suite *GraphConnectorIntegrationSuite) TestEventsSerializationRegression() {
	t := suite.T()
	connector := loadConnector(t)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Events([]string{suite.user}, selectors.Any()))
	scopes := sel.Scopes()
	suite.Equal(len(scopes), 1)
	collections, err := connector.createCollections(context.Background(), scopes[0])
	require.NoError(t, err)
	for _, edc := range collections {
		streamChannel := edc.Items()
		number := 0
		for stream := range streamChannel {
			testName := fmt.Sprintf("%s_Event_%d", edc.FullPath()[2], number)
			suite.T().Run(testName, func(t *testing.T) {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				assert.NoError(t, err)
				assert.NotZero(t, read)
				event, err := support.CreateEventFromBytes(buf.Bytes())
				assert.NotNil(t, event)
				assert.NoError(t, err)
			})
		}
	}
	status := connector.AwaitStatus()
	suite.NotNil(status)
	suite.Equal(status.ObjectCount, status.Successful)
}

// TestRestoreMessages uses mock data to ensure GraphConnector
// is able to restore a several messageable item to a Mailbox.
// The result should be all successful items restored within the same folder.
func (suite *GraphConnectorIntegrationSuite) TestRestoreMessages() {
	t := suite.T()
	connector := loadConnector(t)
	user := tester.M365UserID(t)
	if len(user) == 0 {
		suite.T().Skip("Environment not configured: missing m365 test user")
	}

	collection := make([]data.Collection, 0)
	for i := 0; i < 3; i++ {
		mdc := mockconnector.NewMockExchangeCollection([]string{"tenant", user, mailCategory, "Inbox"}, 1)
		collection = append(collection, mdc)
	}

	err := connector.RestoreMessages(context.Background(), collection)
	assert.NoError(suite.T(), err)
	status := connector.AwaitStatus()
	assert.NotNil(t, status)
	assert.Equal(t, status.ObjectCount, status.Successful)
	assert.Equal(t, status.FolderCount, 1)
}

// TestAccessOfInboxAllUsers verifies that GraphConnector can
// support `--all-users` for backup operations. Selector.DiscreteScopes
// returns all of the users within one scope. Only users who have
// messages in their inbox will have a collection returned.
// The final test insures that more than a 75% of the user collections are
// returned. If an error was experienced, the test will fail overall
func (suite *GraphConnectorIntegrationSuite) TestAccessOfInboxAllUsers() {
	t := suite.T()
	connector := loadConnector(t)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.MailFolders(selectors.Any(), []string{"Inbox"}))
	scopes := sel.DiscreteScopes(connector.GetUsers())
	for _, scope := range scopes {
		users := scope.Get(selectors.ExchangeUser)
		standard := (len(users) / 4) * 3
		collections, err := connector.createCollections(context.Background(), scope)
		require.NoError(t, err)
		suite.Greater(len(collections), standard)
	}
}

///------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------

//  TestCreateAndDeleteMailFolder ensures GraphConnector has the ability
//  to create and remove folders within the tenant
func (suite *GraphConnectorIntegrationSuite) TestCreateAndDeleteMailFolder() {
	now := time.Now()
	folderName := "TestFolder: " + common.FormatSimpleDateTime(now)
	aFolder, err := exchange.CreateMailFolder(&suite.connector.graphService, suite.user, folderName)
	assert.NoError(suite.T(), err, support.ConnectorStackErrorTrace(err))
	if aFolder != nil {
		err = exchange.DeleteMailFolder(suite.connector.Service(), suite.user, *aFolder.GetId())
		assert.NoError(suite.T(), err)
	}
}

// TestCreateAndDeleteContactFolder ensures GraphConnector has the ability
// to create and remove contact folders within the tenant
func (suite *GraphConnectorIntegrationSuite) TestCreateAndDeleteContactFolder() {
	now := time.Now()
	folderName := "TestContactFolder: " + common.FormatSimpleDateTime(now)
	aFolder, err := exchange.CreateContactFolder(suite.connector.Service(), suite.user, folderName)
	assert.NoError(suite.T(), err)
	if aFolder != nil {
		err = exchange.DeleteContactFolder(suite.connector.Service(), suite.user, *aFolder.GetId())
		assert.NoError(suite.T(), err)
	}
}
