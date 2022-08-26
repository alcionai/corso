package mockconnector_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	buf := &bytes.Buffer{}

	for stream := range mdc.Items() {
		_, err := buf.ReadFrom(stream.ToReader())
		assert.NoError(t, err)

		byteArray := buf.Bytes()
		something, err := support.CreateFromBytes(byteArray, models.CreateMessageFromDiscriminatorValue)
		assert.NoError(t, err)
		assert.NotNil(t, something)
	}
}

type MockExchangeDataSuite struct {
	suite.Suite
}

func TestMockExchangeDataSuite(t *testing.T) {
	suite.Run(t, new(MockExchangeDataSuite))
}

func (suite *MockExchangeDataSuite) TestMockExchangeData() {
	data := []byte("foo")
	id := "bar"

	table := []struct {
		name   string
		reader *mockconnector.MockExchangeData
		check  require.ErrorAssertionFunc
	}{
		{
			name: "NoError",
			reader: &mockconnector.MockExchangeData{
				ID:     id,
				Reader: io.NopCloser(bytes.NewReader(data)),
			},
			check: require.NoError,
		},
		{
			name: "Error",
			reader: &mockconnector.MockExchangeData{
				ID:      id,
				ReadErr: assert.AnError,
			},
			check: require.Error,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, id, test.reader.UUID())
			buf, err := ioutil.ReadAll(test.reader.ToReader())

			test.check(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, data, buf)
		})
	}
}
