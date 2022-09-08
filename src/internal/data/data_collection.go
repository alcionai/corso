package data

import (
	"io"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

// A Collection represents a compilation of data from the
// same type application (e.g. mail)
type Collection interface {
	// Items returns a channel from which items in the collection can be read.
	// Each returned struct contains the next item in the collection
	// The channel is closed when there are no more items in the collection or if
	// an unrecoverable error caused an early termination in the sender.
	Items() <-chan Stream
	// FullPath returns a slice of strings that act as metadata tags for this
	// DataCollection. Returned items should be ordered from most generic to least
	// generic. For example, a DataCollection for emails from a specific user
	// would be {"<tenant id>", "<user ID>", "emails"}.
	FullPath() []string
}

// DataStream represents a single item within a DataCollection
// that can be consumed as a stream (it embeds io.Reader)
type Stream interface {
	// ToReader returns an io.Reader for the DataStream
	ToReader() io.ReadCloser
	// UUID provides a unique identifier for this data
	UUID() string
}

// DataStreamInfo is used to provide service specific
// information about the DataStream
type StreamInfo interface {
	Info() details.ItemInfo
}
