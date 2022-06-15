package mockconnector

import (
	"bytes"
	"io"

	"github.com/google/uuid"

	"github.com/alcionai/corso/internal/connector"
)

// MockExchangeDataCollection represents a mock exchange mailbox
type MockExchangeDataCollection struct {
	fullPath     []string
	messageCount int
	messagesRead int
}

var (
	_ connector.DataCollection = &MockExchangeDataCollection{}
	_ connector.DataStream     = &MockExchangeData{}
)

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated
func NewMockExchangeDataCollection(pathRepresentation []string, numMessagesToReturn int) *MockExchangeDataCollection {
	collection := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		messagesRead: 0,
	}
	return collection
}

func (medc *MockExchangeDataCollection) FullPath() []string {
	return append([]string{}, medc.fullPath...)
}

// NextItem returns either the next item in the collection or an error if one occurred.
// If not more items are available in the collection, returns (nil, nil).
func (medc *MockExchangeDataCollection) NextItem() (connector.DataStream, error) {
	if medc.messagesRead < medc.messageCount {
		medc.messagesRead++
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		m := []byte("test message")
		return &MockExchangeData{uuid.NewString(), io.NopCloser(bytes.NewReader(m))}, nil
	}
	return nil, io.EOF
}

// ExchangeData represents a single item retrieved from exchange
type MockExchangeData struct {
	id     string
	reader io.ReadCloser
}

func (med *MockExchangeData) UUID() string {
	return med.id
}

func (med *MockExchangeData) ToReader() io.ReadCloser {
	return med.reader
}
