package flags_test

import (
	"testing"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/cli/flags"
)

var table = []flags.CliFlag{
	{
		Name:        "not-global-not-required-string",
		Short:       "s",
		Description: "desc",
		Var:         pVarByType(int(flags.StringType)),
		VarType:     flags.StringType,
	},
	{
		Name:        "not-global-not-required-int",
		Short:       "i",
		Description: "desc",
		Var:         pVarByType(int(flags.IntType)),
		VarType:     flags.IntType,
	},
	{
		Name:        "not-global-not-required-bool",
		Short:       "b",
		Description: "desc",
		Var:         pVarByType(int(flags.BoolType)),
		VarType:     flags.BoolType,
	},
	{
		Name:    "global-not-required",
		Global:  true,
		Var:     pVarByType(int(flags.StringType)),
		VarType: flags.StringType,
	},
	{
		Name:     "required-not-global",
		Required: true,
		Var:      pVarByType(int(flags.StringType)),
		VarType:  flags.StringType,
	},
	{
		Name:     "global-and-required",
		Global:   true,
		Required: true,
		Var:      pVarByType(int(flags.StringType)),
		VarType:  flags.StringType,
	},
}

func pVarByType(ft int) any {
	switch ft {
	case int(flags.StringType):
		var s string
		return &s
	case int(flags.IntType):
		var i int
		return &i
	case int(flags.BoolType):
		var b bool
		return &b
	default:
		return nil
	}
}

func TestAdd(t *testing.T) {
	for _, flag := range table {
		t.Run(flag.Name, func(t *testing.T) {
			cmd := &cobra.Command{}
			flags.Add(flag, cmd)
			testFlagInCmd(t, flag, cmd)
		})
	}
}

func TestAddAll(t *testing.T) {
	cmd := &cobra.Command{}
	flags.AddAll(table, cmd)

	for _, flag := range table {
		t.Run(flag.Name, func(t *testing.T) {
			testFlagInCmd(t, flag, cmd)
		})
	}
}

func testFlagInCmd(t *testing.T, f flags.CliFlag, cmd *cobra.Command) {
	result := cmd.Flag(f.Name)
	if result == nil {
		t.Fatalf("expected flag [%s] to have been added to the command", f.Name)
	}

	if f.Global && !cmd.HasPersistentFlags() {
		t.Errorf("expected flag [%s] to be marked persistent", f.Name)
	}

	if f.Required {
		var isRequired bool
		for _, annVal := range result.Annotations[cobra.BashCompOneRequiredFlag] {
			if annVal == "true" {
				isRequired = true
				break
			}
		}
		if !isRequired {
			t.Errorf("expected flag [%s] to be marked required", f.Name)
		}
	}
}
