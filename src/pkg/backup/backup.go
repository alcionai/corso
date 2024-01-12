package backup

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
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

	// TODO: in process of gaining support, most cases will still use
	// ResourceOwner and ResourceOwnerName.
	ProtectedResourceID   string `json:"protectedResourceID,omitempty"`
	ProtectedResourceName string `json:"protectedResourceName,omitempty"`

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

	// prefer protectedResource
	ResourceOwnerID   string `json:"resourceOwnerID,omitempty"`
	ResourceOwnerName string `json:"resourceOwnerName,omitempty"`
}

// interface compliance checks
var _ print.Printable = &Backup{}

func New(
	snapshotID, streamStoreID, status string,
	backupVersion int,
	id model.StableID,
	selector selectors.Selector,
	ownerID, ownerName string,
	rw stats.ReadWrites,
	se stats.StartAndEndTime,
	fe *fault.Errors,
	tags map[string]string,
) *Backup {
	if fe == nil {
		fe = &fault.Errors{}
	}

	var (
		errCount  = len(fe.Items)
		skipCount = len(fe.Skipped)
		failMsg   string

		malware, invalidONFile, otherSkips int
	)

	if fe.Failure != nil {
		failMsg = fe.Failure.Msg
		errCount++
	}

	for _, s := range fe.Skipped {
		switch true {
		case s.HasCause(fault.SkipMalware):
			malware++
		case s.HasCause(fault.SkipOneNote):
			invalidONFile++
		default:
			otherSkips++
		}
	}

	return &Backup{
		BaseModel: model.BaseModel{
			ID:   id,
			Tags: tags,
		},

		ResourceOwnerID:   ownerID,
		ResourceOwnerName: ownerName,

		Version:       backupVersion,
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
			SkippedInvalidOneNoteFile: invalidONFile,
		},
	}
}

// Type returns the type of the backup according to the value stored in the
// Backup's tags. Backup type is used during base finding to determine if a
// given backup is eligible to be used for the upcoming incremental backup.
func (b Backup) Type() string {
	t, ok := b.Tags[model.BackupTypeTag]

	// Older backups didn't set the backup type tag because we only persisted
	// backup models for the MergeBackup type. Corso started adding the backup
	// type tag when it was producing v8 backups. Any backup newer than that that
	// doesn't have a backup type should just return an empty type and let the
	// caller figure out what to do.
	if !ok &&
		b.Version != version.NoBackup &&
		b.Version <= version.All8MigrateUserPNToID {
		t = model.MergeBackup
	}

	return t
}

// --------------------------------------------------------------------------------
// CLI Output
// --------------------------------------------------------------------------------

// ----- print backups

// Print writes the Backup to StdOut, in the format requested by the caller.
func (b Backup) Print(ctx context.Context) {
	print.Item(ctx, b)
}

// PrintProperties writes the Backup to StdOut, in the format requested by the caller.
// Unlike Print, it skips the ID of the Backup
func (b Backup) PrintProperties(ctx context.Context) {
	print.ItemProperties(ctx, b)
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
	ID                    model.StableID `json:"id"`
	Status                string         `json:"status"`
	Version               string         `json:"version"`
	ProtectedResourceID   string         `json:"protectedResourceID,omitempty"`
	ProtectedResourceName string         `json:"protectedResourceName,omitempty"`
	Owner                 string         `json:"owner,omitempty"`
	Stats                 backupStats    `json:"stats"`
}

// ToPrintable reduces the Backup to its minimally printable details.
func (b Backup) ToPrintable() Printable {
	return Printable{
		ID:                    b.ID,
		Status:                b.Status,
		Version:               "0",
		ProtectedResourceID:   b.Selector.DiscreteOwner,
		ProtectedResourceName: b.Selector.DiscreteOwnerName,
		Owner:                 b.Selector.DiscreteOwner,
		Stats:                 b.toStats(),
	}
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (b Backup) MinimumPrintable() any {
	return b.ToPrintable()
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (b Backup) Headers(skipID bool) []string {
	headers := []string{
		"ID",
		"Started at",
		"Duration",
		"Status",
		"Protected resource",
		"Data",
	}

	if skipID {
		headers = headers[1:]
	}

	return headers
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (b Backup) Values(skipID bool) []string {
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

		if b.SkippedMalware+b.SkippedInvalidOneNoteFile > 0 {
			status += ": "
		}
	}

	skipped := []string{}

	if b.SkippedMalware > 0 {
		skipped = append(skipped, fmt.Sprintf("%d malware", b.SkippedMalware))
	}

	if b.SkippedInvalidOneNoteFile > 0 {
		skipped = append(skipped, fmt.Sprintf("%d invalid OneNote file", b.SkippedInvalidOneNoteFile))
	}

	status += strings.Join(skipped, ", ")

	if errCount+b.TotalSkippedItems > 0 {
		status += (")")
	}

	name := str.First(
		b.ProtectedResourceName,
		b.ResourceOwnerName,
		b.ProtectedResourceID,
		b.ResourceOwnerID,
		b.Selector.Name())

	bs := b.toStats()

	reasons, err := b.Selector.Reasons("doesnt-matter", false)
	if err != nil {
		logger.CtxErr(context.Background(), err).Error("getting reasons from selector")
	}

	reasonCats := []string{}

	for _, r := range reasons {
		reasonCats = append(reasonCats, r.Category().HumanString())
	}

	values := []string{
		string(b.ID),
		dttm.FormatToTabularDisplay(b.StartedAt),
		bs.EndedAt.Sub(bs.StartedAt).String(),
		status,
		name,
		strings.Join(reasonCats, ","),
	}

	if skipID {
		values = values[1:]
	}

	return values
}

// ----- print backup stats

func (b Backup) toStats() backupStats {
	return backupStats{
		ID:            string(b.ID),
		BytesRead:     b.BytesRead,
		BytesUploaded: b.NonMetaBytesUploaded,
		EndedAt:       b.CompletedAt,
		ErrorCount:    b.ErrorCount,
		ItemsRead:     b.ItemsRead,
		ItemsSkipped:  b.TotalSkippedItems,
		ItemsWritten:  b.NonMetaItemsWritten,
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

// PrintProperties writes the Backup to StdOut, in the format requested by the caller.
// Unlike Print, it skips the ID of backupStats
func (bs backupStats) PrintProperties(ctx context.Context) {
	print.ItemProperties(ctx, bs)
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (bs backupStats) MinimumPrintable() any {
	return bs
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (bs backupStats) Headers(skipID bool) []string {
	headers := []string{
		"Bytes Uploaded",
		"Items Uploaded",
		"Items Skipped",
		"Errors",
	}

	if skipID {
		return headers
	}

	return append([]string{"ID"}, headers...)
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (bs backupStats) Values(skipID bool) []string {
	values := []string{
		humanize.Bytes(uint64(bs.BytesUploaded)),
		strconv.Itoa(bs.ItemsWritten),
		strconv.Itoa(bs.ItemsSkipped),
		strconv.Itoa(bs.ErrorCount),
	}

	if skipID {
		return values
	}

	return append([]string{bs.ID}, values...)
}
