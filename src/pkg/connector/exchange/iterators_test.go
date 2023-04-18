package exchange

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	exchMock "github.com/alcionai/corso/src/pkg/connector/exchange/mock"
	"github.com/alcionai/corso/src/pkg/connector/graph"
	"github.com/alcionai/corso/src/pkg/connector/support"
)

type ExchangeIteratorSuite struct {
	tester.Suite
}

func TestExchangeIteratorSuite(t *testing.T) {
	suite.Run(t, &ExchangeIteratorSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExchangeIteratorSuite) TestDisplayable() {
	t := suite.T()
	bytes := exchMock.ContactBytes("Displayable")
	contact, err := support.CreateContactFromBytes(bytes)
	require.NoError(t, err, clues.ToCore(err))

	aDisplayable, ok := contact.(graph.Displayable)
	assert.True(t, ok)
	assert.NotNil(t, aDisplayable.GetId())
	assert.NotNil(t, aDisplayable.GetDisplayName())
}

func (suite *ExchangeIteratorSuite) TestDescendable() {
	t := suite.T()
	bytes := exchMock.MessageBytes("Descendable")
	message, err := support.CreateMessageFromBytes(bytes)
	require.NoError(t, err, clues.ToCore(err))

	aDescendable, ok := message.(graph.Descendable)
	assert.True(t, ok)
	assert.NotNil(t, aDescendable.GetId())
	assert.NotNil(t, aDescendable.GetParentFolderId())
}
