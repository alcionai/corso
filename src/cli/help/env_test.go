package help

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/tester"
)

type EnvSuite struct {
	suite.Suite
}

func TestEnvSuite(t *testing.T) {
	suite.Run(t, new(EnvSuite))
}

func (suite *EnvSuite) TestAddEnvCommands() {
	t := suite.T()
	cmd := &cobra.Command{Use: "root"}

	AddCommands(cmd)

	c := envCmd()
	require.NotNil(t, c)

	cmds := cmd.Commands()
	require.Len(t, cmds, 1)

	assert.Equal(t, "env", c.Use)
	assert.Equal(t, envCmd().Short, c.Short)
	tester.AreSameFunc(t, handleEnvCmd, c.RunE)
}
