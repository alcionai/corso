package connector

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/support"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
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

type DisconnectedGraphConnectorSuite struct {
	suite.Suite
}

func TestDisconnectedGraphSuite(t *testing.T) {
	suite.Run(t, new(DisconnectedGraphConnectorSuite))
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_setTenantUsers() {
	err := suite.connector.setTenantUsers()
	assert.NoError(suite.T(), err)
	suite.Greater(len(suite.connector.Users), 0)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_ExchangeDataCollection() {
	if err := ctesting.RunOnAny(ctesting.CorsoCITests); err != nil {
		suite.T().Skip(err)
	}
	collectionList, err := suite.connector.ExchangeDataCollection(context.Background(), "lidiah@8qzvrj.onmicrosoft.com")
	assert.NotNil(suite.T(), collectionList, "collection list")
	assert.Nil(suite.T(), err) // TODO Remove after https://github.com/alcionai/corso/issues/140
	assert.NotNil(suite.T(), suite.connector.status, "connector status")
	exchangeData := collectionList[0]
	suite.Greater(len(exchangeData.FullPath()), 2)
}

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
	ds := ExchangeData{id: "test", message: bytes}
	edc := NewExchangeDataCollection("tenant", []string{"tenantId", evs[user], mailCategory, "Inbox"})
	edc.PopulateCollection(&ds)
	edc.FinishPopulation()
	err = suite.connector.RestoreMessages(context.Background(), &edc)
	assert.NoError(suite.T(), err)
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
	var dc DataCollection
	concrete := NewExchangeDataCollection("Check", []string{"interface", "works"})
	dc = &concrete
	assert.NotNil(suite.T(), dc)

}

func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_Status() {
	gc := GraphConnector{}
	suite.Equal(len(gc.Status()), 0)
	status, err := support.CreateStatus(support.Restore, 12, 9, 8,
		support.WrapAndAppend("tres", errors.New("three"), support.WrapAndAppend("arc376", errors.New("one"), errors.New("two"))))
	assert.NoError(suite.T(), err)
	gc.SetStatus(*status)
	suite.Greater(len(gc.Status()), 0)
}
func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_ErrorChecking() {
	tests := []struct {
		name                 string
		err                  error
		returnRecoverable    bool
		returnNonRecoverable bool
	}{
		{
			name:                 "Neither Option",
			err:                  errors.New("regular error"),
			returnRecoverable:    false,
			returnNonRecoverable: false,
		},
		{
			name:                 "Validate Recoverable",
			err:                  support.SetRecoverableError(errors.New("Recoverable")),
			returnRecoverable:    true,
			returnNonRecoverable: false,
		},
		{name: "Validate NonRecoverable",
			err:                  support.SetNonRecoverableError(errors.New("Non-recoverable")),
			returnRecoverable:    false,
			returnNonRecoverable: true,
		},
		{
			name: "Wrapped Recoverable",
			err: support.SetRecoverableError(support.WrapAndAppend(
				"Wrapped Recoverable", errors.New("Recoverable"), nil)),
			returnRecoverable:    true,
			returnNonRecoverable: false,
		},
		{
			name:                 "On Nil",
			err:                  nil,
			returnRecoverable:    false,
			returnNonRecoverable: false,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			recoverable := IsRecoverableError(test.err)
			nonRecoverable := IsNonRecoverableError(test.err)
			suite.Equal(recoverable, test.returnRecoverable, "Expected: %v received %v", test.returnRecoverable, recoverable)
			suite.Equal(nonRecoverable, test.returnNonRecoverable)
		})
	}
}

func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_TaskList() {
	tasks := NewTaskList()
	tasks.AddTask("person1", "Go to store")
	tasks.AddTask("person1", "drop off mail")
	values := tasks.GetTasks("person1")
	suite.Equal(len(values), 2)
	nonValues := tasks.GetTasks("unknown")
	suite.Zero(len(nonValues))
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
			suite.T().Logf("%v", err)
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
			suite.T().Logf("%v", err)
			suite.Equal(test.isError, err != nil)
		})

	}

}
