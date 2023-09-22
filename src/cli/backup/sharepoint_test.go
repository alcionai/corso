package backup

import (
	"bytes"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type SharePointUnitSuite struct {
	tester.Suite
}

func TestSharePointUnitSuite(t *testing.T) {
	suite.Run(t, &SharePointUnitSuite{tester.NewUnitSuite(t)})
}

func (suite *SharePointUnitSuite) TestAddSharePointCommands() {
	expectUse := sharePointServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			name:        "create sharepoint",
			use:         createCommand,
			expectUse:   expectUse + " " + sharePointServiceCommandCreateUseSuffix,
			expectShort: sharePointCreateCmd().Short,
			expectRunE:  createSharePointCmd,
		},
		{
			name:        "list sharepoint",
			use:         listCommand,
			expectUse:   expectUse,
			expectShort: sharePointListCmd().Short,
			expectRunE:  listSharePointCmd,
		},
		{
			name:        "details sharepoint",
			use:         detailsCommand,
			expectUse:   expectUse + " " + sharePointServiceCommandDetailsUseSuffix,
			expectShort: sharePointDetailsCmd().Short,
			expectRunE:  detailsSharePointCmd,
		},
		{
			name:        "delete sharepoint",
			use:         deleteCommand,
			expectUse:   expectUse + " " + sharePointServiceCommandDeleteUseSuffix,
			expectShort: sharePointDeleteCmd().Short,
			expectRunE:  deleteSharePointCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addSharePointCommands(cmd)
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

func (suite *SharePointUnitSuite) TestBackupCreateFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: createCommand}

	// persistent flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// non-persistent flags not added by addCommands
	flags.AddAllProviderFlags(c)
	flags.AddAllStorageFlags(c)

	flagsTD.WithFlags(
		cmd,
		sharePointServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.SiteIDFN, flagsTD.FlgInputs(flagsTD.SiteIDInput),
			"--" + flags.SiteFN, flagsTD.FlgInputs(flagsTD.WebURLInput),
			"--" + flags.CategoryDataFN, flagsTD.FlgInputs(flagsTD.SharepointCategoryDataInput),
			"--" + flags.FailFastFN,
			"--" + flags.DisableIncrementalsFN,
			"--" + flags.ForceItemDataDownloadFN,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	opts := utils.MakeSharePointOpts(cmd)
	co := utils.Control()

	assert.ElementsMatch(t, []string{strings.Join(flagsTD.SiteIDInput, ",")}, opts.SiteID)
	assert.ElementsMatch(t, flagsTD.WebURLInput, opts.WebURL)
	// no assertion for category data input

	// bool flags
	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *SharePointUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: listCommand}

	// persistent flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// non-persistent flags not added by addCommands
	flags.AddAllProviderFlags(c)
	flags.AddAllStorageFlags(c)

	flagsTD.WithFlags(
		cmd,
		sharePointServiceCommand,
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

func (suite *SharePointUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: detailsCommand}

	// persistent flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// non-persistent flags not added by addCommands
	flags.AddAllProviderFlags(c)
	flags.AddAllStorageFlags(c)

	flagsTD.WithFlags(
		cmd,
		sharePointServiceCommand,
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

func (suite *SharePointUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: deleteCommand}

	// persistent flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// non-persistent flags not added by addCommands
	flags.AddAllProviderFlags(c)
	flags.AddAllStorageFlags(c)

	flagsTD.WithFlags(
		cmd,
		sharePointServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *SharePointUnitSuite) TestValidateSharePointBackupCreateFlags() {
	table := []struct {
		name   string
		site   []string
		weburl []string
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "no sites or urls",
			expect: assert.Error,
		},
		{
			name:   "sites",
			site:   []string{"smarf"},
			expect: assert.NoError,
		},
		{
			name:   "urls",
			weburl: []string{"fnord"},
			expect: assert.NoError,
		},
		{
			name:   "both",
			site:   []string{"smarf"},
			weburl: []string{"fnord"},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := validateSharePointBackupCreateFlags(test.site, test.weburl, nil)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *SharePointUnitSuite) TestSharePointBackupCreateSelectors() {
	const (
		id1  = "id_1"
		id2  = "id_2"
		url1 = "url_1/foo"
		url2 = "url_2/bar"
	)

	var (
		ins     = idname.NewCache(map[string]string{id1: url1, id2: url2})
		bothIDs = []string{id1, id2}
	)

	table := []struct {
		name   string
		site   []string
		weburl []string
		data   []string
		expect []string
	}{
		{
			name:   "no sites or urls",
			expect: selectors.None(),
		},
		{
			name:   "empty sites and urls",
			site:   []string{},
			weburl: []string{},
			expect: selectors.None(),
		},
		{
			name:   "site wildcard",
			site:   []string{flags.Wildcard},
			expect: bothIDs,
		},
		{
			name:   "url wildcard",
			weburl: []string{flags.Wildcard},
			expect: bothIDs,
		},
		{
			name:   "sites",
			site:   []string{id1, id2},
			expect: []string{id1, id2},
		},
		{
			name:   "urls",
			weburl: []string{url1, url2},
			expect: []string{url1, url2},
		},
		{
			name:   "mix sites and urls",
			site:   []string{id1},
			weburl: []string{url2},
			expect: []string{id1, url2},
		},
		{
			name:   "duplicate sites and urls",
			site:   []string{id1, id2},
			weburl: []string{url1, url2},
			expect: []string{id1, id2, url1, url2},
		},
		{
			name:   "unnecessary site wildcard",
			site:   []string{id1, flags.Wildcard},
			weburl: []string{url1, url2},
			expect: bothIDs,
		},
		{
			name:   "unnecessary url wildcard",
			site:   []string{id1},
			weburl: []string{url1, flags.Wildcard},
			expect: bothIDs,
		},
		{
			name:   "Pages",
			site:   bothIDs,
			data:   []string{flags.DataPages},
			expect: bothIDs,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel, err := sharePointBackupCreateSelectors(ctx, ins, test.site, test.weburl, test.data)
			require.NoError(t, err, clues.ToCore(err))
			assert.ElementsMatch(t, test.expect, sel.ResourceOwners.Targets)
		})
	}
}
