package selectors

import (
	"strings"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/backup"
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

// NewExchange produces a new Selector with the service set to ServiceExchange.
func NewExchangeBackup() *ExchangeBackup {
	src := ExchangeBackup{
		exchange{
			newSelector(ServiceExchange),
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

// NewExchangeRestore produces a new Selector with the service set to ServiceExchange.
func NewExchangeRestore() *ExchangeRestore {
	src := ExchangeRestore{
		exchange{
			newSelector(ServiceExchange),
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

// -------------------
// Exclude/Includes

// Include appends the provided scopes to the selector's inclusion set.
//
// All parts of the scope must match for data to be included.
// Ex: Mail(u1, f1, m1) => only includes mail if it is owned by user u1,
// located in folder f1, and ID'd as m1.  Use selectors.Any() to wildcard
// a scope value.  No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to child
// properties.
// ex: User(u1) is the same as Mail(u1, Any(), Any()).
func (s *exchange) Include(scopes ...[]exchangeScope) {
	if s.Includes == nil {
		s.Includes = []map[string]string{}
	}
	concat := []exchangeScope{}
	for _, scopeSl := range scopes {
		concat = append(concat, extendExchangeScopeValues(scopeSl)...)
	}
	for _, sc := range concat {
		s.Includes = append(s.Includes, map[string]string(sc))
	}
}

// Exclude appends the provided scopes to the selector's exclusion set.
// Every Exclusion scope applies globally, affecting all inclusion scopes.
//
// All parts of the scope must match for data to be excluded.
// Ex: Mail(u1, f1, m1) => only excludes mail that is owned by user u1,
// located in folder f1, and ID'd as m1.  Use selectors.Any() to wildcard
// a scope value.  No value will match if selectors.None() is provided.
//
// Group-level scopes will automatically apply the Any() wildcard to
// child properties.
// ex: User(u1) automatically includes all mail, events, and contacts,
// therefore it is the same as selecting all of the following:
// Mail(u1, Any(), Any()), Event(u1, Any()), Contacts(u1, Any(), Any())
func (s *exchange) Exclude(scopes ...[]exchangeScope) {
	if s.Excludes == nil {
		s.Excludes = []map[string]string{}
	}
	concat := []exchangeScope{}
	for _, scopeSl := range scopes {
		concat = append(concat, extendExchangeScopeValues(scopeSl)...)
	}
	for _, sc := range concat {
		s.Excludes = append(s.Excludes, map[string]string(sc))
	}
}

// completes population for certain scope properties, according to the
// expecations of Include and Exclude behavior.
func extendExchangeScopeValues(es []exchangeScope) []exchangeScope {
	v := join(Any()...)
	for i := range es {
		switch es[i].Category() {
		case ExchangeContactFolder:
			es[i][ExchangeContact.String()] = v
		case ExchangeMailFolder:
			es[i][ExchangeMail.String()] = v
		case ExchangeUser:
			es[i][ExchangeContactFolder.String()] = v
			es[i][ExchangeContact.String()] = v
			es[i][ExchangeEvent.String()] = v
			es[i][ExchangeMailFolder.String()] = v
			es[i][ExchangeMail.String()] = v
		}
	}
	return es
}

// -------------------
// Scope Factory

// Produces one or more exchange contact scopes.
// One scope is created per combination of users,folders,contacts.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *exchange) Contacts(users, folders, contacts []string) []exchangeScope {
	users = normalize(users)
	folders = normalize(folders)
	contacts = normalize(contacts)
	scopes := []exchangeScope{}
	for _, u := range users {
		for _, f := range folders {
			scopes = append(scopes, exchangeScope{
				scopeKeyGranularity:            Item,
				scopeKeyCategory:               ExchangeContact.String(),
				ExchangeUser.String():          u,
				ExchangeContactFolder.String(): f,
				ExchangeContact.String():       join(contacts...),
			})
		}
	}
	return scopes
}

// Produces one or more exchange contact folder scopes.
// One scope is created per combination of users,folders.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *exchange) ContactFolders(users, folders []string) []exchangeScope {
	users = normalize(users)
	folders = normalize(folders)
	scopes := []exchangeScope{}
	for _, u := range users {
		scopes = append(scopes, exchangeScope{
			scopeKeyGranularity:            Group,
			scopeKeyCategory:               ExchangeContactFolder.String(),
			ExchangeUser.String():          u,
			ExchangeContactFolder.String(): join(folders...),
		})
	}
	return scopes
}

// Produces one or more exchange event scopes.
// One scope is created per combination of users,events.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *exchange) Events(users, events []string) []exchangeScope {
	users = normalize(users)
	events = normalize(events)
	scopes := []exchangeScope{}
	for _, u := range users {
		scopes = append(scopes, exchangeScope{
			scopeKeyGranularity:    Item,
			scopeKeyCategory:       ExchangeEvent.String(),
			ExchangeUser.String():  u,
			ExchangeEvent.String(): join(events...),
		})
	}
	return scopes
}

// Produces one or more mail scopes.
// One scope is created per combination of users,folders,mails.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *exchange) Mails(users, folders, mails []string) []exchangeScope {
	users = normalize(users)
	folders = normalize(folders)
	mails = normalize(mails)
	scopes := []exchangeScope{}
	for _, u := range users {
		for _, f := range folders {
			scopes = append(scopes, exchangeScope{
				scopeKeyGranularity:         Item,
				scopeKeyCategory:            ExchangeMail.String(),
				ExchangeUser.String():       u,
				ExchangeMailFolder.String(): f,
				ExchangeMail.String():       join(mails...),
			})
		}
	}
	return scopes
}

// Produces one or more exchange mail folder scopes.
// One scope is created per combination of users,folders.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *exchange) MailFolders(users, folders []string) []exchangeScope {
	users = normalize(users)
	folders = normalize(folders)
	scopes := []exchangeScope{}
	for _, u := range users {
		scopes = append(scopes, exchangeScope{
			scopeKeyGranularity:         Group,
			scopeKeyCategory:            ExchangeMailFolder.String(),
			ExchangeUser.String():       u,
			ExchangeMailFolder.String(): join(folders...),
		})
	}
	return scopes
}

// Produces one or more exchange contact user scopes.
// One scope is created per user entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *exchange) Users(users []string) []exchangeScope {
	users = normalize(users)
	scopes := []exchangeScope{}
	scopes = append(scopes, exchangeScope{
		scopeKeyGranularity:   Group,
		scopeKeyCategory:      ExchangeUser.String(),
		ExchangeUser.String(): join(users...),
	})
	return scopes
}

// Produces one or more exchange mail sender filter scopes.
// Matches any mail whose sender is equal to one of the provided strings.
// One scope is created per senderID entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) MailSender(senderID []string) []exchangeScope {
	scopes := []exchangeScope{}
	scopes = append(scopes, exchangeScope{
		scopeKeyGranularity:             Item,
		scopeKeyCategory:                ExchangeMail.String(),
		scopeKeyInfoFilter:              ExchangeInfoMailSender.String(),
		ExchangeInfoMailSender.String(): join(senderID...),
	})
	return scopes
}

// Produces one or more exchange mail subject filter scopes.
// Matches any mail whose mail subject contains one of the provided strings.
// One scope is created per subject entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) MailSubject(subjectSubstring []string) []exchangeScope {
	scopes := []exchangeScope{}
	scopes = append(scopes, exchangeScope{
		scopeKeyGranularity:              Item,
		scopeKeyCategory:                 ExchangeMail.String(),
		scopeKeyInfoFilter:               ExchangeInfoMailSubject.String(),
		ExchangeInfoMailSubject.String(): join(subjectSubstring...),
	})
	return scopes
}

// Produces one or more exchange mail received-after filter scopes.
// Matches any mail which was received after the timestring.
// One scope is created per timeString entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) MailReceivedAfter(timeString []string) []exchangeScope {
	scopes := []exchangeScope{}
	scopes = append(scopes, exchangeScope{
		scopeKeyGranularity:                    Item,
		scopeKeyCategory:                       ExchangeMail.String(),
		scopeKeyInfoFilter:                     ExchangeInfoMailReceivedAfter.String(),
		ExchangeInfoMailReceivedAfter.String(): join(timeString...),
	})
	return scopes
}

// Produces one or more exchange mail received-before filter scopes.
// Matches any mail which was received before the timestring.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (sr *ExchangeRestore) MailReceivedBefore(timeString []string) []exchangeScope {
	scopes := []exchangeScope{}
	scopes = append(scopes, exchangeScope{
		scopeKeyGranularity:                     Item,
		scopeKeyCategory:                        ExchangeMail.String(),
		scopeKeyInfoFilter:                      ExchangeInfoMailReceivedBefore.String(),
		ExchangeInfoMailReceivedBefore.String(): join(timeString...),
	})
	return scopes
}

// ---------------------------------------------------------------------------
// Destination
// ---------------------------------------------------------------------------

type ExchangeDestination Destination

func NewExchangeDestination() ExchangeDestination {
	return ExchangeDestination{}
}

// GetOrDefault gets the destination of the provided category.  If no
// destination is set, returns the current value.
func (d ExchangeDestination) GetOrDefault(cat exchangeCategory, current string) string {
	dest, ok := d[cat.String()]
	if !ok {
		return current
	}
	return dest
}

// Sets the destination value of the provided category.  Returns an error
// if a destination is already declared for that category.
func (d ExchangeDestination) Set(cat exchangeCategory, dest string) error {
	if len(dest) == 0 {
		return nil
	}
	cs := cat.String()
	if curr, ok := d[cs]; ok {
		return existingDestinationErr(cs, curr)
	}
	d[cs] = dest
	return nil
}

// ---------------------------------------------------------------------------
// Scopes
// ---------------------------------------------------------------------------

type (
	// exchangeScope specifies the data available
	// when interfacing with the Exchange service.
	exchangeScope map[string]string
	// exchangeCategory enumerates the type of the lowest level
	// of data () in a scope.
	exchangeCategory int
)

// Scopes retrieves the list of exchangeScopes in the selector.
func (s *exchange) Scopes() []exchangeScope {
	scopes := []exchangeScope{}
	for _, v := range s.Includes {
		scopes = append(scopes, exchangeScope(v))
	}
	return scopes
}

//go:generate stringer -type=exchangeCategory
const (
	ExchangeCategoryUnknown exchangeCategory = iota
	// types of data identified by exchange
	ExchangeContact
	ExchangeContactFolder
	ExchangeEvent
	ExchangeMail
	ExchangeMailFolder
	ExchangeUser
	// filterable topics identified by exchange
	ExchangeInfoMailSender exchangeCategory = iota + 100 // offset to pad out future data additions
	ExchangeInfoMailSubject
	ExchangeInfoMailReceivedAfter
	ExchangeInfoMailReceivedBefore
)

func exchangeCatAtoI(s string) exchangeCategory {
	switch s {
	// data types
	case ExchangeContact.String():
		return ExchangeContact
	case ExchangeContactFolder.String():
		return ExchangeContactFolder
	case ExchangeEvent.String():
		return ExchangeEvent
	case ExchangeMail.String():
		return ExchangeMail
	case ExchangeMailFolder.String():
		return ExchangeMailFolder
	case ExchangeUser.String():
		return ExchangeUser
	// filters
	case ExchangeInfoMailSender.String():
		return ExchangeInfoMailSender
	case ExchangeInfoMailSubject.String():
		return ExchangeInfoMailSubject
	case ExchangeInfoMailReceivedAfter.String():
		return ExchangeInfoMailReceivedAfter
	case ExchangeInfoMailReceivedBefore.String():
		return ExchangeInfoMailReceivedBefore
	default:
		return ExchangeCategoryUnknown
	}
}

// Granularity describes the granularity (directory || item)
// of the data in scope
func (s exchangeScope) Granularity() string {
	return s[scopeKeyGranularity]
}

// Category describes the type of the data in scope.
func (s exchangeScope) Category() exchangeCategory {
	return exchangeCatAtoI(s[scopeKeyCategory])
}

// Filer describes the specific filter, and its target values.
func (s exchangeScope) Filter() exchangeCategory {
	return exchangeCatAtoI(s[scopeKeyInfoFilter])
}

// IncludeCategory checks whether the scope includes a
// certain category of data.
// Ex: to check if the scope includes mail data:
// s.IncludesCategory(selector.ExchangeMail)
func (s exchangeScope) IncludesCategory(cat exchangeCategory) bool {
	sCat := s.Category()
	if cat == ExchangeCategoryUnknown || sCat == ExchangeCategoryUnknown {
		return false
	}
	if cat == ExchangeUser || sCat == ExchangeUser {
		return true
	}
	switch sCat {
	case ExchangeContact, ExchangeContactFolder:
		return cat == ExchangeContact || cat == ExchangeContactFolder
	case ExchangeEvent:
		return cat == ExchangeEvent
	case ExchangeMail, ExchangeMailFolder:
		return cat == ExchangeMail || cat == ExchangeMailFolder
	}
	return false
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// ExchangeUser category.
func (s exchangeScope) Get(cat exchangeCategory) []string {
	v, ok := s[cat.String()]
	if !ok {
		return None()
	}
	return split(v)
}

var categoryPathSet = map[exchangeCategory][]exchangeCategory{
	ExchangeContact: {ExchangeUser, ExchangeContactFolder, ExchangeContact},
	ExchangeEvent:   {ExchangeUser, ExchangeEvent},
	ExchangeMail:    {ExchangeUser, ExchangeMailFolder, ExchangeMail},
	ExchangeUser:    {ExchangeUser},
}

// matches returns true if either the path or the info matches the scope details.
func (s exchangeScope) matches(cat exchangeCategory, path []string, info *backup.ExchangeInfo) bool {
	return s.matchesPath(cat, path) || s.matchesInfo(cat, info)
}

// matchesInfo handles the standard behavior when comparing a scope and an exchangeInfo
// returns true if the scope and info match for the provided category.
func (s exchangeScope) matchesInfo(cat exchangeCategory, info *backup.ExchangeInfo) bool {
	// we need values to match against
	if info == nil {
		return false
	}
	// the scope must define targets to match on
	filterCat := s.Filter()
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
		case ExchangeInfoMailSender:
			if target == info.Sender {
				return true
			}
		case ExchangeInfoMailSubject:
			if strings.Contains(info.Subject, target) {
				return true
			}
		case ExchangeInfoMailReceivedAfter:
			if target < common.FormatTime(info.Received) {
				return true
			}
		case ExchangeInfoMailReceivedBefore:
			if target > common.FormatTime(info.Received) {
				return true
			}
		}
	}
	return false
}

// matchesPath handles the standard behavior when comparing a scope and a path
// returns true if the scope and path match for the provided category.
func (s exchangeScope) matchesPath(cat exchangeCategory, path []string) bool {
	pathIDs := exchangeIDPath(cat, path)
	for _, c := range categoryPathSet[cat] {
		target := s.Get(c)
		// the scope must define the targets to match on
		if len(target) == 0 {
			return false
		}
		// None() fails all matches
		if target[0] == NoneTgt {
			return false
		}
		// the path must contain a value to match against
		id, ok := pathIDs[c]
		if !ok {
			return false
		}
		// all parts of the scope must match
		isAny := target[0] == AnyTgt
		if !isAny {
			if !contains(target, id) {
				return false
			}
		}
	}
	return true
}

// temporary helper until filters replace string values for scopes.
func contains(super []string, sub string) bool {
	for _, s := range super {
		if s == sub {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------------
// Restore Point Filtering
// ---------------------------------------------------------------------------

// transforms a path to a map of identified properties.
// Malformed (ie, short len) paths will return incomplete results.
// Example:
// [tenantID, userID, "mail", mailFolder, mailID]
// => {exchUser: userID, exchMailFolder: mailFolder, exchMail: mailID}
func exchangeIDPath(cat exchangeCategory, path []string) map[exchangeCategory]string {
	m := map[exchangeCategory]string{}
	if len(path) == 0 {
		return m
	}
	m[ExchangeUser] = path[1]
	/*
		TODO/Notice:
		Mail and Contacts contain folder structures, identified
		in this code as being at index 3.  This assumes a single
		folder, while in reality users can express subfolder
		hierarchies of arbirary depth.  Subfolder handling is coming
		at a later time.
	*/
	switch cat {
	case ExchangeContact:
		if len(path) < 5 {
			return m
		}
		m[ExchangeContactFolder] = path[3]
		m[ExchangeContact] = path[4]
	case ExchangeEvent:
		if len(path) < 4 {
			return m
		}
		m[ExchangeEvent] = path[3]
	case ExchangeMail:
		if len(path) < 5 {
			return m
		}
		m[ExchangeMailFolder] = path[3]
		m[ExchangeMail] = path[4]
	}
	return m
}

// FilterDetails reduces the entries in a backupDetails struct to only
// those that match the inclusions and exclusions in the selector.
func (s *ExchangeRestore) FilterDetails(deets *backup.Details) *backup.Details {
	if deets == nil {
		return nil
	}

	entIncs := exchangeScopesByCategory(s.Includes)
	entExcs := exchangeScopesByCategory(s.Excludes)

	ents := []backup.DetailsEntry{}

	for _, ent := range deets.Entries {
		// todo: use Path pkg for this
		path := strings.Split(ent.RepoRef, "/")
		// not all paths will be len=3.  Most should be longer.
		// This just protects us from panicing four lines later.
		if len(path) < 3 {
			continue
		}
		var cat exchangeCategory
		switch path[2] {
		case "contact":
			cat = ExchangeContact
		case "event":
			cat = ExchangeEvent
		case "mail":
			cat = ExchangeMail
		}
		matched := matchExchangeEntry(
			cat,
			path,
			ent.Exchange,
			entIncs[cat.String()],
			entExcs[cat.String()])
		if matched {
			ents = append(ents, ent)
		}
	}

	deets.Entries = ents
	return deets
}

// groups each scope by its category of data (contact, event, or mail).
// user-level scopes will duplicate to all three categories.
func exchangeScopesByCategory(scopes []map[string]string) map[string][]exchangeScope {
	m := map[string][]exchangeScope{
		ExchangeContact.String(): {},
		ExchangeEvent.String():   {},
		ExchangeMail.String():    {},
	}
	for _, msc := range scopes {
		sc := exchangeScope(msc)
		if sc.IncludesCategory(ExchangeContact) {
			m[ExchangeContact.String()] = append(m[ExchangeContact.String()], sc)
		}
		if sc.IncludesCategory(ExchangeEvent) {
			m[ExchangeEvent.String()] = append(m[ExchangeEvent.String()], sc)
		}
		if sc.IncludesCategory(ExchangeMail) {
			m[ExchangeMail.String()] = append(m[ExchangeMail.String()], sc)
		}
	}
	return m
}

// compare each path to the included and excluded exchange scopes.  Returns true
// if the path is included, and not excluded.
func matchExchangeEntry(
	cat exchangeCategory,
	path []string,
	info *backup.ExchangeInfo,
	incs, excs []exchangeScope,
) bool {
	var included bool
	for _, inc := range incs {
		if inc.matches(cat, path, info) {
			included = true
			break
		}
	}
	if !included {
		return false
	}

	var excluded bool
	for _, exc := range excs {
		if exc.matches(cat, path, info) {
			excluded = true
			break
		}
	}

	return !excluded
}
