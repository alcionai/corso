package details

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph/metadata"
)

// Max number of items for which we will print details. If there are
// more than this, then we just show a summary.
const maxPrintLimit = 50

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
		ItemInfo:    info,
	}

	// Use the item name and the path for the ShortRef. This ensures that renames
	// within a directory generate unique ShortRefs.
	if info.isDriveItem() {
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

// ---------------------------------------------------------------------------
// LocationIDer
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

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

// ExtensionData stores extension data associated with an item
type ExtensionData struct {
	Data map[string]any `json:"data,omitempty"`
}
