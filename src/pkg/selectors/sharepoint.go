package selectors

import (
	"context"

	"github.com/alcionai/corso/src/pkg/backup/details"
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

var _ Reducer = &SharePointRestore{}

// NewSharePointBackup produces a new Selector with the service set to ServiceSharePoint.
func NewSharePointBackup() *SharePointBackup {
	src := SharePointBackup{
		sharePoint{
			newSelector(ServiceSharePoint),
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

// NewSharePointRestore produces a new Selector with the service set to ServiceSharePoint.
func NewSharePointRestore() *SharePointRestore {
	src := SharePointRestore{
		sharePoint{
			newSelector(ServiceSharePoint),
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

// Printable creates the minimized display of a selector, formatted for human readability.
func (s sharePoint) Printable() Printable {
	return toPrintable[SharePointScope](s.Selector)
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

// DiscreteScopes retrieves the list of sharePointScopes in the selector.
// If any Include scope's Site category is set to Any, replaces that
// scope's value with the list of siteIDs instead.
func (s *sharePoint) DiscreteScopes(siteIDs []string) []SharePointScope {
	return discreteScopes[SharePointScope](s.Selector, SharePointSite, siteIDs)
}

// -------------------
// Scope Factories

// Produces one or more SharePoint site scopes.
// One scope is created per site entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *sharePoint) Sites(sites []string) []SharePointScope {
	scopes := []SharePointScope{}

	scopes = append(scopes, makeScope[SharePointScope](SharePointFolder, sites, Any()))

	return scopes
}

// Folders produces one or more SharePoint folder scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *sharePoint) Folders(sites, folders []string, opts ...option) []SharePointScope {
	var (
		scopes = []SharePointScope{}
		os     = append([]option{pathType()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointFolder, sites, folders, os...),
	)

	return scopes
}

// Items produces one or more SharePoint item scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *sharePoint) Items(sites, folders, items []string, opts ...option) []SharePointScope {
	scopes := []SharePointScope{}

	scopes = append(
		scopes,
		makeScope[SharePointScope](SharePointItem, sites, items).
			set(SharePointFolder, folders, opts...),
	)

	return scopes
}

// -------------------
// Filter Factories

// WebURL produces a SharePoint item webURL filter scope.
// Matches any item where the webURL contains the substring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *sharePoint) WebURL(substring string) []SharePointScope {
	return []SharePointScope{
		makeFilterScope[SharePointScope](
			SharePointItem,
			SharePointFilterWebURL,
			[]string{substring},
			wrapFilter(filters.Less)),
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
	// types of data identified by SharePoint
	SharePointSite   sharePointCategory = "SharePointSite"
	SharePointFolder sharePointCategory = "SharePointFolder"
	SharePointItem   sharePointCategory = "SharePointItem"

	// filterable topics identified by SharePoint
	SharePointFilterWebURL sharePointCategory = "SharePointFilterWebURL"
)

// sharePointLeafProperties describes common metadata of the leaf categories
var sharePointLeafProperties = map[categorizer]leafProperty{
	SharePointItem: {
		pathKeys: []categorizer{SharePointSite, SharePointFolder, SharePointItem},
		pathType: path.FilesCategory,
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
	case SharePointFolder, SharePointItem,
		SharePointFilterWebURL:
		return SharePointItem
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

// isLeaf is true if the category is a SharePointItem category.
func (c sharePointCategory) isLeaf() bool {
	// return c == c.leafCat()??
	return c == SharePointItem
}

// pathValues transforms a path to a map of identified properties.
//
// Example:
// [tenantID, service, siteID, category, folder, itemID]
// => {spSite: siteID, spFolder: folder, spItemID: itemID}
func (c sharePointCategory) pathValues(p path.Path) map[categorizer]string {
	return map[categorizer]string{
		SharePointSite:   p.ResourceOwner(),
		SharePointFolder: p.Folder(),
		SharePointItem:   p.Item(),
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
	if cat == SharePointFolder {
		os = append(os, pathType())
	}

	return set(s, cat, v, append(os, opts...)...)
}

// setDefaults ensures that site scopes express `AnyTgt` for their child category types.
func (s SharePointScope) setDefaults() {
	switch s.Category() {
	case SharePointSite:
		s[SharePointFolder.String()] = passAny
		s[SharePointItem.String()] = passAny
	case SharePointFolder:
		s[SharePointItem.String()] = passAny
	}
}

// matchesInfo handles the standard behavior when comparing a scope and an sharePointInfo
// returns true if the scope and info match for the provided category.
func (s SharePointScope) matchesInfo(dii details.ItemInfo) bool {
	// info := dii.SharePoint
	// if info == nil {
	// 	return false
	// }
	var (
		filterCat = s.FilterCategory()
		i         = ""
	)

	// switch filterCat {
	// case FileFilterCreatedAfter, FileFilterCreatedBefore:
	// 	i = common.FormatTime(info.Created)
	// case FileFilterModifiedAfter, FileFilterModifiedBefore:
	// 	i = common.FormatTime(info.Modified)
	// }

	return s.Matches(filterCat, i)
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
			// TODO: need to figure out the path Category(s) for sharepoint.
			path.FilesCategory: SharePointItem,
		},
	)
}
