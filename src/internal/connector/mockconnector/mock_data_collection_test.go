package mockconnector_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type MockExchangeCollectionSuite struct {
	tester.Suite
}

func TestMockExchangeCollectionSuite(t *testing.T) {
	suite.Run(t, &MockExchangeCollectionSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MockExchangeCollectionSuite) TestMockExchangeCollection() {
	ctx, flush := tester.NewContext()
	defer flush()

	mdc := mockconnector.NewMockExchangeCollection(nil, nil, 2)
	messagesRead := 0

	for item := range mdc.Items(ctx, fault.New(true)) {
		_, err := io.ReadAll(item.ToReader())
		assert.NoError(suite.T(), err, clues.ToCore(err))
		messagesRead++
	}

	assert.Equal(suite.T(), 2, messagesRead)
}

func (suite *MockExchangeCollectionSuite) TestMockExchangeCollectionItemSize() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	mdc := mockconnector.NewMockExchangeCollection(nil, nil, 2)
	mdc.Data[1] = []byte("This is some buffer of data so that the size is different than the default")

	for item := range mdc.Items(ctx, fault.New(true)) {
		buf, err := io.ReadAll(item.ToReader())
		assert.NoError(t, err, clues.ToCore(err))

		assert.Implements(t, (*data.StreamSize)(nil), item)
		s := item.(data.StreamSize)
		assert.Equal(t, int64(len(buf)), s.Size())
	}
}

// NewExchangeCollectionMail_Hydration tests that mock exchange mail data collection can be used for restoration
// functions by verifying no failures on (de)serializing steps using kiota serialization library
func (suite *MockExchangeCollectionSuite) TestMockExchangeCollection_NewExchangeCollectionMail_Hydration() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	mdc := mockconnector.NewMockExchangeCollection(nil, nil, 3)
	buf := &bytes.Buffer{}

	for stream := range mdc.Items(ctx, fault.New(true)) {
		_, err := buf.ReadFrom(stream.ToReader())
		assert.NoError(t, err, clues.ToCore(err))

		byteArray := buf.Bytes()
		something, err := support.CreateFromBytes(byteArray, models.CreateMessageFromDiscriminatorValue)
		assert.NoError(t, err, clues.ToCore(err))
		assert.NotNil(t, something)
	}
}

type MockExchangeDataSuite struct {
	tester.Suite
}

func TestMockExchangeDataSuite(t *testing.T) {
	suite.Run(t, &MockExchangeDataSuite{Suite: tester.NewUnitSuite(t)})
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
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, id, test.reader.UUID())
			buf, err := io.ReadAll(test.reader.ToReader())

			test.check(t, err, clues.ToCore(err))
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
				require.NoError(t, err, clues.ToCore(err))

				bytes, err := writer.GetSerializedContent()
				require.NoError(t, err, clues.ToCore(err))

				_, err = support.CreateListFromBytes(bytes)

				return err
			},
		},
		{
			name: "SharePoint: List 6 Items",
			transformation: func(t *testing.T) error {
				bytes, err := mockconnector.GetMockListBytes(subject)
				require.NoError(t, err, clues.ToCore(err))
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
		suite.Run(test.name, func() {
			t := suite.T()

			err := test.transformation(t)
			assert.NoError(t, err, clues.ToCore(err))
		})
	}
}
