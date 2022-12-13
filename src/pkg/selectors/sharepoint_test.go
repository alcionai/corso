package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

type SharePointSelectorSuite struct {
	suite.Suite
}

func TestSharePointSelectorSuite(t *testing.T) {
	suite.Run(t, new(SharePointSelectorSuite))
}

func (suite *SharePointSelectorSuite) TestNewSharePointBackup() {
	t := suite.T()
	ob := NewSharePointBackup()
	assert.Equal(t, ob.Service, ServiceSharePoint)
	assert.NotZero(t, ob.Scopes())
}

func (suite *SharePointSelectorSuite) TestToSharePointBackup() {
	t := suite.T()
	ob := NewSharePointBackup()
	s := ob.Selector
	ob, err := s.ToSharePointBackup()
	require.NoError(t, err)
	assert.Equal(t, ob.Service, ServiceSharePoint)
	assert.NotZero(t, ob.Scopes())
}

func (suite *SharePointSelectorSuite) TestSharePointBackup_DiscreteScopes() {
	sites := []string{"s1", "s2"}
	table := []struct {
		name     string
		include  []string
		discrete []string
		expect   []string
	}{
		{
			name:     "any site",
			include:  Any(),
			discrete: sites,
			expect:   sites,
		},
		{
			name:     "discrete sitet",
			include:  []string{"s3"},
			discrete: sites,
			expect:   []string{"s3"},
		},
		{
			name:     "nil discrete slice",
			include:  Any(),
			discrete: nil,
			expect:   Any(),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			eb := NewSharePointBackup()
			eb.Include(eb.Sites(test.include))

			scopes := eb.DiscreteScopes(test.discrete)
			for _, sc := range scopes {
				sites := sc.Get(SharePointSite)
				assert.Equal(t, test.expect, sites)
			}
		})
	}
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_Sites() {
	t := suite.T()
	sel := NewSharePointBackup()

	const (
		s1 = "s1"
		s2 = "s2"
	)

	siteScopes := sel.Sites([]string{s1, s2})
	for _, scope := range siteScopes {
		// Scope value is either s1 or s2
		assert.Contains(t, join(s1, s2), scope[SharePointSite.String()].Target)
	}

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
		{"Filter Scopes", sel.Filters},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			require.Len(t, test.scopesToCheck, 2)
			for _, scope := range test.scopesToCheck {
				// Scope value is s1,s2
				assert.Contains(t, join(s1, s2), scope[SharePointSite.String()].Target)
			}
		})
	}
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_Include_WebURLs() {
	t := suite.T()
	sel := NewSharePointRestore()

	const (
		s1 = "s1"
		s2 = "s2"
	)

	sel.Include(sel.WebURL([]string{s1, s2}))
	scopes := sel.Includes
	require.Len(t, scopes, 2)

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
		suite.T().Run(test.name, func(t *testing.T) {
			sel := NewSharePointRestore()
			sel.Include(sel.WebURL(test.in))
			scopes := sel.Includes
			require.Len(t, scopes, 2)

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
	sel := NewSharePointRestore()

	const (
		s1 = "s1"
		s2 = "s2"
	)

	sel.Exclude(sel.WebURL([]string{s1, s2}))
	scopes := sel.Excludes
	// @keepers: Question Should this be 2 as in check for SharePoint Lists and SharePoint Drives?
	require.Len(t, scopes, 2)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			SharePointScope(sc),
			map[categorizer]string{SharePointWebURL: join(s1, s2)},
		)
	}
}

// TestSharePointselector_Include_Sites ensures that the scopes of
// SharePoint Libraries & SharePoint Lists are created.
func (suite *SharePointSelectorSuite) TestSharePointSelector_Include_Sites() {
	t := suite.T()
	sel := NewSharePointBackup()

	const (
		s1 = "s1"
		s2 = "s2"
	)

	sel.Include(sel.Sites([]string{s1, s2}))
	scopes := sel.Includes
	require.Len(t, scopes, 2)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			SharePointScope(sc),
			map[categorizer]string{SharePointSite: join(s1, s2)},
		)
	}
}

func (suite *SharePointSelectorSuite) TestSharePointSelector_Exclude_Sites() {
	t := suite.T()
	sel := NewSharePointBackup()

	const (
		s1 = "s1"
		s2 = "s2"
	)

	sel.Exclude(sel.Sites([]string{s1, s2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 2)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			SharePointScope(sc),
			map[categorizer]string{SharePointSite: join(s1, s2)},
		)
	}
}

func (suite *SharePointSelectorSuite) TestNewSharePointRestore() {
	t := suite.T()
	or := NewSharePointRestore()
	assert.Equal(t, or.Service, ServiceSharePoint)
	assert.NotZero(t, or.Scopes())
}

func (suite *SharePointSelectorSuite) TestToSharePointRestore() {
	t := suite.T()
	eb := NewSharePointRestore()
	s := eb.Selector
	or, err := s.ToSharePointRestore()
	require.NoError(t, err)
	assert.Equal(t, or.Service, ServiceSharePoint)
	assert.NotZero(t, or.Scopes())
}

func (suite *SharePointSelectorSuite) TestSharePointRestore_Reduce() {
	var (
		item  = stubRepoRef(path.SharePointService, path.LibrariesCategory, "uid", "/folderA/folderB", "item")
		item2 = stubRepoRef(path.SharePointService, path.LibrariesCategory, "uid", "/folderA/folderC", "item2")
		item3 = stubRepoRef(path.SharePointService, path.LibrariesCategory, "uid", "/folderD/folderE", "item3")
	)

	deets := &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: []details.DetailsEntry{
				{
					RepoRef: item,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointItem,
						},
					},
				},
				{
					RepoRef: item2,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointItem,
						},
					},
				},
				{
					RepoRef: item3,
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointItem,
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
			"all",
			deets,
			func() *SharePointRestore {
				odr := NewSharePointRestore()
				odr.Include(odr.Sites(Any()))
				return odr
			},
			arr(item, item2, item3),
		},
		{
			"only match item",
			deets,
			func() *SharePointRestore {
				odr := NewSharePointRestore()
				odr.Include(odr.LibraryItems(Any(), Any(), []string{"item2"}))
				return odr
			},
			arr(item2),
		},
		{
			"only match folder",
			deets,
			func() *SharePointRestore {
				odr := NewSharePointRestore()
				odr.Include(odr.Libraries([]string{"uid"}, []string{"folderA/folderB", "folderA/folderC"}))
				return odr
			},
			arr(item, item2),
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			sel := test.makeSelector()
			results := sel.Reduce(ctx, test.deets)
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *SharePointSelectorSuite) TestSharePointCategory_PathValues() {
	t := suite.T()

	pathBuilder := path.Builder{}.Append("dir1", "dir2", "item")
	itemPath, err := pathBuilder.ToDataLayerSharePointPath("tenant", "site", path.LibrariesCategory, true)
	require.NoError(t, err)

	expected := map[categorizer]string{
		SharePointSite:        "site",
		SharePointLibrary:     "dir1/dir2",
		SharePointLibraryItem: "item",
	}

	assert.Equal(t, expected, SharePointLibraryItem.pathValues(itemPath))
}

func (suite *SharePointSelectorSuite) TestSharePointScope_MatchesInfo() {
	var (
		ods  = NewSharePointRestore()
		host = "www.website.com"
		pth  = "/foo"
		url  = host + pth
	)

	table := []struct {
		name    string
		infoURL string
		scope   []SharePointScope
		expect  assert.BoolAssertionFunc
	}{
		{"host match", host, ods.WebURL([]string{host}), assert.True},
		{"url match", url, ods.WebURL([]string{url}), assert.True},
		{"url contains host", url, ods.WebURL([]string{host}), assert.True},
		{"host suffixes host", host, ods.WebURL([]string{host}, SuffixMatch()), assert.True},
		{"url does not suffix host", url, ods.WebURL([]string{host}, SuffixMatch()), assert.False},
		{"url contains path", url, ods.WebURL([]string{pth}), assert.True},
		{"url has path suffix", url, ods.WebURL([]string{pth}, SuffixMatch()), assert.True},
		{"host does not contain substring", host, ods.WebURL([]string{"website"}), assert.False},
		{"url does not suffix substring", url, ods.WebURL([]string{"oo"}), assert.False},
		{"host mismatch", host, ods.WebURL([]string{"www.google.com"}), assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			itemInfo := details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType: details.SharePointItem,
					WebURL:   test.infoURL,
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
		{SharePointLibrary, path.LibrariesCategory},
		{SharePointLibraryItem, path.LibrariesCategory},
		{SharePointList, path.ListsCategory},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, test.pathType, test.cat.PathType())
		})
	}
}
