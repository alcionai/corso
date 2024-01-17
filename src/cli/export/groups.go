package export

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// called by export.go to map subcommands to provider-specific handling.
func addGroupsCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case exportCommand:
		c, _ = utils.AddCommand(cmd, groupsExportCmd(), utils.MarkPreviewCommand())

		c.Use = c.Use + " " + groupsServiceCommandUseSuffix

		flags.AddBackupIDFlag(c, true)
		flags.AddSiteFlag(c, false)
		flags.AddSiteIDFlag(c, false)
		flags.AddSharePointDetailsAndRestoreFlags(c)
		flags.AddGroupDetailsAndRestoreFlags(c)
		flags.AddExportConfigFlags(c)
		flags.AddFailFastFlag(c)
	}

	return c
}

const (
	groupsServiceCommand          = "groups"
	teamsServiceCommand           = "teams"
	groupsServiceCommandUseSuffix = "<destination> --backup <backupId>"

	//nolint:lll
	groupsServiceCommandExportExamples = `# Export a message in Marketing's last backup (1234abcd...) to /my-exports
corso export groups my-exports --backup 1234abcd-12ab-cd34-56de-1234abcd --message 98765abcdef

# Export all messages named in channel "Finance Reports" to the current directory
corso export groups . --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --message '*' --channel "Finance Reports"

# Export all messages in channel "Finance Reports" that were created before 2020 to /my-exports
corso export groups my-exports --backup 1234abcd-12ab-cd34-56de-1234abcd
    --channel "Finance Reports" --message-created-before 2020-01-01T00:00:00

# Export all files and folders in folder "Documents/Finance Reports" that were created before 2020 to /my-exports
corso export groups my-exports --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00`
)

// `corso export groups [<flag>...] <destination>`
func groupsExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:     groupsServiceCommand,
		Aliases: []string{teamsServiceCommand},
		Short:   "Export M365 Groups service data",
		RunE:    exportGroupsCmd,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("missing export destination")
			}

			return nil
		},
		Example: groupsServiceCommandExportExamples,
	}
}

// processes an groups service export.
func exportGroupsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.MakeGroupsOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateGroupsRestoreFlags(flags.BackupIDFV, opts, false); err != nil {
		return err
	}

	sel := utils.IncludeGroupsRestoreDataSelectors(ctx, opts)
	utils.FilterGroupsRestoreInfoSelectors(sel, opts)

	// TODO(pandeyabs): Exclude conversations from export since they are not
	// supported yet. https://github.com/alcionai/corso/issues/4822
	sel.Exclude(sel.Conversation(selectors.Any()))

	acceptedGroupsFormatTypes := []string{
		string(control.DefaultFormat),
		string(control.JSONFormat),
	}

	return runExport(
		ctx,
		cmd,
		args,
		opts.ExportCfg,
		sel.Selector,
		flags.BackupIDFV,
		"Groups",
		acceptedGroupsFormatTypes)
}
