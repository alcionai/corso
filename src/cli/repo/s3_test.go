package repo

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type S3Suite struct {
	tester.Suite
}

func TestS3Suite(t *testing.T) {
	suite.Run(t, &S3Suite{Suite: tester.NewUnitSuite(t)})
}

func (suite *S3Suite) TestAddS3Commands() {
	expectUse := s3ProviderCommand + " " + s3ProviderCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"init s3", initCommand, expectUse, s3InitCmd().Short, initS3Cmd},
		{"connect s3", connectCommand, expectUse, s3ConnectCmd().Short, connectS3Cmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addS3Commands(cmd)
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
