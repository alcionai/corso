package backup

import (
	"strconv"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/cli/utils"
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
			name:   "conversations",
			cats:   []string{flags.DataConversations},
			expect: assert.NoError,
		},
		{
			name: "all allowed",
			cats: []string{
				flags.DataLibraries,
				flags.DataMessages,
				flags.DataConversations,
			},
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

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: createCommand},
		addGroupsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			groupsServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.GroupFN, flagsTD.FlgInputs(flagsTD.GroupsInput),
				"--" + flags.CategoryDataFN, flagsTD.FlgInputs(flagsTD.GroupsCategoryDataInput),
				"--" + flags.FetchParallelismFN, flagsTD.FetchParallelism,
				"--" + flags.DisableDeltaFN,
			},
			flagsTD.PreparedGenericBackupFlags(),
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	opts := utils.MakeGroupsOpts(cmd)
	co := utils.Control()
	backupOpts := utils.ParseBackupOptions()

	// TODO(ashmrtn): Remove flag checks on control.Options to control.Backup once
	// restore flags are switched over too and we no longer parse flags beyond
	// connection info into control.Options.
	assert.Equal(t, flagsTD.FetchParallelism, strconv.Itoa(backupOpts.Parallelism.ItemFetch))
	assert.Equal(t, control.FailFast, backupOpts.FailureHandling)
	assert.True(t, backupOpts.Incrementals.ForceFullEnumeration)
	assert.True(t, backupOpts.Incrementals.ForceItemDataRefresh)
	assert.True(t, backupOpts.M365.DisableDeltaEndpoint)

	assert.Equal(t, flagsTD.FetchParallelism, strconv.Itoa(co.Parallelism.ItemFetch))
	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)
	assert.True(t, co.ToggleFeatures.DisableDelta)

	assert.ElementsMatch(t, flagsTD.GroupsInput, opts.Groups)
	flagsTD.AssertGenericBackupFlags(t, cmd)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *GroupsUnitSuite) TestBackupCreateDefaultControlFlags() {
	t := suite.T()

	cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: createCommand},
		addGroupsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			groupsServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
			}))

	co := utils.Control()
	backupOpts := utils.ParseBackupOptions()

	assert.Equal(t, co.Parallelism.ItemFetch, backupOpts.Parallelism.ItemFetch)
	assert.Equal(t, co.DeltaPageSize, backupOpts.M365.DeltaPageSize)
	assert.Equal(t, co.FailureHandling, backupOpts.FailureHandling)
	assert.Equal(
		t,
		co.ToggleFeatures.DisableIncrementals,
		backupOpts.Incrementals.ForceFullEnumeration)
	assert.Equal(
		t,
		co.ToggleFeatures.ForceItemDataDownload,
		backupOpts.Incrementals.ForceItemDataRefresh)
	assert.Equal(
		t,
		co.ToggleFeatures.DisableDelta,
		backupOpts.M365.DisableDeltaEndpoint)
	assert.Equal(
		t,
		co.ToggleFeatures.ExchangeImmutableIDs,
		backupOpts.M365.ExchangeImmutableIDs)
	assert.Equal(
		t,
		co.ToggleFeatures.DisableSlidingWindowLimiter,
		backupOpts.ServiceRateLimiter.DisableSlidingWindowLimiter)
}

func (suite *GroupsUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: listCommand},
		addGroupsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			groupsServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
			},
			flagsTD.PreparedBackupListFlags(),
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
	flagsTD.AssertBackupListFlags(t, cmd)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *GroupsUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: detailsCommand},
		addGroupsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			groupsServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
				"--" + flags.SkipReduceFN,
			},
			flagsTD.PreparedChannelFlags(),
			flagsTD.PreparedConversationFlags(),
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags(),
			flagsTD.PreparedLibraryFlags()))

	co := utils.Control()

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
	assert.True(t, co.SkipReduce)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
	flagsTD.AssertChannelFlags(t, cmd)
	flagsTD.AssertConversationFlags(t, cmd)
	flagsTD.AssertLibraryFlags(t, cmd)
}

func (suite *GroupsUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: deleteCommand},
		addGroupsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			groupsServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
			},
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}
