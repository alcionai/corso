package data

import (
	"context"
	"strconv"

	"github.com/alcionai/corso/src/cli/print"
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

// interface compliance checks
var _ print.Printable = &CollectionStats{}

// Print writes the Backup to StdOut, in the format requested by the caller.
func (cs CollectionStats) Print(ctx context.Context) {
	print.Item(ctx, cs)
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (cs CollectionStats) MinimumPrintable() any {
	return cs
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (cs CollectionStats) Headers() []string {
	return []string{
		"Folders",
		"Objects",
	}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (cs CollectionStats) Values() []string {
	return []string{
		strconv.Itoa(cs.Folders),
		strconv.Itoa(cs.Objects),
	}
}
