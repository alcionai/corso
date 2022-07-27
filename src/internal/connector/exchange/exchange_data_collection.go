// Package exchange provides support for retrieving M365 Exchange objects
// from M365 servers using the Graph API. M365 object support centers
// on the applications: Mail, Contacts, and Calendar.
package exchange

import (
	"bytes"
	"context"
	"io"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
)

var _ data.Collection = &Collection{}
var _ data.Stream = &Stream{}
var _ data.StreamInfo = &Stream{}

const (
	collectionChannelBufferSize = 1000
)

type Service interface {
	// Client() returns msgraph Service client that can be used to process and execute
	// the majority of the queries to the M365 Backstore
	Client() *msgraphsdk.GraphServiceClient
	// Adapter() returns GraphRequest adapter used to process large requests, create batches
	// and page iterators
	Adapter() *msgraphsdk.GraphRequestAdapter
}

// PopulateFunc are a class of functions that can be used to fill exchange.Collections with
// the corresponding information
type PopulateFunc func(context.Context, Service, Collection, chan *support.ConnectorOperationStatus)

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type Collection struct {
	// M365 user
	User         string // M365 user
	Data         chan data.Stream
	tasks        []string
	updateCh     chan support.ConnectorOperationStatus
	service      Service
	populateFunc PopulateFunc

	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	fullPath []string
}

// NewExchangeDataCollection creates an ExchangeDataCollection with fullPath is annotated
func NewCollection(aUser string, pathRepresentation []string) Collection {
	collection := Collection{
		User:     aUser,
		Data:     make(chan data.Stream, collectionChannelBufferSize),
		fullPath: pathRepresentation,
	}
	return collection
}

func (eoc *Collection) PopulateCollection(newData *Stream) {
	eoc.Data <- newData
}

// FinishPopulation is used to indicate data population of the collection is complete
// TODO: This should be an internal method once we move the message retrieval logic into `ExchangeDataCollection`
func (eoc *Collection) FinishPopulation() {
	if eoc.Data != nil {
		close(eoc.Data)
	}
}

func (eoc *Collection) Items() <-chan data.Stream {
	return eoc.Data
}

func (edc *Collection) FullPath() []string {
	return append([]string{}, edc.fullPath...)
}

// Stream represents a single item retrieved from exchange
type Stream struct {
	Id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	Message []byte
	Inf     *details.ExchangeInfo //temporary change to bring populate function into directory
}

func (od *Stream) UUID() string {
	return od.Id
}

func (od *Stream) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(od.Message))
}

func (od *Stream) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: od.Inf}
}
