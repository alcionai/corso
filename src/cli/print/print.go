package print

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"github.com/tomlazar/table"
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
	if ctx == nil {
		return &cobra.Command{}
	}

	cmdIface := ctx.Value(rootCmdCtx{})
	cmd, ok := cmdIface.(*cobra.Command)

	if cmd == nil || !ok {
		return &cobra.Command{}
	}

	return cmd
}

// adds the persistent flag --output to the provided command.
func AddOutputFlag(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.BoolVar(&outputAsJSON, "json", false, "output data in JSON format")
	fs.BoolVar(&outputAsJSONDebug, "json-debug", false, "output all internal and debugging data in JSON format")
	cobra.CheckErr(fs.MarkHidden("json-debug"))
}

// JSONFormat returns true if the printer plans to output as json.
func JSONFormat() bool {
	return outputAsJSON || outputAsJSONDebug
}

// StderrWriter returns the stderr writer used in the root
// cmd.  Returns nil if no root command is seeded.
func StderrWriter(ctx context.Context) io.Writer {
	return getRootCmd(ctx).ErrOrStderr()
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

// Out prints the params to cobra's output writer (stdOut by default)
// if s is nil, prints nothing.
func Out(ctx context.Context, s ...any) {
	out(getRootCmd(ctx).OutOrStdout(), s...)
}

// out is the testable core of Out()
func out(w io.Writer, s ...any) {
	if len(s) == 0 {
		return
	}

	fmt.Fprint(w, s...)
	fmt.Fprintf(w, "\n")
}

// Out prints the formatted strings to cobra's output writer (stdOut by default)
// if t is empty, prints nothing.
func Outf(ctx context.Context, t string, s ...any) {
	outf(getRootCmd(ctx).OutOrStdout(), t, s...)
}

// outf is the testable core of Outf()
func outf(w io.Writer, t string, s ...any) {
	if len(t) == 0 {
		return
	}

	fmt.Fprintf(w, t, s...)
	fmt.Fprintf(w, "\n")
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

// Item prints the printable, according to the caller's requested format.
func Item(ctx context.Context, p Printable) {
	printItem(getRootCmd(ctx).OutOrStdout(), p)
}

// print prints the printable items,
// according to the caller's requested format.
func printItem(w io.Writer, p Printable) {
	if outputAsJSON || outputAsJSONDebug {
		outputJSON(w, p, outputAsJSONDebug)
		return
	}

	outputTable(w, []Printable{p})
}

// All prints the slice of printable items,
// according to the caller's requested format.
func All(ctx context.Context, ps ...Printable) {
	printAll(getRootCmd(ctx).OutOrStdout(), ps)
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
