package connector

import (
	"bytes"
	"io"
)

const (
	// TODO: Reduce this when https://github.com/alcionai/corso/issues/124 is closed
	// and we make channel population async (decouple from collection initialization)
	collectionChannelBufferSize = 1000
)

// A DataCollection represents a collection of data of the
// same type (e.g. mail)
type DataCollection interface {
	// Returns either the next item in the collection or an error if one occurred.
	// If not more items are available in the collection, returns (nil, nil).
	NextItem() (DataStream, error)
}

// DataStream represents a single item within a DataCollection
// that can be consumed as a stream (it embeds io.Reader)
type DataStream interface {
	// Returns an io.Reader for the DataStream
	ToReader() io.Reader
	// Provides a unique identifier for this data
	UUID() string
}

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type ExchangeDataCollection struct {
	// M365 user
	user string
	// TODO: We would want to replace this with a channel so that we
	// don't need to wait for all data to be retrieved before reading it out
	data chan ExchangeData
	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	FullPath []string
}

// NewExchangeDataCollection creates an ExchangeDataCollection where
// the FullPath is confgured
func NewExchangeDataCollection(aUser string, pathRepresentation []string) ExchangeDataCollection {
	collection := ExchangeDataCollection{
		user:     aUser,
		data:     make(chan ExchangeData, collectionChannelBufferSize),
		FullPath: pathRepresentation,
	}
	return collection
}

func (edc *ExchangeDataCollection) PopulateCollection(newData ExchangeData) {
	edc.data <- newData
}

// FinishPopulation is used to indicate data population of the collection is complete
// TODO: This should be an internal method once we move the message retrieval logic into `ExchangeDataCollection`
func (edc *ExchangeDataCollection) FinishPopulation() {
	close(edc.data)
}

func (edc *ExchangeDataCollection) Length() int {
	return len(edc.data)
}

// NextItem returns either the next item in the collection or an error if one occurred.
// If not more items are available in the collection, returns (nil, nil).
func (edc *ExchangeDataCollection) NextItem() (DataStream, error) {
	item, ok := <-edc.data
	if !ok {
		return nil, io.EOF
	}
	return &item, nil
}

// ExchangeData represents a single item retrieved from exchange
type ExchangeData struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	message []byte
}

func (ed *ExchangeData) UUID() string {
	return ed.id
}

func (ed *ExchangeData) ToReader() io.Reader {
	return bytes.NewReader(ed.message)
}
