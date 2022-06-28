package mockconnector_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/mockconnector"
)

type MockExchangeDataCollectionSuite struct {
	suite.Suite
}

func TestMockExchangeDataCollectionSuite(t *testing.T) {
	suite.Run(t, new(MockExchangeDataCollectionSuite))
}

func (suite *MockExchangeDataCollectionSuite) TestMockExchangeDataCollection() {
	mdc := mockconnector.NewMockExchangeDataCollection([]string{"foo", "bar"}, 2)

	messagesRead := 0

	for item := range mdc.Items() {
		_, err := ioutil.ReadAll(item.ToReader())
		assert.NoError(suite.T(), err)
		messagesRead++
	}
	assert.Equal(suite.T(), 2, messagesRead)
}
