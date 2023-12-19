package selectors

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// Selectors
// ---------------------------------------------------------------------------

type (
	// oneDrive provides an api for selecting
	// data scopes applicable to the OneDrive service.
	oneDrive struct {
		Selector
	}

	// OneDriveBackup provides an api for selecting
	// data scopes applicable to the OneDrive service,
	// plus backup-specific methods.
	OneDriveBackup struct {
		oneDrive
	}

	// OneDriveRestorep provides an api for selecting
	// data scopes applicable to the OneDrive service,
	// plus restore-specific methods.
	OneDriveRestore struct {
		oneDrive
	}
)

var (
	_ Reducer        = &OneDriveRestore{}
	_ pathCategorier = &OneDriveRestore{}
	_ reasoner       = &OneDriveRestore{}
)

// NewOneDriveBackup produces a new Selector with the service set to ServiceOneDrive.
func NewOneDriveBackup(users []string) *OneDriveBackup {
	src := OneDriveBackup{
		oneDrive{
			newSelector(ServiceOneDrive, users),
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

	src := OneDriveBackup{oneDrive{s}}

	return &src, nil
}

func (s OneDriveBackup) SplitByResourceOwner(users []string) []OneDriveBackup {
	sels := splitByProtectedResource[OneDriveScope](s.Selector, users, OneDriveUser)

	ss := make([]OneDriveBackup, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, OneDriveBackup{oneDrive{sel}})
	}

	return ss
}

// NewOneDriveRestore produces a new Selector with the service set to ServiceOneDrive.
func NewOneDriveRestore(users []string) *OneDriveRestore {
	src := OneDriveRestore{
		oneDrive{
			newSelector(ServiceOneDrive, users),
		},
	}

	return &src
}

// ToOneDriveRestore transforms the generic selector into an OneDriveRestore.
// Errors if the service defined by the selector is not ServiceOneDrive.
func (s Selector) ToOneDriveRestore() (*OneDriveRestore, error) {
	if s.Service != ServiceOneDrive {
		return nil, badCastErr(ServiceOneDrive, s.Service)
	}

	src := OneDriveRestore{oneDrive{s}}

	return &src, nil
}

func (s OneDriveRestore) SplitByResourceOwner(users []string) []OneDriveRestore {
	sels := splitByProtectedResource[OneDriveScope](s.Selector, users, OneDriveUser)

	ss := make([]OneDriveRestore, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, OneDriveRestore{oneDrive{sel}})
	}

	return ss
}

// PathCategories produces the aggregation of discrete users described by each type of scope.
func (s oneDrive) PathCategories() selectorPathCategories {
	return selectorPathCategories{
		Excludes: pathCategoriesIn[OneDriveScope, oneDriveCategory](s.Excludes),
		Filters:  pathCategoriesIn[OneDriveScope, oneDriveCategory](s.Filters),
		Includes: pathCategoriesIn[OneDriveScope, oneDriveCategory](s.Includes),
	}
}

// Reasons returns a deduplicated set of the backup reasons produced
// using the selector's discrete owner and each scopes' service and
// category types.
func (s oneDrive) Reasons(tenantID string, useOwnerNameForID bool) []identity.Reasoner {
	return reasonsFor(s, tenantID, useOwnerNameForID)
}

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

func (s OneDriveScope) Conceal() string             { return conceal(s) }
func (s OneDriveScope) Format(fs fmt.State, r rune) { format(s, fs, r) }
func (s OneDriveScope) String() string              { return conceal(s) }
func (s OneDriveScope) PlainString() string         { return plainString(s) }

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
func (s *oneDrive) Include(scopes ...[]OneDriveScope) {
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
func (s *oneDrive) Exclude(scopes ...[]OneDriveScope) {
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
func (s *oneDrive) Filter(scopes ...[]OneDriveScope) {
	s.Filters = appendScopes(s.Filters, scopes...)
}

// Scopes retrieves the list of oneDriveScopes in the selector.
func (s *oneDrive) Scopes() []OneDriveScope {
	return scopes[OneDriveScope](s.Selector)
}

// -------------------
// Scope Factories

// Retrieves all OneDrive data.
// One scope is created per user entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *oneDrive) AllData() []OneDriveScope {
	scopes := []OneDriveScope{}

	scopes = append(scopes, makeScope[OneDriveScope](OneDriveFolder, Any()))

	return scopes
}

// Folders produces one or more OneDrive folder scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *oneDrive) Folders(folders []string, opts ...option) []OneDriveScope {
	var (
		scopes = []OneDriveScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[OneDriveScope](OneDriveFolder, folders, os...))

	return scopes
}

// Items produces one or more OneDrive item scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *oneDrive) Items(folders, items []string, opts ...option) []OneDriveScope {
	scopes := []OneDriveScope{}

	scopes = append(
		scopes,
		makeScope[OneDriveScope](OneDriveItem, items, defaultItemOptions(s.Cfg)...).
			set(OneDriveFolder, folders, opts...))

	return scopes
}

// -------------------
// Filter Factories

// CreatedAfter produces a OneDrive item created-after info scope.
// Matches any item where the created time is after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *oneDrive) CreatedAfter(timeStrings string) []OneDriveScope {
	return []OneDriveScope{
		makeInfoScope[OneDriveScope](
			OneDriveItem,
			FileInfoCreatedAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

// CreatedBefore produces a OneDrive item created-before info scope.
// Matches any item where the created time is before the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *oneDrive) CreatedBefore(timeStrings string) []OneDriveScope {
	return []OneDriveScope{
		makeInfoScope[OneDriveScope](
			OneDriveItem,
			FileInfoCreatedBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

// ModifiedAfter produces a OneDrive item modified-after info scope.
// Matches any item where the modified time is after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *oneDrive) ModifiedAfter(timeStrings string) []OneDriveScope {
	return []OneDriveScope{
		makeInfoScope[OneDriveScope](
			OneDriveItem,
			FileInfoModifiedAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

// ModifiedBefore produces a OneDrive item modified-before info scope.
// Matches any item where the modified time is before the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *oneDrive) ModifiedBefore(timeStrings string) []OneDriveScope {
	return []OneDriveScope{
		makeInfoScope[OneDriveScope](
			OneDriveItem,
			FileInfoModifiedBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

// ---------------------------------------------------------------------------
// Categories
// ---------------------------------------------------------------------------

// oneDriveCategory enumerates the type of the lowest level
// of data () in a scope.
type oneDriveCategory string

// interface compliance checks
var _ categorizer = OneDriveCategoryUnknown

const (
	OneDriveCategoryUnknown oneDriveCategory = ""

	// types of data in OneDrive
	OneDriveUser   oneDriveCategory = "OneDriveUser"
	OneDriveItem   oneDriveCategory = "OneDriveItem"
	OneDriveFolder oneDriveCategory = "OneDriveFolder"

	// details.ItemInfo comparables
	FileInfoCreatedAfter   oneDriveCategory = "FileInfoCreatedAfter"
	FileInfoCreatedBefore  oneDriveCategory = "FileInfoCreatedBefore"
	FileInfoModifiedAfter  oneDriveCategory = "FileInfoModifiedAfter"
	FileInfoModifiedBefore oneDriveCategory = "FileInfoModifiedBefore"
)

// oneDriveLeafProperties describes common metadata of the leaf categories
var oneDriveLeafProperties = map[categorizer]leafProperty{
	OneDriveItem: {
		pathKeys: []categorizer{OneDriveFolder, OneDriveItem},
		pathType: path.FilesCategory,
	},
	OneDriveUser: { // the root category must be represented, even though it isn't a leaf
		pathKeys: []categorizer{OneDriveUser},
		pathType: path.UnknownCategory,
	},
}

func (c oneDriveCategory) String() string {
	return string(c)
}

// leafCat returns the leaf category of the receiver.
// If the receiver category has multiple leaves (ex: User) or no leaves,
// (ex: Unknown), the receiver itself is returned.
// Ex: ServiceTypeFolder.leafCat() => ServiceTypeItem
// Ex: ServiceUser.leafCat() => ServiceUser
func (c oneDriveCategory) leafCat() categorizer {
	switch c {
	case OneDriveFolder, OneDriveItem,
		FileInfoCreatedAfter, FileInfoCreatedBefore,
		FileInfoModifiedAfter, FileInfoModifiedBefore:
		return OneDriveItem
	}

	return c
}

// rootCat returns the root category type.
func (c oneDriveCategory) rootCat() categorizer {
	return OneDriveUser
}

// unknownCat returns the unknown category type.
func (c oneDriveCategory) unknownCat() categorizer {
	return OneDriveCategoryUnknown
}

// isUnion returns true if c is a user
func (c oneDriveCategory) isUnion() bool {
	return c == c.rootCat()
}

// isLeaf is true if the category is a OneDriveItem category.
func (c oneDriveCategory) isLeaf() bool {
	// return c == c.leafCat()??
	return c == OneDriveItem
}

// pathValues transforms the two paths to maps of identified properties.
//
// Example:
// [tenantID, service, userPN, category, folder, fileID]
// => {odFolder: folder, odFileID: fileID}
func (c oneDriveCategory) pathValues(
	repo path.Path,
	ent details.Entry,
	cfg Config,
) (map[categorizer][]string, error) {
	if ent.OneDrive == nil {
		return nil, clues.New("no OneDrive ItemInfo in details")
	}

	// Ignore `drives/<driveID>/root:` for folder comparison
	rFld := ent.OneDrive.ParentPath

	item := ent.ItemRef
	if len(item) == 0 {
		item = repo.Item()
	}

	if cfg.OnlyMatchItemNames {
		item = ent.ItemInfo.OneDrive.ItemName
	}

	result := map[categorizer][]string{
		OneDriveFolder: {rFld},
		OneDriveItem:   {item, ent.ShortRef},
	}

	if len(ent.LocationRef) > 0 {
		result[OneDriveFolder] = append(result[OneDriveFolder], ent.LocationRef)
	}

	return result, nil
}

// pathKeys returns the path keys recognized by the receiver's leaf type.
func (c oneDriveCategory) pathKeys() []categorizer {
	return oneDriveLeafProperties[c.leafCat()].pathKeys
}

// PathType converts the category's leaf type into the matching path.CategoryType.
func (c oneDriveCategory) PathType() path.CategoryType {
	return oneDriveLeafProperties[c.leafCat()].pathType
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
func (s OneDriveScope) Category() oneDriveCategory {
	return oneDriveCategory(getCategory(s))
}

// categorizer type is a generic wrapper around Category.
// Primarily used by scopes.go to for abstract comparisons.
func (s OneDriveScope) categorizer() categorizer {
	return s.Category()
}

// InfoCategory returns the category enum of the scope info.
// If the scope is not an info type, returns OneDriveUnknownCategory.
func (s OneDriveScope) InfoCategory() oneDriveCategory {
	return oneDriveCategory(getInfoCategory(s))
}

// IncludeCategory checks whether the scope includes a
// certain category of data.
// Ex: to check if the scope includes file data:
// s.IncludesCategory(selector.OneDriveFile)
func (s OneDriveScope) IncludesCategory(cat oneDriveCategory) bool {
	return categoryMatches(s.Category(), cat)
}

// Matches returns true if the category is included in the scope's
// data type, and the target string matches that category's comparator.
func (s OneDriveScope) Matches(cat oneDriveCategory, target string) bool {
	return matches(s, cat, target)
}

// returns true if the category is included in the scope's data type,
// and the value is set to Any().
func (s OneDriveScope) IsAny(cat oneDriveCategory) bool {
	return IsAnyTarget(s, cat)
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// OneDriveUser category.
func (s OneDriveScope) Get(cat oneDriveCategory) []string {
	return getCatValue(s, cat)
}

// sets a value by category to the scope.  Only intended for internal use.
func (s OneDriveScope) set(cat oneDriveCategory, v []string, opts ...option) OneDriveScope {
	os := []option{}
	if cat == OneDriveFolder {
		os = append(os, pathComparator())
	}

	return set(s, cat, v, append(os, opts...)...)
}

// setDefaults ensures that user scopes express `AnyTgt` for their child category types.
func (s OneDriveScope) setDefaults() {
	switch s.Category() {
	case OneDriveUser:
		s[OneDriveFolder.String()] = passAny
		s[OneDriveItem.String()] = passAny
	case OneDriveFolder:
		s[OneDriveItem.String()] = passAny
	}
}

// ---------------------------------------------------------------------------
// Backup Details Filtering
// ---------------------------------------------------------------------------

// Reduce filters the entries in a details struct to only those that match the
// inclusions, filters, and exclusions in the selector.
func (s oneDrive) Reduce(
	ctx context.Context,
	deets *details.Details,
	errs *fault.Bus,
) *details.Details {
	return reduce[OneDriveScope](
		ctx,
		deets,
		s.Selector,
		map[path.CategoryType]oneDriveCategory{
			path.FilesCategory: OneDriveItem,
		},
		errs)
}

// matchesInfo handles the standard behavior when comparing a scope and an oneDriveInfo
// returns true if the scope and info match for the provided category.
func (s OneDriveScope) matchesInfo(dii details.ItemInfo) bool {
	info := dii.OneDrive
	if info == nil {
		return false
	}

	infoCat := s.InfoCategory()

	i := ""

	switch infoCat {
	case FileInfoCreatedAfter, FileInfoCreatedBefore:
		i = dttm.Format(info.Created)
	case FileInfoModifiedAfter, FileInfoModifiedBefore:
		i = dttm.Format(info.Modified)
	}

	return s.Matches(infoCat, i)
}
