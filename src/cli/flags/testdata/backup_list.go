package testdata

import (
	"testing"

	"github.com/spf13/cobra"
	"gotest.tools/v3/assert"

	"github.com/alcionai/corso/src/cli/flags"
)

func PreparedBackupListFlags() []string {
	return []string{
		"--" + flags.FailedItemsFN, flags.Show,
		"--" + flags.SkippedItemsFN, flags.Show,
		"--" + flags.RecoveredErrorsFN, flags.Show,
	}
}

func AssertBackupListFlags(t *testing.T, cmd *cobra.Command) {
	assert.Equal(t, flags.Show, flags.ListFailedItemsFV)
	assert.Equal(t, flags.Show, flags.ListSkippedItemsFV)
	assert.Equal(t, flags.Show, flags.ListRecoveredErrorsFV)
}
