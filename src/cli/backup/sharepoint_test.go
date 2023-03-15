package backup

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type SharePointSuite struct {
	tester.Suite
}

func TestSharePointSuite(t *testing.T) {
	suite.Run(t, &SharePointSuite{tester.NewUnitSuite(t)})
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
		suite.Run(test.name, func() {
			err := validateSharePointBackupCreateFlags(test.site, test.weburl, nil)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupCreateSelectors() {
	comboString := []string{"id_1", "id_2"}
	gc := &connector.GraphConnector{
		Sites: map[string]string{
			"url_1": "id_1",
			"url_2": "id_2",
		},
	}

	table := []struct {
		name            string
		site            []string
		weburl          []string
		data            []string
		expect          []string
		expectScopesLen int
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
			name:            "site wildcard",
			site:            []string{utils.Wildcard},
			expect:          selectors.Any(),
			expectScopesLen: 2,
		},
		{
			name:            "url wildcard",
			weburl:          []string{utils.Wildcard},
			expect:          selectors.Any(),
			expectScopesLen: 2,
		},
		{
			name:            "sites",
			site:            []string{"id_1", "id_2"},
			expect:          []string{"id_1", "id_2"},
			expectScopesLen: 2,
		},
		{
			name:            "urls",
			weburl:          []string{"url_1", "url_2"},
			expect:          []string{"id_1", "id_2"},
			expectScopesLen: 2,
		},
		{
			name:            "mix sites and urls",
			site:            []string{"id_1"},
			weburl:          []string{"url_2"},
			expect:          []string{"id_1", "id_2"},
			expectScopesLen: 2,
		},
		{
			name:            "duplicate sites and urls",
			site:            []string{"id_1", "id_2"},
			weburl:          []string{"url_1", "url_2"},
			expect:          comboString,
			expectScopesLen: 2,
		},
		{
			name:            "unnecessary site wildcard",
			site:            []string{"id_1", utils.Wildcard},
			weburl:          []string{"url_1", "url_2"},
			expect:          selectors.Any(),
			expectScopesLen: 2,
		},
		{
			name:            "unnecessary url wildcard",
			site:            comboString,
			weburl:          []string{"url_1", utils.Wildcard},
			expect:          selectors.Any(),
			expectScopesLen: 2,
		},
		{
			name:            "Pages",
			site:            comboString,
			data:            []string{dataPages},
			expect:          comboString,
			expectScopesLen: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			sel, err := sharePointBackupCreateSelectors(ctx, test.site, test.weburl, test.data, gc)
			require.NoError(t, err, clues.ToCore(err))

			assert.ElementsMatch(t, test.expect, sel.DiscreteResourceOwners())
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupDetailsSelectors() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.SharePointOptionDetailLookups {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsSharePointCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
				false)
			assert.NoError(t, err, clues.ToCore(err))
			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *SharePointSuite) TestSharePointBackupDetailsSelectorsBadFormats() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.BadSharePointOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsSharePointCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
				false)
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, output)
		})
	}
}
