package flags

import (
	"github.com/spf13/cobra"
)

type flagType int

const (
	UnknownType flagType = iota
	StringType
	IntType
	BoolType
)

// CliFlag structures a cli flag declaration for attachment to a command using cobra.
type CliFlag struct {
	Name        string   // Long form name of the flag.  Ex: "flag-name" => --flag-name
	Short       string   // Short form name of the flag.  Ex: "f" => -f
	Description string   // Help menu description of the flag.
	Fallback    any      // Default value, in case the flag is not provided.
	Var         any      // The variable which the flag will populate (expects a pointer).
	VarType     flagType // Specifies the type of flag variable.
	Global      bool     // When true, all child args will inherit usage of the flag.
	Required    bool     // In case the flag must be populated for the command to run.
}

// AddAll flags in the slice to the provided command.
func AddAll(flags []CliFlag, cmd *cobra.Command) {
	for _, f := range flags {
		Add(f, cmd)
	}
}

// Add a single flag to the provided command.
func Add(flag CliFlag, cmd *cobra.Command) {
	flags := cmd.Flags()
	require := cmd.MarkFlagRequired
	if flag.Global {
		flags = cmd.PersistentFlags()
		require = cmd.MarkPersistentFlagRequired
	}

	switch flag.VarType {
	case StringType:
		flags.StringVarP(
			flag.Var.(*string),
			flag.Name, flag.Short,
			orZero(flag.Fallback, StringType).(string),
			flag.Description)
	case IntType:
		flags.IntVarP(
			flag.Var.(*int),
			flag.Name, flag.Short,
			orZero(flag.Fallback, IntType).(int),
			flag.Description)
	case BoolType:
		flags.BoolVarP(
			flag.Var.(*bool),
			flag.Name, flag.Short,
			orZero(flag.Fallback, BoolType).(bool),
			flag.Description)
	default:
		return
	}

	if flag.Required {
		require(flag.Name)
	}
}

// prevents nil pointer errors in case of a nil flag.Fallback.
func orZero(fallback any, t flagType) any {
	if fallback != nil {
		return fallback
	}
	switch t {
	case StringType:
		return ""
	case IntType:
		return 0
	default:
		return false
	}
}
