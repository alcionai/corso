package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/alcionai/corso/cli"
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

const fmTemplate = `---
title: "%s"
hide_title: true
---
`

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
	identity := func(s string) string { return s }
	filePrepender := func(filename string) string {
		name := filepath.Base(filename)
		base := strings.TrimSuffix(name, filepath.Ext(name))

		return fmt.Sprintf(fmTemplate, strings.Replace(base, "_", " ", -1))
	}

	if err := makeDir(cliMarkdownDir); err != nil {
		fatal(errors.Wrap(err, "preparing directory for markdown generation"))
	}

	corsoCmd := cli.CorsoCommand()
	corsoCmd.DisableAutoGenTag = true

	err := doc.GenMarkdownTreeCustom(corsoCmd, cliMarkdownDir, filePrepender, identity)
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
