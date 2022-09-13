package restore

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
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
		fs.StringVar(&backupID, "backup", "", "ID of the backup to restore")
		cobra.CheckErr(c.MarkFlagRequired("backup"))

		// per-data-type flags
		fs.StringSliceVar(
			&contact,
			"contact", nil,
			"Restore contacts by ID; accepts "+utils.Wildcard+" to select all contacts")
		fs.StringSliceVar(
			&contactFolder,
			"contact-folder", nil,
			"Restore all contacts within the folder ID; accepts "+utils.Wildcard+" to select all contact folders")
		fs.StringSliceVar(&email,
			"email", nil,
			"Restore emails by ID; accepts "+utils.Wildcard+" to select all emails")
		fs.StringSliceVar(
			&emailFolder,
			"email-folder", nil,
			"Restore all emails by folder ID; accepts "+utils.Wildcard+" to select all email folders")
		fs.StringSliceVar(&event,
			"event", nil,
			"Restore events by ID; accepts "+utils.Wildcard+" to select all events")
		fs.StringSliceVar(
			&eventCalendar,
			"event-calendar", nil,
			"Restore events by calendar ID; accepts "+utils.Wildcard+" to select all event calendars")
		fs.StringSliceVar(&user,
			"user", nil,
			"Restore all data by user ID; accepts "+utils.Wildcard+" to select all users")

		// exchange-info flags
		fs.StringVar(
			&contactName,
			"contact-name", "",
			"Restore contacts where the contact name contains this value")
		fs.StringVar(
			&emailReceivedAfter,
			"email-received-after", "",
			"Restore mail where the email was received after this datetime")
		fs.StringVar(
			&emailReceivedBefore,
			"email-received-before", "",
			"Restore mail where the email was received before this datetime")
		fs.StringVar(
			&emailSender,
			"email-sender", "",
			"Restore mail where the email sender matches this user id")
		fs.StringVar(
			&emailSubject,
			"email-subject", "",
			"Restore mail where the email subject lines contain this value")
		fs.StringVar(
			&eventOrganizer,
			"event-organizer", "",
			"Restore events where the event organizer user id contains this value")
		fs.StringVar(
			&eventRecurs,
			"event-recurs", "",
			"Restore events if the event recurs.  Use --event-recurs false to select where the event does not recur.")
		fs.StringVar(
			&eventStartsAfter,
			"event-starts-after", "",
			"Restore events where the event starts after this datetime")
		fs.StringVar(
			&eventStartsBefore,
			"event-starts-before", "",
			"Restore events where the event starts before this datetime")
		fs.StringVar(
			&eventSubject,
			"event-subject", "",
			"Restore events where the event subject contains this value")

		// others
		options.AddOperationFlags(c)
	}

	return c
}

const exchangeServiceCommand = "exchange"

// `corso restore exchange [<flag>...]`
func exchangeRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Restore M365 Exchange service data",
		RunE:  restoreExchangeCmd,
		Args:  cobra.NoArgs,
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

	r, err := repository.Connect(ctx, a, s)
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	sel := selectors.NewExchangeRestore()
	utils.IncludeExchangeRestoreDataSelectors(
		sel,
		contact,
		contactFolder,
		email,
		emailFolder,
		event,
		eventCalendar,
		user)
	utils.FilterExchangeRestoreInfoSelectors(
		sel,
		contactName,
		emailReceivedAfter,
		emailReceivedBefore,
		emailSender,
		emailSubject,
		eventOrganizer,
		eventRecurs,
		eventStartsAfter,
		eventStartsBefore,
		eventSubject)

	// if no selector flags were specified, get all data in the service.
	if len(sel.Scopes()) == 0 {
		sel.Include(sel.Users(selectors.Any()))
	}

	ro, err := r.NewRestore(ctx, backupID, sel.Selector, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to initialize Exchange restore"))
	}

	if err := ro.Run(ctx); err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to run Exchange restore"))
	}

	Infof(ctx, "Restored Exchange in %s for user %s.\n", s.Provider, user)

	return nil
}
