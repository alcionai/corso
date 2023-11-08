package debug

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var subCommandFuncs = []func() *cobra.Command{
	metadataFilesCmd,
}

var debugCommands = []func(cmd *cobra.Command) *cobra.Command{
	addOneDriveCommands,
	addSharePointCommands,
	addGroupsCommands,
	addExchangeCommands,
}

// AddCommands attaches all `corso debug * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	debugC, _ := utils.AddCommand(cmd, debugCmd(), utils.MarkDebugCommand())

	for _, sc := range subCommandFuncs {
		subCommand := sc()
		utils.AddCommand(debugC, subCommand, utils.MarkDebugCommand())

		for _, addTo := range debugCommands {
			servCmd := addTo(subCommand)
			flags.AddAllProviderFlags(servCmd)
			flags.AddAllStorageFlags(servCmd)
		}
	}
}

// ---------------------------------------------------------------------------
// Commands
// ---------------------------------------------------------------------------

const debugCommand = "debug"

// The debug category of commands.
// `corso debug [<subcommand>] [<flag>...]`
func debugCmd() *cobra.Command {
	return &cobra.Command{
		Use:   debugCommand,
		Short: "debugging & troubleshooting utilities",
		Long:  `debug the data stored in corso.`,
		RunE:  handledebugCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso debug`.
// Produces the same output as `corso debug --help`.
func handledebugCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The debug metadataFiles subcommand.
// `corso debug metadata-files <service> [<flag>...]`
var metadataFilesCommand = "metadata-files"

func metadataFilesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   metadataFilesCommand,
		Short: "display all the metadata file contents stored by the service",
		RunE:  handleMetadataFilesCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso debug metadata-files`.
// Produces the same output as `corso debug metadata-files --help`.
func handleMetadataFilesCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// ---------------------------------------------------------------------------
// runners
// ---------------------------------------------------------------------------

type base struct {
	snapshotID manifest.ID
	reasons    []identity.Reasoner
}

func (b base) GetReasons() []identity.Reasoner {
	return b.reasons
}

func (b base) GetSnapshotID() manifest.ID {
	return b.snapshotID
}

type metadataFile struct {
	name string
	path string
	data any
}

type mdDeserialize func(
	ctx context.Context,
	metadataCollections []data.RestoreCollection,
) ([]metadataFile, error)

func genericMetadataFiles(
	ctx context.Context,
	cmd *cobra.Command,
	args []string,
	sel selectors.Selector,
	backupID string,
	metadataDeserializer mdDeserialize,
) error {
	ctx = clues.Add(ctx, "backup_id", backupID)

	r, repoDeets, err := utils.GetAccountAndConnect(ctx, cmd, sel.PathService())
	if err != nil {
		return Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	bup, err := r.Backup(ctx, backupID)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "looking up backup"))
	}

	// read metadata
	sel = sel.SetDiscreteOwnerIDName(bup.ResourceOwnerID, bup.ResourceOwnerName)

	reasons, err := sel.Reasons(repoDeets.Repo.Account.ID(), false)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "constructing backup reasons"))
	}

	rp := r.DataStore()

	paths, err := r.DataProvider().GetMetadataPaths(
		ctx,
		rp,
		&base{manifest.ID(bup.SnapshotID), reasons},
		fault.New(true))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "retrieving metadata files"))
	}

	colls, err := rp.ProduceRestoreCollections(
		ctx,
		bup.SnapshotID,
		paths,
		nil,
		fault.New(true))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "looking up metadata file content"))
	}

	Info(ctx, "Reading Metadata From:")

	for _, coll := range colls {
		Infof(ctx, "%s", coll.FullPath())
	}

	// print metadata
	files, err := metadataDeserializer(ctx, colls)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "deserializing metadata file content"))
	}

	for _, file := range files {
		Infof(ctx, "\n------------------------------")
		Info(ctx, file.name)
		Info(ctx, file.path)
		Pretty(ctx, file.data)
	}

	return nil
}
