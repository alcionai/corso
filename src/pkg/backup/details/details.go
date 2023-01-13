package details

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/path"
)

type folderEntry struct {
	RepoRef   string
	ShortRef  string
	ParentRef string
	Updated   bool
	Info      ItemInfo
}

// --------------------------------------------------------------------------------
// Model
// --------------------------------------------------------------------------------

// DetailsModel describes what was stored in a Backup
type DetailsModel struct {
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

// Builder should be used to create a details model.
type Builder struct {
	d            Details
	mu           sync.Mutex             `json:"-"`
	knownFolders map[string]folderEntry `json:"-"`
}

func (b *Builder) Add(repoRef, shortRef, parentRef string, updated bool, info ItemInfo) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.d.add(repoRef, shortRef, parentRef, updated, info)
}

func (b *Builder) Details() *Details {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Write the cached folder entries to details
	for _, folder := range b.knownFolders {
		b.d.addFolder(folder)
	}

	return &b.d
}

// TODO(ashmrtn): If we never need to pre-populate the modified time of a folder
// we should just merge this with AddFoldersForItem, have Add call
// AddFoldersForItem, and unexport AddFoldersForItem.
func FolderEntriesForPath(parent *path.Builder) []folderEntry {
	folders := []folderEntry{}

	for len(parent.Elements()) > 0 {
		nextParent := parent.Dir()

		folders = append(folders, folderEntry{
			RepoRef:   parent.String(),
			ShortRef:  parent.ShortRef(),
			ParentRef: nextParent.ShortRef(),
			Info: ItemInfo{
				Folder: &FolderInfo{
					ItemType:    FolderItem,
					DisplayName: parent.Elements()[len(parent.Elements())-1],
				},
			},
		})

		parent = nextParent
	}

	return folders
}

// AddFoldersForItem adds entries for the given folders. It skips adding entries that
// have been added by previous calls.
func (b *Builder) AddFoldersForItem(folders []folderEntry, itemInfo ItemInfo, updated bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.knownFolders == nil {
		b.knownFolders = map[string]folderEntry{}
	}

	for _, folder := range folders {
		if existing, ok := b.knownFolders[folder.ShortRef]; ok {
			// We've seen this folder before for a different item.
			// Update the "cached" folder entry
			folder = existing
		}

		// Update the folder's size and modified time
		itemModified := itemInfo.Modified()

		folder.Info.Folder.Size += itemInfo.size()

		if folder.Info.Folder.Modified.Before(itemModified) {
			folder.Info.Folder.Modified = itemModified
		}

		// If the item being added was "updated" - propagate that to the
		// folder entries
		if updated {
			folder.Updated = true
		}

		b.knownFolders[folder.ShortRef] = folder
	}
}

// --------------------------------------------------------------------------------
// Details
// --------------------------------------------------------------------------------

// Details augments the core with a mutex for processing.
// Should be sliced back to d.DetailsModel for storage and
// printing.
type Details struct {
	DetailsModel
}

func (d *Details) add(repoRef, shortRef, parentRef string, updated bool, info ItemInfo) {
	d.Entries = append(d.Entries, DetailsEntry{
		RepoRef:   repoRef,
		ShortRef:  shortRef,
		ParentRef: parentRef,
		Updated:   updated,
		ItemInfo:  info,
	})
}

// addFolder adds an entry for the given folder.
func (d *Details) addFolder(folder folderEntry) {
	d.Entries = append(d.Entries, DetailsEntry{
		RepoRef:   folder.RepoRef,
		ShortRef:  folder.ShortRef,
		ParentRef: folder.ParentRef,
		ItemInfo:  folder.Info,
		Updated:   folder.Updated,
	})
}

// --------------------------------------------------------------------------------
// Entry
// --------------------------------------------------------------------------------

// DetailsEntry describes a single item stored in a Backup
type DetailsEntry struct {
	// RepoRef is the full storage path of the item in Kopia
	RepoRef   string `json:"repoRef"`
	ShortRef  string `json:"shortRef"`
	ParentRef string `json:"parentRef,omitempty"`
	// Indicates the item was added or updated in this backup
	// Always `true` for full backups
	Updated bool `json:"updated"`
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
	hs := []string{"ID"}

	if de.ItemInfo.Folder != nil {
		hs = append(hs, de.ItemInfo.Folder.Headers()...)
	}

	if de.ItemInfo.Exchange != nil {
		hs = append(hs, de.ItemInfo.Exchange.Headers()...)
	}

	if de.ItemInfo.SharePoint != nil {
		hs = append(hs, de.ItemInfo.SharePoint.Headers()...)
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

	if de.ItemInfo.SharePoint != nil {
		vs = append(vs, de.ItemInfo.SharePoint.Values()...)
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

	SharePointItem ItemType = iota + 100

	OneDriveItem ItemType = iota + 200

	FolderItem ItemType = iota + 300
)

func UpdateItem(item *ItemInfo, newPath path.Path) error {
	// Only OneDrive and SharePoint have information about parent folders
	// contained in them.
	switch item.infoType() {
	case SharePointItem:
		return item.SharePoint.UpdateParentPath(newPath)

	case OneDriveItem:
		return item.OneDrive.UpdateParentPath(newPath)
	}

	return nil
}

// ItemInfo is a oneOf that contains service specific
// information about the item it tracks
type ItemInfo struct {
	Folder     *FolderInfo     `json:"folder,omitempty"`
	Exchange   *ExchangeInfo   `json:"exchange,omitempty"`
	SharePoint *SharePointInfo `json:"sharePoint,omitempty"`
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

	case i.SharePoint != nil:
		return i.SharePoint.ItemType

	case i.OneDrive != nil:
		return i.OneDrive.ItemType
	}

	return UnknownType
}

func (i ItemInfo) size() int64 {
	switch {
	case i.Exchange != nil:
		return i.Exchange.Size

	case i.OneDrive != nil:
		return i.OneDrive.Size

	case i.SharePoint != nil:
		return i.SharePoint.Size

	case i.Folder != nil:
		return i.Folder.Size
	}

	return 0
}

func (i ItemInfo) Modified() time.Time {
	switch {
	case i.Exchange != nil:
		return i.Exchange.Modified

	case i.OneDrive != nil:
		return i.OneDrive.Modified

	case i.SharePoint != nil:
		return i.SharePoint.Modified

	case i.Folder != nil:
		return i.Folder.Modified
	}

	return time.Time{}
}

type FolderInfo struct {
	ItemType    ItemType  `json:"itemType,omitempty"`
	DisplayName string    `json:"displayName"`
	Modified    time.Time `json:"modified,omitempty"`
	Size        int64     `json:"size,omitempty"`
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
	EventEnd    time.Time `json:"eventEnd,omitempty"`
	Organizer   string    `json:"organizer,omitempty"`
	ContactName string    `json:"contactName,omitempty"`
	EventRecurs bool      `json:"eventRecurs,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	Modified    time.Time `json:"modified,omitempty"`
	Size        int64     `json:"size,omitempty"`
}

// Headers returns the human-readable names of properties in an ExchangeInfo
// for printing out to a terminal in a columnar display.
func (i ExchangeInfo) Headers() []string {
	switch i.ItemType {
	case ExchangeEvent:
		return []string{"Organizer", "Subject", "Starts", "Ends", "Recurring"}

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
			common.FormatTabularDisplayTime(i.EventEnd),
			strconv.FormatBool(i.EventRecurs),
		}

	case ExchangeContact:
		return []string{i.ContactName}

	case ExchangeMail:
		return []string{
			i.Sender, i.Subject,
			common.FormatTabularDisplayTime(i.Received),
		}
	}

	return []string{}
}

// SharePointInfo describes a sharepoint item
type SharePointInfo struct {
	Created    time.Time `json:"created,omitempty"`
	ItemName   string    `json:"itemName,omitempty"`
	DriveName  string    `json:"driveName,omitempty"`
	ItemType   ItemType  `json:"itemType,omitempty"`
	Modified   time.Time `josn:"modified,omitempty"`
	Owner      string    `json:"owner,omitempty"`
	ParentPath string    `json:"parentPath,omitempty"`
	Size       int64     `json:"size,omitempty"`
	WebURL     string    `json:"webUrl,omitempty"`
}

// Headers returns the human-readable names of properties in a SharePointInfo
// for printing out to a terminal in a columnar display.
func (i SharePointInfo) Headers() []string {
	return []string{"ItemName", "Drive", "ParentPath", "Size", "WebURL", "Created", "Modified"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i SharePointInfo) Values() []string {
	return []string{
		i.ItemName,
		i.DriveName,
		i.ParentPath,
		humanize.Bytes(uint64(i.Size)),
		i.WebURL,
		common.FormatTabularDisplayTime(i.Created),
		common.FormatTabularDisplayTime(i.Modified),
	}
}

func (i *SharePointInfo) UpdateParentPath(newPath path.Path) error {
	newParent, err := path.GetDriveFolderPath(newPath)
	if err != nil {
		return errors.Wrapf(err, "making sharepoint path from %s", newPath)
	}

	i.ParentPath = newParent

	return nil
}

// OneDriveInfo describes a oneDrive item
type OneDriveInfo struct {
	Created    time.Time `json:"created,omitempty"`
	ItemName   string    `json:"itemName,omitempty"`
	DriveName  string    `json:"driveName,omitempty"`
	ItemType   ItemType  `json:"itemType,omitempty"`
	Modified   time.Time `json:"modified,omitempty"`
	Owner      string    `json:"owner,omitempty"`
	ParentPath string    `json:"parentPath"`
	Size       int64     `json:"size,omitempty"`
}

// Headers returns the human-readable names of properties in a OneDriveInfo
// for printing out to a terminal in a columnar display.
func (i OneDriveInfo) Headers() []string {
	return []string{"ItemName", "ParentPath", "Size", "Owner", "Created", "Modified"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i OneDriveInfo) Values() []string {
	return []string{
		i.ItemName,
		i.ParentPath,
		humanize.Bytes(uint64(i.Size)),
		i.Owner,
		common.FormatTabularDisplayTime(i.Created),
		common.FormatTabularDisplayTime(i.Modified),
	}
}

func (i *OneDriveInfo) UpdateParentPath(newPath path.Path) error {
	newParent, err := path.GetDriveFolderPath(newPath)
	if err != nil {
		return errors.Wrapf(err, "making drive path from %s", newPath)
	}

	i.ParentPath = newParent

	return nil
}
