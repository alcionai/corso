package mockconnector

import (
	"bytes"
	"io"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

// MockExchangeDataCollection represents a mock exchange mailbox
type MockExchangeDataCollection struct {
	fullPath     path.Path
	messageCount int
	Data         [][]byte
	Names        []string
	ModTimes     []time.Time
}

var (
	_ data.Collection = &MockExchangeDataCollection{}
	_ data.Stream     = &MockExchangeData{}
	_ data.StreamInfo = &MockExchangeData{}
	_ data.StreamSize = &MockExchangeData{}
)

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated. Exchange type mail
func NewMockExchangeCollection(pathRepresentation path.Path, numMessagesToReturn int) *MockExchangeDataCollection {
	c := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
		ModTimes:     []time.Time{},
	}
	baseTime := time.Now()

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, GetMockMessageBytes("From: NewMockExchangeCollection"))
		c.Names = append(c.Names, uuid.NewString())
		c.ModTimes = append(c.ModTimes, baseTime.Add(1*time.Hour))
	}

	return c
}

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated. Exchange type mail
func NewMockContactCollection(pathRepresentation path.Path, numMessagesToReturn int) *MockExchangeDataCollection {
	c := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
	}

	rand.Seed(time.Now().UnixNano())

	middleNames := []string{
		"Argon",
		"Bernard",
		"Carleton",
		"Daphenius",
		"Ernesto",
		"Farraday",
		"Ghimley",
		"Irgot",
		"Jannes",
		"Knox",
		"Levi",
		"Milton",
	}

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, GetMockContactBytes(middleNames[rand.Intn(len(middleNames))]))
		c.Names = append(c.Names, uuid.NewString())
	}

	return c
}

func (medc *MockExchangeDataCollection) FullPath() path.Path {
	return medc.fullPath
}

// TODO(ashmrtn): May want to allow setting this in the future for testing.
func (medc MockExchangeDataCollection) PreviousPath() path.Path {
	return nil
}

// TODO(ashmrtn): May want to allow setting this in the future for testing.
func (medc MockExchangeDataCollection) State() data.CollectionState {
	return data.NewState
}

// Items returns a channel that has the next items in the collection. The
// channel is closed when there are no more items available.
func (medc *MockExchangeDataCollection) Items() <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		for i := 0; i < medc.messageCount; i++ {
			res <- &MockExchangeData{
				ID:           medc.Names[i],
				Reader:       io.NopCloser(bytes.NewReader(medc.Data[i])),
				size:         int64(len(medc.Data[i])),
				modifiedTime: medc.ModTimes[i],
			}
		}
	}()

	return res
}

// ExchangeData represents a single item retrieved from exchange
type MockExchangeData struct {
	ID           string
	Reader       io.ReadCloser
	ReadErr      error
	size         int64
	modifiedTime time.Time
}

func (med *MockExchangeData) UUID() string {
	return med.ID
}

// TODO(ashmrtn): May want to allow setting this in the future for testing.
func (med MockExchangeData) Deleted() bool {
	return false
}

func (med *MockExchangeData) ToReader() io.ReadCloser {
	if med.ReadErr != nil {
		return io.NopCloser(errReader{med.ReadErr})
	}

	return med.Reader
}

func (med *MockExchangeData) Info() details.ItemInfo {
	return details.ItemInfo{
		Exchange: &details.ExchangeInfo{
			Sender:   "foo@bar.com",
			Subject:  "Hello world!",
			Received: time.Now(),
		},
	}
}

func (med *MockExchangeData) Size() int64 {
	return med.size
}

func (med *MockExchangeData) ModTime() time.Time {
	return med.modifiedTime
}

type errReader struct {
	readErr error
}

func (er errReader) Read([]byte) (int, error) {
	return 0, er.readErr
}
