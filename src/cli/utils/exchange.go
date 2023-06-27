package utils

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type ExchangeOpts struct {
	Users []string

	Contact       []string
	ContactFolder []string
	ContactName   string

	Email               []string
	EmailFolder         []string
	EmailReceivedAfter  string
	EmailReceivedBefore string
	EmailSender         string
	EmailSubject        string

	Event             []string
	EventCalendar     []string
	EventOrganizer    string
	EventRecurs       string
	EventStartsAfter  string
	EventStartsBefore string
	EventSubject      string

	Populated flags.PopulatedFlags
}

// populates an ExchangeOpts struct with the command's current flags.
func MakeExchangeOpts(cmd *cobra.Command) ExchangeOpts {
	return ExchangeOpts{
		Users: flags.UserFV,

		Contact:       flags.ContactFV,
		ContactFolder: flags.ContactFolderFV,
		ContactName:   flags.ContactNameFV,

		Email:               flags.EmailFV,
		EmailFolder:         flags.EmailFolderFV,
		EmailReceivedAfter:  flags.EmailReceivedAfterFV,
		EmailReceivedBefore: flags.EmailReceivedBeforeFV,
		EmailSender:         flags.EmailSenderFV,
		EmailSubject:        flags.EmailSubjectFV,

		Event:             flags.EventFV,
		EventCalendar:     flags.EventCalendarFV,
		EventOrganizer:    flags.EventOrganizerFV,
		EventRecurs:       flags.EventRecursFV,
		EventStartsAfter:  flags.EventStartsAfterFV,
		EventStartsBefore: flags.EventStartsBeforeFV,
		EventSubject:      flags.EventSubjectFV,

		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// AddExchangeInclude adds the scope of the provided values to the selector's
// inclusion set.  Any unpopulated slice will be replaced with selectors.Any()
// to act as a wildcard.
func AddExchangeInclude(
	sel *selectors.ExchangeRestore,
	folders, items []string,
	eisc selectors.ExchangeItemScopeConstructor,
) {
	lf, li := len(folders), len(items)

	// only use the inclusion if either a folder or item of
	// this type is specified.
	if lf+li == 0 {
		return
	}

	if li == 0 {
		items = selectors.Any()
	}

	containsFolders, prefixFolders := splitFoldersIntoContainsAndPrefix(folders)

	if len(containsFolders) > 0 {
		sel.Include(eisc(containsFolders, items))
	}

	if len(prefixFolders) > 0 {
		sel.Include(eisc(prefixFolders, items, selectors.PrefixMatch()))
	}
}

// AddExchangeInfo adds the scope of the provided values to the selector's
// filter set
func AddExchangeInfo(
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
func ValidateExchangeRestoreFlags(backupID string, opts ExchangeOpts) error {
	if len(backupID) == 0 {
		return clues.New("a backup ID is required")
	}

	if _, ok := opts.Populated[flags.EmailReceivedAfterFN]; ok && !IsValidTimeFormat(opts.EmailReceivedAfter) {
		return clues.New("invalid time format for email-received-after")
	}

	if _, ok := opts.Populated[flags.EmailReceivedBeforeFN]; ok && !IsValidTimeFormat(opts.EmailReceivedBefore) {
		return clues.New("invalid time format for email-received-before")
	}

	if _, ok := opts.Populated[flags.EventStartsAfterFN]; ok && !IsValidTimeFormat(opts.EventStartsAfter) {
		return clues.New("invalid time format for event-starts-after")
	}

	if _, ok := opts.Populated[flags.EventStartsBeforeFN]; ok && !IsValidTimeFormat(opts.EventStartsBefore) {
		return clues.New("invalid time format for event-starts-before")
	}

	if _, ok := opts.Populated[flags.EventRecursFN]; ok && !IsValidBool(opts.EventRecurs) {
		return clues.New("invalid format for event-recurs")
	}

	return nil
}

// IncludeExchangeRestoreDataSelectors builds the common data-selector
// inclusions for exchange commands.
func IncludeExchangeRestoreDataSelectors(opts ExchangeOpts) *selectors.ExchangeRestore {
	users := opts.Users
	if len(users) == 0 {
		users = selectors.Any()
	}

	sel := selectors.NewExchangeRestore(users)

	lc, lcf := len(opts.Contact), len(opts.ContactFolder)
	le, lef := len(opts.Email), len(opts.EmailFolder)
	lev, lec := len(opts.Event), len(opts.EventCalendar)
	// either scope the request to a set of users
	if lc+lcf+le+lef+lev+lec == 0 {
		sel.Include(sel.AllData())
		return sel
	}

	opts.EmailFolder = trimFolderSlash(opts.EmailFolder)

	// or add selectors for each type of data
	AddExchangeInclude(sel, opts.ContactFolder, opts.Contact, sel.Contacts)
	AddExchangeInclude(sel, opts.EmailFolder, opts.Email, sel.Mails)
	AddExchangeInclude(sel, opts.EventCalendar, opts.Event, sel.Events)

	return sel
}

// FilterExchangeRestoreInfoSelectors builds the common info-selector filters.
func FilterExchangeRestoreInfoSelectors(
	sel *selectors.ExchangeRestore,
	opts ExchangeOpts,
) {
	AddExchangeInfo(sel, opts.ContactName, sel.ContactName)
	AddExchangeInfo(sel, opts.EmailReceivedAfter, sel.MailReceivedAfter)
	AddExchangeInfo(sel, opts.EmailReceivedBefore, sel.MailReceivedBefore)
	AddExchangeInfo(sel, opts.EmailSender, sel.MailSender)
	AddExchangeInfo(sel, opts.EmailSubject, sel.MailSubject)
	AddExchangeInfo(sel, opts.EventOrganizer, sel.EventOrganizer)
	AddExchangeInfo(sel, opts.EventRecurs, sel.EventRecurs)
	AddExchangeInfo(sel, opts.EventStartsAfter, sel.EventStartsAfter)
	AddExchangeInfo(sel, opts.EventStartsBefore, sel.EventStartsBefore)
	AddExchangeInfo(sel, opts.EventSubject, sel.EventSubject)
}
