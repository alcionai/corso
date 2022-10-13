package exchange

import (
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

	aDisplayable, ok := contact.(graph.Displayable)
	assert.True(t, ok)
	assert.NotNil(t, aDisplayable.GetId())
	assert.NotNil(t, aDisplayable.GetDisplayName())
}

func (suite *ExchangeIteratorSuite) TestDescendable() {
	t := suite.T()
	bytes := mockconnector.GetMockMessageBytes("Descendable")
	message, err := support.CreateMessageFromBytes(bytes)
	require.NoError(t, err)

	aDescendable, ok := message.(graph.Descendable)
	assert.True(t, ok)
	assert.NotNil(t, aDescendable.GetId())
	assert.NotNil(t, aDescendable.GetParentFolderId())
}

func loadService(t *testing.T) *exchangeService {
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := createService(m365, false)
	require.NoError(t, err)

	return service
}

// TestIterativeFunctions verifies that GraphQuery to Iterate
// functions are valid for current versioning of msgraph-go-sdk.
// Tests for mail have been moved to graph_connector_test.go.
func (suite *ExchangeIteratorSuite) TestIterativeFunctions() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t                                   = suite.T()
		mailScope, contactScope, eventScope []selectors.ExchangeScope
		userID                              = tester.M365UserID(t)
		sel                                 = selectors.NewExchangeBackup()
	)

	eb, err := sel.ToExchangeBackup()
	require.NoError(suite.T(), err)

	contactScope = sel.ContactFolders([]string{userID}, []string{DefaultContactFolder})
	eventScope = sel.EventCalendars([]string{userID}, []string{DefaultCalendar})
	mailScope = sel.MailFolders([]string{userID}, []string{DefaultMailFolder})

	eb.Include(contactScope, eventScope, mailScope)

	tests := []struct {
		name              string
		queryFunction     GraphQuery
		iterativeFunction GraphIterateFunc
		scope             selectors.ExchangeScope
		transformer       absser.ParsableFactory
		folderNames       map[string]struct{}
	}{
		{
			name:              "Contacts Iterative Check",
			queryFunction:     GetAllContactFolderNamesForUser,
			iterativeFunction: IterateSelectAllContactsForCollections,
			scope:             contactScope[0],
			transformer:       models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Contact Folder Traversal",
			queryFunction:     GetAllContactFolderNamesForUser,
			iterativeFunction: IterateSelectAllContactsForCollections,
			scope:             contactScope[0],
			transformer:       models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Events Iterative Check",
			queryFunction:     GetAllCalendarNamesForUser,
			iterativeFunction: IterateSelectAllEventsFromCalendars,
			scope:             eventScope[0],
			transformer:       models.CreateCalendarCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Folder Iterative Check Contacts",
			queryFunction:     GetAllContactFolderNamesForUser,
			iterativeFunction: IterateFilterContainersForCollections,
			scope:             contactScope[0],
			transformer:       models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
		}, {
			name:              "Default Contacts Folder",
			queryFunction:     GetDefaultContactFolderForUser,
			iterativeFunction: IterateSelectAllContactsForCollections,
			scope:             contactScope[0],
			transformer:       models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			service := loadService(t)
			response, err := test.queryFunction(ctx, service, userID)
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
				nil,
				nil,
			)

			iterateError := pageIterator.Iterate(ctx, callbackFunc)
			assert.NoError(t, iterateError)
			assert.NoError(t, errs)
		})
	}
}
