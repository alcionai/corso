package export

import (
	"context"
	"errors"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var exportCommands = []func(cmd *cobra.Command) *cobra.Command{
	addOneDriveCommands,
	addSharePointCommands,
	addGroupsCommands,
	addExchangeCommands,
}

var defaultAcceptedFormatTypes = []string{string(control.DefaultFormat)}

// AddCommands attaches all `corso export * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	subCommand := exportCmd()
	cmd.AddCommand(subCommand)

	for _, addExportTo := range exportCommands {
		sc := addExportTo(subCommand)
		flags.AddAllStorageFlags(sc)
	}
}

const exportCommand = "export"

// The export category of commands.
// `corso export [<subcommand>] [<flag>...]`
func exportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   exportCommand,
		Short: "Export your service data",
		Long:  `Export the data stored in one of your M365 services.`,
		RunE:  handleExportCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso export`.
// Produces the same output as `corso export --help`.
func handleExportCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func runExport(
	ctx context.Context,
	cmd *cobra.Command,
	args []string,
	ueco utils.ExportCfgOpts,
	sel selectors.Selector,
	backupID, serviceName string,
	acceptedFormatTypes []string,
) error {
	if err := utils.ValidateExportConfigFlags(&ueco, acceptedFormatTypes); err != nil {
		return Only(ctx, err)
	}

	r, _, err := utils.GetAccountAndConnect(ctx, cmd, sel.PathService())
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	exportLocation := args[0]
	if len(exportLocation) == 0 {
		// This should not be possible, but adding it just in case.
		exportLocation = control.DefaultRestoreLocation + dttm.FormatNow(dttm.HumanReadableDriveItem)
	}

	Infof(ctx, "Exporting to folder %s", exportLocation)

	eo, err := r.NewExport(
		ctx,
		backupID,
		sel,
		utils.MakeExportConfig(ctx, ueco))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to initialize "+serviceName+" export"))
	}

	expColl, err := eo.Run(ctx)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return Only(ctx, clues.New("Backup or backup details missing for id "+backupID))
		}

		return Only(ctx, clues.Wrap(err, "Failed to run "+serviceName+" export"))
	}

	// It would be better to give a progressbar than a spinner, but we
	// have any way of knowing how many files are available as of now.
	diskWriteComplete := observe.MessageWithCompletion(ctx, "Writing data to disk")

	err = export.ConsumeExportCollections(ctx, exportLocation, expColl, eo.Errors)

	// The progressbar has to be closed before we move on as the Infof
	// below flushes progressbar to prevent clobbering the output and
	// that causes the entire export operation to stall indefinitely.
	// https://github.com/alcionai/corso/blob/8102523dc62c001b301cd2ab4e799f86146ab1a0/src/cli/print/print.go#L151
	close(diskWriteComplete)

	if err != nil {
		return Only(ctx, err)
	}

	stats := eo.GetStats()
	if len(stats) > 0 {
		Infof(ctx, "\nExport details")
	}

	for k, s := range stats {
		Infof(ctx, "%s: %d items (%s)", k.HumanString(), s.ResourceCount, humanize.Bytes(uint64(s.BytesRead)))
	}

	return nil
}
