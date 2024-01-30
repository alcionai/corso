package testdata

import (
	"testing"

	"github.com/spf13/cobra"
)

func PreparedTeamsChatsFlags() []string {
	return []string{
		// FIXME: populate when adding filters
		// "--" + flags.ChatCreatedAfterFN, ChatCreatedAfterInput,
		// "--" + flags.ChatCreatedBeforeFN, ChatCreatedBeforeInput,
		// "--" + flags.ChatLastMessageAfterFN, ChatLastMessageAfterInput,
		// "--" + flags.ChatLastMessageBeforeFN, ChatLastMessageBeforeInput,
	}
}

func AssertTeamsChatsFlags(t *testing.T, cmd *cobra.Command) {
	// FIXME: populate when adding filters
	// assert.Equal(t, ChatCreatedAfterInput, flags.ChatCreatedAfterFV)
	// assert.Equal(t, ChatCreatedBeforeInput, flags.ChatCreatedBeforeFV)
	// assert.Equal(t, ChatLastMessageAfterInput, flags.ChatLastMessageAfterFV)
	// assert.Equal(t, ChatLastMessageBeforeInput, flags.ChatLastMessageBeforeFV)
}
