package testdata

import (
	"strings"

	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/dttm"
)

const RestoreFolderPrefix = "Corso_Test"

func DefaultRestoreConfig(namespace string) control.RestoreConfig {
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
