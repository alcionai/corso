package connector

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
)

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
	var dc DataCollection
	concrete := NewExchangeDataCollection("Check", []string{"interface", "works"})
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
		})
	}
}

func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_TaskList() {
	tasks := NewTaskList()
	tasks.AddTask("person1", "Go to store")
	tasks.AddTask("person1", "drop off mail")
	values := tasks["person1"]
	suite.Equal(len(values), 2)
	nonValues := tasks["unknown"]
	suite.Empty(nonValues)
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
