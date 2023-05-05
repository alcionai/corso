package impl

import (
	"strings"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var spFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Generate OneDrive files",
	RunE:  handleSharePointLibraryFileFactory,
}

func AddSharePointCommands(cmd *cobra.Command) {
	cmd.AddCommand(spFilesCmd)
}

func handleSharePointLibraryFileFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.SharePointService
		category = path.LibrariesCategory
		errs     = fault.New(false)
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, acct, inp, err := getGCAndVerifyResourceOwner(ctx, connector.Sites, Site)
	if err != nil {
		return Only(ctx, err)
	}

	sel := selectors.NewSharePointBackup([]string{Site}).Selector
	sel.SetDiscreteOwnerIDName(inp.ID(), inp.Name())

	deets, err := generateAndRestoreDriveItems(
		gc,
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
