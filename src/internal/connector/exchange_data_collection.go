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
// that can be consumed as a steam
type DataStream interface {
	// ToReader returns a reader for the data stream
	ToReader() io.Reader
	// Provides a unique identifier for this data
	UUID() string
}

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type ExchangeDataCollection struct {
	user string
}

// NextItem returns either the next item in the collection or an error if one occurred.
// If not more items are available in the collection, returns (nil, nil).
func (*ExchangeDataCollection) NextItem() (DataStream, error) {
	// TODO: Return the next "to be read" item in the collection as a
	// DataStream
	return nil, nil
}
