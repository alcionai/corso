package details

import (
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

// NewGroupsLocationIDer builds a LocationIDer for the groups.
func NewGroupsLocationIDer(
	driveID string,
	escapedFolders ...string,
) uniqueLoc {
	// TODO: implement
	return uniqueLoc{}
}

// GroupsInfo describes a groups item
type GroupsInfo struct {
	Created    time.Time `json:"created,omitempty"`
	DriveName  string    `json:"driveName,omitempty"`
	DriveID    string    `json:"driveID,omitempty"`
	ItemName   string    `json:"itemName,omitempty"`
	ItemType   ItemType  `json:"itemType,omitempty"`
	Modified   time.Time `json:"modified,omitempty"`
	Owner      string    `json:"owner,omitempty"`
	ParentPath string    `json:"parentPath,omitempty"`
	Size       int64     `json:"size,omitempty"`
}

// Headers returns the human-readable names of properties in a SharePointInfo
// for printing out to a terminal in a columnar display.
func (i GroupsInfo) Headers() []string {
	return []string{"Created", "Modified"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i GroupsInfo) Values() []string {
	return []string{
		dttm.FormatToTabularDisplay(i.Created),
		dttm.FormatToTabularDisplay(i.Modified),
	}
}

func (i *GroupsInfo) UpdateParentPath(newLocPath *path.Builder) {
	i.ParentPath = newLocPath.PopFront().String()
}

func (i *GroupsInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	return nil, clues.New("not yet implemented")
}

func (i *GroupsInfo) updateFolder(f *FolderInfo) error {
	return clues.New("not yet implemented")
}
