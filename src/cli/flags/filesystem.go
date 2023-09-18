package flags

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/storage"
)

// filesystem flag names
const (
	FilesystemPathFN = "path"
)

// filesystem flag values
var (
	FilesystemPathFV string
)

func AddFilesystemFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	fs.StringVar(
		&FilesystemPathFV,
		FilesystemPathFN,
		"",
		"path to local or network storage")
	cobra.CheckErr(cmd.MarkFlagRequired(FilesystemPathFN))

	fs.BoolVar(
		&SucceedIfExistsFV,
		SucceedIfExistsFN,
		false,
		"Exit with success if the repo has already been initialized.")
	cobra.CheckErr(fs.MarkHidden("succeed-if-exists"))
}

func FilesystemFlagOverrides(cmd *cobra.Command) map[string]string {
	fs := GetPopulatedFlags(cmd)
	return PopulateFilesystemFlags(fs)
}

func PopulateFilesystemFlags(flagset PopulatedFlags) map[string]string {
	fsOverrides := map[string]string{
		storage.StorageProviderTypeKey: storage.ProviderFilesystem.String(),
	}

	if _, ok := flagset[FilesystemPathFN]; ok {
		fsOverrides[FilesystemPathFN] = FilesystemPathFV
	}

	return fsOverrides
}
