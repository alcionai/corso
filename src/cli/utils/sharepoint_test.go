package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type SharePointUtilsSuite struct {
	tester.Suite
}

func TestSharePointUtilsSuite(t *testing.T) {
	suite.Run(t, &SharePointUtilsSuite{Suite: tester.NewUnitSuite(t)})
}

// Tests selector build for SharePoint properly
// differentiates between the 3 categories: Pages, Libraries and Lists CLI
func (suite *SharePointUtilsSuite) TestIncludeSharePointRestoreDataSelectors() {
	var (
		empty             = []string{}
		single            = []string{"single"}
		multi             = []string{"more", "than", "one"}
		containsOnly      = []string{"contains"}
		prefixOnly        = []string{"/prefix"}
		containsAndPrefix = []string{"contains", "/prefix"}
		onlySlash         = []string{string(path.PathSeparator)}
	)

	table := []struct {
		name             string
		opts             utils.SharePointOpts
		expectIncludeLen int
	}{
		{
			name:             "no inputs",
			opts:             utils.SharePointOpts{},
			expectIncludeLen: 3,
		},
		{
			name: "single inputs",
			opts: utils.SharePointOpts{
				FileName:   single,
				FolderPath: single,
				SiteID:     single,
				WebURL:     single,
			},
			expectIncludeLen: 4,
		},
		{
			name: "single extended",
			opts: utils.SharePointOpts{
				FileName:   single,
				FolderPath: single,
				ListItem:   single,
				ListFolder: single,
				SiteID:     single,
				WebURL:     single,
			},
			expectIncludeLen: 5,
		},
		{
			name: "multi inputs",
			opts: utils.SharePointOpts{
				FileName:   multi,
				FolderPath: multi,
				SiteID:     multi,
				WebURL:     multi,
			},
			expectIncludeLen: 4,
		},
		{
			name: "library folder contains",
			opts: utils.SharePointOpts{
				FileName:   empty,
				FolderPath: containsOnly,
				SiteID:     empty,
				WebURL:     empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library folder prefixes",
			opts: utils.SharePointOpts{
				FileName:   empty,
				FolderPath: prefixOnly,
				SiteID:     empty,
				WebURL:     empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library folder prefixes and contains",
			opts: utils.SharePointOpts{
				FileName:   empty,
				FolderPath: containsAndPrefix,
				SiteID:     empty,
				WebURL:     empty,
			},
			expectIncludeLen: 2,
		},
		{
			name: "list contains",
			opts: utils.SharePointOpts{
				FileName:   empty,
				FolderPath: empty,
				ListItem:   empty,
				ListFolder: containsOnly,
				SiteID:     empty,
				WebURL:     empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes",
			opts: utils.SharePointOpts{
				ListFolder: prefixOnly,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes and contains",
			opts: utils.SharePointOpts{
				ListFolder: containsAndPrefix,
			},
			expectIncludeLen: 2,
		},
		{
			name: "weburl contains",
			opts: utils.SharePointOpts{
				FileName:   empty,
				FolderPath: empty,
				SiteID:     empty,
				WebURL:     containsOnly,
			},
			expectIncludeLen: 3,
		},
		{
			name: "library folder suffixes",
			opts: utils.SharePointOpts{
				FileName:   empty,
				FolderPath: empty,
				SiteID:     empty,
				WebURL:     prefixOnly, // prefix pattern matches suffix pattern
			},
			expectIncludeLen: 3,
		},
		{
			name: "library folder suffixes and contains",
			opts: utils.SharePointOpts{
				FileName:   empty,
				FolderPath: empty,
				SiteID:     empty,
				WebURL:     containsAndPrefix, // prefix pattern matches suffix pattern
			},
			expectIncludeLen: 3,
		},
		{
			name: "Page Folder",
			opts: utils.SharePointOpts{
				PageFolder: single,
			},
			expectIncludeLen: 1,
		},
		{
			name: "Site Page ",
			opts: utils.SharePointOpts{
				Page: single,
			},
			expectIncludeLen: 1,
		},
		{
			name: "Page & library Files",
			opts: utils.SharePointOpts{
				PageFolder: single,
				FileName:   multi,
			},
			expectIncludeLen: 2,
		},
		{
			name: "folder with just /",
			opts: utils.SharePointOpts{
				FolderPath: onlySlash,
			},
			expectIncludeLen: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel := utils.IncludeSharePointRestoreDataSelectors(ctx, test.opts)
			assert.Len(suite.T(), sel.Includes, test.expectIncludeLen)
		})
	}
}

func (suite *SharePointUtilsSuite) TestIncludeSharePointRestoreDataSelectors_normalizedWebURLs() {
	table := []struct {
		name   string
		weburl string
		expect []string
	}{
		{
			name:   "blank",
			weburl: "",
			expect: []string{""},
		},
		{
			name:   "wildcard",
			weburl: "*",
			expect: []string{"*"},
		},
		{
			name:   "no scheme",
			weburl: "www.corsobackup.io/path",
			expect: []string{"https://www.corsobackup.io/path"},
		},
		{
			name:   "no path",
			weburl: "https://www.corsobackup.io",
			expect: []string{"https://www.corsobackup.io"},
		},
		{
			name:   "http",
			weburl: "http://www.corsobackup.io/path",
			expect: []string{"https://www.corsobackup.io/path"},
		},
		{
			name:   "https",
			weburl: "https://www.corsobackup.io/path",
			expect: []string{"https://www.corsobackup.io/path"},
		},
		{
			name:   "path only",
			weburl: "/path",
			expect: []string{"/path"},
		},
		{
			name:   "host only",
			weburl: "www.corsobackup.io",
			expect: []string{"https://www.corsobackup.io"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				opts = utils.SharePointOpts{WebURL: []string{test.weburl}}
				sel  = utils.IncludeSharePointRestoreDataSelectors(ctx, opts)
			)

			for _, scope := range sel.Scopes() {
				if scope.InfoCategory() != selectors.SharePointWebURL {
					continue
				}

				assert.ElementsMatch(suite.T(), test.expect, scope.Get(selectors.SharePointWebURL))
			}
		})
	}
}

func (suite *SharePointUtilsSuite) TestValidateSharePointRestoreFlags() {
	table := []struct {
		name     string
		backupID string
		opts     utils.SharePointOpts
		expect   assert.ErrorAssertionFunc
	}{
		{
			name:     "no opts",
			backupID: "id",
			opts:     utils.SharePointOpts{},
			expect:   assert.NoError,
		},
		{
			name:     "all valid",
			backupID: "id",
			opts: utils.SharePointOpts{
				WebURL:             []string{"www.corsobackup.io/sites/foo"},
				FileCreatedAfter:   dttm.Now(),
				FileCreatedBefore:  dttm.Now(),
				FileModifiedAfter:  dttm.Now(),
				FileModifiedBefore: dttm.Now(),
				Populated: flags.PopulatedFlags{
					flags.SiteFN:               struct{}{},
					flags.FileCreatedAfterFN:   struct{}{},
					flags.FileCreatedBeforeFN:  struct{}{},
					flags.FileModifiedAfterFN:  struct{}{},
					flags.FileModifiedBeforeFN: struct{}{},
				},
			},
			expect: assert.NoError,
		},
		{
			name:     "no backupID",
			backupID: "",
			opts:     utils.SharePointOpts{},
			expect:   assert.Error,
		},
		{
			name:     "invalid weburl",
			backupID: "id",
			opts: utils.SharePointOpts{
				WebURL: []string{"slander://:vree.garbles/:"},
				Populated: flags.PopulatedFlags{
					flags.SiteFN: struct{}{},
				},
			},
			expect: assert.Error,
		},
		{
			name:     "invalid file created after",
			backupID: "id",
			opts: utils.SharePointOpts{
				FileCreatedAfter: "1235",
				Populated: flags.PopulatedFlags{
					flags.FileCreatedAfterFN: struct{}{},
				},
			},
			expect: assert.Error,
		},
		{
			name:     "invalid file created before",
			backupID: "id",
			opts: utils.SharePointOpts{
				FileCreatedBefore: "1235",
				Populated: flags.PopulatedFlags{
					flags.FileCreatedBeforeFN: struct{}{},
				},
			},
			expect: assert.Error,
		},
		{
			name:     "invalid file modified after",
			backupID: "id",
			opts: utils.SharePointOpts{
				FileModifiedAfter: "1235",
				Populated: flags.PopulatedFlags{
					flags.FileModifiedAfterFN: struct{}{},
				},
			},
			expect: assert.Error,
		},
		{
			name:     "invalid file modified before",
			backupID: "id",
			opts: utils.SharePointOpts{
				FileModifiedBefore: "1235",
				Populated: flags.PopulatedFlags{
					flags.FileModifiedBeforeFN: struct{}{},
				},
			},
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			test.expect(t, utils.ValidateSharePointRestoreFlags(test.backupID, test.opts))
		})
	}
}
