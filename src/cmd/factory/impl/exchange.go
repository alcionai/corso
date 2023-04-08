package impl

import (
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
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

	gc, acct, err := getGCAndVerifyUser(ctx, User)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		acct,
		service,
		category,
		selectors.NewExchangeRestore([]string{User}).Selector,
		Tenant, User, Destination,
		Count,
		func(id, now, subject, body string) []byte {
			return mockconnector.GetMockMessageWith(
				User, User, User,
				subject, body, body,
				now, now, now, now)
		},
		control.Options{},
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

	gc, acct, err := getGCAndVerifyUser(ctx, User)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		acct,
		service,
		category,
		selectors.NewExchangeRestore([]string{User}).Selector,
		Tenant, User, Destination,
		Count,
		func(id, now, subject, body string) []byte {
			return mockconnector.GetMockEventWith(
				User, subject, body, body,
				now, now, `null`, `[]`, false)
		},
		control.Options{},
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

	gc, acct, err := getGCAndVerifyUser(ctx, User)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		acct,
		service,
		category,
		selectors.NewExchangeRestore([]string{User}).Selector,
		Tenant, User, Destination,
		Count,
		func(id, now, subject, body string) []byte {
			given, mid, sur := id[:8], id[9:13], id[len(id)-12:]

			return mockconnector.GetMockContactBytesWith(
				given+" "+sur,
				sur+", "+given,
				given, mid, sur,
				"123-456-7890",
			)
		},
		control.Options{},
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
