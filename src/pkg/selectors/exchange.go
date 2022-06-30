package selectors

import (
	"strconv"
)

// Exchange provides an api for scoping
// data in the Exchange service.
type Exchange struct {
	Selector
}

// ToExchange transforms the generic selector into an Exchange.
// Errors if the service defined by the selector is not ServiceExchange.
func (s Selector) ToExchange() (*Exchange, error) {
	if s.service != ServiceExchange {
		return nil, badCastErr(ServiceExchange, s.service)
	}
	src := Exchange{s}
	return &src, nil
}

// NewExchange produces a new Selector with the service set to ServiceExchange.
func NewExchange(tenantID string) *Exchange {
	src := Exchange{
		newSelector(tenantID, ServiceExchange),
	}
	return &src
}

// Scopes retrieves the list of exchangeScopes in the selector.
func (s *Exchange) Scopes() []exchangeScope {
	scopes := []exchangeScope{}
	for _, v := range s.scopes {
		scopes = append(scopes, exchangeScope(v))
	}
	return scopes
}

// the following are called by the client to specify the constraints
// each call appends one or more scopes to the selector.

// Users selects the specified users.  All of their data is included.
func (s *Exchange) Users(us ...string) {
	// todo
}

// Contacts selects the specified contacts owned by the user.
func (s *Exchange) Contacts(u string, vs ...string) {
	// todo
}

// Events selects the specified events owned by the user.
func (s *Exchange) Events(u string, vs ...string) {
	// todo
}

// MailFolders selects the specified mail folders owned by the user.
func (s *Exchange) MailFolders(u string, vs ...string) {
	// todo
}

// MailMessages selects the specified mail messages within the given folder,
// owned by the user.
func (s *Exchange) MailMessages(u, f string, vs ...string) {
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
	return exchangeCategory(getIota(s, scopeKeyCategory))
}

// Granularity describes the breadth of data in scope.
func (s exchangeScope) Granularity() scopeGranularity {
	return granularityOf(s)
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
