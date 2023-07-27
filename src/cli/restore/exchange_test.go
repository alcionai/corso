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

			// normally a persistent flag from the root.
			// required to ensure a dry run.
			flags.AddRunModeFlag(cmd, true)

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
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, testdata.BackupInput,

				"--" + flags.ContactFN, testdata.FlgInputs(testdata.ContactInput),
				"--" + flags.ContactFolderFN, testdata.FlgInputs(testdata.ContactFldInput),
				"--" + flags.ContactNameFN, testdata.ContactNameInput,

				"--" + flags.EmailFN, testdata.FlgInputs(testdata.EmailInput),
				"--" + flags.EmailFolderFN, testdata.FlgInputs(testdata.EmailFldInput),
				"--" + flags.EmailReceivedAfterFN, testdata.EmailReceivedAfterInput,
				"--" + flags.EmailReceivedBeforeFN, testdata.EmailReceivedBeforeInput,
				"--" + flags.EmailSenderFN, testdata.EmailSenderInput,
				"--" + flags.EmailSubjectFN, testdata.EmailSubjectInput,

				"--" + flags.EventFN, testdata.FlgInputs(testdata.EventInput),
				"--" + flags.EventCalendarFN, testdata.FlgInputs(testdata.EventCalInput),
				"--" + flags.EventOrganizerFN, testdata.EventOrganizerInput,
				"--" + flags.EventRecursFN, testdata.EventRecursInput,
				"--" + flags.EventStartsAfterFN, testdata.EventStartsAfterInput,
				"--" + flags.EventStartsBeforeFN, testdata.EventStartsBeforeInput,
				"--" + flags.EventSubjectFN, testdata.EventSubjectInput,

				"--" + flags.CollisionsFN, testdata.Collisions,
				"--" + flags.DestinationFN, testdata.Destination,
				"--" + flags.ToResourceFN, testdata.ToResource,

				"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
				"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
				"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

				"--" + flags.AzureClientIDFN, testdata.AzureClientID,
				"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
				"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

				"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeExchangeOpts(cmd)
			assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

			assert.ElementsMatch(t, testdata.ContactInput, opts.Contact)
			assert.ElementsMatch(t, testdata.ContactFldInput, opts.ContactFolder)
			assert.Equal(t, testdata.ContactNameInput, opts.ContactName)

			assert.ElementsMatch(t, testdata.EmailInput, opts.Email)
			assert.ElementsMatch(t, testdata.EmailFldInput, opts.EmailFolder)
			assert.Equal(t, testdata.EmailReceivedAfterInput, opts.EmailReceivedAfter)
			assert.Equal(t, testdata.EmailReceivedBeforeInput, opts.EmailReceivedBefore)
			assert.Equal(t, testdata.EmailSenderInput, opts.EmailSender)
			assert.Equal(t, testdata.EmailSubjectInput, opts.EmailSubject)

			assert.ElementsMatch(t, testdata.EventInput, opts.Event)
			assert.ElementsMatch(t, testdata.EventCalInput, opts.EventCalendar)
			assert.Equal(t, testdata.EventOrganizerInput, opts.EventOrganizer)
			assert.Equal(t, testdata.EventRecursInput, opts.EventRecurs)
			assert.Equal(t, testdata.EventStartsAfterInput, opts.EventStartsAfter)
			assert.Equal(t, testdata.EventStartsBeforeInput, opts.EventStartsBefore)
			assert.Equal(t, testdata.EventSubjectInput, opts.EventSubject)

			assert.Equal(t, testdata.Collisions, opts.RestoreCfg.Collisions)
			assert.Equal(t, testdata.Destination, opts.RestoreCfg.Destination)
			assert.Equal(t, testdata.ToResource, opts.RestoreCfg.ProtectedResource)

			assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
			assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
			assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

			assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
			assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
			assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

			assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)
		})
	}
}
