package selectors

import (
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type SharePointSelectorSuite struct {
	tester.Suite
}

func TestSharePointSelectorSuite(t *testing.T) {
	suite.Run(t, &SharePointSelectorSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SharePointSelectorSuite) TestNewSharePointBackup() {
	t := suite.T()
	ob := NewSharePointBackup(nil)
	assert.Equal(t, ob.Service, ServiceSharePoint)
	assert.NotZero(t, ob.Scopes())
}

func (suite *SharePointSelectorSuite) TestToSharePointBackup() {
	t := suite.T()
	ob := NewSharePointBackup(nil)
	s := ob.Selector
	ob, err := s.ToSharePointBackup()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ob.Service, ServiceSharePoint)
	assert.NotZero(t, ob.Scopes())
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_AllData() {
	t := suite.T()

	sites := []string{"s1", "s2"}

	sel := NewSharePointBackup(sites)
	siteScopes := sel.AllData()

	assert.ElementsMatch(t, sites, sel.DiscreteResourceOwners())

	// Initialize the selector Include, Exclude, Filter
	sel.Exclude(siteScopes)
	sel.Include(siteScopes)
	sel.Filter(siteScopes)

	table := []struct {
		name          string
		scopesToCheck []scope
	}{
		{"Include Scopes", sel.Includes},
		{"Exclude Scopes", sel.Excludes},
		{"info scopes", sel.Filters},
	}
	for _, test := range table {
		require.Len(t, test.scopesToCheck, 3)

		for _, scope := range test.scopesToCheck {
			var (
				spsc = SharePointScope(scope)
				cat  = spsc.Category()
			)

			suite.Run(test.name+"-"+cat.String(), func() {
				t := suite.T()

				switch cat {
				case SharePointLibraryItem:
					scopeMustHave(
						t,
						spsc,
						map[categorizer][]string{
							SharePointLibraryItem:   Any(),
							SharePointLibraryFolder: Any(),
						},
					)
				case SharePointListItem:
					scopeMustHave(
						t,
						spsc,
						map[categorizer][]string{
							SharePointListItem: Any(),
							SharePointList:     Any(),
						},
					)
				}
			})
		}
	}
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_Include_WebURLs() {
	t := suite.T()

	const (
		s1 = "s1"
		s2 = "s2"
	)

	s12 := []string{s1, s2}

	sel := NewSharePointRestore(s12)
	sel.Include(sel.WebURL(s12))
	scopes := sel.Includes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			SharePointScope(sc),
			map[categorizer][]string{SharePointWebURL: s12},
		)
	}
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_Include_WebURLs_anyNone() {
	table := []struct {
		name   string
		in     []string
		expect []string
	}{
		{
			name:   "any",
			in:     []string{AnyTgt},
			expect: Any(),
		},
		{
			name:   "none",
			in:     []string{NoneTgt},
			expect: None(),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			sel := NewSharePointRestore(Any())
			sel.Include(sel.WebURL(test.in))
			scopes := sel.Includes
			require.Len(t, scopes, 3)

			for _, sc := range scopes {
				scopeMustHave(
					t,
					SharePointScope(sc),
					map[categorizer][]string{SharePointWebURL: test.expect},
				)
			}
		})
	}
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_Exclude_WebURLs() {
	t := suite.T()

	const (
		s1 = "s1"
		s2 = "s2"
	)

	s12 := []string{s1, s2}

	sel := NewSharePointRestore(s12)
	sel.Exclude(sel.WebURL(s12))
	scopes := sel.Excludes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			SharePointScope(sc),
			map[categorizer][]string{SharePointWebURL: s12},
		)
	}
}

func (suite *SharePointSelectorSuite) TestNewSharePointRestore() {
	t := suite.T()
	or := NewSharePointRestore(nil)
	assert.Equal(t, or.Service, ServiceSharePoint)
	assert.NotZero(t, or.Scopes())
}

func (suite *SharePointSelectorSuite) TestToSharePointRestore() {
	t := suite.T()
	eb := NewSharePointRestore(nil)
	s := eb.Selector
	or, err := s.ToSharePointRestore()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, or.Service, ServiceSharePoint)
	assert.NotZero(t, or.Scopes())
}

func (suite *SharePointSelectorSuite) TestSharePointRestore_Reduce() {
	toRR := func(cat path.CategoryType, siteID string, folders []string, item string) string {
		folderElems := make([]string, 0, len(folders))

		for _, f := range folders {
			folderElems = append(folderElems, f+".d")
		}

		return stubRepoRef(
			path.SharePointService,
			cat,
			siteID,
			strings.Join(folderElems, "/"),
			item)
	}

	var (
		prefixElems = []string{
			"drive",
			"drive!id",
			"root:",
		}
		itemElems1 = []string{"folderA", "folderB"}
		itemElems2 = []string{"folderA", "folderC"}
		itemElems3 = []string{"folderD", "folderE"}
		pairAC     = "folderA/folderC"
		pairGH     = "folderG/folderH"
		item       = toRR(
			path.LibrariesCategory,
			"sid",
			append(slices.Clone(prefixElems), itemElems1...),
			"item")
		item2 = toRR(
			path.LibrariesCategory,
			"sid",
			append(slices.Clone(prefixElems), itemElems2...),
			"item2")
		item3 = toRR(
			path.LibrariesCategory,
			"sid",
			append(slices.Clone(prefixElems), itemElems3...),
			"item3")
		item4 = stubRepoRef(path.SharePointService, path.PagesCategory, "sid", pairGH, "item4")
		item5 = stubRepoRef(path.SharePointService, path.PagesCategory, "sid", pairGH, "item5")
	)

	deets := &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: []details.Entry{
				{
					RepoRef:     item,
					ItemRef:     "item",
					LocationRef: strings.Join(append([]string{"root:"}, itemElems1...), "/"),
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType:   details.SharePointLibrary,
							ItemName:   "itemName",
							ParentPath: strings.Join(itemElems1, "/"),
						},
					},
				},
				{
					RepoRef:     item2,
					LocationRef: strings.Join(append([]string{"root:"}, itemElems2...), "/"),
					// ItemRef intentionally blank to test fallback case
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType:   details.SharePointLibrary,
							ItemName:   "itemName2",
							ParentPath: strings.Join(itemElems2, "/"),
						},
					},
				},
				{
					RepoRef:     item3,
					ItemRef:     "item3",
					LocationRef: strings.Join(append([]string{"root:"}, itemElems3...), "/"),
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType:   details.SharePointLibrary,
							ItemName:   "itemName3",
							ParentPath: strings.Join(itemElems3, "/"),
						},
					},
				},
				{
					RepoRef:     item4,
					LocationRef: pairGH,
					ItemRef:     "item4",
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType:   details.SharePointPage,
							ItemName:   "itemName4",
							ParentPath: pairGH,
						},
					},
				},
				{
					RepoRef:     item5,
					LocationRef: pairGH,
					// ItemRef intentionally blank to test fallback case
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType:   details.SharePointPage,
							ItemName:   "itemName5",
							ParentPath: pairGH,
						},
					},
				},
			},
		},
	}

	arr := func(s ...string) []string {
		return s
	}

	table := []struct {
		name         string
		makeSelector func() *SharePointRestore
		expect       []string
		cfg          Config
	}{
		{
			name: "all",
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore(Any())
				odr.Include(odr.AllData())
				return odr
			},
			expect: arr(item, item2, item3, item4, item5),
		},
		{
			name: "only match item",
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore(Any())
				odr.Include(odr.LibraryItems(Any(), []string{"item2"}))
				return odr
			},
			expect: arr(item2),
		},
		{
			name: "id doesn't match name",
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore(Any())
				odr.Include(odr.LibraryItems(Any(), []string{"item2"}))
				return odr
			},
			expect: []string{},
			cfg:    Config{OnlyMatchItemNames: true},
		},
		{
			name: "only match item name",
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore(Any())
				odr.Include(odr.LibraryItems(Any(), []string{"itemName2"}))
				return odr
			},
			expect: arr(item2),
			cfg:    Config{OnlyMatchItemNames: true},
		},
		{
			name: "name doesn't match",
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore(Any())
				odr.Include(odr.LibraryItems(Any(), []string{"itemName2"}))
				return odr
			},
			expect: []string{},
		},
		{
			name: "only match folder",
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore([]string{"sid"})
				odr.Include(odr.LibraryFolders([]string{"folderA/folderB", pairAC}))
				return odr
			},
			expect: arr(item, item2),
		},
		{
			name: "pages match folder",
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore([]string{"sid"})
				odr.Include(odr.Pages([]string{pairGH, pairAC}))
				return odr
			},
			expect: arr(item4, item5),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			sel := test.makeSelector()
			sel.Configure(test.cfg)
			results := sel.Reduce(ctx, deets, fault.New(true))
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *SharePointSelectorSuite) TestSharePointCategory_PathValues() {
	var (
		itemName   = "item"
		itemID     = "item-id"
		shortRef   = "short"
		driveElems = []string{"drive", "drive!id", "root:.d", "dir1.d", "dir2.d", itemID}
		elems      = []string{"dir1", "dir2", itemID}
	)

	table := []struct {
		name       string
		sc         sharePointCategory
		pathElems  []string
		locRef     string
		parentPath string
		expected   map[categorizer][]string
		cfg        Config
	}{
		{
			name:       "SharePoint Libraries",
			sc:         SharePointLibraryItem,
			pathElems:  driveElems,
			locRef:     "root:/dir1/dir2",
			parentPath: "dir1/dir2",
			expected: map[categorizer][]string{
				SharePointLibraryFolder: {"dir1/dir2"},
				SharePointLibraryItem:   {itemID, shortRef},
			},
			cfg: Config{},
		},
		{
			name:       "SharePoint Libraries w/ name",
			sc:         SharePointLibraryItem,
			pathElems:  driveElems,
			locRef:     "root:/dir1/dir2",
			parentPath: "dir1/dir2",
			expected: map[categorizer][]string{
				SharePointLibraryFolder: {"dir1/dir2"},
				SharePointLibraryItem:   {itemName, shortRef},
			},
			cfg: Config{OnlyMatchItemNames: true},
		},
		{
			name:      "SharePoint Lists",
			sc:        SharePointListItem,
			pathElems: elems,
			locRef:    "dir1/dir2",
			expected: map[categorizer][]string{
				SharePointList:     {"dir1/dir2"},
				SharePointListItem: {itemID, shortRef},
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
				path.SharePointService,
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
					SharePoint: &details.SharePointInfo{
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

func (suite *SharePointSelectorSuite) TestSharePointScope_MatchesInfo() {
	var (
		sel          = NewSharePointRestore(Any())
		host         = "www.website.com"
		pth          = "/foo"
		url          = host + pth
		epoch        = time.Time{}
		now          = time.Now()
		modification = now.Add(15 * time.Minute)
		future       = now.Add(45 * time.Minute)
	)

	table := []struct {
		name    string
		infoURL string
		scope   []SharePointScope
		expect  assert.BoolAssertionFunc
	}{
		{"host match", host, sel.WebURL([]string{host}), assert.True},
		{"url match", url, sel.WebURL([]string{url}), assert.True},
		{"host suffixes host", host, sel.WebURL([]string{host}, SuffixMatch()), assert.True},
		{"url does not suffix host", url, sel.WebURL([]string{host}, SuffixMatch()), assert.False},
		{"url has path suffix", url, sel.WebURL([]string{pth}, SuffixMatch()), assert.True},
		{"host does not contain substring", host, sel.WebURL([]string{"website"}), assert.False},
		{"url does not suffix substring", url, sel.WebURL([]string{"oo"}, SuffixMatch()), assert.False},
		{"host mismatch", host, sel.WebURL([]string{"www.google.com"}), assert.False},
		{"file create after the epoch", host, sel.CreatedAfter(dttm.Format(epoch)), assert.True},
		{"file create after now", host, sel.CreatedAfter(dttm.Format(now)), assert.False},
		{"file create after later", url, sel.CreatedAfter(dttm.Format(future)), assert.False},
		{"file create before future", host, sel.CreatedBefore(dttm.Format(future)), assert.True},
		{"file create before now", host, sel.CreatedBefore(dttm.Format(now)), assert.False},
		{"file create before modification", host, sel.CreatedBefore(dttm.Format(modification)), assert.True},
		{"file create before epoch", host, sel.CreatedBefore(dttm.Format(now)), assert.False},
		{"file modified after the epoch", host, sel.ModifiedAfter(dttm.Format(epoch)), assert.True},
		{"file modified after now", host, sel.ModifiedAfter(dttm.Format(now)), assert.True},
		{"file modified after later", host, sel.ModifiedAfter(dttm.Format(future)), assert.False},
		{"file modified before future", host, sel.ModifiedBefore(dttm.Format(future)), assert.True},
		{"file modified before now", host, sel.ModifiedBefore(dttm.Format(now)), assert.False},
		{"file modified before epoch", host, sel.ModifiedBefore(dttm.Format(now)), assert.False},
		{"in library", host, sel.Library("included-library"), assert.True},
		{"not in library", host, sel.Library("not-included-library"), assert.False},
		{"library id", host, sel.Library("1234"), assert.True},
		{"not library id", host, sel.Library("abcd"), assert.False},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			itemInfo := details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:  details.SharePointPage,
					WebURL:    test.infoURL,
					Created:   now,
					Modified:  modification,
					DriveName: "included-library",
					DriveID:   "1234",
				},
			}

			scopes := setScopesToDefault(test.scope)
			for _, scope := range scopes {
				test.expect(t, scope.matchesInfo(itemInfo))
			}
		})
	}
}

func (suite *SharePointSelectorSuite) TestCategory_PathType() {
	table := []struct {
		cat      sharePointCategory
		pathType path.CategoryType
	}{
		{SharePointCategoryUnknown, path.UnknownCategory},
		{SharePointWebURL, path.UnknownCategory},
		{SharePointSite, path.UnknownCategory},
		{SharePointLibraryFolder, path.LibrariesCategory},
		{SharePointLibraryItem, path.LibrariesCategory},
		{SharePointList, path.ListsCategory},
	}
	for _, test := range table {
		suite.Run(test.cat.String(), func() {
			assert.Equal(suite.T(), test.pathType, test.cat.PathType())
		})
	}
}
