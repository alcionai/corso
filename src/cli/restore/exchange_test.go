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
	"github.com/alcionai/corso/src/pkg/credentials"
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
				"--" + utils.BackupFN, testdata.BackupInput,

				"--" + utils.ContactFN, testdata.FlgInputs(testdata.ContactInput),
				"--" + utils.ContactFolderFN, testdata.FlgInputs(testdata.ContactFldInput),
				"--" + utils.ContactNameFN, testdata.ContactNameInput,

				"--" + utils.EmailFN, testdata.FlgInputs(testdata.EmailInput),
				"--" + utils.EmailFolderFN, testdata.FlgInputs(testdata.EmailFldInput),
				"--" + utils.EmailReceivedAfterFN, testdata.EmailReceivedAfterInput,
				"--" + utils.EmailReceivedBeforeFN, testdata.EmailReceivedBeforeInput,
				"--" + utils.EmailSenderFN, testdata.EmailSenderInput,
				"--" + utils.EmailSubjectFN, testdata.EmailSubjectInput,

				"--" + utils.EventFN, testdata.FlgInputs(testdata.EventInput),
				"--" + utils.EventCalendarFN, testdata.FlgInputs(testdata.EventCalInput),
				"--" + utils.EventOrganizerFN, testdata.EventOrganizerInput,
				"--" + utils.EventRecursFN, testdata.EventRecursInput,
				"--" + utils.EventStartsAfterFN, testdata.EventStartsAfterInput,
				"--" + utils.EventStartsBeforeFN, testdata.EventStartsBeforeInput,
				"--" + utils.EventSubjectFN, testdata.EventSubjectInput,

				"--" + utils.AWSAccessKeyFN, testdata.AWSAccessKeyID,
				"--" + utils.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
				"--" + utils.AWSSessionTokenFN, testdata.AWSSessionToken,

				"--" + utils.AzureClientIDFN, testdata.AzureClientID,
				"--" + utils.AzureClientTenantFN, testdata.AzureTenantID,
				"--" + utils.AzureClientSecretFN, testdata.AzureClientSecret,

				"--" + credentials.CorsoPassphraseFN, testdata.CorsoPassphrase,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeExchangeOpts(cmd)
			assert.Equal(t, testdata.BackupInput, utils.BackupIDFV)

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

			assert.Equal(t, testdata.AWSAccessKeyID, utils.AWSAccessKeyFV)
			assert.Equal(t, testdata.AWSSecretAccessKey, utils.AWSSecretAccessKeyFV)
			assert.Equal(t, testdata.AWSSessionToken, utils.AWSSessionTokenFV)

			assert.Equal(t, testdata.AzureClientID, credentials.AzureClientIDFV)
			assert.Equal(t, testdata.AzureTenantID, credentials.AzureClientTenantFV)
			assert.Equal(t, testdata.AzureClientSecret, credentials.AzureClientSecretFV)

			assert.Equal(t, testdata.CorsoPassphrase, credentials.CorsoPassphraseFV)
		})
	}
}
