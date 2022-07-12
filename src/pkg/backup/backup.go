package backup

import (
	"sync"
	"time"

	"github.com/alcionai/corso/internal/model"
)

// Backup represents the result of a backup operation
type Backup struct {
	model.BaseModel
	CreationTime time.Time `json:"creationTime"`

	// SnapshotID is the kopia snapshot ID
	SnapshotID string `json:"snapshotID"`

	// Reference to `Details`
	// We store the ModelStoreID since Details is immutable
	DetailsID string `json:"detailsId"`

	// TODO:
	// - Backup "Specification"
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (b Backup) Headers() []string {
	return []string{
		"Creation Time",
		"Stable ID",
		"Snapshot ID",
		"Details ID",
	}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (b Backup) Values() []string {
	return []string{
		b.CreationTime.Format(time.RFC3339Nano),
		string(b.StableID),
		b.SnapshotID,
		b.DetailsID,
	}
}

func New(snapshotID, detailsID string) *Backup {
	return &Backup{
		CreationTime: time.Now(),
		SnapshotID:   snapshotID,
		DetailsID:    detailsID,
	}
}

// DetailsModel describes what was stored in a Backup
type DetailsModel struct {
	model.BaseModel
	Entries []DetailsEntry `json:"entries"`
}

// Details augments the core with a mutext for processing.
// Should be sliced back to d.DetailsModel for storage and
// printing.
type Details struct {
	DetailsModel

	// internal
	mu sync.Mutex `json:"-"`
}

// DetailsEntry describes a single item stored in a Backup
type DetailsEntry struct {
	// TODO: `RepoRef` is currently the full path to the item in Kopia
	// This can be optimized.
	RepoRef string `json:"repoRef"`
	ItemInfo
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
	return vs
}

// ItemInfo is a oneOf that contains service specific
// information about the item it tracks
type ItemInfo struct {
	Exchange   *ExchangeInfo   `json:"exchange,omitempty"`
	Sharepoint *SharepointInfo `json:"sharepoint,omitempty"`
}

// ExchangeInfo describes an exchange item
type ExchangeInfo struct {
	Sender   string    `json:"sender"`
	Subject  string    `json:"subject"`
	Received time.Time `json:"received"`
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

func (d *Details) Add(repoRef string, info ItemInfo) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entries = append(d.Entries, DetailsEntry{RepoRef: repoRef, ItemInfo: info})
}

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
