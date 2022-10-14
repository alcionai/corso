package main

import (
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Generate OneDrive files",
	RunE:  handleOneDriveFileFactory,
}

func addOneDriveCommands(parent *cobra.Command) {
	parent.AddCommand(filesCmd)
}

func handleOneDriveFileFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	return nil
}
