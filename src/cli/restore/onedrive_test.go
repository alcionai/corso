package restore

import (
	"bytes"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
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
	expectUse := oneDriveServiceCommand + " " + oneDriveServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"restore onedrive", restoreCommand, expectUse, oneDriveRestoreCmd().Short, restoreOneDriveCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// normally a persisten flag from the root.
			// required to ensure a dry run.
			utils.AddRunModeFlag(cmd, true)

			c := addOneDriveCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			cmd.SetArgs([]string{
				"onedrive",
				"--" + utils.RunModeFN, utils.RunModeFlagTest,
				"--" + utils.BackupFN, testdata.BackupInpt,

				"--" + utils.FileFN, testdata.FlgInpts(testdata.FileNamesInpt),
				"--" + utils.FolderFN, testdata.FlgInpts(testdata.FolderPathsInpt),
				"--" + utils.FileCreatedAfterFN, testdata.FileCreatedAfterInpt,
				"--" + utils.FileCreatedBeforeFN, testdata.FileCreatedBeforeInpt,
				"--" + utils.FileModifiedAfterFN, testdata.FileModifiedAfterInpt,
				"--" + utils.FileModifiedBeforeFN, testdata.FileModifiedBeforeInpt,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeOneDriveOpts(cmd)
			assert.Equal(t, testdata.BackupInpt, utils.BackupID)

			assert.ElementsMatch(t, testdata.FileNamesInpt, opts.FileNames)
			assert.ElementsMatch(t, testdata.FolderPathsInpt, opts.FolderPaths)
			assert.Equal(t, testdata.FileCreatedAfterInpt, opts.FileCreatedAfter)
			assert.Equal(t, testdata.FileCreatedBeforeInpt, opts.FileCreatedBefore)
			assert.Equal(t, testdata.FileModifiedAfterInpt, opts.FileModifiedAfter)
			assert.Equal(t, testdata.FileModifiedBeforeInpt, opts.FileModifiedBefore)
		})
	}
}
