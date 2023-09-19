package backup

import (
	"bytes"
	"strconv"
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
	"github.com/alcionai/corso/src/pkg/control"
)

type GroupsUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsUnitSuite) TestAddGroupsCommands() {
	expectUse := groupsServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			name:        "create groups",
			use:         createCommand,
			expectUse:   expectUse + " " + groupsServiceCommandCreateUseSuffix,
			expectShort: groupsCreateCmd().Short,
			expectRunE:  createGroupsCmd,
		},
		{
			name:        "list groups",
			use:         listCommand,
			expectUse:   expectUse,
			expectShort: groupsListCmd().Short,
			expectRunE:  listGroupsCmd,
		},
		{
			name:        "details groups",
			use:         detailsCommand,
			expectUse:   expectUse + " " + groupsServiceCommandDetailsUseSuffix,
			expectShort: groupsDetailsCmd().Short,
			expectRunE:  detailsGroupsCmd,
		},
		{
			name:        "delete groups",
			use:         deleteCommand,
			expectUse:   expectUse + " " + groupsServiceCommandDeleteUseSuffix,
			expectShort: groupsDeleteCmd().Short,
			expectRunE:  deleteGroupsCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addGroupsCommands(cmd)
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

func (suite *GroupsUnitSuite) TestValidateGroupsBackupCreateFlags() {
	table := []struct {
		name   string
		cats   []string
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "none",
			cats:   []string{},
			expect: assert.NoError,
		},
		{
			name:   "libraries",
			cats:   []string{flags.DataLibraries},
			expect: assert.NoError,
		},
		{
			name:   "messages",
			cats:   []string{flags.DataMessages},
			expect: assert.NoError,
		},
		{
			name:   "all allowed",
			cats:   []string{flags.DataLibraries, flags.DataMessages},
			expect: assert.NoError,
		},
		{
			name:   "bad inputs",
			cats:   []string{"foo"},
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := validateGroupsBackupCreateFlags([]string{"*"}, test.cats)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *GroupsUnitSuite) TestBackupCreateFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: createCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		groupsServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,

		"--" + flags.GroupFN, testdata.FlgInputs(testdata.GroupsInput),
		"--" + flags.CategoryDataFN, testdata.FlgInputs(testdata.GroupsCategoryDataInput),

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

		"--" + flags.FetchParallelismFN, testdata.FetchParallelism,

		// bool flags
		"--" + flags.FailFastFN,
		"--" + flags.DisableIncrementalsFN,
		"--" + flags.ForceItemDataDownloadFN,
		"--" + flags.DisableDeltaFN,
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	opts := utils.MakeGroupsOpts(cmd)
	co := utils.Control()

	assert.ElementsMatch(t, testdata.GroupsInput, opts.Groups)
	// no assertion for category data input

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

	assert.Equal(t, testdata.FetchParallelism, strconv.Itoa(co.Parallelism.ItemFetch))

	// bool flags
	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)
	assert.True(t, co.ToggleFeatures.DisableDelta)
}

func (suite *GroupsUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: listCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		groupsServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,
		"--" + flags.BackupFN, testdata.BackupInput,

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

		// bool flags
		"--" + flags.FailedItemsFN, "show",
		"--" + flags.SkippedItemsFN, "show",
		"--" + flags.RecoveredErrorsFN, "show",
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

	assert.Equal(t, flags.ListFailedItemsFV, "show")
	assert.Equal(t, flags.ListSkippedItemsFV, "show")
	assert.Equal(t, flags.ListRecoveredErrorsFV, "show")
}

func (suite *GroupsUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: detailsCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		groupsServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,
		"--" + flags.BackupFN, testdata.BackupInput,

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

		// bool flags
		"--" + flags.SkipReduceFN,
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	co := utils.Control()

	assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

	assert.True(t, co.SkipReduce)
}

func (suite *GroupsUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: deleteCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		groupsServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,
		"--" + flags.BackupFN, testdata.BackupInput,

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

	assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)
}
