package backup

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// Backup represents the result of a backup operation
type Backup struct {
	model.BaseModel
	CreationTime time.Time `json:"creationTime"`

	// SnapshotID is the kopia snapshot ID
	SnapshotID string `json:"snapshotID"`

	// Reference to the backup details storage location.
	// Used to read backup.Details from the streamstore.
	DetailsID string `json:"detailsID"`

	// Reference to the fault errors storage location.
	// Used to read fault.Errors from the streamstore.
	ErrorsID string `json:"errorsID"`

	// Status of the operation, eg: completed, failed, etc
	Status string `json:"status"`

	// Selector used in this operation
	Selector selectors.Selector `json:"selectors"`

	// Version represents the version of the backup format
	Version int `json:"version"`

	FailFast bool `json:"failFast"`

	// the quantity of errors, both hard failure and recoverable.
	ErrorCount int `json:"errorCount"`

	// the non-recoverable failure message, only populated if one occurred.
	Failure string `json:"failure"`

	// stats are embedded so that the values appear as top-level properties
	stats.ReadWrites
	stats.StartAndEndTime
	stats.SkippedCounts
}

// interface compliance checks
var _ print.Printable = &Backup{}

func New(
	snapshotID, detailsID, errorsID, status string,
	id model.StableID,
	selector selectors.Selector,
	rw stats.ReadWrites,
	se stats.StartAndEndTime,
	errs *fault.Bus,
) *Backup {
	var (
		ee = errs.Errors()
		// TODO: count errData.Items(), not all recovered errors.
		errCount   = len(ee.Recovered)
		failMsg    string
		malware    int
		otherSkips int
	)

	if ee.Failure != nil {
		failMsg = ee.Failure.Error()
		errCount++
	}

	for _, s := range ee.Skipped {
		if s.HasCause(fault.SkipMalware) {
			malware++
		} else {
			otherSkips++
		}
	}

	return &Backup{
		BaseModel: model.BaseModel{
			ID: id,
			Tags: map[string]string{
				model.ServiceTag: selector.PathService().String(),
			},
		},

		Version:    version.Backup,
		SnapshotID: snapshotID,
		DetailsID:  detailsID,
		ErrorsID:   errorsID,

		CreationTime: time.Now(),
		Status:       status,

		Selector: selector,
		FailFast: errs.FailFast(),

		ErrorCount: errCount,
		Failure:    failMsg,

		ReadWrites:      rw,
		StartAndEndTime: se,
		SkippedCounts: stats.SkippedCounts{
			TotalSkippedItems: len(ee.Skipped),
			SkippedMalware:    malware,
		},
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
	Owner         string         `json:"owner"`
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (b Backup) MinimumPrintable() any {
	return Printable{
		ID:            b.ID,
		ErrorCount:    b.ErrorCount,
		StartedAt:     b.StartedAt,
		Status:        b.Status,
		Version:       "0",
		BytesRead:     b.BytesRead,
		BytesUploaded: b.BytesUploaded,
		Owner:         b.Selector.DiscreteOwner,
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
	var (
		status   = b.Status
		errCount = b.ErrorCount
	)

	if errCount+b.TotalSkippedItems > 0 {
		status += (" (")
	}

	if errCount > 0 {
		status += fmt.Sprintf("%d errors", errCount)
	}

	if errCount > 0 && b.TotalSkippedItems > 0 {
		status += ", "
	}

	if b.TotalSkippedItems > 0 {
		status += fmt.Sprintf("%d skipped", b.TotalSkippedItems)
	}

	if b.TotalSkippedItems > 0 && b.SkippedMalware > 0 {
		status += ": "
	}

	if b.SkippedMalware > 0 {
		status += fmt.Sprintf("%d malware", b.SkippedMalware)
	}

	if errCount+b.TotalSkippedItems > 0 {
		status += (")")
	}

	return []string{
		common.FormatTabularDisplayTime(b.StartedAt),
		string(b.ID),
		status,
		b.Selector.DiscreteOwner,
	}
}
