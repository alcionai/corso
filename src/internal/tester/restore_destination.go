package tester

import (
	"context"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/control"
)

func DefaultTestRestoreDestination(ctx context.Context) control.RestoreDestination {
	// Use microsecond granularity to help reduce collisions.
	return control.DefaultRestoreDestination(ctx, common.SimpleTimeTesting)
}
