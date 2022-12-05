package data

import (
	"io"
	"time"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

// ------------------------------------------------------------------------------------------------
// standard ifaces
// ------------------------------------------------------------------------------------------------

// A Collection represents a compilation of data from the
// same type application (e.g. mail)
type Collection interface {
	// Items returns a channel from which items in the collection can be read.
	// Each returned struct contains the next item in the collection
	// The channel is closed when there are no more items in the collection or if
	// an unrecoverable error caused an early termination in the sender.
	Items() <-chan Stream
	// FullPath returns a path struct that acts as a metadata tag for this
	// DataCollection. Returned items should be ordered from most generic to least
	// generic. For example, a DataCollection for emails from a specific user
	// would be {"<tenant id>", "exchange", "<user ID>", "emails"}.
	FullPath() path.Path
}

// Stream represents a single item within a Collection
// that can be consumed as a stream (it embeds io.Reader)
type Stream interface {
	// ToReader returns an io.Reader for the DataStream
	ToReader() io.ReadCloser
	// UUID provides a unique identifier for this data
	UUID() string
}

// StreamInfo is used to provide service specific
// information about the Stream
type StreamInfo interface {
	Info() details.ItemInfo
}

// StreamSize is used to provide size
// information about the Stream
type StreamSize interface {
	Size() int64
}

// StreamModTime is used to provide the modified time of the stream's data.
type StreamModTime interface {
	ModTime() time.Time
}

// ------------------------------------------------------------------------------------------------
// functionality
// ------------------------------------------------------------------------------------------------

// ResourceOwnerSet extracts the set of unique resource owners from the
// slice of Collections.
func ResourceOwnerSet(cs []Collection) []string {
	rs := map[string]struct{}{}

	for _, c := range cs {
		fp := c.FullPath()
		rs[fp.ResourceOwner()] = struct{}{}
	}

	rss := make([]string, 0, len(rs))

	for k := range rs {
		rss = append(rss, k)
	}

	return rss
}
