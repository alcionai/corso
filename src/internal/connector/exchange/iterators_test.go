package exchange

import (
	"context"
	"testing"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type ExchangeIteratorSuite struct {
	suite.Suite
}

func TestExchangeIteratorSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ExchangeIteratorSuite))
}

func (suite *ExchangeIteratorSuite) TestDisplayable() {
	t := suite.T()
	bytes := mockconnector.GetMockContactBytes("Displayable")
	contact, err := support.CreateContactFromBytes(bytes)
	require.NoError(t, err)

	aDisplayable, ok := contact.(displayable)
	assert.True(t, ok)
	assert.NotNil(t, aDisplayable.GetId())
	assert.NotNil(t, aDisplayable.GetDisplayName())
}

func (suite *ExchangeIteratorSuite) TestDescendable() {
	t := suite.T()
	bytes := mockconnector.GetMockMessageBytes("Descendable")
	message, err := support.CreateMessageFromBytes(bytes)
	require.NoError(t, err)

	aDescendable, ok := message.(descendable)
	assert.True(t, ok)
	assert.NotNil(t, aDescendable.GetId())
	assert.NotNil(t, aDescendable.GetParentFolderId())
}

func loadService(t *testing.T) *exchangeService {
	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(t, err)

	a := tester.NewM365Account(t)
	require.NoError(t, err)

	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := createService(m365, false)
	require.NoError(t, err)

	return service
}

// TestIterativeFunctions verifies that GraphQuery to Iterate
// functions are valid for current versioning of msgraph-go-sdk
func (suite *ExchangeIteratorSuite) TestIterativeFunctions() {
	var (
		expectedFolderNames     = make(map[string]struct{}, 0)
		ctx                     = context.Background()
		t                       = suite.T()
		mailScope, contactScope selectors.ExchangeScope
		userID                  = tester.M365UserID(t)
		sel                     = selectors.NewExchangeBackup()
		service                 = loadService(t)
	)

	sel.Include(sel.Users([]string{userID}))

	eb, err := sel.ToExchangeBackup()
	require.NoError(suite.T(), err)

	scopes := eb.Scopes()

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
			response, err := test.queryFunction(service, userID)
			require.NoError(t, err)
			// Create Iterator
			pageIterator, err := msgraphgocore.NewPageIterator(response,
				&service.adapter,
				test.transformer)
			require.NoError(t, err)

			qp := graph.QueryParams{
				User:        userID,
				Scope:       test.scope,
				Credentials: service.credentials,
				FailFast:    false,
			}
			// Create collection for iterate test
			collections := make(map[string]*Collection)
			var errs error
			// callbackFunc iterates through all models.Messageable and fills exchange.Collection.jobs[]
			// with corresponding item IDs. New collections are created for each directory
			callbackFunc := test.iterativeFunction(
				ctx,
				qp,
				errs,
				collections,
				nil)

			iterateError := pageIterator.Iterate(callbackFunc)
			assert.NoError(t, iterateError)
			assert.NoError(t, errs)

			// TODO(ashmrtn): Only check Exchange Mail folder names right now because
			// other resolvers aren't implemented. Once they are we can expand these
			// checks, potentially by breaking things out into separate tests per
			// category.
			if !test.scope.IncludesCategory(selectors.ExchangeMail) {
				return
			}

			// Mail Version
			if test.name == "Mail Iterative Check" {
				expectedFolderNames = map[string]struct{}{
					"Inbox":      {},
					"Sent Items": {},
				}
			}

			if test.name == "Folder Iterative Check" {
				expectedFolderNames = map[string]struct{}{
					"Inbox":         {},
					"Sent Items":    {},
					"Deleted Items": {},
				}
			}

			for _, c := range collections {
				// TODO(ashmrtn): Update these checks when collections support path.Path.
				require.Greater(t, len(c.FullPath()), 4)

				folder := c.FullPath()[4]
				if _, ok := expectedFolderNames[folder]; ok {
					delete(expectedFolderNames, folder)
				}
			}

			assert.Empty(t, expectedFolderNames)
		})
	}
}
