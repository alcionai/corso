package backup

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// Backup represents the result of a backup operation
type Backup struct {
	model.BaseModel
	CreationTime time.Time `json:"creationTime"`

	// SnapshotID is the kopia snapshot ID
	SnapshotID string `json:"snapshotID"`

	// Reference to `Details`
	// We store the ModelStoreID since Details is immutable
	DetailsID string `json:"detailsID"`

	// Status of the operation
	Status string `json:"status"`

	// Selectors used in this operation
	Selectors selectors.Selector `json:"selectors"`

	// stats are embedded so that the values appear as top-level properties
	stats.ReadWrites
	stats.Errs
	stats.StartAndEndTime
}

// interface compliance checks
var _ print.Printable = &Backup{}

func New(
	snapshotID, detailsID, status string,
	selector selectors.Selector,
	rw stats.ReadWrites,
	se stats.StartAndEndTime,
) *Backup {
	return &Backup{
		CreationTime:    time.Now(),
		SnapshotID:      snapshotID,
		DetailsID:       detailsID,
		Status:          status,
		Selectors:       selector,
		ReadWrites:      rw,
		StartAndEndTime: se,
	}
}

// --------------------------------------------------------------------------------
// CLI Output
// --------------------------------------------------------------------------------

// Print writes the Backup to StdOut, in the format requested by the caller.
func (b Backup) Print(ctx context.Context) {
	print.Item(ctx, b)
}

// PrintAll writes the slice of Backups to StdOut, in the format requested by the caller.
func PrintAll(ctx context.Context, bs []Backup) {
	ps := []print.Printable{}
	for _, b := range bs {
		ps = append(ps, print.Printable(b))
	}

	print.All(ctx, ps...)
}

type Printable struct {
	ID         model.StableID      `json:"id"`
	ErrorCount int                 `json:"errorCount"`
	StartedAt  time.Time           `json:"started at"`
	Status     string              `json:"status"`
	Version    string              `json:"version"`
	Selectors  selectors.Printable `json:"selectors"`
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (b Backup) MinimumPrintable() any {
	return Printable{
		ID:         b.ID,
		ErrorCount: support.GetNumberOfErrors(b.ReadErrors) + support.GetNumberOfErrors(b.WriteErrors),
		StartedAt:  b.StartedAt,
		Status:     b.Status,
		Version:    "0",
		Selectors:  b.Selectors.ToPrintable(),
	}
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (b Backup) Headers() []string {
	return []string{
		"Started At",
		"ID",
		"Status",
		"Selectors",
	}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (b Backup) Values() []string {
	errCount := support.GetNumberOfErrors(b.ReadErrors) + support.GetNumberOfErrors(b.WriteErrors)
	status := fmt.Sprintf("%s (%d errors)", b.Status, errCount)

	return []string{
		common.FormatTabularDisplayTime(b.StartedAt),
		string(b.ID),
		status,
		b.Selectors.ToPrintable().Resources(),
	}
}
