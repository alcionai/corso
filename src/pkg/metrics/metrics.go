package metrics

import (
	"io"

	"github.com/alcionai/corso/src/internal/common/syncd"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
)

type KindStats struct {
	BytesRead     int64
	ResourceCount int64
}

var (
	bytesRead count.Key = "bytes-read"
	resources count.Key = "resources"
)

type ExportStats struct {
	// data is kept private so that we can enforce atomic int updates
	data syncd.MapOf[path.CategoryType, *count.Bus]
}

func NewExportStats() *ExportStats {
	return &ExportStats{
		data: syncd.NewMapOf[path.CategoryType, *count.Bus](),
	}
}

func (es *ExportStats) UpdateBytes(kind path.CategoryType, numBytes int64) {
	es.getCB(kind).Add(bytesRead, numBytes)
}

func (es *ExportStats) UpdateResourceCount(kind path.CategoryType) {
	es.getCB(kind).Inc(resources)
}

func (es *ExportStats) getCB(kind path.CategoryType) *count.Bus {
	es.data.LazyInit()

	cb, ok := es.data.Load(kind)
	if !ok {
		cb = count.New()
		es.data.Store(kind, cb)
	}

	return cb
}

func (es *ExportStats) GetStats() map[path.CategoryType]KindStats {
	toKindStats := map[path.CategoryType]KindStats{}

	for k, cb := range es.data.Values() {
		toKindStats[k] = KindStats{
			BytesRead:     cb.Get(bytesRead),
			ResourceCount: cb.Get(resources),
		}
	}

	return toKindStats
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
