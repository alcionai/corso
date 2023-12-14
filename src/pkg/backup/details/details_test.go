package details

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
)

// ------------------------------------------------------------
// unit tests
// ------------------------------------------------------------

type DetailsUnitSuite struct {
	tester.Suite
}

func TestDetailsUnitSuite(t *testing.T) {
	suite.Run(t, &DetailsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DetailsUnitSuite) TestDetailsEntry_HeadersValues() {
	initial := time.Now()
	nowStr := dttm.FormatTo(initial, dttm.TabularOutput)
	now, err := dttm.ParseTime(nowStr)
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name     string
		entry    Entry
		expectHs []string
		expectVs []string
	}{
		{
			name: "no info",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
			},
			expectHs: []string{"ID"},
			expectVs: []string{"deadbeef"},
		},
		{
			name: "exchange event info",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					Exchange: &ExchangeInfo{
						ItemType:    ExchangeEvent,
						EventStart:  now,
						EventEnd:    now,
						Organizer:   "organizer",
						EventRecurs: true,
						Subject:     "subject",
					},
				},
			},
			expectHs: []string{"ID", "Organizer", "Subject", "Starts", "Ends", "Recurring"},
			expectVs: []string{"deadbeef", "organizer", "subject", nowStr, nowStr, "true"},
		},
		{
			name: "exchange contact info",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					Exchange: &ExchangeInfo{
						ItemType:    ExchangeContact,
						ContactName: "contactName",
					},
				},
			},
			expectHs: []string{"ID", "Contact Name"},
			expectVs: []string{"deadbeef", "contactName"},
		},
		{
			name: "exchange mail info",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					Exchange: &ExchangeInfo{
						ItemType:   ExchangeMail,
						Sender:     "sender",
						ParentPath: "Parent",
						Recipient:  []string{"receiver"},
						Subject:    "subject",
						Received:   now,
					},
				},
			},
			expectHs: []string{"ID", "Sender", "Folder", "Subject", "Received"},
			expectVs: []string{"deadbeef", "sender", "Parent", "subject", nowStr},
		},
		{
			name: "sharepoint library info",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					SharePoint: &SharePointInfo{
						ItemName:   "itemName",
						ItemType:   SharePointLibrary,
						ParentPath: "parentPath",
						Size:       1000,
						WebURL:     "https://not.a.real/url",
						DriveName:  "aLibrary",
						Owner:      "user@email.com",
						Created:    now,
						Modified:   now,
					},
				},
			},
			expectHs: []string{"ID", "ItemName", "Library", "ParentPath", "Size", "Owner", "Created", "Modified"},
			expectVs: []string{
				"deadbeef",
				"itemName",
				"aLibrary",
				"parentPath",
				"1.0 kB",
				"user@email.com",
				nowStr,
				nowStr,
			},
		},
		{
			name: "sharepoint list info for genericList template",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					SharePoint: &SharePointInfo{
						ItemType:     SharePointList,
						ItemName:     "itemName",
						ItemCount:    50,
						ItemTemplate: "genericList",
						WebURL:       "https://10rqc2.sharepoint.com/sites/site-4754-small-lists/Lists/itemName",
						Owner:        "user@email.com",
						Created:      now,
						Modified:     now,
					},
				},
			},
			expectHs: []string{"ID", "ListName", "ListItemsCount", "SiteURL", "Template", "Owner", "Created", "Modified"},
			expectVs: []string{
				"deadbeef",
				"itemName",
				"50",
				"https://10rqc2.sharepoint.com/sites/site-4754-small-lists",
				"genericList",
				"user@email.com",
				nowStr,
				nowStr,
			},
		},
		{
			name: "sharepoint list info for documentLibrary template",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					SharePoint: &SharePointInfo{
						ItemType:     SharePointList,
						ItemName:     "Shared%20Documents",
						ItemCount:    50,
						ItemTemplate: "documentLibrary",
						WebURL:       "https://10rqc2.sharepoint.com/sites/site-4754-small-lists/Shared%20Documents",
						Owner:        "user@email.com",
						Created:      now,
						Modified:     now,
					},
				},
			},
			expectHs: []string{"ID", "ListName", "ListItemsCount", "SiteURL", "Template", "Owner", "Created", "Modified"},
			expectVs: []string{
				"deadbeef",
				"Shared%20Documents",
				"50",
				"https://10rqc2.sharepoint.com/sites/site-4754-small-lists",
				"documentLibrary",
				"user@email.com",
				nowStr,
				nowStr,
			},
		},
		{
			name: "oneDrive info",
			entry: Entry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{
						ItemName:   "itemName",
						ParentPath: "parentPath",
						Size:       1000,
						Owner:      "user@email.com",
						Created:    now,
						Modified:   now,
					},
				},
			},
			expectHs: []string{"ID", "ItemName", "ParentPath", "Size", "Owner", "Created", "Modified"},
			expectVs: []string{"deadbeef", "itemName", "parentPath", "1.0 kB", "user@email.com", nowStr, nowStr},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			hs := test.entry.Headers(false)
			assert.Equal(t, test.expectHs, hs)
			vs := test.entry.Values(false)
			assert.Equal(t, test.expectVs, vs)

			hs = test.entry.Headers(true)
			assert.Equal(t, test.expectHs[1:], hs)
			vs = test.entry.Values(true)
			assert.Equal(t, test.expectVs[1:], vs)
		})
	}
}

func exchangeEntry(t *testing.T, id string, size int, it ItemType) Entry {
	rr := makeItemPath(
		t,
		path.ExchangeService,
		path.EmailCategory,
		"tenant-id",
		"user-id",
		[]string{"Inbox", "folder1", id})

	return Entry{
		RepoRef:     rr.String(),
		ShortRef:    rr.ShortRef(),
		ParentRef:   rr.ToBuilder().Dir().ShortRef(),
		LocationRef: rr.Folder(true),
		ItemRef:     rr.Item(),
		ItemInfo: ItemInfo{
			Exchange: &ExchangeInfo{
				ItemType: it,
				Modified: time.Now(),
				Size:     int64(size),
			},
		},
	}
}

func oneDriveishEntry(t *testing.T, id string, size int, it ItemType, service path.ServiceType) Entry {
	var (
		category path.CategoryType
		info     ItemInfo
	)

	switch it {
	case OneDriveItem:
		category = path.FilesCategory
		info = ItemInfo{
			OneDrive: &OneDriveInfo{
				ItemName:  "bar",
				DriveID:   "drive-id",
				DriveName: "drive-name",
				Modified:  time.Now(),
				ItemType:  it,
				Size:      int64(size),
			},
		}
	case SharePointLibrary:
		category = path.LibrariesCategory

		switch service {
		case path.SharePointService:
			info = ItemInfo{
				SharePoint: &SharePointInfo{
					ItemName:  "bar",
					DriveID:   "drive-id",
					DriveName: "drive-name",
					Modified:  time.Now(),
					ItemType:  it,
					Size:      int64(size),
				},
			}
		case path.GroupsService:
			info = ItemInfo{
				Groups: &GroupsInfo{
					ItemName:  "bar",
					DriveID:   "drive-id",
					DriveName: "drive-name",
					Modified:  time.Now(),
					ItemType:  it,
					Size:      int64(size),
				},
			}
		}
	}

	rr := makeItemPath(
		t,
		service,
		category,
		"tenant-id",
		"user-id",
		[]string{
			odConsts.DrivesPathDir,
			"drive-id",
			odConsts.RootPathDir,
			"Inbox",
			"folder1",
			id,
		})

	loc := path.Builder{}.Append(rr.Folders()...).PopFront().PopFront()

	return Entry{
		RepoRef:     rr.String(),
		ShortRef:    rr.ShortRef(),
		ParentRef:   rr.ToBuilder().Dir().ShortRef(),
		LocationRef: loc.String(),
		ItemRef:     rr.Item(),
		ItemInfo:    info,
	}
}

func (suite *DetailsUnitSuite) TestDetailsAdd_NoLocationFolders() {
	itemID := "foo"

	t := suite.T()
	table := []struct {
		name  string
		entry Entry
	}{
		{
			name:  "Exchange Email",
			entry: exchangeEntry(t, itemID, 42, ExchangeMail),
		},
		{
			name:  "OneDrive File",
			entry: oneDriveishEntry(t, itemID, 42, OneDriveItem, path.OneDriveService),
		},
		{
			name:  "SharePoint File",
			entry: oneDriveishEntry(t, itemID, 42, SharePointLibrary, path.SharePointService),
		},
		{
			name: "Legacy SharePoint File",
			entry: func() Entry {
				res := oneDriveishEntry(t, itemID, 42, SharePointLibrary, path.SharePointService)
				res.SharePoint.ItemType = OneDriveItem

				return res
			}(),
		},
		{
			name:  "Group SharePoint File",
			entry: oneDriveishEntry(t, itemID, 42, SharePointLibrary, path.GroupsService),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			rr, err := path.FromDataLayerPath(test.entry.RepoRef, true)
			require.NoError(t, err, clues.ToCore(err))

			db := &Builder{}

			// Make a local copy so we can modify it.
			localItem := test.entry

			err = db.Add(rr, &path.Builder{}, localItem.ItemInfo)
			require.NoError(t, err, clues.ToCore(err))

			// Clear LocationRef that's automatically populated since we passed an
			// empty builder above.
			localItem.LocationRef = ""

			expectedShortRef := localItem.ShortRef
			localItem.ShortRef = ""

			deets := db.Details()
			assert.Len(t, deets.Entries, 1)

			got := deets.Entries[0]
			gotShortRef := got.ShortRef
			got.ShortRef = ""

			assert.Equal(t, localItem, got, "DetailsEntry")
			assert.Equal(t, expectedShortRef, gotShortRef, "ShortRef")
		})
	}
}

func (suite *DetailsUnitSuite) TestDetailsAdd_LocationFolders() {
	t := suite.T()

	exchangeMail1 := exchangeEntry(t, "foo1", 42, ExchangeMail)
	oneDrive1 := oneDriveishEntry(t, "foo1", 42, OneDriveItem, path.OneDriveService)
	sharePoint1 := oneDriveishEntry(t, "foo1", 42, SharePointLibrary, path.SharePointService)
	sharePointLegacy1 := oneDriveishEntry(t, "foo1", 42, SharePointLibrary, path.SharePointService)
	sharePointLegacy1.SharePoint.ItemType = OneDriveItem
	group1 := oneDriveishEntry(t, "foo1", 42, SharePointLibrary, path.GroupsService)
	// Sleep for a little so we get a larger difference in mod times between the
	// earlier and later entries.
	time.Sleep(100 * time.Millisecond)

	// Get fresh item IDs so we can check that folders populate with the latest
	// mod time. Also the details API is built with the idea that duplicate items
	// aren't added (it has no checking for that).
	exchangeMail2 := exchangeEntry(t, "foo2", 43, ExchangeMail)
	exchangeContact1 := exchangeEntry(t, "foo3", 44, ExchangeContact)

	exchangeFolders := []Entry{
		{
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: "Inbox",
					ItemType:    FolderItem,
					DataType:    ExchangeMail,
				},
			},
		},
		{
			LocationRef: "Inbox",
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: "folder1",
					ItemType:    FolderItem,
					DataType:    ExchangeMail,
				},
			},
		},
	}

	exchangeContactFolders := []Entry{
		{
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: "Inbox",
					ItemType:    FolderItem,
					DataType:    ExchangeContact,
				},
			},
		},
		{
			LocationRef: "Inbox",
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: "folder1",
					ItemType:    FolderItem,
					DataType:    ExchangeContact,
				},
			},
		},
	}

	oneDriveishFolders := []Entry{
		{
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: odConsts.RootPathDir,
					ItemType:    FolderItem,
					DriveName:   "drive-name",
					DriveID:     "drive-id",
				},
			},
		},
		{
			LocationRef: odConsts.RootPathDir,
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: "Inbox",
					ItemType:    FolderItem,
					DriveName:   "drive-name",
					DriveID:     "drive-id",
				},
			},
		},
		{
			LocationRef: "root:/Inbox",
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: "folder1",
					ItemType:    FolderItem,
					DriveName:   "drive-name",
					DriveID:     "drive-id",
				},
			},
		},
	}

	table := []struct {
		name         string
		entries      func() []Entry
		expectedDirs func() []Entry
	}{
		{
			name: "One Exchange Email",
			entries: func() []Entry {
				e := exchangeMail1
				ei := *exchangeMail1.Exchange
				e.Exchange = &ei

				return []Entry{e}
			},
			expectedDirs: func() []Entry {
				res := []Entry{}

				for _, entry := range exchangeFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.Size = exchangeMail1.Exchange.Size
					e.Folder.Modified = exchangeMail1.Exchange.Modified

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "Two Exchange Emails",
			entries: func() []Entry {
				res := []Entry{}

				for _, entry := range []Entry{exchangeMail1, exchangeMail2} {
					e := entry
					ei := *entry.Exchange
					e.Exchange = &ei

					res = append(res, e)
				}

				return res
			},
			expectedDirs: func() []Entry {
				res := []Entry{}

				for _, entry := range exchangeFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.Size = exchangeMail1.Exchange.Size + exchangeMail2.Exchange.Size
					e.Folder.Modified = exchangeMail2.Exchange.Modified

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "One Email And One Contact",
			entries: func() []Entry {
				res := []Entry{}

				for _, entry := range []Entry{exchangeMail1, exchangeContact1} {
					e := entry
					ei := *entry.Exchange
					e.Exchange = &ei

					res = append(res, e)
				}

				return res
			},
			expectedDirs: func() []Entry {
				res := []Entry{}

				for _, entry := range exchangeFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.Size = exchangeMail1.Exchange.Size
					e.Folder.Modified = exchangeMail1.Exchange.Modified

					res = append(res, e)
				}

				for _, entry := range exchangeContactFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.Size = exchangeContact1.Exchange.Size
					e.Folder.Modified = exchangeContact1.Exchange.Modified

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "One OneDrive Item",
			entries: func() []Entry {
				e := oneDrive1
				ei := *oneDrive1.OneDrive
				e.OneDrive = &ei

				return []Entry{e}
			},
			expectedDirs: func() []Entry {
				res := []Entry{}

				for _, entry := range oneDriveishFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.DataType = OneDriveItem
					e.Folder.Size = oneDrive1.OneDrive.Size
					e.Folder.Modified = oneDrive1.OneDrive.Modified

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "One SharePoint Item",
			entries: func() []Entry {
				e := sharePoint1
				ei := *sharePoint1.SharePoint
				e.SharePoint = &ei

				return []Entry{e}
			},
			expectedDirs: func() []Entry {
				res := []Entry{}

				for _, entry := range oneDriveishFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.DataType = SharePointLibrary
					e.Folder.Size = sharePoint1.SharePoint.Size
					e.Folder.Modified = sharePoint1.SharePoint.Modified

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "One SharePoint Legacy Item",
			entries: func() []Entry {
				e := sharePoint1
				ei := *sharePoint1.SharePoint
				e.SharePoint = &ei

				return []Entry{e}
			},
			expectedDirs: func() []Entry {
				res := []Entry{}

				for _, entry := range oneDriveishFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.DataType = SharePointLibrary
					e.Folder.Size = sharePoint1.SharePoint.Size
					e.Folder.Modified = sharePoint1.SharePoint.Modified

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "One Group SharePoint Item",
			entries: func() []Entry {
				e := group1
				ei := *group1.Groups
				e.Groups = &ei

				return []Entry{e}
			},
			expectedDirs: func() []Entry {
				res := []Entry{}

				for _, entry := range oneDriveishFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.DataType = SharePointLibrary
					e.Folder.Size = group1.Groups.Size
					e.Folder.Modified = group1.Groups.Modified

					res = append(res, e)
				}

				return res
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			db := &Builder{}

			for _, entry := range test.entries() {
				rr, err := path.FromDataLayerPath(entry.RepoRef, true)
				require.NoError(t, err, clues.ToCore(err))

				loc, err := path.Builder{}.SplitUnescapeAppend(entry.LocationRef)
				require.NoError(t, err, clues.ToCore(err))

				err = db.Add(rr, loc, entry.ItemInfo)
				require.NoError(t, err, clues.ToCore(err))
			}

			deets := db.Details()
			gotDirs := []Entry{}

			for _, entry := range deets.Entries {
				// Other test checks items are populated properly.
				if entry.infoType() != FolderItem {
					continue
				}

				// Not Comparing these right now.
				entry.RepoRef = ""
				entry.ShortRef = ""
				entry.ParentRef = ""

				gotDirs = append(gotDirs, entry)
			}

			assert.ElementsMatch(t, test.expectedDirs(), gotDirs)
		})
	}
}

var pathItemsTable = []struct {
	name               string
	ents               []Entry
	expectRepoRefs     []string
	expectLocationRefs []string
}{
	{
		name:               "nil entries",
		ents:               nil,
		expectRepoRefs:     []string{},
		expectLocationRefs: []string{},
	},
	{
		name: "single entry",
		ents: []Entry{
			{
				RepoRef:     "abcde",
				LocationRef: "locationref",
				ItemRef:     "itemref",
			},
		},
		expectRepoRefs:     []string{"abcde"},
		expectLocationRefs: []string{"locationref"},
	},
	{
		name: "multiple entries",
		ents: []Entry{
			{
				RepoRef:     "abcde",
				LocationRef: "locationref",
				ItemRef:     "itemref",
			},
			{
				RepoRef:     "12345",
				LocationRef: "locationref2",
				ItemRef:     "itemref2",
			},
		},
		expectRepoRefs:     []string{"abcde", "12345"},
		expectLocationRefs: []string{"locationref", "locationref2"},
	},
	{
		name: "multiple entries with folder",
		ents: []Entry{
			{
				RepoRef:     "abcde",
				LocationRef: "locationref",
				ItemRef:     "itemref",
			},
			{
				RepoRef:     "12345",
				LocationRef: "locationref2",
				ItemRef:     "itemref2",
			},
			{
				RepoRef:     "deadbeef",
				LocationRef: "locationref3",
				ItemInfo: ItemInfo{
					Folder: &FolderInfo{
						DisplayName: "test folder",
					},
				},
			},
		},
		expectRepoRefs:     []string{"abcde", "12345"},
		expectLocationRefs: []string{"locationref", "locationref2"},
	},
	{
		name: "multiple entries with meta file",
		ents: []Entry{
			{
				RepoRef:     "abcde",
				LocationRef: "locationref",
			},
			{
				RepoRef:     "foo.meta",
				LocationRef: "locationref.dirmeta",
				ItemRef:     "itemref.meta",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: false},
				},
			},
			{
				RepoRef:     "is-meta-file",
				LocationRef: "locationref-meta-file",
				ItemRef:     "itemref-meta-file",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: true},
				},
			},
		},
		expectRepoRefs:     []string{"abcde", "foo.meta"},
		expectLocationRefs: []string{"locationref", "locationref.dirmeta"},
	},
	{
		name: "multiple entries with folder and meta file",
		ents: []Entry{
			{
				RepoRef:     "abcde",
				LocationRef: "locationref",
				ItemRef:     "itemref",
			},
			{
				RepoRef:     "12345",
				LocationRef: "locationref2",
				ItemRef:     "itemref2",
			},
			{
				RepoRef:     "foo.meta",
				LocationRef: "locationref.dirmeta",
				ItemRef:     "itemref.dirmeta",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: false},
				},
			},
			{
				RepoRef:     "is-meta-file",
				LocationRef: "locationref-meta-file",
				ItemRef:     "itemref-meta-file",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: true},
				},
			},
			{
				RepoRef:     "deadbeef",
				LocationRef: "locationref3",
				ItemRef:     "itemref3",
				ItemInfo: ItemInfo{
					Folder: &FolderInfo{
						DisplayName: "test folder",
					},
				},
			},
		},
		expectRepoRefs:     []string{"abcde", "12345", "foo.meta"},
		expectLocationRefs: []string{"locationref", "locationref2", "locationref.dirmeta"},
	},
}

func (suite *DetailsUnitSuite) TestDetailsModel_Path() {
	for _, test := range pathItemsTable {
		suite.Run(test.name, func() {
			t := suite.T()

			d := Details{
				DetailsModel: DetailsModel{
					Entries: test.ents,
				},
			}
			assert.ElementsMatch(t, test.expectRepoRefs, d.Paths())
		})
	}
}

func (suite *DetailsUnitSuite) TestDetailsModel_Items() {
	for _, test := range pathItemsTable {
		suite.Run(test.name, func() {
			t := suite.T()

			d := Details{
				DetailsModel: DetailsModel{
					Entries: test.ents,
				},
			}

			ents := d.Items()
			assert.Len(t, ents, len(test.expectRepoRefs))

			for _, e := range ents {
				assert.Contains(t, test.expectRepoRefs, e.RepoRef)
				assert.Contains(t, test.expectLocationRefs, e.LocationRef)
			}
		})
	}
}

func (suite *DetailsUnitSuite) TestDetailsModel_FilterMetaFiles() {
	t := suite.T()

	d := &DetailsModel{
		Entries: []Entry{
			{
				RepoRef: "a.data",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: false},
				},
			},
			{
				RepoRef: "b.meta",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: false},
				},
			},
			{
				RepoRef: "c.meta",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: true},
				},
			},
		},
	}

	d2 := d.FilterMetaFiles()

	assert.Len(t, d2.Entries, 2)
	assert.Len(t, d.Entries, 3)
}

func (suite *DetailsUnitSuite) TestBuilder_Add_shortRefsUniqueFromFolder() {
	t := suite.T()

	b := Builder{}
	name := "itemName"
	info := ItemInfo{
		OneDrive: &OneDriveInfo{
			ItemType:  OneDriveItem,
			ItemName:  name,
			DriveName: "drive-name",
			DriveID:   "drive-id",
		},
	}

	itemPath := makeItemPath(
		t,
		path.OneDriveService,
		path.FilesCategory,
		"a-tenant",
		"a-user",
		[]string{
			"drive-id",
			odConsts.RootPathDir,
			"folder",
			name + "-id",
		})

	otherItemPath := makeItemPath(
		t,
		path.OneDriveService,
		path.FilesCategory,
		"a-tenant",
		"a-user",
		[]string{
			"drive-id",
			odConsts.RootPathDir,
			"folder",
			name + "-id",
			name,
		})

	err := b.Add(
		itemPath,
		// Don't need to generate folders for this entry we just want the ShortRef.
		&path.Builder{},
		info)
	require.NoError(t, err, clues.ToCore(err))

	items := b.Details().Items()
	require.Len(t, items, 1)

	// If the ShortRefs match then it means it's possible for the user to
	// construct folder names such that they'll generate a ShortRef collision.
	assert.NotEqual(t, otherItemPath.ShortRef(), items[0].ShortRef, "same ShortRef as subfolder item")
}

func (suite *DetailsUnitSuite) TestBuilder_Add_cleansFileIDSuffixes() {
	var (
		t    = suite.T()
		b    = Builder{}
		svc  = path.OneDriveService
		cat  = path.FilesCategory
		info = ItemInfo{
			OneDrive: &OneDriveInfo{
				ItemType:  OneDriveItem,
				ItemName:  "in",
				DriveName: "dn",
				DriveID:   "d",
			},
		}

		dataSfx    = makeItemPath(t, svc, cat, "t", "u", []string{"d", "r:", "f", "i1" + metadata.DataFileSuffix})
		dirMetaSfx = makeItemPath(t, svc, cat, "t", "u", []string{"d", "r:", "f", "i1" + metadata.DirMetaFileSuffix})
		metaSfx    = makeItemPath(t, svc, cat, "t", "u", []string{"d", "r:", "f", "i1" + metadata.MetaFileSuffix})
	)

	// Don't need to generate folders for this entry, we just want the itemRef
	loc := &path.Builder{}

	err := b.Add(dataSfx, loc, info)
	require.NoError(t, err, clues.ToCore(err))

	err = b.Add(dirMetaSfx, loc, info)
	require.NoError(t, err, clues.ToCore(err))

	err = b.Add(metaSfx, loc, info)
	require.NoError(t, err, clues.ToCore(err))

	for _, ent := range b.Details().Items() {
		assert.False(t, strings.HasSuffix(ent.ItemRef, metadata.DirMetaFileSuffix))
		assert.False(t, strings.HasSuffix(ent.ItemRef, metadata.MetaFileSuffix))
		assert.False(t, strings.HasSuffix(ent.ItemRef, metadata.DataFileSuffix))
	}
}

func (suite *DetailsUnitSuite) TestBuilder_DetailsNoDuplicate() {
	var (
		t    = suite.T()
		b    = Builder{}
		svc  = path.OneDriveService
		cat  = path.FilesCategory
		info = ItemInfo{
			OneDrive: &OneDriveInfo{
				ItemType:  OneDriveItem,
				ItemName:  "in",
				DriveName: "dn",
				DriveID:   "d",
			},
		}

		dataSfx    = makeItemPath(t, svc, cat, "t", "u", []string{"d", "r:", "f", "i1" + metadata.DataFileSuffix})
		dataSfx2   = makeItemPath(t, svc, cat, "t", "u", []string{"d", "r:", "f", "i2" + metadata.DataFileSuffix})
		dirMetaSfx = makeItemPath(t, svc, cat, "t", "u", []string{"d", "r:", "f", "i1" + metadata.DirMetaFileSuffix})
		metaSfx    = makeItemPath(t, svc, cat, "t", "u", []string{"d", "r:", "f", "i1" + metadata.MetaFileSuffix})
	)

	// Don't need to generate folders for this entry, we just want the itemRef
	loc := &path.Builder{}

	err := b.Add(dataSfx, loc, info)
	require.NoError(t, err, clues.ToCore(err))

	err = b.Add(dataSfx2, loc, info)
	require.NoError(t, err, clues.ToCore(err))

	err = b.Add(dirMetaSfx, loc, info)
	require.NoError(t, err, clues.ToCore(err))

	err = b.Add(metaSfx, loc, info)
	require.NoError(t, err, clues.ToCore(err))

	b.knownFolders = map[string]Entry{
		"dummy": {
			RepoRef:     "xyz",
			ShortRef:    "abcd",
			ParentRef:   "1234",
			LocationRef: "ab",
			ItemRef:     "cd",
			ItemInfo:    info,
		},
		"dummy2": {
			RepoRef:     "xyz2",
			ShortRef:    "abcd2",
			ParentRef:   "12342",
			LocationRef: "ab2",
			ItemRef:     "cd2",
			ItemInfo:    info,
		},
		"dummy3": {
			RepoRef:     "xyz3",
			ShortRef:    "abcd3",
			ParentRef:   "12343",
			LocationRef: "ab3",
			ItemRef:     "cd3",
			ItemInfo:    info,
		},
	}

	// mark the capacity prior to calling details.
	// if the entries slice gets modified and grows to a
	// 5th space, then the capacity would grow as well.
	capCheck := cap(b.d.Entries)

	assert.Len(t, b.Details().Entries, 7) // 4 ents + 3 known folders
	assert.Len(t, b.Details().Entries, 7) // possible reason for err: knownFolders got added twice

	assert.Len(t, b.d.Entries, 4)               // len should not have grown
	assert.Equal(t, capCheck, cap(b.d.Entries)) // capacity should not have grown
}

func makeItemPath(
	t *testing.T,
	service path.ServiceType,
	category path.CategoryType,
	tenant, resourceOwner string,
	elems []string,
) path.Path {
	t.Helper()

	p, err := path.Build(
		tenant,
		resourceOwner,
		service,
		category,
		true,
		elems...)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func (suite *DetailsUnitSuite) TestUpdateItem() {
	const (
		folder1 = "f1"
		folder2 = "f2"
	)

	newExchangePB := path.Builder{}.Append(folder2)
	newOneDrivePB := path.Builder{}.Append(odConsts.RootPathDir, folder2)

	table := []struct {
		name         string
		input        ItemInfo
		locPath      *path.Builder
		expectedItem ItemInfo
	}{
		{
			name: "ExchangeEvent",
			input: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeEvent,
					ParentPath: folder1,
				},
			},
			locPath: newExchangePB,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeEvent,
					ParentPath: folder2,
				},
			},
		},
		{
			name: "ExchangeContact",
			input: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeContact,
					ParentPath: folder1,
				},
			},
			locPath: newExchangePB,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeContact,
					ParentPath: folder2,
				},
			},
		},
		{
			name: "ExchangeMail",
			input: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeMail,
					ParentPath: folder1,
				},
			},
			locPath: newExchangePB,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeMail,
					ParentPath: folder2,
				},
			},
		},
		{
			name: "OneDrive",
			input: ItemInfo{
				OneDrive: &OneDriveInfo{
					ItemType:   OneDriveItem,
					ParentPath: folder1,
				},
			},
			locPath: newOneDrivePB,
			expectedItem: ItemInfo{
				OneDrive: &OneDriveInfo{
					ItemType:   OneDriveItem,
					ParentPath: folder2,
				},
			},
		},
		{
			name: "SharePoint",
			input: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   SharePointLibrary,
					ParentPath: folder1,
				},
			},
			locPath: newOneDrivePB,
			expectedItem: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   SharePointLibrary,
					ParentPath: folder2,
				},
			},
		},
		{
			name: "SharePoint Old Format",
			input: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   OneDriveItem,
					ParentPath: folder1,
				},
			},
			locPath: newOneDrivePB,
			expectedItem: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   SharePointLibrary,
					ParentPath: folder2,
				},
			},
		},
		{
			name:         "Empty Item Doesn't Fail",
			input:        ItemInfo{},
			locPath:      newOneDrivePB,
			expectedItem: ItemInfo{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			item := test.input
			UpdateItem(&item, test.locPath)
			assert.Equal(t, test.expectedItem, item)
		})
	}
}

func (suite *DetailsUnitSuite) TestDetails_Marshal() {
	for _, test := range pathItemsTable {
		suite.Run(test.name, func() {
			d := &Details{DetailsModel: DetailsModel{
				Entries: test.ents,
			}}

			bs, err := d.Marshal()
			require.NoError(suite.T(), err, clues.ToCore(err))
			assert.NotEmpty(suite.T(), bs)
		})
	}
}

func (suite *DetailsUnitSuite) TestUnarshalTo() {
	for _, test := range pathItemsTable {
		suite.Run(test.name, func() {
			orig := &Details{DetailsModel: DetailsModel{
				Entries: test.ents,
			}}

			bs, err := orig.Marshal()
			require.NoError(suite.T(), err, clues.ToCore(err))
			assert.NotEmpty(suite.T(), bs)

			var result Details
			umt := UnmarshalTo(&result)
			err = umt(io.NopCloser(bytes.NewReader(bs)))

			t := suite.T()
			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, result)
			assert.ElementsMatch(t, orig.Entries, result.Entries)
		})
	}
}

func (suite *DetailsUnitSuite) TestLocationIDer_FromEntry() {
	const (
		rrString = "tenant-id/%s/user-id/%s/drives/drive-id/root:/some/folder/stuff/item"
		driveID  = "driveID"

		expectedUniqueLocFmt         = "%s/" + driveID + "/root:/some/folder/stuff"
		expectedExchangeUniqueLocFmt = "%s/root:/some/folder/stuff"
		expectedDetailsLoc           = "root:/some/folder/stuff"
	)

	table := []struct {
		name              string
		service           string
		category          string
		itemInfo          ItemInfo
		hasLocRef         bool
		backupVersion     int
		expectedErr       require.ErrorAssertionFunc
		expectedUniqueLoc string
	}{
		{
			name:     "OneDrive With Drive ID Old Version",
			service:  path.OneDriveService.String(),
			category: path.FilesCategory.String(),
			itemInfo: ItemInfo{
				OneDrive: &OneDriveInfo{
					ItemType: OneDriveItem,
					DriveID:  driveID,
				},
			},
			backupVersion:     version.OneDrive7LocationRef - 1,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedUniqueLocFmt, path.FilesCategory),
		},
		{
			name:     "OneDrive With Drive ID And LocationRef",
			service:  path.OneDriveService.String(),
			category: path.FilesCategory.String(),
			itemInfo: ItemInfo{
				OneDrive: &OneDriveInfo{
					ItemType: OneDriveItem,
					DriveID:  driveID,
				},
			},
			backupVersion:     version.OneDrive7LocationRef,
			hasLocRef:         true,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedUniqueLocFmt, path.FilesCategory),
		},
		{
			name:     "OneDrive With Drive ID New Version Errors",
			service:  path.OneDriveService.String(),
			category: path.FilesCategory.String(),
			itemInfo: ItemInfo{
				OneDrive: &OneDriveInfo{
					ItemType: OneDriveItem,
					DriveID:  driveID,
				},
			},
			backupVersion: version.OneDrive7LocationRef,
			expectedErr:   require.Error,
		},
		{
			name:     "SharePoint With Drive ID Old Version",
			service:  path.SharePointService.String(),
			category: path.LibrariesCategory.String(),
			itemInfo: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType: SharePointLibrary,
					DriveID:  driveID,
				},
			},
			backupVersion:     version.OneDrive7LocationRef - 1,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedUniqueLocFmt, path.LibrariesCategory),
		},
		{
			name:     "SharePoint With Drive ID And LocationRef",
			service:  path.SharePointService.String(),
			category: path.LibrariesCategory.String(),
			itemInfo: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType: SharePointLibrary,
					DriveID:  driveID,
				},
			},
			backupVersion:     version.OneDrive7LocationRef,
			hasLocRef:         true,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedUniqueLocFmt, path.LibrariesCategory),
		},
		{
			name:     "SharePoint With Drive ID New Version Errors",
			service:  path.SharePointService.String(),
			category: path.LibrariesCategory.String(),
			itemInfo: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType: SharePointLibrary,
					DriveID:  driveID,
				},
			},
			backupVersion: version.OneDrive7LocationRef,
			expectedErr:   require.Error,
		},
		{
			name:     "Exchange Email With LocationRef Old Version",
			service:  path.ExchangeService.String(),
			category: path.EmailCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			backupVersion:     version.OneDrive7LocationRef - 1,
			hasLocRef:         true,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedExchangeUniqueLocFmt, path.EmailCategory),
		},
		{
			name:     "Exchange Email With LocationRef New Version",
			service:  path.ExchangeService.String(),
			category: path.EmailCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			backupVersion:     version.OneDrive7LocationRef,
			hasLocRef:         true,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedExchangeUniqueLocFmt, path.EmailCategory),
		},
		{
			name:     "Exchange Email Without LocationRef Old Version",
			service:  path.ExchangeService.String(),
			category: path.EmailCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			backupVersion:     version.OneDrive7LocationRef - 1,
			hasLocRef:         true,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedExchangeUniqueLocFmt, path.EmailCategory),
		},
		{
			name:     "Exchange Email Without LocationRef New Version",
			service:  path.ExchangeService.String(),
			category: path.EmailCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			backupVersion:     version.OneDrive7LocationRef,
			hasLocRef:         true,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedExchangeUniqueLocFmt, path.EmailCategory),
		},
		{
			name:     "Exchange Email Bad RepoRef Fails",
			service:  path.OneDriveService.String(),
			category: path.EmailCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			backupVersion: version.OneDrive7LocationRef,
			expectedErr:   require.Error,
		},
		{
			name:     "Exchange Event Empty LocationRef New Version Fails",
			service:  path.ExchangeService.String(),
			category: path.EventsCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeEvent,
				},
			},
			backupVersion: 2,
			expectedErr:   require.Error,
		},
		{
			name:     "Exchange Event Empty LocationRef Old Version",
			service:  path.ExchangeService.String(),
			category: path.EventsCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeEvent,
				},
			},
			backupVersion:     version.OneDrive1DataAndMetaFiles,
			hasLocRef:         true,
			expectedErr:       require.NoError,
			expectedUniqueLoc: fmt.Sprintf(expectedExchangeUniqueLocFmt, path.EventsCategory),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			entry := Entry{
				RepoRef:  fmt.Sprintf(rrString, test.service, test.category),
				ItemInfo: test.itemInfo,
			}

			if test.hasLocRef {
				entry.LocationRef = expectedDetailsLoc
			}

			loc, err := entry.ToLocationIDer(test.backupVersion)
			test.expectedErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(
				t,
				test.expectedUniqueLoc,
				loc.ID().String(),
				"unique location")
			assert.Equal(
				t,
				expectedDetailsLoc,
				loc.InDetails().String(),
				"details location")
		})
	}
}
