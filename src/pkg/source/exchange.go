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

type exchangeGranularity int

// exchangeGranularity specifies the granularity (ie, whether the data involves
// a wildcard, and at what depth of nested items) of the data in the scope.
//
//	Examples:
// 	AllUserData is the same as a FullPath like /tenantID/userID/*
//	MailFolder is equivalent to /tenantID/userID/mail/folderID/*
//  SingleMailMesssage is equivalent to /tenantID/userID/mail/folderID/messageID
const (
	UnknownEG exchangeGranularity = iota
	AllUserDataEG
	AllContactsEG
	SingleContactEG
	AllEventsEG
	SingleEventEG
	AllMailEG
	MailFolderEG
	SingleMailMessageEG
)

const (
	exchangeScopeKeyUserID    = "u"
	exchangeScopeKeyContactID = "c"
	exchangeScopeKeyEventID   = "e"
	exchangeScopeKeyFolderID  = "f"
	exchangeScopeKeyMessageID = "m"
)

// String complies with the stringer interface, so that exchangeGranularities
// can be added into the scope map.
func (eg exchangeGranularity) String() string {
	return strconv.Itoa(int(eg))
}

// Granularity describes the granularity of the data scope.
func (s exchangeScope) Granularity() exchangeGranularity {
	g := s[scopeKeyGranularity]
	v, err := strconv.Atoi(g)
	if err != nil {
		return UnknownEG
	}
	return exchangeGranularity(v)
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
