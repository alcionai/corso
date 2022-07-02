package selectors

import (
	"strings"
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
func NewExchangeRestore(restorePointID string) *ExchangeRestore {
	src := ExchangeRestore{
		exchange{
			newSelector(ServiceExchange, restorePointID),
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

// IncludeContacts selects the specified contacts owned by the user.
func (s *exchange) IncludeContacts(u string, vs ...string) {
	// todo
}

// IncludeContactFolders selects the specified contactFolders owned by the user.
func (s *exchange) IncludeContactFolders(u string, vs ...string) {
	// todo
}

// IncludeEvents selects the specified events owned by the user.
func (s *exchange) IncludeEvents(u string, vs ...string) {
	// todo
}

// IncludeMail selects the specified mail messages within the given folder,
// owned by the user.
func (s *exchange) IncludeMail(u, f string, vs ...string) {
	// todo
}

// IncludeMailFolders selects the specified mail folders owned by the user.
func (s *exchange) IncludeMailFolders(u string, vs ...string) {
	// todo
}

// IncludeUsers selects the specified users.  All of their data is included.
func (s *exchange) IncludeUsers(us ...string) {
	// todo
}

// ExcludeContacts selects the specified contacts owned by the user.
func (s *exchange) ExcludeContacts(u string, vs ...string) {
	// todo
}

// ExcludeContactFolders selects the specified contactFolders owned by the user.
func (s *exchange) ExcludeContactFolders(u string, vs ...string) {
	// todo
}

// ExcludeEvents selects the specified events owned by the user.
func (s *exchange) ExcludeEvents(u string, vs ...string) {
	// todo
}

// ExcludeMail selects the specified mail messages within the given folder,
// owned by the user.
func (s *exchange) ExcludeMail(u, f string, vs ...string) {
	// todo
}

// ExcludeMailFolders selects the specified mail folders owned by the user.
func (s *exchange) ExcludeMailFolders(u string, vs ...string) {
	// todo
}

// ExcludeUsers selects the specified users.  All of their data is excluded.
func (s *exchange) ExcludeUsers(us ...string) {
	// todo
}

// ---------------------------------------------------------------------------
// Destination
// ---------------------------------------------------------------------------

type ExchangeDestination Destination

func NewExchangeDestination() ExchangeDestination {
	return ExchangeDestination{}
}

// GetsOrDefault gets the destination of the provided category.  If no
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
	if cat == ExchangeUser {
		return true
	}
	switch sCat {
	case ExchangeUser:
		return true
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
	return strings.Split(v, ",")
}
