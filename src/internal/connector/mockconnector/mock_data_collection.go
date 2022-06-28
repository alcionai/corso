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
	}
	return collection
}

func (medc *MockExchangeDataCollection) FullPath() []string {
	return append([]string{}, medc.fullPath...)
}

// Items returns a channel that has the next items in the collection. The
// channel is closed when there are no more items available.
func (medc *MockExchangeDataCollection) Items() <-chan connector.DataStream {
	res := make(chan connector.DataStream)

	go func() {
		defer close(res)

		for i := 0; i < medc.messageCount; i++ {
			// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
			m := []byte("test message")
			res <- &MockExchangeData{uuid.NewString(), io.NopCloser(bytes.NewReader(m))}
		}
	}()

	return res
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
