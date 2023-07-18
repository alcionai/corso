package backup

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/path"
)

// ------------------------------------------------------------------------------------------------
// setup and globals
// ------------------------------------------------------------------------------------------------

const (
	teamsServiceCommand                 = "team"
	teamsServiceCommandCreateUseSuffix  = "--team <teamsName> | '" + flags.Wildcard + "'"
	teamsServiceCommandDeleteUseSuffix  = "--backup <backupId>"
	teamsServiceCommandDetailsUseSuffix = "--backup <backupId>"
)

const (
	teamsServiceCommandCreateExamples = `# Backup all Teams data for Alice
corso backup create team --team alice@example.com 

# Backup only Teams contacts for Alice and Bob
corso backup create team --team testTeams1,testTeams2 --data contacts

# Backup all Teams data for all M365 users 
corso backup create team --team '*'`

	teamsServiceCommandDeleteExamples = `# Delete Teams backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete team --backup 1234abcd-12ab-cd34-56de-1234abcd`

	teamsServiceCommandDetailsExamples = `# Explore items in Alice's latest backup (1234abcd...)
corso backup details team --backup 1234abcd-12ab-cd34-56de-1234abcd

# Explore calendar events occurring after start of 2022
corso backup details teams --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --event-starts-after 2022-01-01T00:00:00

# Explore contacts named Andy
corso backup details teams --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --contact-name Andy`
)

// called by backup.go to map subcommands to provider-specific handling.
func addTeamsCommands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case createCommand:
		c, fs = utils.AddCommand(cmd, teamsCreateCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + teamsServiceCommandCreateUseSuffix
		c.Example = teamsServiceCommandCreateExamples

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// TODO Neha: add teams flag
		flags.AddDataFlag(c, []string{dataEmail, dataContacts, dataEvents}, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		flags.AddFetchParallelismFlag(c)
		flags.AddFailFastFlag(c)

	case listCommand:
		c, fs = utils.AddCommand(cmd, teamsListCmd())
		fs.SortFlags = false

		flags.AddBackupIDFlag(c, false)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)
		addFailedItemsFN(c)
		addSkippedItemsFN(c)
		addRecoveredErrorsFN(c)

	case detailsCommand:
		c, fs = utils.AddCommand(cmd, teamsDetailsCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + teamsServiceCommandDetailsUseSuffix
		c.Example = teamsServiceCommandDetailsExamples

		flags.AddSkipReduceFlag(c)

		// Flags addition ordering should follow the order we want them to appear in help and docs:
		// More generic (ex: --user) and more frequently used flags take precedence.
		flags.AddBackupIDFlag(c, true)
		flags.AddCorsoPassphaseFlags(c)
		flags.AddAWSCredsFlags(c)
		flags.AddAzureCredsFlags(c)

	case deleteCommand:
		c, fs = utils.AddCommand(cmd, teamsDeleteCmd())
		fs.SortFlags = false

		c.Use = c.Use + " " + teamsServiceCommandDeleteUseSuffix
		c.Example = teamsServiceCommandDeleteExamples

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

// `corso backup create team [<flag>...]`
func teamsCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "Backup M365 Team service data",
		RunE:  createTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes an team service backup.
func createTeamsCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// ------------------------------------------------------------------------------------------------
// backup list
// ------------------------------------------------------------------------------------------------

// `corso backup list team [<flag>...]`
func teamsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "List the history of M365 Teams service backups",
		RunE:  listTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// lists the history of backup operations
func listTeamsCmd(cmd *cobra.Command, args []string) error {
	return genericListCommand(cmd, flags.BackupIDFV, path.TeamsService, args)
}

// ------------------------------------------------------------------------------------------------
// backup details
// ------------------------------------------------------------------------------------------------

// `corso backup details team [<flag>...]`
func teamsDetailsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "Shows the details of a M365 Teams service backup",
		RunE:  detailsTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// processes an team service backup.
func detailsTeamsCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// ------------------------------------------------------------------------------------------------
// backup delete
// ------------------------------------------------------------------------------------------------

// `corso backup delete team [<flag>...]`
func teamsDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   teamsServiceCommand,
		Short: "Delete backed-up M365 Teams service data",
		RunE:  deleteTeamsCmd,
		Args:  cobra.NoArgs,
	}
}

// deletes an teams service backup.
func deleteTeamsCmd(cmd *cobra.Command, args []string) error {
	return genericDeleteCommand(cmd, path.TeamsService, flags.BackupIDFV, "Teams", args)
}
