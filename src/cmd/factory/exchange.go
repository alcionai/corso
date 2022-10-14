package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/pkg/control"
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

func addExchangeCommands(parent *cobra.Command) {
	parent.AddCommand(emailsCmd)
	parent.AddCommand(eventsCmd)
	parent.AddCommand(contactsCmd)
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

	items := make([]item, 0, count)

	for i := 0; i < count; i++ {
		var (
			now       = common.Now()
			nowLegacy = common.FormatLegacyTime(time.Now())
			id        = uuid.NewString()
			subject   = "automated " + now[:16] + " - " + id[:8]
			body      = "automated mail generation for " + user + " at " + now + " - " + id
		)

		items = append(items, item{
			name: id,
			// TODO: allow flags that specify a different "from" user, rather than duplicating
			data: mockconnector.GetMockMessageWith(
				user, user, user,
				subject, body,
				nowLegacy, nowLegacy, nowLegacy, nowLegacy),
		})
	}

	collections := []collection{{
		pathElements: []string{destination},
		category:     category,
		items:        items,
	}}

	// TODO: fit the desination to the containers
	dest := control.DefaultRestoreDestination(common.SimpleTimeTesting)
	dest.ContainerName = destination

	dataColls, err := buildCollections(
		service,
		tenantID, user,
		dest,
		collections,
	)
	if err != nil {
		return Only(ctx, err)
	}

	Infof(ctx, "Generating %d emails in %s\n", count, destination)

	sel := selectors.NewExchangeRestore().Selector

	deets, err := gc.RestoreDataCollections(ctx, sel, dest, dataColls)
	if err != nil {
		return Only(ctx, err)
	}

	deets.PrintEntries(ctx)

	return nil
}

func handleExchangeCalendarEventFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// generate mocked events

	return nil
}

func handleExchangeContactFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// generate mocked contacts

	return nil
}
