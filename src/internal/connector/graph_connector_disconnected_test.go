package connector

import (
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
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------
// Disconnected Test Section
// ---------------------------------------------------------------
type DisconnectedGraphConnectorSuite struct {
	suite.Suite
}

func TestDisconnectedGraphSuite(t *testing.T) {
	tester.LogTimeOfTest(t)
	suite.Run(t, new(DisconnectedGraphConnectorSuite))
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
			gc, err := NewGraphConnector(ctx, test.acct(t), Users)
			assert.Nil(t, gc, test.name+" failed")
			assert.NotNil(t, err, test.name+"failed")
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
			Objects:    objects,
			Successes:  success,
			TotalBytes: 0,
		},
		support.WrapAndAppend(
			"tres",
			errors.New("three"),
			support.WrapAndAppend("arc376", errors.New("one"), errors.New("two")),
		),
		"statusTestTask",
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

func (suite *DisconnectedGraphConnectorSuite) TestRestoreFailsBadService() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		dest = tester.DefaultTestRestoreDestination()
		gc   = GraphConnector{wg: &sync.WaitGroup{}}
		sel  = selectors.Selector{
			Service: selectors.ServiceUnknown,
		}
	)

	deets, err := gc.RestoreDataCollections(ctx, acct, sel, dest, nil)
	assert.Error(t, err)
	assert.NotNil(t, deets)

	status := gc.AwaitStatus()
	assert.Equal(t, 0, status.ObjectCount)
	assert.Equal(t, 0, status.FolderCount)
	assert.Equal(t, 0, status.Successful)
}

func (suite *DisconnectedGraphConnectorSuite) TestVerifyBackupInputs() {
	users := []string{
		"elliotReid@someHospital.org",
		"chrisTurk@someHospital.org",
		"carlaEspinosa@someHospital.org",
		"bobKelso@someHospital.org",
		"johnDorian@someHospital.org",
	}

	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
		checkError  assert.ErrorAssertionFunc
	}{
		{
			name:       "No scopes",
			checkError: assert.Error,
			getSelector: func(t *testing.T) selectors.Selector {
				return selectors.NewExchangeBackup(nil).Selector
			},
		},
		{
			name:       "Valid Single User",
			checkError: assert.NoError,
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup([]string{"bobkelso@someHospital.org"})
				sel.Include(sel.MailFolders([]string{"bobkelso@someHospital.org"}, selectors.Any()))
				return sel.Selector
			},
		},
		{
			name:       "Partial invalid user",
			checkError: assert.Error,
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup([]string{"bobkelso@someHospital.org", "janitor@someHospital.org"})
				sel.Include(sel.MailFolders([]string{"bobkelso@someHospital.org", "janitor@someHospital.org"}, selectors.Any()))
				return sel.Selector
			},
		},
		{
			name:       "Multiple Valid Users",
			checkError: assert.NoError,
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup(
					[]string{"elliotReid@someHospital.org", "johnDorian@someHospital.org", "christurk@somehospital.org"},
				)
				sel.Include(
					sel.Users([]string{"elliotReid@someHospital.org", "johnDorian@someHospital.org", "christurk@somehospital.org"}))

				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			err := verifyBackupInputs(test.getSelector(t), users, nil)
			test.checkError(t, err)
		})
	}
}

func (suite *DisconnectedGraphConnectorSuite) TestVerifyBackupInputs_allServices() {
	users := []string{"elliotReid@someHospital.org"}
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
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org"})
				sel.Exclude(sel.Folders([]string{"elliotReid@someHospital.org"}, selectors.Any()))
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org"})
				sel.Filter(sel.Folders([]string{"elliotReid@someHospital.org"}, selectors.Any()))
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org"})
				sel.Include(sel.Folders([]string{"elliotReid@someHospital.org"}, selectors.Any()))
				return sel.Selector
			},
		},
		{
			name:       "Invalid User",
			checkError: assert.Error,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Exclude(sel.Folders([]string{"foo@SomeCompany.org"}, selectors.Any()))
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Filter(sel.Folders([]string{"foo@SomeCompany.org"}, selectors.Any()))
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Include(sel.Folders([]string{"foo@SomeCompany.org"}, selectors.Any()))
				return sel.Selector
			},
		},
		{
			name:       "valid sites",
			checkError: assert.NoError,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.Exclude(sel.Sites([]string{"abc.site.foo", "bar.site.baz"}))
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.Filter(sel.Sites([]string{"abc.site.foo", "bar.site.baz"}))
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.Include(sel.Sites([]string{"abc.site.foo", "bar.site.baz"}))
				return sel.Selector
			},
		},
		{
			name:       "invalid sites",
			checkError: assert.Error,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Exclude(sel.Sites([]string{"fnords.smarfs.brawnhilda"}))
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Filter(sel.Sites([]string{"fnords.smarfs.brawnhilda"}))
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Include(sel.Sites([]string{"fnords.smarfs.brawnhilda"}))
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			err := verifyBackupInputs(test.excludes(t), users, sites)
			test.checkError(t, err)
			err = verifyBackupInputs(test.filters(t), users, sites)
			test.checkError(t, err)
			err = verifyBackupInputs(test.includes(t), users, sites)
			test.checkError(t, err)
		})
	}
}
