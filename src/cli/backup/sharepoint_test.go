package backup

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
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
		flags       []string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create sharepoint",
			createCommand,
			expectUse + " " + sharePointServiceCommandCreateUseSuffix,
			sharePointCreateCmd().Short,
			[]string{"site", "disable-incrementals", "fail-fast"},
			createSharePointCmd,
		},
		{
			"list sharepoint",
			listCommand,
			expectUse,
			sharePointListCmd().Short,
			[]string{"backup", "failed-items", "skipped-items", "recovered-errors"},
			listSharePointCmd,
		},
		{
			"details sharepoint",
			detailsCommand,
			expectUse + " " + sharePointServiceCommandDetailsUseSuffix,
			sharePointDetailsCmd().Short,
			[]string{
				"backup",
				"library",
				"folder",
				"file",
				"file-created-after",
				"file-created-before",
				"file-modified-after",
				"file-modified-before",
			},
			detailsSharePointCmd,
		},
		{
			"delete sharepoint",
			deleteCommand,
			expectUse + " " + sharePointServiceCommandDeleteUseSuffix,
			sharePointDeleteCmd().Short,
			[]string{"backup"},
			deleteSharePointCmd,
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

			for _, f := range test.flags {
				assert.NotNil(t, c.Flag(f), f+" flag")
			}
		})
	}
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
		ins = common.IDsNames{
			IDToName: map[string]string{id1: url1, id2: url2},
			NameToID: map[string]string{url1: id1, url2: id2},
		}
		bothIDs = []string{id1, id2}
	)

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
			expect:          bothIDs,
			expectScopesLen: 2,
		},
		{
			name:            "url wildcard",
			weburl:          []string{utils.Wildcard},
			expect:          bothIDs,
			expectScopesLen: 2,
		},
		{
			name:            "sites",
			site:            []string{id1, id2},
			expect:          []string{id1, id2},
			expectScopesLen: 2,
		},
		{
			name:            "urls",
			weburl:          []string{url1, url2},
			expect:          []string{url1, url2},
			expectScopesLen: 2,
		},
		{
			name:            "mix sites and urls",
			site:            []string{id1},
			weburl:          []string{url2},
			expect:          []string{id1, url2},
			expectScopesLen: 2,
		},
		{
			name:            "duplicate sites and urls",
			site:            []string{id1, id2},
			weburl:          []string{url1, url2},
			expect:          []string{id1, id2, url1, url2},
			expectScopesLen: 2,
		},
		{
			name:            "unnecessary site wildcard",
			site:            []string{id1, utils.Wildcard},
			weburl:          []string{url1, url2},
			expect:          bothIDs,
			expectScopesLen: 2,
		},
		{
			name:            "unnecessary url wildcard",
			site:            []string{id1},
			weburl:          []string{url1, utils.Wildcard},
			expect:          bothIDs,
			expectScopesLen: 2,
		},
		{
			name:            "Pages",
			site:            bothIDs,
			data:            []string{dataPages},
			expect:          bothIDs,
			expectScopesLen: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			sel, err := sharePointBackupCreateSelectors(ctx, ins, test.site, test.weburl, test.data)
			require.NoError(t, err, clues.ToCore(err))
			assert.ElementsMatch(t, test.expect, sel.DiscreteResourceOwners())
		})
	}
}

func (suite *SharePointUnitSuite) TestSharePointBackupDetailsSelectors() {
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

func (suite *SharePointUnitSuite) TestSharePointBackupDetailsSelectorsBadFormats() {
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
