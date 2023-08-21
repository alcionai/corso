package details

import (
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

// NewGroupsLocationIDer builds a LocationIDer for the groups.
func NewGroupsLocationIDer(
	category path.CategoryType,
	driveID string,
	escapedFolders ...string,
) (uniqueLoc, error) {
	// TODO(meain): path fixes
	if err := path.ValidateServiceAndCategory(path.GroupsService, category); err != nil {
		return uniqueLoc{}, clues.Wrap(err, "making groups LocationIDer")
	}

	pb := path.Builder{}.Append(category.String())
	prefixElems := 1

	if driveID != "" { // non sp paths don't have driveID
		pb.Append(driveID)

		prefixElems = 2
	}

	pb.Append(escapedFolders...)

	return uniqueLoc{pb, prefixElems}, nil
}

// GroupsInfo describes a groups item
// TODO(meain): Add channel name and id
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
	var category path.CategoryType

	switch i.ItemType {
	case SharePointLibrary:
		category = path.LibrariesCategory

		if len(i.DriveID) == 0 {
			return nil, clues.New("empty drive ID")
		}
	}

	loc, err := NewGroupsLocationIDer(category, i.DriveID, baseLoc.Elements()...)

	return &loc, err
}

func (i *GroupsInfo) updateFolder(f *FolderInfo) error {
	// TODO(meain): path updates if any
	if i.ItemType == SharePointLibrary {
		return updateFolderWithinDrive(SharePointLibrary, i.DriveName, i.DriveID, f)
	}

	return clues.New("unsupported ItemType for GroupsInfo").With("item_type", i.ItemType)
}
