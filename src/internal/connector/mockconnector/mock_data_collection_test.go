package mockconnector_test

import (
	"bytes"
	"io"
	"testing"

	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
)

type MockExchangeCollectionSuite struct {
	suite.Suite
}

func TestMockExchangeCollectionSuite(t *testing.T) {
	suite.Run(t, new(MockExchangeCollectionSuite))
}

func (suite *MockExchangeCollectionSuite) TestMockExchangeCollection() {
	mdc := mockconnector.NewMockExchangeCollection(nil, 2)

	messagesRead := 0

	for item := range mdc.Items() {
		_, err := io.ReadAll(item.ToReader())
		assert.NoError(suite.T(), err)
		messagesRead++
	}

	assert.Equal(suite.T(), 2, messagesRead)
}

func (suite *MockExchangeCollectionSuite) TestMockExchangeCollectionItemSize() {
	t := suite.T()
	mdc := mockconnector.NewMockExchangeCollection(nil, 2)

	mdc.Data[1] = []byte("This is some buffer of data so that the size is different than the default")

	for item := range mdc.Items() {
		buf, err := io.ReadAll(item.ToReader())
		assert.NoError(t, err)

		assert.Implements(t, (*data.StreamSize)(nil), item)
		s := item.(data.StreamSize)
		assert.Equal(t, int64(len(buf)), s.Size())
	}
}

// NewExchangeCollectionMail_Hydration tests that mock exchange mail data collection can be used for restoration
// functions by verifying no failures on (de)serializing steps using kiota serialization library
func (suite *MockExchangeCollectionSuite) TestMockExchangeCollection_NewExchangeCollectionMail_Hydration() {
	t := suite.T()
	mdc := mockconnector.NewMockExchangeCollection(nil, 3)
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
	itemData := []byte("foo")
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
				Reader: io.NopCloser(bytes.NewReader(itemData)),
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
			buf, err := io.ReadAll(test.reader.ToReader())

			test.check(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, itemData, buf)
		})
	}
}

func (suite *MockExchangeDataSuite) TestMockByteHydration() {
	subject := "Mock Hydration"
	tests := []struct {
		name           string
		transformation func(t *testing.T) error
	}{
		{
			name: "Message Bytes",
			transformation: func(t *testing.T) error {
				bytes := mockconnector.GetMockMessageBytes(subject)
				_, err := support.CreateMessageFromBytes(bytes)
				return err
			},
		},
		{
			name: "Event Message Response: Regression",
			transformation: func(t *testing.T) error {
				bytes := mockconnector.GetMockEventMessageResponse(subject)
				_, err := support.CreateMessageFromBytes(bytes)
				return err
			},
		},
		{
			name: "Event Message Request: Regression",
			transformation: func(t *testing.T) error {
				bytes := mockconnector.GetMockEventMessageRequest(subject)
				_, err := support.CreateMessageFromBytes(bytes)
				return err
			},
		},
		{
			name: "Contact Bytes",
			transformation: func(t *testing.T) error {
				bytes := mockconnector.GetMockContactBytes(subject)
				_, err := support.CreateContactFromBytes(bytes)
				return err
			},
		},
		{
			name: "Event No Attendees Bytes",
			transformation: func(t *testing.T) error {
				bytes := mockconnector.GetDefaultMockEventBytes(subject)
				_, err := support.CreateEventFromBytes(bytes)
				return err
			},
		},
		{
			name: "Event w/ Attendees Bytes",
			transformation: func(t *testing.T) error {
				bytes := mockconnector.GetMockEventWithAttendeesBytes(subject)
				_, err := support.CreateEventFromBytes(bytes)
				return err
			},
		},
		{
			name: "SharePoint: List Empty",
			transformation: func(t *testing.T) error {
				emptyMap := make(map[string]string)
				temp := mockconnector.GetMockList(subject, "Artist", emptyMap)
				writer := kioser.NewJsonSerializationWriter()
				err := writer.WriteObjectValue("", temp)
				require.NoError(t, err)

				bytes, err := writer.GetSerializedContent()
				require.NoError(suite.T(), err)

				_, err = support.CreateListFromBytes(bytes)

				return err
			},
		},
		{
			name: "SharePoint: List 6 Items",
			transformation: func(t *testing.T) error {
				bytes, err := mockconnector.GetMockListBytes(subject)
				require.NoError(suite.T(), err)
				_, err = support.CreateListFromBytes(bytes)
				return err
			},
		},
		{
			name: "SharePoint: Page",
			transformation: func(t *testing.T) error {
				bytes := mockconnector.GetMockPage(subject)
				_, err := support.CreatePageFromBytes(bytes)

				return err
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			err := test.transformation(t)
			assert.NoError(t, err)
		})
	}
}
