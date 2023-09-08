package impl

import (
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/spf13/cobra"
)

var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Generate groups channel messages",
	RunE:  handleGroupChannelFactory,
}

func AddGroupsCommands(cmd *cobra.Command) {
	cmd.AddCommand(channelCmd)
}

func handleGroupChannelFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx  = cmd.Context()
		errs = fault.New(false)
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctrl, acct, inp, err := getControllerAndVerifyResourceOwner(ctx, resource.Groups, Group, path.GroupsService)

	if err != nil {
		return Only(ctx, err)
	}

	sel := selectors.NewGroupsBackup([]string{Group}).Selector
	sel.SetDiscreteOwnerIDName(inp.ID(), inp.Name())

	deets, err := generateAndCreateChannelItems(
		ctrl,
		inp,
		acct,
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
