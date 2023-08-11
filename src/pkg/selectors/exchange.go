package selectors

import (
	"context"
	"fmt"
	"strconv"

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
	// exchange provides an api for selecting
	// data scopes applicable to the Exchange service.
	exchange struct {
		Selector
	}

	// ExchangeBackup provides an api for selecting
	// data scopes applicable to the Exchange service,
	// plus backup-specific methods.
	ExchangeBackup struct {
		exchange
	}

	// ExchangeRestore provides an api for selecting
	// data scopes applicable to the Exchange service,
	// plus restore-specific methods.
	ExchangeRestore struct {
		exchange
	}
)

var (
	_ Reducer        = &ExchangeRestore{}
	_ pathCategorier = &ExchangeRestore{}
	_ reasoner       = &ExchangeRestore{}
)

// NewExchange produces a new Selector with the service set to ServiceExchange.
func NewExchangeBackup(users []string) *ExchangeBackup {
	src := ExchangeBackup{
		exchange{
			newSelector(ServiceExchange, users),
		},
	}

	return &src
}

// ToExchangeBackup transforms the generic selector into an ExchangeBackup.
// Errors if the service defined by the selector is not ServiceExchange.
func (s Selector) ToExchangeBackup() (*ExchangeBackup, error) {
	if s.Service != ServiceExchange {
		return nil, badCastErr(ServiceExchange, s.Service)
	}

	src := ExchangeBackup{exchange{s}}

	return &src, nil
}

func (s ExchangeBackup) SplitByResourceOwner(users []string) []ExchangeBackup {
	sels := splitByProtectedResource[ExchangeScope](s.Selector, users, ExchangeUser)

	ss := make([]ExchangeBackup, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, ExchangeBackup{exchange{sel}})
	}

	return ss
}

// NewExchangeRestore produces a new Selector with the service set to ServiceExchange.
func NewExchangeRestore(users []string) *ExchangeRestore {
	src := ExchangeRestore{
		exchange{
			newSelector(ServiceExchange, users),
		},
	}

	return &src
}

// ToExchangeRestore transforms the generic selector into an ExchangeRestore.
// Errors if the service defined by the selector is not ServiceExchange.
func (s Selector) ToExchangeRestore() (*ExchangeRestore, error) {
	if s.Service != ServiceExchange {
		return nil, badCastErr(ServiceExchange, s.Service)
	}

	src := ExchangeRestore{exchange{s}}

	return &src, nil
}

func (sr ExchangeRestore) SplitByResourceOwner(users []string) []ExchangeRestore {
	sels := splitByProtectedResource[ExchangeScope](sr.Selector, users, ExchangeUser)

	ss := make([]ExchangeRestore, 0, len(sels))
	for _, sel := range sels {
		ss = append(ss, ExchangeRestore{exchange{sel}})
	}

	return ss
}

// PathCategories produces the aggregation of discrete users described by each type of scope.
func (s exchange) PathCategories() selectorPathCategories {
	return selectorPathCategories{
		Excludes: pathCategoriesIn[ExchangeScope, exchangeCategory](s.Excludes),
		Filters:  pathCategoriesIn[ExchangeScope, exchangeCategory](s.Filters),
		Includes: pathCategoriesIn[ExchangeScope, exchangeCategory](s.Includes),
	}
}

// Reasons returns a deduplicated set of the backup reasons produced
// using the selector's discrete owner and each scopes' service and
// category types.
func (s exchange) Reasons(tenantID string, useOwnerNameForID bool) []identity.Reasoner {
	return reasonsFor(s, tenantID, useOwnerNameForID)
}

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

func (s ExchangeScope) Conceal() string             { return conceal(s) }
func (s ExchangeScope) Format(fs fmt.State, r rune) { format(s, fs, r) }
func (s ExchangeScope) String() string              { return conceal(s) }
func (s ExchangeScope) PlainString() string         { return plainString(s) }

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
// ex: User(u1) automatically cascades to all mail, events, and contacts,
// therefore it is the same as selecting all of the following:
// Mail(u1, Any(), Any()), Event(u1, Any()), Contacts(u1, Any(), Any())
func (s *exchange) Exclude(scopes ...[]ExchangeScope) {
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
// ex: User(u1) automatically cascades to all mail, events, and contacts,
// therefore it is the same as selecting all of the following:
// Mail(u1, Any(), Any()), Event(u1, Any()), Contacts(u1, Any(), Any())
func (s *exchange) Filter(scopes ...[]ExchangeScope) {
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
// ex: User(u1) automatically cascades to all mail, events, and contacts,
// therefore it is the same as selecting all of the following:
// Mail(u1, Any(), Any()), Event(u1, Any()), Contacts(u1, Any(), Any())
func (s *exchange) Include(scopes ...[]ExchangeScope) {
	s.Includes = appendScopes(s.Includes, scopes...)
}

// Scopes retrieves the list of exchangeScopes in the selector.
func (s *exchange) Scopes() []ExchangeScope {
	return scopes[ExchangeScope](s.Selector)
}

type ExchangeItemScopeConstructor func([]string, []string, ...option) []ExchangeScope

// -------------------
// Scope Factories

// Contacts produces one or more exchange contact scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *exchange) Contacts(folders, contacts []string, opts ...option) []ExchangeScope {
	scopes := []ExchangeScope{}

	scopes = append(
		scopes,
		makeScope[ExchangeScope](ExchangeContact, contacts, defaultItemOptions(s.Cfg)...).
			set(ExchangeContactFolder, folders, opts...))

	return scopes
}

// Contactfolders produces one or more exchange contact folder scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *exchange) ContactFolders(folders []string, opts ...option) []ExchangeScope {
	var (
		scopes = []ExchangeScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[ExchangeScope](ExchangeContactFolder, folders, os...))

	return scopes
}

// Events produces one or more exchange event scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *exchange) Events(calendars, events []string, opts ...option) []ExchangeScope {
	scopes := []ExchangeScope{}

	scopes = append(
		scopes,
		makeScope[ExchangeScope](ExchangeEvent, events, defaultItemOptions(s.Cfg)...).
			set(ExchangeEventCalendar, calendars, opts...))

	return scopes
}

// EventCalendars produces one or more exchange event calendar scopes.
// Calendars act as folders to contain Events
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *exchange) EventCalendars(events []string, opts ...option) []ExchangeScope {
	var (
		scopes = []ExchangeScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[ExchangeScope](ExchangeEventCalendar, events, os...))

	return scopes
}

// Produces one or more mail scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *exchange) Mails(folders, mails []string, opts ...option) []ExchangeScope {
	scopes := []ExchangeScope{}

	scopes = append(
		scopes,
		makeScope[ExchangeScope](ExchangeMail, mails, defaultItemOptions(s.Cfg)...).
			set(ExchangeMailFolder, folders, opts...))

	return scopes
}

// Produces one or more exchange mail folder scopes.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
// options are only applied to the folder scopes.
func (s *exchange) MailFolders(folders []string, opts ...option) []ExchangeScope {
	var (
		scopes = []ExchangeScope{}
		os     = append([]option{pathComparator()}, opts...)
	)

	scopes = append(
		scopes,
		makeScope[ExchangeScope](ExchangeMailFolder, folders, os...))

	return scopes
}

// Retrieves all exchange data.
// Each user id generates three scopes, one for each data type: contact, event, and mail.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *exchange) AllData() []ExchangeScope {
	scopes := []ExchangeScope{}

	scopes = append(scopes,
		makeScope[ExchangeScope](ExchangeContactFolder, Any()),
		makeScope[ExchangeScope](ExchangeEventCalendar, Any()),
		makeScope[ExchangeScope](ExchangeMailFolder, Any()))

	return scopes
}

// -------------------
// Info Factories

// ContactName produces one or more exchange contact name info scopes.
// Matches any contact whose name contains the provided string.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) ContactName(senderID string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeContact,
			ExchangeInfoContactName,
			[]string{senderID},
			filters.In),
	}
}

// EventSubject produces one or more exchange event subject info scopes.
// Matches any event where the event subject contains one of the provided strings.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) EventOrganizer(organizer string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeEvent,
			ExchangeInfoEventOrganizer,
			[]string{organizer},
			filters.In),
	}
}

// EventRecurs produces one or more exchange event recurrence info scopes.
// Matches any event if the comparator flag matches the event recurrence flag.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) EventRecurs(recurs string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeEvent,
			ExchangeInfoEventRecurs,
			[]string{recurs},
			filters.Equal),
	}
}

// EventStartsAfter produces an exchange event starts-after info scope.
// Matches any event where the start time is after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (sr *ExchangeRestore) EventStartsAfter(timeStrings string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeEvent,
			ExchangeInfoEventStartsAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

// EventStartsBefore produces an exchange event starts-before info scope.
// Matches any event where the start time is before the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (sr *ExchangeRestore) EventStartsBefore(timeStrings string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeEvent,
			ExchangeInfoEventStartsBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

// EventSubject produces one or more exchange event subject info scopes.
// Matches any event where the event subject contains one of the provided strings.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) EventSubject(subject string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeEvent,
			ExchangeInfoEventSubject,
			[]string{subject},
			filters.In),
	}
}

// MailReceivedAfter produces an exchange mail received-after info scope.
// Matches any mail which was received after the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (sr *ExchangeRestore) MailReceivedAfter(timeStrings string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeMail,
			ExchangeInfoMailReceivedAfter,
			[]string{timeStrings},
			filters.Less),
	}
}

// MailReceivedBefore produces an exchange mail received-before info scope.
// Matches any mail which was received before the timestring.
// If the input equals selectors.Any, the scope will match all times.
// If the input is empty or selectors.None, the scope will always fail comparisons.
func (sr *ExchangeRestore) MailReceivedBefore(timeStrings string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeMail,
			ExchangeInfoMailReceivedBefore,
			[]string{timeStrings},
			filters.Greater),
	}
}

// MailSender produces one or more exchange mail sender info scopes.
// Matches any mail whose sender contains one of the provided strings.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) MailSender(sender string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeMail,
			ExchangeInfoMailSender,
			[]string{sender},
			filters.In),
	}
}

// MailSubject produces one or more exchange mail subject line info scopes.
// Matches any mail whose subject contains one of the provided strings.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) MailSubject(subject string) []ExchangeScope {
	return []ExchangeScope{
		makeInfoScope[ExchangeScope](
			ExchangeMail,
			ExchangeInfoMailSubject,
			[]string{subject},
			filters.In),
	}
}

// ---------------------------------------------------------------------------
// Categories
// ---------------------------------------------------------------------------

// exchangeCategory enumerates the type of the lowest level
// of data specified by the scope.
type exchangeCategory string

// interface compliance checks
var _ categorizer = ExchangeCategoryUnknown

const (
	ExchangeCategoryUnknown exchangeCategory = ""

	// types of data identified by exchange
	ExchangeContact       exchangeCategory = "ExchangeContact"
	ExchangeContactFolder exchangeCategory = "ExchangeContactFolder"
	ExchangeEvent         exchangeCategory = "ExchangeEvent"
	ExchangeEventCalendar exchangeCategory = "ExchangeEventCalendar"
	ExchangeMail          exchangeCategory = "ExchangeMail"
	ExchangeMailFolder    exchangeCategory = "ExchangeMailFolder"
	ExchangeUser          exchangeCategory = "ExchangeUser"

	// data contained within details.ItemInfo
	ExchangeInfoMailSender         exchangeCategory = "ExchangeInfoMailSender"
	ExchangeInfoMailSubject        exchangeCategory = "ExchangeInfoMailSubject"
	ExchangeInfoMailReceivedAfter  exchangeCategory = "ExchangeInfoMailReceivedAfter"
	ExchangeInfoMailReceivedBefore exchangeCategory = "ExchangeInfoMailReceivedBefore"
	ExchangeInfoContactName        exchangeCategory = "ExchangeInfoContactName"
	ExchangeInfoEventOrganizer     exchangeCategory = "ExchangeInfoEventOrganizer"
	ExchangeInfoEventRecurs        exchangeCategory = "ExchangeInfoEventRecurs"
	ExchangeInfoEventStartsAfter   exchangeCategory = "ExchangeInfoEventStartsAfter"
	ExchangeInfoEventStartsBefore  exchangeCategory = "ExchangeInfoEventStartsBefore"
	ExchangeInfoEventSubject       exchangeCategory = "ExchangeInfoEventSubject"
)

// exchangeLeafProperties describes common metadata of the leaf categories
var exchangeLeafProperties = map[categorizer]leafProperty{
	ExchangeContact: {
		pathKeys: []categorizer{ExchangeContactFolder, ExchangeContact},
		pathType: path.ContactsCategory,
	},
	ExchangeEvent: {
		pathKeys: []categorizer{ExchangeEventCalendar, ExchangeEvent},
		pathType: path.EventsCategory,
	},
	ExchangeMail: {
		pathKeys: []categorizer{ExchangeMailFolder, ExchangeMail},
		pathType: path.EmailCategory,
	},
	ExchangeUser: { // the root category must be represented, even though it isn't a leaf
		pathKeys: []categorizer{ExchangeUser},
		pathType: path.UnknownCategory,
	},
}

func (ec exchangeCategory) String() string {
	return string(ec)
}

// leafCat returns the leaf category of the receiver.
// If the receiver category has multiple leaves (ex: User) or no leaves,
// (ex: Unknown), the receiver itself is returned.
// If the receiver category is an info type (ex: ExchangeInfoMailSubject),
// returns the category covered by the info.
// Ex: ExchangeContactFolder.leafCat() => ExchangeContact
// Ex: ExchangeEvent.leafCat() => ExchangeEvent
// Ex: ExchangeUser.leafCat() => ExchangeUser
func (ec exchangeCategory) leafCat() categorizer {
	switch ec {
	case ExchangeContact, ExchangeContactFolder, ExchangeInfoContactName:
		return ExchangeContact

	case ExchangeEvent, ExchangeEventCalendar, ExchangeInfoEventOrganizer, ExchangeInfoEventRecurs,
		ExchangeInfoEventStartsAfter, ExchangeInfoEventStartsBefore, ExchangeInfoEventSubject:
		return ExchangeEvent

	case ExchangeMail, ExchangeMailFolder, ExchangeInfoMailReceivedAfter,
		ExchangeInfoMailReceivedBefore, ExchangeInfoMailSender, ExchangeInfoMailSubject:
		return ExchangeMail
	}

	return ec
}

// rootCat returns the root category type.
func (ec exchangeCategory) rootCat() categorizer {
	return ExchangeUser
}

// unknownCat returns the unknown category type.
func (ec exchangeCategory) unknownCat() categorizer {
	return ExchangeCategoryUnknown
}

// isUnion returns true if c is a user
func (ec exchangeCategory) isUnion() bool {
	return ec == ec.rootCat()
}

// isLeaf is true if the category is a mail, event, or contact category.
func (ec exchangeCategory) isLeaf() bool {
	return ec == ec.leafCat()
}

// pathValues transforms the two paths to maps of identified properties.
//
// Example:
// [tenantID, service, userPN, category, mailFolder, mailID]
// => {exchMailFolder: mailFolder, exchMail: mailID}
func (ec exchangeCategory) pathValues(
	repo path.Path,
	ent details.Entry,
	cfg Config,
) (map[categorizer][]string, error) {
	var folderCat, itemCat categorizer

	switch ec {
	case ExchangeContact:
		folderCat, itemCat = ExchangeContactFolder, ExchangeContact

	case ExchangeEvent:
		folderCat, itemCat = ExchangeEventCalendar, ExchangeEvent

	case ExchangeMail:
		folderCat, itemCat = ExchangeMailFolder, ExchangeMail

	default:
		return nil, clues.New("bad exchanageCategory").With("category", ec)
	}

	item := ent.ItemRef
	if len(item) == 0 {
		item = repo.Item()
	}

	items := []string{ent.ShortRef, item}

	// only include the item ID when the user is NOT matching
	// item names. Exchange data does not contain an item name,
	// only an ID, and we don't want to mix up the two.
	if cfg.OnlyMatchItemNames {
		items = []string{ent.ShortRef}
	}

	// Will hit the if-condition when we're at a top-level folder, but we'll get
	// the same result when we extract from the RepoRef.
	folder := ent.LocationRef
	if len(folder) == 0 {
		folder = repo.Folder(true)
	}

	result := map[categorizer][]string{
		folderCat: {folder},
		itemCat:   items,
	}

	return result, nil
}

// pathKeys returns the path keys recognized by the receiver's leaf type.
func (ec exchangeCategory) pathKeys() []categorizer {
	return exchangeLeafProperties[ec.leafCat()].pathKeys
}

// PathType converts the category's leaf type into the matching path.CategoryType.
func (ec exchangeCategory) PathType() path.CategoryType {
	return exchangeLeafProperties[ec.leafCat()].pathType
}

// ---------------------------------------------------------------------------
// Scopes
// ---------------------------------------------------------------------------

// ExchangeScope specifies the data available
// when interfacing with the Exchange service.
type ExchangeScope scope

// interface compliance checks
var _ scoper = &ExchangeScope{}

// Category describes the type of the data in scope.
func (s ExchangeScope) Category() exchangeCategory {
	return exchangeCategory(getCategory(s))
}

// categorizer type is a generic wrapper around Category.
// Primarily used by scopes.go to for abstract comparisons.
func (s ExchangeScope) categorizer() categorizer {
	return s.Category()
}

// Matches returns true if the category is included in the scope's
// data type, and the target string matches that category's comparator.
func (s ExchangeScope) Matches(cat exchangeCategory, target string) bool {
	return matches(s, cat, target)
}

// InfoCategory returns the category enum of the scope info.
// If the scope is not an info type, returns ExchangeUnknownCategory.
func (s ExchangeScope) InfoCategory() exchangeCategory {
	return exchangeCategory(getInfoCategory(s))
}

// IncludeCategory checks whether the scope includes a certain category of data.
// Ex: to check if the scope includes mail data:
// s.IncludesCategory(selector.ExchangeMail)
func (s ExchangeScope) IncludesCategory(cat exchangeCategory) bool {
	return categoryMatches(s.Category(), cat)
}

// returns true if the category is included in the scope's data type,
// and the value is set to Any().
func (s ExchangeScope) IsAny(cat exchangeCategory) bool {
	return isAnyTarget(s, cat)
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// ExchangeUser category.
func (s ExchangeScope) Get(cat exchangeCategory) []string {
	return getCatValue(s, cat)
}

// sets a value by category to the scope.  Only intended for internal use.
func (s ExchangeScope) set(cat exchangeCategory, v []string, opts ...option) ExchangeScope {
	os := []option{}
	if cat == ExchangeContactFolder || cat == ExchangeEventCalendar || cat == ExchangeMailFolder {
		os = append(os, pathComparator())
	}

	return set(s, cat, v, append(os, opts...)...)
}

// setDefaults ensures that contact folder, mail folder, and user category
// scopes all express `AnyTgt` for their child category types.
func (s ExchangeScope) setDefaults() {
	switch s.Category() {
	case ExchangeContactFolder:
		s[ExchangeContact.String()] = passAny

	case ExchangeEventCalendar:
		s[ExchangeEvent.String()] = passAny

	case ExchangeMailFolder:
		s[ExchangeMail.String()] = passAny

	case ExchangeUser:
		s[ExchangeContactFolder.String()] = passAny
		s[ExchangeContact.String()] = passAny
		s[ExchangeEvent.String()] = passAny
		s[ExchangeMailFolder.String()] = passAny
		s[ExchangeMail.String()] = passAny
	}
}

// ---------------------------------------------------------------------------
// Backup Details Filtering
// ---------------------------------------------------------------------------

// Reduce filters the entries in a details struct to only those that match the
// inclusions, filters, and exclusions in the selector.
func (s exchange) Reduce(
	ctx context.Context,
	deets *details.Details,
	errs *fault.Bus,
) *details.Details {
	return reduce[ExchangeScope](
		ctx,
		deets,
		s.Selector,
		map[path.CategoryType]exchangeCategory{
			path.ContactsCategory: ExchangeContact,
			path.EventsCategory:   ExchangeEvent,
			path.EmailCategory:    ExchangeMail,
		},
		errs)
}

// matchesInfo handles the standard behavior when comparing a scope and an ExchangeInfo
// returns true if the scope and info match for the provided category.
func (s ExchangeScope) matchesInfo(dii details.ItemInfo) bool {
	info := dii.Exchange
	if info == nil {
		return false
	}

	infoCat := s.InfoCategory()

	cfpc := categoryFromItemType(info.ItemType)
	if !typeAndCategoryMatches(infoCat, cfpc) {
		return false
	}

	i := ""

	switch infoCat {
	case ExchangeInfoContactName:
		i = info.ContactName
	case ExchangeInfoEventOrganizer:
		i = info.Organizer
	case ExchangeInfoEventRecurs:
		i = strconv.FormatBool(info.EventRecurs)
	case ExchangeInfoEventStartsAfter, ExchangeInfoEventStartsBefore:
		i = dttm.Format(info.EventStart)
	case ExchangeInfoEventSubject:
		i = info.Subject
	case ExchangeInfoMailSender:
		i = info.Sender
	case ExchangeInfoMailSubject:
		i = info.Subject
	case ExchangeInfoMailReceivedAfter, ExchangeInfoMailReceivedBefore:
		i = dttm.Format(info.Received)
	}

	return s.Matches(infoCat, i)
}

// categoryFromItemType interprets the category represented by the ExchangeInfo
// struct.  Since every ExchangeInfo can hold all exchange data info, the exact
// type that the struct represents must be compared using its ItemType prop.
func categoryFromItemType(pct details.ItemType) exchangeCategory {
	switch pct {
	case details.ExchangeContact:
		return ExchangeContact
	case details.ExchangeMail:
		return ExchangeMail
	case details.ExchangeEvent:
		return ExchangeEvent
	}

	return ExchangeCategoryUnknown
}
