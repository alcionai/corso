package backup

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GroupsUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsUnitSuite) TestAddGroupsCommands() {
	expectUse := groupsServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		flags       []string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create groups",
			createCommand,
			expectUse + " " + groupsServiceCommandCreateUseSuffix,
			groupsCreateCmd().Short,
			[]string{
				flags.CategoryDataFN,
				flags.FailFastFN,
				flags.FetchParallelismFN,
				flags.SkipReduceFN,
				flags.NoStatsFN,
			},
			createGroupsCmd,
		},
		{
			"list groups",
			listCommand,
			expectUse,
			groupsListCmd().Short,
			[]string{
				flags.BackupFN,
				flags.FailedItemsFN,
				flags.SkippedItemsFN,
				flags.RecoveredErrorsFN,
			},
			listGroupsCmd,
		},
		{
			"details groups",
			detailsCommand,
			expectUse + " " + groupsServiceCommandDetailsUseSuffix,
			groupsDetailsCmd().Short,
			[]string{
				flags.BackupFN,
			},
			detailsGroupsCmd,
		},
		{
			"delete groups",
			deleteCommand,
			expectUse + " " + groupsServiceCommandDeleteUseSuffix,
			groupsDeleteCmd().Short,
			[]string{flags.BackupFN},
			deleteGroupsCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addGroupsCommands(cmd)
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

func (suite *GroupsUnitSuite) TestValidateGroupsBackupCreateFlags() {
	table := []struct {
		name   string
		cats   []string
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "none",
			cats:   []string{},
			expect: assert.NoError,
		},
		{
			name:   "libraries",
			cats:   []string{dataLibraries},
			expect: assert.NoError,
		},
		{
			name:   "messages",
			cats:   []string{dataMessages},
			expect: assert.NoError,
		},
		{
			name:   "all allowed",
			cats:   []string{dataLibraries, dataMessages},
			expect: assert.NoError,
		},
		{
			name:   "bad inputs",
			cats:   []string{"foo"},
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := validateGroupsBackupCreateFlags([]string{"*"}, test.cats)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *GroupsUnitSuite) TestAddGroupsCategories() {
	table := []struct {
		name           string
		cats           []string
		expectScopeLen int
	}{
		{
			name:           "none",
			cats:           []string{},
			expectScopeLen: 2,
		},
		{
			name:           "libraries",
			cats:           []string{dataLibraries},
			expectScopeLen: 1,
		},
		{
			name:           "messages",
			cats:           []string{dataMessages},
			expectScopeLen: 1,
		},
		{
			name:           "all allowed",
			cats:           []string{dataLibraries, dataMessages},
			expectScopeLen: 2,
		},
		{
			name:           "bad inputs",
			cats:           []string{"foo"},
			expectScopeLen: 0,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			sel := addGroupsCategories(selectors.NewGroupsBackup(selectors.Any()), test.cats)
			scopes := sel.Scopes()
			assert.Len(suite.T(), scopes, test.expectScopeLen)
		})
	}
}
