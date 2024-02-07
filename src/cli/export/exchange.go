package export

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/cli/utils"
)

// called by export.go to map subcommands to provider-specific handling.
func addExchangeCommands(cmd *cobra.Command) *cobra.Command {
	var c *cobra.Command

	switch cmd.Use {
	case exportCommand:
		c, _ = utils.AddCommand(cmd, exchangeExportCmd())

		c.Use = c.Use + " " + exchangeServiceCommandUseSuffix

		flags.AddBackupIDFlag(c, true)
		flags.AddExchangeDetailsAndRestoreFlags(c, true)
		flags.AddExportConfigFlags(c)
		flags.AddFailFastFlag(c)
	}

	return c
}

const (
	exchangeServiceCommand          = "exchange"
	exchangeServiceCommandUseSuffix = "<destination> --backup <backupId>"

	// TODO(meain): remove message about only supporting email exports once others are added
	//nolint:lll
	exchangeServiceCommandExportExamples = `> Only email exports are supported as of now.

# Export emails with ID 98765abcdef and 12345abcdef from Alice's last backup (1234abcd...) to my-folder
corso export exchange my-folder --backup 1234abcd-12ab-cd34-56de-1234abcd --email 98765abcdef,12345abcdef

# Export emails with subject containing "Hello world" in the "Inbox" to my-folder
corso export exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
    --email-subject "Hello world" --email-folder Inbox my-folder`

// TODO(meain): Uncomment once support for these are added
// 		`# Export an entire calendar to my-folder
// corso export exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
//     --event-calendar Calendar my-folder

// # Export the contact with ID abdef0101 to my-folder
// corso export exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --contact abdef0101 my-folder`
)

// `corso export exchange [<flag>...] <destination>`
func exchangeExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exchangeServiceCommand,
		Short: "Export M365 Exchange service data",
		RunE:  exportExchangeCmd,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("missing export destination")
			}

			return nil
		},
		Example: exchangeServiceCommandExportExamples,
	}
}

// processes an exchange service export.
func exportExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	opts := utils.MakeExchangeOpts(cmd)

	if flags.RunModeFV == flags.RunModeFlagTest {
		return nil
	}

	if err := utils.ValidateExchangeRestoreFlags(flags.BackupIDFV, opts); err != nil {
		return err
	}

	sel := utils.IncludeExchangeRestoreDataSelectors(opts)
	utils.FilterExchangeRestoreInfoSelectors(sel, opts)

	return runExport(
		ctx,
		cmd,
		args,
		opts.ExportCfg,
		sel.Selector,
		flags.BackupIDFV,
		"Exchange",
		defaultAcceptedFormatTypes)
}
