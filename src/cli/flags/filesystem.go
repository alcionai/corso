package flags

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/canario/src/pkg/storage"
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

	AddAzureCredsFlags(cmd)
	AddCorsoPassphaseFlags(cmd)

	fs.StringVar(
		&FilesystemPathFV,
		FilesystemPathFN,
		"",
		"path to local or network storage")
	cobra.CheckErr(cmd.MarkFlagRequired(FilesystemPathFN))
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
