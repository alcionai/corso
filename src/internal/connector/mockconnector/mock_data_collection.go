package mockconnector

import (
	"bytes"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
)

// MockExchangeDataCollection represents a mock exchange mailbox
type MockExchangeDataCollection struct {
	fullPath     []string
	messageCount int
	Data         [][]byte
	Names        []string
}

var (
	_ data.Collection = &MockExchangeDataCollection{}
	_ data.Stream     = &MockExchangeData{}
	_ data.StreamInfo = &MockExchangeData{}
)

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated
func NewMockExchangeDataCollection(pathRepresentation []string, numMessagesToReturn int) *MockExchangeDataCollection {
	c := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
	}

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, []byte("test message"))
		c.Names = append(c.Names, uuid.NewString())
	}
	return c
}

func (medc *MockExchangeDataCollection) FullPath() []string {
	return append([]string{}, medc.fullPath...)
}

// Items returns a channel that has the next items in the collection. The
// channel is closed when there are no more items available.
func (medc *MockExchangeDataCollection) Items() <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)
		for i := 0; i < medc.messageCount; i++ {
			res <- &MockExchangeData{
				medc.Names[i],
				io.NopCloser(bytes.NewReader(medc.Data[i])),
			}
		}
	}()

	return res
}

// ExchangeData represents a single item retrieved from exchange
type MockExchangeData struct {
	ID     string
	Reader io.ReadCloser
}

func (med *MockExchangeData) UUID() string {
	return med.ID
}

func (med *MockExchangeData) ToReader() io.ReadCloser {
	return med.Reader
}

func (med *MockExchangeData) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: &details.ExchangeInfo{Sender: "foo@bar.com", Subject: "Hello world!", Received: time.Now()}}
}
