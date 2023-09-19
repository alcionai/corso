package backup

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	dtd "github.com/alcionai/corso/src/pkg/backup/details/testdata"
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

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		sharePointServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,

		"--" + flags.SiteIDFN, testdata.FlgInputs(testdata.SiteIDInput),
		"--" + flags.SiteFN, testdata.FlgInputs(testdata.WebURLInput),
		"--" + flags.CategoryDataFN, testdata.FlgInputs(testdata.SharepointCategoryDataInput),

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

		// bool flags
		"--" + flags.FailFastFN,
		"--" + flags.DisableIncrementalsFN,
		"--" + flags.ForceItemDataDownloadFN,
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	opts := utils.MakeSharePointOpts(cmd)
	co := utils.Control()

	assert.ElementsMatch(t, []string{strings.Join(testdata.SiteIDInput, ",")}, opts.SiteID)
	assert.ElementsMatch(t, testdata.WebURLInput, opts.WebURL)
	// no assertion for category data input

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

	// bool flags
	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)
}

func (suite *SharePointUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: listCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		sharePointServiceCommand,
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

func (suite *SharePointUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: detailsCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		sharePointServiceCommand,
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

func (suite *SharePointUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: deleteCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addSharePointCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		sharePointServiceCommand,
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

func (suite *SharePointUnitSuite) TestSharePointBackupDetailsSelectors() {
	for v := 0; v <= version.Backup; v++ {
		suite.Run(fmt.Sprintf("version%d", v), func() {
			for _, test := range testdata.SharePointOptionDetailLookups {
				suite.Run(test.Name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					bg := testdata.VersionedBackupGetter{
						Details: dtd.GetDetailsSetForVersion(t, v),
					}

					output, err := runDetailsSharePointCmd(
						ctx,
						bg,
						"backup-ID",
						test.Opts(t, v),
						false)
					assert.NoError(t, err, clues.ToCore(err))
					assert.ElementsMatch(t, test.Expected(t, v), output.Entries)
				})
			}
		})
	}
}

func (suite *SharePointUnitSuite) TestSharePointBackupDetailsSelectorsBadFormats() {
	for _, test := range testdata.BadSharePointOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			output, err := runDetailsSharePointCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts(t, version.Backup),
				false)
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, output)
		})
	}
}
