package utils

import (
	"errors"
	"strconv"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/pkg/dttm"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/selectors"
)

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
		if p == string(path.PathSeparator) {
			res = selectors.Any()
			break
		}

		// Use path package because it has logic to handle escaping already.
		res = append(res, path.TrimTrailingSlash(p))
	}

	return res
}

func validateCommonTimeFlags(opts any) error {
	timeFlags := []string{
		flags.FileCreatedAfterFN,
		flags.FileCreatedBeforeFN,
		flags.FileModifiedAfterFN,
		flags.FileModifiedBeforeFN,
		flags.ListCreatedAfterFN,
		flags.ListCreatedBeforeFN,
		flags.ListModifiedAfterFN,
		flags.ListModifiedBeforeFN,
	}

	isFlagPopulated := func(opts any, flag string) bool {
		switch opts := opts.(type) {
		case GroupsOpts:
			_, ok := opts.Populated[flag]
			return ok
		case SharePointOpts:
			_, ok := opts.Populated[flag]
			return ok
		default:
			return false
		}
	}

	getTimeField := func(opts any, flag string) (string, error) {
		switch opts := opts.(type) {
		case GroupsOpts:
			return opts.GetFileTimeField(flag), nil
		case SharePointOpts:
			return opts.GetFileTimeField(flag), nil
		default:
			return "", errors.New("unsupported type")
		}
	}

	for _, flag := range timeFlags {
		if populated := isFlagPopulated(opts, flag); populated {
			timeField, err := getTimeField(opts, flag)
			if err != nil {
				return err
			}

			if !IsValidTimeFormat(timeField) {
				return clues.New("invalid time format for " + flag)
			}
		}
	}

	return nil
}
