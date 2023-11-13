package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/cmd/sanity_test/export"
	"github.com/alcionai/corso/src/cmd/sanity_test/restore"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// root command
// ---------------------------------------------------------------------------

func rootCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "sanity-test",
		Short:             "run the sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestRoot,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("running", cmd.UseLine())
		},
	}
}

func sanityTestRoot(cmd *cobra.Command, args []string) error {
	return print.Only(cmd.Context(), clues.New("must specify a kind of test"))
}

func main() {
	ls := logger.Settings{
		File:        logger.GetLogFile(""),
		Level:       logger.LLInfo,
		PIIHandling: logger.PIIPlainText,
	}

	ctx, log := logger.Seed(context.Background(), ls)
	defer func() {
		_ = log.Sync() // flush all logs in the buffer
	}()

	// TODO: only needed for exchange
	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	root := rootCMD()

	restCMD := restoreCMD()

	restCMD.AddCommand(restoreExchangeCMD())
	restCMD.AddCommand(restoreOneDriveCMD())
	restCMD.AddCommand(restoreSharePointCMD())
	restCMD.AddCommand(restoreGroupsCMD())
	root.AddCommand(restCMD)

	expCMD := exportCMD()

	expCMD.AddCommand(exportOneDriveCMD())
	expCMD.AddCommand(exportSharePointCMD())
	expCMD.AddCommand(exportGroupsCMD())
	root.AddCommand(expCMD)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

// ---------------------------------------------------------------------------
// restore/export command
// ---------------------------------------------------------------------------

func exportCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "export",
		Short:             "run the post-export sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestExport,
	}
}

func sanityTestExport(cmd *cobra.Command, args []string) error {
	return print.Only(cmd.Context(), clues.New("must specify a service"))
}

func restoreCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "restore",
		Short:             "run the post-restore sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestRestore,
	}
}

func sanityTestRestore(cmd *cobra.Command, args []string) error {
	return print.Only(cmd.Context(), clues.New("must specify a service"))
}

// ---------------------------------------------------------------------------
// service commands - export
// ---------------------------------------------------------------------------

func exportGroupsCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "groups",
		Short:             "run the groups export sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestExportGroups,
	}
}

func sanityTestExportGroups(cmd *cobra.Command, args []string) error {
	ctx := common.SetDebug(cmd.Context())
	envs := common.EnvVars(ctx)

	ac, err := common.GetAC()
	if err != nil {
		return print.Only(ctx, err)
	}

	export.CheckGroupsExport(ctx, ac, envs)

	return nil
}

func exportOneDriveCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "onedrive",
		Short:             "run the onedrive export sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestExportOneDrive,
	}
}

func sanityTestExportOneDrive(cmd *cobra.Command, args []string) error {
	ctx := common.SetDebug(cmd.Context())
	envs := common.EnvVars(ctx)

	ac, err := common.GetAC()
	if err != nil {
		return print.Only(ctx, err)
	}

	export.CheckOneDriveExport(ctx, ac, envs)

	return nil
}

func exportSharePointCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "sharepoint",
		Short:             "run the sharepoint export sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestExportSharePoint,
	}
}

func sanityTestExportSharePoint(cmd *cobra.Command, args []string) error {
	ctx := common.SetDebug(cmd.Context())
	envs := common.EnvVars(ctx)

	ac, err := common.GetAC()
	if err != nil {
		return print.Only(ctx, err)
	}

	export.CheckSharePointExport(ctx, ac, envs)

	return nil
}

// ---------------------------------------------------------------------------
// service commands - restore
// ---------------------------------------------------------------------------

func restoreExchangeCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "exchange",
		Short:             "run the exchange restore sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestRestoreExchange,
	}
}

func sanityTestRestoreExchange(cmd *cobra.Command, args []string) error {
	ctx := common.SetDebug(cmd.Context())
	envs := common.EnvVars(ctx)

	ac, err := common.GetAC()
	if err != nil {
		return print.Only(ctx, err)
	}

	restore.CheckEmailRestoration(ctx, ac, envs)

	return nil
}

func restoreOneDriveCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "onedrive",
		Short:             "run the onedrive restore sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestRestoreOneDrive,
	}
}

func sanityTestRestoreOneDrive(cmd *cobra.Command, args []string) error {
	ctx := common.SetDebug(cmd.Context())
	envs := common.EnvVars(ctx)

	ac, err := common.GetAC()
	if err != nil {
		return print.Only(ctx, err)
	}

	restore.CheckOneDriveRestoration(ctx, ac, envs)

	return nil
}

func restoreSharePointCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "sharepoint",
		Short:             "run the sharepoint restore sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestRestoreSharePoint,
	}
}

func sanityTestRestoreSharePoint(cmd *cobra.Command, args []string) error {
	ctx := common.SetDebug(cmd.Context())
	envs := common.EnvVars(ctx)

	ac, err := common.GetAC()
	if err != nil {
		return print.Only(ctx, err)
	}

	restore.CheckSharePointRestoration(ctx, ac, envs)

	return nil
}

func restoreGroupsCMD() *cobra.Command {
	return &cobra.Command{
		Use:               "groups",
		Short:             "run the groups restore sanity tests",
		DisableAutoGenTag: true,
		RunE:              sanityTestRestoreGroups,
	}
}

func sanityTestRestoreGroups(cmd *cobra.Command, args []string) error {
	ctx := common.SetDebug(cmd.Context())
	envs := common.EnvVars(ctx)

	ac, err := common.GetAC()
	if err != nil {
		return print.Only(ctx, err)
	}

	restore.CheckGroupsRestoration(ctx, ac, envs)

	return nil
}
