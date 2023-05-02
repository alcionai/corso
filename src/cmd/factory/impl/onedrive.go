package impl

import (
	"strings"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Generate OneDrive files",
	RunE:  handleOneDriveFileFactory,
}

func AddOneDriveCommands(cmd *cobra.Command) {
	cmd.AddCommand(filesCmd)
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

	gc, acct, inp, err := getGCAndVerifyUser(ctx, User)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreOnedriveItems(
		gc,
		User,
		inp.ID(),
		strings.ToLower(SecondaryUser),
		acct,
		service,
		category,
		selectors.NewOneDriveBackup([]string{User}).Selector,
		Tenant,
		Destination,
		Count,
		errs)
	if err != nil {
		return Only(ctx, err)
	}

	for _, e := range errs.Recovered() {
		logger.CtxErr(ctx, err).Error(e.Error())
	}

	deets.PrintEntries(ctx)

	return nil
}
