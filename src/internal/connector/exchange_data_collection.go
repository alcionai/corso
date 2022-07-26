package connector

import (
	"bytes"
	"io"

	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
)

var _ data.DataCollection = &ExchangeDataCollection{}
var _ data.DataStream = &ExchangeData{}
var _ data.DataStreamInfo = &ExchangeData{}

const (
	collectionChannelBufferSize = 120
)

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type ExchangeDataCollection struct {
	// M365 user
	user         string
	data         chan data.DataStream
	tasks        []string
	updateCh     chan support.ConnectorOperationStatus
	service      graphService
	populateFunc PopulateFunc

	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	fullPath []string
}

// NewExchangeDataCollection creates an ExchangeDataCollection with fullPath is annotated
func NewExchangeDataCollection(aUser string, pathRepresentation []string) ExchangeDataCollection {
	collection := ExchangeDataCollection{
		user:     aUser,
		data:     make(chan data.DataStream, collectionChannelBufferSize),
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

func (edc *ExchangeDataCollection) Items() <-chan data.DataStream {
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
	info    *details.ExchangeInfo
}

func (ed *ExchangeData) UUID() string {
	return ed.id
}

func (ed *ExchangeData) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(ed.message))
}

func (ed *ExchangeData) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: ed.info}
}
