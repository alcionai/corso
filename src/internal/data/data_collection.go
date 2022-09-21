package data

import (
	"io"

	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/backup/details"
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

// ByteCounter tracks a total quantity of bytes.  Structs that comply with Collection
// should strive to comply with ByteCounter for metrics gathering.  Bytes are normally
// counted using StreamSize when reading from Items().
type ByteCounter interface {
	CountBytes(c int64)
	BytesCounted() int64
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

// CountAllBytes returns the total count of bytes across all collections.
// Bytes are only counted if the Collection can be cast to a ByteCounter.
func CountAllBytes(cs []Collection) int64 {
	var totalBytes int64

	for _, c := range cs {
		bc, ok := c.(ByteCounter)
		if !ok {
			continue
		}

		totalBytes += bc.BytesCounted()
	}

	return totalBytes
}
