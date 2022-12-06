package main

import (
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
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

func addExchangeCommands(cmd *cobra.Command) {
	cmd.AddCommand(emailsCmd)
	cmd.AddCommand(eventsCmd)
	cmd.AddCommand(contactsCmd)
}

func handleExchangeEmailFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.EmailCategory
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, tenantID, err := getGCAndVerifyUser(ctx, user)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		service,
		category,
		selectors.NewExchangeRestore().Selector,
		tenantID, user, destination,
		count,
		func(id, now, subject, body string) []byte {
			return mockconnector.GetMockMessageWith(
				user, user, user,
				subject, body, body,
				now, now, now, now)
		},
	)
	if err != nil {
		return Only(ctx, err)
	}

	deets.PrintEntries(ctx)

	return nil
}

func handleExchangeCalendarEventFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.EventsCategory
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, tenantID, err := getGCAndVerifyUser(ctx, user)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		service,
		category,
		selectors.NewExchangeRestore().Selector,
		tenantID, user, destination,
		count,
		func(id, now, subject, body string) []byte {
			return mockconnector.GetMockEventWith(
				user, subject, body, body,
				now, now, false)
		},
	)
	if err != nil {
		return Only(ctx, err)
	}

	deets.PrintEntries(ctx)

	return nil
}

func handleExchangeContactFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.ContactsCategory
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, tenantID, err := getGCAndVerifyUser(ctx, user)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		service,
		category,
		selectors.NewExchangeRestore().Selector,
		tenantID, user, destination,
		count,
		func(id, now, subject, body string) []byte {
			given, mid, sur := id[:8], id[9:13], id[len(id)-12:]

			return mockconnector.GetMockContactBytesWith(
				given+" "+sur,
				sur+", "+given,
				given, mid, sur,
				"123-456-7890",
			)
		},
	)
	if err != nil {
		return Only(ctx, err)
	}

	deets.PrintEntries(ctx)

	return nil
}
