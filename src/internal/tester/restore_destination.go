package tester

import (
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/control"
)

func DefaultTestRestoreDestination() control.RestoreDestination {
	// Use microsecond granularity to help reduce collisions.
	return control.DefaultRestoreDestination(dttm.SafeForTesting)
}
