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

var outputAsJSON bool

// adds the --output flag to the provided command.
func AddOutputFlag(parent *cobra.Command) {
	parent.PersistentFlags().BoolVar(&outputAsJSON, "json", false, "output data in JSON format")
}

type Printable interface {
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
	if outputAsJSON {
		outputJSON(ps)
		return
	}
	outputTable(ps)
}

// output to stdout the list of printable structs as json
func outputJSON(ps []Printable) {
	bs, err := json.Marshal(ps)
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
