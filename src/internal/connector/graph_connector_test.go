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
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
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
	t := suite.T()
	now := time.Now()
	folderName := "TestFolder: " + common.FormatSimpleDateTime(now)
	aFolder, err := exchange.CreateMailFolder(ctx, suite.connector.Service(), suite.user, folderName)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))

	if aFolder != nil {
		secondFolder, err := exchange.CreateMailFolderWithParent(
			ctx,
			suite.connector.Service(),
			suite.user,
			"SubFolder",
			*aFolder.GetId(),
		)
		assert.NoError(t, err)
		assert.True(t, *secondFolder.GetParentFolderId() == *aFolder.GetId())

		err = exchange.DeleteMailFolder(ctx, suite.connector.Service(), suite.user, *aFolder.GetId())
		assert.NoError(t, err)

		if err != nil {
			t.Log(support.ConnectorStackErrorTrace(err))
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

func (suite *GraphConnectorIntegrationSuite) TestEmptyCollections() {
	dest := tester.DefaultTestRestoreDestination()
	table := []struct {
		name string
		col  []data.Collection
		sel  selectors.Selector
	}{
		{
			name: "ExchangeNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceExchange,
			},
		},
		{
			name: "ExchangeEmpty",
			col:  []data.Collection{},
			sel: selectors.Selector{
				Service: selectors.ServiceExchange,
			},
		},
		{
			name: "OneDriveNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceOneDrive,
			},
		},
		{
			name: "OneDriveEmpty",
			col:  []data.Collection{},
			sel: selectors.Selector{
				Service: selectors.ServiceOneDrive,
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx := context.Background()

			deets, err := suite.connector.RestoreDataCollections(ctx, test.sel, dest, test.col)
			require.NoError(t, err)
			assert.NotNil(t, deets)

			stats := suite.connector.AwaitStatus()
			assert.Zero(t, stats.ObjectCount)
			assert.Zero(t, stats.FolderCount)
			assert.Zero(t, stats.Successful)
		})
	}
}

/*
func (suite *GraphConnectorIntegrationSuite) TestRestoreAndBackup() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

	table := []struct {
		name                   string
		service                path.ServiceType
		collections            []colInfo
		backupSelFunc          func(dest control.RestoreDestination, backupUser string) selectors.Selector
		expectedRestoreFolders int
	}{
		{
			name:                   "EmailsWithAttachments",
			service:                path.ExchangeService,
			expectedRestoreFolders: 1,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithDirectAttachment(
								subjectText + "-1",
							),
							lookupKey: subjectText + "-1",
						},
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithTwoAttachments(
								subjectText + "-2",
							),
							lookupKey: subjectText + "-2",
						},
					},
				},
			},
			// TODO(ashmrtn): Generalize this once we know the path transforms that
			// occur during restore.
			backupSelFunc: func(dest control.RestoreDestination, backupUser string) selectors.Selector {
				backupSel := selectors.NewExchangeBackup()
				backupSel.Include(backupSel.MailFolders(
					[]string{backupUser},
					[]string{dest.ContainerName},
				))

				return backupSel.Selector
			},
		},
		{
			name:                   "MultipleEmailsSingleFolder",
			service:                path.ExchangeService,
			expectedRestoreFolders: 1,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-1",
								bodyText+" 1.",
							),
							lookupKey: subjectText + "-1",
						},
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-2",
								bodyText+" 2.",
							),
							lookupKey: subjectText + "-2",
						},
						{
							name: "someencodeditemID3",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-3",
								bodyText+" 3.",
							),
							lookupKey: subjectText + "-3",
						},
					},
				},
			},
			// TODO(ashmrtn): Generalize this once we know the path transforms that
			// occur during restore.
			backupSelFunc: func(dest control.RestoreDestination, backupUser string) selectors.Selector {
				backupSel := selectors.NewExchangeBackup()
				backupSel.Include(backupSel.MailFolders(
					[]string{backupUser},
					[]string{dest.ContainerName},
				))

				return backupSel.Selector
			},
		},
		{
			name:                   "MultipleContactsSingleFolder",
			service:                path.ExchangeService,
			expectedRestoreFolders: 1,
			collections: []colInfo{
				{
					pathElements: []string{"Contacts"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
			},
			// TODO(ashmrtn): Generalize this once we know the path transforms that
			// occur during restore.
			backupSelFunc: func(dest control.RestoreDestination, backupUser string) selectors.Selector {
				backupSel := selectors.NewExchangeBackup()
				backupSel.Include(backupSel.ContactFolders(
					[]string{backupUser},
					[]string{dest.ContainerName},
				))

				return backupSel.Selector
			},
		},
		{
			name:                   "MultipleContactsMutlipleFolders",
			service:                path.ExchangeService,
			expectedRestoreFolders: 1,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID4",
							data:      mockconnector.GetMockContactBytes("Argon"),
							lookupKey: "Argon",
						},
						{
							name:      "someencodeditemID5",
							data:      mockconnector.GetMockContactBytes("Bernard"),
							lookupKey: "Bernard",
						},
					},
				},
			},
			// TODO(ashmrtn): Generalize this once we know the path transforms that
			// occur during restore.
			backupSelFunc: func(dest control.RestoreDestination, backupUser string) selectors.Selector {
				backupSel := selectors.NewExchangeBackup()
				backupSel.Include(backupSel.ContactFolders(
					[]string{backupUser},
					[]string{dest.ContainerName},
				))

				return backupSel.Selector
			},
		},
		{
			name:                   "MultipleEventsSingleCalendar",
			service:                path.ExchangeService,
			expectedRestoreFolders: 1,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.EventsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockEventWithSubjectBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
			},
			// TODO(ashmrtn): Generalize this once we know the path transforms that
			// occur during restore.
			backupSelFunc: func(dest control.RestoreDestination, backupUser string) selectors.Selector {
				backupSel := selectors.NewExchangeBackup()
				backupSel.Include(backupSel.EventCalendars(
					[]string{backupUser},
					[]string{dest.ContainerName},
				))

				return backupSel.Selector
			},
		},
		{
			name:                   "MultipleEventsMultipleCalendars",
			service:                path.ExchangeService,
			expectedRestoreFolders: 2,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.EventsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockEventWithSubjectBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.EventsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID4",
							data:      mockconnector.GetMockEventWithSubjectBytes("Argon"),
							lookupKey: "Argon",
						},
						{
							name:      "someencodeditemID5",
							data:      mockconnector.GetMockEventWithSubjectBytes("Bernard"),
							lookupKey: "Bernard",
						},
					},
				},
			},
			// TODO(ashmrtn): Generalize this once we know the path transforms that
			// occur during restore.
			backupSelFunc: func(dest control.RestoreDestination, backupUser string) selectors.Selector {
				backupSel := selectors.NewExchangeBackup()
				backupSel.Include(backupSel.EventCalendars(
					[]string{backupUser},
					[]string{dest.ContainerName},
				))

				return backupSel.Selector
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			// Get a dest per test so they're independent.
			dest := tester.DefaultTestRestoreDestination()

			totalItems, collections, expectedData := collectionsForInfo(
				t,
				test.service,
				suite.connector.tenant,
				suite.user,
				dest,
				test.collections,
			)

			t.Logf("Restoring collections to %s\n", dest.ContainerName)

			restoreGC := loadConnector(ctx, t)
			restoreSel := getSelectorWith(test.service)
			deets, err := restoreGC.RestoreDataCollections(ctx, restoreSel, dest, collections)
			require.NoError(t, err)
			assert.NotNil(t, deets)

			status := restoreGC.AwaitStatus()
			assert.Equal(t, test.expectedRestoreFolders, status.FolderCount, "status.FolderCount")
			assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
			assert.Equal(t, totalItems, status.Successful, "status.Successful")
			assert.Equal(
				t, totalItems, len(deets.Entries),
				"details entries contains same item count as total successful items restored")

			t.Logf("Restore complete\n")

			// Run a backup and compare its output with what we put in.

			backupGC := loadConnector(ctx, t)
			backupSel := test.backupSelFunc(dest, suite.user)
			t.Logf("Selective backup of %s\n", backupSel)

			dcs, err := backupGC.DataCollections(ctx, backupSel)
			require.NoError(t, err)

			t.Logf("Backup enumeration complete\n")

			// Pull the data prior to waiting for the status as otherwise it will
			// deadlock.
			checkCollections(t, totalItems, expectedData, dcs)

			status = backupGC.AwaitStatus()
			// TODO(ashmrtn): This will need to change when the restore layout is
			// updated.
			assert.Equal(t, 1, status.FolderCount, "status.FolderCount")
			assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
			assert.Equal(t, totalItems, status.Successful, "status.Successful")
		})
	}
}
*/

/*
func (suite *GraphConnectorIntegrationSuite) TestMultiFolderBackupDifferentNames() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

	// TODO(ashmrtn): Update if we start mixing categories during backup/restore.
	backupSelFunc := func(
		dests []control.RestoreDestination,
		category path.CategoryType,
		backupUser string,
	) selectors.Selector {
		destNames := make([]string, 0, len(dests))

		for _, d := range dests {
			destNames = append(destNames, d.ContainerName)
		}

		backupSel := selectors.NewExchangeBackup()

		switch category {
		case path.EmailCategory:
			backupSel.Include(backupSel.MailFolders(
				[]string{backupUser},
				destNames,
			))
		case path.ContactsCategory:
			backupSel.Include(backupSel.ContactFolders(
				[]string{backupUser},
				destNames,
			))
		case path.EventsCategory:
			backupSel.Include(backupSel.EventCalendars(
				[]string{backupUser},
				destNames,
			))
		}

		return backupSel.Selector
	}

	table := []struct {
		name     string
		service  path.ServiceType
		category path.CategoryType
		// Each collection will be restored separately, creating multiple folders to
		// backup later.
		collections []colInfo
	}{
		{
			name:     "Email",
			service:  path.ExchangeService,
			category: path.EmailCategory,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-1",
								bodyText+" 1.",
							),
							lookupKey: subjectText + "-1",
						},
					},
				},
				{
					pathElements: []string{"Archive"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-2",
								bodyText+" 2.",
							),
							lookupKey: subjectText + "-2",
						},
					},
				},
			},
		},
		{
			name:     "Contacts",
			service:  path.ExchangeService,
			category: path.ContactsCategory,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
					},
				},
			},
		},
		{
			name:     "Events",
			service:  path.ExchangeService,
			category: path.EventsCategory,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.EventsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.EventsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
							lookupKey: "Irgot",
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			restoreSel := getSelectorWith(test.service)
			dests := make([]control.RestoreDestination, 0, len(test.collections))
			allItems := 0
			allExpectedData := map[string]map[string][]byte{}

			for i, collection := range test.collections {
				// Get a dest per collection so they're independent.
				dest := tester.DefaultTestRestoreDestination()
				dests = append(dests, dest)

				totalItems, collections, expectedData := collectionsForInfo(
					t,
					test.service,
					suite.connector.tenant,
					suite.user,
					dest,
					[]colInfo{collection},
				)
				allItems += totalItems

				for k, v := range expectedData {
					allExpectedData[k] = v
				}

				t.Logf(
					"Restoring %v/%v collections to %s\n",
					i+1,
					len(test.collections),
					dest.ContainerName,
				)

				restoreGC := loadConnector(ctx, t)
				deets, err := restoreGC.RestoreDataCollections(ctx, restoreSel, dest, collections)
				require.NoError(t, err)
				require.NotNil(t, deets)

				status := restoreGC.AwaitStatus()
				// Always just 1 because it's just 1 collection.
				assert.Equal(t, 1, status.FolderCount, "status.FolderCount")
				assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
				assert.Equal(t, totalItems, status.Successful, "status.Successful")
				assert.Equal(
					t, totalItems, len(deets.Entries),
					"details entries contains same item count as total successful items restored")

				t.Logf("Restore complete\n")
			}

			// Run a backup and compare its output with what we put in.

			backupGC := loadConnector(ctx, t)
			backupSel := backupSelFunc(dests, test.category, suite.user)
			t.Logf("Selective backup of %s\n", backupSel)

			dcs, err := backupGC.DataCollections(ctx, backupSel)
			require.NoError(t, err)

			t.Logf("Backup enumeration complete\n")

			// Pull the data prior to waiting for the status as otherwise it will
			// deadlock.
			checkCollections(t, allItems, allExpectedData, dcs)

			status := backupGC.AwaitStatus()
			assert.Equal(t, len(test.collections), status.FolderCount, "status.FolderCount")
			assert.Equal(t, allItems, status.ObjectCount, "status.ObjectCount")
			assert.Equal(t, allItems, status.Successful, "status.Successful")
		})
	}
}
*/
