package details

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

// NewSharePointLocationIDer builds a LocationIDer for the drive and folder
// path. The path denoted by the folders should be unique within the drive.
func NewSharePointLocationIDer(
	category path.CategoryType,
	driveID string,
	escapedFolders ...string,
) uniqueLoc {
	pb := path.Builder{}.Append(category.String())
	prefixElems := 1

	if len(driveID) > 0 { // for library category
		pb = pb.Append(driveID)

		prefixElems = 2
	}

	pb = pb.Append(escapedFolders...)

	return uniqueLoc{
		pb:          pb,
		prefixElems: prefixElems,
	}
}

// SharePointInfo describes a sharepoint item
type SharePointInfo struct {
	Created      time.Time `json:"created,omitempty"`
	DriveName    string    `json:"driveName,omitempty"`
	DriveID      string    `json:"driveID,omitempty"`
	ItemName     string    `json:"itemName,omitempty"`
	ItemType     ItemType  `json:"itemType,omitempty"`
	ItemCount    int64     `json:"itemCount,omitempty"`
	ItemTemplate string    `json:"itemTemplate,omitempty"`
	Modified     time.Time `json:"modified,omitempty"`
	Owner        string    `json:"owner,omitempty"`
	ParentPath   string    `json:"parentPath,omitempty"`
	Size         int64     `json:"size,omitempty"`
	WebURL       string    `json:"webUrl,omitempty"`
	SiteID       string    `json:"siteID,omitempty"`
}

// Headers returns the human-readable names of properties in a SharePointInfo
// for printing out to a terminal in a columnar display.
func (i SharePointInfo) Headers() []string {
	switch i.ItemType {
	case SharePointLibrary:
		return []string{"ItemName", "Library", "ParentPath", "Size", "Owner", "Created", "Modified"}
	case SharePointList:
		return []string{"ListName", "ListItemsCount", "SiteURL", "Template", "Owner", "Created", "Modified"}
	}

	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i SharePointInfo) Values() []string {
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
	case SharePointList:
		return []string{
			i.ItemName,
			fmt.Sprintf("%d", i.ItemCount),
			getSiteURL(i.WebURL),
			i.ItemTemplate,
			i.Owner,
			dttm.FormatToTabularDisplay(i.Created),
			dttm.FormatToTabularDisplay(i.Modified),
		}
	}

	return []string{}
}

func (i *SharePointInfo) UpdateParentPath(newLocPath *path.Builder) {
	i.ParentPath = newLocPath.PopFront().String()
}

func (i *SharePointInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	loc := uniqueLoc{}

	switch i.ItemType {
	case SharePointLibrary, OneDriveItem:
		loc = NewSharePointLocationIDer(path.LibrariesCategory, i.DriveID, baseLoc.Elements()...)
	case SharePointList:
		loc = NewSharePointLocationIDer(path.ListsCategory, "", baseLoc.Elements()...)
	}

	return &loc, nil
}

func (i *SharePointInfo) updateFolder(f *FolderInfo) error {
	// TODO(ashmrtn): Change to just SharePointLibrary when the code that
	// generates the item type is fixed.
	switch i.ItemType {
	case OneDriveItem, SharePointLibrary:
		return updateFolderWithinDrive(SharePointLibrary, i.DriveName, i.DriveID, f)
	case SharePointList:
		return nil
	}

	return clues.New("unsupported non-SharePoint ItemType").With("item_type", i.ItemType)
}

func getSiteURL(itemWebURL string) string {
	siteURLWithSchemeAndHost := ""

	webURLParts, err := url.Parse(itemWebURL)
	if err != nil {
		return ""
	}

	pathParts := strings.Split(webURLParts.Path, "/")
	if len(pathParts) < 2 {
		return ""
	}

	siteURLWithSchemeAndHost, err = url.JoinPath(
		webURLParts.Scheme+":",
		webURLParts.Host,
		strings.Join(pathParts[:3], "/"),
	)
	if err != nil {
		return ""
	}

	return siteURLWithSchemeAndHost
}
