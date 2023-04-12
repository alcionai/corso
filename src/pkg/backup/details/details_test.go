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
			},
			{
				RepoRef:     "12345",
				LocationRef: "locationref2",
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
			},
			{
				RepoRef:     "12345",
				LocationRef: "locationref2",
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
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: false},
				},
			},
			{
				RepoRef:     "is-meta-file",
				LocationRef: "locationref-meta-file",
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
			},
			{
				RepoRef:     "12345",
				LocationRef: "locationref2",
			},
			{
				RepoRef:     "foo.meta",
				LocationRef: "locationref.dirmeta",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: false},
				},
			},
			{
				RepoRef:     "is-meta-file",
				LocationRef: "locationref-meta-file",
				ItemInfo: ItemInfo{
					OneDrive: &OneDriveInfo{IsMeta: true},
				},
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

func (suite *DetailsUnitSuite) TestDetails_Add_ShortRefs() {
	itemNames := []string{
		"item1",
		"item2",
	}

	table := []struct {
		name               string
		service            path.ServiceType
		category           path.CategoryType
		itemInfoFunc       func(name string) ItemInfo
		expectedUniqueRefs int
	}{
		{
			name:     "OneDrive",
			service:  path.OneDriveService,
			category: path.FilesCategory,
			itemInfoFunc: func(name string) ItemInfo {
				return ItemInfo{
					OneDrive: &OneDriveInfo{
						ItemType: OneDriveItem,
						ItemName: name,
					},
				}
			},
			expectedUniqueRefs: len(itemNames),
		},
		{
			name:     "SharePoint",
			service:  path.SharePointService,
			category: path.LibrariesCategory,
			itemInfoFunc: func(name string) ItemInfo {
				return ItemInfo{
					SharePoint: &SharePointInfo{
						ItemType: SharePointLibrary,
						ItemName: name,
					},
				}
			},
			expectedUniqueRefs: len(itemNames),
		},
		{
			name:     "SharePoint List",
			service:  path.SharePointService,
			category: path.ListsCategory,
			itemInfoFunc: func(name string) ItemInfo {
				return ItemInfo{
					SharePoint: &SharePointInfo{
						ItemType: SharePointList,
						ItemName: name,
					},
				}
			},
			// Should all end up as the starting shortref.
			expectedUniqueRefs: 1,
		},
		{
			name:     "Exchange no change",
			service:  path.ExchangeService,
			category: path.EmailCategory,
			itemInfoFunc: func(name string) ItemInfo {
				return ItemInfo{
					Exchange: &ExchangeInfo{
						ItemType:  ExchangeMail,
						Sender:    "a-person@foo.com",
						Subject:   name,
						Recipient: []string{"another-person@bar.com"},
					},
				}
			},
			// Should all end up as the starting shortref.
			expectedUniqueRefs: 1,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			b := Builder{}

			for _, name := range itemNames {
				item := test.itemInfoFunc(name)
				itemPath := makeItemPath(
					suite.T(),
					test.service,
					test.category,
					"a-tenant",
					"a-user",
					[]string{
						"drive-id",
						"root:",
						"folder",
						name + "-id",
					},
				)

				require.NoError(t, b.Add(
					itemPath.String(),
					"deadbeef",
					itemPath.ToBuilder().Dir().String(),
					itemPath.String(),
					false,
					item,
				))
			}

			deets := b.Details()
			shortRefs := map[string]struct{}{}

			for _, d := range deets.Items() {
				shortRefs[d.ShortRef] = struct{}{}
			}

			assert.Len(t, shortRefs, test.expectedUniqueRefs, "items don't have unique ShortRefs")
		})
	}
}

func (suite *DetailsUnitSuite) TestDetails_Add_ShortRefs_Unique_From_Folder() {
	t := suite.T()

	b := Builder{}
	name := "itemName"
	info := ItemInfo{
		OneDrive: &OneDriveInfo{
			ItemType: OneDriveItem,
			ItemName: name,
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
		},
	)

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
		},
	)

	err := b.Add(
		itemPath.String(),
		"deadbeef",
		itemPath.ToBuilder().Dir().String(),
		itemPath.String(),
		false,
		info)
	require.NoError(t, err)

	items := b.Details().Items()
	require.Len(t, items, 1)

	// If the ShortRefs match then it means it's possible for the user to
	// construct folder names such that they'll generate a ShortRef collision.
	assert.NotEqual(t, otherItemPath.ShortRef(), items[0].ShortRef, "same ShortRef as subfolder item")
}

func (suite *DetailsUnitSuite) TestDetails_AddFolders() {
	itemTime := time.Date(2022, 10, 21, 10, 0, 0, 0, time.UTC)
	folderTimeOlderThanItem := time.Date(2022, 9, 21, 10, 0, 0, 0, time.UTC)
	folderTimeNewerThanItem := time.Date(2022, 11, 21, 10, 0, 0, 0, time.UTC)

	itemInfo := ItemInfo{
		Exchange: &ExchangeInfo{
			Size:     20,
			Modified: itemTime,
		},
	}

	table := []struct {
		name               string
		folders            []folderEntry
		expectedShortRefs  []string
		expectedFolderInfo map[string]FolderInfo
	}{
		{
			name: "MultipleFolders",
			folders: []folderEntry{
				{
					RepoRef:     "rr1",
					ShortRef:    "sr1",
					ParentRef:   "pr1",
					LocationRef: "lr1",
					Info: ItemInfo{
						Folder: &FolderInfo{
							Modified: folderTimeOlderThanItem,
						},
					},
				},
				{
					RepoRef:     "rr2",
					ShortRef:    "sr2",
					ParentRef:   "pr2",
					LocationRef: "lr2",
					Info: ItemInfo{
						Folder: &FolderInfo{
							Modified: folderTimeNewerThanItem,
						},
					},
				},
			},
			expectedShortRefs: []string{"sr1", "sr2"},
			expectedFolderInfo: map[string]FolderInfo{
				"sr1": {Size: 20, Modified: itemTime},
				"sr2": {Size: 20, Modified: folderTimeNewerThanItem},
			},
		},
		{
			name: "MultipleFoldersWithRepeats",
			folders: []folderEntry{
				{
					RepoRef:     "rr1",
					ShortRef:    "sr1",
					ParentRef:   "pr1",
					LocationRef: "lr1",
					Info: ItemInfo{
						Folder: &FolderInfo{
							Modified: folderTimeOlderThanItem,
						},
					},
				},
				{
					RepoRef:     "rr2",
					ShortRef:    "sr2",
					ParentRef:   "pr2",
					LocationRef: "lr2",
					Info: ItemInfo{
						Folder: &FolderInfo{
							Modified: folderTimeOlderThanItem,
						},
					},
				},
				{
					RepoRef:     "rr1",
					ShortRef:    "sr1",
					ParentRef:   "pr1",
					LocationRef: "lr1",
					Info: ItemInfo{
						Folder: &FolderInfo{
							Modified: folderTimeOlderThanItem,
						},
					},
				},
				{
					RepoRef:     "rr3",
					ShortRef:    "sr3",
					ParentRef:   "pr3",
					LocationRef: "lr3",
					Info: ItemInfo{
						Folder: &FolderInfo{
							Modified: folderTimeNewerThanItem,
						},
					},
				},
			},
			expectedShortRefs: []string{"sr1", "sr2", "sr3"},
			expectedFolderInfo: map[string]FolderInfo{
				// Two items were added
				"sr1": {Size: 40, Modified: itemTime},
				"sr2": {Size: 20, Modified: itemTime},
				"sr3": {Size: 20, Modified: folderTimeNewerThanItem},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			builder := Builder{}
			builder.AddFoldersForItem(test.folders, itemInfo, true)
			deets := builder.Details()
			assert.Len(t, deets.Entries, len(test.expectedShortRefs))

			for _, e := range deets.Entries {
				assert.Contains(t, test.expectedShortRefs, e.ShortRef)
				assert.Equal(t, test.expectedFolderInfo[e.ShortRef].Size, e.Folder.Size)
				assert.Equal(t, test.expectedFolderInfo[e.ShortRef].Modified, e.Folder.Modified)
			}
		})
	}
}

func (suite *DetailsUnitSuite) TestDetails_AddFoldersUpdate() {
	itemInfo := ItemInfo{
		Exchange: &ExchangeInfo{},
	}

	table := []struct {
		name                       string
		folders                    []folderEntry
		itemUpdated                bool
		expectedFolderUpdatedValue map[string]bool
	}{
		{
			name: "ItemNotUpdated_NoChange",
			folders: []folderEntry{
				{
					RepoRef:     "rr1",
					ShortRef:    "sr1",
					ParentRef:   "pr1",
					LocationRef: "lr1",
					Info: ItemInfo{
						Folder: &FolderInfo{},
					},
					Updated: true,
				},
				{
					RepoRef:     "rr2",
					ShortRef:    "sr2",
					ParentRef:   "pr2",
					LocationRef: "lr2",
					Info: ItemInfo{
						Folder: &FolderInfo{},
					},
				},
			},
			itemUpdated: false,
			expectedFolderUpdatedValue: map[string]bool{
				"sr1": true,
				"sr2": false,
			},
		},
		{
			name: "ItemUpdated",
			folders: []folderEntry{
				{
					RepoRef:     "rr1",
					ShortRef:    "sr1",
					ParentRef:   "pr1",
					LocationRef: "lr1",
					Info: ItemInfo{
						Folder: &FolderInfo{},
					},
				},
				{
					RepoRef:     "rr2",
					ShortRef:    "sr2",
					ParentRef:   "pr2",
					LocationRef: "lr2",
					Info: ItemInfo{
						Folder: &FolderInfo{},
					},
				},
			},
			itemUpdated: true,
			expectedFolderUpdatedValue: map[string]bool{
				"sr1": true,
				"sr2": true,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			builder := Builder{}
			builder.AddFoldersForItem(test.folders, itemInfo, test.itemUpdated)
			deets := builder.Details()
			assert.Len(t, deets.Entries, len(test.expectedFolderUpdatedValue))

			for _, e := range deets.Entries {
				assert.Equalf(
					t,
					test.expectedFolderUpdatedValue[e.ShortRef],
					e.Updated, "%s updated value incorrect",
					e.ShortRef)
			}
		})
	}
}

func (suite *DetailsUnitSuite) TestDetails_AddFoldersDifferentServices() {
	itemTime := time.Date(2022, 10, 21, 10, 0, 0, 0, time.UTC)

	table := []struct {
		name               string
		item               ItemInfo
		expectedFolderInfo FolderInfo
	}{
		{
			name: "Exchange",
			item: ItemInfo{
				Exchange: &ExchangeInfo{
					Size:     20,
					Modified: itemTime,
				},
			},
			expectedFolderInfo: FolderInfo{
				Size:     20,
				Modified: itemTime,
			},
		},
		{
			name: "OneDrive",
			item: ItemInfo{
				OneDrive: &OneDriveInfo{
					Size:     20,
					Modified: itemTime,
				},
			},
			expectedFolderInfo: FolderInfo{
				Size:     20,
				Modified: itemTime,
			},
		},
		{
			name: "SharePoint",
			item: ItemInfo{
				SharePoint: &SharePointInfo{
					Size:     20,
					Modified: itemTime,
				},
			},
			expectedFolderInfo: FolderInfo{
				Size:     20,
				Modified: itemTime,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			folder := folderEntry{
				RepoRef:     "rr1",
				ShortRef:    "sr1",
				ParentRef:   "pr1",
				LocationRef: "lr1",
				Info: ItemInfo{
					Folder: &FolderInfo{},
				},
			}

			builder := Builder{}
			builder.AddFoldersForItem([]folderEntry{folder}, test.item, true)
			deets := builder.Details()
			require.Len(t, deets.Entries, 1)

			got := deets.Entries[0].Folder

			assert.Equal(t, test.expectedFolderInfo, *got)
		})
	}
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
		tenant        = "a-tenant"
		resourceOwner = "a-user"
		driveID       = "abcd"
		folder1       = "f1"
		folder2       = "f2"
		folder3       = "f3"
		item          = "hello.txt"
	)

	// Making both OneDrive paths is alright because right now they're the same as
	// SharePoint path and there's no extra validation.
	newOneDrivePath := makeItemPath(
		suite.T(),
		path.OneDriveService,
		path.FilesCategory,
		tenant,
		resourceOwner,
		[]string{
			"drives",
			driveID,
			"root:",
			folder2,
			item,
		},
	)
	newExchangePB := path.Builder{}.Append(folder3)
	badOneDrivePath := makeItemPath(
		suite.T(),
		path.OneDriveService,
		path.FilesCategory,
		tenant,
		resourceOwner,
		[]string{item},
	)

	table := []struct {
		name         string
		input        ItemInfo
		repoPath     path.Path
		locPath      *path.Builder
		errCheck     assert.ErrorAssertionFunc
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
			repoPath: newOneDrivePath,
			locPath:  newExchangePB,
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeEvent,
					ParentPath: folder3,
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
			repoPath: newOneDrivePath,
			locPath:  newExchangePB,
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeContact,
					ParentPath: folder3,
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
			repoPath: newOneDrivePath,
			locPath:  newExchangePB,
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType:   ExchangeMail,
					ParentPath: folder3,
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
			repoPath: newOneDrivePath,
			locPath:  newExchangePB,
			errCheck: assert.NoError,
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
			repoPath: newOneDrivePath,
			locPath:  newExchangePB,
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   SharePointLibrary,
					ParentPath: folder2,
				},
			},
		},
		{
			name: "OneDriveBadPath",
			input: ItemInfo{
				OneDrive: &OneDriveInfo{
					ItemType:   OneDriveItem,
					ParentPath: folder1,
				},
			},
			repoPath: badOneDrivePath,
			locPath:  newExchangePB,
			errCheck: assert.Error,
		},
		{
			name: "SharePointBadPath",
			input: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   SharePointLibrary,
					ParentPath: folder1,
				},
			},
			repoPath: badOneDrivePath,
			locPath:  newExchangePB,
			errCheck: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			item := test.input

			err := UpdateItem(&item, test.repoPath, test.locPath)
			test.errCheck(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expectedItem, item)
		})
	}
}

var (
	basePath       = path.Builder{}.Append("ten", "serv", "user", "type")
	baseFolderEnts = []folderEntry{
		{
			RepoRef:     basePath.String(),
			ShortRef:    basePath.ShortRef(),
			ParentRef:   basePath.Dir().ShortRef(),
			LocationRef: "",
			Info: ItemInfo{
				Folder: &FolderInfo{
					ItemType:    FolderItem,
					DisplayName: "type",
				},
			},
		},
		{
			RepoRef:     basePath.Dir().String(),
			ShortRef:    basePath.Dir().ShortRef(),
			ParentRef:   basePath.Dir().Dir().ShortRef(),
			LocationRef: "",
			Info: ItemInfo{
				Folder: &FolderInfo{
					ItemType:    FolderItem,
					DisplayName: "user",
				},
			},
		},
		{
			RepoRef:     basePath.Dir().Dir().String(),
			ShortRef:    basePath.Dir().Dir().ShortRef(),
			ParentRef:   basePath.Dir().Dir().Dir().ShortRef(),
			LocationRef: "",
			Info: ItemInfo{
				Folder: &FolderInfo{
					ItemType:    FolderItem,
					DisplayName: "serv",
				},
			},
		},
		{
			RepoRef:     basePath.Dir().Dir().Dir().String(),
			ShortRef:    basePath.Dir().Dir().Dir().ShortRef(),
			ParentRef:   "",
			LocationRef: "",
			Info: ItemInfo{
				Folder: &FolderInfo{
					ItemType:    FolderItem,
					DisplayName: "ten",
				},
			},
		},
	}
)

func folderEntriesFor(pathElems []string, locElems []string) []folderEntry {
	p := basePath.Append(pathElems...)
	l := path.Builder{}.Append(locElems...)

	ents := make([]folderEntry, 0, len(pathElems)+4)

	for range pathElems {
		dn := p.LastElem()
		if l != nil && len(l.Elements()) > 0 {
			dn = l.LastElem()
		}

		fe := folderEntry{
			RepoRef:     p.String(),
			ShortRef:    p.ShortRef(),
			ParentRef:   p.Dir().ShortRef(),
			LocationRef: l.String(),
			Info: ItemInfo{
				Folder: &FolderInfo{
					ItemType:    FolderItem,
					DisplayName: dn,
				},
			},
		}

		l = l.Dir()
		p = p.Dir()

		ents = append(ents, fe)
	}

	return append(ents, baseFolderEnts...)
}

func (suite *DetailsUnitSuite) TestFolderEntriesForPath() {
	var (
		fnords = []string{"fnords"}
		smarf  = []string{"fnords", "smarf"}
		beau   = []string{"beau"}
		regard = []string{"beau", "regard"}
	)

	table := []struct {
		name     string
		parent   *path.Builder
		location *path.Builder
		expect   []folderEntry
	}{
		{
			name:   "base path, parent only",
			parent: basePath,
			expect: baseFolderEnts,
		},
		{
			name:   "single depth parent only",
			parent: basePath.Append(fnords...),
			expect: folderEntriesFor(fnords, nil),
		},
		{
			name:     "single depth with location",
			parent:   basePath.Append(fnords...),
			location: path.Builder{}.Append(beau...),
			expect:   folderEntriesFor(fnords, beau),
		},
		{
			name:   "two depth parent only",
			parent: basePath.Append(smarf...),
			expect: folderEntriesFor(smarf, nil),
		},
		{
			name:     "two depth with location",
			parent:   basePath.Append(smarf...),
			location: path.Builder{}.Append(regard...),
			expect:   folderEntriesFor(smarf, regard),
		},
		{
			name:     "mismatched depth, parent longer",
			parent:   basePath.Append(smarf...),
			location: path.Builder{}.Append(beau...),
			expect:   folderEntriesFor(smarf, beau),
		},
		// We can't handle this right now.  But we don't have any cases
		// which immediately require it, either.  Keeping in the test
		// as a reminder that this might be required at some point.
		// {
		// 	name:     "mismatched depth, location longer",
		// 	parent:   basePath.Append(fnords...),
		// 	location: basePath.Append(regard...),
		// 	expect:   folderEntriesFor(fnords, regard),
		// },
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := FolderEntriesForPath(test.parent, test.location)
			assert.ElementsMatch(t, test.expect, result)
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

func (suite *DetailsUnitSuite) TestUniqueLocation_FromEntry() {
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
			backupVersion:     version.OneDriveXLocationRef - 1,
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
			backupVersion:     version.OneDriveXLocationRef,
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
			backupVersion: version.OneDriveXLocationRef,
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
			backupVersion:     version.OneDriveXLocationRef - 1,
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
			backupVersion:     version.OneDriveXLocationRef,
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
			backupVersion: version.OneDriveXLocationRef,
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
			backupVersion:     version.OneDriveXLocationRef - 1,
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
			backupVersion:     version.OneDriveXLocationRef,
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
			backupVersion: version.OneDriveXLocationRef - 1,
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
			backupVersion: version.OneDriveXLocationRef,
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

			loc, err := entry.UniqueLocation(test.backupVersion)
			test.expectedErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(
				t,
				test.expectedUniqueLoc,
				loc.UniqueLocation().String(),
				"unique location")
			assert.Equal(
				t,
				expectedDetailsLoc,
				loc.DetailsLocation().String(),
				"details location")
		})
	}
}
