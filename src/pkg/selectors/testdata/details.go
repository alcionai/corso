package testdata

import (
	stdpath "path"
	"time"

	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/backup/details"
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
	ItemName1 = "item1"
	ItemName2 = "item2"
)

var (
	Time1 = time.Date(2022, 9, 21, 10, 0, 0, 0, time.UTC)
	Time2 = time.Date(2022, 10, 21, 10, 0, 0, 0, time.UTC)

	ExchangeEmailBasePath  = mustParsePath("tenant-id/exchange/user-id/email/Inbox/subfolder", false)
	ExchangeEmailItemPath1 = mustAppendPath(ExchangeEmailBasePath, ItemName1, true)
	ExchangeEmailItemPath2 = mustAppendPath(ExchangeEmailBasePath, ItemName2, true)

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
	}

	ExchangeContactsBasePath  = mustParsePath("tenant-id/exchange/user-id/contacts/contacts", false)
	ExchangeContactsItemPath1 = mustAppendPath(ExchangeContactsBasePath, ItemName1, true)
	ExchangeContactsItemPath2 = mustAppendPath(ExchangeContactsBasePath, ItemName2, true)

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

	ExchangeEventsBasePath  = mustParsePath("tenant-id/exchange/user-id/events/holidays", false)
	ExchangeEventsItemPath1 = mustAppendPath(ExchangeEventsBasePath, ItemName1, true)
	ExchangeEventsItemPath2 = mustAppendPath(ExchangeEventsBasePath, ItemName2, true)

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

	OneDriveBasePath  = mustParsePath("tenant-id/onedrive/user-id/files/folder/subfolder", false)
	OneDriveItemPath1 = mustAppendPath(OneDriveBasePath, ItemName1, true)
	OneDriveItemPath2 = mustAppendPath(OneDriveBasePath, ItemName2, true)

	OneDriveItems = []details.DetailsEntry{
		{
			RepoRef:   OneDriveItemPath1.String(),
			ShortRef:  OneDriveItemPath1.ShortRef(),
			ParentRef: OneDriveItemPath1.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType: details.OneDriveItem,
					ParentPath: stdpath.Join(
						append(
							[]string{
								"drives",
								"foo",
								"root:",
							},
							OneDriveItemPath1.Folders()...,
						)...,
					),
					ItemName: OneDriveItemPath1.Item() + "name",
				},
			},
		},
		{
			RepoRef:   OneDriveItemPath2.String(),
			ShortRef:  OneDriveItemPath2.ShortRef(),
			ParentRef: OneDriveItemPath2.ToBuilder().Dir().ShortRef(),
			ItemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType: details.OneDriveItem,
					ParentPath: stdpath.Join(
						append(
							[]string{
								"drives",
								"foo",
								"root:",
							},
							OneDriveItemPath2.Folders()...,
						)...,
					),
					ItemName: OneDriveItemPath2.Item() + "name",
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

	return &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: entries,
		},
	}
}
