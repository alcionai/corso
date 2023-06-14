package tester

import (
	"strings"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/control"
)

const RestoreFolderPrefix = "Corso_Test"

func DefaultTestRestoreConfig(namespace string) control.RestoreConfig {
	var (
		restoreCfg = control.DefaultRestoreConfig(dttm.SafeForTesting)
		sft        = dttm.FormatNow(dttm.SafeForTesting)
	)

	parts := []string{RestoreFolderPrefix, namespace, sft}
	if len(namespace) == 0 {
		parts = []string{RestoreFolderPrefix, sft}
	}

	restoreCfg.Location = strings.Join(parts, "_")

	return restoreCfg
}
