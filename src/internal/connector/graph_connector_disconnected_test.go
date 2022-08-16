package connector

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
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
	suite.Contains(first, "Al")
	suite.Contains(first, "Ellen")
	suite.Contains(first, "Axel")
	suite.Contains(last, "Bundy")
	suite.Contains(last, "Ripley")
	suite.Contains(last, "Foley")
}

func (suite *DisconnectedGraphConnectorSuite) TestInterfaceAlignment() {
	var dc data.Collection
	concrete := mockconnector.NewMockExchangeCollection([]string{"a", "path"}, 1)
	dc = concrete
	assert.NotNil(suite.T(), dc)
}

// TestStatus to verify status pipeline connectivity
func (suite *DisconnectedGraphConnectorSuite) TestStatus() {
	gc := GraphConnector{
		statusCh: make(chan *support.ConnectorOperationStatus),
	}
	suite.Equal(len(gc.PrintableStatus()), 0)
	gc.incrementAwaitingMessages()
	go func() {
		status := support.CreateStatus(
			context.Background(),
			support.Restore,
			12, 9, 8,
			support.WrapAndAppend(
				"tres",
				errors.New("three"),
				support.WrapAndAppend("arc376", errors.New("one"), errors.New("two")),
			),
		)
		gc.statusCh <- status
	}()
	temp := <-gc.statusCh
	gc.status = temp
	atomic.AddInt32(&gc.awaitingMessages, -1)
	gc.AwaitStatus()
	suite.Greater(len(gc.PrintableStatus()), 0)
	suite.Greater(gc.Status().ObjectCount, 0)
}

func (suite *DisconnectedGraphConnectorSuite) TestErrorChecking() {
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

// TestLaunchAsyncStatus verifes status updates are populated asynchronously
// and ensures that when status switches from backup to restore that the
// status reflets the change. NOTE: time.Sleep required to allow
// async status update to perform operations.
func (suite *DisconnectedGraphConnectorSuite) TestLaunchAsyncStatus() {
	gc := GraphConnector{
		statusCh: make(chan *support.ConnectorOperationStatus),
	}
	var wg sync.WaitGroup

	testStatusCreate := func(
		objects, success, folders int,
		operation support.Operation,
		statusChannel chan *support.ConnectorOperationStatus,
		wg *sync.WaitGroup,
	) {
		status := support.CreateStatus(
			context.Background(),
			operation,
			objects,
			success,
			folders,
			nil,
		)
		statusChannel <- status
		wg.Done()
	}
	suite.Equal(len(gc.PrintableStatus()), 0)
	// Launches async process for status update
	wg.Add(1)
	go gc.LaunchAsyncStatusUpdate()
	go testStatusCreate(5, 5, 1, support.Backup, gc.statusCh, &wg)
	wg.Wait()
	time.Sleep(1 * time.Second)
	expected := 5

	suite.Equal(gc.status.Successful, expected)
	// Sending 3 more statuses
	additional := 3
	wg.Add(additional)
	for i := 0; i < additional; i++ {
		go testStatusCreate(5, 5, 1, support.Backup, gc.statusCh, &wg)
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	suite.Equal(gc.status.Successful, expected*4)
	// Switch from Backup to Restore status
	wg.Add(1)
	go testStatusCreate(2, 2, 1, support.Restore, gc.statusCh, &wg)
	wg.Wait()
	time.Sleep(1 * time.Second)
	suite.Equal(gc.status.LastOperation, support.Restore)
	suite.Equal(gc.status.Successful, 2)
}
