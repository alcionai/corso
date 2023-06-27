package flags

import (
	"github.com/spf13/cobra"
)

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

// flag values (ie: FV)
var (
	ContactFV       []string
	ContactFolderFV []string
	ContactNameFV   string

	EmailFV               []string
	EmailFolderFV         []string
	EmailReceivedAfterFV  string
	EmailReceivedBeforeFV string
	EmailSenderFV         string
	EmailSubjectFV        string

	EventFV             []string
	EventCalendarFV     []string
	EventOrganizerFV    string
	EventRecursFV       string
	EventStartsAfterFV  string
	EventStartsBeforeFV string
	EventSubjectFV      string
)

// AddExchangeDetailsAndRestoreFlags adds flags that are common to both the
// details and restore commands.
func AddExchangeDetailsAndRestoreFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	// email flags
	fs.StringSliceVar(
		&EmailFV,
		EmailFN, nil,
		"Select email messages by ID; accepts '"+Wildcard+"' to select all emails.")
	fs.StringSliceVar(
		&EmailFolderFV,
		EmailFolderFN, nil,
		"Select emails within a folder; accepts '"+Wildcard+"' to select all email folders.")
	fs.StringVar(
		&EmailSubjectFV,
		EmailSubjectFN, "",
		"Select emails with a subject containing this value.")
	fs.StringVar(
		&EmailSenderFV,
		EmailSenderFN, "",
		"Select emails from a specific sender.")
	fs.StringVar(
		&EmailReceivedAfterFV,
		EmailReceivedAfterFN, "",
		"Select emails received after this datetime.")
	fs.StringVar(
		&EmailReceivedBeforeFV,
		EmailReceivedBeforeFN, "",
		"Select emails received before this datetime.")

	// event flags
	fs.StringSliceVar(
		&EventFV,
		EventFN, nil,
		"Select events by event ID; accepts '"+Wildcard+"' to select all events.")
	fs.StringSliceVar(
		&EventCalendarFV,
		EventCalendarFN, nil,
		"Select events under a calendar; accepts '"+Wildcard+"' to select all events.")
	fs.StringVar(
		&EventSubjectFV,
		EventSubjectFN, "",
		"Select events with a subject containing this value.")
	fs.StringVar(
		&EventOrganizerFV,
		EventOrganizerFN, "",
		"Select events from a specific organizer.")
	fs.StringVar(
		&EventRecursFV,
		EventRecursFN, "",
		"Select recurring events. Use `--event-recurs false` to select non-recurring events.")
	fs.StringVar(
		&EventStartsAfterFV,
		EventStartsAfterFN, "",
		"Select events starting after this datetime.")
	fs.StringVar(
		&EventStartsBeforeFV,
		EventStartsBeforeFN, "",
		"Select events starting before this datetime.")

	// contact flags
	fs.StringSliceVar(
		&ContactFV,
		ContactFN, nil,
		"Select contacts by contact ID; accepts '"+Wildcard+"' to select all contacts.")
	fs.StringSliceVar(
		&ContactFolderFV,
		ContactFolderFN, nil,
		"Select contacts within a folder; accepts '"+Wildcard+"' to select all contact folders.")
	fs.StringVar(
		&ContactNameFV,
		ContactNameFN, "",
		"Select contacts whose contact name contains this value.")
}
