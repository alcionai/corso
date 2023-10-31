package testdata

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/cli/flags"
)

func PreparedLibraryFlags() []string {
	return []string{
		"--" + flags.LibraryFN, LibraryInput,
		"--" + flags.FolderFN, FlgInputs(FolderPathInput),
		"--" + flags.FileFN, FlgInputs(FileNameInput),
		"--" + flags.FileCreatedAfterFN, FileCreatedAfterInput,
		"--" + flags.FileCreatedBeforeFN, FileCreatedBeforeInput,
		"--" + flags.FileModifiedAfterFN, FileModifiedAfterInput,
		"--" + flags.FileModifiedBeforeFN, FileModifiedBeforeInput,
	}
}

func AssertLibraryFlags(t *testing.T, cmd *cobra.Command) {
	assert.Equal(t, LibraryInput, flags.LibraryFV)
	assert.Equal(t, FolderPathInput, flags.FolderPathFV)
	assert.Equal(t, FileNameInput, flags.FileNameFV)
	assert.Equal(t, FileCreatedAfterInput, flags.FileCreatedAfterFV)
	assert.Equal(t, FileCreatedBeforeInput, flags.FileCreatedBeforeFV)
	assert.Equal(t, FileModifiedAfterInput, flags.FileModifiedAfterFV)
	assert.Equal(t, FileModifiedBeforeInput, flags.FileModifiedBeforeFV)
}
