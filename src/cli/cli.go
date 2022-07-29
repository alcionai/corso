package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/cli/backup"
	"github.com/alcionai/corso/cli/config"
	"github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/cli/repo"
	"github.com/alcionai/corso/cli/restore"
	"github.com/alcionai/corso/pkg/logger"
)

// The root-level command.
// `corso <command> [<subcommand>] [<service>] [<flag>...]`
var corsoCmd = &cobra.Command{
	Use:   "corso",
	Short: "Protect your Microsoft 365 data.",
	Long:  `Reliable, secure, and efficient data protection for Microsoft 365.`,
	RunE:  handleCorsoCmd,
}

// the root-level flags
var (
	cfgFile string
	version bool
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := config.InitConfig(cfgFile)
	cobra.CheckErr(err)

	if err := viper.ReadInConfig(); err == nil {
		print.Info("Using config file:", viper.ConfigFileUsed())
	}
}

// Handler for flat calls to `corso`.
// Produces the same output as `corso --help`.
func handleCorsoCmd(cmd *cobra.Command, args []string) error {
	if version {
		print.Infof("Corso\nversion:\tpre-alpha\n")
		return nil
	}
	return cmd.Help()
}

// Handle builds and executes the cli processor.
func Handle() {
	corsoCmd.Flags().BoolP("version", "v", version, "current version info")
	corsoCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", "config file (default is $HOME/.corso)")
	print.SetRootCommand(corsoCmd)
	print.AddOutputFlag(corsoCmd)

	corsoCmd.CompletionOptions.DisableDefaultCmd = true

	repo.AddCommands(corsoCmd)
	backup.AddCommands(corsoCmd)
	restore.AddCommands(corsoCmd)

	ctx, log := logger.Seed(context.Background())
	defer func() {
		_ = log.Sync() // flush all logs in the buffer
	}()

	if err := corsoCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
