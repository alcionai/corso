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

type ExchangeUnitSuite struct {
	tester.Suite
}

func TestExchangeUnitSuite(t *testing.T) {
	suite.Run(t, &ExchangeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExchangeUnitSuite) TestAddExchangeCommands() {
	expectUse := exchangeServiceCommand + " " + exchangeServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"restore exchange", restoreCommand, expectUse, exchangeRestoreCmd().Short, restoreExchangeCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// normally a persisten flag from the root.
			// required to ensure a dry run.
			utils.AddRunModeFlag(cmd, true)

			c := addExchangeCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			// Test arg parsing for few args
			cmd.SetArgs([]string{
				"exchange",
				"--" + utils.RunModeFN, utils.RunModeFlagTest,
				"--" + utils.BackupFN, testdata.BackupInpt,

				"--" + utils.ContactFN, testdata.FlgInpts(testdata.ContactInpt),
				"--" + utils.ContactFolderFN, testdata.FlgInpts(testdata.ContactFldInpt),
				"--" + utils.ContactNameFN, testdata.ContactNameInpt,

				"--" + utils.EmailFN, testdata.FlgInpts(testdata.EmailInpt),
				"--" + utils.EmailFolderFN, testdata.FlgInpts(testdata.EmailFldInpt),
				"--" + utils.EmailReceivedAfterFN, testdata.EmailReceivedAfterInpt,
				"--" + utils.EmailReceivedBeforeFN, testdata.EmailReceivedBeforeInpt,
				"--" + utils.EmailSenderFN, testdata.EmailSenderInpt,
				"--" + utils.EmailSubjectFN, testdata.EmailSubjectInpt,

				"--" + utils.EventFN, testdata.FlgInpts(testdata.EventInpt),
				"--" + utils.EventCalendarFN, testdata.FlgInpts(testdata.EventCalInpt),
				"--" + utils.EventOrganizerFN, testdata.EventOrganizerInpt,
				"--" + utils.EventRecursFN, testdata.EventRecursInpt,
				"--" + utils.EventStartsAfterFN, testdata.EventStartsAfterInpt,
				"--" + utils.EventStartsBeforeFN, testdata.EventStartsBeforeInpt,
				"--" + utils.EventSubjectFN, testdata.EventSubjectInpt,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeExchangeOpts(cmd)
			assert.Equal(t, testdata.BackupInpt, utils.BackupID)

			assert.ElementsMatch(t, testdata.ContactInpt, opts.Contact)
			assert.ElementsMatch(t, testdata.ContactFldInpt, opts.ContactFolder)
			assert.Equal(t, testdata.ContactNameInpt, opts.ContactName)

			assert.ElementsMatch(t, testdata.EmailInpt, opts.Email)
			assert.ElementsMatch(t, testdata.EmailFldInpt, opts.EmailFolder)
			assert.Equal(t, testdata.EmailReceivedAfterInpt, opts.EmailReceivedAfter)
			assert.Equal(t, testdata.EmailReceivedBeforeInpt, opts.EmailReceivedBefore)
			assert.Equal(t, testdata.EmailSenderInpt, opts.EmailSender)
			assert.Equal(t, testdata.EmailSubjectInpt, opts.EmailSubject)

			assert.ElementsMatch(t, testdata.EventInpt, opts.Event)
			assert.ElementsMatch(t, testdata.EventCalInpt, opts.EventCalendar)
			assert.Equal(t, testdata.EventOrganizerInpt, opts.EventOrganizer)
			assert.Equal(t, testdata.EventRecursInpt, opts.EventRecurs)
			assert.Equal(t, testdata.EventStartsAfterInpt, opts.EventStartsAfter)
			assert.Equal(t, testdata.EventStartsBeforeInpt, opts.EventStartsBefore)
			assert.Equal(t, testdata.EventSubjectInpt, opts.EventSubject)
		})
	}
}
