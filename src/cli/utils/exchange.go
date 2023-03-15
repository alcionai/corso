package utils

import (
	"errors"

	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/spf13/cobra"
)

// flag names
const (
	ContactFN       = "contact"
	ContactFolderFN = "contact-folder"
	ContactNameFN   = "contact-name"

	EmailFN               = "email"
	EmailFolderFN         = "email-folder"
	EmailReceivedAfterFN  = "email-received-after"
	EmailReceivedBeforeFN = "email-received-before"
	EmailSenderFN         = "email-sender"
	EmailSubjectFN        = "email-subject"

	EventFN             = "event"
	EventCalendarFN     = "event-calendar"
	EventOrganizerFN    = "event-organizer"
	EventRecursFN       = "event-recurs"
	EventStartsAfterFN  = "event-starts-after"
	EventStartsBeforeFN = "event-starts-before"
	EventSubjectFN      = "event-subject"
)

// flag population values
var (
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
)

type ExchangeOpts struct {
	Contact             []string
	ContactFolder       []string
	Email               []string
	EmailFolder         []string
	Event               []string
	EventCalendar       []string
	Users               []string
	ContactName         string
	EmailReceivedAfter  string
	EmailReceivedBefore string
	EmailSender         string
	EmailSubject        string
	EventOrganizer      string
	EventRecurs         string
	EventStartsAfter    string
	EventStartsBefore   string
	EventSubject        string

	Populated PopulatedFlags
}

// AddExchangeDetailsAndRestoreFlags adds flags that are common to both the
// details and restore commands.
func AddExchangeDetailsAndRestoreFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	// email flags
	fs.StringSliceVar(
		&Email,
		EmailFN, nil,
		"Select emails by email ID; accepts '"+Wildcard+"' to select all emails.")
	fs.StringSliceVar(
		&EmailFolder,
		EmailFolderFN, nil,
		"Select emails within a folder; accepts '"+Wildcard+"' to select all email folders.")
	fs.StringVar(
		&EmailSubject,
		EmailSubjectFN, "",
		"Select emails with a subject containing this value.")
	fs.StringVar(
		&EmailSender,
		EmailSenderFN, "",
		"Select emails from a specific sender.")
	fs.StringVar(
		&EmailReceivedAfter,
		EmailReceivedAfterFN, "",
		"Select emails received after this datetime.")
	fs.StringVar(
		&EmailReceivedBefore,
		EmailReceivedBeforeFN, "",
		"Select emails received before this datetime.")

	// event flags
	fs.StringSliceVar(
		&Event,
		EventFN, nil,
		"Select events by event ID; accepts '"+Wildcard+"' to select all events.")
	fs.StringSliceVar(
		&EventCalendar,
		EventCalendarFN, nil,
		"Select events under a calendar; accepts '"+Wildcard+"' to select all events.")
	fs.StringVar(
		&EventSubject,
		EventSubjectFN, "",
		"Select events with a subject containing this value.")
	fs.StringVar(
		&EventOrganizer,
		EventOrganizerFN, "",
		"Select events from a specific organizer.")
	fs.StringVar(
		&EventRecurs,
		EventRecursFN, "",
		"Select recurring events. Use `--event-recurs false` to select non-recurring events.")
	fs.StringVar(
		&EventStartsAfter,
		EventStartsAfterFN, "",
		"Select events starting after this datetime.")
	fs.StringVar(
		&EventStartsBefore,
		EventStartsBeforeFN, "",
		"Select events starting before this datetime.")

	// contact flags
	fs.StringSliceVar(
		&Contact,
		ContactFN, nil,
		"Select contacts by contact ID; accepts '"+Wildcard+"' to select all contacts.")
	fs.StringSliceVar(
		&ContactFolder,
		ContactFolderFN, nil,
		"Select contacts within a folder; accepts '"+Wildcard+"' to select all contact folders.")
	fs.StringVar(
		&ContactName,
		ContactNameFN, "",
		"Select contacts whose contact name contains this value.")
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
func ValidateExchangeRestoreFlags(backupID string, opts ExchangeOpts) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}

	if _, ok := opts.Populated[EmailReceivedAfterFN]; ok && !IsValidTimeFormat(opts.EmailReceivedAfter) {
		return errors.New("invalid time format for email-received-after")
	}

	if _, ok := opts.Populated[EmailReceivedBeforeFN]; ok && !IsValidTimeFormat(opts.EmailReceivedBefore) {
		return errors.New("invalid time format for email-received-before")
	}

	if _, ok := opts.Populated[EventStartsAfterFN]; ok && !IsValidTimeFormat(opts.EventStartsAfter) {
		return errors.New("invalid time format for event-starts-after")
	}

	if _, ok := opts.Populated[EventStartsBeforeFN]; ok && !IsValidTimeFormat(opts.EventStartsBefore) {
		return errors.New("invalid time format for event-starts-before")
	}

	if _, ok := opts.Populated[EventRecursFN]; ok && !IsValidBool(opts.EventRecurs) {
		return errors.New("invalid format for event-recurs")
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
	AddExchangeFilter(sel, opts.ContactName, sel.ContactName)
	AddExchangeFilter(sel, opts.EmailReceivedAfter, sel.MailReceivedAfter)
	AddExchangeFilter(sel, opts.EmailReceivedBefore, sel.MailReceivedBefore)
	AddExchangeFilter(sel, opts.EmailSender, sel.MailSender)
	AddExchangeFilter(sel, opts.EmailSubject, sel.MailSubject)
	AddExchangeFilter(sel, opts.EventOrganizer, sel.EventOrganizer)
	AddExchangeFilter(sel, opts.EventRecurs, sel.EventRecurs)
	AddExchangeFilter(sel, opts.EventStartsAfter, sel.EventStartsAfter)
	AddExchangeFilter(sel, opts.EventStartsBefore, sel.EventStartsBefore)
	AddExchangeFilter(sel, opts.EventSubject, sel.EventSubject)
}
