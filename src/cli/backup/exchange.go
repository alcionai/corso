package backup

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/options"
	"github.com/alcionai/corso/cli/print"
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
	backupID      string
	exchangeAll   bool
	exchangeData  []string
	contact       []string
	contactFolder []string
	email         []string
	emailFolder   []string
	event         []string
	user          []string
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
		c, fs = utils.AddCommand(parent, exchangeCreateCmd)
		fs.StringArrayVar(&user, "user", nil, "Backup Exchange data by user ID; accepts "+utils.Wildcard+" to select all users")
		fs.BoolVar(&exchangeAll, "all", false, "Backup all Exchange data for all users")
		fs.StringArrayVar(
			&exchangeData,
			"data",
			nil,
			"Select one or more types of data to backup: "+dataEmail+", "+dataContacts+", or "+dataEvents)
		options.AddOperationFlags(c)

	case listCommand:
		c, _ = utils.AddCommand(parent, exchangeListCmd)

	case detailsCommand:
		c, fs = utils.AddCommand(parent, exchangeDetailsCmd)
		fs.StringVar(&backupID, "backup", "", "ID of the backup containing the details to be shown")
		cobra.CheckErr(c.MarkFlagRequired("backup"))
		fs.StringArrayVar(&contact, "contact", nil, "Select backup details by contact ID; accepts "+utils.Wildcard+" to select all contacts")
		fs.StringArrayVar(
			&contactFolder,
			"contact-folder",
			nil,
			"Select backup details by contact folder ID; accepts "+utils.Wildcard+" to select all contact folders")
		fs.StringArrayVar(&email, "email", nil, "Select backup details by emails ID; accepts "+utils.Wildcard+" to select all emails")
		fs.StringArrayVar(
			&emailFolder,
			"email-folder",
			nil,
			"Select backup details by email folder ID; accepts "+utils.Wildcard+" to select all email folderss")
		fs.StringArrayVar(&event, "event", nil, "Select backup details by event ID; accepts "+utils.Wildcard+" to select all events")
		fs.StringArrayVar(&user, "user", nil, "Select backup details by user ID; accepts "+utils.Wildcard+" to select all users")

		// TODO: reveal these flags when their production is supported in GC
		cobra.CheckErr(fs.MarkHidden("contact"))
		cobra.CheckErr(fs.MarkHidden("contact-folder"))
		cobra.CheckErr(fs.MarkHidden("event"))
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create exchange [<flag>...]`
var exchangeCreateCmd = &cobra.Command{
	Use:   exchangeServiceCommand,
	Short: "Backup M365 Exchange service data",
	RunE:  createExchangeCmd,
	Args:  cobra.NoArgs,
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

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return err
	}

	m365, err := acct.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
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

	sel := exchangeExchangeBackupCreateSelectors(exchangeAll, user, exchangeData)

	bo, err := r.NewBackup(ctx, sel, options.OperationOptions())
	if err != nil {
		return errors.Wrap(err, "Failed to initialize Exchange backup")
	}

	err = bo.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to run Exchange backup")
	}

	// todo: revive when backups are hooked up to backupOperation results
	// fmt.Printf("Created backup %s in %s for Exchange user %s.\n", result.SnapshotID, s.Provider, user)
	return nil
}

func exchangeExchangeBackupCreateSelectors(all bool, users, data []string) selectors.Selector {
	sel := selectors.NewExchangeBackup()
	if all {
		sel.Include(sel.Users(selectors.All()))
		return sel.Selector
	}
	if len(data) == 0 {
		sel.Include(sel.ContactFolders(user, selectors.All()))
		sel.Include(sel.MailFolders(user, selectors.All()))
		sel.Include(sel.Events(user, selectors.All()))
	}
	for _, d := range data {
		switch d {
		case dataContacts:
			sel.Include(sel.ContactFolders(users, selectors.All()))
		case dataEmail:
			sel.Include(sel.MailFolders(users, selectors.All()))
		case dataEvents:
			sel.Include(sel.Events(users, selectors.All()))
		}
	}
	return sel.Selector
}

func validateExchangeBackupCreateFlags(all bool, users, data []string) error {
	if len(users) == 0 && !all {
		return errors.New("requries one or more --user ids, the wildcard --user *, or the --all flag.")
	}
	if len(data) > 0 && all {
		return errors.New("--all does a backup on all data, and cannot be reduced with --data")
	}
	for _, d := range data {
		if d != dataContacts && d != dataEmail && d != dataEvents {
			return errors.New(d + " is an unrecognized data type; must be one of " + dataContacts + ", " + dataEmail + ", or " + dataEvents)
		}
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list exchange [<flag>...]`
var exchangeListCmd = &cobra.Command{
	Use:   exchangeServiceCommand,
	Short: "List the history of M365 Exchange service backups",
	RunE:  listExchangeCmd,
	Args:  cobra.NoArgs,
}

// lists the history of backup operations
func listExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return err
	}

	m365, err := acct.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	rps, err := r.Backups(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to list backups in the repository")
	}

	print.Backups(rps)

	return nil
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details exchange [<flag>...]`
var exchangeDetailsCmd = &cobra.Command{
	Use:   exchangeServiceCommand,
	Short: "Shows the details of a M365 Exchange service backup",
	RunE:  detailsExchangeCmd,
	Args:  cobra.NoArgs,
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

	s, acct, err := config.GetStorageAndAccount(true, nil)
	if err != nil {
		return err
	}

	m365, err := acct.M365Config()
	if err != nil {
		return errors.Wrap(err, "Failed to parse m365 account config")
	}

	logger.Ctx(ctx).Debugw(
		"Called - "+cmd.CommandPath(),
		"tenantID", m365.TenantID)

	r, err := repository.Connect(ctx, acct, s)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to the %s repository", s.Provider)
	}
	defer utils.CloseRepo(ctx, r)

	d, _, err := r.BackupDetails(ctx, backupID)
	if err != nil {
		return errors.Wrap(err, "Failed to get backup details in the repository")
	}

	sel := exchangeBackupDetailSelectors(contact, contactFolder, email, emailFolder, event, user)
	erSel, err := sel.ToExchangeRestore()
	if err != nil {
		return err
	}

	ds := erSel.FilterDetails(d)
	print.Entries(ds.Entries)

	return nil
}

func exchangeBackupDetailSelectors(
	contacts, contactFolders, emails, emailFolders, events, users []string,
) selectors.Selector {
	sel := selectors.NewExchangeBackup()
	lc, lcf := len(contacts), len(contactFolders)
	le, lef := len(emails), len(emailFolders)
	lev := len(events)
	lu := len(users)

	// if only the backupID is provided, treat that as an --all query
	if lc+lcf+le+lef+lev+lu == 0 {
		sel.Include(sel.Users(selectors.All()))
		return sel.Selector
	}

	// if only users are provided, we only get one selector
	if lc+lcf+le+lef+lev == 0 {
		sel.Include(sel.Users(users))
		return sel.Selector
	}

	// otherwise, add selectors for each type of data
	includeExchangeContacts(sel, users, contactFolders, contacts)
	includeExchangeEmails(sel, users, emailFolders, email)
	includeExchangeEvents(sel, users, events)

	return sel.Selector
}

func includeExchangeContacts(sel *selectors.ExchangeBackup, users, contactFolders, contacts []string) {
	if len(contactFolders) == 0 {
		return
	}
	if len(contacts) > 0 {
		sel.Include(sel.Contacts(users, contactFolders, contacts))
	} else {
		sel.Include(sel.ContactFolders(users, contactFolders))
	}
}

func includeExchangeEmails(sel *selectors.ExchangeBackup, users, emailFolders, emails []string) {
	if len(emailFolders) == 0 {
		return
	}
	if len(emails) > 0 {
		sel.Include(sel.Mails(users, emailFolders, emails))
	} else {
		sel.Include(sel.MailFolders(users, emailFolders))
	}
}

func includeExchangeEvents(sel *selectors.ExchangeBackup, users, events []string) {
	if len(events) == 0 {
		return
	}
	sel.Include(sel.Events(users, events))
}

func validateExchangeBackupDetailFlags(
	contacts, contactFolders, emails, emailFolders, events, users []string,
	backupID string,
) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is requried")
	}
	lu := len(users)
	lc, lcf := len(contacts), len(contactFolders)
	le, lef := len(emails), len(emailFolders)
	lev := len(events)
	// if only the backupID is populated, that's the same as --all
	if lu+lc+lcf+le+lef+lev == 0 {
		return nil
	}
	if lu == 0 {
		return errors.New("requries one or more --user ids, the wildcard --user *, or the --all flag.")
	}
	if lc > 0 && lcf == 0 {
		return errors.New("one or more --contact-folder ids or the wildcard --contact-folder * must be included to specify a --contact")
	}
	if le > 0 && lef == 0 {
		return errors.New("one or more --email-folder ids or the wildcard --email-folder * must be included to specify a --email")
	}
	return nil
}
