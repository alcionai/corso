package details

import (
	"context"
	"sync"
	"time"

	"github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/internal/model"
)

// --------------------------------------------------------------------------------
// Model
// --------------------------------------------------------------------------------

// DetailsModel describes what was stored in a Backup
type DetailsModel struct {
	model.BaseModel
	Entries []DetailsEntry `json:"entries"`
}

// Print writes the DetailModel Entries to StdOut, in the format
// requested by the caller.
func (dm DetailsModel) PrintEntries(ctx context.Context) {
	ps := []print.Printable{}
	for _, de := range dm.Entries {
		ps = append(ps, de)
	}
	print.All(ctx, ps...)
}

// Paths returns the list of Paths extracted from the Entries slice.
func (dm DetailsModel) Paths() []string {
	ents := dm.Entries
	r := make([]string, len(ents))
	for i := range ents {
		r[i] = ents[i].RepoRef
	}
	return r
}

// --------------------------------------------------------------------------------
// Details
// --------------------------------------------------------------------------------

// Details augments the core with a mutex for processing.
// Should be sliced back to d.DetailsModel for storage and
// printing.
type Details struct {
	DetailsModel

	// internal
	mu sync.Mutex `json:"-"`
}

func (d *Details) Add(repoRef string, info ItemInfo) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entries = append(d.Entries, DetailsEntry{RepoRef: repoRef, ItemInfo: info})
}

// --------------------------------------------------------------------------------
// Entry
// --------------------------------------------------------------------------------

// DetailsEntry describes a single item stored in a Backup
type DetailsEntry struct {
	// TODO: `RepoRef` is currently the full path to the item in Kopia
	// This can be optimized.
	RepoRef string `json:"repoRef"`
	ItemInfo
}

// --------------------------------------------------------------------------------
// CLI Output
// --------------------------------------------------------------------------------

// interface compliance checks
var _ print.Printable = &DetailsEntry{}

// MinimumPrintable DetailsEntries is a passthrough func, because no
// reduction is needed for the json output.
func (de DetailsEntry) MinimumPrintable() any {
	return de
}

// Headers returns the human-readable names of properties in a DetailsEntry
// for printing out to a terminal in a columnar display.
func (de DetailsEntry) Headers() []string {
	hs := []string{"Repo Ref"}
	if de.ItemInfo.Exchange != nil {
		hs = append(hs, de.ItemInfo.Exchange.Headers()...)
	}
	if de.ItemInfo.Sharepoint != nil {
		hs = append(hs, de.ItemInfo.Sharepoint.Headers()...)
	}
	if de.ItemInfo.Onedrive != nil {
		hs = append(hs, de.ItemInfo.Onedrive.Headers()...)
	}
	return hs
}

// Values returns the values matching the Headers list.
func (de DetailsEntry) Values() []string {
	vs := []string{de.RepoRef}
	if de.ItemInfo.Exchange != nil {
		vs = append(vs, de.ItemInfo.Exchange.Values()...)
	}
	if de.ItemInfo.Sharepoint != nil {
		vs = append(vs, de.ItemInfo.Sharepoint.Values()...)
	}
	if de.ItemInfo.Onedrive != nil {
		vs = append(vs, de.ItemInfo.Onedrive.Values()...)
	}
	return vs
}

// ItemInfo is a oneOf that contains service specific
// information about the item it tracks
type ItemInfo struct {
	Exchange   *ExchangeInfo   `json:"exchange,omitempty"`
	Sharepoint *SharepointInfo `json:"sharepoint,omitempty"`
	Onedrive   *OnedriveInfo   `json:"onedrive,omitempty"`
}

// ExchangeInfo describes an exchange item
type ExchangeInfo struct {
	Sender      string    `json:"sender,omitempty"`
	Subject     string    `json:"subject,omitempty"`
	Received    time.Time `json:"received,omitempty"`
	EventStart  time.Time `json:"eventStart,omitempty"`
	Organizer   string    `json:"organizer,omitempty"`
	ContactName string    `json:"contactName,omitempty"`
}

// Headers returns the human-readable names of properties in an ExchangeInfo
// for printing out to a terminal in a columnar display.
func (e ExchangeInfo) Headers() []string {
	return []string{"Sender", "Subject", "Received"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (e ExchangeInfo) Values() []string {
	return []string{e.Sender, e.Subject, e.Received.Format(time.RFC3339Nano)}
}

// SharepointInfo describes a sharepoint item
// TODO: Implement this. This is currently here
// just to illustrate usage
type SharepointInfo struct{}

// Headers returns the human-readable names of properties in a SharepointInfo
// for printing out to a terminal in a columnar display.
func (s SharepointInfo) Headers() []string {
	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (s SharepointInfo) Values() []string {
	return []string{}
}

// OnedriveInfo describes a onedrive item
type OnedriveInfo struct {
	ParentPath string `json:"parentPath"`
	ItemName   string `json:"itemName"`
}

// Headers returns the human-readable names of properties in a OnedriveInfo
// for printing out to a terminal in a columnar display.
func (oi OnedriveInfo) Headers() []string {
	return []string{"ItemName", "ParentPath"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (oi OnedriveInfo) Values() []string {
	return []string{oi.ItemName, oi.ParentPath}
}
