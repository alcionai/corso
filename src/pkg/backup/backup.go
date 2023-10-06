package backup

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/errs"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/backup/identity"
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

	// ** DO NOT CHANGE JSON TAG NAMES **
	// These are in-memory only variable renames of previously persisted fields.
	// ** CHANGING THE JSON TAGS WILL BREAK THINGS BECAUSE THE MODEL WON'T **
	// ** DESERIALIZE PROPERLY **
	ProtectedResourceID   string `json:"resourceOwnerID,omitempty"`
	ProtectedResourceName string `json:"resourceOwnerName,omitempty"`

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

	// MergeBases records the set of merge bases used for this backup and the
	// Reason(s) each merge base was selected. Reasons are serialized the same
	// way that Reason tags are serialized.
	MergeBases map[model.StableID][]string `json:"mergeBases,omitempty"`
	// AssistBases records the set of assist bases used for this backup and the
	// Reason(s) each assist base was selected. Reasons are serialized the same
	// way that Reason tags are serialized.
	AssistBases map[model.StableID][]string `json:"assistBases,omitempty"`

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
	reasons []identity.Reasoner,
	bases BackupBases,
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
		case s.HasCause(fault.SkipBigOneNote):
			invalidONFile++
		default:
			otherSkips++
		}
	}

	// maps.Clone throws an NPE if passed nil on Mac for some reason.
	if tags == nil {
		tags = map[string]string{}
	}

	b := &Backup{
		BaseModel: model.BaseModel{
			ID:   id,
			Tags: maps.Clone(tags),
		},

		ProtectedResourceID:   ownerID,
		ProtectedResourceName: ownerName,

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
			SkippedInvalidOneNoteFile: invalidONFile,
		},
	}

	if bases != nil {
		mergeBases := map[model.StableID][]string{}
		assistBases := map[model.StableID][]string{}

		for _, backup := range bases.Backups() {
			for _, reason := range backup.Reasons {
				mergeBases[backup.ID] = append(
					mergeBases[backup.ID],
					ServiceCatString(reason.Service(), reason.Category()))
			}
		}

		for _, backup := range bases.UniqueAssistBackups() {
			for _, reason := range backup.Reasons {
				assistBases[backup.ID] = append(
					assistBases[backup.ID],
					ServiceCatString(reason.Service(), reason.Category()))
			}
		}

		if len(mergeBases) > 0 {
			b.MergeBases = mergeBases
		}

		if len(assistBases) > 0 {
			b.AssistBases = assistBases
		}
	}

	for _, reason := range reasons {
		for k, v := range reasonTags(reason) {
			b.Tags[k] = v
		}
	}

	return b
}

// PersistedBaseSet contains information extracted from the backup model
// relating to its lineage. It only contains the backup ID and Reasons each
// base was selected instead of the full set of information contained in other
// structs like BackupBases.
type PersistedBaseSet struct {
	Merge  map[model.StableID][]identity.Reasoner
	Assist map[model.StableID][]identity.Reasoner
}

func (b Backup) Bases() (PersistedBaseSet, error) {
	res := PersistedBaseSet{
		Merge:  map[model.StableID][]identity.Reasoner{},
		Assist: map[model.StableID][]identity.Reasoner{},
	}

	for id, reasons := range b.MergeBases {
		for _, reason := range reasons {
			service, cat, err := serviceCatStringToTypes(reason)
			if err != nil {
				return res, clues.Wrap(err, "getting Reason info").With(
					"base_type", "merge",
					"base_backup_id", id,
					"input_string", reason)
			}

			res.Merge[id] = append(res.Merge[id], identity.NewReason(
				// Tenant ID not currently stored in backup model.
				"",
				str.First(
					b.ProtectedResourceID,
					b.Selector.DiscreteOwner),
				service,
				cat))
		}
	}

	for id, reasons := range b.AssistBases {
		for _, reason := range reasons {
			service, cat, err := serviceCatStringToTypes(reason)
			if err != nil {
				return res, clues.Wrap(err, "getting Reason info").With(
					"base_type", "assist",
					"base_backup_id", id,
					"input_string", reason)
			}

			res.Assist[id] = append(res.Assist[id], identity.NewReason(
				// Tenant ID not currently stored in backup model.
				"",
				str.First(
					b.ProtectedResourceID,
					b.Selector.DiscreteOwner),
				service,
				cat))
		}
	}

	return res, nil
}

func (b Backup) Tenant() (string, error) {
	t := b.Tags[TenantIDKey]
	if len(t) == 0 {
		return "", clues.Wrap(errs.NotFound, "getting tenant")
	}

	return t, nil
}

// Reasons returns the set of services and categories this backup encompassed
// for the tenant and protected resource.
func (b Backup) Reasons() ([]identity.Reasoner, error) {
	tenant, err := b.Tenant()
	if err != nil {
		return nil, clues.Stack(err)
	}

	var res []identity.Reasoner

	for tag := range b.Tags {
		service, cat, err := serviceCatStringToTypes(tag)
		if err != nil {
			// Assume it's just not one of the Reason tags.
			if errors.Is(err, errMissingPrefix) {
				continue
			}

			return nil, clues.Wrap(err, "parsing reasons")
		}

		res = append(
			res,
			identity.NewReason(
				tenant,
				str.First(b.ProtectedResourceID, b.Selector.DiscreteOwner),
				service,
				cat))
	}

	return res, nil
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
		b.ProtectedResourceID,
		b.Selector.Name())

	bs := b.toStats()

	return []string{
		string(b.ID),
		dttm.FormatToTabularDisplay(b.StartedAt),
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
