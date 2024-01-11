package utils

import (
	"reflect"
	"strconv"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
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
	}

	switch opts := opts.(type) {
	case GroupsOpts:
		for _, flag := range timeFlags {
			if _, ok := opts.Populated[flag]; ok {
				timeField := reflect.ValueOf(opts).FieldByName(flag).String()
				if !IsValidTimeFormat(timeField) {
					return clues.New("invalid time format for " + flag)
				}
			}
		}
	case SharePointOpts:
		for _, flag := range timeFlags {
			if _, ok := opts.Populated[flag]; ok {
				timeField := reflect.ValueOf(opts).FieldByName(flag).String()
				if !IsValidTimeFormat(timeField) {
					return clues.New("invalid time format for " + flag)
				}
			}
		}
	}

	return nil
}
