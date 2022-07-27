package print

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"github.com/tomlazar/table"
)

var (
	outputAsJSON      bool
	outputAsJSONDebug bool
)

// adds the --output flag to the provided command.
func AddOutputFlag(parent *cobra.Command) {
	fs := parent.PersistentFlags()
	fs.BoolVar(&outputAsJSON, "json", false, "output data in JSON format")
	fs.BoolVar(&outputAsJSONDebug, "json-debug", false, "output all internal and debugging data in JSON format")
	fs.MarkHidden("json-debug")
}

type Printable interface {
	// reduces the struct to a minimized format for easier human consumption
	MinimumPrintable() any
	// should list the property names of the values surfaced in Values()
	Headers() []string
	// list of values for tabular or csv formatting
	// if the backing data is nil or otherwise missing,
	// values should provide an empty string as opposed to skipping entries
	Values() []string
}

// Prints the backups to the terminal with stdout.
func Backups(bs []backup.Backup) {
	ps := []Printable{}
	for _, b := range bs {
		ps = append(ps, b)
	}
	printAll(ps)
}

// Prints the entries to the terminal with stdout.
func Entries(des []details.DetailsEntry) {
	ps := []Printable{}
	for _, de := range des {
		ps = append(ps, de)
	}
	printAll(ps)
}

// printAll prints the slice of printable items,
// according to the caller's requested format.
func printAll(ps []Printable) {
	if len(ps) == 0 {
		return
	}
	if outputAsJSON || outputAsJSONDebug {
		outputJSON(ps, outputAsJSONDebug)
		return
	}
	outputTable(ps)
}

// output to stdout the list of printable structs as json.
// if debug is false, calls ps.Printable() first
func outputJSON(ps []Printable, debug bool) {
	sl := make([]any, 0, len(ps))
	for _, p := range ps {
		if debug {
			sl = append(sl, p)
		} else {
			sl = append(sl, p.MinimumPrintable())
		}
	}
	bs, err := json.Marshal(sl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error formatting results to json: %v\n", err)
		return
	}
	fmt.Println(string(pretty.Pretty(bs)))
}

// output to stdout the list of printable structs in a table
func outputTable(ps []Printable) {
	t := table.Table{
		Headers: ps[0].Headers(),
		Rows:    [][]string{},
	}
	for _, p := range ps {
		t.Rows = append(t.Rows, p.Values())
	}
	_ = t.WriteTable(
		os.Stdout,
		&table.Config{
			ShowIndex:       false,
			Color:           false,
			AlternateColors: false,
		})
}
