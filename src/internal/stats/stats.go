package stats

import (
	"sync/atomic"
	"time"
)

// ReadWrites tracks the total count of reads and writes.  ItemsRead
// and ItemsWritten counts are assumed to be successful reads.
type ReadWrites struct {
	BytesRead            int64 `json:"bytesRead,omitempty"`
	BytesUploaded        int64 `json:"bytesUploaded,omitempty"`
	ItemsRead            int   `json:"itemsRead,omitempty"`
	NonMetaBytesUploaded int64 `json:"nonMetaBytesUploaded,omitempty"`
	NonMetaItemsWritten  int   `json:"nonMetaItemsWritten,omitempty"`
	ItemsWritten         int   `json:"itemsWritten,omitempty"`
	ResourceOwners       int   `json:"resourceOwners,omitempty"`
}

// StartAndEndTime tracks a paired starting time and ending time.
type StartAndEndTime struct {
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
}

type ByteCounter struct {
	NumBytes int64
}

func (bc *ByteCounter) Count(i int64) {
	atomic.AddInt64(&bc.NumBytes, i)
}

type SkippedCounts struct {
	TotalSkippedItems         int `json:"totalSkippedItems"`
	SkippedMalware            int `json:"skippedMalware"`
	SkippedInvalidOneNoteFile int `json:"skippedInvalidOneNoteFile"`
}
