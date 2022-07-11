package selectors

import (
	"strings"

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
			newSelector(ServiceExchange, ""),
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
func NewExchangeRestore(backupID string) *ExchangeRestore {
	src := ExchangeRestore{
		exchange{
			newSelector(ServiceExchange, backupID),
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
func (s *exchange) Include(scopes ...exchangeScope) {
	if s.Includes == nil {
		s.Includes = []map[string]string{}
	}
	for _, sc := range scopes {
		sc = extendExchangeScopeValues(All, sc)
		s.Includes = append(s.Includes, map[string]string(sc))
	}
}

// Exclude appends the provided scopes to the selector's exclusion set.
// Every Exclusion scope applies globally, affecting all inclusion scopes.
func (s *exchange) Exclude(scopes ...exchangeScope) {
	if s.Excludes == nil {
		s.Excludes = []map[string]string{}
	}
	for _, sc := range scopes {
		sc = extendExchangeScopeValues(None, sc)
		s.Excludes = append(s.Excludes, map[string]string(sc))
	}
}

// completes population for certain scope properties, according to the
// expecations of Include and Exclude behavior.
func extendExchangeScopeValues(v string, sc exchangeScope) exchangeScope {
	switch sc.Category() {
	case ExchangeContactFolder:
		sc[ExchangeContact.String()] = v
	case ExchangeMailFolder:
		sc[ExchangeMail.String()] = v
	case ExchangeUser:
		sc[ExchangeContactFolder.String()] = v
		sc[ExchangeContact.String()] = v
		sc[ExchangeEvent.String()] = v
		sc[ExchangeMailFolder.String()] = v
		sc[ExchangeMail.String()] = v
	}
	return sc
}

// -------------------
// Scope Factory

func (s *exchange) Contacts(u, f string, vs ...string) exchangeScope {
	return exchangeScope{
		scopeKeyGranularity:            Item,
		scopeKeyCategory:               ExchangeContact.String(),
		ExchangeUser.String():          u,
		ExchangeContactFolder.String(): f,
		ExchangeContact.String():       join(vs...),
	}
}

func (s *exchange) ContactFolders(u string, vs ...string) exchangeScope {
	return exchangeScope{
		scopeKeyGranularity:            Group,
		scopeKeyCategory:               ExchangeContactFolder.String(),
		ExchangeUser.String():          u,
		ExchangeContactFolder.String(): join(vs...),
	}
}

func (s *exchange) Events(u string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyGranularity:    Item,
		scopeKeyCategory:       ExchangeEvent.String(),
		ExchangeUser.String():  u,
		ExchangeEvent.String(): join(vs...),
	}
}

func (s *exchange) Mails(u, f string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyGranularity:         Item,
		scopeKeyCategory:            ExchangeMail.String(),
		ExchangeUser.String():       u,
		ExchangeMailFolder.String(): f,
		ExchangeMail.String():       join(vs...),
	}
}

func (s *exchange) MailFolders(u string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyGranularity:         Group,
		scopeKeyCategory:            ExchangeMailFolder.String(),
		ExchangeUser.String():       u,
		ExchangeMailFolder.String(): join(vs...),
	}
}

func (s *exchange) Users(vs ...string) map[string]string {
	return map[string]string{
		scopeKeyGranularity:   Group,
		scopeKeyCategory:      ExchangeUser.String(),
		ExchangeUser.String(): join(vs...),
	}
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
	ExchangeContact
	ExchangeContactFolder
	ExchangeEvent
	ExchangeMail
	ExchangeMailFolder
	ExchangeUser
)

func exchangeCatAtoI(s string) exchangeCategory {
	switch s {
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
		return []string{None}
	}
	return split(v)
}

var categoryPathSet = map[exchangeCategory][]exchangeCategory{
	ExchangeContact: {ExchangeUser, ExchangeContactFolder, ExchangeContact},
	ExchangeEvent:   {ExchangeUser, ExchangeEvent},
	ExchangeMail:    {ExchangeUser, ExchangeMailFolder, ExchangeMail},
}

// includesPath returns true if all filters in the scope match the path.
func (s exchangeScope) includesPath(cat exchangeCategory, path []string) bool {
	ids := idPath(cat, path)
	for _, c := range categoryPathSet[cat] {
		target := s.Get(c)
		if len(target) == 0 {
			return false
		}
		id, ok := ids[c]
		if !ok {
			return false
		}
		if target[0] != All && !contains(target, id) {
			return false
		}
	}
	return true
}

// excludesPath returns true if all filters in the scope match the path.
func (s exchangeScope) excludesPath(cat exchangeCategory, path []string) bool {
	ids := idPath(cat, path)
	for _, c := range categoryPathSet[cat] {
		target := s.Get(c)
		if len(target) == 0 {
			return true
		}
		id, ok := ids[c]
		if !ok {
			return true
		}
		if target[0] == All || contains(target, id) {
			return true
		}
	}
	return false
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
func idPath(cat exchangeCategory, path []string) map[exchangeCategory]string {
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
func (s *ExchangeRestore) FilterDetails(deets *backup.Details) []string {
	if deets == nil {
		return []string{}
	}

	entIncs := exchangeScopesByCategory(s.Includes)
	entExcs := exchangeScopesByCategory(s.Excludes)

	refs := []string{}

	for _, ent := range deets.Entries {
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
			entIncs[cat.String()],
			entExcs[cat.String()])
		if matched {
			refs = append(refs, ent.RepoRef)
		}
	}

	return refs
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
func matchExchangeEntry(cat exchangeCategory, path []string, incs, excs []exchangeScope) bool {
	var included bool
	for _, inc := range incs {
		if inc.includesPath(cat, path) {
			included = true
			break
		}
	}
	if !included {
		return false
	}

	var excluded bool
	for _, exc := range excs {
		if exc.excludesPath(cat, path) {
			excluded = true
			break
		}
	}

	return !excluded
}
