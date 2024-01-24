package details

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
)

// Add a new type so we can transparently use PrintAll in different situations.
type entrySet []*Entry

func (ents entrySet) PrintEntries(ctx context.Context) {
	printEntries(ctx, ents)
}

// MaybePrintEntries is same as PrintEntries, but only prints if we
// have less than 15 items or is not json output.
func (ents entrySet) MaybePrintEntries(ctx context.Context) {
	if len(ents) <= maxPrintLimit ||
		print.DisplayJSONFormat() ||
		print.DisplayVerbose() {
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

	ItemInfo
}

// ToLocationIDer takes a backup version and produces the unique location for
// this entry if possible. Reasons it may not be possible to produce the unique
// location include an unsupported backup version or missing information.
//
// TODO(ashmrtn): Remove this function completely if we ever decide to sunset
// older corso versions that didn't populate LocationRef.
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

	rr, err := path.FromDataLayerPath(de.RepoRef, true)
	if err != nil {
		return nil, clues.Wrap(err, "getting item RepoRef").
			With("repo_ref", de.RepoRef)
	}

	var baseLoc *path.Builder

	switch de.ItemInfo.infoType() {
	case ExchangeEvent:
		if backupVersion >= 2 {
			return nil, clues.New("no previous location for calendar entry").
				With("repo_ref", rr)
		}

		fallthrough
	case ExchangeMail, ExchangeContact:
		baseLoc = path.Builder{}.Append(rr.Folders()...)

	case OneDriveItem, SharePointLibrary:
		if backupVersion >= version.OneDrive7LocationRef {
			return nil, clues.New("no previous location for drive entry").
				With("repo_ref", rr)
		}

		p, err := path.ToDrivePath(rr)
		if err != nil {
			return nil, clues.New("converting RepoRef to drive path").
				With("repo_ref", rr)
		}

		baseLoc = path.Builder{}.Append(p.Root).Append(p.Folders...)
	}

	if baseLoc == nil {
		return nil, clues.New("unable to extract LocationRef from RepoRef").
			With("repo_ref", rr)
	}

	// Individual services may add additional info to the base and return that.
	return de.ItemInfo.uniqueLocation(baseLoc)
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
func (de Entry) Headers(skipID bool) []string {
	var hs []string

	switch {
	case de.ItemInfo.Folder != nil:
		hs = de.ItemInfo.Folder.Headers()
	case de.ItemInfo.Exchange != nil:
		hs = de.ItemInfo.Exchange.Headers()
	case de.ItemInfo.SharePoint != nil:
		hs = de.ItemInfo.SharePoint.Headers()
	case de.ItemInfo.OneDrive != nil:
		hs = de.ItemInfo.OneDrive.Headers()
	case de.ItemInfo.Groups != nil:
		hs = de.ItemInfo.Groups.Headers()
	case de.ItemInfo.TeamsChats != nil:
		hs = de.ItemInfo.TeamsChats.Headers()
	default:
		hs = []string{"ERROR - Service not recognized"}
	}

	if skipID {
		return hs
	}

	return append([]string{"ID"}, hs...)
}

// Values returns the values matching the Headers list.
func (de Entry) Values(skipID bool) []string {
	var vs []string

	switch {
	case de.ItemInfo.Folder != nil:
		vs = de.ItemInfo.Folder.Values()
	case de.ItemInfo.Exchange != nil:
		vs = de.ItemInfo.Exchange.Values()
	case de.ItemInfo.SharePoint != nil:
		vs = de.ItemInfo.SharePoint.Values()
	case de.ItemInfo.OneDrive != nil:
		vs = de.ItemInfo.OneDrive.Values()
	case de.ItemInfo.Groups != nil:
		vs = de.ItemInfo.Groups.Values()
	case de.ItemInfo.TeamsChats != nil:
		vs = de.ItemInfo.TeamsChats.Values()
	default:
		vs = []string{"ERROR - Service not recognized"}
	}

	if skipID {
		return vs
	}

	return append([]string{de.ShortRef}, vs...)
}
