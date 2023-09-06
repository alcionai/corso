package details

import (
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/path"
)

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
	UnknownType ItemType = 0

	// Exchange (00x)
	ExchangeContact ItemType = 1
	ExchangeEvent   ItemType = 2
	ExchangeMail    ItemType = 3

	// SharePoint (10x)
	SharePointLibrary ItemType = 101 // also used for groups
	SharePointList    ItemType = 102
	SharePointPage    ItemType = 103

	// OneDrive (20x)
	OneDriveItem ItemType = 205

	// Folder Management(30x)
	FolderItem ItemType = 306

	// Groups/Teams(40x)
	GroupsChannelMessage ItemType = 401
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
	} else if item.Groups != nil {
		item.Groups.UpdateParentPath(newLocPath)
	}
}

// ItemInfo is a oneOf that contains service specific
// information about the item it tracks
type ItemInfo struct {
	Folder     *FolderInfo     `json:"folder,omitempty"`
	Exchange   *ExchangeInfo   `json:"exchange,omitempty"`
	SharePoint *SharePointInfo `json:"sharePoint,omitempty"`
	OneDrive   *OneDriveInfo   `json:"oneDrive,omitempty"`
	Groups     *GroupsInfo     `json:"groups,omitempty"`
	// Optional item extension data
	Extension *ExtensionData `json:"extension,omitempty"`
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

	case i.Groups != nil:
		return i.Groups.ItemType
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

	case i.Groups != nil:
		return i.Groups.Size

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

	case i.Groups != nil:
		return i.Groups.Modified

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

	case i.Groups != nil:
		return i.Groups.uniqueLocation(baseLoc)

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

	case i.Groups != nil:
		return i.Groups.updateFolder(f)

	default:
		return clues.New("unsupported type")
	}
}
