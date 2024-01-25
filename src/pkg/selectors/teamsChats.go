package selectors

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"

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
	// teamsChats provides an api for selecting
	// data scopes applicable to the TeamsChats service.
	teamsChats struct {
		Selector
	}

	// TeamsChatsBackup provides an api for selecting
	// data scopes applicable to the TeamsChats service,
	// plus backup-specific methods.
	TeamsChatsBackup struct {
		teamsChats
	}

	// TeamsChatsRestore provides an api for selecting
	// data scopes applicable to the TeamsChats service,
	// plus restore-specific methods.
	TeamsChatsRestore struct {
		teamsChats
	}
)

var (
	_ Reducer        = &TeamsChatsRestore{}
	_ pathCategorier = &TeamsChatsRestore{}
	_ reasoner       = &TeamsChatsRestore{}
)

// NewTeamsChats produces a new Selector with the service set to ServiceTeamsChats.
func NewTeamsChatsBackup(users []string) *TeamsChatsBackup {
	src := TeamsChatsBackup{
		teamsChats{
			newSelector(ServiceTeamsChats, users),
		},
	}

	return &src
}

// ToTeamsChatsBackup transforms the generic selector into an TeamsChatsBackup.
// Errors if the service defined by the selector is not ServiceTeamsChats.
func (s Selector) ToTeamsChatsBackup() (*TeamsChatsBackup, error) {
	if s.Service != ServiceTeamsChats {
		return nil, badCastErr(ServiceTeamsChats, s.Service)
	}

	src := TeamsChatsBackup{teamsChats{s}}

	return &src, nil
}

func (s TeamsChatsBackup) SplitByResourceOwner(users []string) []TeamsChatsBackup {
	sels := splitByProtectedResource[TeamsChatsScope](s.Selector, users, TeamsChatsUser)

	ss := make([]TeamsChatsBackup, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, TeamsChatsBackup{teamsChats{sel}})
	}

	return ss
}

// NewTeamsChatsRestore produces a new Selector with the service set to ServiceTeamsChats.
func NewTeamsChatsRestore(users []string) *TeamsChatsRestore {
	src := TeamsChatsRestore{
		teamsChats{
			newSelector(ServiceTeamsChats, users),
		},
	}

	return &src
}

// ToTeamsChatsRestore transforms the generic selector into an TeamsChatsRestore.
// Errors if the service defined by the selector is not ServiceTeamsChats.
func (s Selector) ToTeamsChatsRestore() (*TeamsChatsRestore, error) {
	if s.Service != ServiceTeamsChats {
		return nil, badCastErr(ServiceTeamsChats, s.Service)
	}

	src := TeamsChatsRestore{teamsChats{s}}

	return &src, nil
}

func (sr TeamsChatsRestore) SplitByResourceOwner(users []string) []TeamsChatsRestore {
	sels := splitByProtectedResource[TeamsChatsScope](sr.Selector, users, TeamsChatsUser)

	ss := make([]TeamsChatsRestore, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, TeamsChatsRestore{teamsChats{sel}})
	}

	return ss
}

// PathCategories produces the aggregation of discrete users described by each type of scope.
func (s teamsChats) PathCategories() selectorPathCategories {
	return selectorPathCategories{
		Excludes: pathCategoriesIn[TeamsChatsScope, teamsChatsCategory](s.Excludes),
		Filters:  pathCategoriesIn[TeamsChatsScope, teamsChatsCategory](s.Filters),
		Includes: pathCategoriesIn[TeamsChatsScope, teamsChatsCategory](s.Includes),
	}
}

// Reasons returns a deduplicated set of the backup reasons produced
// using the selector's discrete owner and each scopes' service and
// category types.
func (s teamsChats) Reasons(tenantID string, useOwnerNameForID bool) []identity.Reasoner {
	return reasonsFor(s, tenantID, useOwnerNameForID)
}

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

func (s TeamsChatsScope) Conceal() string             { return conceal(s) }
func (s TeamsChatsScope) Format(fs fmt.State, r rune) { format(s, fs, r) }
func (s TeamsChatsScope) String() string              { return conceal(s) }
func (s TeamsChatsScope) PlainString() string         { return plainString(s) }

// -------------------
// Exclude/Includes

// Exclude appends the provided scopes to the selector's exclusion set.
// Every Exclusion scope applies globally, affecting all inclusion scopes.
// Data is excluded if it matches ANY exclusion (of the same data category).
//
// All parts of the scope must match for data to be exclucded.
// Ex: Mail(u1, f1, m1) => only excludes mail if it is owned by user u1,
// located in folder f1, and ID'd as m1.  MailSender(foo) => only excludes
// mail whose sender is foo.  Use selectors.Any() to wildcard a scope value.
// No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: User(u1) automatically cascades to all chats,
func (s *teamsChats) Exclude(scopes ...[]TeamsChatsScope) {
	s.Excludes = appendScopes(s.Excludes, scopes...)
}

// Filter appends the provided scopes to the selector's filters set.
// A selector with >0 filters and 0 inclusions will include any data
// that passes all filters.
// A selector with >0 filters and >0 inclusions will reduce the
// inclusion set to only the data that passes all filters.
// Data is retained if it passes ALL filters (of the same data category).
//
// All parts of the scope must match for data to pass the filter.
// Ex: Mail(u1, f1, m1) => only passes mail that is owned by user u1,
// located in folder f1, and ID'd as m1.  MailSender(foo) => only passes
// mail whose sender is foo.  Use selectors.Any() to wildcard a scope value.
// No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: User(u1) automatically cascades to all chats,
func (s *teamsChats) Filter(scopes ...[]TeamsChatsScope) {
	s.Filters = appendScopes(s.Filters, scopes...)
}

// Include appends the provided scopes to the selector's inclusion set.
// Data is included if it matches ANY inclusion.
// The inclusion set is later filtered (all included data must pass ALL
// filters) and excluded (all included data must not match ANY exclusion).
// Data is included if it matches ANY inclusion (of the same data category).
//
// All parts of the scope must match for data to be included.
// Ex: Mail(u1, f1, m1) => only includes mail if it is owned by user u1,
// located in folder f1, and ID'd as m1.  MailSender(foo) => only includes
// mail whose sender is foo.  Use selectors.Any() to wildcard a scope value.
// No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: User(u1) automatically cascades to all chats,
func (s *teamsChats) Include(scopes ...[]TeamsChatsScope) {
	s.Includes = appendScopes(s.Includes, scopes...)
}

// Scopes retrieves the list of teamsChatsScopes in the selector.
func (s *teamsChats) Scopes() []TeamsChatsScope {
	return scopes[TeamsChatsScope](s.Selector)
}

type TeamsChatsItemScopeConstructor func([]string, []string, ...option) []TeamsChatsScope

// -------------------
// Scope Factories

// Chats produces one or more teamsChats scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *teamsChats) Chats(chats []string, opts ...option) []TeamsChatsScope {
	scopes := []TeamsChatsScope{}

	scopes = append(
		scopes,
		makeScope[TeamsChatsScope](TeamsChatsChat, chats, defaultItemOptions(s.Cfg)...))

	return scopes
}

// Retrieves all teamsChats data.
// Each user id generates a scope for each data type: chats (only one data type at this time).
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *teamsChats) AllData() []TeamsChatsScope {
	scopes := []TeamsChatsScope{}

	scopes = append(scopes, makeScope[TeamsChatsScope](TeamsChatsChat, Any()))

	return scopes
}

// -------------------
// ItemInfo Factories

// ChatMember produces one or more teamsChats chat member info scopes.
// Matches any chat member whose email contains the provided string.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *TeamsChatsRestore) ChatMember(memberID string) []TeamsChatsScope {
	return []TeamsChatsScope{
		makeInfoScope[TeamsChatsScope](
			TeamsChatsChat,
			TeamsChatsInfoChatMember,
			[]string{memberID},
			filters.In),
	}
}

// ChatName produces one or more teamsChats chat name info scopes.
// Matches any chat whose name contains the provided string.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *TeamsChatsRestore) ChatName(memberID string) []TeamsChatsScope {
	return []TeamsChatsScope{
		makeInfoScope[TeamsChatsScope](
			TeamsChatsChat,
			TeamsChatsInfoChatName,
			[]string{memberID},
			filters.In),
	}
}

// ---------------------------------------------------------------------------
// Categories
// ---------------------------------------------------------------------------

// teamsChatsCategory enumerates the type of the lowest level
// of data specified by the scope.
type teamsChatsCategory string

// interface compliance checks
var _ categorizer = TeamsChatsCategoryUnknown

const (
	TeamsChatsCategoryUnknown teamsChatsCategory = ""

	// types of data identified by teamsChats
	TeamsChatsUser teamsChatsCategory = "TeamsChatsUser"
	TeamsChatsChat teamsChatsCategory = "TeamsChatsChat"

	// data contained within details.ItemInfo
	TeamsChatsInfoChatMember teamsChatsCategory = "TeamsChatsInfoChatMember"
	TeamsChatsInfoChatName   teamsChatsCategory = "TeamsChatsInfoChatName"
)

// teamsChatsLeafProperties describes common metadata of the leaf categories
var teamsChatsLeafProperties = map[categorizer]leafProperty{
	TeamsChatsChat: {
		pathKeys: []categorizer{TeamsChatsChat},
		pathType: path.ChatsCategory,
	},
	TeamsChatsUser: { // the root category must be represented, even though it isn't a leaf
		pathKeys: []categorizer{TeamsChatsUser},
		pathType: path.UnknownCategory,
	},
}

func (ec teamsChatsCategory) String() string {
	return string(ec)
}

// leafCat returns the leaf category of the receiver.
// If the receiver category has multiple leaves (ex: User) or no leaves,
// (ex: Unknown), the receiver itself is returned.
// If the receiver category is an info type (ex: TeamsChatsInfoChatMember),
// returns the category covered by the info.
// Ex: TeamsChatsChatFolder.leafCat() => TeamsChatsChat
// Ex: TeamsChatsUser.leafCat() => TeamsChatsUser
func (ec teamsChatsCategory) leafCat() categorizer {
	switch ec {
	case TeamsChatsChat, TeamsChatsInfoChatMember, TeamsChatsInfoChatName:
		return TeamsChatsChat
	}

	return ec
}

// rootCat returns the root category type.
func (ec teamsChatsCategory) rootCat() categorizer {
	return TeamsChatsUser
}

// unknownCat returns the unknown category type.
func (ec teamsChatsCategory) unknownCat() categorizer {
	return TeamsChatsCategoryUnknown
}

// isUnion returns true if c is a user
func (ec teamsChatsCategory) isUnion() bool {
	return ec == ec.rootCat()
}

// isLeaf is true if the category is a mail, event, or contact category.
func (ec teamsChatsCategory) isLeaf() bool {
	return ec == ec.leafCat()
}

// pathValues transforms the two paths to maps of identified properties.
//
// Example:
// [tenantID, service, userID, category, chatID]
// => {teamsChat: chatID}
func (ec teamsChatsCategory) pathValues(
	repo path.Path,
	ent details.Entry,
	cfg Config,
) (map[categorizer][]string, error) {
	var itemCat categorizer

	switch ec {
	case TeamsChatsChat:
		itemCat = TeamsChatsChat

	default:
		return nil, clues.New("bad Chat Category").With("category", ec)
	}

	item := ent.ItemRef
	if len(item) == 0 {
		item = repo.Item()
	}

	items := []string{ent.ShortRef, item}

	// only include the item ID when the user is NOT matching
	// item names. TeamsChats data does not contain an item name,
	// only an ID, and we don't want to mix up the two.
	if cfg.OnlyMatchItemNames {
		items = []string{ent.ShortRef}
	}

	result := map[categorizer][]string{
		itemCat: items,
	}

	return result, nil
}

// pathKeys returns the path keys recognized by the receiver's leaf type.
func (ec teamsChatsCategory) pathKeys() []categorizer {
	return teamsChatsLeafProperties[ec.leafCat()].pathKeys
}

// PathType converts the category's leaf type into the matching path.CategoryType.
func (ec teamsChatsCategory) PathType() path.CategoryType {
	return teamsChatsLeafProperties[ec.leafCat()].pathType
}

// ---------------------------------------------------------------------------
// Scopes
// ---------------------------------------------------------------------------

// TeamsChatsScope specifies the data available
// when interfacing with the TeamsChats service.
type TeamsChatsScope scope

// interface compliance checks
var _ scoper = &TeamsChatsScope{}

// Category describes the type of the data in scope.
func (s TeamsChatsScope) Category() teamsChatsCategory {
	return teamsChatsCategory(getCategory(s))
}

// categorizer type is a generic wrapper around Category.
// Primarily used by scopes.go to for abstract comparisons.
func (s TeamsChatsScope) categorizer() categorizer {
	return s.Category()
}

// Matches returns true if the category is included in the scope's
// data type, and the target string matches that category's comparator.
func (s TeamsChatsScope) Matches(cat teamsChatsCategory, target string) bool {
	return matches(s, cat, target)
}

// InfoCategory returns the category enum of the scope info.
// If the scope is not an info type, returns TeamsChatsUnknownCategory.
func (s TeamsChatsScope) InfoCategory() teamsChatsCategory {
	return teamsChatsCategory(getInfoCategory(s))
}

// IncludeCategory checks whether the scope includes a certain category of data.
// Ex: to check if the scope includes mail data:
// s.IncludesCategory(selector.TeamsChatsMail)
func (s TeamsChatsScope) IncludesCategory(cat teamsChatsCategory) bool {
	return categoryMatches(s.Category(), cat)
}

// returns true if the category is included in the scope's data type,
// and the value is set to Any().
func (s TeamsChatsScope) IsAny(cat teamsChatsCategory) bool {
	return IsAnyTarget(s, cat)
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// TeamsChatsUser category.
func (s TeamsChatsScope) Get(cat teamsChatsCategory) []string {
	return getCatValue(s, cat)
}

// commenting out for now, but leaving in place; it's likely to return when we add filters
// // sets a value by category to the scope.  Only intended for internal use.
// func (s TeamsChatsScope) set(cat teamsChatsCategory, v []string, opts ...option) TeamsChatsScope {
// 	return set(s, cat, v, opts...)
// }

// setDefaults ensures that contact folder, mail folder, and user category
// scopes all express `AnyTgt` for their child category types.
func (s TeamsChatsScope) setDefaults() {
	switch s.Category() {
	case TeamsChatsUser:
		s[TeamsChatsChat.String()] = passAny
	}
}

// ---------------------------------------------------------------------------
// Backup Details Filtering
// ---------------------------------------------------------------------------

// Reduce filters the entries in a details struct to only those that match the
// inclusions, filters, and exclusions in the selector.
func (s teamsChats) Reduce(
	ctx context.Context,
	deets *details.Details,
	errs *fault.Bus,
) *details.Details {
	return reduce[TeamsChatsScope](
		ctx,
		deets,
		s.Selector,
		map[path.CategoryType]teamsChatsCategory{
			path.ChatsCategory: TeamsChatsChat,
		},
		errs)
}

// matchesInfo handles the standard behavior when comparing a scope and an TeamsChatsInfo
// returns true if the scope and info match for the provided category.
func (s TeamsChatsScope) matchesInfo(dii details.ItemInfo) bool {
	info := dii.TeamsChats
	if info == nil {
		return false
	}

	infoCat := s.InfoCategory()

	cfpc := teamsChatsCategoryFromItemType(info.ItemType)
	if !typeAndCategoryMatches(infoCat, cfpc) {
		return false
	}

	i := ""

	switch infoCat {
	case TeamsChatsInfoChatMember:
		i = strings.Join(info.Chat.Members, ",")
	case TeamsChatsInfoChatName:
		i = info.Chat.Name
	}

	return s.Matches(infoCat, i)
}

// teamsChatsCategoryFromItemType interprets the category represented by the TeamsChatsInfo
// struct.  Since every TeamsChatsInfo can hold all teamsChats data info, the exact
// type that the struct represents must be compared using its ItemType prop.
func teamsChatsCategoryFromItemType(pct details.ItemType) teamsChatsCategory {
	switch pct {
	case details.TeamsChat:
		return TeamsChatsChat
	}

	return TeamsChatsCategoryUnknown
}
