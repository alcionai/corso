package details

import (
	"time"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

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

// OneDriveInfo describes a oneDrive item
type OneDriveInfo struct {
	Created    time.Time `json:"created,omitempty"`
	DriveID    string    `json:"driveID,omitempty"`
	DriveName  string    `json:"driveName,omitempty"`
	IsMeta     bool      `json:"isMeta,omitempty"`
	ItemName   string    `json:"itemName,omitempty"`
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
