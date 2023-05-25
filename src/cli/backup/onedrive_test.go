package backup

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/options"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	dtd "github.com/alcionai/corso/src/pkg/backup/details/testdata"
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
			[]string{
				utils.UserFN,
				options.DisableIncrementalsFN,
				options.FailFastFN,
			},
			createOneDriveCmd,
		},
		{
			"list onedrive",
			listCommand,
			expectUse,
			oneDriveListCmd().Short,
			[]string{
				utils.BackupFN,
				failedItemsFN,
				skippedItemsFN,
				recoveredErrorsFN,
			},
			listOneDriveCmd,
		},
		{
			"details onedrive",
			detailsCommand,
			expectUse + " " + oneDriveServiceCommandDetailsUseSuffix,
			oneDriveDetailsCmd().Short,
			[]string{
				utils.BackupFN,
				utils.FolderFN,
				utils.FileFN,
				utils.FileCreatedAfterFN,
				utils.FileCreatedBeforeFN,
				utils.FileModifiedAfterFN,
				utils.FileModifiedBeforeFN,
			},
			detailsOneDriveCmd,
		},
		{
			"delete onedrive",
			deleteCommand,
			expectUse + " " + oneDriveServiceCommandDeleteUseSuffix,
			oneDriveDeleteCmd().Short,
			[]string{utils.BackupFN},
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
	for v := 0; v <= version.Backup; v++ {
		suite.Run(fmt.Sprintf("version%d", v), func() {
			for _, test := range testdata.OneDriveOptionDetailLookups {
				suite.Run(test.Name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					bg := testdata.VersionedBackupGetter{
						Details: dtd.GetDetailsSetForVersion(t, v),
					}

					output, err := runDetailsOneDriveCmd(
						ctx,
						bg,
						"backup-ID",
						test.Opts(t, v),
						false)
					assert.NoError(t, err, clues.ToCore(err))
					assert.ElementsMatch(t, test.Expected(t, v), output.Entries)
				})
			}
		})
	}
}

func (suite *OneDriveUnitSuite) TestOneDriveBackupDetailsSelectorsBadFormats() {
	for _, test := range testdata.BadOneDriveOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			output, err := runDetailsOneDriveCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts(t, version.Backup),
				false)
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, output)
		})
	}
}
