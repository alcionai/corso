package export

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/control"
)

// called by export.go to map subcommands to provider-specific handling.
func addTeamsChatsCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case exportCommand:
		c, _ = utils.AddCommand(cmd, teamschatsExportCmd(), utils.MarkPreviewCommand())

		c.Use = c.Use + " " + teamschatsServiceCommandUseSuffix

		flags.AddBackupIDFlag(c, true)
		flags.AddTeamsChatsDetailsAndRestoreFlags(c)
		flags.AddExportConfigFlags(c)
		flags.AddFailFastFlag(c)
	}

	return c
}

const (
	teamschatsServiceCommand          = "chats"
	teamschatsServiceCommandUseSuffix = "<destination> --backup <backupId>"

	//nolint:lll
	teamschatsServiceCommandExportExamples = `# Export a specific chat from the last backup (1234abcd...) to /my-exports
corso export chats my-exports --backup 1234abcd-12ab-cd34-56de-1234abcd --chat 98765abcdef

# Export all of Bob's chats to the current directory
corso export chats . --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --chat '*'

# Export all chats that were created before 2020 to /my-exports
corso export chats my-exports --backup 1234abcd-12ab-cd34-56de-1234abcd
    --chat-created-before 2020-01-01T00:00:00`
)

// `corso export chats [<flag>...] <destination>`
func teamschatsExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:     teamschatsServiceCommand,
		Aliases: []string{teamsServiceCommand},
		Short:   "Export M365 Chats data",
		RunE:    exportTeamsChatsCmd,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("missing export destination")
			}

			return nil
		},
		Example: teamschatsServiceCommandExportExamples,
	}
}

// processes an teamschats service export.
func exportTeamsChatsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.MakeTeamsChatsOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateTeamsChatsRestoreFlags(flags.BackupIDFV, opts, false); err != nil {
		return err
	}

	sel := utils.IncludeTeamsChatsRestoreDataSelectors(ctx, opts)
	utils.FilterTeamsChatsRestoreInfoSelectors(sel, opts)

	acceptedTeamsChatsFormatTypes := []string{
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
		"Chats",
		acceptedTeamsChatsFormatTypes)
}
