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
		ctx                                 = context.Background()
		t                                   = suite.T()
		mailScope, contactScope, eventScope selectors.ExchangeScope
		userID                              = tester.M365UserID(t)
		sel                                 = selectors.NewExchangeBackup()
		service                             = loadService(t)
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

		if scope.IncludesCategory(selectors.ExchangeEvent) {
			eventScope = scope
		}
	}

	tests := []struct {
		name              string
		queryFunction     GraphQuery
		iterativeFunction GraphIterateFunc
		scope             selectors.ExchangeScope
		transformer       absser.ParsableFactory
		folderNames       map[string]struct{}
	}{
		{
			name:              "Mail Iterative Check",
			queryFunction:     GetAllMessagesForUser,
			iterativeFunction: IterateSelectAllDescendablesForCollections,
			scope:             mailScope,
			transformer:       models.CreateMessageCollectionResponseFromDiscriminatorValue,
			folderNames: map[string]struct{}{
				"Inbox":      {},
				"Sent Items": {},
			},
		}, {
			name:              "Contacts Iterative Check",
			queryFunction:     GetAllContactsForUser,
			iterativeFunction: IterateSelectAllDescendablesForCollections,
			scope:             contactScope,
			transformer:       models.CreateContactFromDiscriminatorValue,
		}, {
			name:              "Contact Folder Traversal",
			queryFunction:     GetAllContactFolderNamesForUser,
			iterativeFunction: IterateSelectAllContactsForCollections,
			scope:             contactScope,
			transformer:       models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Events Iterative Check",
			queryFunction:     GetAllCalendarNamesForUser,
			iterativeFunction: IterateSelectAllEventsFromCalendars,
			scope:             eventScope,
			transformer:       models.CreateCalendarCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Folder Iterative Check Mail",
			queryFunction:     GetAllFolderNamesForUser,
			iterativeFunction: IterateFilterFolderDirectoriesForCollections,
			scope:             mailScope,
			transformer:       models.CreateMailFolderCollectionResponseFromDiscriminatorValue,
			folderNames: map[string]struct{}{
				"Inbox":         {},
				"Sent Items":    {},
				"Deleted Items": {},
			},
		}, {
			name:              "Folder Iterative Check Contacts",
			queryFunction:     GetAllContactFolderNamesForUser,
			iterativeFunction: IterateFilterFolderDirectoriesForCollections,
			scope:             contactScope,
			transformer:       models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
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
			errUpdater := func(id string, err error) {
				errs = support.WrapAndAppend(id, err, errs)
			}
			// callbackFunc iterates through all models.Messageable and fills exchange.Collection.jobs[]
			// with corresponding item IDs. New collections are created for each directory
			callbackFunc := test.iterativeFunction(
				ctx,
				qp,
				errUpdater,
				collections,
				nil)

			iterateError := pageIterator.Iterate(callbackFunc)
			assert.NoError(t, iterateError)
			assert.NoError(t, errs)

			// TODO(ashmrtn): Only check Exchange Mail folder names right now because
			// other resolvers aren't implemented. Once they are we can expand these
			// checks, potentially by breaking things out into separate tests per
			// category.
			for _, dad := range collections {
				t.Logf("Collections: %v\n", dad.FullPath().Folder())
			}
			if !test.scope.IncludesCategory(selectors.ExchangeMail) {
				return
			}

			for _, c := range collections {
				require.NotEmpty(t, c.FullPath().Folder())

				folder := c.FullPath().Folder()
				if _, ok := test.folderNames[folder]; ok {
					delete(test.folderNames, folder)
				}
			}

			assert.Empty(t, test.folderNames)
		})
	}
}
