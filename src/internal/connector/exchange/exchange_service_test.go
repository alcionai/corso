package exchange

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type ExchangeServiceSuite struct {
	suite.Suite
	es *exchangeService
}

func TestExchangeServiceSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ExchangeServiceSuite))
}

func (suite *ExchangeServiceSuite) SetupSuite() {
	t := suite.T()
	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(t, err)

	a := tester.NewM365Account(t)
	require.NoError(t, err)
	m365, err := a.M365Config()
	require.NoError(t, err)
	service, err := createService(m365, false)
	require.NoError(t, err)

	suite.es = service
}

// TestCreateService verifies that services are created
// when called with the correct range of params. NOTE:
// incorrect tenant or password information will NOT generate
// an error.
func (suite *ExchangeServiceSuite) TestCreateService() {
	creds := suite.es.credentials
	invalidCredentials := suite.es.credentials
	invalidCredentials.ClientSecret = ""

	tests := []struct {
		name        string
		credentials account.M365Config
		checkErr    assert.ErrorAssertionFunc
	}{
		{
			name:        "Valid Service Creation",
			credentials: creds,
			checkErr:    assert.NoError,
		},
		{
			name:        "Invalid Service Creation",
			credentials: invalidCredentials,
			checkErr:    assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			t.Log(test.credentials.ClientSecret)
			_, err := createService(test.credentials, false)
			test.checkErr(t, err)
		})
	}
}

func (suite *ExchangeServiceSuite) TestOptionsForCalendars() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Empty Literal",
			params:     []string{},
			checkError: assert.NoError,
		},
		{
			name:       "Invalid Parameter",
			params:     []string{"status"},
			checkError: assert.Error,
		},
		{
			name:       "Invalid Parameters",
			params:     []string{"status", "height", "month"},
			checkError: assert.Error,
		},
		{
			name:       "Valid Parameters",
			params:     []string{"changeKey", "events", "owner"},
			checkError: assert.NoError,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := optionsForCalendars(test.params)
			test.checkError(t, err)
		})
	}
}

// TestOptionsForMessages checks to ensure approved query
// options are added to the type specific RequestBuildConfiguration. Expected
// will be +1 on all select parameters
func (suite *ExchangeServiceSuite) TestOptionsForMessages() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Valid Message Option",
			params:     []string{"subject"},
			checkError: assert.NoError,
		},
		{
			name:       "Multiple Message Options: Accepted",
			params:     []string{"webLink", "parentFolderId"},
			checkError: assert.NoError,
		},
		{
			name:       "Invalid Message Parameter",
			params:     []string{"status"},
			checkError: assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			config, err := optionsForMessages(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(len(config.QueryParameters.Select), len(test.params)+1)
			}
		})
	}
}

// TestOptionsForFolders ensures that approved query options
// are added to the RequestBuildConfiguration. Expected will always be +1
// on than the input as "id" are always included within the select parameters
func (suite *ExchangeServiceSuite) TestOptionsForFolders() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
		expected   int
	}{
		{
			name:       "Valid Folder Option",
			params:     []string{"parentFolderId"},
			checkError: assert.NoError,
			expected:   2,
		},
		{
			name:       "Multiple Folder Options: Valid",
			params:     []string{"displayName", "isHidden"},
			checkError: assert.NoError,
			expected:   3,
		},
		{
			name:       "Invalid Folder option param",
			params:     []string{"status"},
			checkError: assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			config, err := optionsForMailFolders(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(test.expected, len(config.QueryParameters.Select))
			}
		})
	}
}

// TestOptionsForContacts similar to TestExchangeService_optionsForFolders
func (suite *ExchangeServiceSuite) TestOptionsForContacts() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
		expected   int
	}{
		{
			name:       "Valid Contact Option",
			params:     []string{"displayName"},
			checkError: assert.NoError,
			expected:   2,
		},
		{
			name:       "Multiple Contact Options: Valid",
			params:     []string{"displayName", "parentFolderId"},
			checkError: assert.NoError,
			expected:   3,
		},
		{
			name:       "Invalid Contact Option param",
			params:     []string{"status"},
			checkError: assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			options, err := optionsForContacts(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(test.expected, len(options.QueryParameters.Select))
			}
		})
	}
}

// TestSetupExchangeCollection ensures SetupExchangeCollectionVars returns a non-nil variable for
// the following selector types:
// - Mail
// - Contacts
// - Events
func (suite *ExchangeServiceSuite) TestSetupExchangeCollection() {
	userID := tester.M365UserID(suite.T())
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{userID}))
	eb, err := sel.ToExchangeBackup()
	require.NoError(suite.T(), err)

	scopes := eb.Scopes()

	for _, test := range scopes {
		suite.T().Run(test.Category().String(), func(t *testing.T) {
			discriminateFunc, graphQuery, iterFunc, err := SetupExchangeCollectionVars(test)
			assert.NoError(t, err)
			assert.NotNil(t, discriminateFunc)
			assert.NotNil(t, graphQuery)
			assert.NotNil(t, iterFunc)
		})
	}
}

// TestGraphQueryFunctions verifies if Query functions APIs
// through Microsoft Graph are functional
func (suite *ExchangeServiceSuite) TestGraphQueryFunctions() {
	ctx := context.Background()
	userID := tester.M365UserID(suite.T())
	tests := []struct {
		name     string
		function GraphQuery
	}{
		{
			name:     "GraphQuery: Get All Messages For User",
			function: GetAllMessagesForUser,
		},
		{
			name:     "GraphQuery: Get All Contacts For User",
			function: GetAllContactsForUser,
		},
		{
			name:     "GraphQuery: Get All Folders",
			function: GetAllFolderNamesForUser,
		},
		{
			name:     "GraphQuery: Get All Users",
			function: GetAllUsersForTenant,
		},
		{
			name:     "GraphQuery: Get All ContactFolders",
			function: GetAllContactFolderNamesForUser,
		},
		{
			name:     "GraphQuery: Get All Events for User",
			function: GetAllEventsForUser,
		},
		{
			name:     "GraphQuery: Get All Calendars for User",
			function: GetAllCalendarNamesForUser,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			response, err := test.function(ctx, suite.es, userID)
			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}

// TestGetMailFolderID verifies the ability to retrieve folder ID of folders
// at the top level of the file tree
func (suite *ExchangeServiceSuite) TestGetContainerID() {
	userID := tester.M365UserID(suite.T())
	ctx := context.Background()
	tests := []struct {
		name          string
		containerName string
		// category references the current optionId :: TODO --> use selector fields
		category   optionIdentifier
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:          "Mail Valid",
			containerName: DefaultMailFolder,
			category:      messages,
			checkError:    assert.NoError,
		},
		{
			name:          "Mail Invalid",
			containerName: "FolderThatIsNotHere",
			category:      messages,
			checkError:    assert.Error,
		},
		{
			name:          "Contact Invalid",
			containerName: "FolderThatIsNotHereContacts",
			category:      contacts,
			checkError:    assert.Error,
		},
		{
			name:          "Contact Valid",
			containerName: "TrialFolder",
			category:      contacts,
			checkError:    assert.NoError,
		},
		{
			name:          "Event Invalid",
			containerName: "NotAValid?@V'vCalendar",
			category:      events,
			checkError:    assert.Error,
		},
		{
			name:          "Event Valid",
			containerName: DefaultCalendar,
			category:      events,
			checkError:    assert.NoError,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := GetContainerID(
				ctx,
				suite.es,
				test.containerName,
				userID,
				test.category)
			test.checkError(t, err, "error with container: "+test.containerName)
		})
	}
}

// Restore Functions
// TestRestoreMessages uses mock data to ensure GraphConnector
// is able to restore a several messageable item to a Mailbox.
// The result should be all successful items restored within the same folder.
func (suite *ExchangeServiceSuite) TestRestoreMessages() {
	ctx := context.Background()
	userID := tester.M365UserID(suite.T())
	now := time.Now()

	folderName := "TestRestoreMessage: " + common.FormatSimpleDateTime(now)
	folder, err := CreateMailFolder(ctx, suite.es, userID, folderName)
	require.NoError(suite.T(), err)

	folderID := *folder.GetId()
	tests := []struct {
		name  string
		bytes []byte
	}{
		{
			name:  "Simple Message",
			bytes: mockconnector.GetMockMessageBytes(folderName),
		},
		{
			name:  "One Direct Attachment",
			bytes: mockconnector.GetMockMessageWithDirectAttachment(folderName),
		},
		{
			name:  "Two Attachments",
			bytes: mockconnector.GetMockMessageWithTwoAttachments(folderName),
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			err = RestoreMailMessage(context.Background(),
				test.bytes,
				suite.es,
				control.Copy,
				folderID,
				userID,
			)
			require.NoError(t, err, support.ConnectorStackErrorTrace(err))
		})
	}

	err = DeleteMailFolder(ctx, suite.es, userID, folderID)
	assert.NoError(suite.T(), err, "Failure during folder clean-up")
}

// TestRestoreContact ensures contact object can be created, placed into
// the Corso Folder. The function handles test clean-up.
func (suite *ExchangeServiceSuite) TestRestoreContact() {
	t := suite.T()
	ctx := context.Background()
	userID := tester.M365UserID(t)
	now := time.Now()

	folderName := "TestRestoreContact: " + common.FormatSimpleDateTime(now)
	aFolder, err := CreateContactFolder(ctx, suite.es, userID, folderName)
	require.NoError(t, err)

	folderID := *aFolder.GetId()
	err = RestoreExchangeContact(context.Background(),
		mockconnector.GetMockContactBytes("Corso TestContact"),
		suite.es,
		control.Copy,
		folderID,
		userID)
	assert.NoError(t, err)
	// Removes folder containing contact prior to exiting test
	err = DeleteContactFolder(ctx, suite.es, userID, folderID)
	assert.NoError(t, err)
}

// TestRestoreEvent verifies that event object is able to created
// and sent into the test account of the Corso user in the newly created Corso Calendar
func (suite *ExchangeServiceSuite) TestRestoreEvent() {
	t := suite.T()
	ctx := context.Background()
	userID := tester.M365UserID(t)
	name := "TestRestoreEvent: " + common.FormatSimpleDateTime(time.Now())
	calendar, err := CreateCalendar(ctx, suite.es, userID, name)
	require.NoError(t, err)

	calendarID := *calendar.GetId()
	err = RestoreExchangeEvent(context.Background(),
		mockconnector.GetMockEventWithAttendeesBytes(name),
		suite.es,
		control.Copy,
		calendarID,
		userID)
	assert.NoError(t, err)
	// Removes calendar containing events created during the test
	err = DeleteCalendar(ctx, suite.es, userID, *calendar.GetId())
	assert.NoError(t, err)
}

// TestGetRestoreContainer checks the ability to Create a "container" for the
// GraphConnector's Restore Workflow based on OptionIdentifier.
func (suite *ExchangeServiceSuite) TestGetRestoreContainer() {
	ctx := context.Background()
	dest := tester.DefaultTestRestoreDestination()
	tests := []struct {
		name        string
		option      path.CategoryType
		checkError  assert.ErrorAssertionFunc
		cleanupFunc func(context.Context, graph.Service, string, string) error
	}{
		{
			name:        "Establish User Restore Folder",
			option:      path.CategoryType(-1),
			checkError:  assert.Error,
			cleanupFunc: nil,
		},

		// TODO: #884 - reinstate when able to specify root folder by name
		// {
		// 	name:        "Establish Event Restore Location",
		// 	option:      path.EventsCategory,
		// 	checkError:  assert.NoError,
		// 	cleanupFunc: DeleteCalendar,
		// },
		{
			name:        "Establish Restore Folder for Unknown",
			option:      path.UnknownCategory,
			checkError:  assert.Error,
			cleanupFunc: nil,
		},
		{
			name:        "Establish Restore folder for Mail",
			option:      path.EmailCategory,
			checkError:  assert.NoError,
			cleanupFunc: DeleteMailFolder,
		},
		// TODO: #884 - reinstate when able to specify root folder by name
		// {
		// 	name:        "Establish Restore folder for Contacts",
		// 	option:      path.ContactsCategory,
		// 	checkError:  assert.NoError,
		// 	cleanupFunc: DeleteContactFolder,
		// },
	}

	userID := tester.M365UserID(suite.T())

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			containerID, err := GetRestoreContainer(ctx, suite.es, userID, test.option, dest.ContainerName)
			require.True(t, test.checkError(t, err, support.ConnectorStackErrorTrace(err)))

			if test.cleanupFunc != nil {
				err = test.cleanupFunc(ctx, suite.es, userID, containerID)
				assert.NoError(t, err)
			}
		})
	}
}

// TestRestoreExchangeObject verifies path.Category usage for restored objects
func (suite *ExchangeServiceSuite) TestRestoreExchangeObject() {
	ctx := context.Background()
	t := suite.T()
	userID := tester.M365UserID(t)
	service := loadService(t)
	now := time.Now()
	tests := []struct {
		name        string
		bytes       []byte
		category    path.CategoryType
		cleanupFunc func(context.Context, graph.Service, string, string) error
		destination func() string
	}{
		{
			name:        "Test Mail",
			bytes:       mockconnector.GetMockMessageBytes("Restore Exchange Object"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func() string {
				folderName := "TestRestoreMailObject: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.es, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		// TODO: #884 - reinstate when able to specify root folder by name
		// {
		// 	name:        "Test Contact",
		// 	bytes:       mockconnector.GetMockContactBytes("Test_Omega"),
		// 	category:    path.ContactsCategory,
		// 	cleanupFunc: DeleteContactFolder,
		// 	destination: func() string {
		// 		folderName := "TestRestoreContactObject: " + common.FormatSimpleDateTime(now)
		// 		folder, err := CreateContactFolder(suite.es, userID, folderName)
		// 		require.NoError(t, err)

		// 		return *folder.GetId()
		// 	},
		// },
		// {
		// 	name:        "Test Events",
		// 	bytes:       mockconnector.GetMockEventBytes("Restored Event Object"),
		// 	category:    path.EventsCategory,
		// 	cleanupFunc: DeleteCalendar,
		// 	destination: func() string {
		// 		calendarName := "TestRestoreEventObject: " + common.FormatSimpleDateTime(now)
		// 		calendar, err := CreateCalendar(suite.es, userID, calendarName)
		// 		require.NoError(t, err)

		// 		return *calendar.GetId()
		// 	},
		// },
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			destination := test.destination()
			err := RestoreExchangeObject(
				ctx,
				test.bytes,
				test.category,
				control.Copy,
				service,
				destination,
				userID,
			)
			assert.NoError(t, err)
			cleanupError := test.cleanupFunc(ctx, service, userID, destination)
			assert.NoError(t, cleanupError)
		})
	}
}
