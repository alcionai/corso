package selectors

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// Selectors
// ---------------------------------------------------------------------------

type (
	// groups provides an api for selecting
	// data scopes applicable to the groups service.
	groups struct {
		Selector
	}

	// groups provides an api for selecting
	// data scopes applicable to the groups service,
	// plus backup-specific methods.
	GroupsBackup struct {
		groups
	}

	// GroupsRestorep provides an api for selecting
	// data scopes applicable to the Groups service,
	// plus restore-specific methods.
	GroupsRestore struct {
		groups
	}
)

var (
	_ Reducer        = &GroupsRestore{}
	_ pathCategorier = &GroupsRestore{}
	_ reasoner       = &GroupsRestore{}
)

// NewGroupsBackup produces a new Selector with the service set to ServiceGroups.
func NewGroupsBackup(resources []string) *GroupsBackup {
	src := GroupsBackup{
		groups{
			newSelector(ServiceGroups, resources),
		},
	}

	return &src
}

// ToGroupsBackup transforms the generic selector into an GroupsBackup.
// Errors if the service defined by the selector is not ServiceGroups.
func (s Selector) ToGroupsBackup() (*GroupsBackup, error) {
	if s.Service != ServiceGroups {
		return nil, badCastErr(ServiceGroups, s.Service)
	}

	src := GroupsBackup{groups{s}}

	return &src, nil
}

func (s GroupsBackup) SplitByResourceOwner(resources []string) []GroupsBackup {
	sels := splitByProtectedResource[GroupsScope](s.Selector, resources, GroupsGroup)

	ss := make([]GroupsBackup, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, GroupsBackup{groups{sel}})
	}

	return ss
}

// NewGroupsRestore produces a new Selector with the service set to ServiceGroups.
func NewGroupsRestore(resources []string) *GroupsRestore {
	src := GroupsRestore{
		groups{
			newSelector(ServiceGroups, resources),
		},
	}

	return &src
}

// ToGroupsRestore transforms the generic selector into an GroupsRestore.
// Errors if the service defined by the selector is not ServiceGroups.
func (s Selector) ToGroupsRestore() (*GroupsRestore, error) {
	if s.Service != ServiceGroups {
		return nil, badCastErr(ServiceGroups, s.Service)
	}

	src := GroupsRestore{groups{s}}

	return &src, nil
}

func (s GroupsRestore) SplitByResourceOwner(resources []string) []GroupsRestore {
	sels := splitByProtectedResource[GroupsScope](s.Selector, resources, GroupsGroup)

	ss := make([]GroupsRestore, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, GroupsRestore{groups{sel}})
	}

	return ss
}

// PathCategories produces the aggregation of discrete resources described by each type of scope.
func (s groups) PathCategories() selectorPathCategories {
	return selectorPathCategories{
		Excludes: pathCategoriesIn[GroupsScope, groupsCategory](s.Excludes),
		Filters:  pathCategoriesIn[GroupsScope, groupsCategory](s.Filters),
		Includes: pathCategoriesIn[GroupsScope, groupsCategory](s.Includes),
	}
}

// Reasons returns a deduplicated set of the backup reasons produced
// using the selector's discrete owner and each scopes' service and
// category types.
func (s groups) Reasons(tenantID string, useOwnerNameForID bool) []identity.Reasoner {
	return reasonsFor(s, tenantID, useOwnerNameForID)
}

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

func (s GroupsScope) Conceal() string             { return conceal(s) }
func (s GroupsScope) Format(fs fmt.State, r rune) { format(s, fs, r) }
func (s GroupsScope) String() string              { return conceal(s) }
func (s GroupsScope) PlainString() string         { return plainString(s) }

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
func (s *groups) Include(scopes ...[]GroupsScope) {
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
func (s *groups) Exclude(scopes ...[]GroupsScope) {
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
func (s *groups) Filter(scopes ...[]GroupsScope) {
	s.Filters = appendScopes(s.Filters, scopes...)
}

// Scopes retrieves the list of groupsScopes in the selector.
func (s *groups) Scopes() []GroupsScope {
	return scopes[GroupsScope](s.Selector)
}

// -------------------
// Scope Factories

// Produces one or more Groups scopes.
// One scope is created per group entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *groups) AllData() []GroupsScope {
	scopes := []GroupsScope{}

	scopes = append(
		scopes,
		makeScope[GroupsScope](GroupsLibraryFolder, Any()),
		makeScope[GroupsScope](GroupsChannel, Any()))

	return scopes
}

// Channels produces one or more SharePoint channel scopes, where the channel
// matches upon a given channel by ID or Name.  In order to ensure channel selection
// this should always be embedded within the Filter() set; include(channel()) will
// select all items in the channel without further filtering.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *groups) Channels(channels []string) []GroupsScope {
	return []GroupsScope{
		makeInfoScope[GroupsScope](
			GroupsChannel,
			GroupsInfoChannel,
			channels,
			filters.Equal),
	}
}

// ChannelMessages produces one or more Groups channel message scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *groups) ChannelMessages(channels, messages []string, opts ...option) []GroupsScope {
	var (
		scopes = []GroupsScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[GroupsScope](GroupsChannelMessage, messages, os...).
			set(GroupsChannel, channels, opts...))

	return scopes
}

// Library produces one or more Group library scopes, where the library
// matches upon a given drive by ID or Name.  In order to ensure library selection
// this should always be embedded within the Filter() set; include(Library()) will
// select all items in the library without further filtering.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *groups) Library(library string) []GroupsScope {
	return []GroupsScope{
		makeInfoScope[GroupsScope](
			GroupsLibraryItem,
			GroupsInfoSiteLibraryDrive,
			[]string{library},
			filters.Equal),
	}
}

// LibraryFolders produces one or more SharePoint libraryFolder scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *groups) LibraryFolders(libraryFolders []string, opts ...option) []GroupsScope {
	var (
		scopes = []GroupsScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[GroupsScope](GroupsLibraryFolder, libraryFolders, os...))

	return scopes
}

// LibraryItems produces one or more Groups library item scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the library scopes.
func (s *groups) LibraryItems(libraries, items []string, opts ...option) []GroupsScope {
	scopes := []GroupsScope{}

	scopes = append(
		scopes,
		makeScope[GroupsScope](GroupsLibraryItem, items, defaultItemOptions(s.Cfg)...).
			set(GroupsLibraryFolder, libraries, opts...))

	return scopes
}

// -------------------
// ItemInfo Factories

// MessageCreator produces one or more groups channelMessage info scopes.
// Matches any channel message created by the specified user.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *GroupsRestore) MessageCreator(creator string) []GroupsScope {
	return []GroupsScope{
		makeInfoScope[GroupsScope](
			GroupsChannelMessage,
			GroupsInfoChannelMessageCreator,
			[]string{creator},
			filters.In),
	}
}

// MessageCreatedAfter produces a channel message created-after info scope.
// Matches any message where the creation time is after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *GroupsRestore) MessageCreatedAfter(timeStrings string) []GroupsScope {
	return []GroupsScope{
		makeInfoScope[GroupsScope](
			GroupsChannelMessage,
			GroupsInfoChannelMessageCreatedAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

// MessageCreatedBefore produces a channel message created-before info scope.
// Matches any message where the creation time is after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *GroupsRestore) MessageCreatedBefore(timeStrings string) []GroupsScope {
	return []GroupsScope{
		makeInfoScope[GroupsScope](
			GroupsChannelMessage,
			GroupsInfoChannelMessageCreatedBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

// MessageLastReplyAfter produces a channel message last-response-after info scope.
// Matches any message where last response time is after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *GroupsRestore) MessageLastReplyAfter(timeStrings string) []GroupsScope {
	return []GroupsScope{
		makeInfoScope[GroupsScope](
			GroupsChannelMessage,
			GroupsInfoChannelMessageLastReplyAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

// MessageLastReplyBefore produces a channel message last-response-before info scope.
// Matches any message where last response time is after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (s *GroupsRestore) MessageLastReplyBefore(timeStrings string) []GroupsScope {
	return []GroupsScope{
		makeInfoScope[GroupsScope](
			GroupsChannelMessage,
			GroupsInfoChannelMessageLastReplyBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

// ---------------------------------------------------------------------------
// Categories
// ---------------------------------------------------------------------------

// groupsCategory enumerates the type of the lowest level
// of data () in a scope.
type groupsCategory string

// interface compliance checks
var _ categorizer = GroupsCategoryUnknown

const (
	GroupsCategoryUnknown groupsCategory = ""

	// types of data in Groups
	GroupsGroup          groupsCategory = "GroupsGroup"
	GroupsChannel        groupsCategory = "GroupsChannel"
	GroupsChannelMessage groupsCategory = "GroupsChannelMessage"
	GroupsLibraryFolder  groupsCategory = "GroupsLibraryFolder"
	GroupsLibraryItem    groupsCategory = "GroupsLibraryItem"

	// details.itemInfo comparables

	// channel and drive selection
	GroupsInfoSiteLibraryDrive groupsCategory = "GroupsInfoSiteLibraryDrive"
	GroupsInfoChannel          groupsCategory = "GroupsInfoChannel"

	// data contained within details.ItemInfo
	GroupsInfoChannelMessageCreatedAfter    groupsCategory = "GroupsInfoChannelMessageCreatedAfter"
	GroupsInfoChannelMessageCreatedBefore   groupsCategory = "GroupsInfoChannelMessageCreatedBefore"
	GroupsInfoChannelMessageCreator         groupsCategory = "GroupsInfoChannelMessageCreator"
	GroupsInfoChannelMessageLastReplyAfter  groupsCategory = "GroupsInfoChannelMessageLastReplyAfter"
	GroupsInfoChannelMessageLastReplyBefore groupsCategory = "GroupsInfoChannelMessageLastReplyBefore"
)

// groupsLeafProperties describes common metadata of the leaf categories
var groupsLeafProperties = map[categorizer]leafProperty{
	GroupsChannelMessage: { // the root category must be represented, even though it isn't a leaf
		pathKeys: []categorizer{GroupsChannel, GroupsChannelMessage},
		pathType: path.ChannelMessagesCategory,
	},
	GroupsLibraryItem: {
		pathKeys: []categorizer{GroupsLibraryFolder, GroupsLibraryItem},
		pathType: path.LibrariesCategory,
	},
	GroupsGroup: { // the root category must be represented, even though it isn't a leaf
		pathKeys: []categorizer{GroupsGroup},
		pathType: path.UnknownCategory,
	},
}

func (c groupsCategory) String() string {
	return string(c)
}

// leafCat returns the leaf category of the receiver.
// If the receiver category has multiple leaves (ex: User) or no leaves,
// (ex: Unknown), the receiver itself is returned.
// Ex: ServiceTypeFolder.leafCat() => ServiceTypeItem
// Ex: ServiceUser.leafCat() => ServiceUser
func (c groupsCategory) leafCat() categorizer {
	switch c {
	// TODO: if channels ever contain more than one type of item,
	// we'll need to fix this up.
	case GroupsChannel, GroupsChannelMessage,
		GroupsInfoChannelMessageCreatedAfter, GroupsInfoChannelMessageCreatedBefore, GroupsInfoChannelMessageCreator,
		GroupsInfoChannelMessageLastReplyAfter, GroupsInfoChannelMessageLastReplyBefore:
		return GroupsChannelMessage
	case GroupsLibraryFolder, GroupsLibraryItem, GroupsInfoSiteLibraryDrive:
		return GroupsLibraryItem
	}

	return c
}

// rootCat returns the root category type.
func (c groupsCategory) rootCat() categorizer {
	return GroupsGroup
}

// unknownCat returns the unknown category type.
func (c groupsCategory) unknownCat() categorizer {
	return GroupsCategoryUnknown
}

// isUnion returns true if the category is a site or a webURL, which
// can act as an alternative identifier to siteID across all site types.
func (c groupsCategory) isUnion() bool {
	return c == c.rootCat()
}

// isLeaf is true if the category is a GroupsItem category.
func (c groupsCategory) isLeaf() bool {
	return c == c.leafCat()
}

// pathValues transforms the two paths to maps of identified properties.
//
// Example:
// [tenantID, service, groupID, site, siteID, category, folder, itemID]
// => {spFolder: folder, spItemID: itemID}
func (c groupsCategory) pathValues(
	repo path.Path,
	ent details.Entry,
	cfg Config,
) (map[categorizer][]string, error) {
	var (
		folderCat, itemCat categorizer
		itemID             string
		rFld               string
	)

	if ent.Groups == nil {
		return nil, clues.New("no Groups ItemInfo in details")
	}

	switch c {
	case GroupsChannel, GroupsChannelMessage:
		folderCat, itemCat = GroupsChannel, GroupsChannelMessage
		rFld = ent.Groups.ParentPath
	case GroupsLibraryFolder, GroupsLibraryItem:
		folderCat, itemCat = GroupsLibraryFolder, GroupsLibraryItem
		rFld = ent.Groups.ParentPath
	default:
		return nil, clues.New("unrecognized groupsCategory").With("category", c)
	}

	item := ent.ItemRef
	if len(item) == 0 {
		item = repo.Item()
	}

	if cfg.OnlyMatchItemNames {
		item = ent.ItemInfo.Groups.ItemName
	}

	result := map[categorizer][]string{
		folderCat: {rFld},
		itemCat:   {item, ent.ShortRef},
	}

	if len(itemID) > 0 {
		result[itemCat] = append(result[itemCat], itemID)
	}

	return result, nil
}

// pathKeys returns the path keys recognized by the receiver's leaf type.
func (c groupsCategory) pathKeys() []categorizer {
	return groupsLeafProperties[c.leafCat()].pathKeys
}

// PathType converts the category's leaf type into the matching path.CategoryType.
func (c groupsCategory) PathType() path.CategoryType {
	return groupsLeafProperties[c.leafCat()].pathType
}

// ---------------------------------------------------------------------------
// Scopes
// ---------------------------------------------------------------------------

// GroupsScope specifies the data available
// when interfacing with the Groups service.
type GroupsScope scope

// interface compliance checks
var _ scoper = &GroupsScope{}

// Category describes the type of the data in scope.
func (s GroupsScope) Category() groupsCategory {
	return groupsCategory(getCategory(s))
}

// categorizer type is a generic wrapper around Category.
// Primarily used by scopes.go to for abstract comparisons.
func (s GroupsScope) categorizer() categorizer {
	return s.Category()
}

// Matches returns true if the category is included in the scope's
// data type, and the target string matches that category's comparator.
func (s GroupsScope) Matches(cat groupsCategory, target string) bool {
	return matches(s, cat, target)
}

// InfoCategory returns the category enum of the scope info.
// If the scope is not an info type, returns GroupsUnknownCategory.
func (s GroupsScope) InfoCategory() groupsCategory {
	return groupsCategory(getInfoCategory(s))
}

// IncludeCategory checks whether the scope includes a
// certain category of data.
// Ex: to check if the scope includes file data:
// s.IncludesCategory(selector.GroupsFile)
func (s GroupsScope) IncludesCategory(cat groupsCategory) bool {
	return categoryMatches(s.Category(), cat)
}

// returns true if the category is included in the scope's data type,
// and the value is set to Any().
func (s GroupsScope) IsAny(cat groupsCategory) bool {
	return isAnyTarget(s, cat)
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// GroupsUser category.
func (s GroupsScope) Get(cat groupsCategory) []string {
	return getCatValue(s, cat)
}

// sets a value by category to the scope.  Only intended for internal use.
func (s GroupsScope) set(cat groupsCategory, v []string, opts ...option) GroupsScope {
	os := []option{}

	switch cat {
	case GroupsChannel, GroupsLibraryFolder:
		os = append(os, pathComparator())
	}

	return set(s, cat, v, append(os, opts...)...)
}

// setDefaults ensures that site scopes express `AnyTgt` for their child category types.
func (s GroupsScope) setDefaults() {
	switch s.Category() {
	case GroupsGroup:
		s[GroupsChannel.String()] = passAny
		s[GroupsChannelMessage.String()] = passAny
		s[GroupsLibraryFolder.String()] = passAny
		s[GroupsLibraryItem.String()] = passAny
	case GroupsChannel:
		s[GroupsChannelMessage.String()] = passAny
	case GroupsLibraryFolder:
		s[GroupsLibraryItem.String()] = passAny
	}
}

// ---------------------------------------------------------------------------
// Backup Details Filtering
// ---------------------------------------------------------------------------

// Reduce filters the entries in a details struct to only those that match the
// inclusions, filters, and exclusions in the selector.
func (s groups) Reduce(
	ctx context.Context,
	deets *details.Details,
	errs *fault.Bus,
) *details.Details {
	return reduce[GroupsScope](
		ctx,
		deets,
		s.Selector,
		map[path.CategoryType]groupsCategory{
			path.ChannelMessagesCategory: GroupsChannelMessage,
			path.LibrariesCategory:       GroupsLibraryItem,
		},
		errs)
}

// matchesInfo handles the standard behavior when comparing a scope and an groupsInfo
// returns true if the scope and info match for the provided category.
func (s GroupsScope) matchesInfo(dii details.ItemInfo) bool {
	var (
		infoCat = s.InfoCategory()
		i       = ""
		info    = dii.Groups
	)

	if info == nil {
		return false
	}

	switch infoCat {
	case GroupsInfoSiteLibraryDrive:
		ds := []string{}

		if len(info.DriveName) > 0 {
			ds = append(ds, info.DriveName)
		}

		if len(info.DriveID) > 0 {
			ds = append(ds, info.DriveID)
		}

		return matchesAny(s, GroupsInfoSiteLibraryDrive, ds)
	case GroupsInfoChannel:
		ds := []string{}

		if len(info.ChannelID) > 0 {
			ds = append(ds, info.ChannelID)
		}

		if len(info.ChannelName) > 0 {
			ds = append(ds, info.ChannelName)
		}

		return matchesAny(s, GroupsInfoChannel, ds)
	case GroupsInfoChannelMessageCreator:
		i = info.MessageCreator
	case GroupsInfoChannelMessageCreatedAfter, GroupsInfoChannelMessageCreatedBefore:
		i = dttm.Format(info.Created)
	case GroupsInfoChannelMessageLastReplyAfter, GroupsInfoChannelMessageLastReplyBefore:
		i = dttm.Format(info.LastReplyAt)
	}

	return s.Matches(infoCat, i)
}
