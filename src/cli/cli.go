package cli

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/backup"
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/help"
	"github.com/alcionai/corso/src/cli/options"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/repo"
	"github.com/alcionai/corso/src/cli/restore"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/logger"
)

var version = "dev"

// ------------------------------------------------------------------------------------------
// Corso Command
// ------------------------------------------------------------------------------------------

// The root-level command.
// `corso <command> [<subcommand>] [<service>] [<flag>...]`
var corsoCmd = &cobra.Command{
	Use:               "corso",
	Short:             "Free, Secure, Open-Source Backup for M365.",
	Long:              `Free, Secure, and Open-Source Backup for Microsoft 365.`,
	RunE:              handleCorsoCmd,
	PersistentPreRunE: config.InitFunc(),
}

// Handler for flat calls to `corso`.
// Produces the same output as `corso --help`.
func handleCorsoCmd(cmd *cobra.Command, args []string) error {
	v, _ := cmd.Flags().GetBool("version")
	if v {
		print.Outf(cmd.Context(), "Corso\nversion: "+version)
		return nil
	}

	return cmd.Help()
}

// CorsoCommand produces a copy of the cobra command used by Corso.
// The command tree is built and attached to the returned command.
func CorsoCommand() *cobra.Command {
	c := &cobra.Command{}
	*c = *corsoCmd
	BuildCommandTree(c)

	return c
}

// BuildCommandTree builds out the command tree used by the Corso library.
func BuildCommandTree(cmd *cobra.Command) {
	// want to order flags explicitly
	cmd.PersistentFlags().SortFlags = false

	cmd.Flags().BoolP("version", "v", false, "current version info")
	cmd.PersistentPostRunE = config.InitFunc()
	config.AddConfigFlags(cmd)
	logger.AddLogLevelFlag(cmd)
	observe.AddProgressBarFlags(cmd)
	print.AddOutputFlag(cmd)
	options.AddGlobalOperationFlags(cmd)

	cmd.SetUsageTemplate(indentExamplesTemplate(corsoCmd.UsageTemplate()))

	cmd.CompletionOptions.DisableDefaultCmd = true

	repo.AddCommands(cmd)
	backup.AddCommands(cmd)
	restore.AddCommands(cmd)
	help.AddCommands(cmd)
}

// ------------------------------------------------------------------------------------------
// Running Corso
// ------------------------------------------------------------------------------------------

// Handle builds and executes the cli processor.
func Handle() {
	ctx := config.Seed(context.Background())
	ctx = print.SetRootCmd(ctx, corsoCmd)
	observe.SeedWriter(ctx, print.StderrWriter(ctx), observe.PreloadFlags())

	BuildCommandTree(corsoCmd)

	ctx, log := logger.Seed(ctx, logger.PreloadLogLevel())
	defer func() {
		_ = log.Sync() // flush all logs in the buffer
		f, err := os.Create("mem.prof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		PrintMemUsage()
	}()

	if err := corsoCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

// Adjust the default usage template which does not properly indent examples
func indentExamplesTemplate(template string) string {
	cobra.AddTemplateFunc("indent", func(spaces int, v string) string {
		pad := strings.Repeat(" ", spaces)
		return pad + strings.Replace(v, "\n", "\n"+pad, -1)
	})

	e := regexp.MustCompile(`{{\.Example}}`)

	return e.ReplaceAllString(template, "{{.Example | indent 2}}")
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
