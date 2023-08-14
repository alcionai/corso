package mock

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
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

var _ data.BackupCollection = &DataCollection{}

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
) <-chan data.Item {
	res := make(chan data.Item)

	go func() {
		defer close(res)

		for i := 0; i < medc.messageCount; i++ {
			res <- &dataMock.Item{
				ItemID:       medc.Names[i],
				Reader:       io.NopCloser(bytes.NewReader(medc.Data[i])),
				ItemSize:     int64(len(medc.Data[i])),
				ModifiedTime: medc.ModTimes[i],
				DeletedFlag:  medc.DeletedItems[i],
				ItemInfo:     StubMailInfo(),
			}
		}
	}()

	return res
}

func StubMailInfo() details.ItemInfo {
	return details.ItemInfo{
		Exchange: &details.ExchangeInfo{
			ItemType: details.ExchangeMail,
			Sender:   "foo@bar.com",
			Subject:  "Hello world!",
			Received: time.Now(),
		},
	}
}
