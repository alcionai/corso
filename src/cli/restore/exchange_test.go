package restore

import (
	"bytes"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	"github.com/alcionai/corso/src/cli/utils"
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

			// global flags not added by addCommands
			flags.AddRunModeFlag(cmd, true)
			flags.AddAllProviderFlags(cmd)
			flags.AddAllStorageFlags(cmd)

			c := addExchangeCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			flagsTD.WithFlags(
				cmd,
				exchangeServiceCommand,
				[]string{
					"--" + flags.RunModeFN, flags.RunModeFlagTest,
					"--" + flags.BackupFN, flagsTD.BackupInput,

					"--" + flags.ContactFN, flagsTD.FlgInputs(flagsTD.ContactInput),
					"--" + flags.ContactFolderFN, flagsTD.FlgInputs(flagsTD.ContactFldInput),
					"--" + flags.ContactNameFN, flagsTD.ContactNameInput,

					"--" + flags.EmailFN, flagsTD.FlgInputs(flagsTD.EmailInput),
					"--" + flags.EmailFolderFN, flagsTD.FlgInputs(flagsTD.EmailFldInput),
					"--" + flags.EmailReceivedAfterFN, flagsTD.EmailReceivedAfterInput,
					"--" + flags.EmailReceivedBeforeFN, flagsTD.EmailReceivedBeforeInput,
					"--" + flags.EmailSenderFN, flagsTD.EmailSenderInput,
					"--" + flags.EmailSubjectFN, flagsTD.EmailSubjectInput,

					"--" + flags.EventFN, flagsTD.FlgInputs(flagsTD.EventInput),
					"--" + flags.EventCalendarFN, flagsTD.FlgInputs(flagsTD.EventCalInput),
					"--" + flags.EventOrganizerFN, flagsTD.EventOrganizerInput,
					"--" + flags.EventRecursFN, flagsTD.EventRecursInput,
					"--" + flags.EventStartsAfterFN, flagsTD.EventStartsAfterInput,
					"--" + flags.EventStartsBeforeFN, flagsTD.EventStartsBeforeInput,
					"--" + flags.EventSubjectFN, flagsTD.EventSubjectInput,

					"--" + flags.CollisionsFN, flagsTD.Collisions,
					"--" + flags.DestinationFN, flagsTD.Destination,
					"--" + flags.ToResourceFN, flagsTD.ToResource,
				},
				flagsTD.PreparedProviderFlags(),
				flagsTD.PreparedStorageFlags())

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output

			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeExchangeOpts(cmd)
			assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

			assert.ElementsMatch(t, flagsTD.ContactInput, opts.Contact)
			assert.ElementsMatch(t, flagsTD.ContactFldInput, opts.ContactFolder)
			assert.Equal(t, flagsTD.ContactNameInput, opts.ContactName)

			assert.ElementsMatch(t, flagsTD.EmailInput, opts.Email)
			assert.ElementsMatch(t, flagsTD.EmailFldInput, opts.EmailFolder)
			assert.Equal(t, flagsTD.EmailReceivedAfterInput, opts.EmailReceivedAfter)
			assert.Equal(t, flagsTD.EmailReceivedBeforeInput, opts.EmailReceivedBefore)
			assert.Equal(t, flagsTD.EmailSenderInput, opts.EmailSender)
			assert.Equal(t, flagsTD.EmailSubjectInput, opts.EmailSubject)

			assert.ElementsMatch(t, flagsTD.EventInput, opts.Event)
			assert.ElementsMatch(t, flagsTD.EventCalInput, opts.EventCalendar)
			assert.Equal(t, flagsTD.EventOrganizerInput, opts.EventOrganizer)
			assert.Equal(t, flagsTD.EventRecursInput, opts.EventRecurs)
			assert.Equal(t, flagsTD.EventStartsAfterInput, opts.EventStartsAfter)
			assert.Equal(t, flagsTD.EventStartsBeforeInput, opts.EventStartsBefore)
			assert.Equal(t, flagsTD.EventSubjectInput, opts.EventSubject)

			assert.Equal(t, flagsTD.Collisions, opts.RestoreCfg.Collisions)
			assert.Equal(t, flagsTD.Destination, opts.RestoreCfg.Destination)
			assert.Equal(t, flagsTD.ToResource, opts.RestoreCfg.ProtectedResource)

			flagsTD.AssertProviderFlags(t, cmd)
			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
