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
	suite.Suite
}

func TestOneDriveSuite(t *testing.T) {
	suite.Run(t, new(OneDriveSuite))
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
		suite.T().Run(test.name, func(t *testing.T) {
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
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, validateOneDriveBackupCreateFlags(test.user))
		})
	}
}

func (suite *OneDriveSuite) TestOneDriveBackupDetailsSelectors() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.OneDriveOptionDetailLookups {
		suite.T().Run(test.Name, func(t *testing.T) {
			output, errs := runDetailsOneDriveCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts)
			assert.NoError(t, errs.Err())
			assert.Empty(t, errs.Errs())
			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *OneDriveSuite) TestOneDriveBackupDetailsSelectorsBadFormats() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.BadOneDriveOptionsFormats {
		suite.T().Run(test.Name, func(t *testing.T) {
			output, errs := runDetailsOneDriveCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts)
			assert.Error(t, errs.Err())
			assert.Empty(t, errs.Errs())
			assert.Empty(t, output)
		})
	}
}
