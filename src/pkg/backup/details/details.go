package details

import (
	"context"
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
)

// Max number of items for which we will print details. If there are
// more than this, then we just show a summary.
const maxPrintLimit = 15

// LocationIDer provides access to location information but guarantees that it
// can also generate a unique location (among items in the same service but
// possibly across data types within the service) that can be used as a key in
// maps and other structures. The unique location may be different than
// InDetails, the location used in backup details.
type LocationIDer interface {
	ID() *path.Builder
	InDetails() *path.Builder
}

type uniqueLoc struct {
	pb          *path.Builder
	prefixElems int
}

func (ul uniqueLoc) ID() *path.Builder {
	return ul.pb
}

func (ul uniqueLoc) InDetails() *path.Builder {
	return path.Builder{}.Append(ul.pb.Elements()[ul.prefixElems:]...)
}

// elementCount returns the number of non-prefix elements in the LocationIDer
// (i.e. the number of elements in the InDetails path.Builder).
func (ul uniqueLoc) elementCount() int {
	res := len(ul.pb.Elements()) - ul.prefixElems
	if res < 0 {
		res = 0
	}

	return res
}

func (ul *uniqueLoc) dir() {
	if ul.elementCount() == 0 {
		return
	}

	ul.pb = ul.pb.Dir()
}

// lastElem returns the unescaped last element in the location. If the location
// is empty returns an empty string.
func (ul uniqueLoc) lastElem() string {
	if ul.elementCount() == 0 {
		return ""
	}

	return ul.pb.LastElem()
}

// Having service-specific constructors can be kind of clunky, but in this case
// I think they'd be useful to ensure the proper args are used since this
// path.Builder is used as a key in some maps.

// NewExchangeLocationIDer builds a LocationIDer for the given category and
// folder path. The path denoted by the folders should be unique within the
// category.
func NewExchangeLocationIDer(
	category path.CategoryType,
	escapedFolders ...string,
) (uniqueLoc, error) {
	if err := path.ValidateServiceAndCategory(path.ExchangeService, category); err != nil {
		return uniqueLoc{}, clues.Wrap(err, "making exchange LocationIDer")
	}

	pb := path.Builder{}.Append(category.String()).Append(escapedFolders...)

	return uniqueLoc{
		pb:          pb,
		prefixElems: 1,
	}, nil
}

// NewOneDriveLocationIDer builds a LocationIDer for the drive and folder path.
// The path denoted by the folders should be unique within the drive.
func NewOneDriveLocationIDer(
	driveID string,
	escapedFolders ...string,
) uniqueLoc {
	pb := path.Builder{}.
		Append(path.FilesCategory.String(), driveID).
		Append(escapedFolders...)

	return uniqueLoc{
		pb:          pb,
		prefixElems: 2,
	}
}

// NewSharePointLocationIDer builds a LocationIDer for the drive and folder
// path. The path denoted by the folders should be unique within the drive.
func NewSharePointLocationIDer(
	driveID string,
	escapedFolders ...string,
) uniqueLoc {
	pb := path.Builder{}.
		Append(path.LibrariesCategory.String(), driveID).
		Append(escapedFolders...)

	return uniqueLoc{
		pb:          pb,
		prefixElems: 2,
	}
}

// --------------------------------------------------------------------------------
// Model
// --------------------------------------------------------------------------------

// DetailsModel describes what was stored in a Backup
type DetailsModel struct {
	Entries []Entry `json:"entries"`
}

// Print writes the DetailModel Entries to StdOut, in the format
// requested by the caller.
func (dm DetailsModel) PrintEntries(ctx context.Context) {
	printEntries(ctx, dm.Entries)
}

type infoer interface {
	Entry | *Entry
	// Need this here so we can access the infoType function without a type
	// assertion. See https://stackoverflow.com/a/71378366 for more details.
	infoType() ItemType
}

func printEntries[T infoer](ctx context.Context, entries []T) {
	if print.DisplayJSONFormat() {
		printJSON(ctx, entries)
	} else {
		printTable(ctx, entries)
	}
}

func printTable[T infoer](ctx context.Context, entries []T) {
	perType := map[ItemType][]print.Printable{}

	for _, ent := range entries {
		it := ent.infoType()
		ps, ok := perType[it]

		if !ok {
			ps = []print.Printable{}
		}

		perType[it] = append(ps, print.Printable(ent))
	}

	for _, ps := range perType {
		print.All(ctx, ps...)
	}
}

func printJSON[T infoer](ctx context.Context, entries []T) {
	ents := []print.Printable{}

	for _, ent := range entries {
		ents = append(ents, print.Printable(ent))
	}

	print.All(ctx, ents...)
}

// Paths returns the list of Paths for non-folder and non-meta items extracted
// from the Entries slice.
func (dm DetailsModel) Paths() []string {
	r := make([]string, 0, len(dm.Entries))

	for _, ent := range dm.Entries {
		if ent.Folder != nil || ent.isMetaFile() {
			continue
		}

		r = append(r, ent.RepoRef)
	}

	return r
}

// Items returns a slice of *ItemInfo that does not contain any FolderInfo
// entries. Required because not all folders in the details are valid resource
// paths, and we want to slice out metadata.
func (dm DetailsModel) Items() entrySet {
	res := make([]*Entry, 0, len(dm.Entries))

	for i := 0; i < len(dm.Entries); i++ {
		ent := dm.Entries[i]
		if ent.Folder != nil || ent.isMetaFile() {
			continue
		}

		res = append(res, &ent)
	}

	return res
}

// FilterMetaFiles returns a copy of the Details with all of the
// .meta files removed from the entries.
func (dm DetailsModel) FilterMetaFiles() DetailsModel {
	d2 := DetailsModel{
		Entries: []Entry{},
	}

	for _, ent := range dm.Entries {
		if !ent.isMetaFile() {
			d2.Entries = append(d2.Entries, ent)
		}
	}

	return d2
}

// SumNonMetaFileSizes returns the total size of items excluding all the
// .meta files from the items.
func (dm DetailsModel) SumNonMetaFileSizes() int64 {
	var size int64

	// Items will provide only files and filter out folders
	for _, ent := range dm.FilterMetaFiles().Items() {
		size += ent.size()
	}

	return size
}

// Check if a file is a metadata file. These are used to store
// additional data like permissions (in case of Drive items) and are
// not to be treated as regular files.
func (de Entry) isMetaFile() bool {
	// sharepoint types not needed, since sharepoint permissions were
	// added after IsMeta was deprecated.
	// Earlier onedrive backups used to store both metafiles and files in details.
	// So filter out just the onedrive items and check for metafiles
	return de.ItemInfo.OneDrive != nil && de.ItemInfo.OneDrive.IsMeta
}

// ---------------------------------------------------------------------------
// Builder
// ---------------------------------------------------------------------------

// Builder should be used to create a details model.
type Builder struct {
	d            Details
	mu           sync.Mutex       `json:"-"`
	knownFolders map[string]Entry `json:"-"`
}

func (b *Builder) Add(
	repoRef path.Path,
	locationRef *path.Builder,
	updated bool,
	info ItemInfo,
) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	entry, err := b.d.add(
		repoRef,
		locationRef,
		updated,
		info)
	if err != nil {
		return clues.Wrap(err, "adding entry to details")
	}

	if err := b.addFolderEntries(
		repoRef.ToBuilder().Dir(),
		locationRef,
		entry,
	); err != nil {
		return clues.Wrap(err, "adding folder entries")
	}

	return nil
}

func (b *Builder) addFolderEntries(
	repoRef, locationRef *path.Builder,
	entry Entry,
) error {
	if len(repoRef.Elements()) < len(locationRef.Elements()) {
		return clues.New("RepoRef shorter than LocationRef").
			With("repo_ref", repoRef, "location_ref", locationRef)
	}

	if b.knownFolders == nil {
		b.knownFolders = map[string]Entry{}
	}

	// Need a unique location because we want to have separate folders for
	// different drives and categories even if there's duplicate folder names in
	// them.
	uniqueLoc, err := entry.uniqueLocation(locationRef)
	if err != nil {
		return clues.Wrap(err, "getting LocationIDer")
	}

	for uniqueLoc.elementCount() > 0 {
		mapKey := uniqueLoc.ID().ShortRef()

		name := uniqueLoc.lastElem()
		if len(name) == 0 {
			return clues.New("folder with no display name").
				With("repo_ref", repoRef, "location_ref", uniqueLoc.InDetails())
		}

		shortRef := repoRef.ShortRef()
		rr := repoRef.String()

		// Get the parent of this entry to add as the LocationRef for the folder.
		uniqueLoc.dir()

		repoRef = repoRef.Dir()
		parentRef := repoRef.ShortRef()

		folder, ok := b.knownFolders[mapKey]
		if !ok {
			loc := uniqueLoc.InDetails().String()

			folder = Entry{
				RepoRef:     rr,
				ShortRef:    shortRef,
				ParentRef:   parentRef,
				LocationRef: loc,
				ItemInfo: ItemInfo{
					Folder: &FolderInfo{
						ItemType: FolderItem,
						// TODO(ashmrtn): Use the item type returned by the entry once
						// SharePoint properly sets it.
						DisplayName: name,
					},
				},
			}

			if err := entry.updateFolder(folder.Folder); err != nil {
				return clues.Wrap(err, "adding folder").
					With("parent_repo_ref", repoRef, "location_ref", loc)
			}
		}

		folder.Folder.Size += entry.size()
		folder.Updated = folder.Updated || entry.Updated

		itemModified := entry.Modified()
		if folder.Folder.Modified.Before(itemModified) {
			folder.Folder.Modified = itemModified
		}

		// Always update the map because we're storing structs not pointers to
		// structs.
		b.knownFolders[mapKey] = folder
	}

	return nil
}

func (b *Builder) Details() *Details {
	b.mu.Lock()
	defer b.mu.Unlock()

	ents := make([]Entry, len(b.d.Entries))
	copy(ents, b.d.Entries)

	// Write the cached folder entries to details
	details := &Details{
		DetailsModel{
			Entries: append(ents, maps.Values(b.knownFolders)...),
		},
	}

	return details
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

func (d *Details) add(
	repoRef path.Path,
	locationRef *path.Builder,
	updated bool,
	info ItemInfo,
) (Entry, error) {
	if locationRef == nil {
		return Entry{}, clues.New("nil LocationRef").With("repo_ref", repoRef)
	}

	entry := Entry{
		RepoRef:     repoRef.String(),
		ShortRef:    repoRef.ShortRef(),
		ParentRef:   repoRef.ToBuilder().Dir().ShortRef(),
		LocationRef: locationRef.String(),
		ItemRef:     repoRef.Item(),
		Updated:     updated,
		ItemInfo:    info,
	}

	// Use the item name and the path for the ShortRef. This ensures that renames
	// within a directory generate unique ShortRefs.
	if info.infoType() == OneDriveItem || info.infoType() == SharePointLibrary {
		if info.OneDrive == nil && info.SharePoint == nil {
			return entry, clues.New("item is not SharePoint or OneDrive type")
		}

		filename := ""
		if info.OneDrive != nil {
			filename = info.OneDrive.ItemName
		} else if info.SharePoint != nil {
			filename = info.SharePoint.ItemName
		}

		// Make the new path contain all display names and then the M365 item ID.
		// This ensures the path will be unique, thus ensuring the ShortRef will be
		// unique.
		//
		// If we appended the file's display name to the path then it's possible
		// for a folder in the parent directory to have the same display name as the
		// M365 ID of this file and also have a subfolder in the folder with a
		// display name that matches the file's display name. That would result in
		// duplicate ShortRefs, which we can't allow.
		elements := repoRef.Elements()
		elements = append(elements[:len(elements)-1], filename, repoRef.Item())
		entry.ShortRef = path.Builder{}.Append(elements...).ShortRef()

		// clean metadata suffixes from item refs
		entry.ItemRef = withoutMetadataSuffix(entry.ItemRef)
	}

	d.Entries = append(d.Entries, entry)

	return entry, nil
}

// Marshal complies with the marshaller interface in streamStore.
func (d *Details) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

// UnmarshalTo produces a func that complies with the unmarshaller type in streamStore.
func UnmarshalTo(d *Details) func(io.ReadCloser) error {
	return func(rc io.ReadCloser) error {
		return json.NewDecoder(rc).Decode(d)
	}
}

// remove metadata file suffixes from the string.
// assumes only one suffix is applied to any given id.
func withoutMetadataSuffix(id string) string {
	id = strings.TrimSuffix(id, metadata.DirMetaFileSuffix)
	id = strings.TrimSuffix(id, metadata.MetaFileSuffix)
	id = strings.TrimSuffix(id, metadata.DataFileSuffix)

	return id
}

// --------------------------------------------------------------------------------
// Entry
// --------------------------------------------------------------------------------

// Add a new type so we can transparently use PrintAll in different situations.
type entrySet []*Entry

func (ents entrySet) PrintEntries(ctx context.Context) {
	printEntries(ctx, ents)
}

// MaybePrintEntries is same as PrintEntries, but only prints if we
// have less than 15 items or is not json output.
func (ents entrySet) MaybePrintEntries(ctx context.Context) {
	if len(ents) > maxPrintLimit &&
		!print.DisplayJSONFormat() &&
		!print.DisplayVerbose() {
		// TODO: Should we detect if the user is piping the output and
		// print if that is the case?
		print.Outf(ctx, "Restored %d items.", len(ents))
	} else {
		printEntries(ctx, ents)
	}
}

// Entry describes a single item stored in a Backup
type Entry struct {
	// RepoRef is the full storage path of the item in Kopia
	RepoRef   string `json:"repoRef"`
	ShortRef  string `json:"shortRef"`
	ParentRef string `json:"parentRef,omitempty"`

	// LocationRef contains the logical path structure by its human-readable
	// display names.  IE:  If an item is located at "/Inbox/Important", we
	// hold that string in the LocationRef, while the actual IDs of each
	// container are used for the RepoRef.
	// LocationRef only holds the container values, and does not include
	// the metadata prefixes (tenant, service, owner, etc) found in the
	// repoRef.
	// Currently only implemented for Exchange Calendars.
	LocationRef string `json:"locationRef,omitempty"`

	// ItemRef contains the stable id of the item itself.  ItemRef is not
	// guaranteed to be unique within a repository.  Uniqueness guarantees
	// maximally inherit from the source item. Eg: Entries for m365 mail items
	// are only as unique as m365 mail item IDs themselves.
	ItemRef string `json:"itemRef,omitempty"`

	// Indicates the item was added or updated in this backup
	// Always `true` for full backups
	Updated bool `json:"updated"`

	ItemInfo
}

// ToLocationIDer takes a backup version and produces the unique location for
// this entry if possible. Reasons it may not be possible to produce the unique
// location include an unsupported backup version or missing information.
func (de Entry) ToLocationIDer(backupVersion int) (LocationIDer, error) {
	if len(de.LocationRef) > 0 {
		baseLoc, err := path.Builder{}.SplitUnescapeAppend(de.LocationRef)
		if err != nil {
			return nil, clues.Wrap(err, "parsing base location info").
				With("location_ref", de.LocationRef)
		}

		// Individual services may add additional info to the base and return that.
		return de.ItemInfo.uniqueLocation(baseLoc)
	}

	if backupVersion >= version.OneDrive7LocationRef ||
		(de.ItemInfo.infoType() != OneDriveItem &&
			de.ItemInfo.infoType() != SharePointLibrary) {
		return nil, clues.New("no previous location for entry")
	}

	// This is a little hacky, but we only want to try to extract the old
	// location if it's OneDrive or SharePoint libraries and it's known to
	// be an older backup version.
	//
	// TODO(ashmrtn): Remove this code once OneDrive/SharePoint libraries
	// LocationRef code has been out long enough that all delta tokens for
	// previous backup versions will have expired. At that point, either
	// we'll do a full backup (token expired, no newer backups) or have a
	// backup of a higher version with the information we need.
	rr, err := path.FromDataLayerPath(de.RepoRef, true)
	if err != nil {
		return nil, clues.Wrap(err, "getting item RepoRef")
	}

	p, err := path.ToDrivePath(rr)
	if err != nil {
		return nil, clues.New("converting RepoRef to drive path")
	}

	baseLoc := path.Builder{}.Append(p.Root).Append(p.Folders...)

	// Individual services may add additional info to the base and return that.
	return de.ItemInfo.uniqueLocation(baseLoc)
}

// --------------------------------------------------------------------------------
// CLI Output
// --------------------------------------------------------------------------------

// interface compliance checks
var _ print.Printable = &Entry{}

// MinimumPrintable DetailsEntries is a passthrough func, because no
// reduction is needed for the json output.
func (de Entry) MinimumPrintable() any {
	return de
}

// Headers returns the human-readable names of properties in a DetailsEntry
// for printing out to a terminal in a columnar display.
func (de Entry) Headers() []string {
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
func (de Entry) Values() []string {
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

// ItemTypes are enumerated by service (hundredth digit) and data type (ones digit).
// Ex: exchange is 00x where x is the data type.  Sharepoint is 10x, and etc.
// Every item info struct should get its own hundredth enumeration entry.
// Every item category for that service should get its own entry (even if differences
// between types aren't apparent on initial implementation, this future-proofs
// against breaking changes).
// Entries should not be rearranged.
// Additionally, any itemType directly assigned a number should not be altered.
// This applies to  OneDriveItem and FolderItem
const (
	UnknownType ItemType = iota // 0, global unknown value

	// Exchange (00x)
	ExchangeContact
	ExchangeEvent
	ExchangeMail
	// SharePoint (10x)
	SharePointLibrary ItemType = iota + 97 // 100
	SharePointList                         // 101...
	SharePointPage

	// OneDrive (20x)
	OneDriveItem ItemType = 205

	// Folder Management(30x)
	FolderItem ItemType = 306
)

func UpdateItem(item *ItemInfo, newLocPath *path.Builder) {
	// Only OneDrive and SharePoint have information about parent folders
	// contained in them.
	// Can't switch based on infoType because that's been unstable.
	if item.Exchange != nil {
		item.Exchange.UpdateParentPath(newLocPath)
	} else if item.SharePoint != nil {
		// SharePoint used to store library items with the OneDriveItem ItemType.
		// Start switching them over as we see them since there's no point in
		// keeping the old format.
		if item.SharePoint.ItemType == OneDriveItem {
			item.SharePoint.ItemType = SharePointLibrary
		}

		item.SharePoint.UpdateParentPath(newLocPath)
	} else if item.OneDrive != nil {
		item.OneDrive.UpdateParentPath(newLocPath)
	}
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

func (i ItemInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	switch {
	case i.Exchange != nil:
		return i.Exchange.uniqueLocation(baseLoc)

	case i.OneDrive != nil:
		return i.OneDrive.uniqueLocation(baseLoc)

	case i.SharePoint != nil:
		return i.SharePoint.uniqueLocation(baseLoc)

	default:
		return nil, clues.New("unsupported type")
	}
}

func (i ItemInfo) updateFolder(f *FolderInfo) error {
	switch {
	case i.Exchange != nil:
		return i.Exchange.updateFolder(f)

	case i.OneDrive != nil:
		return i.OneDrive.updateFolder(f)

	case i.SharePoint != nil:
		return i.SharePoint.updateFolder(f)

	default:
		return clues.New("unsupported type")
	}
}

type FolderInfo struct {
	ItemType    ItemType  `json:"itemType,omitempty"`
	DisplayName string    `json:"displayName"`
	Modified    time.Time `json:"modified,omitempty"`
	Size        int64     `json:"size,omitempty"`
	DataType    ItemType  `json:"dataType,omitempty"`
	DriveName   string    `json:"driveName,omitempty"`
	DriveID     string    `json:"driveID,omitempty"`
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
	Recipient   []string  `json:"recipient,omitempty"`
	ParentPath  string    `json:"parentPath,omitempty"`
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
		return []string{"Sender", "Folder", "Subject", "Received"}
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
			dttm.FormatToTabularDisplay(i.EventStart),
			dttm.FormatToTabularDisplay(i.EventEnd),
			strconv.FormatBool(i.EventRecurs),
		}

	case ExchangeContact:
		return []string{i.ContactName}

	case ExchangeMail:
		return []string{
			i.Sender, i.ParentPath, i.Subject,
			dttm.FormatToTabularDisplay(i.Received),
		}
	}

	return []string{}
}

func (i *ExchangeInfo) UpdateParentPath(newLocPath *path.Builder) {
	i.ParentPath = newLocPath.String()
}

func (i *ExchangeInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	var category path.CategoryType

	switch i.ItemType {
	case ExchangeEvent:
		category = path.EventsCategory
	case ExchangeContact:
		category = path.ContactsCategory
	case ExchangeMail:
		category = path.EmailCategory
	}

	loc, err := NewExchangeLocationIDer(category, baseLoc.Elements()...)

	return &loc, err
}

func (i *ExchangeInfo) updateFolder(f *FolderInfo) error {
	// Use a switch instead of a rather large if-statement. Just make sure it's an
	// Exchange type. If it's not return an error.
	switch i.ItemType {
	case ExchangeContact, ExchangeEvent, ExchangeMail:
	default:
		return clues.New("unsupported non-Exchange ItemType").
			With("item_type", i.ItemType)
	}

	f.DataType = i.ItemType

	return nil
}

// SharePointInfo describes a sharepoint item
type SharePointInfo struct {
	Created    time.Time `json:"created,omitempty"`
	DriveName  string    `json:"driveName,omitempty"`
	DriveID    string    `json:"driveID,omitempty"`
	ItemName   string    `json:"itemName,omitempty"`
	ItemType   ItemType  `json:"itemType,omitempty"`
	Modified   time.Time `json:"modified,omitempty"`
	Owner      string    `json:"owner,omitempty"`
	ParentPath string    `json:"parentPath,omitempty"`
	Size       int64     `json:"size,omitempty"`
	WebURL     string    `json:"webUrl,omitempty"`
	SiteID     string    `json:"siteID,omitempty"`
}

// Headers returns the human-readable names of properties in a SharePointInfo
// for printing out to a terminal in a columnar display.
func (i SharePointInfo) Headers() []string {
	return []string{"ItemName", "Library", "ParentPath", "Size", "Owner", "Created", "Modified"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i SharePointInfo) Values() []string {
	return []string{
		i.ItemName,
		i.DriveName,
		i.ParentPath,
		humanize.Bytes(uint64(i.Size)),
		i.Owner,
		dttm.FormatToTabularDisplay(i.Created),
		dttm.FormatToTabularDisplay(i.Modified),
	}
}

func (i *SharePointInfo) UpdateParentPath(newLocPath *path.Builder) {
	i.ParentPath = newLocPath.PopFront().String()
}

func (i *SharePointInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	if len(i.DriveID) == 0 {
		return nil, clues.New("empty drive ID")
	}

	loc := NewSharePointLocationIDer(i.DriveID, baseLoc.Elements()...)

	return &loc, nil
}

func (i *SharePointInfo) updateFolder(f *FolderInfo) error {
	// TODO(ashmrtn): Change to just SharePointLibrary when the code that
	// generates the item type is fixed.
	if i.ItemType == OneDriveItem || i.ItemType == SharePointLibrary {
		return updateFolderWithinDrive(SharePointLibrary, i.DriveName, i.DriveID, f)
	}

	return clues.New("unsupported non-SharePoint ItemType").With("item_type", i.ItemType)
}

// OneDriveInfo describes a oneDrive item
type OneDriveInfo struct {
	Created    time.Time      `json:"created,omitempty"`
	DriveID    string         `json:"driveID,omitempty"`
	DriveName  string         `json:"driveName,omitempty"`
	IsMeta     bool           `json:"isMeta,omitempty"`
	ItemName   string         `json:"itemName,omitempty"`
	ItemType   ItemType       `json:"itemType,omitempty"`
	Modified   time.Time      `json:"modified,omitempty"`
	Owner      string         `json:"owner,omitempty"`
	ParentPath string         `json:"parentPath"`
	Size       int64          `json:"size,omitempty"`
	Extension  *ExtensionInfo `json:"extensionData,omitempty"`
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
		dttm.FormatToTabularDisplay(i.Created),
		dttm.FormatToTabularDisplay(i.Modified),
	}
}

func (i *OneDriveInfo) UpdateParentPath(newLocPath *path.Builder) {
	i.ParentPath = newLocPath.PopFront().String()
}

func (i *OneDriveInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	if len(i.DriveID) == 0 {
		return nil, clues.New("empty drive ID")
	}

	loc := NewOneDriveLocationIDer(i.DriveID, baseLoc.Elements()...)

	return &loc, nil
}

func (i *OneDriveInfo) updateFolder(f *FolderInfo) error {
	return updateFolderWithinDrive(OneDriveItem, i.DriveName, i.DriveID, f)
}

func updateFolderWithinDrive(
	t ItemType,
	driveName, driveID string,
	f *FolderInfo,
) error {
	if len(driveName) == 0 {
		return clues.New("empty drive name")
	} else if len(driveID) == 0 {
		return clues.New("empty drive ID")
	}

	f.DriveName = driveName
	f.DriveID = driveID
	f.DataType = t

	return nil
}

// ExtensionInfo describes extension data associated with an item
// TODO: Expose this store behind an interface which can synchrnoize access to the
// underlying map.
type ExtensionInfo struct {
	Data map[string]any `json:"data,omitempty"`
}
