package exchange

import (
	"testing"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/tester"
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

// TestExchangeService_optionsForMessages checks to ensure approved query
// options are added to the type specific RequestBuildConfiguration. Expected
// will be +1 on all select parameters
func (suite *ExchangeServiceSuite) TestExchangeService_optionsForMessages() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Accepted",
			params:     []string{"subject"},
			checkError: assert.NoError,
		},
		{
			name:       "Multiple Accepted",
			params:     []string{"webLink", "parentFolderId"},
			checkError: assert.NoError,
		},
		{
			name:       "Incorrect param",
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

// TestExchangeService_optionsForFolders ensures that approved query options
// are added to the RequestBuildConfiguration. Expected will always be +1
// on than the input as "id" are always included within the select parameters
func (suite *ExchangeServiceSuite) TestExchangeService_optionsForFolders() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
		expected   int
	}{
		{
			name:       "Accepted",
			params:     []string{"displayName"},
			checkError: assert.NoError,
			expected:   2,
		},
		{
			name:       "Multiple Accepted",
			params:     []string{"displayName", "parentFolderId"},
			checkError: assert.NoError,
			expected:   3,
		},
		{
			name:       "Incorrect param",
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
			options, err := optionsForContacts(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(test.expected, len(options.QueryParameters.Select))
			}
		})
	}
}

// TestExchangeService_optionsForContacts similar to TestExchangeService_optionsForFolders
func (suite *ExchangeServiceSuite) TestExchangeService_optionsForContacts() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
		expected   int
	}{
		{
			name:       "Accepted",
			params:     []string{"displayName"},
			checkError: assert.NoError,
			expected:   2,
		},
		{
			name:       "Multiple Accepted",
			params:     []string{"displayName", "parentFolderId"},
			checkError: assert.NoError,
			expected:   3,
		},
		{
			name:       "Incorrect param",
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

// TestExchangeService_SetupExchangeCollection ensures that the helper
// function SetupExchangeCollectionVars returns a non-nil variable for returns
// in regards to the selector.ExchangeScope.
func (suite *ExchangeServiceSuite) TestExchangeService_SetupExchangeCollection() {
	userID := tester.M365UserID(suite.T())
	sel := selectors.NewExchangeBackup()
	sel.Include(sel.Users([]string{userID}))
	eb, err := sel.ToExchangeBackup()
	require.NoError(suite.T(), err)
	scopes := eb.Scopes()

	for _, test := range scopes {
		suite.T().Run(test.Category().String(), func(t *testing.T) {
			discriminateFunc, graphQuery, iterFunc, err := SetupExchangeCollectionVars(test)
			if test.Category() == selectors.ExchangeMailFolder ||
				test.Category() == selectors.ExchangeContactFolder {
				assert.NoError(t, err)
				assert.NotNil(t, discriminateFunc)
				assert.NotNil(t, graphQuery)
				assert.NotNil(t, iterFunc)
			}
		})
	}
}

// TestExchangeService_GraphQueryFunctions verifies if Query functions APIs
// through Microsoft Graph are functional
func (suite *ExchangeServiceSuite) TestExchangeService_GraphQueryFunctions() {
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
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			response, err := test.function(suite.es, userID)
			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}

// TestExchangeService_IterativeFunctions verifies that GraphQuery to Iterate
// functions are valid for current versioning of msgraph-go-sdk
func (suite *ExchangeServiceSuite) TestExchangeService_IterativeFunctions() {
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
			iterativeFunction: IterateSelectAllMessagesForCollections,
			scope:             mailScope,
			transformer:       models.CreateMessageCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Contacts Iterative Check",
			queryFunction:     GetAllContactsForUser,
			iterativeFunction: IterateAllContactsForCollection,
			scope:             contactScope,
			transformer:       models.CreateContactFromDiscriminatorValue,
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
				"testingTenant",
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
