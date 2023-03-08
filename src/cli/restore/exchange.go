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
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
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

// called by restore.go to map subcommands to provider-specific handling.
func addExchangeCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(cmd, exchangeRestoreCmd())

		c.Use = c.Use + " " + exchangeServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		// general flags
		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to restore. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))

		fs.StringSliceVar(&user,
			utils.UserFN, nil,
			"Restore data by user's email address; accepts '"+utils.Wildcard+"' to select all users.")

		// email flags
		fs.StringSliceVar(&email,
			utils.EmailFN, nil,
			"Restore emails by ID; accepts '"+utils.Wildcard+"' to select all emails.")
		fs.StringSliceVar(
			&emailFolder,
			utils.EmailFolderFN, nil,
			"Restore emails within a folder; accepts '"+utils.Wildcard+"' to select all email folders.")
		fs.StringVar(
			&emailSubject,
			utils.EmailSubjectFN, "",
			"Restore emails with a subject containing this value.")
		fs.StringVar(
			&emailSender,
			utils.EmailSenderFN, "",
			"Restore emails from a specific sender.")
		fs.StringVar(
			&emailReceivedAfter,
			utils.EmailReceivedAfterFN, "",
			"Restore emails received after this datetime.")
		fs.StringVar(
			&emailReceivedBefore,
			utils.EmailReceivedBeforeFN, "",
			"Restore emails received before this datetime.")

		// event flags
		fs.StringSliceVar(&event,
			utils.EventFN, nil,
			"Restore events by event ID; accepts '"+utils.Wildcard+"' to select all events.")
		fs.StringSliceVar(
			&eventCalendar,
			utils.EventCalendarFN, nil,
			"Restore events under a calendar; accepts '"+utils.Wildcard+"' to select all event calendars.")
		fs.StringVar(
			&eventSubject,
			utils.EventSubjectFN, "",
			"Restore events with a subject containing this value.")
		fs.StringVar(
			&eventOrganizer,
			utils.EventOrganizerFN, "",
			"Restore events from a specific organizer.")
		fs.StringVar(
			&eventRecurs,
			utils.EventRecursFN, "",
			"Restore recurring events. Use `--event-recurs false` to restore non-recurring events.")
		fs.StringVar(
			&eventStartsAfter,
			utils.EventStartsAfterFN, "",
			"Restore events starting after this datetime.")
		fs.StringVar(
			&eventStartsBefore,
			utils.EventStartsBeforeFN, "",
			"Restore events starting before this datetime.")

		// contacts flags
		fs.StringSliceVar(
			&contact,
			utils.ContactFN, nil,
			"Restore contacts by contact ID; accepts '"+utils.Wildcard+"' to select all contacts.")
		fs.StringSliceVar(
			&contactFolder,
			utils.ContactFolderFN, nil,
			"Restore contacts within a folder; accepts '"+utils.Wildcard+"' to select all contact folders.")
		fs.StringVar(
			&contactName,
			utils.ContactNameFN, "",
			"Restore contacts whose contact name contains this value.")

		// others
		options.AddOperationFlags(c)
	}

	return c
}

const (
	exchangeServiceCommand          = "exchange"
	exchangeServiceCommandUseSuffix = "--backup <backupId>"

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
		Example: exchangeServiceCommandRestoreExamples,
	}
}

// processes an exchange service restore.
func restoreExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.ExchangeOpts{
		Contact:             contact,
		ContactFolder:       contactFolder,
		Email:               email,
		EmailFolder:         emailFolder,
		Event:               event,
		EventCalendar:       eventCalendar,
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

		Populated: utils.GetPopulatedFlags(cmd),
	}

	if err := utils.ValidateExchangeRestoreFlags(backupID, opts); err != nil {
		return err
	}

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	dest := control.DefaultRestoreDestination(common.SimpleDateTime)
	Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	sel := utils.IncludeExchangeRestoreDataSelectors(opts)
	utils.FilterExchangeRestoreInfoSelectors(sel, opts)

	ro, err := r.NewRestore(ctx, backupID, sel.Selector, dest)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to initialize Exchange restore"))
	}

	ds, err := ro.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, errors.Errorf("Backup or backup details missing for id %s", backupID))
		}

		return Only(ctx, errors.Wrap(err, "Failed to run Exchange restore"))
	}

	ds.PrintEntries(ctx)

	return nil
}
