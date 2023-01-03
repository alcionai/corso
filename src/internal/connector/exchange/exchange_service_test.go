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
)

type ExchangeServiceSuite struct {
	suite.Suite
	credentials account.M365Config
	gs          graph.Servicer
}

func TestExchangeServiceSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

	suite.Run(t, new(ExchangeServiceSuite))
}

func (suite *ExchangeServiceSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := createService(m365)
	require.NoError(t, err)

	suite.credentials = m365
	suite.gs = service
}

// TestCreateService verifies that services are created
// when called with the correct range of params. NOTE:
// incorrect tenant or password information will NOT generate
// an error.
func (suite *ExchangeServiceSuite) TestCreateService() {
	creds := suite.credentials
	invalidCredentials := suite.credentials
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
			_, err := createService(test.credentials)
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

// TestGraphQueryFunctions verifies if Query functions APIs
// through Microsoft Graph are functional
func (suite *ExchangeServiceSuite) TestGraphQueryFunctions() {
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
			response, err := test.function(ctx, suite.gs, userID)
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
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t          = suite.T()
		userID     = tester.M365UserID(t)
		now        = time.Now()
		folderName = "TestRestoreContact: " + common.FormatSimpleDateTime(now)
	)

	aFolder, err := CreateContactFolder(ctx, suite.gs, userID, folderName)
	require.NoError(t, err)

	folderID := *aFolder.GetId()

	defer func() {
		// Remove the folder containing contact prior to exiting test
		err = DeleteContactFolder(ctx, suite.gs, userID, folderID)
		assert.NoError(t, err)
	}()

	info, err := RestoreExchangeContact(ctx,
		mockconnector.GetMockContactBytes("Corso TestContact"),
		suite.gs,
		control.Copy,
		folderID,
		userID)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
	assert.NotNil(t, info, "contact item info")
}

// TestRestoreEvent verifies that event object is able to created
// and sent into the test account of the Corso user in the newly created Corso Calendar
func (suite *ExchangeServiceSuite) TestRestoreEvent() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t      = suite.T()
		userID = tester.M365UserID(t)
		name   = "TestRestoreEvent: " + common.FormatSimpleDateTime(time.Now())
	)

	calendar, err := CreateCalendar(ctx, suite.gs, userID, name)
	require.NoError(t, err)

	calendarID := *calendar.GetId()

	defer func() {
		// Removes calendar containing events created during the test
		err = DeleteCalendar(ctx, suite.gs, userID, calendarID)
		assert.NoError(t, err)
	}()

	info, err := RestoreExchangeEvent(ctx,
		mockconnector.GetMockEventWithAttendeesBytes(name),
		suite.gs,
		control.Copy,
		calendarID,
		userID)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
	assert.NotNil(t, info, "event item info")
}

// TestRestoreExchangeObject verifies path.Category usage for restored objects
func (suite *ExchangeServiceSuite) TestRestoreExchangeObject() {
	a := tester.NewM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	service, err := createService(m365)
	require.NoError(suite.T(), err)

	userID := tester.M365UserID(suite.T())
	now := time.Now()
	tests := []struct {
		name        string
		bytes       []byte
		category    path.CategoryType
		cleanupFunc func(context.Context, graph.Servicer, string, string) error
		destination func(*testing.T, context.Context) string
	}{
		{
			name:        "Test Mail",
			bytes:       mockconnector.GetMockMessageBytes("Restore Exchange Object"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailObject: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.gs, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: One Direct Attachment",
			bytes:       mockconnector.GetMockMessageWithDirectAttachment("Restore 1 Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.gs, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: One Large Attachment",
			bytes:       mockconnector.GetMockMessageWithLargeAttachment("Restore Large Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithLargeAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.gs, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: Two Attachments",
			bytes:       mockconnector.GetMockMessageWithTwoAttachments("Restore 2 Attachments"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachments: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.gs, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: Reference(OneDrive) Attachment",
			bytes:       mockconnector.GetMessageWithOneDriveAttachment("Restore Reference(OneDrive) Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithReferenceAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := CreateMailFolder(ctx, suite.gs, userID, folderName)
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
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreContactObject: " + common.FormatSimpleDateTime(now)
				folder, err := CreateContactFolder(ctx, suite.gs, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Events",
			bytes:       mockconnector.GetDefaultMockEventBytes("Restored Event Object"),
			category:    path.EventsCategory,
			cleanupFunc: DeleteCalendar,
			destination: func(t *testing.T, ctx context.Context) string {
				calendarName := "TestRestoreEventObject: " + common.FormatSimpleDateTime(now)
				calendar, err := CreateCalendar(ctx, suite.gs, userID, calendarName)
				require.NoError(t, err)

				return *calendar.GetId()
			},
		},
		{
			name:        "Test Event with Attachment",
			bytes:       mockconnector.GetMockEventWithAttachment("Restored Event Attachment"),
			category:    path.EventsCategory,
			cleanupFunc: DeleteCalendar,
			destination: func(t *testing.T, ctx context.Context) string {
				calendarName := "TestRestoreEventObject_" + common.FormatSimpleDateTime(now)
				calendar, err := CreateCalendar(ctx, suite.gs, userID, calendarName)
				require.NoError(t, err)

				return *calendar.GetId()
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			destination := test.destination(t, ctx)
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

	a := tester.NewM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	connector, err := createService(m365)
	require.NoError(suite.T(), err)

	var (
		user            = tester.M365UserID(suite.T())
		directoryCaches = make(map[path.CategoryType]graph.ContainerResolver)
		folderName      = tester.DefaultTestRestoreDestination().ContainerName
		tests           = []struct {
			name      string
			pathFunc1 func(t *testing.T) path.Path
			pathFunc2 func(t *testing.T) path.Path
			category  path.CategoryType
		}{
			{
				name:     "Mail Cache Test",
				category: path.EmailCategory,
				pathFunc1: func(t *testing.T) path.Path {
					pth, err := path.Builder{}.Append("Griffindor").
						Append("Croix").ToDataLayerExchangePathForCategory(
						suite.credentials.AzureTenantID,
						user,
						path.EmailCategory,
						false,
					)

					require.NoError(t, err)
					return pth
				},
				pathFunc2: func(t *testing.T) path.Path {
					pth, err := path.Builder{}.Append("Griffindor").
						Append("Felicius").ToDataLayerExchangePathForCategory(
						suite.credentials.AzureTenantID,
						user,
						path.EmailCategory,
						false,
					)

					require.NoError(t, err)
					return pth
				},
			},
			{
				name:     "Contact Cache Test",
				category: path.ContactsCategory,
				pathFunc1: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.Append("HufflePuff").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.ContactsCategory,
							false,
						)

					require.NoError(t, err)
					return aPath
				},
				pathFunc2: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.Append("Ravenclaw").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.ContactsCategory,
							false,
						)

					require.NoError(t, err)
					return aPath
				},
			},
			{
				name:     "Event Cache Test",
				category: path.EventsCategory,
				pathFunc1: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.Append("Durmstrang").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.EventsCategory,
							false,
						)
					require.NoError(t, err)
					return aPath
				},
				pathFunc2: func(t *testing.T) path.Path {
					aPath, err := path.Builder{}.Append("Beauxbatons").
						ToDataLayerExchangePathForCategory(
							suite.credentials.AzureTenantID,
							user,
							path.EventsCategory,
							false,
						)
					require.NoError(t, err)
					return aPath
				},
			},
		}
	)

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			folderID, err := CreateContainerDestinaion(
				ctx,
				connector,
				test.pathFunc1(t),
				folderName,
				directoryCaches)

			require.NoError(t, err)
			resolver := directoryCaches[test.category]
			_, err = resolver.IDToPath(ctx, folderID)
			assert.NoError(t, err)

			secondID, err := CreateContainerDestinaion(
				ctx,
				connector,
				test.pathFunc2(t),
				folderName,
				directoryCaches)

			require.NoError(t, err)
			_, err = resolver.IDToPath(ctx, secondID)
			require.NoError(t, err)
			_, ok := resolver.PathInCache(folderName)
			require.True(t, ok)
		})
	}
}
