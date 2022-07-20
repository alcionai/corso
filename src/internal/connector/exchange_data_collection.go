package connector

import (
	"bytes"
	"io"
	"sort"
	"strings"

	"github.com/alcionai/corso/pkg/backup"
)

const (
	maximumMessages = 2500
)

// A DataCollection represents a collection of data of the
// same type (e.g. mail)
type DataCollection interface {
	// Items returns a channel from which items in the collection can be read.
	// Each returned struct contains the next item in the collection
	// The channel is closed when there are no more items in the collection or if
	// an unrecoverable error caused an early termination in the sender.
	Items() <-chan DataStream
	// FullPath returns a slice of strings that act as metadata tags for this
	// DataCollection. Returned items should be ordered from most generic to least
	// generic. For example, a DataCollection for emails from a specific user
	// would be {"<tenant id>", "<user ID>", "emails"}.
	FullPath() []string
}

// DataStream represents a single item within a DataCollection
// that can be consumed as a stream (it embeds io.Reader)
type DataStream interface {
	// ToReader returns an io.Reader for the DataStream
	ToReader() io.ReadCloser
	// UUID provides a unique identifier for this data
	UUID() string
}

// DataStreamInfo is used to provide service specific
// information about the DataStream
type DataStreamInfo interface {
	Info() backup.ItemInfo
}

var _ DataCollection = &ExchangeDataCollection{}
var _ DataStream = &ExchangeData{}
var _ DataStreamInfo = &ExchangeData{}

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type ExchangeDataCollection struct {
	// M365 user
	user string
	data chan DataStream
	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	fullPath []string
}

// SortDataCollections helper method that sorts the collection by the last item in the FullPath string literal
func SortDataCollections(dcs []DataCollection) {
	sort.SliceStable(dcs, func(i, j int) bool {
		a := dcs[i].FullPath()
		b := dcs[j].FullPath()
		return strings.ToLower(a[len(a)-1]) < strings.ToLower(b[len(b)-1])
	})

}

// NewExchangeDataCollection creates an ExchangeDataCollection with fullPath is annotated
func NewExchangeDataCollection(aUser string, pathRepresentation []string) ExchangeDataCollection {
	collection := ExchangeDataCollection{
		user:     aUser,
		data:     make(chan DataStream, maximumMessages),
		fullPath: pathRepresentation,
	}
	return collection
}

func (edc *ExchangeDataCollection) PopulateCollection(newData *ExchangeData) {
	edc.data <- newData
}

// FinishPopulation is used to indicate data population of the collection is complete
// TODO: This should be an internal method once we move the message retrieval logic into `ExchangeDataCollection`
func (edc *ExchangeDataCollection) FinishPopulation() {
	if edc != nil && edc.data != nil {
		close(edc.data)
	}
}

func (edc *ExchangeDataCollection) Items() <-chan DataStream {
	return edc.data
}

func (edc *ExchangeDataCollection) FullPath() []string {
	return append([]string{}, edc.fullPath...)
}

// ExchangeData represents a single item retrieved from exchange
type ExchangeData struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	message []byte
	info    *backup.ExchangeInfo
}

func (ed *ExchangeData) UUID() string {
	return ed.id
}

func (ed *ExchangeData) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(ed.message))
}

func (ed *ExchangeData) Info() backup.ItemInfo {
	return backup.ItemInfo{Exchange: ed.info}
}
