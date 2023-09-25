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
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
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

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		groupsServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	// Test arg parsing for few args
	args := []string{
		groupsServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,

		"--" + flags.GroupFN, flagsTD.FlgInputs(flagsTD.GroupsInput),
		"--" + flags.CategoryDataFN, flagsTD.FlgInputs(flagsTD.GroupsCategoryDataInput),

		"--" + flags.FetchParallelismFN, flagsTD.FetchParallelism,

		// bool flags
		"--" + flags.FailFastFN,
		"--" + flags.DisableIncrementalsFN,
		"--" + flags.ForceItemDataDownloadFN,
		"--" + flags.DisableDeltaFN,
	}

	args = append(args, flagsTD.PreparedProviderFlags()...)
	args = append(args, flagsTD.PreparedStorageFlags()...)

	cmd.SetArgs(args)

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	opts := utils.MakeGroupsOpts(cmd)
	co := utils.Control()

	assert.ElementsMatch(t, flagsTD.GroupsInput, opts.Groups)
	// no assertion for category data input

	assert.Equal(t, flagsTD.FetchParallelism, strconv.Itoa(co.Parallelism.ItemFetch))

	// bool flags
	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)
	assert.True(t, co.ToggleFeatures.DisableDelta)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *GroupsUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: listCommand}

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		groupsServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
		},
		flagsTD.PreparedBackupListFlags(),
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

	flagsTD.AssertBackupListFlags(t, cmd)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *GroupsUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: detailsCommand}

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		groupsServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
			"--" + flags.SkipReduceFN,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	co := utils.Control()

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

	assert.True(t, co.SkipReduce)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *GroupsUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: deleteCommand}

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addGroupsCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		groupsServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	// Test arg parsing for few args
	args := []string{
		groupsServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,
		"--" + flags.BackupFN, flagsTD.BackupInput,
	}

	args = append(args, flagsTD.PreparedProviderFlags()...)
	args = append(args, flagsTD.PreparedStorageFlags()...)

	cmd.SetArgs(args)

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}
