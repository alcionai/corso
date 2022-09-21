package stats

import "time"

// ReadWrites tracks the total count of reads and writes, and of
// read and write errors.  ItemsRead and ItemsWritten counts are
// assumed to be successful, so the total count of items involved
// would be ItemsRead+ReadErrors.
type ReadWrites struct {
	ItemsRead      int   `json:"itemsRead,omitempty"`
	ItemsWritten   int   `json:"itemsWritten,omitempty"`
	ReadErrors     error `json:"readErrors,omitempty"`
	WriteErrors    error `json:"writeErrors,omitempty"`
	ResourceOwners int   `json:"resourceOwners,omitempty"`
}

// StartAndEndTime tracks a paired starting time and ending time.
type StartAndEndTime struct {
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
}
