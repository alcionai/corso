package tester

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/logger"
)

// StubRootCmd builds a stub cobra command to be used as
// the root command for integration testing on the CLI
func StubRootCmd(args ...string) *cobra.Command {
	id := uuid.NewString()
	now := common.FormatTime(time.Now())
	cmdArg := "testing-corso"
	c := &cobra.Command{
		Use:   cmdArg,
		Short: id,
		Long:  id + " - " + now,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "test command args: %+v", args)
			return nil
		},
	}
	c.SetArgs(args)

	return c
}

func NewContext() (context.Context, func()) {
	ctx, _ := logger.SeedLevel(context.Background(), logger.Development)
	return ctx, func() { logger.Flush(ctx) }
}

func WithContext(ctx context.Context) (context.Context, func()) {
	ctx, _ = logger.SeedLevel(ctx, logger.Development)
	return ctx, func() { logger.Flush(ctx) }
}
