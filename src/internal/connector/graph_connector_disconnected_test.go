package connector

import (
	"context"
	"sync"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
)

// ---------------------------------------------------------------
// Disconnected Test Section
// -------------------------
type DisconnectedGraphConnectorSuite struct {
	suite.Suite
}

func TestDisconnectedGraphSuite(t *testing.T) {
	tester.LogTimeOfTest(t)
	suite.Run(t, new(DisconnectedGraphConnectorSuite))
}

func (suite *DisconnectedGraphConnectorSuite) TestBadConnection() {
	ctx := context.Background()
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
			gc, err := NewGraphConnector(ctx, test.acct(t))
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

	suite.Contains(first, "Al")
	suite.Contains(first, "Ellen")
	suite.Contains(first, "Axel")
	suite.Contains(last, "Bundy")
	suite.Contains(last, "Ripley")
	suite.Contains(last, "Foley")
}

func statusTestTask(gc *GraphConnector, objects, success, folder int) {
	status := support.CreateStatus(
		context.Background(),
		support.Restore,
		objects, success, folder,
		support.WrapAndAppend(
			"tres",
			errors.New("three"),
			support.WrapAndAppend("arc376", errors.New("one"), errors.New("two")),
		),
	)
	gc.UpdateStatus(status)
}

func (suite *DisconnectedGraphConnectorSuite) TestGraphConnector_Status() {
	gc := GraphConnector{wg: &sync.WaitGroup{}}

	// Two tasks
	gc.incrementAwaitingMessages()
	gc.incrementAwaitingMessages()

	// Each helper task processes 4 objects, 1 success, 3 errors, 1 folders
	go statusTestTask(&gc, 4, 1, 1)
	go statusTestTask(&gc, 4, 1, 1)

	gc.AwaitStatus()
	suite.NotEmpty(gc.PrintableStatus())
	// Expect 8 objects
	suite.Equal(8, gc.Status().ObjectCount)
	// Expect 2 success
	suite.Equal(2, gc.Status().Successful)
	// Expect 2 folders
	suite.Equal(2, gc.Status().FolderCount)
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
		{
			name:                 "Validate NonRecoverable",
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
