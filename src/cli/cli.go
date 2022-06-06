package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/backup"
	"github.com/alcionai/corso/cli/repo"
)

// The root-level command.
// `corso <command> [<subcommand>] [<service>] [<flag>...]`
var corsoCmd = &cobra.Command{
	Use:   "corso",
	Short: "Protect your Microsoft 365 data.",
	Long:  `Reliable, secure, and efficient data protection for Microsoft 365.`,
	Run:   handleCorsoCmd,
}

// the root-level flags
var (
	version bool
)

// Handler for flat calls to `corso`.
// Produces the same output as `corso --help`.
func handleCorsoCmd(cmd *cobra.Command, args []string) {
	if version {
		fmt.Printf("Corso\nversion:\tpre-alpha\n")
	} else {
		cmd.Help()
	}
}

// Handle builds and executes the cli processor.
func Handle() {
	corsoCmd.Flags().BoolP("version", "v", version, "current version info")

	repo.AddCommands(corsoCmd)
	backup.AddCommands(corsoCmd)

	if err := corsoCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
