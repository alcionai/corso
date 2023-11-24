package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var CategoryDataFV []string

const CategoryDataFN = "data"

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

	// TODO(meain): This is a hacky way to get it to autocomplete multiple items
	cobra.CheckErr(cmd.RegisterFlagCompletionFunc(
		CategoryDataFN,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) ([]string, cobra.ShellCompDirective) {
			added := strings.Split(toComplete, ",")
			last := added[len(added)-1]
			added = added[:len(added)-1]

			if slices.Contains(added, last) {
				added = append(added, last)
				last = ""
			}

			pending := make([]string, 0, len(allowed)-len(added))
			for _, a := range allowed {
				if !slices.Contains(added, a) && strings.HasPrefix(a, last) {
					pending = append(pending, a)
				}
			}

			completions := []string{}
			for _, p := range pending {
				completions = append(completions, strings.Join(append(added, p), ","))
			}

			return completions, cobra.ShellCompDirectiveNoSpace
		}))
}
