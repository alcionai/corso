package connector

import (
	"sync"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------
// Disconnected Test Section
// ---------------------------------------------------------------
type DisconnectedGraphConnectorSuite struct {
	tester.Suite
}

func TestDisconnectedGraphSuite(t *testing.T) {
	s := &DisconnectedGraphConnectorSuite{
		Suite: tester.NewUnitSuite(t),
	}

	suite.Run(t, s)
}

func (suite *DisconnectedGraphConnectorSuite) TestBadConnection() {
	ctx, flush := tester.NewContext()
	defer flush()

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
							AzureClientID:     "Test",
							AzureClientSecret: "without",
						},
						AzureTenantID: "data",
					},
				)
				require.NoError(t, err, clues.ToCore(err))
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
		suite.Run(test.name, func() {
			t := suite.T()

			gc, err := NewGraphConnector(
				ctx,
				graph.HTTPClient(graph.NoTimeout()),
				test.acct(t),
				Sites,
				fault.New(true))
			assert.Nil(t, gc, test.name+" failed")
			assert.NotNil(t, err, test.name+" failed")
		})
	}
}

func statusTestTask(gc *GraphConnector, objects, success, folder int) {
	ctx, flush := tester.NewContext()
	defer flush()

	status := support.CreateStatus(
		ctx,
		support.Restore, folder,
		support.CollectionMetrics{
			Objects:   objects,
			Successes: success,
			Bytes:     0,
		},
		"statusTestTask")
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

	status := gc.Wait()
	t := suite.T()

	assert.NotEmpty(t, gc.PrintableStatus())
	// Expect 8 objects
	assert.Equal(t, 8, status.Metrics.Objects)
	// Expect 2 success
	assert.Equal(t, 2, status.Metrics.Successes)
	// Expect 2 folders
	assert.Equal(t, 2, status.Folders)
}

func (suite *DisconnectedGraphConnectorSuite) TestVerifyBackupInputs_allServices() {
	sites := []string{"abc.site.foo", "bar.site.baz"}

	tests := []struct {
		name       string
		excludes   func(t *testing.T) selectors.Selector
		filters    func(t *testing.T) selectors.Selector
		includes   func(t *testing.T) selectors.Selector
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Valid User",
			checkError: assert.NoError,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Exclude(sel.Folders(selectors.Any()))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Filter(sel.Folders(selectors.Any()))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Include(sel.Folders(selectors.Any()))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
		},
		{
			name:       "Invalid User",
			checkError: assert.NoError,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Exclude(sel.Folders(selectors.Any()))
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Filter(sel.Folders(selectors.Any()))
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Include(sel.Folders(selectors.Any()))
				return sel.Selector
			},
		},
		{
			name:       "valid sites",
			checkError: assert.NoError,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.DiscreteOwner = "abc.site.foo"
				sel.Exclude(sel.AllData())
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.DiscreteOwner = "abc.site.foo"
				sel.Filter(sel.AllData())
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.DiscreteOwner = "abc.site.foo"
				sel.Include(sel.AllData())
				return sel.Selector
			},
		},
		{
			name:       "invalid sites",
			checkError: assert.Error,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Exclude(sel.AllData())
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Filter(sel.AllData())
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Include(sel.AllData())
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := verifyBackupInputs(test.excludes(t), sites)
			test.checkError(t, err, clues.ToCore(err))
			err = verifyBackupInputs(test.filters(t), sites)
			test.checkError(t, err, clues.ToCore(err))
			err = verifyBackupInputs(test.includes(t), sites)
			test.checkError(t, err, clues.ToCore(err))
		})
	}
}
