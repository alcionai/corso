package mock

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type MockSuite struct {
	tester.Suite
}

func TestMockSuite(t *testing.T) {
	suite.Run(t, &MockSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MockSuite) TestMockExchangeCollection() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	mdc := NewCollection(nil, nil, 2)
	messagesRead := 0

	for item := range mdc.Items(ctx, fault.New(true)) {
		_, err := io.ReadAll(item.ToReader())
		assert.NoError(t, err, clues.ToCore(err))
		messagesRead++
	}

	assert.Equal(t, 2, messagesRead)
}

func (suite *MockSuite) TestMockExchangeCollectionItemSize() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	mdc := NewCollection(nil, nil, 2)
	mdc.Data[1] = []byte("This is some buffer of data so that the size is different than the default")

	for item := range mdc.Items(ctx, fault.New(true)) {
		buf, err := io.ReadAll(item.ToReader())
		assert.NoError(t, err, clues.ToCore(err))

		assert.Implements(t, (*data.ItemSize)(nil), item)
		s := item.(data.ItemSize)
		assert.Equal(t, int64(len(buf)), s.Size())
	}
}

// NewExchangeCollectionMail_Hydration tests that mock exchange mail data collection can be used for restoration
// functions by verifying no failures on (de)serializing steps using kiota serialization library
func (suite *MockSuite) TestMockExchangeCollection_NewExchangeCollectionMail_Hydration() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	mdc := NewCollection(nil, nil, 3)

	for stream := range mdc.Items(ctx, fault.New(true)) {
		buf := &bytes.Buffer{}
		_, err := buf.ReadFrom(stream.ToReader())
		assert.NoError(t, err, clues.ToCore(err))

		byteArray := buf.Bytes()
		something, err := api.BytesToMessageable(byteArray)
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
		reader *dataMock.Item
		check  require.ErrorAssertionFunc
	}{
		{
			name: "NoError",
			reader: &dataMock.Item{
				ItemID:   id,
				Reader:   io.NopCloser(bytes.NewReader(itemData)),
				ItemInfo: StubMailInfo(),
			},
			check: require.NoError,
		},
		{
			name: "Error",
			reader: &dataMock.Item{
				ItemID:   id,
				ReadErr:  assert.AnError,
				ItemInfo: StubMailInfo(),
			},
			check: require.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, id, test.reader.ID())
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
				bytes := MessageBytes(subject)
				_, err := api.BytesToMessageable(bytes)
				return err
			},
		},
		{
			name: "Event Message Response: Regression",
			transformation: func(t *testing.T) error {
				bytes := EventMessageResponse(subject)
				_, err := api.BytesToEventable(bytes)
				return err
			},
		},
		{
			name: "Event Message Request: Regression",
			transformation: func(t *testing.T) error {
				bytes := EventMessageRequest(subject)
				_, err := api.BytesToEventable(bytes)
				return err
			},
		},
		{
			name: "Contact Bytes",
			transformation: func(t *testing.T) error {
				bytes := ContactBytes(subject)
				_, err := api.BytesToContactable(bytes)
				return err
			},
		},
		{
			name: "Event No Attendees Bytes",
			transformation: func(t *testing.T) error {
				bytes := EventBytes(subject)
				_, err := api.BytesToEventable(bytes)
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
