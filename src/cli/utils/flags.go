package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

// common flag vars (eg: FV)
var (
	// RunMode describes the type of run, such as:
	// flagtest, dry, run.  Should default to 'run'.
	RunModeFV string

	BackupIDFV string

	FolderPathFV []string
	FileNameFV   []string

	FileCreatedAfterFV   string
	FileCreatedBeforeFV  string
	FileModifiedAfterFV  string
	FileModifiedBeforeFV string

	LibraryFV string
	SiteIDFV  []string
	WebURLFV  []string

	UserFV []string

	// for selection of data by category.  eg: `--data email,contacts`
	CategoryDataFV []string
)

// common flag names (eg: FN)
const (
	RunModeFN = "run-mode"

	BackupFN       = "backup"
	CategoryDataFN = "data"

	SiteFN    = "site"    // site only accepts WebURL values
	SiteIDFN  = "site-id" // site-id accepts actual site ids
	UserFN    = "user"
	MailBoxFN = "mailbox"

	LibraryFN = "library"
	FileFN    = "file"
	FolderFN  = "folder"

	FileCreatedAfterFN   = "file-created-after"
	FileCreatedBeforeFN  = "file-created-before"
	FileModifiedAfterFN  = "file-modified-after"
	FileModifiedBeforeFN = "file-modified-before"
)

// well-known flag values
const (
	RunModeFlagTest = "flag-test"
	RunModeRun      = "run"
)

// AddBackupIDFlag adds the --backup flag.
func AddBackupIDFlag(cmd *cobra.Command, require bool) {
	cmd.Flags().StringVar(&BackupIDFV, BackupFN, "", "ID of the backup to retrieve.")

	if require {
		cobra.CheckErr(cmd.MarkFlagRequired(BackupFN))
	}
}

func AddDataFlag(cmd *cobra.Command, allowed []string, hide bool) {
	var (
		allowedMsg string
		fs         = cmd.Flags()
	)

	switch len(allowed) {
	case 0:
		return
	case 1:
		allowedMsg = allowed[0]
	case 2:
		allowedMsg = fmt.Sprintf("%s or %s", allowed[0], allowed[1])
	default:
		allowedMsg = fmt.Sprintf(
			"%s or %s",
			strings.Join(allowed[:len(allowed)-1], ", "),
			allowed[len(allowed)-1])
	}

	fs.StringSliceVar(
		&CategoryDataFV,
		CategoryDataFN, nil,
		"Select one or more types of data to backup: "+allowedMsg+".")

	if hide {
		cobra.CheckErr(fs.MarkHidden(CategoryDataFN))
	}
}

// AddRunModeFlag adds the hidden --run-mode flag.
func AddRunModeFlag(cmd *cobra.Command, persistent bool) {
	fs := cmd.Flags()
	if persistent {
		fs = cmd.PersistentFlags()
	}

	fs.StringVar(&RunModeFV, RunModeFN, "run", "What mode to run: dry, test, run.  Defaults to run.")
	cobra.CheckErr(fs.MarkHidden(RunModeFN))
}

// AddUserFlag adds the --user flag.
func AddUserFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&UserFV,
		UserFN, nil,
		"Backup a specific user's data; accepts '"+Wildcard+"' to select all users.")
	cobra.CheckErr(cmd.MarkFlagRequired(UserFN))
}

// AddMailBoxFlag adds the --user and --mailbox flag.
func AddMailBoxFlag(cmd *cobra.Command) {
	flags := cmd.Flags()

	flags.StringSliceVar(
		&UserFV,
		UserFN, nil,
		"Backup a specific user's data; accepts '"+Wildcard+"' to select all users.")

	cobra.CheckErr(flags.MarkDeprecated(UserFN, fmt.Sprintf("use --%s instead", MailBoxFN)))

	flags.StringSliceVar(
		&UserFV,
		MailBoxFN, nil,
		"Backup a specific mailbox's data; accepts '"+Wildcard+"' to select all mailbox.")
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
		&SiteIDFV,
		SiteIDFN, nil,
		//nolint:lll
		"Backup data by site ID; accepts '"+Wildcard+"' to select all sites.  Args cannot be comma-delimited and must use multiple flags.")
	cobra.CheckErr(fs.MarkHidden(SiteIDFN))
}

// AddSiteFlag adds the --site flag, which accepts webURL values.
func AddSiteFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&WebURLFV,
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
	_, err := dttm.ParseTime(in)
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
		if len(p) == 1 && p[0] == path.PathSeparator {
			res = []string{}
			break
		}

		// Use path package because it has logic to handle escaping already.
		res = append(res, path.TrimTrailingSlash(p))
	}

	return res
}
