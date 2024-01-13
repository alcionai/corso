package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
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
}
