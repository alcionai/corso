package selectors

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

type GroupsSelectorSuite struct {
	tester.Suite
}

func TestGroupsSelectorSuite(t *testing.T) {
	suite.Run(t, &GroupsSelectorSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsSelectorSuite) TestNewGroupsBackup() {
	t := suite.T()
	ob := NewGroupsBackup(nil)
	assert.Equal(t, ob.Service, ServiceGroups)
	assert.NotZero(t, ob.Scopes())
}

func (suite *GroupsSelectorSuite) TestToGroupsBackup() {
	t := suite.T()
	ob := NewGroupsBackup(nil)
	s := ob.Selector
	ob, err := s.ToGroupsBackup()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ob.Service, ServiceGroups)
	assert.NotZero(t, ob.Scopes())
}

func (suite *GroupsSelectorSuite) TestNewGroupsRestore() {
	t := suite.T()
	or := NewGroupsRestore(nil)
	assert.Equal(t, or.Service, ServiceGroups)
	assert.NotZero(t, or.Scopes())
}

func (suite *GroupsSelectorSuite) TestToGroupsRestore() {
	t := suite.T()
	eb := NewGroupsRestore(nil)
	s := eb.Selector
	or, err := s.ToGroupsRestore()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, or.Service, ServiceGroups)
	assert.NotZero(t, or.Scopes())
}

// TODO(rkeepers): implement
// func (suite *GroupsSelectorSuite) TestGroupsRestore_Reduce() {
// 	toRR := func(cat path.CategoryType, siteID string, folders []string, item string) string {
// 		folderElems := make([]string, 0, len(folders))

// 		for _, f := range folders {
// 			folderElems = append(folderElems, f+".d")
// 		}

// 		return stubRepoRef(
// 			path.GroupsService,
// 			cat,
// 			siteID,
// 			strings.Join(folderElems, "/"),
// 			item)
// 	}

// 	var (
// 		prefixElems = []string{
// 			odConsts.DrivesPathDir,
// 			"drive!id",
// 			odConsts.RootPathDir,
// 		}
// 		itemElems1 = []string{"folderA", "folderB"}
// 		itemElems2 = []string{"folderA", "folderC"}
// 		itemElems3 = []string{"folderD", "folderE"}
// 		pairAC     = "folderA/folderC"
// 		pairGH     = "folderG/folderH"
// 		item       = toRR(
// 			path.LibrariesCategory,
// 			"sid",
// 			append(slices.Clone(prefixElems), itemElems1...),
// 			"item")
// 		item2 = toRR(
// 			path.LibrariesCategory,
// 			"sid",
// 			append(slices.Clone(prefixElems), itemElems2...),
// 			"item2")
// 		item3 = toRR(
// 			path.LibrariesCategory,
// 			"sid",
// 			append(slices.Clone(prefixElems), itemElems3...),
// 			"item3")
// 		item4 = stubRepoRef(path.GroupsService, path.PagesCategory, "sid", pairGH, "item4")
// 		item5 = stubRepoRef(path.GroupsService, path.PagesCategory, "sid", pairGH, "item5")
// 	)

// 	deets := &details.Details{
// 		DetailsModel: details.DetailsModel{
// 			Entries: []details.Entry{
// 				{
// 					RepoRef:     item,
// 					ItemRef:     "item",
// 					LocationRef: strings.Join(append([]string{odConsts.RootPathDir}, itemElems1...), "/"),
// 					ItemInfo: details.ItemInfo{
// 						Groups: &details.GroupsInfo{
// 							ItemType:   details.GroupsLibrary,
// 							ItemName:   "itemName",
// 							ParentPath: strings.Join(itemElems1, "/"),
// 						},
// 					},
// 				},
// 				{
// 					RepoRef:     item2,
// 					LocationRef: strings.Join(append([]string{odConsts.RootPathDir}, itemElems2...), "/"),
// 					// ItemRef intentionally blank to test fallback case
// 					ItemInfo: details.ItemInfo{
// 						Groups: &details.GroupsInfo{
// 							ItemType:   details.GroupsLibrary,
// 							ItemName:   "itemName2",
// 							ParentPath: strings.Join(itemElems2, "/"),
// 						},
// 					},
// 				},
// 				{
// 					RepoRef:     item3,
// 					ItemRef:     "item3",
// 					LocationRef: strings.Join(append([]string{odConsts.RootPathDir}, itemElems3...), "/"),
// 					ItemInfo: details.ItemInfo{
// 						Groups: &details.GroupsInfo{
// 							ItemType:   details.GroupsLibrary,
// 							ItemName:   "itemName3",
// 							ParentPath: strings.Join(itemElems3, "/"),
// 						},
// 					},
// 				},
// 				{
// 					RepoRef:     item4,
// 					LocationRef: pairGH,
// 					ItemRef:     "item4",
// 					ItemInfo: details.ItemInfo{
// 						Groups: &details.GroupsInfo{
// 							ItemType:   details.GroupsPage,
// 							ItemName:   "itemName4",
// 							ParentPath: pairGH,
// 						},
// 					},
// 				},
// 				{
// 					RepoRef:     item5,
// 					LocationRef: pairGH,
// 					// ItemRef intentionally blank to test fallback case
// 					ItemInfo: details.ItemInfo{
// 						Groups: &details.GroupsInfo{
// 							ItemType:   details.GroupsPage,
// 							ItemName:   "itemName5",
// 							ParentPath: pairGH,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	arr := func(s ...string) []string {
// 		return s
// 	}

// 	table := []struct {
// 		name         string
// 		makeSelector func() *GroupsRestore
// 		expect       []string
// 		cfg          Config
// 	}{
// 		{
// 			name: "all",
// 			makeSelector: func() *GroupsRestore {
// 				odr := NewGroupsRestore(Any())
// 				odr.Include(odr.AllData())
// 				return odr
// 			},
// 			expect: arr(item, item2, item3, item4, item5),
// 		},
// 		{
// 			name: "only match item",
// 			makeSelector: func() *GroupsRestore {
// 				odr := NewGroupsRestore(Any())
// 				odr.Include(odr.LibraryItems(Any(), []string{"item2"}))
// 				return odr
// 			},
// 			expect: arr(item2),
// 		},
// 		{
// 			name: "id doesn't match name",
// 			makeSelector: func() *GroupsRestore {
// 				odr := NewGroupsRestore(Any())
// 				odr.Include(odr.LibraryItems(Any(), []string{"item2"}))
// 				return odr
// 			},
// 			expect: []string{},
// 			cfg:    Config{OnlyMatchItemNames: true},
// 		},
// 		{
// 			name: "only match item name",
// 			makeSelector: func() *GroupsRestore {
// 				odr := NewGroupsRestore(Any())
// 				odr.Include(odr.LibraryItems(Any(), []string{"itemName2"}))
// 				return odr
// 			},
// 			expect: arr(item2),
// 			cfg:    Config{OnlyMatchItemNames: true},
// 		},
// 		{
// 			name: "name doesn't match",
// 			makeSelector: func() *GroupsRestore {
// 				odr := NewGroupsRestore(Any())
// 				odr.Include(odr.LibraryItems(Any(), []string{"itemName2"}))
// 				return odr
// 			},
// 			expect: []string{},
// 		},
// 		{
// 			name: "only match folder",
// 			makeSelector: func() *GroupsRestore {
// 				odr := NewGroupsRestore([]string{"sid"})
// 				odr.Include(odr.LibraryFolders([]string{"folderA/folderB", pairAC}))
// 				return odr
// 			},
// 			expect: arr(item, item2),
// 		},
// 		{
// 			name: "pages match folder",
// 			makeSelector: func() *GroupsRestore {
// 				odr := NewGroupsRestore([]string{"sid"})
// 				odr.Include(odr.Pages([]string{pairGH, pairAC}))
// 				return odr
// 			},
// 			expect: arr(item4, item5),
// 		},
// 	}
// 	for _, test := range table {
// 		suite.Run(test.name, func() {
// 			t := suite.T()

// 			ctx, flush := tester.NewContext(t)
// 			defer flush()

// 			sel := test.makeSelector()
// 			sel.Configure(test.cfg)
// 			results := sel.Reduce(ctx, deets, fault.New(true))
// 			paths := results.Paths()
// 			assert.Equal(t, test.expect, paths)
// 		})
// 	}
// }

func (suite *GroupsSelectorSuite) TestGroupsCategory_PathValues() {
	var (
		itemName = "item"
		itemID   = "item-id"
		shortRef = "short"
		elems    = []string{itemID}
	)

	table := []struct {
		name       string
		sc         groupsCategory
		pathElems  []string
		locRef     string
		parentPath string
		expected   map[categorizer][]string
		cfg        Config
	}{
		{
			name:      "Groups Channel Messages",
			sc:        GroupsChannelMessage,
			pathElems: elems,
			locRef:    "",
			expected: map[categorizer][]string{
				GroupsChannel:        {""},
				GroupsChannelMessage: {itemID, shortRef},
			},
			cfg: Config{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			itemPath, err := path.Build(
				"tenant",
				"site",
				path.GroupsService,
				test.sc.PathType(),
				true,
				test.pathElems...)
			require.NoError(t, err, clues.ToCore(err))

			ent := details.Entry{
				RepoRef:     itemPath.String(),
				ShortRef:    shortRef,
				ItemRef:     itemPath.Item(),
				LocationRef: test.locRef,
				ItemInfo: details.ItemInfo{
					Groups: &details.GroupsInfo{
						ItemName:   itemName,
						ParentPath: test.parentPath,
					},
				},
			}

			pv, err := test.sc.pathValues(itemPath, ent, test.cfg)
			require.NoError(t, err)
			assert.Equal(t, test.expected, pv)
		})
	}
}

// TODO(abin): implement
// func (suite *GroupsSelectorSuite) TestGroupsScope_MatchesInfo() {
// 	var (
// 		sel          = NewGroupsRestore(Any())
// 		host         = "www.website.com"
// 		pth          = "/foo"
// 		url          = host + pth
// 		epoch        = time.Time{}
// 		now          = time.Now()
// 		modification = now.Add(15 * time.Minute)
// 		future       = now.Add(45 * time.Minute)
// 	)

// 	table := []struct {
// 		name    string
// 		infoURL string
// 		scope   []GroupsScope
// 		expect  assert.BoolAssertionFunc
// 	}{
// 		{"host match", host, sel.WebURL([]string{host}), assert.True},
// 		{"url match", url, sel.WebURL([]string{url}), assert.True},
// 		{"host suffixes host", host, sel.WebURL([]string{host}, SuffixMatch()), assert.True},
// 		{"url does not suffix host", url, sel.WebURL([]string{host}, SuffixMatch()), assert.False},
// 		{"url has path suffix", url, sel.WebURL([]string{pth}, SuffixMatch()), assert.True},
// 		{"host does not contain substring", host, sel.WebURL([]string{"website"}), assert.False},
// 		{"url does not suffix substring", url, sel.WebURL([]string{"oo"}, SuffixMatch()), assert.False},
// 		{"host mismatch", host, sel.WebURL([]string{"www.google.com"}), assert.False},
// 		{"file create after the epoch", host, sel.CreatedAfter(dttm.Format(epoch)), assert.True},
// 		{"file create after now", host, sel.CreatedAfter(dttm.Format(now)), assert.False},
// 		{"file create after later", url, sel.CreatedAfter(dttm.Format(future)), assert.False},
// 		{"file create before future", host, sel.CreatedBefore(dttm.Format(future)), assert.True},
// 		{"file create before now", host, sel.CreatedBefore(dttm.Format(now)), assert.False},
// 		{"file create before modification", host, sel.CreatedBefore(dttm.Format(modification)), assert.True},
// 		{"file create before epoch", host, sel.CreatedBefore(dttm.Format(now)), assert.False},
// 		{"file modified after the epoch", host, sel.ModifiedAfter(dttm.Format(epoch)), assert.True},
// 		{"file modified after now", host, sel.ModifiedAfter(dttm.Format(now)), assert.True},
// 		{"file modified after later", host, sel.ModifiedAfter(dttm.Format(future)), assert.False},
// 		{"file modified before future", host, sel.ModifiedBefore(dttm.Format(future)), assert.True},
// 		{"file modified before now", host, sel.ModifiedBefore(dttm.Format(now)), assert.False},
// 		{"file modified before epoch", host, sel.ModifiedBefore(dttm.Format(now)), assert.False},
// 		{"in library", host, sel.Library("included-library"), assert.True},
// 		{"not in library", host, sel.Library("not-included-library"), assert.False},
// 		{"library id", host, sel.Library("1234"), assert.True},
// 		{"not library id", host, sel.Library("abcd"), assert.False},
// 	}
// 	for _, test := range table {
// 		suite.Run(test.name, func() {
// 			t := suite.T()

// 			itemInfo := details.ItemInfo{
// 				Groups: &details.GroupsInfo{
// 					ItemType:  details.GroupsPage,
// 					WebURL:    test.infoURL,
// 					Created:   now,
// 					Modified:  modification,
// 					DriveName: "included-library",
// 					DriveID:   "1234",
// 				},
// 			}

// 			scopes := setScopesToDefault(test.scope)
// 			for _, scope := range scopes {
// 				test.expect(t, scope.matchesInfo(itemInfo))
// 			}
// 		})
// 	}
// }

func (suite *GroupsSelectorSuite) TestCategory_PathType() {
	table := []struct {
		cat      groupsCategory
		pathType path.CategoryType
	}{
		{GroupsCategoryUnknown, path.UnknownCategory},
		{GroupsChannel, path.UnknownCategory},
		{GroupsChannelMessage, path.ChannelMessagesCategory},
	}
	for _, test := range table {
		suite.Run(test.cat.String(), func() {
			assert.Equal(suite.T(), test.pathType, test.cat.PathType())
		})
	}
}
