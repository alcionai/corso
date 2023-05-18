package tester

import (
	"strings"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/control"
)

const RestoreFolderPrefix = "Corso_Test"

func DefaultTestRestoreDestination(namespace string) control.RestoreDestination {
	var (
		dest = control.DefaultRestoreDestination(dttm.SafeForTesting)
		sft  = dttm.FormatNow(dttm.SafeForTesting)
	)

	parts := []string{RestoreFolderPrefix, namespace, sft}
	if len(namespace) == 0 {
		parts = []string{RestoreFolderPrefix, sft}
	}

	dest.ContainerName = strings.Join(parts, "_")

	return dest
}
