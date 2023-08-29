package details

import (
	"strconv"
	"time"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"

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
type GroupsInfo struct {
	Created    time.Time `json:"created,omitempty"`
	ItemName   string    `json:"itemName,omitempty"`
	ItemType   ItemType  `json:"itemType,omitempty"`
	Modified   time.Time `json:"modified,omitempty"`
	Owner      string    `json:"owner,omitempty"`
	ParentPath string    `json:"parentPath,omitempty"`
	Size       int64     `json:"size,omitempty"`

	// Channels Specific
	ChannelName    string    `json:"channelName,omitempty"`
	ChannelID      string    `json:"channelID,omitempty"`
	LastReplyAt    time.Time `json:"lastResponseAt,omitempty"`
	MessageCreator string    `json:"messageCreator,omitempty"`
	MessagePreview string    `json:"messagePreview,omitempty"`
	ReplyCount     int       `json:"replyCount,omitempty"`

	// SharePoint specific
	DriveName string `json:"driveName,omitempty"`
	DriveID   string `json:"driveID,omitempty"`
	SiteID    string `json:"siteID,omitempty"`
	WebURL    string `json:"webURL,omitempty"`
}

// Headers returns the human-readable names of properties in a SharePointInfo
// for printing out to a terminal in a columnar display.
func (i GroupsInfo) Headers() []string {
	switch i.ItemType {
	case SharePointLibrary:
		return []string{"ItemName", "Library", "ParentPath", "Size", "Owner", "Created", "Modified"}
	case GroupsChannelMessage:
		return []string{"Message", "Channel", "Replies", "Creator", "Created", "Last Response"}
	}

	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i GroupsInfo) Values() []string {
	switch i.ItemType {
	case SharePointLibrary:
		return []string{
			i.ItemName,
			i.DriveName,
			i.ParentPath,
			humanize.Bytes(uint64(i.Size)),
			i.Owner,
			dttm.FormatToTabularDisplay(i.Created),
			dttm.FormatToTabularDisplay(i.Modified),
		}
	case GroupsChannelMessage:
		return []string{
			i.MessagePreview,
			i.ChannelName,
			strconv.Itoa(i.ReplyCount),
			i.MessageCreator,
			dttm.FormatToTabularDisplay(i.Created),
			dttm.FormatToTabularDisplay(i.Modified),
		}
	}

	return []string{}
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
	case GroupsChannelMessage:
		category = path.ChannelMessagesCategory

		if len(i.ChannelID) == 0 {
			return nil, clues.New("empty channel ID")
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
