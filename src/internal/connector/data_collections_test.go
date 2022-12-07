package connector

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// DataCollection tests
// ---------------------------------------------------------------------------

type ConnectorDataCollectionIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
	user      string
	site      string
}

func TestConnectorDataCollectionIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoConnectorDataCollectionTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ConnectorDataCollectionIntegrationSuite))
}

func (suite *ConnectorDataCollectionIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)
	suite.connector = loadConnector(ctx, suite.T(), AllResources)
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
func (suite *ConnectorDataCollectionIntegrationSuite) TestExchangeDataCollection() {
	ctx, flush := tester.NewContext()
	defer flush()

	selUsers := []string{suite.user}

	connector := loadConnector(ctx, suite.T(), Users)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: suite.user + " Email",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.MailFolders(selUsers, []string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))

				return sel.Selector
			},
		},
		{
			name: suite.user + " Contacts",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.ContactFolders(
					selUsers,
					[]string{exchange.DefaultContactFolder},
					selectors.PrefixMatch()))

				return sel.Selector
			},
		},
		// {
		// 	name: suite.user + " Events",
		// 	getSelector: func(t *testing.T) selectors.Selector {
		// 		sel := selectors.NewExchangeBackup(selUsers)
		// 		sel.Include(sel.EventCalendars(
		// 			selUsers,
		// 			[]string{exchange.DefaultCalendar},
		// 			selectors.PrefixMatch(),
		// 		))

		// 		return sel.Selector
		// 	},
		// },
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collections, err := exchange.DataCollections(
				ctx,
				test.getSelector(t),
				nil,
				[]string{suite.user},
				connector.credentials,
				connector.UpdateStatus,
				control.Options{})
			require.NoError(t, err)

			for range collections {
				connector.incrementAwaitingMessages()
			}

			// Categories with delta endpoints will produce a collection for metadata
			// as well as the actual data pulled.
			assert.GreaterOrEqual(t, len(collections), 1, "expected 1 <= num collections <= 2")
			assert.GreaterOrEqual(t, 2, len(collections), "expected 1 <= num collections <= 2")

			for _, col := range collections {
				for object := range col.Items() {
					buf := &bytes.Buffer{}
					_, err := buf.ReadFrom(object.ToReader())
					assert.NoError(t, err, "received a buf.Read error")
				}
			}

			status := connector.AwaitStatus()
			assert.NotZero(t, status.Successful)
			t.Log(status.String())
		})
	}
}

// TestInvalidUserForDataCollections ensures verification process for users
func (suite *ConnectorDataCollectionIntegrationSuite) TestInvalidUserForDataCollections() {
	ctx, flush := tester.NewContext()
	defer flush()

	invalidUser := "foo@example.com"
	selUsers := []string{invalidUser}

	connector := loadConnector(ctx, suite.T(), Users)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: "invalid exchange backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.MailFolders(selUsers, selectors.Any()))
				return sel.Selector
			},
		},
		{
			name: "Invalid onedrive backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup(selUsers)
				sel.Include(sel.Folders(selUsers, selectors.Any()))
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collections, err := connector.DataCollections(ctx, test.getSelector(t), nil, control.Options{})
			assert.Error(t, err)
			assert.Empty(t, collections)
		})
	}
}

// TestSharePointDataCollection verifies interface between operation and
// GraphConnector remains stable to receive a non-zero amount of Collections
// for the SharePoint Package.
func (suite *ConnectorDataCollectionIntegrationSuite) TestSharePointDataCollection() {
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
				sel.Include(sel.Libraries(selSites, selectors.Any()))
				sel.DiscreteOwner = suite.site
				return sel.Selector
			},
		},
		{
			name:     "Lists",
			expected: 0,
			getSelector: func() selectors.Selector {
				sel := selectors.NewSharePointBackup(selSites)
				sel.Include(sel.Lists(selSites, selectors.Any()))
				sel.DiscreteOwner = suite.site
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collections, err := sharepoint.DataCollections(
				ctx,
				test.getSelector(),
				selSites,
				connector.credentials.AzureTenantID,
				connector.Service,
				connector,
				control.Options{})
			require.NoError(t, err)

			for range collections {
				connector.incrementAwaitingMessages()
			}

			// we don't know an exact count of drives this will produce,
			// but it should be more than one.
			assert.Less(t, test.expected, len(collections))

			for _, coll := range collections {
				for object := range coll.Items() {
					buf := &bytes.Buffer{}
					_, err := buf.ReadFrom(object.ToReader())
					assert.NoError(t, err, "reading item")
				}
			}

			status := connector.AwaitStatus()
			assert.NotZero(t, status.Successful)
			t.Log(status.String())
		})
	}
}

// ---------------------------------------------------------------------------
// CreateSharePointCollection tests
// ---------------------------------------------------------------------------

type ConnectorCreateSharePointCollectionIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
	user      string
}

func TestConnectorCreateSharePointCollectionIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoConnectorCreateSharePointCollectionTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ConnectorCreateSharePointCollectionIntegrationSuite))
}

func (suite *ConnectorCreateSharePointCollectionIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)
	suite.connector = loadConnector(ctx, suite.T(), Sites)
	suite.user = tester.M365UserID(suite.T())
	tester.LogTimeOfTest(suite.T())
}

// TestCreateSharePointCollection. Ensures the proper amount of collections are created based
// on the selector.
func (suite *ConnectorCreateSharePointCollectionIntegrationSuite) TestCreateSharePointCollection() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t       = suite.T()
		siteID  = tester.M365SiteID(t)
		gc      = loadConnector(ctx, t, Sites)
		siteIDs = []string{siteID}
	)

	tables := []struct {
		name       string
		sel        func() selectors.Selector
		comparator assert.ComparisonAssertionFunc
	}{
		{
			name:       "SharePoint.Libraries",
			comparator: assert.Equal,
			sel: func() selectors.Selector {
				sel := selectors.NewSharePointBackup(siteIDs)
				sel.Include(sel.Libraries(
					siteIDs,
					[]string{"foo"},
					selectors.PrefixMatch(),
				))

				return sel.Selector
			},
		},
		{
			name:       "SharePoint.Lists",
			comparator: assert.Less,
			sel: func() selectors.Selector {
				sel := selectors.NewSharePointBackup(siteIDs)
				sel.Include(sel.Lists(
					siteIDs,
					selectors.Any(),
					selectors.PrefixMatch(), // without this option a SEG Fault occurs
				))

				return sel.Selector
			},
		},
	}

	for _, test := range tables {
		t.Run(test.name, func(t *testing.T) {
			cols, err := gc.DataCollections(ctx, test.sel(), nil, control.Options{})
			require.NoError(t, err)
			test.comparator(t, 0, len(cols))

			if test.name == "SharePoint.Lists" {
				for _, collection := range cols {
					t.Logf("Path: %s\n", collection.FullPath().String())
					for item := range collection.Items() {
						t.Log("File: " + item.UUID())

						bytes, err := io.ReadAll(item.ToReader())
						require.NoError(t, err)
						t.Log(string(bytes))

					}
				}
			}
		})
	}
}
