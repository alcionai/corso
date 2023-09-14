package flags

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	// Local / Network Attached Storage
	FilesystemPathFN = "path"
)

var FilesystemPathFV string

func FilesystemFlagOverrides(cmd *cobra.Command) map[string]string {
	fs := GetPopulatedFlags(cmd)
	return PopulateFilesystemFlags(fs)
}

func PopulateFilesystemFlags(flagset PopulatedFlags) map[string]string {
	fsOverrides := make(map[string]string)

	fsOverrides[storage.StorageProviderTypeKey] = storage.ProviderFilesystem.String()

	if _, ok := flagset[FilesystemPathFN]; ok {
		fsOverrides[FilesystemPathFN] = FilesystemPathFV
	}

	return fsOverrides
}
