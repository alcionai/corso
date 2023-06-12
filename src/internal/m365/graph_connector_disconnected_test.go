package m365

import (
	"sync"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
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

func statusTestTask(
	t *testing.T,
	gc *GraphConnector,
	objects, success, folder int,
) {
	ctx, flush := tester.NewContext(t)
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
	t := suite.T()
	gc := GraphConnector{wg: &sync.WaitGroup{}}

	// Two tasks
	gc.incrementAwaitingMessages()
	gc.incrementAwaitingMessages()

	// Each helper task processes 4 objects, 1 success, 3 errors, 1 folders
	go statusTestTask(t, &gc, 4, 1, 1)
	go statusTestTask(t, &gc, 4, 1, 1)

	stats := gc.Wait()

	assert.NotEmpty(t, gc.PrintableStatus())
	// Expect 8 objects
	assert.Equal(t, 8, stats.Objects)
	// Expect 2 success
	assert.Equal(t, 2, stats.Successes)
	// Expect 2 folders
	assert.Equal(t, 2, stats.Folders)
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
				sel.Exclude(selTD.OneDriveBackupFolderScope(sel))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Filter(selTD.OneDriveBackupFolderScope(sel))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Include(selTD.OneDriveBackupFolderScope(sel))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
		},
		{
			name:       "Invalid User",
			checkError: assert.NoError,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Exclude(selTD.OneDriveBackupFolderScope(sel))
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Filter(selTD.OneDriveBackupFolderScope(sel))
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Include(selTD.OneDriveBackupFolderScope(sel))
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
