package selectors

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

// -------------------
// Exclude/Includes

func contactScope(u, f string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyCategory:               ExchangeContact.String(),
		ExchangeUser.String():          u,
		ExchangeContactFolder.String(): f,
		ExchangeContact.String():       join(vs...),
	}
}

// ExcludeContacts selects the specified contacts owned by the user.
func (s *exchange) ExcludeContacts(user, folder string, vs ...string) {
	s.Excludes = append(s.Excludes, contactScope(user, folder, vs...))
}

// IncludeContacts selects the specified contacts owned by the user.
func (s *exchange) IncludeContacts(user, folder string, vs ...string) {
	s.Includes = append(s.Includes, contactScope(user, folder, vs...))
}

func contactFolderScope(u string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyCategory:               ExchangeContactFolder.String(),
		ExchangeUser.String():          u,
		ExchangeContactFolder.String(): join(vs...),
	}
}

// IncludeContactFolders selects the specified contactFolders owned by the user.
func (s *exchange) IncludeContactFolders(user string, vs ...string) {
	s.Includes = append(s.Includes, contactFolderScope(user, vs...))
}

// ExcludeContactFolders selects the specified contactFolders owned by the user.
func (s *exchange) ExcludeContactFolders(user string, vs ...string) {
	s.Excludes = append(s.Excludes, contactFolderScope(user, vs...))
}

func eventScope(u string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyCategory:       ExchangeEvent.String(),
		ExchangeUser.String():  u,
		ExchangeEvent.String(): join(vs...),
	}
}

// ExcludeEvents selects the specified events owned by the user.
func (s *exchange) ExcludeEvents(user string, vs ...string) {
	s.Excludes = append(s.Excludes, eventScope(user, vs...))
}

// IncludeEvents selects the specified events owned by the user.
func (s *exchange) IncludeEvents(user string, vs ...string) {
	s.Includes = append(s.Includes, eventScope(user, vs...))
}

func mailScope(u, f string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyCategory:            ExchangeMail.String(),
		ExchangeUser.String():       u,
		ExchangeMailFolder.String(): f,
		ExchangeMail.String():       join(vs...),
	}
}

// ExcludeMail selects the specified mail messages within the given folder,
// owned by the user.
func (s *exchange) ExcludeMail(user, folder string, vs ...string) {
	s.Excludes = append(s.Excludes, mailScope(user, folder, vs...))
}

// IncludeMail selects the specified mail messages within the given folder,
// owned by the user.
func (s *exchange) IncludeMail(user, folder string, vs ...string) {
	s.Includes = append(s.Includes, mailScope(user, folder, vs...))
}

func mailFolderScope(u string, vs ...string) map[string]string {
	return map[string]string{
		scopeKeyCategory:            ExchangeMailFolder.String(),
		ExchangeUser.String():       u,
		ExchangeMailFolder.String(): join(vs...),
	}
}

// ExcludeMailFolders selects the specified mail folders owned by the user.
func (s *exchange) ExcludeMailFolders(user string, vs ...string) {
	s.Excludes = append(s.Excludes, mailFolderScope(user, vs...))
}

// IncludeMailFolders selects the specified mail folders owned by the user.
func (s *exchange) IncludeMailFolders(user string, vs ...string) {
	s.Includes = append(s.Includes, mailFolderScope(user, vs...))
}

func userScope(vs ...string) map[string]string {
	return map[string]string{
		scopeKeyCategory:      ExchangeUser.String(),
		ExchangeUser.String(): join(vs...),
	}
}

// ExcludeUsers selects the specified users.  All of their data is excluded.
func (s *exchange) ExcludeUsers(vs ...string) {
	s.Excludes = append(s.Excludes, userScope(vs...))
}

// IncludeUsers selects the specified users.  All of their data is included.
func (s *exchange) IncludeUsers(vs ...string) {
	s.Includes = append(s.Includes, userScope(vs...))
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
