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
	BackupID string

	FolderPaths []string
	FileNames   []string

	FileCreatedAfter   string
	FileCreatedBefore  string
	FileModifiedAfter  string
	FileModifiedBefore string

	Library string
	Site    []string
	WebURL  []string

	User []string
)

// common flag names
const (
	BackupFN  = "backup"
	DataFN    = "data"
	LibraryFN = "library"
	SiteFN    = "site"    // site only accepts WebURL values
	SiteIDFN  = "site-id" // site-id accepts actual site ids
	UserFN    = "user"

	FileFN   = "file"
	FolderFN = "folder"

	FileCreatedAfterFN   = "file-created-after"
	FileCreatedBeforeFN  = "file-created-before"
	FileModifiedAfterFN  = "file-modified-after"
	FileModifiedBeforeFN = "file-modified-before"
)

// AddBackupIDFlag adds the --backup flag.
func AddBackupIDFlag(cmd *cobra.Command, require bool) {
	cmd.Flags().StringVar(&BackupID, BackupFN, "", "ID of the backup to retrieve.")

	if require {
		cobra.CheckErr(cmd.MarkFlagRequired(BackupFN))
	}
}

// AddUserFlag adds the --user flag.
func AddUserFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&User,
		UserFN, nil,
		"Backup a specific user's data; accepts '"+Wildcard+"' to select all users.")
	cobra.CheckErr(cmd.MarkFlagRequired(UserFN))
}

// AddSiteIDFlag adds the --site-id flag, which accepts site ID values.
// This flag is hidden, since we expect users to prefer the --site url
// and do not want to encourage confusion.
func AddSiteIDFlag(cmd *cobra.Command) {
	fs := cmd.Flags()

	// note string ARRAY var.  IDs naturally contain commas, so we cannot accept
	// duplicate values within a flag declaration.  ie: --site-id a,b,c does not
	// work.  Users must call --site-id a --site-id b --site-id c.
	fs.StringArrayVar(
		&Site,
		SiteIDFN, nil,
		//nolint:lll
		"Backup data by site ID; accepts '"+Wildcard+"' to select all sites.  Args cannot be comma-delimited and must use multiple flags.")
	cobra.CheckErr(fs.MarkHidden(SiteIDFN))
}

// AddSiteFlag adds the --site flag, which accepts webURL values.
func AddSiteFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&WebURL,
		SiteFN, nil,
		"Backup data by site URL; accepts '"+Wildcard+"' to select all sites.")
}

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
