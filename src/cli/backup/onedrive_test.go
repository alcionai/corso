package backup

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
)

type OneDriveUnitSuite struct {
	tester.Suite
}

func TestOneDriveUnitSuite(t *testing.T) {
	suite.Run(t, &OneDriveUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDriveUnitSuite) TestAddOneDriveCommands() {
	expectUse := oneDriveServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		flags       []string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create onedrive",
			createCommand,
			expectUse + " " + oneDriveServiceCommandCreateUseSuffix,
			oneDriveCreateCmd().Short,
			[]string{"user", "disable-incrementals"},
			createOneDriveCmd,
		},
		{
			"list onedrive",
			listCommand,
			expectUse,
			oneDriveListCmd().Short,
			[]string{"backup", "failed-items", "skipped-items", "recovered-errors"},
			listOneDriveCmd,
		},
		{
			"details onedrive",
			detailsCommand,
			expectUse + " " + oneDriveServiceCommandDetailsUseSuffix,
			oneDriveDetailsCmd().Short,
			[]string{
				"backup",
				"folder",
				"file",
				"file-created-after",
				"file-created-before",
				"file-modified-after",
				"file-modified-before",
			},
			detailsOneDriveCmd,
		},
		{
			"delete onedrive",
			deleteCommand,
			expectUse + " " + oneDriveServiceCommandDeleteUseSuffix,
			oneDriveDeleteCmd().Short,
			[]string{"backup"},
			deleteOneDriveCmd,
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

			for _, f := range test.flags {
				assert.NotNil(t, c.Flag(f), f+" flag")
			}
		})
	}
}

func (suite *OneDriveUnitSuite) TestValidateOneDriveBackupCreateFlags() {
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
			err := validateOneDriveBackupCreateFlags(test.user)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *OneDriveUnitSuite) TestOneDriveBackupDetailsSelectors() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.OneDriveOptionDetailLookups {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsOneDriveCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
				false)
			assert.NoError(t, err, clues.ToCore(err))
			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *OneDriveUnitSuite) TestOneDriveBackupDetailsSelectorsBadFormats() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.BadOneDriveOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsOneDriveCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
				false)
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, output)
		})
	}
}
