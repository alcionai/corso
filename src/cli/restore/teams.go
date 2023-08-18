package restore

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
)

// called by restore.go to map subcommands to provider-specific handling.
func addTeamsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case restoreCommand:
		c, fs = utils.AddCommand(cmd, teamsRestoreCmd(), utils.HideCommand())

		c.Use = c.Use + " " + teamsServiceCommandUseSuffix

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, true)
		flags.AddRestorePermissionsFlag(c)
		flags.AddRestoreConfigFlags(c)
		flags.AddFailFastFlag(c)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
	}

	return c
}

// TODO: correct examples
const (
	teamsServiceCommand          = "teams"
	teamsServiceCommandUseSuffix = "--backup <backupId>"

	teamsServiceCommandRestoreExamples = `# Restore file with ID 98765abcdef in Bob's last backup (1234abcd...)
corso restore teams --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Restore the file with ID 98765abcdef along with its associated permissions
corso restore teams --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef --restore-permissions

# Restore files named "FY2021 Planning.xlsx" in "Documents/Finance Reports"
corso restore teams --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --file "FY2021 Planning.xlsx" --folder "Documents/Finance Reports"

# Restore all files and folders in folder "Documents/Finance Reports" that were created before 2020
corso restore teams --backup 1234abcd-12ab-cd34-56de-1234abcd 
    --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00`
)

// `corso restore teams [<flag>...]`
func teamsRestoreCmd() *cobra.Command {
	return &cobra.Command{
		Use:     teamsServiceCommand,
		Short:   "Restore M365 Teams service data",
		RunE:    restoreTeamsCmd,
		Args:    cobra.NoArgs,
		Example: teamsServiceCommandRestoreExamples,
	}
}

// processes an teams service restore.
func restoreTeamsCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	return Only(ctx, clues.New("not yet implemented"))
}
