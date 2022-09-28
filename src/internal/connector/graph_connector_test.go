package connector

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
	user      string
}

func loadConnector(ctx context.Context, t *testing.T) *GraphConnector {
	a := tester.NewM365Account(t)
	connector, err := NewGraphConnector(ctx, a)
	require.NoError(t, err)

	return connector
}

func TestGraphConnectorIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(GraphConnectorIntegrationSuite))
}

func (suite *GraphConnectorIntegrationSuite) SetupSuite() {
	if err := tester.RunOnAny(tester.CorsoCITests); err != nil {
		suite.T().Skip(err)
	}

	ctx := context.Background()
	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)
	suite.connector = loadConnector(ctx, suite.T())
	suite.user = tester.M365UserID(suite.T())
	tester.LogTimeOfTest(suite.T())
}

// TestSetTenantUsers verifies GraphConnector's ability to query
// the users associated with the credentials
func (suite *GraphConnectorIntegrationSuite) TestSetTenantUsers() {
	newConnector := GraphConnector{
		tenant:      "test_tenant",
		Users:       make(map[string]string, 0),
		credentials: suite.connector.credentials,
	}
	ctx := context.Background()
	service, err := newConnector.createService(false)
	require.NoError(suite.T(), err)

	newConnector.graphService = *service

	suite.Equal(len(newConnector.Users), 0)
	err = newConnector.setTenantUsers(ctx)
	assert.NoError(suite.T(), err)
	suite.Greater(len(newConnector.Users), 0)
}

// TestExchangeDataCollection verifies interface between operation and
// GraphConnector remains stable to receive a non-zero amount of Collections
// for the Exchange Package. Enabled exchange applications:
// - mail
// - contacts
// - events
func (suite *GraphConnectorIntegrationSuite) TestExchangeDataCollection() {
	ctx := context.Background()
	connector := loadConnector(ctx, suite.T())
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: suite.user + " Email",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.MailFolders([]string{suite.user}, []string{exchange.DefaultMailFolder}))

				return sel.Selector
			},
		},
		{
			name: suite.user + " Contacts",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.ContactFolders([]string{suite.user}, []string{exchange.DefaultContactFolder}))

				return sel.Selector
			},
		},
		{
			name: suite.user + " Events",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{suite.user}, []string{exchange.DefaultCalendar}))

				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			collection, err := connector.ExchangeDataCollection(ctx, test.getSelector(t))
			require.NoError(t, err)
			assert.Equal(t, len(collection), 1)
			channel := collection[0].Items()
			for object := range channel {
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

// TestMailSerializationRegression verifies that all mail data stored in the
// test account can be successfully downloaded into bytes and restored into
// M365 mail objects
func (suite *GraphConnectorIntegrationSuite) TestMailSerializationRegression() {
	ctx := context.Background()
	t := suite.T()
	connector := loadConnector(ctx, t)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.MailFolders([]string{suite.user}, []string{exchange.DefaultMailFolder}))
	collection, err := connector.createCollections(context.Background(), sel.Scopes()[0])
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
func (suite *GraphConnectorIntegrationSuite) TestContactSerializationRegression() {
	connector := loadConnector(context.Background(), suite.T())

	tests := []struct {
		name          string
		getCollection func(t *testing.T) []*exchange.Collection
	}{
		{
			name: "Default Contact Folder",
			getCollection: func(t *testing.T) []*exchange.Collection {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.ContactFolders([]string{suite.user}, []string{exchange.DefaultContactFolder}))
				collections, err := connector.createCollections(context.Background(), sel.Scopes()[0])
				require.NoError(t, err)

				return collections
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			edcs := test.getCollection(t)
			assert.Equal(t, len(edcs), 1)
			edc := edcs[0]
			assert.Equal(t, edc.FullPath().Folder(), exchange.DefaultContactFolder)
			streamChannel := edc.Items()
			count := 0
			for stream := range streamChannel {
				buf := &bytes.Buffer{}
				read, err := buf.ReadFrom(stream.ToReader())
				assert.NoError(t, err)
				assert.NotZero(t, read)
				contact, err := support.CreateContactFromBytes(buf.Bytes())
				assert.NotNil(t, contact)
				assert.NoError(t, err, "error on converting contact bytes: "+string(buf.Bytes()))
				count++
			}
			assert.NotZero(t, count)

			status := connector.AwaitStatus()
			suite.NotNil(status)
			suite.Equal(status.ObjectCount, status.Successful)
		})
	}
}

// TestEventsSerializationRegression ensures functionality of createCollections
// to be able to successfully query, download and restore event objects
func (suite *GraphConnectorIntegrationSuite) TestEventsSerializationRegression() {
	connector := loadConnector(context.Background(), suite.T())

	tests := []struct {
		name, expected string
		getCollection  func(t *testing.T) []*exchange.Collection
	}{
		{
			name:     "Default Event Calendar",
			expected: exchange.DefaultCalendar,
			getCollection: func(t *testing.T) []*exchange.Collection {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{suite.user}, []string{exchange.DefaultCalendar}))
				collections, err := connector.createCollections(context.Background(), sel.Scopes()[0])
				require.NoError(t, err)

				return collections
			},
		},
		{
			name:     "Birthday Calendar",
			expected: "Birthdays",
			getCollection: func(t *testing.T) []*exchange.Collection {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{suite.user}, []string{"Birthdays"}))
				collections, err := connector.createCollections(context.Background(), sel.Scopes()[0])
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
				assert.NoError(t, err, "experienced error parsing event bytes: "+string(buf.Bytes()))
			}

			status := connector.AwaitStatus()
			suite.NotNil(status)
			suite.Equal(status.ObjectCount, status.Successful)
		})
	}
}

// TestAccessOfInboxAllUsers verifies that GraphConnector can
// support `--all-users` for backup operations. Selector.DiscreteScopes
// returns all of the users within one scope. Only users who have
// messages in their inbox will have a collection returned.
// The final test insures that more than a 75% of the user collections are
// returned. If an error was experienced, the test will fail overall
func (suite *GraphConnectorIntegrationSuite) TestAccessOfInboxAllUsers() {
	ctx := context.Background()
	t := suite.T()
	connector := loadConnector(ctx, t)
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.MailFolders(selectors.Any(), []string{exchange.DefaultMailFolder}))
	scopes := sel.DiscreteScopes(connector.GetUsers())

	for _, scope := range scopes {
		users := scope.Get(selectors.ExchangeUser)
		standard := (len(users) / 4) * 3
		collections, err := connector.createCollections(context.Background(), scope)
		require.NoError(t, err)
		suite.Greater(len(collections), standard)
	}
}

///------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------

// TestCreateAndDeleteMailFolder ensures GraphConnector has the ability
// to create and remove folders within the tenant
func (suite *GraphConnectorIntegrationSuite) TestCreateAndDeleteMailFolder() {
	ctx := context.Background()
	now := time.Now()
	folderName := "TestFolder: " + common.FormatSimpleDateTime(now)
	aFolder, err := exchange.CreateMailFolder(ctx, suite.connector.Service(), suite.user, folderName)
	assert.NoError(suite.T(), err, support.ConnectorStackErrorTrace(err))

	if aFolder != nil {
		err = exchange.DeleteMailFolder(ctx, suite.connector.Service(), suite.user, *aFolder.GetId())
		assert.NoError(suite.T(), err)

		if err != nil {
			suite.T().Log(support.ConnectorStackErrorTrace(err))
		}
	}
}

// TestCreateAndDeleteContactFolder ensures GraphConnector has the ability
// to create and remove contact folders within the tenant
func (suite *GraphConnectorIntegrationSuite) TestCreateAndDeleteContactFolder() {
	ctx := context.Background()
	now := time.Now()
	folderName := "TestContactFolder: " + common.FormatSimpleDateTime(now)
	aFolder, err := exchange.CreateContactFolder(ctx, suite.connector.Service(), suite.user, folderName)
	assert.NoError(suite.T(), err)

	if aFolder != nil {
		err = exchange.DeleteContactFolder(ctx, suite.connector.Service(), suite.user, *aFolder.GetId())
		assert.NoError(suite.T(), err)

		if err != nil {
			suite.T().Log(support.ConnectorStackErrorTrace(err))
		}
	}
}

// TestCreateAndDeleteCalendar verifies GraphConnector has the ability to create and remove
// exchange.Event.Calendars within the tenant
func (suite *GraphConnectorIntegrationSuite) TestCreateAndDeleteCalendar() {
	ctx := context.Background()
	now := time.Now()
	service := suite.connector.Service()
	calendarName := "TestCalendar: " + common.FormatSimpleDateTime(now)
	calendar, err := exchange.CreateCalendar(ctx, service, suite.user, calendarName)
	assert.NoError(suite.T(), err)

	if calendar != nil {
		err = exchange.DeleteCalendar(ctx, service, suite.user, *calendar.GetId())
		assert.NoError(suite.T(), err)

		if err != nil {
			suite.T().Log(support.ConnectorStackErrorTrace(err))
		}
	}
}

func (suite *GraphConnectorIntegrationSuite) TestRestoreContact() {
	t := suite.T()
	sel := selectors.NewExchangeRestore()
	fullpath, err := path.Builder{}.Append("testing").
		ToDataLayerExchangePathForCategory(
			suite.connector.tenant,
			suite.user,
			path.ContactsCategory,
			false,
		)

	require.NoError(t, err)
	aPath, err := path.Builder{}.Append("validator").ToDataLayerExchangePathForCategory(
		suite.connector.tenant,
		suite.user,
		path.ContactsCategory,
		false,
	)
	require.NoError(t, err)

	dcs := mockconnector.NewMockContactCollection(fullpath, 3)
	two := mockconnector.NewMockContactCollection(aPath, 2)
	collections := []data.Collection{dcs, two}
	ctx := context.Background()
	connector := loadConnector(ctx, suite.T())
	dest := control.DefaultRestoreDestination(common.SimpleDateTimeFormat)
	err = connector.RestoreDataCollections(ctx, sel.Selector, dest, collections)
	assert.NoError(suite.T(), err)

	value := connector.AwaitStatus()
	assert.Equal(t, value.FolderCount, 1)
	suite.T().Log(value.String())
}
