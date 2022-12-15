package connector

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
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

	connector := loadConnector(ctx, suite.T(), Users)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: suite.user + " Email",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.MailFolders([]string{suite.user}, []string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))

				return sel.Selector
			},
		},
		{
			name: suite.user + " Contacts",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.ContactFolders(
					[]string{suite.user},
					[]string{exchange.DefaultContactFolder},
					selectors.PrefixMatch()))

				return sel.Selector
			},
		},
		// {
		// 	name: suite.user + " Events",
		// 	getSelector: func(t *testing.T) selectors.Selector {
		// 		sel := selectors.NewExchangeBackup()
		// 		sel.Include(sel.EventCalendars(
		// 			[]string{suite.user},
		// 			[]string{exchange.DefaultCalendar},
		// 			selectors.PrefixMatch(),
		// 		))

		// 		return sel.Selector
		// 	},
		// },
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collection, err := connector.ExchangeDataCollection(ctx, test.getSelector(t), nil, control.Options{})
			require.NoError(t, err)
			// Categories with delta endpoints will produce a collection for metadata
			// as well as the actual data pulled.
			assert.GreaterOrEqual(t, len(collection), 1, "expected 1 <= num collections <= 2")
			assert.GreaterOrEqual(t, 2, len(collection), "expected 1 <= num collections <= 2")

			for _, col := range collection {
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
	connector := loadConnector(ctx, suite.T(), Users)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: "invalid exchange backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.MailFolders([]string{invalidUser}, selectors.Any()))
				return sel.Selector
			},
		},
		{
			name: "Invalid onedrive backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup()
				sel.Include(sel.Folders([]string{invalidUser}, selectors.Any()))
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

	connector := loadConnector(ctx, suite.T(), Sites)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: "Libraries",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup()
				sel.Include(sel.Libraries([]string{suite.site}, selectors.Any()))

				return sel.Selector
			},
		},
		{
			name: "Lists",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup()
				sel.Include(sel.Sites([]string{suite.site}))

				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collection, err := sharepoint.DataCollections(
				ctx,
				test.getSelector(t),
				[]string{suite.site},
				connector.credentials.AzureTenantID,
				connector.Service,
				connector,
				control.Options{})
			require.NoError(t, err)

			// we don't know an exact count of drives this will produce,
			// but it should be more than one.
			assert.Less(t, 1, len(collection))

			// the test only reads the firstt collection
			connector.incrementAwaitingMessages()

			for object := range collection[0].Items() {
				buf := &bytes.Buffer{}
				_, err := buf.ReadFrom(object.ToReader())
				assert.NoError(t, err, "received a buf.Read error")
			}

			status := connector.AwaitStatus()
			assert.NotZero(t, status.Successful)

			t.Log(status.String())
		})
	}
}

// ---------------------------------------------------------------------------
// CreateExchangeCollection tests
// ---------------------------------------------------------------------------

type ConnectorCreateExchangeCollectionIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
	user      string
	site      string
}

func TestConnectorCreateExchangeCollectionIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoConnectorCreateExchangeCollectionTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ConnectorCreateExchangeCollectionIntegrationSuite))
}

func (suite *ConnectorCreateExchangeCollectionIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)
	suite.connector = loadConnector(ctx, suite.T(), Users)
	suite.user = tester.M365UserID(suite.T())
	suite.site = tester.M365SiteID(suite.T())
	tester.LogTimeOfTest(suite.T())
}

func (suite *ConnectorCreateExchangeCollectionIntegrationSuite) TestMailFetch() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t      = suite.T()
		userID = tester.M365UserID(t)
	)

	tests := []struct {
		name        string
		scope       selectors.ExchangeScope
		folderNames map[string]struct{}
	}{
		{
			name: "Folder Iterative Check Mail",
			scope: selectors.NewExchangeBackup().MailFolders(
				[]string{userID},
				[]string{exchange.DefaultMailFolder},
				selectors.PrefixMatch(),
			)[0],
			folderNames: map[string]struct{}{
				exchange.DefaultMailFolder: {},
			},
		},
	}

	gc := loadConnector(ctx, t, Users)

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collections, err := gc.createExchangeCollections(ctx, test.scope, nil, control.Options{})
			require.NoError(t, err)

			for _, c := range collections {
				if c.FullPath().Service() == path.ExchangeMetadataService {
					continue
				}

				require.NotEmpty(t, c.FullPath().Folder())
				folder := c.FullPath().Folder()

				delete(test.folderNames, folder)
			}

			assert.Empty(t, test.folderNames)
		})
	}
}

func (suite *ConnectorCreateExchangeCollectionIntegrationSuite) TestDelta() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		userID = tester.M365UserID(suite.T())
		gc     = loadConnector(ctx, suite.T(), Users)
	)

	tests := []struct {
		name  string
		scope selectors.ExchangeScope
	}{
		{
			name: "Mail",
			scope: selectors.NewExchangeBackup().MailFolders(
				[]string{userID},
				[]string{exchange.DefaultMailFolder},
				selectors.PrefixMatch(),
			)[0],
		},
		{
			name: "Contacts",
			scope: selectors.NewExchangeBackup().ContactFolders(
				[]string{userID},
				[]string{exchange.DefaultContactFolder},
				selectors.PrefixMatch(),
			)[0],
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			// get collections without providing any delta history (ie: full backup)
			collections, err := gc.createExchangeCollections(ctx, test.scope, nil, control.Options{})
			require.NoError(t, err)
			assert.Less(t, 1, len(collections), "retrieved metadata and data collections")

			var metadata data.Collection

			for _, coll := range collections {
				if coll.FullPath().Service() == path.ExchangeMetadataService {
					metadata = coll
				}
			}

			require.NotNil(t, metadata, "collections contains a metadata collection")

			_, deltas, err := exchange.ParseMetadataCollections(ctx, []data.Collection{metadata})
			require.NoError(t, err)

			// now do another backup with the previous delta tokens,
			// which should only contain the difference.
			collections, err = gc.createExchangeCollections(ctx, test.scope, deltas, control.Options{})
			require.NoError(t, err)

			// TODO(keepers): this isn't a very useful test at the moment.  It needs to
			// investigate the items in the original and delta collections to at least
			// assert some minimum assumptions, such as "deltas should retrieve fewer items".
			// Delta usage is commented out at the moment, anyway.  So this is currently
			// a sanity check that the minimum behavior won't break.
			for _, coll := range collections {
				if coll.FullPath().Service() != path.ExchangeMetadataService {
					ec, ok := coll.(*exchange.Collection)
					require.True(t, ok, "collection is *exchange.Collection")
					assert.NotNil(t, ec)
				}
			}
		})
	}
}

// TestMailSerializationRegression verifies that all mail data stored in the
// test account can be successfully downloaded into bytes and restored into
// M365 mail objects
func (suite *ConnectorCreateExchangeCollectionIntegrationSuite) TestMailSerializationRegression() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	connector := loadConnector(ctx, t, Users)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.MailFolders([]string{suite.user}, []string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))
	collection, err := connector.createExchangeCollections(ctx, sel.Scopes()[0], nil, control.Options{})
	require.NoError(t, err)

	for _, edc := range collection {
		suite.T().Run(edc.FullPath().String(), func(t *testing.T) {
			streamChannel := edc.Items()
			// Verify that each message can be restored
			for stream := range streamChannel {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				assert.NoError(t, err)
				assert.NotZero(t, read)
				message, err := support.CreateMessageFromBytes(buf.Bytes())
				assert.NotNil(t, message)
				assert.NoError(t, err)
			}
		})
	}

	status := connector.AwaitStatus()
	suite.NotNil(status)
	suite.Equal(status.ObjectCount, status.Successful)
}

// TestContactSerializationRegression verifies ability to query contact items
// and to store contact within Collection. Downloaded contacts are run through
// a regression test to ensure that downloaded items can be uploaded.
func (suite *ConnectorCreateExchangeCollectionIntegrationSuite) TestContactSerializationRegression() {
	ctx, flush := tester.NewContext()
	defer flush()

	connector := loadConnector(ctx, suite.T(), Users)

	tests := []struct {
		name          string
		getCollection func(t *testing.T) []data.Collection
	}{
		{
			name: "Default Contact Folder",
			getCollection: func(t *testing.T) []data.Collection {
				scope := selectors.
					NewExchangeBackup().
					ContactFolders([]string{suite.user}, []string{exchange.DefaultContactFolder}, selectors.PrefixMatch())[0]
				collections, err := connector.createExchangeCollections(ctx, scope, nil, control.Options{})
				require.NoError(t, err)

				return collections
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			edcs := test.getCollection(t)
			require.GreaterOrEqual(t, len(edcs), 1, "expected 1 <= num collections <= 2")
			require.GreaterOrEqual(t, 2, len(edcs), "expected 1 <= num collections <= 2")

			for _, edc := range edcs {
				isMetadata := edc.FullPath().Service() == path.ExchangeMetadataService
				count := 0

				for stream := range edc.Items() {
					buf := &bytes.Buffer{}
					read, err := buf.ReadFrom(stream.ToReader())
					assert.NoError(t, err)
					assert.NotZero(t, read)

					if isMetadata {
						continue
					}

					contact, err := support.CreateContactFromBytes(buf.Bytes())
					assert.NotNil(t, contact)
					assert.NoError(t, err, "error on converting contact bytes: "+buf.String())
					count++
				}

				if isMetadata {
					continue
				}

				assert.Equal(t, edc.FullPath().Folder(), exchange.DefaultContactFolder)
				assert.NotZero(t, count)
			}

			status := connector.AwaitStatus()
			suite.NotNil(status)
			suite.Equal(status.ObjectCount, status.Successful)
		})
	}
}

// TestEventsSerializationRegression ensures functionality of createCollections
// to be able to successfully query, download and restore event objects
func (suite *ConnectorCreateExchangeCollectionIntegrationSuite) TestEventsSerializationRegression() {
	ctx, flush := tester.NewContext()
	defer flush()

	connector := loadConnector(ctx, suite.T(), Users)

	tests := []struct {
		name, expected string
		getCollection  func(t *testing.T) []data.Collection
	}{
		{
			name:     "Default Event Calendar",
			expected: exchange.DefaultCalendar,
			getCollection: func(t *testing.T) []data.Collection {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{suite.user}, []string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
				collections, err := connector.createExchangeCollections(ctx, sel.Scopes()[0], nil, control.Options{})
				require.NoError(t, err)

				return collections
			},
		},
		{
			name:     "Birthday Calendar",
			expected: "Birthdays",
			getCollection: func(t *testing.T) []data.Collection {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{suite.user}, []string{"Birthdays"}, selectors.PrefixMatch()))
				collections, err := connector.createExchangeCollections(ctx, sel.Scopes()[0], nil, control.Options{})
				require.NoError(t, err)

				return collections
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collections := test.getCollection(t)
			require.Equal(t, len(collections), 1)
			edc := collections[0]
			assert.Equal(t, edc.FullPath().Folder(), test.expected)
			streamChannel := edc.Items()

			for stream := range streamChannel {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				assert.NoError(t, err)
				assert.NotZero(t, read)
				event, err := support.CreateEventFromBytes(buf.Bytes())
				assert.NotNil(t, event)
				assert.NoError(t, err, "experienced error parsing event bytes: "+buf.String())
			}

			status := connector.AwaitStatus()
			suite.NotNil(status)
			suite.Equal(status.ObjectCount, status.Successful)
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
		t      = suite.T()
		siteID = tester.M365SiteID(t)
		gc     = loadConnector(ctx, t, Sites)
		sel    = selectors.NewSharePointBackup()
	)

	tables := []struct {
		name     string
		sel      func(t *testing.T) selectors.Selector
		expected int
	}{
		{
			name:     "SharePoint.Libraries",
			expected: 0,
			sel: func(t *testing.T) selectors.Selector {
				sel.Include(sel.Libraries(
					[]string{siteID},
					[]string{"foo"},
					selectors.PrefixMatch(),
				))

				return sel.Selector
			},
		},
		{
			name:     "SharePoint.Lists",
			expected: 1,
			sel: func(t *testing.T) selectors.Selector {
				sel.Include(sel.Lists(
					[]string{siteID},
					selectors.Any(),
					selectors.PrefixMatch(),
				))

				return sel.Selector
			},
		},
	}

	for _, test := range tables {
		t.Run(test.name, func(t *testing.T) {
			cols, err := gc.DataCollections(ctx, test.sel(t), nil, control.Options{})
			require.NoError(t, err)
			assert.Equal(t, test.expected, len(cols))
		})
	}
}
