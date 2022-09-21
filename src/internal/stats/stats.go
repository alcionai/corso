package stats

import "time"

// ReadWrites tracks the total count of reads and writes, and of
// read and write errors.  ItemsRead and ItemsWritten counts are
// assumed to be successful, so the total count of items involved
// would be ItemsRead+ReadErrors.
type ReadWrites struct {
	CollectedBytes int64 `json:"collectedBytes,omitempty"`
	ItemsRead      int   `json:"itemsRead,omitempty"`
	ItemsWritten   int   `json:"itemsWritten,omitempty"`
	ResourceOwners int   `json:"resourceOwners,omitempty"`
}

// Errs tracks the errors which occurred during an operation.
type Errs struct {
	ReadErrors  error `json:"readErrors,omitempty"`
	WriteErrors error `json:"writeErrors,omitempty"`
}

// StartAndEndTime tracks a paired starting time and ending time.
type StartAndEndTime struct {
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
}
