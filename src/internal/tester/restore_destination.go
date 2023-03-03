package tester

import (
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/control"
)

func DefaultTestRestoreDestination() control.RestoreDestination {
	// Use microsecond granularity to help reduce collisions.
	return control.DefaultRestoreDestination(common.SimpleTimeTesting)
}
