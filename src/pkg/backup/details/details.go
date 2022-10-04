package details

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/model"
)

type FolderEntry struct {
	RepoRef   string
	ShortRef  string
	ParentRef string
	Info      ItemInfo
}

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
	if print.JSONFormat() {
		printJSON(ctx, dm)
	} else {
		printTable(ctx, dm)
	}
}

func printTable(ctx context.Context, dm DetailsModel) {
	perType := map[ItemType][]print.Printable{}

	for _, de := range dm.Entries {
		it := de.infoType()
		ps, ok := perType[it]

		if !ok {
			ps = []print.Printable{}
		}

		perType[it] = append(ps, print.Printable(de))
	}

	for _, ps := range perType {
		print.All(ctx, ps...)
	}
}

func printJSON(ctx context.Context, dm DetailsModel) {
	ents := []print.Printable{}

	for _, ent := range dm.Entries {
		ents = append(ents, print.Printable(ent))
	}

	print.All(ctx, ents...)
}

// Paths returns the list of Paths for non-folder items extracted from the
// Entries slice.
func (dm DetailsModel) Paths() []string {
	r := make([]string, 0, len(dm.Entries))

	for _, ent := range dm.Entries {
		if ent.Folder != nil {
			continue
		}

		r = append(r, ent.RepoRef)
	}

	return r
}

// Items returns a slice of *ItemInfo that does not contain any FolderInfo
// entries. Required because not all folders in the details are valid resource
// paths.
func (dm DetailsModel) Items() []*DetailsEntry {
	res := make([]*DetailsEntry, 0, len(dm.Entries))

	for i := 0; i < len(dm.Entries); i++ {
		if dm.Entries[i].Folder != nil {
			continue
		}

		res = append(res, &dm.Entries[i])
	}

	return res
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
	mu           sync.Mutex          `json:"-"`
	knownFolders map[string]struct{} `json:"-"`
}

func (d *Details) Add(repoRef, shortRef, parentRef string, info ItemInfo) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entries = append(d.Entries, DetailsEntry{
		RepoRef:   repoRef,
		ShortRef:  shortRef,
		ParentRef: parentRef,
		ItemInfo:  info,
	})
}

// AddFolders adds entries for the given folders. It skips adding entries that
// have been added by previous calls.
func (d *Details) AddFolders(folders []FolderEntry) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.knownFolders == nil {
		d.knownFolders = map[string]struct{}{}
	}

	for _, folder := range folders {
		if _, ok := d.knownFolders[folder.ShortRef]; ok {
			// Entry already exists, nothing to do.
			continue
		}

		d.knownFolders[folder.ShortRef] = struct{}{}
		d.Entries = append(d.Entries, DetailsEntry{
			RepoRef:   folder.RepoRef,
			ShortRef:  folder.ShortRef,
			ParentRef: folder.ParentRef,
			ItemInfo:  folder.Info,
		})
	}
}

// --------------------------------------------------------------------------------
// Entry
// --------------------------------------------------------------------------------

// DetailsEntry describes a single item stored in a Backup
type DetailsEntry struct {
	// TODO: `RepoRef` is currently the full path to the item in Kopia
	// This can be optimized.
	RepoRef   string `json:"repoRef"`
	ShortRef  string `json:"shortRef"`
	ParentRef string `json:"parentRef,omitempty"`
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
	hs := []string{"Reference"}

	if de.ItemInfo.Folder != nil {
		hs = append(hs, de.ItemInfo.Folder.Headers()...)
	}

	if de.ItemInfo.Exchange != nil {
		hs = append(hs, de.ItemInfo.Exchange.Headers()...)
	}

	if de.ItemInfo.Sharepoint != nil {
		hs = append(hs, de.ItemInfo.Sharepoint.Headers()...)
	}

	if de.ItemInfo.OneDrive != nil {
		hs = append(hs, de.ItemInfo.OneDrive.Headers()...)
	}

	return hs
}

// Values returns the values matching the Headers list.
func (de DetailsEntry) Values() []string {
	vs := []string{de.ShortRef}

	if de.ItemInfo.Folder != nil {
		vs = append(vs, de.ItemInfo.Folder.Values()...)
	}

	if de.ItemInfo.Exchange != nil {
		vs = append(vs, de.ItemInfo.Exchange.Values()...)
	}

	if de.ItemInfo.Sharepoint != nil {
		vs = append(vs, de.ItemInfo.Sharepoint.Values()...)
	}

	if de.ItemInfo.OneDrive != nil {
		vs = append(vs, de.ItemInfo.OneDrive.Values()...)
	}

	return vs
}

type ItemType int

const (
	UnknownType ItemType = iota

	// separate each service by a factor of 100 for padding
	ExchangeContact
	ExchangeEvent
	ExchangeMail

	SharepointItem ItemType = iota + 100

	OneDriveItem ItemType = iota + 200

	FolderItem ItemType = iota + 300
)

// ItemInfo is a oneOf that contains service specific
// information about the item it tracks
type ItemInfo struct {
	Folder     *FolderInfo     `json:"folder,omitempty"`
	Exchange   *ExchangeInfo   `json:"exchange,omitempty"`
	Sharepoint *SharepointInfo `json:"sharepoint,omitempty"`
	OneDrive   *OneDriveInfo   `json:"oneDrive,omitempty"`
}

// typedInfo should get embedded in each sesrvice type to track
// the type of item it stores for multi-item service support.

// infoType provides internal categorization for collecting like-typed ItemInfos.
// It should return the most granular value type (ex: "event" for an exchange
// calendar event).
func (i ItemInfo) infoType() ItemType {
	switch {
	case i.Folder != nil:
		return i.Folder.ItemType

	case i.Exchange != nil:
		return i.Exchange.ItemType

	case i.Sharepoint != nil:
		return i.Sharepoint.ItemType

	case i.OneDrive != nil:
		return i.OneDrive.ItemType
	}

	return UnknownType
}

type FolderInfo struct {
	ItemType    ItemType `json:"itemType,omitempty"`
	DisplayName string   `json:"displayName"`
}

func (i FolderInfo) Headers() []string {
	return []string{"Display Name"}
}

func (i FolderInfo) Values() []string {
	return []string{i.DisplayName}
}

// ExchangeInfo describes an exchange item
type ExchangeInfo struct {
	ItemType    ItemType  `json:"itemType,omitempty"`
	Sender      string    `json:"sender,omitempty"`
	Subject     string    `json:"subject,omitempty"`
	Received    time.Time `json:"received,omitempty"`
	EventStart  time.Time `json:"eventStart,omitempty"`
	Organizer   string    `json:"organizer,omitempty"`
	ContactName string    `json:"contactName,omitempty"`
	EventRecurs bool      `json:"eventRecurs,omitempty"`
}

// Headers returns the human-readable names of properties in an ExchangeInfo
// for printing out to a terminal in a columnar display.
func (i ExchangeInfo) Headers() []string {
	switch i.ItemType {
	case ExchangeEvent:
		return []string{"Organizer", "Subject", "Starts", "Recurring"}

	case ExchangeContact:
		return []string{"Contact Name"}

	case ExchangeMail:
		return []string{"Sender", "Subject", "Received"}
	}

	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i ExchangeInfo) Values() []string {
	switch i.ItemType {
	case ExchangeEvent:
		return []string{
			i.Organizer,
			i.Subject,
			common.FormatTabularDisplayTime(i.EventStart),
			strconv.FormatBool(i.EventRecurs),
		}

	case ExchangeContact:
		return []string{i.ContactName}

	case ExchangeMail:
		return []string{i.Sender, i.Subject, common.FormatTabularDisplayTime(i.Received)}
	}

	return []string{}
}

// SharepointInfo describes a sharepoint item
// TODO: Implement this. This is currently here
// just to illustrate usage
type SharepointInfo struct {
	ItemType ItemType `json:"itemType,omitempty"`
}

// Headers returns the human-readable names of properties in a SharepointInfo
// for printing out to a terminal in a columnar display.
func (i SharepointInfo) Headers() []string {
	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i SharepointInfo) Values() []string {
	return []string{}
}

// OneDriveInfo describes a oneDrive item
type OneDriveInfo struct {
	ItemType   ItemType  `json:"itemType,omitempty"`
	ParentPath string    `json:"parentPath"`
	ItemName   string    `json:"itemName"`
	Size       int64     `json:"size,omitempty"`
	Created    time.Time `json:"created,omitempty"`
	Modified   time.Time `json:"modified,omitempty"`
}

// Headers returns the human-readable names of properties in a OneDriveInfo
// for printing out to a terminal in a columnar display.
func (i OneDriveInfo) Headers() []string {
	return []string{"ItemName", "ParentPath", "Size", "Created", "Modified"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i OneDriveInfo) Values() []string {
	return []string{
		i.ItemName, i.ParentPath, humanize.Bytes(uint64(i.Size)),
		common.FormatTabularDisplayTime(i.Created), common.FormatTabularDisplayTime(i.Modified),
	}
}
