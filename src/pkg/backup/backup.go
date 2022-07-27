package backup

import (
	"strconv"
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

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (b Backup) Headers() []string {
	return []string{
		"Creation Time",
		"Stable ID",
		"Snapshot ID",
		"Details ID",
		"Status",
		"Selectors",
		"Items Read",
		"Items Written",
		"Read Errors",
		"Write Errors",
		"Started At",
		"Completed At",
	}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (b Backup) Values() []string {
	return []string{
		common.FormatTime(b.CreationTime),
		string(b.ID),
		b.SnapshotID,
		b.DetailsID,
		b.Status,
		b.Selectors.String(),
		strconv.Itoa(b.ReadWrites.ItemsRead),
		strconv.Itoa(b.ReadWrites.ItemsWritten),
		strconv.Itoa(support.GetNumberOfErrors(b.ReadWrites.ReadErrors)),
		strconv.Itoa(support.GetNumberOfErrors(b.ReadWrites.WriteErrors)),
		common.FormatTime(b.StartAndEndTime.StartedAt),
		common.FormatTime(b.StartAndEndTime.CompletedAt),
	}
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
