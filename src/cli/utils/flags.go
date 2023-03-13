package utils

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/path"
)

// common flag vars
var (
	FolderPaths []string
	FileNames   []string

	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	Library string
)

// common flag names
const (
	BackupFN  = "backup"
	DataFN    = "data"
	LibraryFN = "library"
	SiteFN    = "site"
	UserFN    = "user"

	FileFN   = "file"
	FolderFN = "folder"

	FileCreatedAfterFN   = "file-created-after"
	FileCreatedBeforeFN  = "file-created-before"
	FileModifiedAfterFN  = "file-modified-after"
	FileModifiedBeforeFN = "file-modified-before"
)

type PopulatedFlags map[string]struct{}

func (fs PopulatedFlags) populate(pf *pflag.Flag) {
	if pf == nil {
		return
	}

	if pf.Changed {
		fs[pf.Name] = struct{}{}
	}
}

// GetPopulatedFlags returns a map of flags that have been
// populated by the user.  Entry keys match the flag's long
// name.  Values are empty.
func GetPopulatedFlags(cmd *cobra.Command) PopulatedFlags {
	pop := PopulatedFlags{}

	fs := cmd.Flags()
	if fs == nil {
		return pop
	}

	fs.VisitAll(pop.populate)

	return pop
}

// IsValidTimeFormat returns true if the input is recognized as a
// supported format by the common time parser.
func IsValidTimeFormat(in string) bool {
	_, err := common.ParseTime(in)
	return err == nil
}

// IsValidTimeFormat returns true if the input is recognized as a
// boolean.
func IsValidBool(in string) bool {
	_, err := strconv.ParseBool(in)
	return err == nil
}

// trimFolderSlash takes a set of folder paths and returns a set of folder paths
// with any unescaped trailing `/` characters removed.
func trimFolderSlash(folders []string) []string {
	res := make([]string, 0, len(folders))

	for _, p := range folders {
		// Use path package because it has logic to handle escaping already.
		res = append(res, path.TrimTrailingSlash(p))
	}

	return res
}
