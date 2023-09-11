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

type GroupsUtilsSuite struct {
	tester.Suite
}

func TestGroupsUtilsSuite(t *testing.T) {
	suite.Run(t, &GroupsUtilsSuite{Suite: tester.NewUnitSuite(t)})
}

// Tests selector build for Groups properly
// differentiates between the 3 categories: Pages, Libraries and Lists CLI
func (suite *GroupsUtilsSuite) TestIncludeGroupsRestoreDataSelectors() {
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
		opts             utils.GroupsOpts
		expectIncludeLen int
	}{
		{
			name:             "no inputs",
			opts:             utils.GroupsOpts{},
			expectIncludeLen: 2,
		},
		{
			name: "empty",
			opts: utils.GroupsOpts{
				Groups: empty,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single inputs",
			opts: utils.GroupsOpts{
				Groups: single,
			},
			expectIncludeLen: 2,
		},
		{
			name: "multi inputs",
			opts: utils.GroupsOpts{
				Groups: multi,
			},
			expectIncludeLen: 2,
		},
		{
			name: "library folder contains",
			opts: utils.GroupsOpts{
				FileName:   empty,
				FolderPath: containsOnly,
				SiteID:     empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library folder prefixes",
			opts: utils.GroupsOpts{
				FileName:   empty,
				FolderPath: prefixOnly,
				SiteID:     empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library folder prefixes and contains",
			opts: utils.GroupsOpts{
				FileName:   empty,
				FolderPath: containsAndPrefix,
				SiteID:     empty,
			},
			expectIncludeLen: 2,
		},
		{
			name: "list contains",
			opts: utils.GroupsOpts{
				FileName:   empty,
				FolderPath: empty,
				ListItem:   empty,
				ListFolder: containsOnly,
				SiteID:     empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes",
			opts: utils.GroupsOpts{
				ListFolder: prefixOnly,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes and contains",
			opts: utils.GroupsOpts{
				ListFolder: containsAndPrefix,
			},
			expectIncludeLen: 2,
		},
		{
			name: "library folder suffixes",
			opts: utils.GroupsOpts{
				FileName:   empty,
				FolderPath: empty,
				// SiteID:     empty,  // TODO(meain): Update once we support multiple sites
			},
			expectIncludeLen: 2,
		},
		{
			name: "library folder suffixes and contains",
			opts: utils.GroupsOpts{
				FileName:   empty,
				FolderPath: empty,
				// SiteID:     empty, // TODO(meain): update once we support multiple sites
			},
			expectIncludeLen: 2,
		},
		{
			name: "Page Folder",
			opts: utils.GroupsOpts{
				PageFolder: single,
			},
			expectIncludeLen: 1,
		},
		{
			name: "Site Page ",
			opts: utils.GroupsOpts{
				Page: single,
			},
			expectIncludeLen: 1,
		},
		{
			name: "Page & library Files",
			opts: utils.GroupsOpts{
				PageFolder: single,
				FileName:   multi,
			},
			expectIncludeLen: 2,
		},
		{
			name: "folder with just /",
			opts: utils.GroupsOpts{
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

			sel := utils.IncludeGroupsRestoreDataSelectors(ctx, test.opts)
			assert.Len(suite.T(), sel.Includes, test.expectIncludeLen)
		})
	}
}

func (suite *GroupsUtilsSuite) TestValidateGroupsRestoreFlags() {
	table := []struct {
		name     string
		backupID string
		opts     utils.GroupsOpts
		expect   assert.ErrorAssertionFunc
	}{
		{
			name:     "no opts",
			backupID: "id",
			opts:     utils.GroupsOpts{},
			expect:   assert.NoError,
		},
		{
			name:     "no backupID",
			backupID: "",
			opts:     utils.GroupsOpts{},
			expect:   assert.Error,
		},
		{
			name:     "all valid",
			backupID: "id",
			opts: utils.GroupsOpts{
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
			name:     "invalid file created after",
			backupID: "id",
			opts: utils.GroupsOpts{
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
			opts: utils.GroupsOpts{
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
			opts: utils.GroupsOpts{
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
			opts: utils.GroupsOpts{
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
			test.expect(t, utils.ValidateGroupsRestoreFlags(test.backupID, test.opts))
		})
	}
}

func (suite *GroupsUtilsSuite) TestAddGroupsCategories() {
	table := []struct {
		name           string
		cats           []string
		expectScopeLen int
	}{
		{
			name:           "none",
			cats:           []string{},
			expectScopeLen: 2,
		},
		{
			name:           "libraries",
			cats:           []string{flags.DataLibraries},
			expectScopeLen: 1,
		},
		{
			name:           "messages",
			cats:           []string{flags.DataMessages},
			expectScopeLen: 1,
		},
		{
			name:           "all allowed",
			cats:           []string{flags.DataLibraries, flags.DataMessages},
			expectScopeLen: 2,
		},
		{
			name:           "bad inputs",
			cats:           []string{"foo"},
			expectScopeLen: 0,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			sel := utils.AddGroupsCategories(selectors.NewGroupsBackup(selectors.Any()), test.cats)
			scopes := sel.Scopes()
			assert.Len(suite.T(), scopes, test.expectScopeLen)
		})
	}
}
