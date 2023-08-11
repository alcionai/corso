package details

import (
	"time"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

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
