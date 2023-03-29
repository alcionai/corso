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

type SharePointUnitSuite struct {
	tester.Suite
}

func TestSharePointUnitSuite(t *testing.T) {
	suite.Run(t, &SharePointUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SharePointUnitSuite) TestAddSharePointCommands() {
	expectUse := sharePointServiceCommand + " " + sharePointServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"restore onedrive", restoreCommand, expectUse, sharePointRestoreCmd().Short, restoreSharePointCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// normally a persisten flag from the root.
			// required to ensure a dry run.
			utils.AddRunModeFlag(cmd, true)

			c := addSharePointCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			cmd.SetArgs([]string{
				"sharepoint",
				"--" + utils.RunModeFN, utils.RunModeFlagTest,
				"--" + utils.BackupFN, testdata.BackupInpt,

				"--" + utils.LibraryFN, testdata.LibraryInpt,
				"--" + utils.FileFN, testdata.FlgInpts(testdata.FileNamesInpt),
				"--" + utils.FolderFN, testdata.FlgInpts(testdata.FolderPathsInpt),
				"--" + utils.FileCreatedAfterFN, testdata.FileCreatedAfterInpt,
				"--" + utils.FileCreatedBeforeFN, testdata.FileCreatedBeforeInpt,
				"--" + utils.FileModifiedAfterFN, testdata.FileModifiedAfterInpt,
				"--" + utils.FileModifiedBeforeFN, testdata.FileModifiedBeforeInpt,

				"--" + utils.ListItemFN, testdata.FlgInpts(testdata.ListItemInpt),
				"--" + utils.ListFolderFN, testdata.FlgInpts(testdata.ListFolderInpt),

				"--" + utils.PageFN, testdata.FlgInpts(testdata.PageInpt),
				"--" + utils.PageFolderFN, testdata.FlgInpts(testdata.PageFolderInpt),
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeSharePointOpts(cmd)
			assert.Equal(t, testdata.BackupInpt, utils.BackupID)

			assert.Equal(t, testdata.LibraryInpt, opts.Library)
			assert.ElementsMatch(t, testdata.FileNamesInpt, opts.FileName)
			assert.ElementsMatch(t, testdata.FolderPathsInpt, opts.FolderPath)
			assert.Equal(t, testdata.FileCreatedAfterInpt, opts.FileCreatedAfter)
			assert.Equal(t, testdata.FileCreatedBeforeInpt, opts.FileCreatedBefore)
			assert.Equal(t, testdata.FileModifiedAfterInpt, opts.FileModifiedAfter)
			assert.Equal(t, testdata.FileModifiedBeforeInpt, opts.FileModifiedBefore)

			assert.ElementsMatch(t, testdata.ListItemInpt, opts.ListItem)
			assert.ElementsMatch(t, testdata.ListFolderInpt, opts.ListFolder)

			assert.ElementsMatch(t, testdata.PageInpt, opts.Page)
			assert.ElementsMatch(t, testdata.PageFolderInpt, opts.PageFolder)
		})
	}
}
