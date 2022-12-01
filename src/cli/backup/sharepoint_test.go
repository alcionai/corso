package backup

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
)

type SharePointSuite struct {
	suite.Suite
}

func TestSharePointSuite(t *testing.T) {
	suite.Run(t, new(SharePointSuite))
}

func (suite *SharePointSuite) TestAddSharePointCommands() {
	expectUse := sharePointServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create sharepoint", createCommand, expectUse + " " + sharePointServiceCommandCreateUseSuffix,
			sharePointCreateCmd().Short, createSharePointCmd,
		},
		{
			"list sharepoint", listCommand, expectUse,
			sharePointListCmd().Short, listSharePointCmd,
		},
		{
			"details sharepoint", detailsCommand, expectUse + " " + sharePointServiceCommandDetailsUseSuffix,
			sharePointDetailsCmd().Short, detailsSharePointCmd,
		},
		{
			"delete sharepoint", deleteCommand, expectUse + " " + sharePointServiceCommandDeleteUseSuffix,
			sharePointDeleteCmd().Short, deleteSharePointCmd,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cmd := &cobra.Command{Use: test.use}

			c := addSharePointCommands(cmd)
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

func (suite *SharePointSuite) TestValidateSharePointBackupCreateFlags() {
	table := []struct {
		name   string
		site   []string
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "no sites",
			expect: assert.Error,
		},
		{
			name:   "sites",
			site:   []string{"fnord"},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, validateSharePointBackupCreateFlags(test.site))
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupDetailsSelectors() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.SharePointOptionDetailLookups {
		suite.T().Run(test.Name, func(t *testing.T) {
			output, err := runDetailsSharePointCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
			)
			assert.NoError(t, err)

			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupDetailsSelectorsBadFormats() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.BadSharePointOptionsFormats {
		suite.T().Run(test.Name, func(t *testing.T) {
			output, err := runDetailsSharePointCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
			)

			assert.Error(t, err)
			assert.Empty(t, output)
		})
	}
}
