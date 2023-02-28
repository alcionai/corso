package backup

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/options"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365"
	"github.com/alcionai/corso/src/pkg/store"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

// exchange bucket info from flags
var (
	backupID     string
	exchangeData []string
	user         []string

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

const (
	dataContacts = "contacts"
	dataEmail    = "email"
	dataEvents   = "events"
)

const (
	exchangeServiceCommand                 = "exchange"
	exchangeServiceCommandCreateUseSuffix  = "--user <email> | '" + utils.Wildcard + "'"
	exchangeServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	exchangeServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	exchangeServiceCommandCreateExamples = `# Backup all Exchange data for Alice
corso backup create exchange --user alice@example.com

# Backup only Exchange contacts for Alice and Bob
corso backup create exchange --user alice@example.com,bob@example.com --data contacts

# Backup all Exchange data for all M365 users 
corso backup create exchange --user '*'`

	exchangeServiceCommandDeleteExamples = `# Delete Exchange backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete exchange --backup 1234abcd-12ab-cd34-56de-1234abcd`

	exchangeServiceCommandDetailsExamples = `# Explore Alice's items in backup 1234abcd-12ab-cd34-56de-1234abcd 
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --user alice@example.com

# Explore Alice's emails with subject containing "Hello world" in folder "Inbox" from a specific backup 
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --email-subject "Hello world" --email-folder Inbox

# Explore Bobs's events occurring after start of 2022 from a specific backup
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user bob@example.com --event-starts-after 2022-01-01T00:00:00

# Explore Alice's contacts with name containing Andy from a specific backup
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --contact-name Andy`
)

// called by backup.go to map subcommands to provider-specific handling.
func addExchangeCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, exchangeCreateCmd())
		options.AddFeatureToggle(cmd, options.DisableIncrementals())

		c.Use = c.Use + " " + exchangeServiceCommandCreateUseSuffix
		c.Example = exchangeServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.StringSliceVar(
			&user,
			utils.UserFN, nil,
			"Backup Exchange data by a user's email; accepts '"+utils.Wildcard+"' to select all users")
		fs.StringSliceVar(
			&exchangeData,
			utils.DataFN, nil,
			"Select one or more types of data to backup: "+dataEmail+", "+dataContacts+", or "+dataEvents)
		options.AddOperationFlags(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, exchangeListCmd())

		fs.StringVar(&backupID,
			"backup", "",
			"ID of the backup to retrieve.")

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, exchangeDetailsCmd())

		c.Use = c.Use + " " + exchangeServiceCommandDetailsUseSuffix
		c.Example = exchangeServiceCommandDetailsExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to explore. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))
		fs.StringSliceVar(
			&user,
			utils.UserFN, nil,
			"Select backup details by user ID; accepts '"+utils.Wildcard+"' to select all users.")

		// email flags
		fs.StringSliceVar(
			&email,
			utils.EmailFN, nil,
			"Select backup details for emails by email ID; accepts '"+utils.Wildcard+"' to select all emails.")
		fs.StringSliceVar(
			&emailFolder,
			utils.EmailFolderFN, nil,
			"Select backup details for emails within a folder; accepts '"+utils.Wildcard+"' to select all email folders.")
		fs.StringVar(
			&emailSubject,
			utils.EmailSubjectFN, "",
			"Select backup details for emails with a subject containing this value.")
		fs.StringVar(
			&emailSender,
			utils.EmailSenderFN, "",
			"Select backup details for emails from a specific sender.")
		fs.StringVar(
			&emailReceivedAfter,
			utils.EmailReceivedAfterFN, "",
			"Select backup details for emails received after this datetime.")
		fs.StringVar(
			&emailReceivedBefore,
			utils.EmailReceivedBeforeFN, "",
			"Select backup details for emails received before this datetime.")

		// event flags
		fs.StringSliceVar(
			&event,
			utils.EventFN, nil,
			"Select backup details for events by event ID; accepts '"+utils.Wildcard+"' to select all events.")
		fs.StringSliceVar(
			&eventCalendar,
			utils.EventCalendarFN, nil,
			"Select backup details for events under a calendar; accepts '"+utils.Wildcard+"' to select all events.")
		fs.StringVar(
			&eventSubject,
			utils.EventSubjectFN, "",
			"Select backup details for events with a subject containing this value.")
		fs.StringVar(
			&eventOrganizer,
			utils.EventOrganizerFN, "",
			"Select backup details for events from a specific organizer.")
		fs.StringVar(
			&eventRecurs,
			utils.EventRecursFN, "",
			"Select backup details for recurring events. Use `--event-recurs false` to select non-recurring events.")
		fs.StringVar(
			&eventStartsAfter,
			utils.EventStartsAfterFN, "",
			"Select backup details for events starting after this datetime.")
		fs.StringVar(
			&eventStartsBefore,
			utils.EventStartsBeforeFN, "",
			"Select backup details for events starting before this datetime.")

		// contact flags
		fs.StringSliceVar(
			&contact,
			utils.ContactFN, nil,
			"Select backup details for contacts by contact ID; accepts '"+utils.Wildcard+"' to select all contacts.")
		fs.StringSliceVar(
			&contactFolder,
			utils.ContactFolderFN, nil,
			"Select backup details for contacts within a folder; accepts '"+utils.Wildcard+"' to select all contact folders.")

		fs.StringVar(
			&contactName,
			utils.ContactNameFN, "",
			"Select backup details for contacts whose contact name contains this value.")

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, exchangeDeleteCmd())

		c.Use = c.Use + " " + exchangeServiceCommandDeleteUseSuffix
		c.Example = exchangeServiceCommandDeleteExamples

		fs.StringVar(&backupID,
			utils.BackupFN, "",
			"ID of the backup to delete. (required)")
		cobra.CheckErr(c.MarkFlagRequired(utils.BackupFN))
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

	if err := validateExchangeBackupCreateFlags(user, exchangeData); err != nil {
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

	sel := exchangeBackupCreateSelectors(user, exchangeData)

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	users, err := m365.UserPNs(ctx, cfg.Account, errs)
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to retrieve M365 user(s)"))
	}

	var (
		merrs *multierror.Error
		bIDs  []model.StableID
	)

	for _, discSel := range sel.SplitByResourceOwner(users) {
		bo, err := r.NewBackup(ctx, discSel.Selector)
		if err != nil {
			merrs = multierror.Append(merrs, errors.Wrapf(
				err,
				"Failed to initialize Exchange backup for user %s",
				discSel.DiscreteOwner,
			))

			continue
		}

		err = bo.Run(ctx)
		if err != nil {
			merrs = multierror.Append(merrs, errors.Wrapf(
				err,
				"Failed to run Exchange backup for user %s",
				discSel.DiscreteOwner,
			))

			continue
		}

		bIDs = append(bIDs, bo.Results.BackupID)
	}

	bups, ferrs := r.Backups(ctx, bIDs)
	// TODO: print/log recoverable errors
	if ferrs.Failure() != nil {
		return Only(ctx, errors.Wrap(ferrs.Failure(), "Unable to retrieve backup results from storage"))
	}

	backup.PrintAll(ctx, bups)

	if e := merrs.ErrorOrNil(); e != nil {
		return Only(ctx, e)
	}

	return nil
}

func exchangeBackupCreateSelectors(userIDs, cats []string) *selectors.ExchangeBackup {
	sel := selectors.NewExchangeBackup(userIDs)

	if len(cats) == 0 {
		sel.Include(sel.ContactFolders(selectors.Any()))
		sel.Include(sel.MailFolders(selectors.Any()))
		sel.Include(sel.EventCalendars(selectors.Any()))
	}

	for _, d := range cats {
		switch d {
		case dataContacts:
			sel.Include(sel.ContactFolders(selectors.Any()))
		case dataEmail:
			sel.Include(sel.MailFolders(selectors.Any()))
		case dataEvents:
			sel.Include(sel.EventCalendars(selectors.Any()))
		}
	}

	return sel
}

func validateExchangeBackupCreateFlags(userIDs, cats []string) error {
	if len(userIDs) == 0 {
		return errors.New("--user requires one or more email addresses or the wildcard '*'")
	}

	for _, d := range cats {
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

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	if len(backupID) > 0 {
		b, err := r.Backup(ctx, model.StableID(backupID))
		if err != nil {
			if errors.Is(err, data.ErrNotFound) {
				return Only(ctx, errors.Errorf("No backup exists with the id %s", backupID))
			}

			return Only(ctx, errors.Wrap(err, "Failed to find backup "+backupID))
		}

		b.Print(ctx)

		return nil
	}

	bs, err := r.BackupsByTag(ctx, store.Service(path.ExchangeService))
	if err != nil {
		return Only(ctx, errors.Wrap(err, "Failed to list backups in the repository"))
	}

	backup.PrintAll(ctx, bs)

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
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctx := cmd.Context()
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

	cfg, err := config.GetConfigRepoDetails(ctx, true, nil)
	if err != nil {
		return Only(ctx, err)
	}

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, options.Control())
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "Failed to connect to the %s repository", cfg.Storage.Provider))
	}

	defer utils.CloseRepo(ctx, r)

	ds, err := runDetailsExchangeCmd(ctx, r, backupID, opts)
	if err != nil {
		return Only(ctx, err)
	}

	if len(ds.Entries) == 0 {
		Info(ctx, selectors.ErrorNoMatchingItems)
		return nil
	}

	ds.PrintEntries(ctx)

	return nil
}

// runDetailsExchangeCmd actually performs the lookup in backup details.
// the fault.Errors return is always non-nil.  Callers should check if
// errs.Failure() == nil.
func runDetailsExchangeCmd(
	ctx context.Context,
	r repository.BackupGetter,
	backupID string,
	opts utils.ExchangeOpts,
) (*details.Details, error) {
	if err := utils.ValidateExchangeRestoreFlags(backupID, opts); err != nil {
		return nil, err
	}

	d, _, errs := r.BackupDetails(ctx, backupID)
	// TODO: log/track recoverable errors
	if errs.Failure() != nil {
		if errors.Is(errs.Failure(), data.ErrNotFound) {
			return nil, errors.Errorf("No backup exists with the id %s", backupID)
		}

		return nil, errors.Wrap(errs.Failure(), "Failed to get backup details in the repository")
	}

	sel := utils.IncludeExchangeRestoreDataSelectors(opts)
	utils.FilterExchangeRestoreInfoSelectors(sel, opts)

	return sel.Reduce(ctx, d, errs), nil
}

// ------------------------------------------------------------------------------------------------
// backup delete
// ------------------------------------------------------------------------------------------------

// `corso backup delete exchange [<flag>...]`
func exchangeDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Delete backed-up M365 Exchange service data",
		RunE:  deleteExchangeCmd,
		Args:  cobra.NoArgs,
	}
}

// deletes an exchange service backup.
func deleteExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
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

	if err := r.DeleteBackup(ctx, model.StableID(backupID)); err != nil {
		return Only(ctx, errors.Wrapf(err, "Deleting backup %s", backupID))
	}

	Info(ctx, "Deleted Exchange backup ", backupID)

	return nil
}
