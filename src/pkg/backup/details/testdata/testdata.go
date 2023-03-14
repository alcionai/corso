package testdata

import (
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
	newP, err := p.Append(newElement, isItem)
	if err != nil {
		panic(err)
	}

	return newP
}

const (
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

	ExchangeEmailInboxPath = mustParsePath("tenant-id/exchange/user-id/email/Inbox", false)
	ExchangeEmailBasePath  = mustAppendPath(ExchangeEmailInboxPath, "subfolder", false)
	ExchangeEmailBasePath2 = mustAppendPath(ExchangeEmailInboxPath, "othersubfolder/", false)
	ExchangeEmailBasePath3 = mustAppendPath(ExchangeEmailBasePath2, "subsubfolder", false)
	ExchangeEmailItemPath1 = mustAppendPath(ExchangeEmailBasePath, ItemName1, true)
	ExchangeEmailItemPath2 = mustAppendPath(ExchangeEmailBasePath2, ItemName2, true)
	ExchangeEmailItemPath3 = mustAppendPath(ExchangeEmailBasePath3, ItemName3, true)

	ExchangeEmailItems = []details.DetailsEntry{
		{
			RepoRef:   ExchangeEmailItemPath1.String(),
			ShortRef:  ExchangeEmailItemPath1.ShortRef(),
			ParentRef: ExchangeEmailItemPath1.ToBuilder().Dir().ShortRef(),
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
			RepoRef:   ExchangeEmailItemPath2.String(),
			ShortRef:  ExchangeEmailItemPath2.ShortRef(),
			ParentRef: ExchangeEmailItemPath2.ToBuilder().Dir().ShortRef(),
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
			RepoRef:   ExchangeEmailItemPath3.String(),
			ShortRef:  ExchangeEmailItemPath3.ShortRef(),
			ParentRef: ExchangeEmailItemPath3.ToBuilder().Dir().ShortRef(),
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

	ExchangeContactsRootPath  = mustParsePath("tenant-id/exchange/user-id/contacts/contacts", false)
	ExchangeContactsBasePath  = mustAppendPath(ExchangeContactsRootPath, "contacts", false)
	ExchangeContactsBasePath2 = mustAppendPath(ExchangeContactsRootPath, "morecontacts", false)
	ExchangeContactsItemPath1 = mustAppendPath(ExchangeContactsBasePath, ItemName1, true)
	ExchangeContactsItemPath2 = mustAppendPath(ExchangeContactsBasePath2, ItemName2, true)

	ExchangeContactsItems = []details.DetailsEntry{
		{
			RepoRef:   ExchangeContactsItemPath1.String(),
			ShortRef:  ExchangeContactsItemPath1.ShortRef(),
			ParentRef: ExchangeContactsItemPath1.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType:    details.ExchangeContact,
					ContactName: "a-person",
				},
			},
		},
		{
			RepoRef:   ExchangeContactsItemPath2.String(),
			ShortRef:  ExchangeContactsItemPath2.ShortRef(),
			ParentRef: ExchangeContactsItemPath2.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					ItemType:    details.ExchangeContact,
					ContactName: "another-person",
				},
			},
		},
	}

	ExchangeEventsRootPath  = mustParsePath("tenant-id/exchange/user-id/events/holidays", false)
	ExchangeEventsBasePath  = mustAppendPath(ExchangeEventsRootPath, "holidays", false)
	ExchangeEventsBasePath2 = mustAppendPath(ExchangeEventsRootPath, "moreholidays", false)
	ExchangeEventsItemPath1 = mustAppendPath(ExchangeEventsBasePath, ItemName1, true)
	ExchangeEventsItemPath2 = mustAppendPath(ExchangeEventsBasePath2, ItemName2, true)

	ExchangeEventsItems = []details.DetailsEntry{
		{
			RepoRef:   ExchangeEventsItemPath1.String(),
			ShortRef:  ExchangeEventsItemPath1.ShortRef(),
			ParentRef: ExchangeEventsItemPath1.ToBuilder().Dir().ShortRef(),
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
			RepoRef:   ExchangeEventsItemPath2.String(),
			ShortRef:  ExchangeEventsItemPath2.ShortRef(),
			ParentRef: ExchangeEventsItemPath2.ToBuilder().Dir().ShortRef(),
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

	OneDriveRootPath   = mustParsePath("tenant-id/onedrive/user-id/files/drives/foo/root:", false)
	OneDriveFolderPath = mustAppendPath(OneDriveRootPath, "folder", false)
	OneDriveBasePath1  = mustAppendPath(OneDriveFolderPath, "a", false)
	OneDriveBasePath2  = mustAppendPath(OneDriveFolderPath, "b", false)

	OneDriveItemPath1 = mustAppendPath(OneDriveFolderPath, ItemName1, true)
	OneDriveItemPath2 = mustAppendPath(OneDriveBasePath1, ItemName2, true)
	OneDriveItemPath3 = mustAppendPath(OneDriveBasePath2, ItemName3, true)

	OneDriveFolderFolder  = stdpath.Join(OneDriveFolderPath.Folders()[3:]...)
	OneDriveParentFolder1 = stdpath.Join(OneDriveBasePath1.Folders()[3:]...)
	OneDriveParentFolder2 = stdpath.Join(OneDriveBasePath2.Folders()[3:]...)

	OneDriveItems = []details.DetailsEntry{
		{
			RepoRef:   OneDriveItemPath1.String(),
			ShortRef:  OneDriveItemPath1.ShortRef(),
			ParentRef: OneDriveItemPath1.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:   details.OneDriveItem,
					ParentPath: OneDriveFolderFolder,
					ItemName:   OneDriveItemPath1.Item() + "name",
					Size:       int64(23),
					Owner:      UserEmail1,
					Created:    Time2,
					Modified:   Time4,
				},
			},
		},
		{
			RepoRef:   OneDriveItemPath2.String(),
			ShortRef:  OneDriveItemPath2.ShortRef(),
			ParentRef: OneDriveItemPath2.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:   details.OneDriveItem,
					ParentPath: OneDriveParentFolder1,
					ItemName:   OneDriveItemPath2.Item() + "name",
					Size:       int64(42),
					Owner:      UserEmail1,
					Created:    Time1,
					Modified:   Time3,
				},
			},
		},
		{
			RepoRef:   OneDriveItemPath3.String(),
			ShortRef:  OneDriveItemPath3.ShortRef(),
			ParentRef: OneDriveItemPath3.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:   details.OneDriveItem,
					ParentPath: OneDriveParentFolder2,
					ItemName:   OneDriveItemPath3.Item() + "name",
					Size:       int64(19),
					Owner:      UserEmail2,
					Created:    Time2,
					Modified:   Time4,
				},
			},
		},
	}

	SharePointRootPath    = mustParsePath("tenant-id/sharepoint/site-id/libraries/drives/foo/root:", false)
	SharePointLibraryPath = mustAppendPath(SharePointRootPath, "library", false)
	SharePointBasePath1   = mustAppendPath(SharePointLibraryPath, "a", false)
	SharePointBasePath2   = mustAppendPath(SharePointLibraryPath, "b", false)

	SharePointLibraryItemPath1 = mustAppendPath(SharePointLibraryPath, ItemName1, true)
	SharePointLibraryItemPath2 = mustAppendPath(SharePointBasePath1, ItemName2, true)
	SharePointLibraryItemPath3 = mustAppendPath(SharePointBasePath2, ItemName3, true)

	SharePointLibraryFolder  = stdpath.Join(SharePointLibraryPath.Folders()[3:]...)
	SharePointParentLibrary1 = stdpath.Join(SharePointBasePath1.Folders()[3:]...)
	SharePointParentLibrary2 = stdpath.Join(SharePointBasePath2.Folders()[3:]...)

	SharePointLibraryItems = []details.DetailsEntry{
		{
			RepoRef:   SharePointLibraryItemPath1.String(),
			ShortRef:  SharePointLibraryItemPath1.ShortRef(),
			ParentRef: SharePointLibraryItemPath1.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:   details.SharePointLibrary,
					ParentPath: SharePointLibraryFolder,
					ItemName:   SharePointLibraryItemPath1.Item() + "name",
					Size:       int64(23),
					Owner:      UserEmail1,
					Created:    Time2,
					Modified:   Time4,
				},
			},
		},
		{
			RepoRef:   SharePointLibraryItemPath2.String(),
			ShortRef:  SharePointLibraryItemPath2.ShortRef(),
			ParentRef: SharePointLibraryItemPath2.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:   details.SharePointLibrary,
					ParentPath: SharePointParentLibrary1,
					ItemName:   SharePointLibraryItemPath2.Item() + "name",
					Size:       int64(42),
					Owner:      UserEmail1,
					Created:    Time1,
					Modified:   Time3,
				},
			},
		},
		{
			RepoRef:   SharePointLibraryItemPath3.String(),
			ShortRef:  SharePointLibraryItemPath3.ShortRef(),
			ParentRef: SharePointLibraryItemPath3.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:   details.SharePointLibrary,
					ParentPath: SharePointParentLibrary2,
					ItemName:   SharePointLibraryItemPath3.Item() + "name",
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
	entries := []details.DetailsEntry{}

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
