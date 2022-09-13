package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
)

// AddExchangeInclude adds the scope of the provided values to the selector's
// inclusion set.  Any unpopulated slice will be replaced with selectors.Any()
// to act as a wildcard.
func AddExchangeInclude(
	sel *selectors.ExchangeRestore,
	resource, folders, items []string,
	incl func([]string, []string, []string) []selectors.ExchangeScope,
) {
	lf, li := len(folders), len(items)

	// only use the inclusion if either a folder or item of
	// this type is specified.
	if lf+li == 0 {
		return
	}

	if len(resource) == 0 {
		resource = selectors.Any()
	}

	if lf == 0 {
		folders = selectors.Any()
	}

	if li == 0 {
		items = selectors.Any()
	}

	sel.Include(incl(resource, folders, items))
}

// AddExchangeFilter adds the scope of the provided values to the selector's
// filter set
func AddExchangeFilter(
	sel *selectors.ExchangeRestore,
	v string,
	f func(string) []selectors.ExchangeScope,
) {
	if len(v) == 0 {
		return
	}

	sel.Filter(f(v))
}

// ValidateExchangeRestoreFlags checks common flags for correctness and interdependencies
func ValidateExchangeRestoreFlags(backupID string) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}

	return nil
}

// IncludeExchangeRestoreDataSelectors builds the common data-selector
// inclusions for exchange commands.
func IncludeExchangeRestoreDataSelectors(
	sel *selectors.ExchangeRestore,
	contacts, contactFolders, emails, emailFolders, events, eventCalendars, users []string,
) {
	lc, lcf := len(contacts), len(contactFolders)
	le, lef := len(emails), len(emailFolders)
	lev, lec := len(events), len(eventCalendars)
	// either scope the request to a set of users
	if lc+lcf+le+lef+lev+lec == 0 {
		if len(users) == 0 {
			users = selectors.Any()
		}

		sel.Include(sel.Users(users))

		return
	}

	// or add selectors for each type of data
	AddExchangeInclude(sel, users, contactFolders, contacts, sel.Contacts)
	AddExchangeInclude(sel, users, emailFolders, emails, sel.Mails)
	AddExchangeInclude(sel, users, eventCalendars, events, sel.Events)
}

// FilterExchangeRestoreInfoSelectors builds the common info-selector filters.
func FilterExchangeRestoreInfoSelectors(
	sel *selectors.ExchangeRestore,
	contactName,
	emailReceivedAfter, emailReceivedBefore, emailSender, emailSubject,
	eventOrganizer, eventRecurs, eventStartsAfter, eventStartsBefore, eventSubject string,
) {
	AddExchangeFilter(sel, contactName, sel.ContactName)
	AddExchangeFilter(sel, emailReceivedAfter, sel.MailReceivedAfter)
	AddExchangeFilter(sel, emailReceivedBefore, sel.MailReceivedBefore)
	AddExchangeFilter(sel, emailSender, sel.MailSender)
	AddExchangeFilter(sel, emailSubject, sel.MailSubject)
	AddExchangeFilter(sel, eventOrganizer, sel.EventOrganizer)
	AddExchangeFilter(sel, eventRecurs, sel.EventRecurs)
	AddExchangeFilter(sel, eventStartsAfter, sel.EventStartsAfter)
	AddExchangeFilter(sel, eventStartsBefore, sel.EventStartsBefore)
	AddExchangeFilter(sel, eventSubject, sel.EventSubject)
}
