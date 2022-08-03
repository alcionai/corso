package print

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"github.com/tomlazar/table"

	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/backup/details"
)

var (
	outputAsJSON      bool
	outputAsJSONDebug bool
)

var rootCmd = &cobra.Command{}

func SetRootCommand(root *cobra.Command) {
	rootCmd = root
}

// adds the persistent flag --output to the provided command.
func AddOutputFlag(parent *cobra.Command) {
	fs := parent.PersistentFlags()
	fs.BoolVar(&outputAsJSON, "json", false, "output data in JSON format")
	fs.BoolVar(&outputAsJSONDebug, "json-debug", false, "output all internal and debugging data in JSON format")
	cobra.CheckErr(fs.MarkHidden("json-debug"))
}

// ---------------------------------------------------------------------------------------------------------
// Helper funcs
// ---------------------------------------------------------------------------------------------------------

// Only tells the CLI to only display this error, preventing the usage
// (ie, help) menu from displaying as well.
func Only(e error) error {
	rootCmd.SilenceUsage = true
	return e
}

// Err prints the params to cobra's error writer (stdErr by default)
// if s is nil, prints nothing.
// Prepends the message with "Error: "
func Err(s ...any) {
	err(rootCmd.ErrOrStderr(), s...)
}

// err is the testable core of Err()
func err(w io.Writer, s ...any) {
	if len(s) == 0 {
		return
	}
	msg := append([]any{"Error: "}, s...)
	fmt.Fprint(w, msg...)
}

// Info prints the params to cobra's error writer (stdErr by default)
// if s is nil, prints nothing.
func Info(s ...any) {
	info(rootCmd.ErrOrStderr(), s...)
}

// info is the testable core of Info()
func info(w io.Writer, s ...any) {
	if len(s) == 0 {
		return
	}
	fmt.Fprint(w, s...)
	fmt.Fprintf(w, "\n")
}

// Info prints the formatted strings to cobra's error writer (stdErr by default)
// if t is empty, prints nothing.
func Infof(t string, s ...any) {
	infof(rootCmd.ErrOrStderr(), t, s...)
}

// infof is the testable core of Infof()
func infof(w io.Writer, t string, s ...any) {
	if len(t) == 0 {
		return
	}
	fmt.Fprintf(w, t, s...)
	fmt.Fprintf(w, "\n")
}

// ---------------------------------------------------------------------------------------------------------
// Output control for backup list/details
// ---------------------------------------------------------------------------------------------------------

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

func print(p Printable) {
	if outputAsJSON || outputAsJSONDebug {
		outputJSON(p, outputAsJSONDebug)
		return
	}
	outputTable([]Printable{p})
}

// printAll prints the slice of printable items,
// according to the caller's requested format.
func printAll(ps []Printable) {
	if len(ps) == 0 {
		return
	}
	if outputAsJSON || outputAsJSONDebug {
		outputJSONArr(ps, outputAsJSONDebug)
		return
	}
	outputTable(ps)
}

// ------------------------------------------------------------------------------------------
// Type Formatters (TODO: migrate to owning packages)
// ------------------------------------------------------------------------------------------

// Prints the backup to the terminal with stdout.
func OutputBackup(b backup.Backup) {
	print(b)
}

// Prints the backups to the terminal with stdout.
func OutputBackups(bs []backup.Backup) {
	ps := []Printable{}
	for _, b := range bs {
		ps = append(ps, b)
	}
	printAll(ps)
}

// Prints the entries to the terminal with stdout.
func OutputEntries(des []details.DetailsEntry) {
	ps := []Printable{}
	for _, de := range des {
		ps = append(ps, de)
	}
	printAll(ps)
}

// ------------------------------------------------------------------------------------------
// Tabular
// ------------------------------------------------------------------------------------------

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
		rootCmd.OutOrStdout(),
		&table.Config{
			ShowIndex:       false,
			Color:           false,
			AlternateColors: false,
		})
}

// ------------------------------------------------------------------------------------------
// JSON
// ------------------------------------------------------------------------------------------

func outputJSON(p Printable, debug bool) {
	if debug {
		printJSON(p)
		return
	}
	printJSON(p.MinimumPrintable())
}

func outputJSONArr(ps []Printable, debug bool) {
	sl := make([]any, 0, len(ps))
	for _, p := range ps {
		if debug {
			sl = append(sl, p)
		} else {
			sl = append(sl, p.MinimumPrintable())
		}
	}
	printJSON(sl)
}

// output to stdout the list of printable structs as json.
func printJSON(a any) {
	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Fprintf(rootCmd.OutOrStderr(), "error formatting results to json: %v\n", err)
		return
	}
	fmt.Fprintln(
		rootCmd.OutOrStdout(),
		string(pretty.Pretty(bs)))
}
