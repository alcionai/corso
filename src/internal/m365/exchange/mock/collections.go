package mock

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// DataCollection represents a mock exchange mailbox
type DataCollection struct {
	fullPath     path.Path
	LocPath      path.Path
	messageCount int
	Data         [][]byte
	Names        []string
	ModTimes     []time.Time
	ColState     data.CollectionState
	PrevPath     path.Path
	DeletedItems []bool
	DoNotMerge   bool
}

var (
	_ data.BackupCollection = &DataCollection{}
	_ data.Stream           = &Data{}
	_ data.StreamInfo       = &Data{}
	_ data.StreamSize       = &Data{}
)

func (medc DataCollection) FullPath() path.Path { return medc.fullPath }

func (medc DataCollection) LocationPath() *path.Builder {
	if medc.LocPath == nil {
		return nil
	}

	return path.Builder{}.Append(medc.LocPath.Folders()...)
}

func (medc DataCollection) PreviousPath() path.Path     { return medc.PrevPath }
func (medc DataCollection) State() data.CollectionState { return medc.ColState }
func (medc DataCollection) DoNotMergeItems() bool       { return medc.DoNotMerge }

// NewCollection creates an data collection that will return the specified number of
// mock messages when iterated. Exchange type mail
func NewCollection(
	storagePath path.Path,
	locationPath path.Path,
	numMessagesToReturn int,
) *DataCollection {
	c := &DataCollection{
		fullPath:     storagePath,
		LocPath:      locationPath,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
		ModTimes:     []time.Time{},
		DeletedItems: []bool{},
	}
	baseTime := time.Now()

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, MessageBytes("From: NewMockCollection"))
		c.Names = append(c.Names, uuid.NewString())
		c.ModTimes = append(c.ModTimes, baseTime.Add(1*time.Hour))
		c.DeletedItems = append(c.DeletedItems, false)
	}

	return c
}

// NewContactCollection creates an data collection that will return the specified number of
// mock messages when iterated. Exchange type mail
func NewContactCollection(pathRepresentation path.Path, numMessagesToReturn int) *DataCollection {
	c := &DataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
	}

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
		c.Data = append(c.Data, ContactBytes(middleNames[rand.Intn(len(middleNames))]))
		c.Names = append(c.Names, uuid.NewString())
	}

	return c
}

// Items returns a channel that has the next items in the collection. The
// channel is closed when there are no more items available.
func (medc *DataCollection) Items(
	ctx context.Context,
	_ *fault.Bus, // unused
) <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		for i := 0; i < medc.messageCount; i++ {
			res <- &Data{
				ID:           medc.Names[i],
				Reader:       io.NopCloser(bytes.NewReader(medc.Data[i])),
				size:         int64(len(medc.Data[i])),
				modifiedTime: medc.ModTimes[i],
				deleted:      medc.DeletedItems[i],
			}
		}
	}()

	return res
}

// TODO: move to data/mock for service-agnostic mocking
// Data represents a single item retrieved from exchange
type Data struct {
	ID           string
	Reader       io.ReadCloser
	ReadErr      error
	size         int64
	modifiedTime time.Time
	deleted      bool
}

func (med *Data) UUID() string       { return med.ID }
func (med *Data) Deleted() bool      { return med.deleted }
func (med *Data) Size() int64        { return med.size }
func (med *Data) ModTime() time.Time { return med.modifiedTime }

func (med *Data) ToReader() io.ReadCloser {
	if med.ReadErr != nil {
		return io.NopCloser(errReader{med.ReadErr})
	}

	return med.Reader
}

func (med *Data) Info() details.ItemInfo {
	return details.ItemInfo{
		Exchange: &details.ExchangeInfo{
			ItemType: details.ExchangeMail,
			Sender:   "foo@bar.com",
			Subject:  "Hello world!",
			Received: time.Now(),
		},
	}
}

type errReader struct {
	readErr error
}

func (er errReader) Read([]byte) (int, error) {
	return 0, er.readErr
}
