package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra/doc"

	"github.com/alcionai/corso/cli"
)

const (
	// generate markdown files in the current directory.
	// callers of this func can then migrate the files
	// to where they need.
	cliMarkdownDir = "./cmd/mdgen/cli_markdown"
)

func main() {
	if err := makeDir(cliMarkdownDir); err != nil {
		fatal(errors.Wrap(err, "preparing directory for markdown generation"))
	}
	cmd := cli.CorsoCommand()
	err := doc.GenMarkdownTree(cmd, cliMarkdownDir)
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
