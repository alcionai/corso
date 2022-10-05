package restore

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// exchange bucket info from flags
var (
	backupID string
	user     []string

	contact       []string
	contactFolder []string
	contactName   string

	email               []string
	emailFolder         []string
	emailReceivedAfter  string
	emailReceivedBefore string
	emailSender         string
	emailSubject        string

	event             []string
	eventCalendar     []string
	eventOrganizer    string
	eventRecurs       string
	eventStartsAfter  string
	eventStartsBefore string
	eventSubject      string
)

// called by restore.go to map parent subcommands to provider-specific handling.
func addExchangeCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(parent, exchangeRestoreCmd())

		c.Use = c.Use + exchangeServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --all) and more frequently used flags take precedence.
		// general flags
		fs.StringVar(&backupID, "backup", "", "ID of the backup to restore. (required)")
		cobra.CheckErr(c.MarkFlagRequired("backup"))

		fs.StringSliceVar(&user,
			"user", nil,
			"Restore data by user ID; accepts '"+utils.Wildcard+"' to select all users.")

		// email flags
		fs.StringSliceVar(&email,
			"email", nil,
			"Restore emails by ID; accepts '"+utils.Wildcard+"' to select all emails.")
		fs.StringSliceVar(
			&emailFolder,
			"email-folder", nil,
			"Restore emails within a folder; accepts '"+utils.Wildcard+"' to select all email folders.")
		fs.StringVar(
			&emailSubject,
			"email-subject", "",
			"Restore emails with a subject containing this value.")
		fs.StringVar(
			&emailSender,
			"email-sender", "",
			"Restore emails from a specific sender.")
		fs.StringVar(
			&emailReceivedAfter,
			"email-received-after", "",
			"Restore emails received after this datetime.")
		fs.StringVar(
			&emailReceivedBefore,
			"email-received-before", "",
			"Restore emails received before this datetime.")

		// event flags
		fs.StringSliceVar(&event,
			"event", nil,
			"Restore events by event ID; accepts '"+utils.Wildcard+"' to select all events.")
		fs.StringSliceVar(
			&eventCalendar,
			"event-calendar", nil,
			"Restore events under a calendar; accepts '"+utils.Wildcard+"' to select all event calendars.")
		fs.StringVar(
			&eventSubject,
			"event-subject", "",
			"Restore events with a subject containing this value.")
		fs.StringVar(
			&eventOrganizer,
			"event-organizer", "",
			"Restore events from a specific organizer.")
		fs.StringVar(
			&eventRecurs,
			"event-recurs", "",
			"Restore recurring events. Use `--event-recurs false` to restore non-recurring events.")
		fs.StringVar(
			&eventStartsAfter,
			"event-starts-after", "",
			"Restore events starting after this datetime.")
		fs.StringVar(
			&eventStartsBefore,
			"event-starts-before", "",
			"Restore events starting before this datetime.")

		// contacts flags
		fs.StringSliceVar(
			&contact,
			"contact", nil,
			"Restore contacts by contact ID; accepts '"+utils.Wildcard+"' to select all contacts.")
		fs.StringSliceVar(
			&contactFolder,
			"contact-folder", nil,
			"Restore contacts within a folder; accepts '"+utils.Wildcard+"' to select all contact folders.")
		fs.StringVar(
			&contactName,
			"contact-name", "",
			"Restore contacts whose contact name contains this value.")

		// others
		options.AddOperationFlags(c)
	}

	return c
}

const (
	exchangeServiceCommand          = "exchange"
	exchangeServiceCommandUseSuffix = " --backup <backupId>"

	exchangeServiceCommandRestoreExamples = `# Restore emails with ID 98765abcdef and 12345abcdef from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --email 98765abcdef,12345abcdef

# Restore Alice's emails with subject containing "Hello world" in "Inbox" from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
	--user alice@example.com --email-subject "Hello world" --email-folder Inbox

# Restore Bobs's entire calendar from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
	--user bob@example.com --event-calendar Calendar

# Restore contact with ID abdef0101 from a specific backup
corso restore exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --contact abdef0101`
)

// `corso restore exchange [<flag>...]`
func exchangeRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:     exchangeServiceCommand,
		Short:   "Restore M365 Exchange service data",
		RunE:    restoreExchangeCmd,
		Args:    cobra.NoArgs,
		Example: utils.IndentExamples(exchangeServiceCommandRestoreExamples),
	}
}

// processes an exchange service restore.
func restoreExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := utils.ValidateExchangeRestoreFlags(backupID); err != nil {
		return err
	}

	s, a, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, a, s, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	opts := utils.ExchangeOpts{
		Contacts:            contact,
		ContactFolders:      contactFolder,
		Emails:              email,
		EmailFolders:        emailFolder,
		Events:              event,
		EventCalendars:      eventCalendar,
		Users:               user,
		ContactName:         contactName,
		EmailReceivedAfter:  emailReceivedAfter,
		EmailReceivedBefore: emailReceivedBefore,
		EmailSender:         emailSender,
		EmailSubject:        emailSubject,
		EventOrganizer:      eventOrganizer,
		EventRecurs:         eventRecurs,
		EventStartsAfter:    eventStartsAfter,
		EventStartsBefore:   eventStartsBefore,
		EventSubject:        eventSubject,
	}

	sel := selectors.NewExchangeRestore()
	utils.IncludeExchangeRestoreDataSelectors(sel, opts)
	utils.FilterExchangeRestoreInfoSelectors(sel, opts)

	// if no selector flags were specified, get all data in the service.
	if len(sel.Scopes()) == 0 {
		sel.Include(sel.Users(selectors.Any()))
	}

	restoreDest := control.DefaultRestoreDestination(common.SimpleDateTimeFormat)

	ro, err := r.NewRestore(ctx, backupID, sel.Selector, restoreDest)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to initialize Exchange restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to run Exchange restore"))
	}

	Infof(ctx, "Restored OneDrive in %s for user %s.\n", s.Provider, sel.ToPrintable().Resources())
	ds.PrintEntries(ctx)

	return nil
}
