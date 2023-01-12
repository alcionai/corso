package selectors

import (
	"context"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// Selectors
// ---------------------------------------------------------------------------

type (
	// sharePoint provides an api for selecting
	// data scopes applicable to the SharePoint service.
	sharePoint struct {
		Selector
	}

	// SharePointBackup provides an api for selecting
	// data scopes applicable to the SharePoint service,
	// plus backup-specific methods.
	SharePointBackup struct {
		sharePoint
	}

	// SharePointRestorep provides an api for selecting
	// data scopes applicable to the SharePoint service,
	// plus restore-specific methods.
	SharePointRestore struct {
		sharePoint
	}
)

var (
	_ Reducer        = &SharePointRestore{}
	_ pathCategorier = &SharePointRestore{}
)

// NewSharePointBackup produces a new Selector with the service set to ServiceSharePoint.
func NewSharePointBackup(sites []string) *SharePointBackup {
	src := SharePointBackup{
		sharePoint{
			newSelector(ServiceSharePoint, sites),
		},
	}

	return &src
}

// ToSharePointBackup transforms the generic selector into an SharePointBackup.
// Errors if the service defined by the selector is not ServiceSharePoint.
func (s Selector) ToSharePointBackup() (*SharePointBackup, error) {
	if s.Service != ServiceSharePoint {
		return nil, badCastErr(ServiceSharePoint, s.Service)
	}

	src := SharePointBackup{sharePoint{s}}

	return &src, nil
}

func (s SharePointBackup) SplitByResourceOwner(sites []string) []SharePointBackup {
	sels := splitByResourceOwner[ExchangeScope](s.Selector, sites, SharePointSite)

	ss := make([]SharePointBackup, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, SharePointBackup{sharePoint{sel}})
	}

	return ss
}

// NewSharePointRestore produces a new Selector with the service set to ServiceSharePoint.
func NewSharePointRestore(sites []string) *SharePointRestore {
	src := SharePointRestore{
		sharePoint{
			newSelector(ServiceSharePoint, sites),
		},
	}

	return &src
}

// ToSharePointRestore transforms the generic selector into an SharePointRestore.
// Errors if the service defined by the selector is not ServiceSharePoint.
func (s Selector) ToSharePointRestore() (*SharePointRestore, error) {
	if s.Service != ServiceSharePoint {
		return nil, badCastErr(ServiceSharePoint, s.Service)
	}

	src := SharePointRestore{sharePoint{s}}

	return &src, nil
}

func (s SharePointRestore) SplitByResourceOwner(users []string) []SharePointRestore {
	sels := splitByResourceOwner[ExchangeScope](s.Selector, users, ExchangeUser)

	ss := make([]SharePointRestore, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, SharePointRestore{sharePoint{sel}})
	}

	return ss
}

// PathCategories produces the aggregation of discrete users described by each type of scope.
func (s sharePoint) PathCategories() selectorPathCategories {
	return selectorPathCategories{
		Excludes: pathCategoriesIn[SharePointScope, sharePointCategory](s.Excludes),
		Filters:  pathCategoriesIn[SharePointScope, sharePointCategory](s.Filters),
		Includes: pathCategoriesIn[SharePointScope, sharePointCategory](s.Includes),
	}
}

// -------------------
// Scope Factories

// Include appends the provided scopes to the selector's inclusion set.
// Data is included if it matches ANY inclusion.
// The inclusion set is later filtered (all included data must pass ALL
// filters) and excluded (all included data must not match ANY exclusion).
// Data is included if it matches ANY inclusion (of the same data category).
//
// All parts of the scope must match for data to be exclucded.
// Ex: File(s1, f1, i1) => only excludes an item if it is owned by site s1,
// located in folder f1, and ID'd as i1.  Use selectors.Any() to wildcard
// a scope value. No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: Site(u1) automatically cascades to all folders and files owned
// by s1.
func (s *sharePoint) Include(scopes ...[]SharePointScope) {
	s.Includes = appendScopes(s.Includes, scopes...)
}

// Exclude appends the provided scopes to the selector's exclusion set.
// Every Exclusion scope applies globally, affecting all inclusion scopes.
// Data is excluded if it matches ANY exclusion.
//
// All parts of the scope must match for data to be exclucded.
// Ex: File(s1, f1, i1) => only excludes an item if it is owned by site s1,
// located in folder f1, and ID'd as i1.  Use selectors.Any() to wildcard
// a scope value. No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: Site(u1) automatically cascades to all folders and files owned
// by s1.
func (s *sharePoint) Exclude(scopes ...[]SharePointScope) {
	s.Excludes = appendScopes(s.Excludes, scopes...)
}

// Filter appends the provided scopes to the selector's filters set.
// A selector with >0 filters and 0 inclusions will include any data
// that passes all filters.
// A selector with >0 filters and >0 inclusions will reduce the
// inclusion set to only the data that passes all filters.
// Data is retained if it passes ALL filters.
//
// All parts of the scope must match for data to be exclucded.
// Ex: File(s1, f1, i1) => only excludes an item if it is owned by site s1,
// located in folder f1, and ID'd as i1.  Use selectors.Any() to wildcard
// a scope value. No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: Site(u1) automatically cascades to all folders and files owned
// by s1.
func (s *sharePoint) Filter(scopes ...[]SharePointScope) {
	s.Filters = appendScopes(s.Filters, scopes...)
}

// Scopes retrieves the list of sharePointScopes in the selector.
func (s *sharePoint) Scopes() []SharePointScope {
	return scopes[SharePointScope](s.Selector)
}

// -------------------
// Scope Factories

// Produces one or more SharePoint webURL scopes.
// One scope is created per webURL entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *SharePointRestore) WebURL(urlSuffixes []string, opts ...option) []SharePointScope {
	scopes := []SharePointScope{}

	scopes = append(
		scopes,
		makeFilterScope[SharePointScope](
			SharePointLibraryItem,
			SharePointWebURL,
			urlSuffixes,
			pathFilterFactory(opts...)),
		makeFilterScope[SharePointScope](
			SharePointListItem,
			SharePointWebURL,
			urlSuffixes,
			pathFilterFactory(opts...)),
	)

	return scopes
}

// Produces one or more SharePoint site scopes.
// One scope is created per site entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *sharePoint) AllData() []SharePointScope {
	scopes := []SharePointScope{}

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointLibrary, Any()),
		makeScope[SharePointScope](SharePointList, Any()),
	)

	return scopes
}

// Lists produces one or more SharePoint list scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// Any empty slice defaults to [selectors.None]
func (s *sharePoint) Lists(lists []string, opts ...option) []SharePointScope {
	var (
		scopes = []SharePointScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(scopes, makeScope[SharePointScope](SharePointList, lists, os...))

	return scopes
}

// ListItems produces one or more SharePoint list item scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the list scopes.
func (s *sharePoint) ListItems(lists, items []string, opts ...option) []SharePointScope {
	scopes := []SharePointScope{}

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointListItem, items).
			set(SharePointList, lists, opts...),
	)

	return scopes
}

// Libraries produces one or more SharePoint library scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *sharePoint) Libraries(libraries []string, opts ...option) []SharePointScope {
	var (
		scopes = []SharePointScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointLibrary, libraries, os...),
	)

	return scopes
}

// LibraryItems produces one or more SharePoint library item scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the library scopes.
func (s *sharePoint) LibraryItems(libraries, items []string, opts ...option) []SharePointScope {
	scopes := []SharePointScope{}

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointLibraryItem, items).
			set(SharePointLibrary, libraries, opts...),
	)

	return scopes
}

// -------------------
// Filter Factories

// ---------------------------------------------------------------------------
// Categories
// ---------------------------------------------------------------------------

// sharePointCategory enumerates the type of the lowest level
// of data () in a scope.
type sharePointCategory string

// interface compliance checks
var _ categorizer = SharePointCategoryUnknown

const (
	SharePointCategoryUnknown sharePointCategory = ""

	// types of data identified by SharePoint
	SharePointWebURL      sharePointCategory = "SharePointWebURL"
	SharePointSite        sharePointCategory = "SharePointSite"
	SharePointList        sharePointCategory = "SharePointList"
	SharePointListItem    sharePointCategory = "SharePointListItem"
	SharePointLibrary     sharePointCategory = "SharePointLibrary"
	SharePointLibraryItem sharePointCategory = "SharePointLibraryItem"

	// filterable topics identified by SharePoint
)

// sharePointLeafProperties describes common metadata of the leaf categories
var sharePointLeafProperties = map[categorizer]leafProperty{
	SharePointLibraryItem: {
		pathKeys: []categorizer{SharePointLibrary, SharePointLibraryItem},
		pathType: path.LibrariesCategory,
	},
	SharePointSite: { // the root category must be represented, even though it isn't a leaf
		pathKeys: []categorizer{SharePointSite},
		pathType: path.UnknownCategory,
	},
	SharePointListItem: {
		pathKeys: []categorizer{SharePointSite, SharePointList, SharePointListItem},
		pathType: path.ListsCategory,
	},
}

func (c sharePointCategory) String() string {
	return string(c)
}

// leafCat returns the leaf category of the receiver.
// If the receiver category has multiple leaves (ex: User) or no leaves,
// (ex: Unknown), the receiver itself is returned.
// Ex: ServiceTypeFolder.leafCat() => ServiceTypeItem
// Ex: ServiceUser.leafCat() => ServiceUser
func (c sharePointCategory) leafCat() categorizer {
	switch c {
	case SharePointLibrary, SharePointLibraryItem:
		return SharePointLibraryItem
	case SharePointList, SharePointListItem:
		return SharePointListItem
	}

	return c
}

// rootCat returns the root category type.
func (c sharePointCategory) rootCat() categorizer {
	return SharePointSite
}

// unknownCat returns the unknown category type.
func (c sharePointCategory) unknownCat() categorizer {
	return SharePointCategoryUnknown
}

// isUnion returns true if the category is a site or a webURL, which
// can act as an alternative identifier to siteID across all site types.
func (c sharePointCategory) isUnion() bool {
	return c == SharePointWebURL || c == c.rootCat()
}

// isLeaf is true if the category is a SharePointItem category.
func (c sharePointCategory) isLeaf() bool {
	return c == c.leafCat()
}

// pathValues transforms a path to a map of identified properties.
//
// Example:
// [tenantID, service, siteID, category, folder, itemID]
// => {spSite: siteID, spFolder: folder, spItemID: itemID}
func (c sharePointCategory) pathValues(p path.Path) map[categorizer]string {
	var folderCat, itemCat categorizer

	switch c {
	case SharePointLibrary, SharePointLibraryItem:
		folderCat, itemCat = SharePointLibrary, SharePointLibraryItem
	case SharePointList, SharePointListItem:
		folderCat, itemCat = SharePointList, SharePointListItem
	}

	return map[categorizer]string{
		folderCat: p.Folder(),
		itemCat:   p.Item(),
	}
}

// pathKeys returns the path keys recognized by the receiver's leaf type.
func (c sharePointCategory) pathKeys() []categorizer {
	return sharePointLeafProperties[c.leafCat()].pathKeys
}

// PathType converts the category's leaf type into the matching path.CategoryType.
func (c sharePointCategory) PathType() path.CategoryType {
	return sharePointLeafProperties[c.leafCat()].pathType
}

// ---------------------------------------------------------------------------
// Scopes
// ---------------------------------------------------------------------------

// SharePointScope specifies the data available
// when interfacing with the SharePoint service.
type SharePointScope scope

// interface compliance checks
var _ scoper = &SharePointScope{}

// Category describes the type of the data in scope.
func (s SharePointScope) Category() sharePointCategory {
	return sharePointCategory(getCategory(s))
}

// categorizer type is a generic wrapper around Category.
// Primarily used by scopes.go to for abstract comparisons.
func (s SharePointScope) categorizer() categorizer {
	return s.Category()
}

// FilterCategory returns the category enum of the scope filter.
// If the scope is not a filter type, returns SharePointUnknownCategory.
func (s SharePointScope) FilterCategory() sharePointCategory {
	return sharePointCategory(getFilterCategory(s))
}

// IncludeCategory checks whether the scope includes a
// certain category of data.
// Ex: to check if the scope includes file data:
// s.IncludesCategory(selector.SharePointFile)
func (s SharePointScope) IncludesCategory(cat sharePointCategory) bool {
	return categoryMatches(s.Category(), cat)
}

// Matches returns true if the category is included in the scope's
// data type, and the target string matches that category's comparator.
func (s SharePointScope) Matches(cat sharePointCategory, target string) bool {
	return matches(s, cat, target)
}

// returns true if the category is included in the scope's data type,
// and the value is set to Any().
func (s SharePointScope) IsAny(cat sharePointCategory) bool {
	return isAnyTarget(s, cat)
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// SharePointUser category.
func (s SharePointScope) Get(cat sharePointCategory) []string {
	return getCatValue(s, cat)
}

// sets a value by category to the scope.  Only intended for internal use.
func (s SharePointScope) set(cat sharePointCategory, v []string, opts ...option) SharePointScope {
	os := []option{}

	switch cat {
	case SharePointLibrary, SharePointList:
		os = append(os, pathComparator())
	}

	return set(s, cat, v, append(os, opts...)...)
}

// setDefaults ensures that site scopes express `AnyTgt` for their child category types.
func (s SharePointScope) setDefaults() {
	switch s.Category() {
	case SharePointSite:
		s[SharePointLibrary.String()] = passAny
		s[SharePointLibraryItem.String()] = passAny
		s[SharePointList.String()] = passAny
		s[SharePointListItem.String()] = passAny
	case SharePointLibrary:
		s[SharePointLibraryItem.String()] = passAny
	case SharePointList:
		s[SharePointListItem.String()] = passAny
	}
}

// DiscreteCopy makes a shallow clone of the scope, then replaces the clone's
// site comparison with only the provided site.
func (s SharePointScope) DiscreteCopy(site string) SharePointScope {
	return discreteCopy(s, site)
}

// ---------------------------------------------------------------------------
// Backup Details Filtering
// ---------------------------------------------------------------------------

// Reduce filters the entries in a details struct to only those that match the
// inclusions, filters, and exclusions in the selector.
func (s sharePoint) Reduce(ctx context.Context, deets *details.Details) *details.Details {
	return reduce[SharePointScope](
		ctx,
		deets,
		s.Selector,
		map[path.CategoryType]sharePointCategory{
			path.LibrariesCategory: SharePointLibraryItem,
			path.ListsCategory:     SharePointListItem,
		},
	)
}

// matchesInfo handles the standard behavior when comparing a scope and an sharePointInfo
// returns true if the scope and info match for the provided category.
func (s SharePointScope) matchesInfo(dii details.ItemInfo) bool {
	var (
		filterCat = s.FilterCategory()
		i         = ""
		info      = dii.SharePoint
	)

	if info == nil {
		return false
	}

	switch filterCat {
	case SharePointWebURL:
		i = info.WebURL
	}

	return s.Matches(filterCat, i)
}
