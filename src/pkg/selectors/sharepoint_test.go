package selectors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common"
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

			suite.T().Run(test.name+"-"+cat.String(), func(t *testing.T) {
				switch cat {
				case SharePointLibraryItem:
					scopeMustHave(
						t,
						spsc,
						map[categorizer]string{
							SharePointLibraryItem:   AnyTgt,
							SharePointLibraryFolder: AnyTgt,
						},
					)
				case SharePointListItem:
					scopeMustHave(
						t,
						spsc,
						map[categorizer]string{
							SharePointListItem: AnyTgt,
							SharePointList:     AnyTgt,
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

	sel := NewSharePointRestore([]string{s1, s2})
	sel.Include(sel.WebURL([]string{s1, s2}))
	scopes := sel.Includes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			SharePointScope(sc),
			map[categorizer]string{SharePointWebURL: join(s1, s2)},
		)
	}
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_Include_WebURLs_anyNone() {
	table := []struct {
		name   string
		in     []string
		expect string
	}{
		{
			name:   "any",
			in:     []string{AnyTgt},
			expect: AnyTgt,
		},
		{
			name:   "none",
			in:     []string{NoneTgt},
			expect: NoneTgt,
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
					map[categorizer]string{SharePointWebURL: test.expect},
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

	sel := NewSharePointRestore([]string{s1, s2})
	sel.Exclude(sel.WebURL([]string{s1, s2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			SharePointScope(sc),
			map[categorizer]string{SharePointWebURL: join(s1, s2)},
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
	var (
		pairAC = "folderA/folderC"
		pairGH = "folderG/folderH"
		item   = stubRepoRef(path.SharePointService, path.LibrariesCategory, "sid", "folderA/folderB", "item")
		item2  = stubRepoRef(path.SharePointService, path.LibrariesCategory, "sid", pairAC, "item2")
		item3  = stubRepoRef(path.SharePointService, path.LibrariesCategory, "sid", "folderD/folderE", "item3")
		item4  = stubRepoRef(path.SharePointService, path.PagesCategory, "sid", pairGH, "item4")
		item5  = stubRepoRef(path.SharePointService, path.PagesCategory, "sid", pairGH, "item5")
	)

	deets := &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: []details.DetailsEntry{
				{
					RepoRef: item,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointLibrary,
							ItemName: "itemName",
						},
					},
				},
				{
					RepoRef: item2,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointLibrary,
							ItemName: "itemName2",
						},
					},
				},
				{
					RepoRef: item3,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointLibrary,
							ItemName: "itemName3",
						},
					},
				},
				{
					RepoRef: item4,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointPage,
							ItemName: "itemName4",
						},
					},
				},
				{
					RepoRef: item5,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointPage,
							ItemName: "itemName5",
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
		deets        *details.Details
		makeSelector func() *SharePointRestore
		expect       []string
	}{
		{
			name:  "all",
			deets: deets,
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore(Any())
				odr.Include(odr.AllData())
				return odr
			},
			expect: arr(item, item2, item3, item4, item5),
		},
		{
			name:  "only match item",
			deets: deets,
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore(Any())
				odr.Include(odr.LibraryItems(Any(), []string{"itemName2"}))
				return odr
			},
			expect: arr(item2),
		},
		{
			name:  "only match folder",
			deets: deets,
			makeSelector: func() *SharePointRestore {
				odr := NewSharePointRestore([]string{"sid"})
				odr.Include(odr.LibraryFolders([]string{"folderA/folderB", pairAC}))
				return odr
			},
			expect: arr(item, item2),
		},
		{
			name:  "pages match folder",
			deets: deets,
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
			results := sel.Reduce(ctx, test.deets, fault.New(true))
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *SharePointSelectorSuite) TestSharePointCategory_PathValues() {
	itemName := "item"
	shortRef := "short"
	elems := []string{"dir1", "dir2", itemName + "-id"}

	table := []struct {
		name     string
		sc       sharePointCategory
		expected map[categorizer][]string
	}{
		{
			name: "SharePoint Libraries",
			sc:   SharePointLibraryItem,
			expected: map[categorizer][]string{
				SharePointLibraryFolder: {"dir1/dir2"},
				SharePointLibraryItem:   {itemName, shortRef},
			},
		},
		{
			name: "SharePoint Lists",
			sc:   SharePointListItem,
			expected: map[categorizer][]string{
				SharePointList:     {"dir1/dir2"},
				SharePointListItem: {"item-id", shortRef},
			},
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
				elems...)
			require.NoError(t, err, clues.ToCore(err))

			ent := details.DetailsEntry{
				RepoRef:  itemPath.String(),
				ShortRef: shortRef,
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemName: itemName,
					},
				},
			}

			pv, err := test.sc.pathValues(itemPath, ent)
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
		{"url contains host", url, sel.WebURL([]string{host}), assert.True},
		{"host suffixes host", host, sel.WebURL([]string{host}, SuffixMatch()), assert.True},
		{"url does not suffix host", url, sel.WebURL([]string{host}, SuffixMatch()), assert.False},
		{"url contains path", url, sel.WebURL([]string{pth}), assert.True},
		{"url has path suffix", url, sel.WebURL([]string{pth}, SuffixMatch()), assert.True},
		{"host does not contain substring", host, sel.WebURL([]string{"website"}), assert.False},
		{"url does not suffix substring", url, sel.WebURL([]string{"oo"}), assert.False},
		{"host mismatch", host, sel.WebURL([]string{"www.google.com"}), assert.False},
		{"file create after the epoch", host, sel.CreatedAfter(common.FormatTime(epoch)), assert.True},
		{"file create after now", host, sel.CreatedAfter(common.FormatTime(now)), assert.False},
		{"file create after later", url, sel.CreatedAfter(common.FormatTime(future)), assert.False},
		{"file create before future", host, sel.CreatedBefore(common.FormatTime(future)), assert.True},
		{"file create before now", host, sel.CreatedBefore(common.FormatTime(now)), assert.False},
		{"file create before modification", host, sel.CreatedBefore(common.FormatTime(modification)), assert.True},
		{"file create before epoch", host, sel.CreatedBefore(common.FormatTime(now)), assert.False},
		{"file modified after the epoch", host, sel.ModifiedAfter(common.FormatTime(epoch)), assert.True},
		{"file modified after now", host, sel.ModifiedAfter(common.FormatTime(now)), assert.True},
		{"file modified after later", host, sel.ModifiedAfter(common.FormatTime(future)), assert.False},
		{"file modified before future", host, sel.ModifiedBefore(common.FormatTime(future)), assert.True},
		{"file modified before now", host, sel.ModifiedBefore(common.FormatTime(now)), assert.False},
		{"file modified before epoch", host, sel.ModifiedBefore(common.FormatTime(now)), assert.False},
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
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, test.pathType, test.cat.PathType())
		})
	}
}
