package backup

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/path"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	groupsServiceCommand                 = "groups"
	groupsServiceCommandCreateUseSuffix  = "--group <groupsName> | '" + flags.Wildcard + "'"
	groupsServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	groupsServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

// TODO: correct examples
const (
	groupsServiceCommandCreateExamples = `# Backup all Groups data for Alice
corso backup create groups --group alice@example.com 

# Backup only Groups contacts for Alice and Bob
corso backup create groups --group engineering,sales --data contacts

# Backup all Groups data for all M365 users 
corso backup create groups --group '*'`

	groupsServiceCommandDeleteExamples = `# Delete Groups backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete groups --backup 1234abcd-12ab-cd34-56de-1234abcd`

	groupsServiceCommandDetailsExamples = `# Explore items in Alice's latest backup (1234abcd...)
corso backup details groups --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore calendar events occurring after start of 2022
corso backup details groups --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --event-starts-after 2022-01-01T00:00:00`
)

// called by backup.go to map subcommands to provider-specific handling.
func addGroupsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, groupsCreateCmd(), utils.HideCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandCreateUseSuffix
		c.Example = groupsServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		flags.AddGroupFlag(c)
		flags.AddDataFlag(c, []string{dataLibraries}, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		flags.AddFetchParallelismFlag(c)
		flags.AddFailFastFlag(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, groupsListCmd(), utils.HideCommand())
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		addFailedItemsFN(c)
		addSkippedItemsFN(c)
		addRecoveredErrorsFN(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, groupsDetailsCmd(), utils.HideCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandDetailsUseSuffix
		c.Example = groupsServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, groupsDeleteCmd(), utils.HideCommand())
		fs.SortFlags = false

		c.Use = c.Use + " " + groupsServiceCommandDeleteUseSuffix
		c.Example = groupsServiceCommandDeleteExamples

		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
	}

	return c
}

// ------------------------------------------------------------------------------------------------
// backup create
// ------------------------------------------------------------------------------------------------

// `corso backup create groups [<flag>...]`
func groupsCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "Backup M365 Group service data",
		RunE:  createGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes a groups service backup.
func createGroupsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	return Only(ctx, clues.New("not yet implemented"))
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list groups [<flag>...]`
func groupsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "List the history of M365 Groups service backups",
		RunE:  listGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listGroupsCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, flags.BackupIDFV, path.GroupsService, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details groups [<flag>...]`
func groupsDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "Shows the details of a M365 Groups service backup",
		RunE:  detailsGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes a groups service backup.
func detailsGroupsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	if err := validateGroupBackupCreateFlags(flags.GroupFV); err != nil {
		return Only(ctx, err)
	}

	return Only(ctx, clues.New("not yet implemented"))
}

// ------------------------------------------------------------------------------------------------
// backup delete
// ------------------------------------------------------------------------------------------------

// `corso backup delete groups [<flag>...]`
func groupsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   groupsServiceCommand,
		Short: "Delete backed-up M365 Groups service data",
		RunE:  deleteGroupsCmd,
		Args:  cobra.NoArgs,
	}
}

// deletes an groups service backup.
func deleteGroupsCmd(cmd *cobra.Command, args []string) error {
	return genericDeleteCommand(cmd, path.GroupsService, flags.BackupIDFV, "Groups", args)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func validateGroupBackupCreateFlags(groups []string) error {
	if len(groups) == 0 {
		return clues.New(
			"requires one or more --" +
				flags.GroupFN + " ids, or the wildcard --" +
				flags.GroupFN + " *",
		)
	}

	// TODO(meain)
	// for _, d := range cats {
	// 	if d != dataLibraries {
	// 		return clues.New(
	// 			d + " is an unrecognized data type; only  " + dataLibraries + " is supported"
	// 		)
	// 	}
	// }

	return nil
}
