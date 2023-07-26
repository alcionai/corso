package export

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

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
	if exportLocation == "" {
		// This should not be possible, but adding it just in case.
		exportLocation = control.DefaultRestoreLocation + dttm.FormatNow(dttm.HumanReadableDriveItem)
	}

	Infof(ctx, "Exporting to folder %s", exportLocation)

	eo, err := r.NewExport(
		ctx,
		backupID,
		sel,
		utils.MakeExportConfig(ctx, ueco),
	)
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
	// have know way of knowing how many files are available as of now.
	diskWriteComplete := observe.MessageWithCompletion(ctx, "Writing data to disk")
	defer func() {
		diskWriteComplete <- struct{}{}
		close(diskWriteComplete)
	}()

	err = writeExportCollections(ctx, exportLocation, expColl)
	if err != nil {
		return err
	}

	return nil
}

func writeExportCollections(
	ctx context.Context,
	exportLocation string,
	expColl []export.Collection,
) error {
	for _, col := range expColl {
		folder := filepath.Join(exportLocation, col.BasePath())

		for item := range col.Items(ctx) {
			err := item.Error
			if err != nil {
				return Only(ctx, clues.Wrap(err, "getting item").With("dir_name", folder))
			}

			err = writeExportItem(ctx, item, folder)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// writeExportItem writes an ExportItem to disk in the specified folder.
func writeExportItem(ctx context.Context, item export.Item, folder string) error {
	name := item.Data.Name
	fpath := filepath.Join(folder, name)

	progReader, pclose := observe.ItemSpinner(
		ctx,
		item.Data.Body,
		observe.ItemExportMsg,
		clues.Hide(name))

	defer item.Data.Body.Close()
	defer pclose()

	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "creating directory").With("dir_name", folder))
	}

	// In case the user tries to restore to a non-clean
	// directory, we might run into collisions an fail.
	f, err := os.Create(fpath)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "creating file").With("file_name", name, "file_dir", folder))
	}

	_, err = io.Copy(f, progReader)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "writing file").With("file_name", name, "file_dir", folder))
	}

	return nil
}
