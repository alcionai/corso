package backup

import (
	"fmt"
	"time"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/internal/stats"
	"github.com/alcionai/corso/pkg/selectors"
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
	stats.StartAndEndTime
}

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
	// todo: implement printable backup struct
	return Printable{
		ID:         b.ID,
		ErrorCount: support.GetNumberOfErrors(b.ReadErrors) + support.GetNumberOfErrors(b.WriteErrors),
		StartedAt:  b.StartedAt,
		Status:     b.Status,
		Version:    "0",
		Selectors:  b.Selectors.Printable(),
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
		common.FormatTime(b.StartedAt),
		string(b.ID),
		status,
		b.Selectors.Printable().Resources(),
	}
}
