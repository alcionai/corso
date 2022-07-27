// Package exchange provides support for retrieving M365 Exchange objects
// from M365 servers using the Graph API. M365 object support centers
// on the applications: Mail, Contacts, and Calendar.
package exchange

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
)

var _ data.Collection = ObjectCollection{}
var _ data.Stream = &ObjectData{}
var _ data.StreamInfo = &ObjectData{}

const (
	collectionChannelBufferSize = 120
)

type Service interface {
	// Items returns a channel from which items in the collection can be read.
	// Each returned struct contains the next item in the collection
	// The channel is closed when there are no more items in the collection or if
	// an unrecoverable error caused an early termination in the sender.
	Client()
	// FullPath returns a slice of strings that act as metadata tags for this
	// DataCollection. Returned items should be ordered from most generic to least
	// generic. For example, a DataCollection for emails from a specific user
	// would be {"<tenant id>", "<user ID>", "emails"}.
	Adapter() []string
}
type PopulateFunc func(context.Context, Service, ObjectCollection, chan *support.ConnectorOperationStatus)

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type ObjectCollection struct {
	// M365 user
	user         string // M365 user
	data         chan data.Stream
	tasks        []string
	updateCh     chan support.ConnectorOperationStatus
	service      Service
	populateFunc PopulateFunc

	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	fullPath []string
}

// NewExchangeDataCollection creates an ExchangeDataCollection with fullPath is annotated
func NewObjectCollection(aUser string, pathRepresentation []string) ObjectCollection {
	collection := ObjectCollection{
		user:     aUser,
		data:     make(chan data.Stream, collectionChannelBufferSize),
		fullPath: pathRepresentation,
	}
	return collection
}

func (eoc *ObjectCollection) PopulateCollection(newData *ObjectData) {
	eoc.data <- newData
}

// FinishPopulation is used to indicate data population of the collection is complete
// TODO: This should be an internal method once we move the message retrieval logic into `ExchangeDataCollection`
func (eoc *ObjectCollection) FinishPopulation() {
	if eoc.data != nil {
		close(eoc.data)
	}
}

func (eoc *ObjectCollection) Items() <-chan data.Stream {
	return eoc.data
}

func (edc *ObjectCollection) FullPath() []string {
	return append([]string{}, edc.fullPath...)
}

// ExchangeData represents a single item retrieved from exchange
type ObjectData struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	message []byte
	info    *details.ExchangeInfo
}

func (od *ObjectData) UUID() string {
	return od.id
}

func (od *ObjectData) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(od.message))
}

func (od *ObjectData) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: od.info}
}
