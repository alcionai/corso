package connector

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/selectors/testdata"
)

// ---------------------------------------------------------------------------
// DataCollection tests
// ---------------------------------------------------------------------------

type DataCollectionIntgSuite struct {
	tester.Suite
	user string
	site string
}

func TestDataCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &DataCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *DataCollectionIntgSuite) SetupSuite() {
	suite.user = tester.M365UserID(suite.T())
	suite.site = tester.M365SiteID(suite.T())

	tester.LogTimeOfTest(suite.T())
}

// TestExchangeDataCollection verifies interface between operation and
// GraphConnector remains stable to receive a non-zero amount of Collections
// for the Exchange Package. Enabled exchange applications:
// - mail
// - contacts
// - events
func (suite *DataCollectionIntgSuite) TestExchangeDataCollection() {
	ctx, flush := tester.NewContext()
	defer flush()

	selUsers := []string{suite.user}

	connector := loadConnector(ctx, suite.T(), Users)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: "Email",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.user
				return sel.Selector
			},
		},
		{
			name: "Contacts",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.ContactFolders([]string{exchange.DefaultContactFolder}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.user
				return sel.Selector
			},
		},
		// {
		// 	name: "Events",
		// 	getSelector: func(t *testing.T) selectors.Selector {
		// 		sel := selectors.NewExchangeBackup(selUsers)
		// 		sel.Include(sel.EventCalendars([]string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
		// 		sel.DiscreteOwner = suite.user
		// 		return sel.Selector
		// 	},
		// },
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			sel := test.getSelector(t)

			collections, excludes, err := exchange.DataCollections(
				ctx,
				sel,
				sel,
				nil,
				connector.credentials,
				connector.UpdateStatus,
				control.Options{},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.Empty(t, excludes)

			for range collections {
				connector.incrementAwaitingMessages()
			}

			// Categories with delta endpoints will produce a collection for metadata
			// as well as the actual data pulled, and the "temp" root collection.
			assert.GreaterOrEqual(t, len(collections), 1, "expected 1 <= num collections <= 2")
			assert.GreaterOrEqual(t, 3, len(collections), "expected 1 <= num collections <= 3")

			for _, col := range collections {
				for object := range col.Items(ctx, fault.New(true)) {
					buf := &bytes.Buffer{}
					_, err := buf.ReadFrom(object.ToReader())
					assert.NoError(t, err, "received a buf.Read error", clues.ToCore(err))
				}
			}

			status := connector.Wait()
			assert.NotZero(t, status.Successes)
			t.Log(status.String())
		})
	}
}

// TestInvalidUserForDataCollections ensures verification process for users
func (suite *DataCollectionIntgSuite) TestDataCollections_invalidResourceOwner() {
	ctx, flush := tester.NewContext()
	defer flush()

	owners := []string{"snuffleupagus"}

	connector := loadConnector(ctx, suite.T(), Users)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: "invalid exchange backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(owners)
				sel.Include(sel.MailFolders(selectors.Any()))
				return sel.Selector
			},
		},
		{
			name: "Invalid onedrive backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup(owners)
				sel.Include(sel.Folders(selectors.Any()))
				return sel.Selector
			},
		},
		{
			name: "Invalid sharepoint backup site",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup(owners)
				sel.Include(testdata.SharePointBackupFolderScope(sel))
				return sel.Selector
			},
		},
		{
			name: "missing exchange backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(owners)
				sel.Include(sel.MailFolders(selectors.Any()))
				sel.DiscreteOwner = ""
				return sel.Selector
			},
		},
		{
			name: "missing onedrive backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup(owners)
				sel.Include(sel.Folders(selectors.Any()))
				sel.DiscreteOwner = ""
				return sel.Selector
			},
		},
		{
			name: "missing sharepoint backup site",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup(owners)
				sel.Include(testdata.SharePointBackupFolderScope(sel))
				sel.DiscreteOwner = ""
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			collections, excludes, err := connector.ProduceBackupCollections(
				ctx,
				test.getSelector(t),
				test.getSelector(t),
				nil,
				control.Options{},
				fault.New(true))
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, collections)
			assert.Empty(t, excludes)
		})
	}
}

// TestSharePointDataCollection verifies interface between operation and
// GraphConnector remains stable to receive a non-zero amount of Collections
// for the SharePoint Package.
func (suite *DataCollectionIntgSuite) TestSharePointDataCollection() {
	ctx, flush := tester.NewContext()
	defer flush()

	selSites := []string{suite.site}

	connector := loadConnector(ctx, suite.T(), Sites)
	tests := []struct {
		name        string
		expected    int
		getSelector func() selectors.Selector
	}{
		{
			name: "Libraries",
			getSelector: func() selectors.Selector {
				sel := selectors.NewSharePointBackup(selSites)
				sel.Include(testdata.SharePointBackupFolderScope(sel))
				return sel.Selector
			},
		},
		{
			name:     "Lists",
			expected: 0,
			getSelector: func() selectors.Selector {
				sel := selectors.NewSharePointBackup(selSites)
				sel.Include(sel.Lists(selectors.Any()))
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()
			sel := test.getSelector()

			collections, excludes, err := sharepoint.DataCollections(
				ctx,
				graph.NoTimeoutHTTPWrapper(),
				sel,
				connector.credentials,
				connector.Service,
				connector,
				control.Options{},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			// Not expecting excludes as this isn't an incremental backup.
			assert.Empty(t, excludes)

			for range collections {
				connector.incrementAwaitingMessages()
			}

			// we don't know an exact count of drives this will produce,
			// but it should be more than one.
			assert.Less(t, test.expected, len(collections))

			for _, coll := range collections {
				for object := range coll.Items(ctx, fault.New(true)) {
					buf := &bytes.Buffer{}
					_, err := buf.ReadFrom(object.ToReader())
					assert.NoError(t, err, "reading item", clues.ToCore(err))
				}
			}

			status := connector.Wait()
			assert.NotZero(t, status.Successes)
			t.Log(status.String())
		})
	}
}

// ---------------------------------------------------------------------------
// CreateSharePointCollection tests
// ---------------------------------------------------------------------------

type SPCollectionIntgSuite struct {
	tester.Suite
	connector *GraphConnector
	user      string
}

func TestSPCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &SPCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *SPCollectionIntgSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	suite.connector = loadConnector(ctx, suite.T(), Sites)
	suite.user = tester.M365UserID(suite.T())

	tester.LogTimeOfTest(suite.T())
}

func (suite *SPCollectionIntgSuite) TestCreateSharePointCollection_Libraries() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t       = suite.T()
		siteID  = tester.M365SiteID(t)
		gc      = loadConnector(ctx, t, Sites)
		siteIDs = []string{siteID}
	)

	id, name, err := gc.PopulateOwnerIDAndNamesFrom(ctx, siteID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.LibraryFolders([]string{"foo"}, selectors.PrefixMatch()))

	sel.SetDiscreteOwnerIDName(id, name)

	cols, excludes, err := gc.ProduceBackupCollections(
		ctx,
		sel.Selector,
		sel.Selector,
		nil,
		control.Options{},
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	require.Len(t, cols, 2) // 1 collection, 1 path prefix directory to ensure the root path exists.
	// No excludes yet as this isn't an incremental backup.
	assert.Empty(t, excludes)

	t.Logf("cols[0] Path: %s\n", cols[0].FullPath().String())
	assert.Equal(
		t,
		path.SharePointMetadataService.String(),
		cols[0].FullPath().Service().String())

	t.Logf("cols[1] Path: %s\n", cols[1].FullPath().String())
	assert.Equal(
		t,
		path.SharePointService.String(),
		cols[1].FullPath().Service().String())
}

func (suite *SPCollectionIntgSuite) TestCreateSharePointCollection_Lists() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t       = suite.T()
		siteID  = tester.M365SiteID(t)
		gc      = loadConnector(ctx, t, Sites)
		siteIDs = []string{siteID}
	)

	id, name, err := gc.PopulateOwnerIDAndNamesFrom(ctx, siteID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.Lists(selectors.Any()))

	sel.SetDiscreteOwnerIDName(id, name)

	cols, excludes, err := gc.ProduceBackupCollections(
		ctx,
		sel.Selector,
		sel.Selector,
		nil,
		control.Options{},
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.Less(t, 0, len(cols))
	// No excludes yet as this isn't an incremental backup.
	assert.Empty(t, excludes)

	for _, collection := range cols {
		t.Logf("Path: %s\n", collection.FullPath().String())

		for item := range collection.Items(ctx, fault.New(true)) {
			t.Log("File: " + item.UUID())

			bs, err := io.ReadAll(item.ToReader())
			require.NoError(t, err, clues.ToCore(err))
			t.Log(string(bs))
		}
	}
}
