package repo

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type FilesystemSuite struct {
	tester.Suite
}

func TestFilesystemSuite(t *testing.T) {
	suite.Run(t, &FilesystemSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *FilesystemSuite) TestAddFilesystemCommands() {
	expectUse := fsProviderCommand + " " + fsProviderCmdUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"init filesystem", initCommand, expectUse, filesystemInitCmd().Short, initFilesystemCmd},
		{"connect filesystem", connectCommand, expectUse, filesystemConnectCmd().Short, connectFilesystemCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addFilesystemCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)
		})
	}
}
