package exchange

import (
	"context"
	"testing"
	"time"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/tester"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/control"
	"github.com/alcionai/corso/pkg/selectors"
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
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			response, err := test.function(suite.es, userID)
			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}

// TestParseCalendarIDFromEvent verifies that parse function
// works on the current accepted reference format of
// additional data["calendar@odata.associationLink"]
func (suite *ExchangeServiceSuite) TestParseCalendarIDFromEvent() {
	tests := []struct {
		name       string
		input      string
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Empty string",
			input:      "",
			checkError: assert.Error,
		},
		{
			name:       "Invalid string",
			input:      "https://github.com/whyNot/calendarNot Used",
			checkError: assert.Error,
		},
		{
			name: "Missing calendarID not found",
			input: "https://graph.microsoft.com/v1.0/users" +
				"('invalid@onmicrosoft.com')/calendars(" +
				"'')/$ref",
			checkError: assert.Error,
		},
		{
			name: "Valid string",
			input: "https://graph.microsoft.com/v1.0/users" +
				"('valid@onmicrosoft.com')/calendars(" +
				"'AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAA" +
				"DCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEGAADSEBNbUIB9RL6ePDeF3FIYAAAZkDq1AAA=')/$ref",
			checkError: assert.NoError,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := parseCalendarIDFromEvent(test.input)
			test.checkError(t, err)
		})
	}
}

// TestGetMailFolderID verifies the ability to retrieve folder ID of folders
// at the top level of the file tree
func (suite *ExchangeServiceSuite) TestGetFolderID() {
	userID := tester.M365UserID(suite.T())
	tests := []struct {
		name       string
		folderName string
		// category references the current optionId :: TODO --> use selector fields
		category   optionIdentifier
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Mail Valid",
			folderName: "Inbox",
			category:   messages,
			checkError: assert.NoError,
		},
		{
			name:       "Mail Invalid",
			folderName: "FolderThatIsNotHere",
			category:   messages,
			checkError: assert.Error,
		},
		{
			name:       "Contact Invalid",
			folderName: "FolderThatIsNotHereContacts",
			category:   contacts,
			checkError: assert.Error,
		},
		{
			name:       "Contact Valid",
			folderName: "TrialFolder",
			category:   contacts,
			checkError: assert.NoError,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := GetFolderID(
				suite.es,
				test.folderName,
				userID,
				test.category)
			test.checkError(t, err, "Unable to find folder: "+test.folderName)
		})
	}
}

// TestIterativeFunctions verifies that GraphQuery to Iterate
// functions are valid for current versioning of msgraph-go-sdk
func (suite *ExchangeServiceSuite) TestIterativeFunctions() {
	userID := tester.M365UserID(suite.T())
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{userID}))
	eb, err := sel.ToExchangeBackup()
	require.NoError(suite.T(), err)
	scopes := eb.Scopes()
	var mailScope, contactScope selectors.ExchangeScope
	for _, scope := range scopes {
		if scope.IncludesCategory(selectors.ExchangeContactFolder) {
			contactScope = scope
		}
		if scope.IncludesCategory(selectors.ExchangeMail) {
			mailScope = scope
		}
	}

	tests := []struct {
		name              string
		queryFunction     GraphQuery
		iterativeFunction GraphIterateFunc
		scope             selectors.ExchangeScope
		transformer       absser.ParsableFactory
	}{
		{
			name:              "Mail Iterative Check",
			queryFunction:     GetAllMessagesForUser,
			iterativeFunction: IterateSelectAllDescendablesForCollections,
			scope:             mailScope,
			transformer:       models.CreateMessageCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Contacts Iterative Check",
			queryFunction:     GetAllContactsForUser,
			iterativeFunction: IterateSelectAllDescendablesForCollections,
			scope:             contactScope,
			transformer:       models.CreateContactFromDiscriminatorValue,
		}, {
			name:              "Folder Iterative Check",
			queryFunction:     GetAllFolderNamesForUser,
			iterativeFunction: IterateFilterFolderDirectoriesForCollections,
			scope:             mailScope,
			transformer:       models.CreateMailFolderCollectionResponseFromDiscriminatorValue,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			response, err := test.queryFunction(suite.es, userID)
			require.NoError(t, err)
			// Create Iterator
			pageIterator, err := msgraphgocore.NewPageIterator(response,
				&suite.es.adapter,
				test.transformer)
			require.NoError(t, err)
			// Create collection for iterate test
			collections := make(map[string]*Collection)
			var errs error
			// callbackFunc iterates through all models.Messageable and fills exchange.Collection.jobs[]
			// with corresponding item IDs. New collections are created for each directory
			callbackFunc := test.iterativeFunction(
				userID,
				test.scope,
				errs, false,
				suite.es.credentials,
				collections,
				nil)

			iterateError := pageIterator.Iterate(callbackFunc)
			require.NoError(t, iterateError)
		})
	}
}

// TestRestoreContact ensures contact object can be created, placed into
// the Corso Folder. The function handles test clean-up.
func (suite *ExchangeServiceSuite) TestRestoreContact() {
	t := suite.T()
	userID := tester.M365UserID(suite.T())
	now := time.Now()

	folderName := "TestRestoreContact: " + common.FormatSimpleDateTime(now)
	aFolder, err := CreateContactFolder(suite.es, userID, folderName)
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
	err = DeleteContactFolder(suite.es, userID, folderID)
	assert.NoError(t, err)
}

// TestEstablishFolder checks the ability to Create a "container" for the
// GraphConnector's Restore Workflow based on OptionIdentifier.
func (suite *ExchangeServiceSuite) TestEstablishFolder() {
	tests := []struct {
		name       string
		option     optionIdentifier
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Establish User Restore Folder",
			option:     users,
			checkError: assert.Error,
		},
		{
			name:       "Establish Event Restore Location",
			option:     events,
			checkError: assert.Error,
		},
		{
			name:       "Establish Restore Folder for Unknown",
			option:     unknown,
			checkError: assert.Error,
		},
		{
			name:       "Establish Restore folder for Mail",
			option:     messages,
			checkError: assert.NoError,
		},
		{
			name:       "Establish Restore folder for Contacts",
			option:     contacts,
			checkError: assert.NoError,
		},
	}
	now := time.Now()
	folderName := "CorsoEstablishFolder" + common.FormatSimpleDateTime(now)
	userID := tester.M365UserID(suite.T())
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			folderID, err := establishFolder(suite.es, folderName, userID, test.option)
			require.True(t, test.checkError(t, err))
			if folderID != "" {
				switch test.option {
				case messages:
					err = DeleteMailFolder(suite.es, userID, folderID)
					assert.NoError(t, err)
				case contacts:
					err = DeleteContactFolder(suite.es, userID, folderID)
					assert.NoError(t, err)
				default:
					assert.NoError(t,
						errors.New("unsupported type received folderID: "+test.option.String()),
					)
				}
			}
		})
	}
}
