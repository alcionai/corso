package exchange

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
	contact, err := api.BytesToContactable(bytes)
	require.NoError(t, err, clues.ToCore(err))

	aDisplayable, ok := contact.(graph.Displayable)
	assert.True(t, ok)
	assert.NotNil(t, aDisplayable.GetId())
	assert.NotNil(t, aDisplayable.GetDisplayName())
}

func (suite *ExchangeIteratorSuite) TestDescendable() {
	t := suite.T()
	bytes := exchMock.MessageBytes("Descendable")
	message, err := api.BytesToMessageable(bytes)
	require.NoError(t, err, clues.ToCore(err))

	aDescendable, ok := message.(graph.Descendable)
	assert.True(t, ok)
	assert.NotNil(t, aDescendable.GetId())
	assert.NotNil(t, aDescendable.GetParentFolderId())
}
