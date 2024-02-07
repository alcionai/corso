package testdata

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/internal/tester"
)

// StubRootCmd builds a stub cobra command to be used as
// the root command for integration testing on the CLI
func StubRootCmd(args ...string) *cobra.Command {
	id := uuid.NewString()
	now := time.Now().UTC().Format(time.RFC3339Nano)
	cmdArg := "testing-canario"
	c := &cobra.Command{
		Use:   cmdArg,
		Short: id,
		Long:  id + " - " + now,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "test command args: %+v", args)
			return nil
		},
	}
	c.SetArgs(args)

	return c
}

type UseCobraCommandFn func(*cobra.Command)

func SetUpCmdHasFlags(
	t *testing.T,
	parentCmd *cobra.Command,
	addChildCommand func(*cobra.Command) *cobra.Command,
	addFlags []UseCobraCommandFn,
	setArgs UseCobraCommandFn,
) *cobra.Command {
	parentCmd.PersistentPreRun = func(c *cobra.Command, args []string) {
		t.Log("testing args:")

		for _, arg := range args {
			t.Log(arg)
		}
	}

	// persistent flags not added by addCommands
	flags.AddRunModeFlag(parentCmd, true)

	cmd := addChildCommand(parentCmd)
	require.NotNil(t, cmd)

	cul := cmd.UseLine()
	require.Truef(
		t,
		strings.HasPrefix(cul, parentCmd.Use+" "+cmd.Use),
		"child command has expected usage format 'parent child', got %q",
		cul)

	for _, af := range addFlags {
		af(cmd)
	}

	setArgs(parentCmd)

	parentCmd.SetOut(new(bytes.Buffer)) // drop output
	parentCmd.SetErr(new(bytes.Buffer)) // drop output

	err := parentCmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	return cmd
}

type CobraRunEFn func(cmd *cobra.Command, args []string) error

func CheckCmdChild(
	t *testing.T,
	cmd *cobra.Command,
	expectChildCount int,
	expectUse string,
	expectShort string,
	expectRunE CobraRunEFn,
) {
	var (
		cmds  = cmd.Commands()
		child *cobra.Command
	)

	for _, cc := range cmds {
		if cc.Use == expectUse {
			child = cc
			break
		}
	}

	require.Len(
		t,
		cmds,
		expectChildCount,
		"parent command should have the correct child command count")

	require.NotNil(t, child, "should have found expected child command")

	assert.Equal(t, expectShort, child.Short)
	tester.AreSameFunc(t, expectRunE, child.RunE)
}
