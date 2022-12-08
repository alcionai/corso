package backup

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type SharePointSuite struct {
	suite.Suite
}

func TestSharePointSuite(t *testing.T) {
	suite.Run(t, new(SharePointSuite))
}

func (suite *SharePointSuite) TestAddSharePointCommands() {
	expectUse := sharePointServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create sharepoint", createCommand, expectUse + " " + sharePointServiceCommandCreateUseSuffix,
			sharePointCreateCmd().Short, createSharePointCmd,
		},
		{
			"list sharepoint", listCommand, expectUse,
			sharePointListCmd().Short, listSharePointCmd,
		},
		{
			"details sharepoint", detailsCommand, expectUse + " " + sharePointServiceCommandDetailsUseSuffix,
			sharePointDetailsCmd().Short, detailsSharePointCmd,
		},
		{
			"delete sharepoint", deleteCommand, expectUse + " " + sharePointServiceCommandDeleteUseSuffix,
			sharePointDeleteCmd().Short, deleteSharePointCmd,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
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

func (suite *SharePointSuite) TestValidateSharePointBackupCreateFlags() {
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
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, validateSharePointBackupCreateFlags(test.site, test.weburl))
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupCreateSelectors() {
	gc := &connector.GraphConnector{
		Sites: map[string]string{
			"url_1": "id_1",
			"url_2": "id_2",
		},
	}

	table := []struct {
		name   string
		site   []string
		weburl []string
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
			site:   []string{utils.Wildcard},
			expect: selectors.Any(),
		},
		{
			name:   "url wildcard",
			weburl: []string{utils.Wildcard},
			expect: selectors.Any(),
		},
		{
			name:   "sites",
			site:   []string{"id_1", "id_2"},
			expect: []string{"id_1", "id_2"},
		},
		{
			name:   "urls",
			weburl: []string{"url_1", "url_2"},
			expect: []string{"id_1", "id_2"},
		},
		{
			name:   "mix sites and urls",
			site:   []string{"id_1"},
			weburl: []string{"url_2"},
			expect: []string{"id_1", "id_2"},
		},
		{
			name:   "duplicate sites and urls",
			site:   []string{"id_1", "id_2"},
			weburl: []string{"url_1", "url_2"},
			expect: []string{"id_1", "id_2"},
		},
		{
			name:   "unnecessary site wildcard",
			site:   []string{"id_1", utils.Wildcard},
			weburl: []string{"url_1", "url_2"},
			expect: selectors.Any(),
		},
		{
			name:   "unnecessary url wildcard",
			site:   []string{"id_1", "id_2"},
			weburl: []string{"url_1", utils.Wildcard},
			expect: selectors.Any(),
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			sel, err := sharePointBackupCreateSelectors(ctx, test.site, test.weburl, gc)
			require.NoError(t, err)

			scope := sel.Scopes()[0]
			targetSites := scope.Get(selectors.SharePointSite)

			assert.ElementsMatch(t, test.expect, targetSites)
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupDetailsSelectors() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.SharePointOptionDetailLookups {
		suite.T().Run(test.Name, func(t *testing.T) {
			output, err := runDetailsSharePointCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
			)
			assert.NoError(t, err)

			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupDetailsSelectorsBadFormats() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.BadSharePointOptionsFormats {
		suite.T().Run(test.Name, func(t *testing.T) {
			output, err := runDetailsSharePointCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
			)

			assert.Error(t, err)
			assert.Empty(t, output)
		})
	}
}
