package backup

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/tester"
)

type TeamsUnitSuite struct {
	tester.Suite
}

func TestTeamsUnitSuite(t *testing.T) {
	suite.Run(t, &TeamsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *TeamsUnitSuite) TestAddTeamsCommands() {
	expectUse := teamsServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		flags       []string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create teams",
			createCommand,
			expectUse + " " + teamsServiceCommandCreateUseSuffix,
			teamsCreateCmd().Short,
			[]string{
				flags.CategoryDataFN,
				flags.FailFastFN,
				flags.FetchParallelismFN,
				flags.SkipReduceFN,
				flags.NoStatsFN,
			},
			createTeamsCmd,
		},
		{
			"list teams",
			listCommand,
			expectUse,
			teamsListCmd().Short,
			[]string{
				flags.BackupFN,
				flags.FailedItemsFN,
				flags.SkippedItemsFN,
				flags.RecoveredErrorsFN,
			},
			listTeamsCmd,
		},
		{
			"details teams",
			detailsCommand,
			expectUse + " " + teamsServiceCommandDetailsUseSuffix,
			teamsDetailsCmd().Short,
			[]string{
				flags.BackupFN,
			},
			detailsTeamsCmd,
		},
		{
			"delete teams",
			deleteCommand,
			expectUse + " " + teamsServiceCommandDeleteUseSuffix,
			teamsDeleteCmd().Short,
			[]string{flags.BackupFN},
			deleteTeamsCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addTeamsCommands(cmd)
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
