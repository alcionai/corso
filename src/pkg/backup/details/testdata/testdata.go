package testdata

import (
	"strings"
	"time"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

// mustParsePath takes a string representing a resource path and returns a path
// instance. Panics if the path cannot be parsed. Useful for simple variable
// assignments.
func mustParsePath(ref string, isItem bool) path.Path {
	p, err := path.FromDataLayerPath(ref, isItem)
	if err != nil {
		panic(err)
	}

	return p
}

// mustAppendPath takes a Path, string representing a path element, and whether
// the element is an item and returns a path instance representing the original
// path with the element appended to it. Panics if the path cannot be parsed.
// Useful for simple variable assignments.
func mustAppendPath(p path.Path, newElement string, isItem bool) path.Path {
	newP, err := p.Append(isItem, newElement)
	if err != nil {
		panic(err)
	}

	return newP
}

func locFromRepo(rr path.Path, isItem bool) *path.Builder {
	loc := &path.Builder{}

	for _, e := range rr.Folders() {
		loc = loc.Append(strings.TrimSuffix(e, folderSuffix))
	}

	if rr.Service() == path.OneDriveService || rr.Category() == path.LibrariesCategory {
		loc = loc.PopFront()
	}

	// Folders don't have their final element in the location.
	if !isItem {
		loc = loc.Dir()
	}

	return loc
}

type repoRefAndLocRef struct {
	RR  path.Path
	Loc *path.Builder
}

func (p repoRefAndLocRef) MustAppend(newElement string, isItem bool) repoRefAndLocRef {
	e := newElement + folderSuffix

	if isItem {
		e = newElement + fileSuffix
	}

	res := repoRefAndLocRef{
		RR: mustAppendPath(p.RR, e, isItem),
	}

	res.Loc = locFromRepo(res.RR, isItem)

	return res
}

func (p repoRefAndLocRef) ItemLocation() string {
	return strings.TrimSuffix(p.RR.Item(), fileSuffix)
}

func (p repoRefAndLocRef) FolderLocation() string {
	lastElem := p.RR.ToBuilder().LastElem()

	if len(p.RR.Item()) > 0 {
		f := p.RR.Folders()
		lastElem = f[len(f)-2]
	}

	return p.Loc.Append(strings.TrimSuffix(lastElem, folderSuffix)).String()
}

func mustPathRep(ref string, isItem bool) repoRefAndLocRef {
	res := repoRefAndLocRef{}
	tmp := mustParsePath(ref, isItem)

	// Now append stuff to the RepoRef elements so we have distinct LocationRef
	// and RepoRef elements to simulate using IDs in the path instead of display
	// names.
	rrPB := &path.Builder{}
	for _, e := range tmp.Folders() {
		rrPB = rrPB.Append(e + folderSuffix)
	}

	if isItem {
		rrPB = rrPB.Append(tmp.Item() + fileSuffix)
	}

	rr, err := rrPB.ToDataLayerPath(
		tmp.Tenant(),
		tmp.ResourceOwner(),
		tmp.Service(),
		tmp.Category(),
		isItem)
	if err != nil {
		panic(err)
	}

	res.RR = rr
	res.Loc = locFromRepo(rr, isItem)

	return res
}

const (
	folderSuffix = ".d"
	fileSuffix   = ".f"

	ItemName1  = "item1"
	ItemName2  = "item2"
	ItemName3  = "item3"
	UserEmail1 = "user1@email.com"
	UserEmail2 = "user2@email.com"
)

var (
	Time1 = time.Date(2022, 9, 21, 10, 0, 0, 0, time.UTC)
	Time2 = time.Date(2022, 10, 21, 10, 0, 0, 0, time.UTC)
	Time3 = time.Date(2023, 9, 21, 10, 0, 0, 0, time.UTC)
	Time4 = time.Date(2023, 10, 21, 10, 0, 0, 0, time.UTC)

	ExchangeEmailInboxPath = mustPathRep("tenant-id/exchange/user-id/email/Inbox", false)
	ExchangeEmailBasePath  = ExchangeEmailInboxPath.MustAppend("subfolder", false)
	ExchangeEmailBasePath2 = ExchangeEmailInboxPath.MustAppend("othersubfolder/", false)
	ExchangeEmailBasePath3 = ExchangeEmailBasePath2.MustAppend("subsubfolder", false)
	ExchangeEmailItemPath1 = ExchangeEmailBasePath.MustAppend(ItemName1, true)
	ExchangeEmailItemPath2 = ExchangeEmailBasePath2.MustAppend(ItemName2, true)
	ExchangeEmailItemPath3 = ExchangeEmailBasePath3.MustAppend(ItemName3, true)

	ExchangeEmailItems = []details.Entry{
		{
			RepoRef:     ExchangeEmailItemPath1.RR.String(),
			ShortRef:    ExchangeEmailItemPath1.RR.ShortRef(),
			ParentRef:   ExchangeEmailItemPath1.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     ExchangeEmailItemPath1.ItemLocation(),
			LocationRef: ExchangeEmailItemPath1.Loc.String(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType: details.ExchangeMail,
					Sender:   "a-person",
					Subject:  "foo",
					Received: Time1,
				},
			},
		},
		{
			RepoRef:     ExchangeEmailItemPath2.RR.String(),
			ShortRef:    ExchangeEmailItemPath2.RR.ShortRef(),
			ParentRef:   ExchangeEmailItemPath2.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     ExchangeEmailItemPath2.ItemLocation(),
			LocationRef: ExchangeEmailItemPath2.Loc.String(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType: details.ExchangeMail,
					Sender:   "a-person",
					Subject:  "bar",
					Received: Time2,
				},
			},
		},
		{
			RepoRef:     ExchangeEmailItemPath3.RR.String(),
			ShortRef:    ExchangeEmailItemPath3.RR.ShortRef(),
			ParentRef:   ExchangeEmailItemPath3.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     ExchangeEmailItemPath3.ItemLocation(),
			LocationRef: ExchangeEmailItemPath3.Loc.String(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType: details.ExchangeMail,
					Sender:   "another-person",
					Subject:  "baz",
					Received: Time2,
				},
			},
		},
	}

	ExchangeContactsRootPath  = mustPathRep("tenant-id/exchange/user-id/contacts/contacts", false)
	ExchangeContactsBasePath  = ExchangeContactsRootPath.MustAppend("contacts", false)
	ExchangeContactsBasePath2 = ExchangeContactsRootPath.MustAppend("morecontacts", false)
	ExchangeContactsItemPath1 = ExchangeContactsBasePath.MustAppend(ItemName1, true)
	ExchangeContactsItemPath2 = ExchangeContactsBasePath2.MustAppend(ItemName2, true)

	ExchangeContactsItems = []details.Entry{
		{
			RepoRef:     ExchangeContactsItemPath1.RR.String(),
			ShortRef:    ExchangeContactsItemPath1.RR.ShortRef(),
			ParentRef:   ExchangeContactsItemPath1.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     ExchangeContactsItemPath1.ItemLocation(),
			LocationRef: ExchangeContactsItemPath1.Loc.String(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType:    details.ExchangeContact,
					ContactName: "a-person",
				},
			},
		},
		{
			RepoRef:     ExchangeContactsItemPath2.RR.String(),
			ShortRef:    ExchangeContactsItemPath2.RR.ShortRef(),
			ParentRef:   ExchangeContactsItemPath2.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     ExchangeContactsItemPath2.ItemLocation(),
			LocationRef: ExchangeContactsItemPath2.Loc.String(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType:    details.ExchangeContact,
					ContactName: "another-person",
				},
			},
		},
	}

	ExchangeEventsBasePath  = mustPathRep("tenant-id/exchange/user-id/events/holidays", false)
	ExchangeEventsBasePath2 = mustPathRep("tenant-id/exchange/user-id/events/moreholidays", false)
	ExchangeEventsItemPath1 = ExchangeEventsBasePath.MustAppend(ItemName1, true)
	ExchangeEventsItemPath2 = ExchangeEventsBasePath2.MustAppend(ItemName2, true)

	ExchangeEventsItems = []details.Entry{
		{
			RepoRef:     ExchangeEventsItemPath1.RR.String(),
			ShortRef:    ExchangeEventsItemPath1.RR.ShortRef(),
			ParentRef:   ExchangeEventsItemPath1.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     ExchangeEventsItemPath1.ItemLocation(),
			LocationRef: ExchangeEventsItemPath1.Loc.String(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType:    details.ExchangeEvent,
					Organizer:   "a-person",
					Subject:     "foo",
					EventStart:  Time1,
					EventRecurs: false,
				},
			},
		},
		{
			RepoRef:     ExchangeEventsItemPath2.RR.String(),
			ShortRef:    ExchangeEventsItemPath2.RR.ShortRef(),
			ParentRef:   ExchangeEventsItemPath2.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     ExchangeEventsItemPath2.ItemLocation(),
			LocationRef: ExchangeEventsItemPath2.Loc.String(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType:    details.ExchangeEvent,
					Organizer:   "a-person",
					Subject:     "foo",
					EventStart:  Time2,
					EventRecurs: true,
				},
			},
		},
	}

	OneDriveRootPath   = mustPathRep("tenant-id/onedrive/user-id/files/drives/foo/root:", false)
	OneDriveFolderPath = OneDriveRootPath.MustAppend("folder", false)
	OneDriveBasePath1  = OneDriveFolderPath.MustAppend("a", false)
	OneDriveBasePath2  = OneDriveFolderPath.MustAppend("b", false)

	OneDriveItemPath1 = OneDriveFolderPath.MustAppend(ItemName1, true)
	OneDriveItemPath2 = OneDriveBasePath1.MustAppend(ItemName2, true)
	OneDriveItemPath3 = OneDriveBasePath2.MustAppend(ItemName3, true)

	OneDriveFolderFolder  = OneDriveFolderPath.Loc.PopFront().String()
	OneDriveParentFolder1 = OneDriveBasePath1.Loc.PopFront().String()
	OneDriveParentFolder2 = OneDriveBasePath2.Loc.PopFront().String()

	OneDriveItems = []details.Entry{
		{
			RepoRef:     OneDriveItemPath1.RR.String(),
			ShortRef:    OneDriveItemPath1.RR.ShortRef(),
			ParentRef:   OneDriveItemPath1.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     OneDriveItemPath1.ItemLocation(),
			LocationRef: OneDriveItemPath1.Loc.String(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:   details.OneDriveItem,
					ParentPath: OneDriveFolderFolder,
					ItemName:   OneDriveItemPath1.ItemLocation() + "name",
					Size:       int64(23),
					Owner:      UserEmail1,
					Created:    Time2,
					Modified:   Time4,
				},
			},
		},
		{
			RepoRef:     OneDriveItemPath2.RR.String(),
			ShortRef:    OneDriveItemPath2.RR.ShortRef(),
			ParentRef:   OneDriveItemPath2.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     OneDriveItemPath2.ItemLocation(),
			LocationRef: OneDriveItemPath2.Loc.String(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:   details.OneDriveItem,
					ParentPath: OneDriveParentFolder1,
					ItemName:   OneDriveItemPath2.ItemLocation() + "name",
					Size:       int64(42),
					Owner:      UserEmail1,
					Created:    Time1,
					Modified:   Time3,
				},
			},
		},
		{
			RepoRef:     OneDriveItemPath3.RR.String(),
			ShortRef:    OneDriveItemPath3.RR.ShortRef(),
			ParentRef:   OneDriveItemPath3.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     OneDriveItemPath3.ItemLocation(),
			LocationRef: OneDriveItemPath3.Loc.String(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:   details.OneDriveItem,
					ParentPath: OneDriveParentFolder2,
					ItemName:   OneDriveItemPath3.ItemLocation() + "name",
					Size:       int64(19),
					Owner:      UserEmail2,
					Created:    Time2,
					Modified:   Time4,
				},
			},
		},
	}

	SharePointRootPath    = mustPathRep("tenant-id/sharepoint/site-id/libraries/drives/foo/root:", false)
	SharePointLibraryPath = SharePointRootPath.MustAppend("library", false)
	SharePointBasePath1   = SharePointLibraryPath.MustAppend("a", false)
	SharePointBasePath2   = SharePointLibraryPath.MustAppend("b", false)

	SharePointLibraryItemPath1 = SharePointLibraryPath.MustAppend(ItemName1, true)
	SharePointLibraryItemPath2 = SharePointBasePath1.MustAppend(ItemName2, true)
	SharePointLibraryItemPath3 = SharePointBasePath2.MustAppend(ItemName3, true)

	SharePointLibraryFolder  = SharePointLibraryPath.Loc.PopFront().String()
	SharePointParentLibrary1 = SharePointBasePath1.Loc.PopFront().String()
	SharePointParentLibrary2 = SharePointBasePath2.Loc.PopFront().String()

	SharePointLibraryItems = []details.Entry{
		{
			RepoRef:     SharePointLibraryItemPath1.RR.String(),
			ShortRef:    SharePointLibraryItemPath1.RR.ShortRef(),
			ParentRef:   SharePointLibraryItemPath1.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     SharePointLibraryItemPath1.ItemLocation(),
			LocationRef: SharePointLibraryItemPath1.Loc.String(),
			ItemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:   details.SharePointLibrary,
					ParentPath: SharePointLibraryFolder,
					ItemName:   SharePointLibraryItemPath1.ItemLocation() + "name",
					Size:       int64(23),
					Owner:      UserEmail1,
					Created:    Time2,
					Modified:   Time4,
				},
			},
		},
		{
			RepoRef:     SharePointLibraryItemPath2.RR.String(),
			ShortRef:    SharePointLibraryItemPath2.RR.ShortRef(),
			ParentRef:   SharePointLibraryItemPath2.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     SharePointLibraryItemPath2.ItemLocation(),
			LocationRef: SharePointLibraryItemPath2.Loc.String(),
			ItemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:   details.SharePointLibrary,
					ParentPath: SharePointParentLibrary1,
					ItemName:   SharePointLibraryItemPath2.ItemLocation() + "name",
					Size:       int64(42),
					Owner:      UserEmail1,
					Created:    Time1,
					Modified:   Time3,
				},
			},
		},
		{
			RepoRef:     SharePointLibraryItemPath3.RR.String(),
			ShortRef:    SharePointLibraryItemPath3.RR.ShortRef(),
			ParentRef:   SharePointLibraryItemPath3.RR.ToBuilder().Dir().ShortRef(),
			ItemRef:     SharePointLibraryItemPath3.ItemLocation(),
			LocationRef: SharePointLibraryItemPath3.Loc.String(),
			ItemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:   details.SharePointLibrary,
					ParentPath: SharePointParentLibrary2,
					ItemName:   SharePointLibraryItemPath3.ItemLocation() + "name",
					Size:       int64(19),
					Owner:      UserEmail2,
					Created:    Time2,
					Modified:   Time4,
				},
			},
		},
	}
)

func GetDetailsSet() *details.Details {
	entries := []details.Entry{}

	for _, e := range ExchangeEmailItems {
		entries = append(entries, e)
	}

	for _, e := range ExchangeContactsItems {
		entries = append(entries, e)
	}

	for _, e := range ExchangeEventsItems {
		entries = append(entries, e)
	}

	for _, e := range OneDriveItems {
		entries = append(entries, e)
	}

	for _, e := range SharePointLibraryItems {
		entries = append(entries, e)
	}

	return &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: entries,
		},
	}
}
