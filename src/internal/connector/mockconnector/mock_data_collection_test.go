package mockconnector_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/connector/support"
)

type MockExchangeCollectionSuite struct {
	suite.Suite
}

func TestMockExchangeCollectionSuite(t *testing.T) {
	suite.Run(t, new(MockExchangeCollectionSuite))
}

func (suite *MockExchangeCollectionSuite) TestMockExchangeCollection() {
	mdc := mockconnector.NewMockExchangeCollection([]string{"foo", "bar"}, 2)

	messagesRead := 0

	for item := range mdc.Items() {
		_, err := ioutil.ReadAll(item.ToReader())
		assert.NoError(suite.T(), err)
		messagesRead++
	}
	assert.Equal(suite.T(), 2, messagesRead)
}

// NewExchangeCollectionMail_Hydration tests that mock exchange mail data collection can be used for restoration
// functions by verifying no failures on (de)serializing steps using kiota serialization library
func (suite *MockExchangeCollectionSuite) TestMockExchangeCollection_NewExchangeCollectionMail_Hydration() {
	t := suite.T()
	mdc := mockconnector.NewMockExchangeCollection([]string{"foo", "bar"}, 3)
	var (
		byteArray []byte
	)
	buf := &bytes.Buffer{}
	for stream := range mdc.Items() {
		_, err := buf.ReadFrom(stream.ToReader())
		assert.NoError(t, err)
		byteArray = buf.Bytes()
		something, err := support.CreateFromBytes(byteArray, models.CreateMessageFromDiscriminatorValue)
		assert.NoError(t, err)
		assert.NotNil(t, something)
	}
}
