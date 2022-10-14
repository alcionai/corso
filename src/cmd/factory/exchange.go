package main

import (
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
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
	Err(cmd.Context(), ErrNotYetImplemeted)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// generate mocked emails

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
	//nolint
	Err(cmd.Context(), ErrNotYetImplemeted)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// generate mocked contacts

	return nil
}
