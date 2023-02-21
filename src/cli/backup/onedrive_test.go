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

type OneDriveSuite struct {
	tester.Suite
}

func TestOneDriveSuite(t *testing.T) {
	suite.Run(t, &OneDriveSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDriveSuite) TestAddOneDriveCommands() {
	expectUse := oneDriveServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create onedrive", createCommand, expectUse + " " + oneDriveServiceCommandCreateUseSuffix,
			oneDriveCreateCmd().Short, createOneDriveCmd,
		},
		{
			"list onedrive", listCommand, expectUse,
			oneDriveListCmd().Short, listOneDriveCmd,
		},
		{
			"details onedrive", detailsCommand, expectUse + " " + oneDriveServiceCommandDetailsUseSuffix,
			oneDriveDetailsCmd().Short, detailsOneDriveCmd,
		},
		{
			"delete onedrive", deleteCommand, expectUse + " " + oneDriveServiceCommandDeleteUseSuffix,
			oneDriveDeleteCmd().Short, deleteOneDriveCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addOneDriveCommands(cmd)
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

func (suite *OneDriveSuite) TestValidateOneDriveBackupCreateFlags() {
	table := []struct {
		name   string
		user   []string
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "no users",
			expect: assert.Error,
		},
		{
			name:   "users",
			user:   []string{"fnord"},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), validateOneDriveBackupCreateFlags(test.user))
		})
	}
}

func (suite *OneDriveSuite) TestOneDriveBackupDetailsSelectors() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.OneDriveOptionDetailLookups {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsOneDriveCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts)
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *OneDriveSuite) TestOneDriveBackupDetailsSelectorsBadFormats() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.BadOneDriveOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsOneDriveCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts)
			assert.Error(t, err)
			assert.Empty(t, output)
		})
	}
}
