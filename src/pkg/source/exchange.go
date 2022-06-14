package source

import (
	"strconv"
)

// ExchangeSource provides an api for scoping
// data in the Exchange service.
type ExchangeSource struct {
	Source
}

func (s Source) ToExchange() (*ExchangeSource, error) {
	if s.service != ServiceExchange {
		return nil, BadCastErr(ServiceExchange, s.service)
	}
	src := ExchangeSource{s}
	return &src, nil
}

func NewExchange(tenantID string) *ExchangeSource {
	src := ExchangeSource{
		newSource(tenantID, ServiceExchange),
	}
	return &src
}

func (s *ExchangeSource) Scopes() []exchangeScope {
	scopes := []exchangeScope{}
	for _, v := range s.scopes {
		scopes = append(scopes, exchangeScope(v))
	}
	return scopes
}

// the following are called by the client to specify the constraints
// each call appends one or more scopes to the source.

// Users selects the specified users.  All of their data is included.
func (s *ExchangeSource) Users(us ...string) {
	// todo
}

// AllContactsForUser selects all contacts data for the user.
func (s *ExchangeSource) AllContactsForUser(u string) {
	// todo
}

// UsersContacts selects the specified contacts owned by the user.
func (s *ExchangeSource) UsersContacts(u string, cs ...string) {
	// todo
}

// AllEventsForUser selects all events for the user.
func (s *ExchangeSource) AllEventsForUser(u string) {
	// todo
}

// UsersEvents selects the specified events owned by the user.
func (s *ExchangeSource) UsersEvents(u string, cs ...string) {
	// todo
}

// AllMailForUser selects all mail folders and messages owned by the user.
func (s *ExchangeSource) AllMailForUser(u string) {
	// todo
}

// UsersMailFolders selects the specified mail folders owned by the user.
func (s *ExchangeSource) UsersMailFolders(u string, fs ...string) {
	// todo
}

// UsersMessages selects the specified mail messages within the given folder,
// owned by the user.
func (s *ExchangeSource) UsersMessages(u, f string, ms ...string) {
	// todo
}

// -----------------------

// exchangeScope specifies the data available
// when interfacing with the Exchange service.
type exchangeScope map[string]string

type exchangeCategory int

// exchangeCategory describes the type of data in scope.
const (
	ExchangeCategoryUnknown exchangeCategory = iota
	ExchangeContact
	ExchangeEvent
	ExchangeFolder
	ExchangeMail
	ExchangeUser
)

// String complies with the stringer interface, so that exchangeCategories
// can be added into the scope map.
func (ec exchangeCategory) String() string {
	return strconv.Itoa(int(ec))
}

var (
	exchangeScopeKeyContactID = ExchangeContact.String()
	exchangeScopeKeyEventID   = ExchangeEvent.String()
	exchangeScopeKeyFolderID  = ExchangeFolder.String()
	exchangeScopeKeyMessageID = ExchangeMail.String()
	exchangeScopeKeyUserID    = ExchangeUser.String()
)

// Category describes the type of the data in scope.
func (s exchangeScope) Category() exchangeCategory {
	return exchangeCategory(valAtoI(s, scopeKeyCategory))
}

// Granularity describes the breadth of data in scope.
func (s exchangeScope) Granularity() scopeGranularity {
	return granularityOf(s)
}

// FullPath returns the full path of data (as much as is known to the scope).
func (s exchangeScope) FullPath() string {
	return s[scopeKeyFullPath]
}

func (s exchangeScope) UserID() string {
	return s[exchangeScopeKeyUserID]
}

func (s exchangeScope) ContactID() string {
	return s[exchangeScopeKeyContactID]
}

func (s exchangeScope) EventID() string {
	return s[exchangeScopeKeyEventID]
}

func (s exchangeScope) FolderID() string {
	return s[exchangeScopeKeyFolderID]
}

func (s exchangeScope) MessageID() string {
	return s[exchangeScopeKeyMessageID]
}
