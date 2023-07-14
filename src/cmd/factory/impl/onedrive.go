package impl

import (
	"strings"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var odFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Generate OneDrive files",
	RunE:  handleOneDriveFileFactory,
}

func AddOneDriveCommands(cmd *cobra.Command) {
	cmd.AddCommand(odFilesCmd)
}

func handleOneDriveFileFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.OneDriveService
		category = path.FilesCategory
		errs     = fault.New(false)
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctrl, acct, inp, err := getControllerAndVerifyResourceOwner(ctx, resource.Users, User, path.OneDriveService)
	if err != nil {
		return Only(ctx, err)
	}

	sel := selectors.NewOneDriveBackup([]string{User}).Selector
	sel.SetDiscreteOwnerIDName(inp.ID(), inp.Name())

	deets, err := generateAndRestoreDriveItems(
		ctrl,
		inp.ID(),
		SecondaryUser,
		strings.ToLower(SecondaryUser),
		acct,
		service,
		category,
		sel,
		Tenant,
		Destination,
		Count,
		errs,
		count.New())
	if err != nil {
		return Only(ctx, err)
	}

	for _, e := range errs.Recovered() {
		logger.CtxErr(ctx, err).Error(e.Error())
	}

	deets.PrintEntries(ctx)

	return nil
}
