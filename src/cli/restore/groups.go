package restore

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/dttm"
)

// called by restore.go to map subcommands to provider-specific handling.
func addGroupsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(cmd, groupsRestoreCmd(), utils.MarkPreviewCommand())

		c.Use = c.Use + " " + groupsServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, true)
		flags.AddSkipPermissionsFlag(c)
		flags.AddSharePointDetailsAndRestoreFlags(c) // for sp restores
		flags.AddSiteIDFlag(c)
		flags.AddRestoreConfigFlags(c)
		flags.AddFailFastFlag(c)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
	}

	return c
}

const (
	groupsServiceCommand          = "groups"
	teamsServiceCommand           = "teams"
	groupsServiceCommandUseSuffix = "--backup <backupId>"

	groupsServiceCommandRestoreExamples = `# Restore file with ID 98765abcdef in Marketing's last backup (1234abcd...)
corso restore groups --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Restore the file with ID 98765abcdef without its associated permissions
corso restore groups --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef --no-permissions

# Restore all files named "FY2021 Planning.xlsx"
corso restore groups --backup 1234abcd-12ab-cd34-56de-1234abcd --file "FY2021 Planning.xlsx"

# Restore all files and folders in folder "Documents/Finance Reports" that were created before 2020
corso restore groups --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00`
)

// `corso restore groups [<flag>...]`
func groupsRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:     groupsServiceCommand,
		Aliases: []string{teamsServiceCommand},
		Short:   "Restore M365 Groups service data",
		RunE:    restoreGroupsCmd,
		Args:    cobra.NoArgs,
		Example: groupsServiceCommandRestoreExamples,
	}
}

// processes an groups service restore.
func restoreGroupsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.MakeGroupsOpts(cmd)
	opts.RestoreCfg.DTTMFormat = dttm.HumanReadableDriveItem

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateGroupsRestoreFlags(flags.BackupIDFV, opts); err != nil {
		return err
	}

	sel := utils.IncludeGroupsRestoreDataSelectors(ctx, opts)
	utils.FilterGroupsRestoreInfoSelectors(sel, opts)

	return runRestore(
		ctx,
		cmd,
		opts.RestoreCfg,
		sel.Selector,
		flags.BackupIDFV,
		"Groups")
}
