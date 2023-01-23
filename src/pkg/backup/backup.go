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
	"github.com/alcionai/corso/src/pkg/fault"
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

	// Selector used in this operation
	Selector selectors.Selector `json:"selectors"`

	// Errors contains all errors aggregated during a backup operation.
	Errors fault.ErrorsData `json:"errors"`

	// stats are embedded so that the values appear as top-level properties
	stats.Errs // Deprecated, replaced with Errors.
	stats.ReadWrites
	stats.StartAndEndTime
}

// interface compliance checks
var _ print.Printable = &Backup{}

func New(
	snapshotID, detailsID, status string,
	id model.StableID,
	selector selectors.Selector,
	rw stats.ReadWrites,
	se stats.StartAndEndTime,
	errs *fault.Errors,
) *Backup {
	return &Backup{
		BaseModel: model.BaseModel{
			ID: id,
			Tags: map[string]string{
				model.ServiceTag: selector.PathService().String(),
			},
		},
		CreationTime:    time.Now(),
		SnapshotID:      snapshotID,
		DetailsID:       detailsID,
		Status:          status,
		Selector:        selector,
		Errors:          errs.Data(),
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
func PrintAll(ctx context.Context, bs []*Backup) {
	if len(bs) == 0 {
		print.Info(ctx, "No backups available")
		return
	}

	ps := []print.Printable{}
	for _, b := range bs {
		ps = append(ps, print.Printable(b))
	}

	print.All(ctx, ps...)
}

type Printable struct {
	ID            model.StableID `json:"id"`
	ErrorCount    int            `json:"errorCount"`
	StartedAt     time.Time      `json:"started at"`
	Status        string         `json:"status"`
	Version       string         `json:"version"`
	BytesRead     int64          `json:"bytesRead"`
	BytesUploaded int64          `json:"bytesUploaded"`
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (b Backup) MinimumPrintable() any {
	return Printable{
		ID:            b.ID,
		ErrorCount:    b.errorCount(),
		StartedAt:     b.StartedAt,
		Status:        b.Status,
		Version:       "0",
		BytesRead:     b.BytesRead,
		BytesUploaded: b.BytesUploaded,
	}
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (b Backup) Headers() []string {
	return []string{
		"Started At",
		"ID",
		"Status",
		"Resource Owner",
	}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (b Backup) Values() []string {
	status := fmt.Sprintf("%s (%d errors)", b.Status, b.errorCount())

	return []string{
		common.FormatTabularDisplayTime(b.StartedAt),
		string(b.ID),
		status,
		b.Selector.DiscreteOwner,
	}
}

func (b Backup) errorCount() int {
	var errCount int

	if b.Errors.Err != nil || len(b.Errors.Errs) > 0 {
		if b.Errors.Err != nil {
			errCount++
		}

		errCount += len(b.Errors.Errs)
	} else {
		errCount = support.GetNumberOfErrors(b.ReadErrors) + support.GetNumberOfErrors(b.WriteErrors)
	}

	return errCount
}
