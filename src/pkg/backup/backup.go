package backup

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
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

	// Reference to the details and fault errors storage location.
	// Used to read backup.Details and fault.Errors from the streamstore.
	StreamStoreID string `json:"streamStoreID"`

	// Status of the operation, eg: completed, failed, etc
	Status string `json:"status"`

	// Selector used in this operation
	Selector selectors.Selector `json:"selectors"`

	// ResourceOwner reference
	ResourceOwnerID   string `json:"resourceOwnerID"`
	ResourceOwnerName string `json:"resourceOwnerName"`

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

	// **Deprecated**
	// Reference to the backup details storage location.
	// Used to read backup.Details from the streamstore.
	DetailsID string `json:"detailsID"`
}

// interface compliance checks
var _ print.Printable = &Backup{}

func New(
	snapshotID, streamStoreID, status string,
	version int,
	id model.StableID,
	selector selectors.Selector,
	ownerID, ownerName string,
	rw stats.ReadWrites,
	se stats.StartAndEndTime,
	fe *fault.Errors,
) *Backup {
	if fe == nil {
		fe = &fault.Errors{}
	}

	var (
		errCount  = len(fe.Items)
		skipCount = len(fe.Skipped)
		failMsg   string

		malware, notFound,
		invalidONFile, otherSkips int
	)

	if fe.Failure != nil {
		failMsg = fe.Failure.Msg
		errCount++
	}

	for _, s := range fe.Skipped {
		switch true {
		case s.HasCause(fault.SkipMalware):
			malware++
		case s.HasCause(fault.SkipNotFound):
			notFound++
		case s.HasCause(fault.SkipBigOneNote):
			invalidONFile++
		default:
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

		ResourceOwnerID:   ownerID,
		ResourceOwnerName: ownerName,

		Version:       version,
		SnapshotID:    snapshotID,
		StreamStoreID: streamStoreID,

		CreationTime: time.Now(),
		Status:       status,

		Selector: selector,
		FailFast: fe.FailFast,

		ErrorCount: errCount,
		Failure:    failMsg,

		ReadWrites:      rw,
		StartAndEndTime: se,
		SkippedCounts: stats.SkippedCounts{
			TotalSkippedItems:         skipCount,
			SkippedMalware:            malware,
			SkippedNotFound:           notFound,
			SkippedInvalidOneNoteFile: invalidONFile,
		},
	}
}

// --------------------------------------------------------------------------------
// CLI Output
// --------------------------------------------------------------------------------

// ----- print backups

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
	ID      model.StableID `json:"id"`
	Status  string         `json:"status"`
	Version string         `json:"version"`
	Owner   string         `json:"owner"`
	Stats   backupStats    `json:"stats"`
}

// ToPrintable reduces the Backup to its minimally printable details.
func (b Backup) ToPrintable() Printable {
	return Printable{
		ID:      b.ID,
		Status:  b.Status,
		Version: "0",
		Owner:   b.Selector.DiscreteOwner,
		Stats:   b.toStats(),
	}
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (b Backup) MinimumPrintable() any {
	return b.ToPrintable()
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (b Backup) Headers() []string {
	return []string{
		"ID",
		"Started At",
		"Duration",
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

		if b.SkippedMalware+b.SkippedNotFound+b.SkippedInvalidOneNoteFile > 0 {
			status += ": "
		}
	}

	skipped := []string{}

	if b.SkippedMalware > 0 {
		skipped = append(skipped, fmt.Sprintf("%d malware", b.SkippedMalware))
	}

	if b.SkippedNotFound > 0 {
		skipped = append(skipped, fmt.Sprintf("%d not found", b.SkippedNotFound))
	}

	if b.SkippedInvalidOneNoteFile > 0 {
		skipped = append(skipped, fmt.Sprintf("%d invalid OneNote file", b.SkippedInvalidOneNoteFile))
	}

	status += strings.Join(skipped, ", ")

	if errCount+b.TotalSkippedItems > 0 {
		status += (")")
	}

	name := b.ResourceOwnerName

	if len(name) == 0 {
		name = b.ResourceOwnerID
	}

	if len(name) == 0 {
		name = b.Selector.DiscreteOwner
	}

	bs := b.toStats()

	return []string{
		string(b.ID),
		common.FormatTabularDisplayTime(b.StartedAt),
		bs.EndedAt.Sub(bs.StartedAt).String(),
		status,
		name,
	}
}

// ----- print backup stats

func (b Backup) toStats() backupStats {
	return backupStats{
		ID:            string(b.ID),
		BytesRead:     b.BytesRead,
		BytesUploaded: b.BytesUploaded,
		EndedAt:       b.CompletedAt,
		ErrorCount:    b.ErrorCount,
		ItemsRead:     b.ItemsRead,
		ItemsSkipped:  b.TotalSkippedItems,
		ItemsWritten:  b.ItemsWritten,
		StartedAt:     b.StartedAt,
	}
}

// interface compliance checks
var _ print.Printable = &backupStats{}

type backupStats struct {
	ID            string    `json:"id"`
	BytesRead     int64     `json:"bytesRead"`
	BytesUploaded int64     `json:"bytesUploaded"`
	EndedAt       time.Time `json:"endedAt"`
	ErrorCount    int       `json:"errorCount"`
	ItemsRead     int       `json:"itemsRead"`
	ItemsSkipped  int       `json:"itemsSkipped"`
	ItemsWritten  int       `json:"itemsWritten"`
	StartedAt     time.Time `json:"startedAt"`
}

// Print writes the Backup to StdOut, in the format requested by the caller.
func (bs backupStats) Print(ctx context.Context) {
	print.Item(ctx, bs)
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (bs backupStats) MinimumPrintable() any {
	return bs
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (bs backupStats) Headers() []string {
	return []string{
		"ID",
		"Bytes Uploaded",
		"Items Uploaded",
		"Items Skipped",
		"Errors",
	}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (bs backupStats) Values() []string {
	return []string{
		bs.ID,
		humanize.Bytes(uint64(bs.BytesUploaded)),
		strconv.Itoa(bs.ItemsWritten),
		strconv.Itoa(bs.ItemsSkipped),
		strconv.Itoa(bs.ErrorCount),
	}
}
