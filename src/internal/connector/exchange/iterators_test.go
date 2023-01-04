package exchange

import (
	"testing"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
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
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

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

// TestCollectionFunctions verifies ability to gather
// containers functions are valid for current versioning of msgraph-go-sdk.
// Tests for mail have been moved to graph_connector_test.go.
// exchange.Mail uses a sequential delta function.
// TODO: Add exchange.Mail when delta iterator functionality implemented
func (suite *ExchangeIteratorSuite) TestCollectionFunctions() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t                                   = suite.T()
		mailScope, contactScope, eventScope []selectors.ExchangeScope
		userID                              = tester.M365UserID(t)
		users                               = []string{userID}
		sel                                 = selectors.NewExchangeBackup(users)
	)

	eb, err := sel.ToExchangeBackup()
	require.NoError(suite.T(), err)

	contactScope = sel.ContactFolders(users, []string{DefaultContactFolder}, selectors.PrefixMatch())
	eventScope = sel.EventCalendars(users, []string{DefaultCalendar}, selectors.PrefixMatch())
	mailScope = sel.MailFolders(users, []string{DefaultMailFolder}, selectors.PrefixMatch())

	eb.Include(contactScope, eventScope, mailScope)

	tests := []struct {
		name              string
		queryFunc         api.GraphQuery
		scope             selectors.ExchangeScope
		iterativeFunction func(
			container map[string]graph.Container,
			aFilter string,
			errUpdater func(string, error)) func(any) bool
		transformer absser.ParsableFactory
	}{
		{
			name:              "Contacts Iterative Check",
			queryFunc:         api.GetAllContactFolderNamesForUser,
			transformer:       models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
			iterativeFunction: IterativeCollectContactContainers,
		},
		{
			name:              "Events Iterative Check",
			queryFunc:         api.GetAllCalendarNamesForUser,
			transformer:       models.CreateCalendarCollectionResponseFromDiscriminatorValue,
			iterativeFunction: IterativeCollectCalendarContainers,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			a := tester.NewM365Account(t)
			m365, err := a.M365Config()
			require.NoError(t, err)

			service, err := createService(m365)
			require.NoError(t, err)

			response, err := test.queryFunc(ctx, service, userID)
			require.NoError(t, err)

			// Iterator Creation
			pageIterator, err := msgraphgocore.NewPageIterator(
				response,
				service.Adapter(),
				test.transformer)
			require.NoError(t, err)

			// Create collection for iterate test
			collections := make(map[string]graph.Container)

			var errs error

			errUpdater := func(id string, err error) {
				errs = support.WrapAndAppend(id, err, errs)
			}

			// callbackFunc iterates through all models.Messageable and fills exchange.Collection.added[]
			// with corresponding item IDs. New collections are created for each directory
			callbackFunc := test.iterativeFunction(collections, "", errUpdater)

			iterateError := pageIterator.Iterate(ctx, callbackFunc)
			assert.NoError(t, iterateError)
			assert.NoError(t, errs)
		})
	}
}
