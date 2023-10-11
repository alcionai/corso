package data

import (
	"io"
	"sync/atomic"

	"github.com/alcionai/corso/src/pkg/path"
)

type CollectionStats struct {
	Folders,
	Objects,
	Successes int
	Bytes   int64
	Details string
}

func (cs CollectionStats) IsZero() bool {
	return cs.Folders+cs.Objects+cs.Successes+int(cs.Bytes) == 0
}

func (cs CollectionStats) String() string {
	return cs.Details
}

type KindStats struct {
	BytesRead     int64
	ResourceCount int64
}

type ExportStats struct {
	// data is kept private so that we can enforce atomic int updates
	data map[path.CategoryType]KindStats
}

func (es *ExportStats) UpdateBytes(kind path.CategoryType, bytesRead int64) {
	if es.data == nil {
		es.data = map[path.CategoryType]KindStats{}
	}

	ks := es.data[kind]
	atomic.AddInt64(&ks.BytesRead, bytesRead)
	es.data[kind] = ks
}

func (es *ExportStats) UpdateResourceCount(kind path.CategoryType) {
	if es.data == nil {
		es.data = map[path.CategoryType]KindStats{}
	}

	ks := es.data[kind]
	atomic.AddInt64(&ks.ResourceCount, 1)
	es.data[kind] = ks
}

func (es *ExportStats) GetStats() map[path.CategoryType]KindStats {
	return es.data
}

type statsReader struct {
	io.ReadCloser
	kind  path.CategoryType
	stats *ExportStats
}

func (sr *statsReader) Read(p []byte) (int, error) {
	n, err := sr.ReadCloser.Read(p)
	sr.stats.UpdateBytes(sr.kind, int64(n))

	return n, err
}

// Create a function that will take a reader and return a reader that
// will update the stats
func ReaderWithStats(
	reader io.ReadCloser,
	kind path.CategoryType,
	stats *ExportStats,
) io.ReadCloser {
	if reader == nil {
		return nil
	}

	return &statsReader{
		ReadCloser: reader,
		kind:       kind,
		stats:      stats,
	}
}
