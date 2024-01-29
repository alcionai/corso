package print

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"github.com/tomlazar/table"

	"github.com/alcionai/corso/src/internal/common/color"
	"github.com/alcionai/corso/src/internal/observe"
)

var (
	outputAsJSON      bool
	outputAsJSONDebug bool
	outputVerbose     bool
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
	fs.BoolVar(&outputVerbose, "verbose", false, "don't hide additional information")
}

// DisplayJSONFormat returns true if the printer plans to output as json.
func DisplayJSONFormat() bool {
	return outputAsJSON || outputAsJSONDebug
}

// DisplayVerbose returns true if verbose output is enabled
func DisplayVerbose() bool {
	return outputVerbose
}

// StderrWriter returns the stderr writer used in the root
// cmd.  Returns nil if no root command is seeded.
func StderrWriter(ctx context.Context) io.Writer {
	return getRootCmd(ctx).ErrOrStderr()
}

// ---------------------------------------------------------------------------------------------------------
// Exported interface
// ---------------------------------------------------------------------------------------------------------

// Only tells the CLI to only display this error, preventing the usage
// (ie, help) menu from displaying as well.
func Only(ctx context.Context, e error) error {
	getRootCmd(ctx).SilenceUsage = true
	return e
}

// Err prints the params to cobra's error writer (stdErr by default)
// if s is nil, prints nothing.
func Err(ctx context.Context, s ...any) {
	cw := color.NewColorableWriter(color.Red, getRootCmd(ctx).ErrOrStderr())

	s = append([]any{"Error:"}, s...)

	out(ctx, cw, s...)
}

// Errf prints the params to cobra's error writer (stdErr by default)
// if s is nil, prints nothing.
// You should ideally be using SimpleError or OperationError.
func Errf(ctx context.Context, tmpl string, s ...any) {
	cw := color.NewColorableWriter(color.Red, getRootCmd(ctx).ErrOrStderr())
	tmpl = "Error: " + tmpl
	outf(ctx, cw, tmpl, s...)
}

// Out prints the params to cobra's output writer (stdOut by default)
// if s is nil, prints nothing.
func Out(ctx context.Context, s ...any) {
	out(ctx, getRootCmd(ctx).OutOrStdout(), s...)
}

// Out prints the formatted strings to cobra's output writer (stdOut by default)
// if t is empty, prints nothing.
func Outf(ctx context.Context, t string, s ...any) {
	outf(ctx, getRootCmd(ctx).OutOrStdout(), t, s...)
}

// Info prints the params to cobra's error writer (stdErr by default)
// if s is nil, prints nothing.
func Info(ctx context.Context, s ...any) {
	out(ctx, getRootCmd(ctx).ErrOrStderr(), s...)
}

// Info prints the formatted strings to cobra's error writer (stdErr by default)
// if t is empty, prints nothing.
func Infof(ctx context.Context, t string, s ...any) {
	outf(ctx, getRootCmd(ctx).ErrOrStderr(), t, s...)
}

// Pretty prettifies and prints the value.
func Pretty(ctx context.Context, a any) {
	if a == nil {
		Err(ctx, "<nil>")
		return
	}

	printPrettyJSON(ctx, getRootCmd(ctx).ErrOrStderr(), a)
}

// PrettyJSON prettifies and prints the value.
func PrettyJSON(ctx context.Context, p minimumPrintabler) {
	if p == nil {
		Err(ctx, "<nil>")
		return
	}

	outputJSON(ctx, getRootCmd(ctx).ErrOrStderr(), p, outputAsJSONDebug)
}

// out is the testable core of exported print funcs
func out(ctx context.Context, w io.Writer, s ...any) {
	if len(s) == 0 {
		return
	}

	// observe bars needs to be flushed before printing
	observe.Flush(ctx)

	fmt.Fprint(w, s...)
	fmt.Fprintf(w, "\n")
}

// outf is the testable core of exported print funcs
func outf(ctx context.Context, w io.Writer, t string, s ...any) {
	if len(t) == 0 {
		return
	}

	// observe bars needs to be flushed before printing
	observe.Flush(ctx)

	fmt.Fprintf(w, t, s...)
	fmt.Fprintf(w, "\n")
}

// ---------------------------------------------------------------------------------------------------------
// Output control for backup list/details
// ---------------------------------------------------------------------------------------------------------

type Printable interface {
	minimumPrintabler
	// should list the property names of the values surfaced in Values()
	Headers(skipID bool) []string
	// list of values for tabular or csv formatting
	// if the backing data is nil or otherwise missing,
	// values should provide an empty string as opposed to skipping entries
	Values(skipID bool) []string
}

type minimumPrintabler interface {
	// reduces the struct to a minimized format for easier human consumption
	MinimumPrintable() any
}

// Item prints the printable, according to the caller's requested format.
func Item(ctx context.Context, p Printable) {
	printItem(ctx, getRootCmd(ctx).OutOrStdout(), p)
}

// print prints the printable items,
// according to the caller's requested format.
func printItem(ctx context.Context, w io.Writer, p Printable) {
	if outputAsJSON || outputAsJSONDebug {
		outputJSON(ctx, w, p, outputAsJSONDebug)
		return
	}

	outputTable(ctx, w, []Printable{p})
}

// ItemProperties prints the printable either as in a single line or a json
// The difference between this and Item is that this one does not print the ID
func ItemProperties(ctx context.Context, p Printable) {
	printItemProperties(ctx, getRootCmd(ctx).OutOrStdout(), p)
}

// print prints the printable items,
// according to the caller's requested format.
func printItemProperties(ctx context.Context, w io.Writer, p Printable) {
	if outputAsJSON || outputAsJSONDebug {
		outputJSON(ctx, w, p, outputAsJSONDebug)
		return
	}

	outputOneLine(ctx, w, []Printable{p})
}

// All prints the slice of printable items,
// according to the caller's requested format.
func All(ctx context.Context, ps ...Printable) {
	printAll(ctx, getRootCmd(ctx).OutOrStdout(), ps)
}

// printAll prints the slice of printable items,
// according to the caller's requested format.
func printAll(ctx context.Context, w io.Writer, ps []Printable) {
	if len(ps) == 0 {
		return
	}

	if outputAsJSON || outputAsJSONDebug {
		outputJSONArr(ctx, w, ps, outputAsJSONDebug)
		return
	}

	outputTable(ctx, w, ps)
}

// ------------------------------------------------------------------------------------------
// Tabular
// ------------------------------------------------------------------------------------------

// Table writes the printables in a tabular format.  Takes headers from
// the 0th printable only.
func Table(ctx context.Context, ps []Printable) {
	outputTable(ctx, getRootCmd(ctx).OutOrStdout(), ps)
}

// output to stdout the list of printable structs in a table
func outputTable(ctx context.Context, w io.Writer, ps []Printable) {
	t := table.Table{
		Headers: ps[0].Headers(false),
		Rows:    [][]string{},
	}

	for _, p := range ps {
		t.Rows = append(t.Rows, p.Values(false))
	}

	// observe bars needs to be flushed before printing
	observe.Flush(ctx)

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

func outputJSON(ctx context.Context, w io.Writer, p minimumPrintabler, debug bool) {
	if debug {
		printJSON(ctx, w, p)
		return
	}

	if debug {
		printJSON(ctx, w, p)
	} else {
		printJSON(ctx, w, p.MinimumPrintable())
	}
}

func outputJSONArr(ctx context.Context, w io.Writer, ps []Printable, debug bool) {
	sl := make([]any, 0, len(ps))

	for _, p := range ps {
		if debug {
			sl = append(sl, p)
		} else {
			sl = append(sl, p.MinimumPrintable())
		}
	}

	printJSON(ctx, w, sl)
}

// output to stdout the list of printable structs as json.
func printJSON(ctx context.Context, w io.Writer, a any) {
	// observe bars needs to be flushed before printing
	observe.Flush(ctx)

	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Fprintf(w, "error formatting results to json: %v\n", err)
		return
	}

	fmt.Fprintln(w, string(pretty.Pretty(bs)))
}

// output to stdout the list of printable structs as prettified json.
func printPrettyJSON(ctx context.Context, w io.Writer, a any) {
	// observe bars needs to be flushed before printing
	observe.Flush(ctx)

	bs, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "error formatting results to json: %v\n", err)
		return
	}

	fmt.Fprintln(w, string(pretty.Pretty(bs)))
}

// -------------------------------------------------------------------------------------------
// One line
// -------------------------------------------------------------------------------------------

// Output in the following format:
// Bytes Uploaded: 401 kB | Items Uploaded: 59 | Items Skipped: 0 | Errors: 0
func outputOneLine(ctx context.Context, w io.Writer, ps []Printable) {
	// observe bars needs to be flushed before printing
	observe.Flush(ctx)

	headers := ps[0].Headers(true)
	rows := [][]string{}

	for _, p := range ps {
		rows = append(rows, p.Values(true))
	}

	printables := []string{}

	for _, row := range rows {
		for i, col := range row {
			printables = append(printables, fmt.Sprintf("%s: %s", headers[i], col))
		}
	}

	fmt.Fprintln(w, strings.Join(printables, " | "))
}
