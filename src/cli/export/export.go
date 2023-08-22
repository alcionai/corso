package export

import (
	"context"
	"errors"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
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
	addTeamsCommands,
}

// AddCommands attaches all `corso export * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	exportC := exportCmd()
	cmd.AddCommand(exportC)

	for _, addExportTo := range exportCommands {
		addExportTo(exportC)
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
) error {
	r, _, _, _, err := utils.GetAccountAndConnect(ctx, sel.PathService(), repo.S3Overrides(cmd))
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
	defer close(diskWriteComplete)

	err = export.ConsumeExportCollections(ctx, exportLocation, expColl, eo.Errors)
	if err != nil {
		return Only(ctx, err)
	}

	return nil
}
