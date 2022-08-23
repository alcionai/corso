package selectors

import (
	"github.com/alcionai/corso/pkg/backup/details"
)

// ---------------------------------------------------------------------------
// Selectors
// ---------------------------------------------------------------------------

type (
	// onedrive provides an api for selecting
	// data scopes applicable to the OneDrive service.
	onedrive struct {
		Selector
	}

	// OneDriveBackup provides an api for selecting
	// data scopes applicable to the OneDrive service,
	// plus backup-specific methods.
	OneDriveBackup struct {
		onedrive
	}
)

// NewOneDriveBackup produces a new Selector with the service set to ServiceOneDrive.
func NewOneDriveBackup() *OneDriveBackup {
	src := OneDriveBackup{
		onedrive{
			newSelector(ServiceOneDrive),
		},
	}
	return &src
}

// ToOneDriveBackup transforms the generic selector into an OneDriveBackup.
// Errors if the service defined by the selector is not ServiceOneDrive.
func (s Selector) ToOneDriveBackup() (*OneDriveBackup, error) {
	if s.Service != ServiceOneDrive {
		return nil, badCastErr(ServiceOneDrive, s.Service)
	}
	src := OneDriveBackup{onedrive{s}}
	return &src, nil
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
// Ex: File(u1, f1, m1) => only excludes a file if it is owned by user u1,
// located in folder f1, and ID'd as m1.  Use selectors.Any() to wildcard
// a scope value. No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: User(u1) automatically cascades to all folders and files owned
// by u1.
func (s *onedrive) Include(scopes ...[]OneDriveScope) {
	s.Includes = appendScopes(s.Includes, scopes...)
}

// Exclude appends the provided scopes to the selector's exclusion set.
// Every Exclusion scope applies globally, affecting all inclusion scopes.
// Data is excluded if it matches ANY exclusion.
//
// All parts of the scope must match for data to be exclucded.
// Ex: File(u1, f1, m1) => only excludes a file if it is owned by user u1,
// located in folder f1, and ID'd as m1.  Use selectors.Any() to wildcard
// a scope value. No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: User(u1) automatically cascades to all folders and files owned
// by u1.
func (s *onedrive) Exclude(scopes ...[]OneDriveScope) {
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
// Ex: File(u1, f1, m1) => only excludes a file if it is owned by user u1,
// located in folder f1, and ID'd as m1.  Use selectors.Any() to wildcard
// a scope value. No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: User(u1) automatically cascades to all folders and files owned
// by u1.
func (s *onedrive) Filter(scopes ...[]OneDriveScope) {
	s.Filters = appendScopes(s.Filters, scopes...)
}

// Produces one or more onedrive user scopes.
// One scope is created per user entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *onedrive) Users(users []string) []OneDriveScope {
	users = normalize(users)
	scopes := []OneDriveScope{}
	for _, u := range users {
		scopes = append(scopes, makeScope[OneDriveScope](u, Group, OneDriveUser, users))
	}
	return scopes
}

// Scopes retrieves the list of onedriveScopes in the selector.
func (s *onedrive) Scopes() []OneDriveScope {
	return scopes[OneDriveScope](s.Selector)
}

// DiscreteScopes retrieves the list of onedriveScopes in the selector.
// If any Include scope's User category is set to Any, replaces that
// scope's value with the list of userIDs instead.
func (s *onedrive) DiscreteScopes(userIDs []string) []OneDriveScope {
	return discreteScopes[OneDriveScope](s.Selector, OneDriveUser, userIDs)
}

// ---------------------------------------------------------------------------
// Categories
// ---------------------------------------------------------------------------

// onedriveCategory enumerates the type of the lowest level
// of data () in a scope.
type onedriveCategory int

// interface compliance checks
var _ categorizer = OneDriveCategoryUnknown

//go:generate go run golang.org/x/tools/cmd/stringer -type=onedriveCategory
const (
	OneDriveCategoryUnknown onedriveCategory = iota
	// types of data identified by OneDrive
	OneDriveUser
)

func onedriveCatAtoI(s string) onedriveCategory {
	switch s {
	// data types
	case OneDriveUser.String():
		return OneDriveUser
	// filters
	default:
		return OneDriveCategoryUnknown
	}
}

// oneDrivePathSet describes the category type keys used in OneDrive paths.
// The order of each slice is important, and should match the order in which
// these types appear in the canonical Path for each type.
var oneDrivePathSet = map[categorizer][]categorizer{
	OneDriveUser: {OneDriveUser}, // the root category must be represented
}

// leafCat returns the leaf category of the receiver.
// If the receiver category has multiple leaves (ex: User) or no leaves,
// (ex: Unknown), the receiver itself is returned.
// Ex: ServiceTypeFolder.leafCat() => ServiceTypeItem
// Ex: ServiceUser.leafCat() => ServiceUser
func (c onedriveCategory) leafCat() categorizer {
	return c
}

// rootCat returns the root category type.
func (c onedriveCategory) rootCat() categorizer {
	return OneDriveUser
}

// unknownCat returns the unknown category type.
func (c onedriveCategory) unknownCat() categorizer {
	return OneDriveCategoryUnknown
}

// pathValues transforms a path to a map of identified properties.
// TODO: this should use service-specific funcs in the Paths pkg.  Instead of
// peeking at the path directly, the caller should compare against values like
// path.UserID() and path.Folders().
//
// Malformed (ie, short len) paths will return incomplete results.
// Example:
// [tenantID, userID, "files", folder, fileID]
// => {odUser: userID, odFolder: folder, odFileID: fileID}
func (c onedriveCategory) pathValues(path []string) map[categorizer]string {
	m := map[categorizer]string{}
	if len(path) < 2 {
		return m
	}
	m[OneDriveUser] = path[1]
	/*
		TODO/Notice:
		Files contain folder structures, identified
		in this code as being at index 3.  This assumes a single
		folder, while in reality users can express subfolder
		hierarchies of arbirary depth.  Subfolder handling is coming
		at a later time.
	*/
	// TODO: populate path values when known.
	return m
}

// pathKeys returns the path keys recognized by the receiver's leaf type.
func (c onedriveCategory) pathKeys() []categorizer {
	return oneDrivePathSet[c.leafCat()]
}

// ---------------------------------------------------------------------------
// Scopes
// ---------------------------------------------------------------------------

// OneDriveScope specifies the data available
// when interfacing with the OneDrive service.
type OneDriveScope scope

// interface compliance checks
var _ scoper = &OneDriveScope{}

// Category describes the type of the data in scope.
func (s OneDriveScope) Category() onedriveCategory {
	return onedriveCatAtoI(s[scopeKeyCategory])
}

// categorizer type is a generic wrapper around Category.
// Primarily used by scopes.go to for abstract comparisons.
func (s OneDriveScope) categorizer() categorizer {
	return s.Category()
}

// FilterCategory returns the category enum of the scope filter.
// If the scope is not a filter type, returns OneDriveUnknownCategory.
func (s OneDriveScope) FilterCategory() onedriveCategory {
	return onedriveCatAtoI(s[scopeKeyInfoFilter])
}

// Granularity describes the granularity (directory || item)
// of the data in scope.
func (s OneDriveScope) Granularity() string {
	return s[scopeKeyGranularity]
}

// IncludeCategory checks whether the scope includes a
// certain category of data.
// Ex: to check if the scope includes file data:
// s.IncludesCategory(selector.OneDriveFile)
func (s OneDriveScope) IncludesCategory(cat onedriveCategory) bool {
	return categoryMatches(s.Category(), cat)
}

// Contains returns true if the category is included in the scope's
// data type, and the target string is included in the scope.
func (s OneDriveScope) Contains(cat onedriveCategory, target string) bool {
	return contains(s, cat, target)
}

// returns true if the category is included in the scope's data type,
// and the value is set to Any().
func (s OneDriveScope) IsAny(cat onedriveCategory) bool {
	return isAnyTarget(s, cat)
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// OneDriveUser category.
func (s OneDriveScope) Get(cat onedriveCategory) []string {
	return getCatValue(s, cat)
}

// sets a value by category to the scope.  Only intended for internal use.
func (s OneDriveScope) set(cat onedriveCategory, v string) OneDriveScope {
	s[cat.String()] = v
	return s
}

// setDefaults ensures that user scopes express `AnyTgt` for their child category types.
func (s OneDriveScope) setDefaults() {
	// no-op while no child scope types below user are identified
}

// matchesEntry returns true if either the path or the info in the onedriveEntry matches the scope details.
func (s OneDriveScope) matchesEntry(
	cat categorizer,
	pathValues map[categorizer]string,
	entry details.DetailsEntry,
) bool {
	// matchesPathValues can be handled generically, thanks to SCIENCE.
	return matchesPathValues(s, cat.(onedriveCategory), pathValues) || s.matchesInfo(entry.Onedrive)
}

// matchesInfo handles the standard behavior when comparing a scope and an onedriveInfo
// returns true if the scope and info match for the provided category.
func (s OneDriveScope) matchesInfo(info *details.OnedriveInfo) bool {
	// we need values to match against
	if info == nil {
		return false
	}
	// the scope must define targets to match on
	filterCat := s.FilterCategory()
	targets := s.Get(filterCat)
	if len(targets) == 0 {
		return false
	}
	if targets[0] == AnyTgt {
		return true
	}
	if targets[0] == NoneTgt {
		return false
	}
	// any of the targets for a given info filter may succeed.
	for _, target := range targets {
		switch filterCat {
		// TODO: populate onedrive filter checks
		default:
			return target != NoneTgt
		}
	}
	return false
}
