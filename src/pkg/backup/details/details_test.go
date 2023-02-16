package details

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
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
	require.NoError(suite.T(), err)

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
						ItemType: ExchangeMail,
						Sender:   "sender",
						Subject:  "subject",
						Received: now,
					},
				},
			},
			expectHs: []string{"ID", "Sender", "Subject", "Received"},
			expectVs: []string{"deadbeef", "sender", "subject", nowStr},
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
						DriveName:  "aDrive",
						Created:    now,
						Modified:   now,
					},
				},
			},
			expectHs: []string{"ID", "ItemName", "Drive", "ParentPath", "Size", "WebURL", "Created", "Modified"},
			expectVs: []string{
				"deadbeef",
				"itemName",
				"aDrive",
				"parentPath",
				"1.0 kB",
				"https://not.a.real/url",
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

	p, err := path.Builder{}.Append(elems...).
		ToDataLayerPath(
			tenant,
			resourceOwner,
			service,
			category,
			true,
		)
	require.NoError(t, err)

	return p
}

func (suite *DetailsUnitSuite) TestUpdateItem() {
	const (
		tenant        = "a-tenant"
		resourceOwner = "a-user"
		driveID       = "abcd"
		folder1       = "f1"
		folder2       = "f2"
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
		locPath      path.Path
		errCheck     assert.ErrorAssertionFunc
		expectedItem ItemInfo
	}{
		{
			name: "ExchangeEventNoChange",
			input: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeEvent,
				},
			},
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeEvent,
				},
			},
		},
		{
			name: "ExchangeContactNoChange",
			input: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeContact,
				},
			},
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeContact,
				},
			},
		},
		{
			name: "ExchangeMailNoChange",
			input: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
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
			locPath:  newOneDrivePath,
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
					ItemType:   SharePointItem,
					ParentPath: folder1,
				},
			},
			repoPath: newOneDrivePath,
			locPath:  newOneDrivePath,
			errCheck: assert.NoError,
			expectedItem: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   SharePointItem,
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
			locPath:  badOneDrivePath,
			errCheck: assert.Error,
		},
		{
			name: "SharePointBadPath",
			input: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType:   SharePointItem,
					ParentPath: folder1,
				},
			},
			repoPath: badOneDrivePath,
			locPath:  badOneDrivePath,
			errCheck: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			item := test.input
			err := UpdateItem(&item, test.repoPath)
			test.errCheck(t, err)

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
			name:     "base path with location",
			parent:   basePath,
			location: basePath,
			expect:   baseFolderEnts,
		},
		{
			name:   "single depth parent only",
			parent: basePath.Append(fnords...),
			expect: folderEntriesFor(fnords, nil),
		},
		{
			name:     "single depth with location",
			parent:   basePath.Append(fnords...),
			location: basePath.Append(beau...),
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
			location: basePath.Append(regard...),
			expect:   folderEntriesFor(smarf, regard),
		},
		{
			name:     "mismatched depth, parent longer",
			parent:   basePath.Append(smarf...),
			location: basePath.Append(beau...),
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
