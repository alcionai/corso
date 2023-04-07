package details

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
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
	nowStr := common.FormatTimeWith(initial, common.TabularOutput)
	now, err := common.ParseTime(nowStr)
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name     string
		entry    DetailsEntry
		expectHs []string
		expectVs []string
	}{
		{
			name: "no info",
			entry: DetailsEntry{
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
			entry: DetailsEntry{
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
			entry: DetailsEntry{
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
			entry: DetailsEntry{
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
			name: "sharepoint info",
			entry: DetailsEntry{
				RepoRef:     "reporef",
				ShortRef:    "deadbeef",
				LocationRef: "locationref",
				ItemRef:     "itemref",
				ItemInfo: ItemInfo{
					SharePoint: &SharePointInfo{
						ItemName:   "itemName",
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
			name: "oneDrive info",
			entry: DetailsEntry{
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

			hs := test.entry.Headers()
			assert.Equal(t, test.expectHs, hs)
			vs := test.entry.Values()
			assert.Equal(t, test.expectVs, vs)
		})
	}
}

func exchangeEntry(t *testing.T, id string, size int, it ItemType) DetailsEntry {
	rr := makeItemPath(
		t,
		path.ExchangeService,
		path.EmailCategory,
		"tenant-id",
		"user-id",
		[]string{"Inbox", "folder1", id})

	return DetailsEntry{
		RepoRef:     rr.String(),
		ShortRef:    rr.ShortRef(),
		ParentRef:   rr.ToBuilder().Dir().ShortRef(),
		LocationRef: rr.Folder(true),
		ItemInfo: ItemInfo{
			Exchange: &ExchangeInfo{
				ItemType: it,
				Modified: time.Now(),
				Size:     int64(size),
			},
		},
	}
}

func oneDriveishEntry(t *testing.T, id string, size int, it ItemType) DetailsEntry {
	service := path.OneDriveService
	category := path.FilesCategory
	info := ItemInfo{
		OneDrive: &OneDriveInfo{
			ItemName:  "bar",
			DriveID:   "drive-id",
			DriveName: "drive-name",
			Modified:  time.Now(),
			ItemType:  it,
			Size:      int64(size),
		},
	}

	if it == SharePointLibrary {
		service = path.SharePointService
		category = path.LibrariesCategory

		info.OneDrive = nil
		info.SharePoint = &SharePointInfo{
			ItemName:  "bar",
			DriveID:   "drive-id",
			DriveName: "drive-name",
			Modified:  time.Now(),
			ItemType:  it,
			Size:      int64(size),
		}
	}

	rr := makeItemPath(
		t,
		service,
		category,
		"tenant-id",
		"user-id",
		[]string{
			"drives",
			"drive-id",
			"root:",
			"Inbox",
			"folder1",
			id,
		})

	loc := path.Builder{}.Append(rr.Folders()...).PopFront().PopFront()

	return DetailsEntry{
		RepoRef:     rr.String(),
		ShortRef:    rr.ShortRef(),
		ParentRef:   rr.ToBuilder().Dir().ShortRef(),
		LocationRef: loc.String(),
		ItemInfo:    info,
	}
}

func (suite *DetailsUnitSuite) TestDetailsAdd_NoLocationFolders() {
	t := suite.T()
	table := []struct {
		name  string
		entry DetailsEntry
		// shortRefEqual allows checking that OneDrive and SharePoint have their
		// ShortRef updated in the returned entry.
		//
		// TODO(ashmrtn): Remove this when we don't need extra munging for
		// OneDrive/SharePoint file name changes.
		shortRefEqual assert.ComparisonAssertionFunc
	}{
		{
			name:          "Exchange Email",
			entry:         exchangeEntry(t, "foo", 42, ExchangeMail),
			shortRefEqual: assert.Equal,
		},
		{
			name:          "OneDrive File",
			entry:         oneDriveishEntry(t, "foo", 42, OneDriveItem),
			shortRefEqual: assert.NotEqual,
		},
		{
			name:          "SharePoint File",
			entry:         oneDriveishEntry(t, "foo", 42, SharePointLibrary),
			shortRefEqual: assert.NotEqual,
		},
		{
			name: "Legacy SharePoint File",
			entry: func() DetailsEntry {
				res := oneDriveishEntry(t, "foo", 42, SharePointLibrary)
				res.SharePoint.ItemType = OneDriveItem

				return res
			}(),
			shortRefEqual: assert.NotEqual,
		},
	}

	for _, test := range table {
		for _, updated := range []bool{false, true} {
			suite.Run(fmt.Sprintf("%s Updated %v", test.name, updated), func() {
				t := suite.T()

				rr, err := path.FromDataLayerPath(test.entry.RepoRef, true)
				require.NoError(t, err, clues.ToCore(err))

				db := &Builder{}

				// Make a local copy so we can modify it.
				localItem := test.entry

				err = db.Add(rr, &path.Builder{}, updated, localItem.ItemInfo)
				require.NoError(t, err, clues.ToCore(err))

				// Clear LocationRef that's automatically populated since we passed an
				// empty builder above.
				localItem.LocationRef = ""
				localItem.Updated = updated

				expectedShortRef := localItem.ShortRef
				localItem.ShortRef = ""

				deets := db.Details()
				assert.Len(t, deets.Entries, 1)

				got := deets.Entries[0]
				gotShortRef := got.ShortRef
				got.ShortRef = ""

				assert.Equal(t, localItem, got, "DetailsEntry")
				test.shortRefEqual(t, expectedShortRef, gotShortRef, "ShortRef")
			})
		}
	}
}

func (suite *DetailsUnitSuite) TestDetailsAdd_LocationFolders() {
	t := suite.T()

	exchangeMail1 := exchangeEntry(t, "foo1", 42, ExchangeMail)
	oneDrive1 := oneDriveishEntry(t, "foo1", 42, OneDriveItem)
	sharePoint1 := oneDriveishEntry(t, "foo1", 42, SharePointLibrary)
	sharePointLegacy1 := oneDriveishEntry(t, "foo1", 42, SharePointLibrary)
	sharePointLegacy1.SharePoint.ItemType = OneDriveItem

	// Sleep for a little so we get a larger difference in mod times between the
	// earlier and later entries.
	time.Sleep(100 * time.Millisecond)

	// Get fresh item IDs so we can check that folders populate with the latest
	// mod time. Also the details API is built with the idea that duplicate items
	// aren't added (it has no checking for that).
	exchangeMail2 := exchangeEntry(t, "foo2", 43, ExchangeMail)
	exchangeContact1 := exchangeEntry(t, "foo3", 44, ExchangeContact)

	exchangeFolders := []DetailsEntry{
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

	exchangeContactFolders := []DetailsEntry{
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

	oneDriveishFolders := []DetailsEntry{
		{
			ItemInfo: ItemInfo{
				Folder: &FolderInfo{
					DisplayName: "root:",
					ItemType:    FolderItem,
					DriveName:   "drive-name",
					DriveID:     "drive-id",
				},
			},
		},
		{
			LocationRef: "root:",
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
		entries      func() []DetailsEntry
		expectedDirs func() []DetailsEntry
	}{
		{
			name: "One Exchange Email None Updated",
			entries: func() []DetailsEntry {
				e := exchangeMail1
				ei := *exchangeMail1.Exchange
				e.Exchange = &ei

				return []DetailsEntry{e}
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

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
			name: "One Exchange Email Updated",
			entries: func() []DetailsEntry {
				e := exchangeMail1
				ei := *exchangeMail1.Exchange
				e.Exchange = &ei
				e.Updated = true

				return []DetailsEntry{e}
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

				for _, entry := range exchangeFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.Size = exchangeMail1.Exchange.Size
					e.Folder.Modified = exchangeMail1.Exchange.Modified
					e.Updated = true

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "Two Exchange Emails None Updated",
			entries: func() []DetailsEntry {
				res := []DetailsEntry{}

				for _, entry := range []DetailsEntry{exchangeMail1, exchangeMail2} {
					e := entry
					ei := *entry.Exchange
					e.Exchange = &ei

					res = append(res, e)
				}

				return res
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

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
			name: "Two Exchange Emails One Updated",
			entries: func() []DetailsEntry {
				res := []DetailsEntry{}

				for i, entry := range []DetailsEntry{exchangeMail1, exchangeMail2} {
					e := entry
					ei := *entry.Exchange
					e.Exchange = &ei
					e.Updated = i == 1

					res = append(res, e)
				}

				return res
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

				for _, entry := range exchangeFolders {
					e := entry
					ei := *entry.Folder

					e.Folder = &ei
					e.Folder.Size = exchangeMail1.Exchange.Size + exchangeMail2.Exchange.Size
					e.Folder.Modified = exchangeMail2.Exchange.Modified
					e.Updated = true

					res = append(res, e)
				}

				return res
			},
		},
		{
			name: "One Email And One Contact None Updated",
			entries: func() []DetailsEntry {
				res := []DetailsEntry{}

				for _, entry := range []DetailsEntry{exchangeMail1, exchangeContact1} {
					e := entry
					ei := *entry.Exchange
					e.Exchange = &ei

					res = append(res, e)
				}

				return res
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

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
			name: "One OneDrive Item None Updated",
			entries: func() []DetailsEntry {
				e := oneDrive1
				ei := *oneDrive1.OneDrive
				e.OneDrive = &ei

				return []DetailsEntry{e}
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

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
			name: "One SharePoint Item None Updated",
			entries: func() []DetailsEntry {
				e := sharePoint1
				ei := *sharePoint1.SharePoint
				e.SharePoint = &ei

				return []DetailsEntry{e}
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

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
			name: "One SharePoint Legacy Item None Updated",
			entries: func() []DetailsEntry {
				e := sharePoint1
				ei := *sharePoint1.SharePoint
				e.SharePoint = &ei

				return []DetailsEntry{e}
			},
			expectedDirs: func() []DetailsEntry {
				res := []DetailsEntry{}

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

				err = db.Add(rr, loc, entry.Updated, entry.ItemInfo)
				require.NoError(t, err, clues.ToCore(err))
			}

			deets := db.Details()
			gotDirs := []DetailsEntry{}

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
	ents               []DetailsEntry
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
		ents: []DetailsEntry{
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
		ents: []DetailsEntry{
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
		ents: []DetailsEntry{
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
		ents: []DetailsEntry{
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
		ents: []DetailsEntry{
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
		Entries: []DetailsEntry{
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

func (suite *DetailsUnitSuite) TestDetails_Add_ShortRefs_Unique_From_Folder() {
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
			"root:",
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
			"root:",
			"folder",
			name + "-id",
			name,
		})

	err := b.Add(
		itemPath,
		// Don't need to generate folders for this entry we just want the ShortRef.
		&path.Builder{},
		false,
		info)
	require.NoError(t, err)

	items := b.Details().Items()
	require.Len(t, items, 1)

	// If the ShortRefs match then it means it's possible for the user to
	// construct folder names such that they'll generate a ShortRef collision.
	assert.NotEqual(t, otherItemPath.ShortRef(), items[0].ShortRef, "same ShortRef as subfolder item")
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
	newOneDrivePB := path.Builder{}.Append("root:", folder2)

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
			name:     "Exchange Email Without LocationRef Old Version Errors",
			service:  path.ExchangeService.String(),
			category: path.EmailCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			backupVersion: version.OneDrive7LocationRef - 1,
			expectedErr:   require.Error,
		},
		{
			name:     "Exchange Email Without LocationRef New Version Errors",
			service:  path.ExchangeService.String(),
			category: path.EmailCategory.String(),
			itemInfo: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			backupVersion: version.OneDrive7LocationRef,
			expectedErr:   require.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			entry := DetailsEntry{
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
