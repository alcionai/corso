package selectors

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
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
	sels := splitByResourceOwner[SharePointScope](s.Selector, sites, SharePointSite)

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

func (s SharePointRestore) SplitByResourceOwner(sites []string) []SharePointRestore {
	sels := splitByResourceOwner[SharePointScope](s.Selector, sites, SharePointSite)

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

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

func (s SharePointScope) Conceal() string             { return conceal(s) }
func (s SharePointScope) Format(fs fmt.State, r rune) { format(s, fs, r) }
func (s SharePointScope) String() string              { return conceal(s) }
func (s SharePointScope) PlainString() string         { return plainString(s) }

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
// Defaults to equals check, on the assumption we identify fully qualified
// urls, and do not want to default to contains.  ie: https://host/sites/foo
// should not match https://host/sites/foo/bar.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *SharePointRestore) WebURL(urls []string, opts ...option) []SharePointScope {
	var (
		scopes = []SharePointScope{}
		os     = append([]option{ExactMatch()}, opts...)
	)

	scopes = append(
		scopes,
		makeInfoScope[SharePointScope](
			SharePointLibraryItem,
			SharePointWebURL,
			urls,
			pathFilterFactory(os...)),
		makeInfoScope[SharePointScope](
			SharePointListItem,
			SharePointWebURL,
			urls,
			pathFilterFactory(os...)),
		makeInfoScope[SharePointScope](
			SharePointPage,
			SharePointWebURL,
			urls,
			pathFilterFactory(os...)),
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
		makeScope[SharePointScope](SharePointLibraryFolder, Any()),
		makeScope[SharePointScope](SharePointList, Any()),
		makeScope[SharePointScope](SharePointPageFolder, Any()),
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

// Library produces one or more SharePoint library scopes, where the library
// matches upon a given drive by ID or Name.  In order to ensure library selection
// this should always be embedded within the Filter() set; include(Library()) will
// select all items in the library without further filtering.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *sharePoint) Library(library string) []SharePointScope {
	return []SharePointScope{
		makeInfoScope[SharePointScope](
			SharePointLibraryItem,
			SharePointInfoLibraryDrive,
			[]string{library},
			filters.Equal),
	}
}

// LibraryFolders produces one or more SharePoint libraryFolder scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *sharePoint) LibraryFolders(libraryFolders []string, opts ...option) []SharePointScope {
	var (
		scopes = []SharePointScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointLibraryFolder, libraryFolders, os...),
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
			set(SharePointLibraryFolder, libraries, opts...),
	)

	return scopes
}

// Pages produces one or more SharePoint page scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *sharePoint) Pages(pages []string, opts ...option) []SharePointScope {
	var (
		scopes = []SharePointScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(scopes, makeScope[SharePointScope](SharePointPageFolder, pages, os...))

	return scopes
}

// PageItems produces one or more SharePoint page item scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the page scopes.
func (s *sharePoint) PageItems(pages, items []string, opts ...option) []SharePointScope {
	scopes := []SharePointScope{}

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointPage, items).
			set(SharePointPageFolder, pages, opts...),
	)

	return scopes
}

// -------------------
// ItemInfo Factories

func (s *sharePoint) CreatedAfter(timeStrings string) []SharePointScope {
	return []SharePointScope{
		makeInfoScope[SharePointScope](
			SharePointLibraryItem,
			SharePointInfoCreatedAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

func (s *sharePoint) CreatedBefore(timeStrings string) []SharePointScope {
	return []SharePointScope{
		makeInfoScope[SharePointScope](
			SharePointLibraryItem,
			SharePointInfoCreatedBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

func (s *sharePoint) ModifiedAfter(timeStrings string) []SharePointScope {
	return []SharePointScope{
		makeInfoScope[SharePointScope](
			SharePointLibraryItem,
			SharePointInfoModifiedAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

func (s *sharePoint) ModifiedBefore(timeStrings string) []SharePointScope {
	return []SharePointScope{
		makeInfoScope[SharePointScope](
			SharePointLibraryItem,
			SharePointInfoModifiedBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

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

	// types of data in SharePoint
	SharePointWebURL        sharePointCategory = "SharePointWebURL"
	SharePointSite          sharePointCategory = "SharePointSite"
	SharePointList          sharePointCategory = "SharePointList"
	SharePointListItem      sharePointCategory = "SharePointListItem"
	SharePointLibraryFolder sharePointCategory = "SharePointLibraryFolder"
	SharePointLibraryItem   sharePointCategory = "SharePointLibraryItem"
	SharePointPageFolder    sharePointCategory = "SharePointPageFolder"
	SharePointPage          sharePointCategory = "SharePointPage"

	// details.itemInfo comparables
	SharePointInfoCreatedAfter   sharePointCategory = "SharePointInfoCreatedAfter"
	SharePointInfoCreatedBefore  sharePointCategory = "SharePointInfoCreatedBefore"
	SharePointInfoModifiedAfter  sharePointCategory = "SharePointInfoModifiedAfter"
	SharePointInfoModifiedBefore sharePointCategory = "SharePointInfoModifiedBefore"

	// library drive selection
	SharePointInfoLibraryDrive sharePointCategory = "SharePointInfoLibraryDrive"
)

// sharePointLeafProperties describes common metadata of the leaf categories
var sharePointLeafProperties = map[categorizer]leafProperty{
	SharePointLibraryItem: {
		pathKeys: []categorizer{SharePointLibraryFolder, SharePointLibraryItem},
		pathType: path.LibrariesCategory,
	},
	SharePointListItem: {
		pathKeys: []categorizer{SharePointList, SharePointListItem},
		pathType: path.ListsCategory,
	},
	SharePointPage: {
		pathKeys: []categorizer{SharePointPageFolder, SharePointPage},
		pathType: path.PagesCategory,
	},
	SharePointSite: { // the root category must be represented, even though it isn't a leaf
		pathKeys: []categorizer{SharePointSite},
		pathType: path.UnknownCategory,
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
	case SharePointLibraryFolder, SharePointLibraryItem, SharePointInfoLibraryDrive,
		SharePointInfoCreatedAfter, SharePointInfoCreatedBefore,
		SharePointInfoModifiedAfter, SharePointInfoModifiedBefore:
		return SharePointLibraryItem
	case SharePointList, SharePointListItem:
		return SharePointListItem
	case SharePointPage, SharePointPageFolder:
		return SharePointPage
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

// pathValues transforms the two paths to maps of identified properties.
//
// Example:
// [tenantID, service, siteID, category, folder, itemID]
// => {spFolder: folder, spItemID: itemID}
func (c sharePointCategory) pathValues(
	repo path.Path,
	ent details.DetailsEntry,
	cfg Config,
) (map[categorizer][]string, error) {
	var (
		folderCat, itemCat    categorizer
		dropDriveFolderPrefix bool
		itemID                string
	)

	switch c {
	case SharePointLibraryFolder, SharePointLibraryItem:
		if ent.SharePoint == nil {
			return nil, clues.New("no SharePoint ItemInfo in details")
		}

		dropDriveFolderPrefix = true
		folderCat, itemCat = SharePointLibraryFolder, SharePointLibraryItem

	case SharePointList, SharePointListItem:
		folderCat, itemCat = SharePointList, SharePointListItem

	case SharePointPage, SharePointPageFolder:
		folderCat, itemCat = SharePointPageFolder, SharePointPage

	default:
		return nil, clues.New("unrecognized sharePointCategory").With("category", c)
	}

	rFld := repo.Folder(false)
	if dropDriveFolderPrefix {
		// like onedrive, ignore `drives/<driveID>/root:` for library folder comparison
		rFld = path.Builder{}.Append(repo.Folders()...).PopFront().PopFront().PopFront().String()
	}

	item := ent.ItemRef
	if len(item) == 0 {
		item = repo.Item()
	}

	if cfg.OnlyMatchItemNames {
		item = ent.ItemInfo.SharePoint.ItemName
	}

	result := map[categorizer][]string{
		folderCat: {rFld},
		itemCat:   {item, ent.ShortRef},
	}

	if len(itemID) > 0 {
		result[itemCat] = append(result[itemCat], itemID)
	}

	if len(ent.LocationRef) > 0 {
		result[folderCat] = append(result[folderCat], ent.LocationRef)
	}

	return result, nil
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

// Matches returns true if the category is included in the scope's
// data type, and the target string matches that category's comparator.
func (s SharePointScope) Matches(cat sharePointCategory, target string) bool {
	return matches(s, cat, target)
}

// InfoCategory returns the category enum of the scope info.
// If the scope is not an info type, returns SharePointUnknownCategory.
func (s SharePointScope) InfoCategory() sharePointCategory {
	return sharePointCategory(getInfoCategory(s))
}

// IncludeCategory checks whether the scope includes a
// certain category of data.
// Ex: to check if the scope includes file data:
// s.IncludesCategory(selector.SharePointFile)
func (s SharePointScope) IncludesCategory(cat sharePointCategory) bool {
	return categoryMatches(s.Category(), cat)
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
	case SharePointLibraryFolder, SharePointList, SharePointPage:
		os = append(os, pathComparator())
	}

	return set(s, cat, v, append(os, opts...)...)
}

// setDefaults ensures that site scopes express `AnyTgt` for their child category types.
func (s SharePointScope) setDefaults() {
	switch s.Category() {
	case SharePointSite:
		s[SharePointLibraryFolder.String()] = passAny
		s[SharePointLibraryItem.String()] = passAny
		s[SharePointList.String()] = passAny
		s[SharePointListItem.String()] = passAny
		s[SharePointPageFolder.String()] = passAny
		s[SharePointPage.String()] = passAny
	case SharePointLibraryFolder:
		s[SharePointLibraryItem.String()] = passAny
	case SharePointList:
		s[SharePointListItem.String()] = passAny
	case SharePointPageFolder:
		s[SharePointPage.String()] = passAny
	}
}

// ---------------------------------------------------------------------------
// Backup Details Filtering
// ---------------------------------------------------------------------------

// Reduce filters the entries in a details struct to only those that match the
// inclusions, filters, and exclusions in the selector.
func (s sharePoint) Reduce(
	ctx context.Context,
	deets *details.Details,
	errs *fault.Bus,
) *details.Details {
	return reduce[SharePointScope](
		ctx,
		deets,
		s.Selector,
		map[path.CategoryType]sharePointCategory{
			path.LibrariesCategory: SharePointLibraryItem,
			path.ListsCategory:     SharePointListItem,
			path.PagesCategory:     SharePointPage,
		},
		errs)
}

// matchesInfo handles the standard behavior when comparing a scope and an sharePointInfo
// returns true if the scope and info match for the provided category.
func (s SharePointScope) matchesInfo(dii details.ItemInfo) bool {
	var (
		infoCat = s.InfoCategory()
		i       = ""
		info    = dii.SharePoint
	)

	if info == nil {
		return false
	}

	switch infoCat {
	case SharePointWebURL:
		i = info.WebURL
	case SharePointInfoCreatedAfter, SharePointInfoCreatedBefore:
		i = common.FormatTime(info.Created)
	case SharePointInfoModifiedAfter, SharePointInfoModifiedBefore:
		i = common.FormatTime(info.Modified)
	case SharePointInfoLibraryDrive:
		ds := []string{}

		if len(info.DriveName) > 0 {
			ds = append(ds, info.DriveName)
		}

		if len(info.DriveID) > 0 {
			ds = append(ds, info.DriveID)
		}

		return matchesAny(s, SharePointInfoLibraryDrive, ds)
	}

	return s.Matches(infoCat, i)
}
