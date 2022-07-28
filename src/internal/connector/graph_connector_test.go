package connector

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
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

//TestGraphConnector_restoreMessages uses mock data to ensure GraphConnector
// is able to restore a messageable item to a Mailbox.
func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_restoreMessages() {
	user := "TEST_GRAPH_USER" // user.GetId()
	evs, err := ctesting.GetRequiredEnvVars(user)
	if err != nil {
		suite.T().Skipf("Environment not configured: %v\n", err)
	}
	mdc := mockconnector.NewMockExchangeCollectionMail([]string{"tenant", evs[user], mailCategory, "Inbox"}, 1)
	err = suite.connector.RestoreMessages(context.Background(), []data.Collection{mdc})
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

// ---------------------------------------------------------------
// Disconnected Test Section
// -------------------------
type DisconnectedGraphConnectorSuite struct {
	suite.Suite
}

func TestDisconnectedGraphSuite(t *testing.T) {
	suite.Run(t, new(DisconnectedGraphConnectorSuite))
}

func (suite *DisconnectedGraphConnectorSuite) TestBadConnection() {

	table := []struct {
		name string
		acct func(t *testing.T) account.Account
	}{
		{
			name: "Invalid Credentials",
			acct: func(t *testing.T) account.Account {
				a, err := account.NewAccount(
					account.ProviderM365,
					account.M365Config{
						M365: credentials.M365{
							ClientID:     "Test",
							ClientSecret: "without",
						},
						TenantID: "data",
					},
				)
				require.NoError(t, err)
				return a
			},
		},
		{
			name: "Empty Credentials",
			acct: func(t *testing.T) account.Account {
				// intentionally swallowing the error here
				a, _ := account.NewAccount(account.ProviderM365)
				return a
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			gc, err := NewGraphConnector(test.acct(t))
			assert.Nil(t, gc, test.name+" failed")
			assert.NotNil(t, err, test.name+"failed")
		})
	}
}

func (suite *DisconnectedGraphConnectorSuite) TestBuild() {
	names := make(map[string]string)
	names["Al"] = "Bundy"
	names["Ellen"] = "Ripley"
	names["Axel"] = "Foley"
	first := buildFromMap(true, names)
	last := buildFromMap(false, names)
	suite.True(Contains(first, "Al"))
	suite.True(Contains(first, "Ellen"))
	suite.True(Contains(first, "Axel"))
	suite.True(Contains(last, "Bundy"))
	suite.True(Contains(last, "Ripley"))
	suite.True(Contains(last, "Foley"))

}

func (suite *DisconnectedGraphConnectorSuite) TestInterfaceAlignment() {
	var dc data.Collection
	concrete := exchange.NewCollection("Check", []string{"interface", "works"})
	dc = &concrete
	assert.NotNil(suite.T(), dc)

}

func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_Status() {
	gc := GraphConnector{}
	suite.Equal(len(gc.PrintableStatus()), 0)
	status := support.CreateStatus(
		context.Background(),
		support.Restore,
		12, 9, 8,
		support.WrapAndAppend("tres", errors.New("three"), support.WrapAndAppend("arc376", errors.New("one"), errors.New("two"))))
	gc.SetStatus(*status)
	suite.Greater(len(gc.PrintableStatus()), 0)
	suite.Greater(gc.Status().ObjectCount, 0)
}
func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_ErrorChecking() {
	tests := []struct {
		name                 string
		err                  error
		returnRecoverable    assert.BoolAssertionFunc
		returnNonRecoverable assert.BoolAssertionFunc
	}{
		{
			name:                 "Neither Option",
			err:                  errors.New("regular error"),
			returnRecoverable:    assert.False,
			returnNonRecoverable: assert.False,
		},
		{
			name:                 "Validate Recoverable",
			err:                  support.SetRecoverableError(errors.New("Recoverable")),
			returnRecoverable:    assert.True,
			returnNonRecoverable: assert.False,
		},
		{name: "Validate NonRecoverable",
			err:                  support.SetNonRecoverableError(errors.New("Non-recoverable")),
			returnRecoverable:    assert.False,
			returnNonRecoverable: assert.True,
		},
		{
			name: "Wrapped Recoverable",
			err: support.WrapAndAppend(
				"Wrapped Recoverable",
				support.SetRecoverableError(errors.New("Recoverable")),
				nil),
			returnRecoverable:    assert.True,
			returnNonRecoverable: assert.False,
		},
		{
			name:                 "On Nil",
			err:                  nil,
			returnRecoverable:    assert.False,
			returnNonRecoverable: assert.False,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			recoverable := IsRecoverableError(test.err)
			nonRecoverable := IsNonRecoverableError(test.err)
			test.returnRecoverable(suite.T(), recoverable, "Test: %s Recoverable-received %v", test.name, recoverable)
			test.returnNonRecoverable(suite.T(), nonRecoverable, "Test: %s non-recoverable: %v", test.name, nonRecoverable)
			t.Logf("Is nil: %v", test.err == nil)
		})
	}
}

func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_TestOptionsForMailFolders() {
	tests := []struct {
		name    string
		params  []string
		isError bool
	}{
		{
			name:    "Accepted",
			params:  []string{"displayName"},
			isError: false,
		},
		{
			name:    "Multiple Accepted",
			params:  []string{"displayName", "parentFolderId"},
			isError: false,
		},
		{
			name:    "Incorrect param",
			params:  []string{"status"},
			isError: true,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := optionsForMailFolders(test.params)
			suite.Equal(test.isError, err != nil)
		})

	}

}

func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_TestOptionsForMessages() {
	tests := []struct {
		name    string
		params  []string
		isError bool
	}{
		{
			name:    "Accepted",
			params:  []string{"subject"},
			isError: false,
		},
		{
			name:    "Multiple Accepted",
			params:  []string{"webLink", "parentFolderId"},
			isError: false,
		},
		{
			name:    "Incorrect param",
			params:  []string{"status"},
			isError: true,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := optionsForMessages(test.params)
			suite.Equal(test.isError, err != nil)
		})
	}
}
