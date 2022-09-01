package restore

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/options"
	. "github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
)

// exchange bucket info from flags
var (
	backupID            string
	contact             []string
	contactFolder       []string
	email               []string
	emailFolder         []string
	emailReceivedAfter  string
	emailReceivedBefore string
	emailSender         string
	emailSubject        string
	event               []string
	eventCalendar       []string
	user                []string
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
			"contact",
			nil,
			"Restore contacts by ID; accepts "+utils.Wildcard+" to select all contacts",
		)
		fs.StringSliceVar(
			&contactFolder,
			"contact-folder",
			nil,
			"Restore all contacts within the folder ID; accepts "+utils.Wildcard+" to select all contact folders")
		fs.StringSliceVar(&email, "email", nil, "Restore emails by ID; accepts "+utils.Wildcard+" to select all emails")
		fs.StringSliceVar(
			&emailFolder,
			"email-folder",
			nil,
			"Restore all emails by folder ID; accepts "+utils.Wildcard+" to select all email folders",
		)
		fs.StringSliceVar(&event, "event", nil, "Restore events by ID; accepts "+utils.Wildcard+" to select all events")
		fs.StringSliceVar(
			&eventCalendar,
			"event-calendar",
			nil,
			"Restore events by calendar ID; accepts "+utils.Wildcard+" to select all event calendars")
		fs.StringSliceVar(&user, "user", nil, "Restore all data by user ID; accepts "+utils.Wildcard+" to select all users")

		// TODO: reveal these flags when their production is supported in GC
		cobra.CheckErr(fs.MarkHidden("contact"))
		cobra.CheckErr(fs.MarkHidden("contact-folder"))
		cobra.CheckErr(fs.MarkHidden("event"))
		cobra.CheckErr(fs.MarkHidden("event-calendar"))

		// exchange-info flags
		fs.StringVar(
			&emailReceivedAfter,
			"email-received-after",
			"",
			"Restore mail where the email was received after this datetime",
		)
		fs.StringVar(
			&emailReceivedBefore,
			"email-received-before",
			"",
			"Restore mail where the email was received before this datetime",
		)
		fs.StringVar(&emailSender, "email-sender", "", "Restore mail where the email sender matches this user id")
		fs.StringVar(&emailSubject, "email-subject", "", "Restore mail where the email subject lines contain this value")

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

	if err := validateExchangeRestoreFlags(
		contact,
		contactFolder,
		email,
		emailFolder,
		event,
		eventCalendar,
		user,
		backupID,
	); err != nil {
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
	includeExchangeRestoreDataSelectors(
		sel,
		contact,
		contactFolder,
		email,
		emailFolder,
		event,
		eventCalendar,
		user)
	filterExchangeRestoreInfoSelectors(
		sel,
		emailReceivedAfter,
		emailReceivedBefore,
		emailSender,
		emailSubject)

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

// builds the data-selector inclusions for `restore exchange`
func includeExchangeRestoreDataSelectors(
	sel *selectors.ExchangeRestore,
	contacts, contactFolders, emails, emailFolders, events, eventCalendars, users []string,
) {
	lc, lcf := len(contacts), len(contactFolders)
	le, lef := len(emails), len(emailFolders)
	lev, lec := len(events), len(eventCalendars)
	lu := len(users)

	if lc+lcf+le+lef+lev+lec+lu == 0 {
		return
	}

	// if only users are provided, we only get one selector
	if lu > 0 && lc+lcf+le+lef+lev+lec == 0 {
		sel.Include(sel.Users(users))
		return
	}

	// otherwise, add selectors for each type of data
	includeExchangeContacts(sel, users, contactFolders, contacts)
	includeExchangeEmails(sel, users, emailFolders, email)
	includeExchangeEvents(sel, users, eventCalendars, events)
}

func includeExchangeContacts(sel *selectors.ExchangeRestore, users, contactFolders, contacts []string) {
	if len(contactFolders) == 0 {
		return
	}

	if len(contacts) == 0 {
		contacts = selectors.Any()
	}

	sel.Include(sel.Contacts(users, contactFolders, contacts))
}

func includeExchangeEmails(sel *selectors.ExchangeRestore, users, emailFolders, emails []string) {
	if len(emailFolders) == 0 {
		return
	}

	if len(emails) == 0 {
		emails = selectors.Any()
	}

	sel.Include(sel.Mails(users, emailFolders, emails))
}

func includeExchangeEvents(sel *selectors.ExchangeRestore, users, eventCalendars, events []string) {
	if len(eventCalendars) == 0 {
		return
	}

	if len(events) == 0 {
		events = selectors.Any()
	}

	sel.Include(sel.Events(users, eventCalendars, events))
}

// builds the info-selector filters for `restore exchange`
func filterExchangeRestoreInfoSelectors(
	sel *selectors.ExchangeRestore,
	emailReceivedAfter, emailReceivedBefore, emailSender, emailSubject string,
) {
	filterExchangeInfoMailReceivedAfter(sel, emailReceivedAfter)
	filterExchangeInfoMailReceivedBefore(sel, emailReceivedBefore)
	filterExchangeInfoMailSender(sel, emailSender)
	filterExchangeInfoMailSubject(sel, emailSubject)
}

func filterExchangeInfoMailReceivedAfter(sel *selectors.ExchangeRestore, receivedAfter string) {
	if len(receivedAfter) == 0 {
		return
	}

	sel.Filter(sel.MailReceivedAfter(receivedAfter))
}

func filterExchangeInfoMailReceivedBefore(sel *selectors.ExchangeRestore, receivedBefore string) {
	if len(receivedBefore) == 0 {
		return
	}

	sel.Filter(sel.MailReceivedBefore(receivedBefore))
}

func filterExchangeInfoMailSender(sel *selectors.ExchangeRestore, sender string) {
	if len(sender) == 0 {
		return
	}

	sel.Filter(sel.MailSender(sender))
}

func filterExchangeInfoMailSubject(sel *selectors.ExchangeRestore, subject string) {
	if len(subject) == 0 {
		return
	}

	sel.Filter(sel.MailSubject(subject))
}

// checks all flags for correctness and interdependencies
func validateExchangeRestoreFlags(
	contacts, contactFolders, emails, emailFolders, events, eventCalendars, users []string,
	backupID string,
) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}

	lu := len(users)
	lc, lcf := len(contacts), len(contactFolders)
	le, lef := len(emails), len(emailFolders)
	lev, lec := len(events), len(eventCalendars)

	// if only the backupID is populated, that's the same as --all
	if lu+lc+lcf+le+lef+lev+lec == 0 {
		return nil
	}

	if lu == 0 {
		return errors.New("requires one or more --user ids, the wildcard --user *, or the --all flag")
	}

	if lc > 0 && lcf == 0 {
		return errors.New(
			"one or more --contact-folder ids or the wildcard --contact-folder * must be included to specify a --contact")
	}

	if le > 0 && lef == 0 {
		return errors.New(
			"one or more --email-folder ids or the wildcard --email-folder * must be included to specify an --email")
	}

	if lev > 0 && lec == 0 {
		return errors.New(
			"one or more --event-calendar ids or the wildcard --event-calendar * must be included to specify an --event")
	}

	return nil
}
