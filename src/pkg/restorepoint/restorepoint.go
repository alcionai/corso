package restorepoint

import (
	"sync"
	"time"

	"github.com/alcionai/corso/internal/model"
)

// RestorePoint represents the result of a backup operation
// that can be restored
type RestorePoint struct {
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

func New(snapshotID string, detailsID string) *RestorePoint {
	return &RestorePoint{
		CreationTime: time.Now(),
		SnapshotID:   snapshotID,
		DetailsID:    detailsID,
	}
}

// Details describes what was stored in a RestorePoint
type Details struct {
	model.BaseModel
	Entries []DetailsEntry `json:"entries"`

	// internal
	mu sync.Mutex
}

// DetailsEntry describes a single item stored in a RestorePoint
type DetailsEntry struct {
	// TODO: `RepoRef` is currently the full path to the item in Kopia
	// This can be optimized.
	RepoRef string `json:"repoRef"`
	ItemInfo
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

// SharepointInfo describes a sharepoint item
// TODO: Implement this. This is currently here
// just to illustrate usage
type SharepointInfo struct{}

func (d *Details) Add(repoRef string, info ItemInfo) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entries = append(d.Entries, DetailsEntry{RepoRef: repoRef, ItemInfo: info})
}
