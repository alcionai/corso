package metadata

import (
	"io"
	"time"
)

type Entity struct {
	ID         string  `json:"id,omitempty"`
	EntityType GV2Type `json:"entityType,omitempty"`
}

type LinkShareLink struct {
	Scope            string `json:"scope,omitempty"`
	Type             string `json:"type,omitempty"`
	WebUrl           string `json:"webUrl,omitempty"` // TODO(meain): Should we store this?
	PreventsDownload bool   `json:"preventsDownload,omitempty"`
}

type LinkShare struct {
	ID          string        `json:"id,omitempty"`
	Link        LinkShareLink `json:"link,omitempty"`
	Roles       []string      `json:"role,omitempty"`
	Entities    []Entity      `json:"entityId,omitempty"`    // this is the resource owner's ID
	HasPassword bool          `json:"hasPassword,omitempty"` // Only anonymous ones require password and we don't restore them
	Expiration  *time.Time    `json:"expiration,omitempty"`
}

// ItemMeta contains metadata about the Item. It gets stored in a
// separate file in kopia
type Metadata struct {
	FileName string `json:"filename,omitempty"`
	// SharingMode denotes what the current mode of sharing is for the object.
	// - inherited: permissions same as parent permissions (no "shared" in delta)
	// - custom: use Permissions to set correct permissions ("shared" has value in delta)
	SharingMode SharingMode  `json:"permissionMode,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
	// TODO(meain): Maybe find a better name?
	LinkShares []LinkShare `json:"linkShares,omitempty"`
}

type Item struct {
	ID   string
	Data io.ReadCloser
	Mod  time.Time
}

// Deleted implements an interface function. However, OneDrive items are marked
// as deleted by adding them to the exclude list so this can always return
// false.
func (i *Item) Deleted() bool           { return false }
func (i *Item) UUID() string            { return i.ID }
func (i *Item) ToReader() io.ReadCloser { return i.Data }
func (i *Item) ModTime() time.Time      { return i.Mod }
