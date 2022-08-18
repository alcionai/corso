package print

import (
	"context"
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

type rootCmdCtx struct{}

// Adds a root cobra command to the context.
// Used to amend output controls like SilenceUsage or to retrieve
// the command's output writer.
func SetRootCmd(ctx context.Context, root *cobra.Command) context.Context {
	return context.WithValue(ctx, rootCmdCtx{}, root)
}

// Gets the root cobra command from the context.
// If no command is found, returns a new, blank command.
func getRootCmd(ctx context.Context) *cobra.Command {
	cmdIface := ctx.Value(rootCmdCtx{})
	cmd, ok := cmdIface.(*cobra.Command)
	if cmd == nil || !ok {
		return &cobra.Command{}
	}
	return cmd
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
func Only(ctx context.Context, e error) error {
	getRootCmd(ctx).SilenceUsage = true
	return e
}

// Err prints the params to cobra's error writer (stdErr by default)
// if s is nil, prints nothing.
// Prepends the message with "Error: "
func Err(ctx context.Context, s ...any) {
	err(getRootCmd(ctx).ErrOrStderr(), s...)
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
func Info(ctx context.Context, s ...any) {
	info(getRootCmd(ctx).ErrOrStderr(), s...)
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
func Infof(ctx context.Context, t string, s ...any) {
	infof(getRootCmd(ctx).ErrOrStderr(), t, s...)
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

//revive:disable:redefines-builtin-id
func print(w io.Writer, p Printable) {
	if outputAsJSON || outputAsJSONDebug {
		outputJSON(w, p, outputAsJSONDebug)
		return
	}
	outputTable(w, []Printable{p})
}

// printAll prints the slice of printable items,
// according to the caller's requested format.
func printAll(w io.Writer, ps []Printable) {
	if len(ps) == 0 {
		return
	}
	if outputAsJSON || outputAsJSONDebug {
		outputJSONArr(w, ps, outputAsJSONDebug)
		return
	}
	outputTable(w, ps)
}

// ------------------------------------------------------------------------------------------
// Type Formatters (TODO: migrate to owning packages)
// ------------------------------------------------------------------------------------------

// Prints the backup to the terminal with stdout.
func OutputBackup(ctx context.Context, b backup.Backup) {
	print(getRootCmd(ctx).OutOrStdout(), b)
}

// Prints the backups to the terminal with stdout.
func OutputBackups(ctx context.Context, bs []backup.Backup) {
	ps := []Printable{}
	for _, b := range bs {
		ps = append(ps, b)
	}
	printAll(getRootCmd(ctx).OutOrStdout(), ps)
}

// Prints the entries to the terminal with stdout.
func OutputEntries(ctx context.Context, des []details.DetailsEntry) {
	ps := []Printable{}
	for _, de := range des {
		ps = append(ps, de)
	}
	printAll(getRootCmd(ctx).OutOrStdout(), ps)
}

// ------------------------------------------------------------------------------------------
// Tabular
// ------------------------------------------------------------------------------------------

// Table writes the printables in a tabular format.  Takes headers from
// the 0th printable only.
func Table(ctx context.Context, ps []Printable) {
	outputTable(getRootCmd(ctx).OutOrStdout(), ps)
}

// output to stdout the list of printable structs in a table
func outputTable(w io.Writer, ps []Printable) {
	t := table.Table{
		Headers: ps[0].Headers(),
		Rows:    [][]string{},
	}
	for _, p := range ps {
		t.Rows = append(t.Rows, p.Values())
	}
	_ = t.WriteTable(
		w,
		&table.Config{
			ShowIndex:       false,
			Color:           false,
			AlternateColors: false,
		})
}

// ------------------------------------------------------------------------------------------
// JSON
// ------------------------------------------------------------------------------------------

func outputJSON(w io.Writer, p Printable, debug bool) {
	if debug {
		printJSON(w, p)
		return
	}
	printJSON(w, p.MinimumPrintable())
}

func outputJSONArr(w io.Writer, ps []Printable, debug bool) {
	sl := make([]any, 0, len(ps))
	for _, p := range ps {
		if debug {
			sl = append(sl, p)
		} else {
			sl = append(sl, p.MinimumPrintable())
		}
	}
	printJSON(w, sl)
}

// output to stdout the list of printable structs as json.
func printJSON(w io.Writer, a any) {
	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Fprintf(w, "error formatting results to json: %v\n", err)
		return
	}
	fmt.Fprintln(w, string(pretty.Pretty(bs)))
}
