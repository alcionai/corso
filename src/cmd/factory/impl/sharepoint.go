package impl

import (
	"strings"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	siteMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var (
	spFilesCmd = &cobra.Command{
		Use:   "files",
		Short: "Generate SharePoint files",
		RunE:  handleSharePointLibraryFileFactory,
	}

	spListsCmd = &cobra.Command{
		Use:   "lists",
		Short: "Generate SharePoint lists",
		RunE:  handleSharepointListsFactory,
	}
)

func AddSharePointCommands(cmd *cobra.Command) {
	cmd.AddCommand(spFilesCmd)
	cmd.AddCommand(spListsCmd)
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

	ctrl, acct, inp, err := getControllerAndVerifyResourceOwner(ctx, Site, path.SharePointService)
	if err != nil {
		return Only(ctx, err)
	}

	sel := selectors.NewSharePointBackup([]string{Site}).Selector
	sel.SetDiscreteOwnerIDName(inp.ID(), inp.Name())

	deets, err := generateAndRestoreDriveItems(
		ctrl,
		inp,
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

func handleSharepointListsFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.SharePointService
		category = path.ListsCategory
		errs     = fault.New(false)
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctrl, _, ins, err := getControllerAndVerifyResourceOwner(ctx, Site, path.SharePointService)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		ctrl,
		service,
		category,
		selectors.NewSharePointRestore([]string{ins.ID()}).Selector,
		Tenant, ins.ID(), Destination,
		Count,
		func(id, now, subject, body string) []byte {
			listBytes, err := siteMock.ListBytes(id)
			if err != nil {
				logger.CtxErr(ctx, err)
				return nil
			}
			return listBytes
		},
		control.DefaultOptions(),
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
