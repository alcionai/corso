package impl

import (
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var (
	emailsCmd = &cobra.Command{
		Use:   "emails",
		Short: "Generate exchange emails",
		RunE:  handleExchangeEmailFactory,
	}

	eventsCmd = &cobra.Command{
		Use:   "events",
		Short: "Generate exchange calendar events",
		RunE:  handleExchangeCalendarEventFactory,
	}

	contactsCmd = &cobra.Command{
		Use:   "contacts",
		Short: "Generate exchange contacts",
		RunE:  handleExchangeContactFactory,
	}
)

func AddExchangeCommands(cmd *cobra.Command) {
	cmd.AddCommand(emailsCmd)
	cmd.AddCommand(eventsCmd)
	cmd.AddCommand(contactsCmd)
}

func handleExchangeEmailFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.EmailCategory
		errs     = fault.New(false)
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctrl, _, _, err := getControllerAndVerifyResourceOwner(ctx, resource.Users, User)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		ctrl,
		service,
		category,
		selectors.NewExchangeRestore([]string{User}).Selector,
		Tenant, User, Destination,
		Count,
		func(id, now, subject, body string) []byte {
			return exchMock.MessageWith(
				User, User, User,
				subject, body, body,
				now, now, now, now)
		},
		control.Defaults(),
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

func handleExchangeCalendarEventFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.EventsCategory
		errs     = fault.New(false)
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctrl, _, _, err := getControllerAndVerifyResourceOwner(ctx, resource.Users, User)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		ctrl,
		service,
		category,
		selectors.NewExchangeRestore([]string{User}).Selector,
		Tenant, User, Destination,
		Count,
		func(id, now, subject, body string) []byte {
			return exchMock.EventWith(
				User, subject, body, body,
				exchMock.NoOriginalStartDate, now, now,
				exchMock.NoRecurrence, exchMock.NoAttendees,
				false, exchMock.NoCancelledOccurrences,
				exchMock.NoExceptionOccurrences)
		},
		control.Defaults(),
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

func handleExchangeContactFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.ContactsCategory
		errs     = fault.New(false)
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	ctrl, _, _, err := getControllerAndVerifyResourceOwner(ctx, resource.Users, User)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		ctrl,
		service,
		category,
		selectors.NewExchangeRestore([]string{User}).Selector,
		Tenant, User, Destination,
		Count,
		func(id, now, subject, body string) []byte {
			given, mid, sur := id[:8], id[9:13], id[len(id)-12:]

			return exchMock.ContactBytesWith(
				given+" "+sur,
				sur+", "+given,
				given, mid, sur,
				"123-456-7890",
			)
		},
		control.Defaults(),
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
