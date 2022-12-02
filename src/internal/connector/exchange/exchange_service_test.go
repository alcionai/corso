package exchange

import (
	"context"
	"testing"
	"time"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
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
)

type ExchangeServiceSuite struct {
	suite.Suite
	es *exchangeService
}

func TestExchangeServiceSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests,
		"flomp",
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
	suite.T().Skip()
	creds := suite.es.credentials
	invalidCredentials := suite.es.credentials
	invalidCredentials.AzureClientSecret = ""

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
			t.Log(test.credentials.AzureClientSecret)
			_, err := createService(test.credentials, false)
			test.checkErr(t, err)
		})
	}
}

func (suite *ExchangeServiceSuite) TestOptionsForCalendars() {
	suite.T().Skip()
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
	suite.T().Skip()
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
	suite.T().Skip()
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
	suite.T().Skip()
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

// TestGraphQueryFunctions verifies if Query functions APIs
// through Microsoft Graph are functional
func (suite *ExchangeServiceSuite) TestGraphQueryFunctions() {
	suite.T().Skip()
	ctx, flush := tester.NewContext()
	defer flush()

	userID := tester.M365UserID(suite.T())
	tests := []struct {
		name     string
		function GraphQuery
	}{
		{
			name:     "GraphQuery: Get All Contacts For User",
			function: GetAllContactsForUser,
		},
		{
			name:     "GraphQuery: Get All Folders",
			function: GetAllFolderNamesForUser,
		},
		{
			name: "GraphQuery: Get All Users",
			function: func(ctx context.Context, gs graph.Service, toss string) (absser.Parsable, error) {
				return GetAllUsersForTenant(ctx, gs)
			},
		},
		{
			name:     "GraphQuery: Get All ContactFolders",
			function: GetAllContactFolderNamesForUser,
		},
		{
			name:     "GraphQuery: Get Default ContactFolder",
			function: GetDefaultContactFolderForUser,
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

//==========================
// Restore Functions
//==========================

// TestRestoreContact ensures contact object can be created, placed into
// the Corso Folder. The function handles test clean-up.
func (suite *ExchangeServiceSuite) TestRestoreContact() {
	suite.T().Skip()
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t          = suite.T()
		userID     = tester.M365UserID(t)
		now        = time.Now()
		folderName = "TestRestoreContact: " + common.FormatSimpleDateTime(now)
	)

	aFolder, err := CreateContactFolder(ctx, suite.es, userID, folderName)
	require.NoError(t, err)

	folderID := *aFolder.GetId()

	defer func() {
		// Remove the folder containing contact prior to exiting test
		err = DeleteContactFolder(ctx, suite.es, userID, folderID)
		assert.NoError(t, err)
	}()

	info, err := RestoreExchangeContact(ctx,
		mockconnector.GetMockContactBytes("Corso TestContact"),
		suite.es,
		control.Copy,
		folderID,
		userID)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
	assert.NotNil(t, info, "contact item info")
}

// TestRestoreEvent verifies that event object is able to created
// and sent into the test account of the Corso user in the newly created Corso Calendar
func (suite *ExchangeServiceSuite) TestRestoreEvent() {
	suite.T().Skip()
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t      = suite.T()
		userID = tester.M365UserID(t)
		name   = "TestRestoreEvent: " + common.FormatSimpleDateTime(time.Now())
	)

	calendar, err := CreateCalendar(ctx, suite.es, userID, name)
	require.NoError(t, err)

	calendarID := *calendar.GetId()

	defer func() {
		// Removes calendar containing events created during the test
		err = DeleteCalendar(ctx, suite.es, userID, calendarID)
		assert.NoError(t, err)
	}()

	info, err := RestoreExchangeEvent(ctx,
		mockconnector.GetMockEventWithAttendeesBytes(name),
		suite.es,
		control.Copy,
		calendarID,
		userID)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
	assert.NotNil(t, info, "event item info")
}

// TestRestoreExchangeObject verifies path.Category usage for restored objects
func (suite *ExchangeServiceSuite) TestRestoreExchangeObject() {
	t := suite.T()
	t.Skip()
	service := loadService(t)

	userID := tester.M365UserID(t)
	now := time.Now()
	tests := []struct {
		name        string
		bytes       []byte
		category    path.CategoryType
		cleanupFunc func(context.Context, graph.Service, string, string) error
		destination func(context.Context) string
	}{
		{
			name:        "Test Mail",
			bytes:       mockconnector.GetMockMessageBytes("Restore Exchange Object"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(ctx context.Context) string {
				folderName := "TestRestoreMailObject: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.es, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: One Direct Attachment",
			bytes:       mockconnector.GetMockMessageWithDirectAttachment("Restore 1 Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.es, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: One Large Attachment",
			bytes:       mockconnector.GetMockMessageWithLargeAttachment("Restore Large Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(ctx context.Context) string {
				folderName := "TestRestoreMailwithLargeAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.es, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: Two Attachments",
			bytes:       mockconnector.GetMockMessageWithTwoAttachments("Restore 2 Attachments"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachments: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.es, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		// TODO: #884 - reinstate when able to specify root folder by name
		{
			name:        "Test Contact",
			bytes:       mockconnector.GetMockContactBytes("Test_Omega"),
			category:    path.ContactsCategory,
			cleanupFunc: DeleteContactFolder,
			destination: func(ctx context.Context) string {
				folderName := "TestRestoreContactObject: " + common.FormatSimpleDateTime(now)
				folder, err := CreateContactFolder(ctx, suite.es, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Events",
			bytes:       mockconnector.GetDefaultMockEventBytes("Restored Event Object"),
			category:    path.EventsCategory,
			cleanupFunc: DeleteCalendar,
			destination: func(ctx context.Context) string {
				calendarName := "TestRestoreEventObject: " + common.FormatSimpleDateTime(now)
				calendar, err := CreateCalendar(ctx, suite.es, userID, calendarName)
				require.NoError(t, err)

				return *calendar.GetId()
			},
		},
		{
			name:        "Test Event with Attachment",
			bytes:       mockconnector.GetMockEventWithAttachment("Restored Event Attachment"),
			category:    path.EventsCategory,
			cleanupFunc: DeleteCalendar,
			destination: func(ctx context.Context) string {
				calendarName := "TestRestoreEventObject_" + common.FormatSimpleDateTime(now)
				calendar, err := CreateCalendar(ctx, suite.es, userID, calendarName)
				require.NoError(t, err)

				return *calendar.GetId()
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			destination := test.destination(ctx)
			info, err := RestoreExchangeObject(
				ctx,
				test.bytes,
				test.category,
				control.Copy,
				service,
				destination,
				userID,
			)
			assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
			assert.NotNil(t, info, "item info is populated")

			cleanupError := test.cleanupFunc(ctx, service, userID, destination)
			assert.NoError(t, cleanupError)
		})
	}
}

// Testing to ensure that cache system works for in multiple different environments
func (suite *ExchangeServiceSuite) TestGetContainerIDFromCache() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t               = suite.T()
		user            = tester.M365UserID(t)
		connector       = loadService(t)
		directoryCaches = make(map[path.CategoryType]graph.ContainerResolver)
		folderName      = tester.DefaultTestRestoreDestination().ContainerName
		tests           = []struct {
			name      string
			pathFunc1 func() path.Path
			pathFunc2 func() path.Path
			category  path.CategoryType
		}{
			{
				name:     "Mail Cache Test",
				category: path.EmailCategory,
				pathFunc1: func() path.Path {
					pth, err := path.Builder{}.Append("Griffindor").
						Append("Croix").ToDataLayerExchangePathForCategory(
						suite.es.credentials.AzureTenantID,
						user,
						path.EmailCategory,
						false,
					)

					require.NoError(suite.T(), err)
					return pth
				},
				pathFunc2: func() path.Path {
					pth, err := path.Builder{}.Append("Griffindor").
						Append("Felicius").ToDataLayerExchangePathForCategory(
						suite.es.credentials.AzureTenantID,
						user,
						path.EmailCategory,
						false,
					)

					require.NoError(suite.T(), err)
					return pth
				},
			},
			{
				name:     "Contact Cache Test",
				category: path.ContactsCategory,
				pathFunc1: func() path.Path {
					aPath, err := path.Builder{}.Append("HufflePuff").
						ToDataLayerExchangePathForCategory(
							suite.es.credentials.AzureTenantID,
							user,
							path.ContactsCategory,
							false,
						)

					require.NoError(suite.T(), err)
					return aPath
				},
				pathFunc2: func() path.Path {
					aPath, err := path.Builder{}.Append("Ravenclaw").
						ToDataLayerExchangePathForCategory(
							suite.es.credentials.AzureTenantID,
							user,
							path.ContactsCategory,
							false,
						)

					require.NoError(suite.T(), err)
					return aPath
				},
			},
			{
				name:     "Event Cache Test",
				category: path.EventsCategory,
				pathFunc1: func() path.Path {
					aPath, err := path.Builder{}.Append("Durmstrang").
						ToDataLayerExchangePathForCategory(
							suite.es.credentials.AzureTenantID,
							user,
							path.EventsCategory,
							false,
						)
					require.NoError(suite.T(), err)
					return aPath
				},
				pathFunc2: func() path.Path {
					aPath, err := path.Builder{}.Append("Beauxbatons").
						ToDataLayerExchangePathForCategory(
							suite.es.credentials.AzureTenantID,
							user,
							path.EventsCategory,
							false,
						)
					require.NoError(suite.T(), err)
					return aPath
				},
			},
		}
	)

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			folderID, err := GetContainerIDFromCache(
				ctx,
				connector,
				test.pathFunc1(),
				folderName,
				directoryCaches,
			)

			require.NoError(t, err)
			resolver := directoryCaches[test.category]
			_, err = resolver.IDToPath(ctx, folderID)
			assert.NoError(t, err)

			secondID, err := GetContainerIDFromCache(
				ctx,
				connector,
				test.pathFunc2(),
				folderName,
				directoryCaches,
			)

			require.NoError(t, err)
			_, err = resolver.IDToPath(ctx, secondID)
			require.NoError(t, err)
			_, ok := resolver.PathInCache(folderName)
			require.True(t, ok)
		})
	}
}
