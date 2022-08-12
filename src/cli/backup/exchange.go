package backup

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/options"
	. "github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/selectors"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

// exchange bucket info from flags
var (
	backupID            string
	exchangeAll         bool
	exchangeData        []string
	contact             []string
	contactFolder       []string
	email               []string
	emailFolder         []string
	emailReceivedAfter  string
	emailReceivedBefore string
	emailSender         string
	emailSubject        string
	event               []string
	user                []string
)

const (
	dataContacts = "contacts"
	dataEmail    = "email"
	dataEvents   = "events"
)

const exchangeServiceCommand = "exchange"

// called by backup.go to map parent subcommands to provider-specific handling.
func addExchangeCommands(parent *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch parent.Use {

	case createCommand:
		c, fs = utils.AddCommand(parent, exchangeCreateCmd())
		fs.StringSliceVar(
			&user,
			"user",
			nil,
			"Backup Exchange data by user ID; accepts "+utils.Wildcard+" to select all users",
		)
		fs.BoolVar(&exchangeAll, "all", false, "Backup all Exchange data for all users")
		fs.StringSliceVar(
			&exchangeData,
			"data",
			nil,
			"Select one or more types of data to backup: "+dataEmail+", "+dataContacts+", or "+dataEvents)
		options.AddOperationFlags(c)

	case listCommand:
		c, _ = utils.AddCommand(parent, exchangeListCmd())

	case detailsCommand:
		c, fs = utils.AddCommand(parent, exchangeDetailsCmd())
		fs.StringVar(&backupID, "backup", "", "ID of the backup containing the details to be shown")
		cobra.CheckErr(c.MarkFlagRequired("backup"))

		// per-data-type flags
		fs.StringSliceVar(
			&contact,
			"contact",
			nil,
			"Select backup details by contact ID; accepts "+utils.Wildcard+" to select all contacts",
		)
		fs.StringSliceVar(
			&contactFolder,
			"contact-folder",
			nil,
			"Select backup details by contact folder ID; accepts "+utils.Wildcard+" to select all contact folders",
		)
		fs.StringSliceVar(
			&email,
			"email",
			nil,
			"Select backup details by emails ID; accepts "+utils.Wildcard+" to select all emails",
		)
		fs.StringSliceVar(
			&emailFolder,
			"email-folder",
			nil,
			"Select backup details by email folder ID; accepts "+utils.Wildcard+" to select all email folderss")
		fs.StringSliceVar(
			&event,
			"event",
			nil,
			"Select backup details by event ID; accepts "+utils.Wildcard+" to select all events",
		)
		fs.StringSliceVar(
			&user,
			"user",
			nil,
			"Select backup details by user ID; accepts "+utils.Wildcard+" to select all users",
		)

		// TODO: reveal these flags when their production is supported in GC
		cobra.CheckErr(fs.MarkHidden("contact"))
		cobra.CheckErr(fs.MarkHidden("contact-folder"))
		cobra.CheckErr(fs.MarkHidden("event"))

		// exchange-info flags
		fs.StringVar(
			&emailReceivedAfter,
			"email-received-after",
			"",
			"Select backup details where the email was received after this datetime",
		)
		fs.StringVar(
			&emailReceivedBefore,
			"email-received-before",
			"",
			"Select backup details where the email was received before this datetime",
		)
		fs.StringVar(&emailSender, "email-sender", "", "Select backup details where the email sender matches this user id")
		fs.StringVar(
			&emailSubject,
			"email-subject",
			"",
			"Select backup details where the email subject lines contain this value",
		)
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create exchange [<flag>...]`
func exchangeCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Backup M365 Exchange service data",
		RunE:  createExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// processes an exchange service backup.
func createExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateExchangeBackupCreateFlags(exchangeAll, user, exchangeData); err != nil {
		return err
	}

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(err)
	}

	m365, err := acct.M365Config()
	if err != nil {
		return Only(errors.Wrap(err, "Failed to parse m365 account config"))
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID,
		"clientID", m365.ClientID,
		"hasClientSecret", len(m365.ClientSecret) > 0)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	sel := exchangeBackupCreateSelectors(exchangeAll, user, exchangeData)

	bo, err := r.NewBackup(ctx, sel, options.Control())
	if err != nil {
		return Only(errors.Wrap(err, "Failed to initialize Exchange backup"))
	}

	err = bo.Run(ctx)
	if err != nil {
		return Only(errors.Wrap(err, "Failed to run Exchange backup"))
	}

	bu, err := r.Backup(ctx, bo.Results.BackupID)
	if err != nil {
		return errors.Wrap(err, "Unable to retrieve backup results from storage")
	}

	OutputBackup(*bu)
	return nil
}

func exchangeBackupCreateSelectors(all bool, users, data []string) selectors.Selector {
	sel := selectors.NewExchangeBackup()
	if all {
		sel.Include(sel.Users(selectors.Any()))
		return sel.Selector
	}
	if len(data) == 0 {
		sel.Include(sel.ContactFolders(user, selectors.Any()))
		sel.Include(sel.MailFolders(user, selectors.Any()))
		sel.Include(sel.Events(user, selectors.Any()))
	}
	for _, d := range data {
		switch d {
		case dataContacts:
			sel.Include(sel.ContactFolders(users, selectors.Any()))
		case dataEmail:
			sel.Include(sel.MailFolders(users, selectors.Any()))
		case dataEvents:
			sel.Include(sel.Events(users, selectors.Any()))
		}
	}
	return sel.Selector
}

func validateExchangeBackupCreateFlags(all bool, users, data []string) error {
	if len(users) == 0 && !all {
		return errors.New("requires one or more --user ids, the wildcard --user *, or the --all flag")
	}
	if len(data) > 0 && all {
		return errors.New("--all does a backup on all data, and cannot be reduced with --data")
	}
	for _, d := range data {
		if d != dataContacts && d != dataEmail && d != dataEvents {
			return errors.New(
				d + " is an unrecognized data type; must be one of " + dataContacts + ", " + dataEmail + ", or " + dataEvents)
		}
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list exchange [<flag>...]`
func exchangeListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "List the history of M365 Exchange service backups",
		RunE:  listExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(err)
	}

	m365, err := acct.M365Config()
	if err != nil {
		return Only(errors.Wrap(err, "Failed to parse m365 account config"))
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return Only(errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}
	defer utils.CloseRepo(ctx, r)

	rps, err := r.Backups(ctx)
	if err != nil {
		return Only(errors.Wrap(err, "Failed to list backups in the repository"))
	}

	OutputBackups(rps)
	return nil
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details exchange [<flag>...]`
func exchangeDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Shows the details of a M365 Exchange service backup",
		RunE:  detailsExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func detailsExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateExchangeBackupDetailFlags(
		contact,
		contactFolder,
		email,
		emailFolder,
		event,
		user,
		backupID,
	); err != nil {
		return err
	}

	s, acct, err := config.GetStorageAndAccount(ctx, true, nil)
	if err != nil {
		return Only(err)
	}

	m365, err := acct.M365Config()
	if err != nil {
		return Only(errors.Wrap(err, "Failed to parse m365 account config"))
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return Only(errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider))
	}
	defer utils.CloseRepo(ctx, r)

	d, _, err := r.BackupDetails(ctx, backupID)
	if err != nil {
		return Only(errors.Wrap(err, "Failed to get backup details in the repository"))
	}

	sel := selectors.NewExchangeRestore()
	includeExchangeBackupDetailDataSelectors(
		sel,
		contact,
		contactFolder,
		email,
		emailFolder,
		event,
		user)
	filterExchangeBackupDetailInfoSelectors(
		sel,
		emailReceivedAfter,
		emailReceivedBefore,
		emailSender,
		emailSubject)

	// if no selector flags were specified, get all data in the service.
	if len(sel.Scopes()) == 0 {
		sel.Include(sel.Users(selectors.Any()))
	}

	ds := sel.Reduce(d)
	if len(ds.Entries) == 0 {
		return Only(errors.New("nothing to display: no items in the backup match the provided selectors"))
	}

	OutputEntries(ds.Entries)
	return nil
}

// builds the data-selector inclusions for `backup details exchange`
func includeExchangeBackupDetailDataSelectors(
	sel *selectors.ExchangeRestore,
	contacts, contactFolders, emails, emailFolders, events, users []string,
) {
	lc, lcf := len(contacts), len(contactFolders)
	le, lef := len(emails), len(emailFolders)
	lev := len(events)
	lu := len(users)

	if lc+lcf+le+lef+lev+lu == 0 {
		return
	}

	// if only users are provided, we only get one selector
	if lu > 0 && lc+lcf+le+lef+lev == 0 {
		sel.Include(sel.Users(users))
		return
	}

	// otherwise, add selectors for each type of data
	includeExchangeContacts(sel, users, contactFolders, contacts)
	includeExchangeEmails(sel, users, emailFolders, email)
	includeExchangeEvents(sel, users, events)
}

func includeExchangeContacts(sel *selectors.ExchangeRestore, users, contactFolders, contacts []string) {
	if len(contactFolders) == 0 {
		return
	}
	if len(contacts) > 0 {
		sel.Include(sel.Contacts(users, contactFolders, contacts))
	} else {
		sel.Include(sel.ContactFolders(users, contactFolders))
	}
}

func includeExchangeEmails(sel *selectors.ExchangeRestore, users, emailFolders, emails []string) {
	if len(emailFolders) == 0 {
		return
	}
	if len(emails) > 0 {
		sel.Include(sel.Mails(users, emailFolders, emails))
	} else {
		sel.Include(sel.MailFolders(users, emailFolders))
	}
}

func includeExchangeEvents(sel *selectors.ExchangeRestore, users, events []string) {
	if len(events) == 0 {
		return
	}
	sel.Include(sel.Events(users, events))
}

// builds the info-selector filters for `backup details exchange`
func filterExchangeBackupDetailInfoSelectors(
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
	sel.Filter(sel.MailSender([]string{sender}))
}

func filterExchangeInfoMailSubject(sel *selectors.ExchangeRestore, subject string) {
	if len(subject) == 0 {
		return
	}
	sel.Filter(sel.MailSubject([]string{subject}))
}

// checks all flags for correctness and interdependencies
func validateExchangeBackupDetailFlags(
	contacts, contactFolders, emails, emailFolders, events, users []string,
	backupID string,
) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}
	lu := len(users)
	lc, lcf := len(contacts), len(contactFolders)
	le, lef := len(emails), len(emailFolders)
	lev := len(events)
	if lu+lc+lcf+le+lef+lev == 0 {
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
			"one or more --email-folder ids or the wildcard --email-folder * must be included to specify a --email")
	}
	return nil
}
