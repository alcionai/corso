package connector

import "io"

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
	io.Reader
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
	data          []ExchangeData
	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleamed from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	FullPath []string
}

// NewExchangeDataCollection creates an ExchangeDataCollection where
// the FullPath is confgured
func NewExchangeDataCollection(aUser string, pathRepresentation []string) ExchangeDataCollection {
	collection := ExchangeDataCollection{
		user:          aUser,
		data:          make([]ExchangeData, 0),
		FullPath:      pathRepresentation,
	}
	return collection
}

func (ec *ExchangeDataCollection) PopulateCollection(newData ExchangeData) {
	ec.data = append(ec.data, newData)
}
func (ec *ExchangeDataCollection) GetLength() int {
	return len(ec.data)
}

// NextItem returns either the next item in the collection or an error if one occurred.
// If not more items are available in the collection, returns (nil, nil).
func (*ExchangeDataCollection) NextItem() (DataStream, error) {
	// TODO: Return the next "to be read" item in the collection as a
	// DataStream
	return nil, nil
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

func (ed *ExchangeData) Read(bytes []byte) (int, error) {
	// TODO: Copy ed.message into []bytes. Will need to take care of partial reads
	return 0, nil
}
