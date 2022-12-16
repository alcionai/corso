package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli"
)

// generate markdown files in the given.
// callers of this func can then migrate the files
// to where they need.
var cliMarkdownDir string

// The root-level command.
// `corso <command> [<subcommand>] [<service>] [<flag>...]`
var cmd = &cobra.Command{
	Use:               "generate",
	Short:             "Autogenerate Corso documentation.",
	DisableAutoGenTag: true,
	Run:               genDocs,
}

func main() {
	cmd.
		PersistentFlags().
		StringVar(
			&cliMarkdownDir,
			"cli-folder",
			"./cmd/mdgen/cli_markdown",
			"relative path to the folder where cli docs will be generated (default: ./cmd/mdgen/cli_markdown)")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func genDocs(cmd *cobra.Command, args []string) {
	if err := makeDir(cliMarkdownDir); err != nil {
		fatal(errors.Wrap(err, "preparing directory for markdown generation"))
	}

	corsoCmd := cli.CorsoCommand()

	err := genMarkdownCorso(corsoCmd, cliMarkdownDir)
	if err != nil {
		fatal(errors.Wrap(err, "generating the Corso CLI markdown"))
	}
}

func makeDir(dir string) error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "finding current working directory")
	}

	if !strings.HasSuffix(wd, "/src") {
		return errors.New("must be called from /corso/src")
	}

	_, err = os.Stat(dir)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return errors.Wrap(err, "unable to discover directory")
	}

	if errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return errors.Wrap(err, "generating directory to hold markdown")
		}
	}

	return nil
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "ERR: %v\n", err)
	os.Exit(1)
}

// Adapted from https://github.com/spf13/cobra/blob/main/doc/md_docs.go for Corso specific formatting
func genMarkdownCorso(cmd *cobra.Command, dir string) error {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}

		if err := genMarkdownCorso(c, dir); err != nil {
			return err
		}
	}

	// Skip docs for non-leaf commands
	if !cmd.Runnable() || cmd.HasSubCommands() {
		return nil
	}

	basename := strings.ReplaceAll(cmd.CommandPath(), " ", "-") + ".md"
	filename := filepath.Join(dir, basename)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return genMarkdownCustomCorso(cmd, f)
}

func genMarkdownCustomCorso(cmd *cobra.Command, w io.Writer) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	buf := new(bytes.Buffer)
	name := cmd.CommandPath()

	// frontMatter section
	buf.WriteString("---\n")
	buf.WriteString(fmt.Sprintf("title: %s\n", name))
	buf.WriteString("---\n\n")

	// actual markdown
	if len(cmd.Long) > 0 {
		buf.WriteString(cmd.Long + "\n")
	} else {
		buf.WriteString(cmd.Short + "\n")
	}

	if cmd.Runnable() {
		buf.WriteString("\n")
		buf.WriteString(fmt.Sprintf("```bash\n%s\n```\n", cmd.UseLine()))
	}

	if cmd.HasExample() {
		buf.WriteString("\n")
		buf.WriteString("## Examples\n\n")
		buf.WriteString(fmt.Sprintf("```bash\n%s\n```\n", cmd.Example))
	}

	flags := cmd.NonInheritedFlags()
	if flags.HasAvailableFlags() {
		buf.WriteString("\n")
		buf.WriteString("## Flags\n\n")
		printFlags(buf, flags)
	}

	parentFlags := cmd.InheritedFlags()
	if parentFlags.HasAvailableFlags() {
		buf.WriteString("\n")
		buf.WriteString("## Global and inherited flags\n\n")
		printFlags(buf, parentFlags)
	}

	_, err := buf.WriteTo(w)

	return err
}

func printFlags(buf *bytes.Buffer, flags *pflag.FlagSet) {
	if !flags.HasAvailableFlags() {
		return
	}

	buf.WriteString("|Flag|Short|Default|Help|\n")
	buf.WriteString("|:----|:-----|:-------|:----|\n")

	flags.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		buf.WriteString("|")
		buf.WriteString(fmt.Sprintf("`--%s`", flag.Name))
		buf.WriteString("|")

		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			buf.WriteString(fmt.Sprintf("`-%s`", flag.Shorthand))
		}

		buf.WriteString("|")

		if flag.DefValue != "" {
			defValue := flag.DefValue
			if defValue == "[]" {
				defValue = ""
			}

			buf.WriteString(fmt.Sprintf("`%s`", defValue))
		}

		buf.WriteString("|")
		buf.WriteString(strings.ReplaceAll(flag.Usage, "(required)", "<div class='required'>Required</div>"))
		buf.WriteString("|\n")
	})
}
