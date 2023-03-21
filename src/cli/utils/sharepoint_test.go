package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/tester"
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
				ListPath:   single,
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
				ListPath:   containsOnly,
				SiteID:     empty,
				WebURL:     empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes",
			opts: utils.SharePointOpts{
				ListPath: prefixOnly,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes and contains",
			opts: utils.SharePointOpts{
				ListPath: containsAndPrefix,
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
			expectIncludeLen: 6,
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
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			sel := utils.IncludeSharePointRestoreDataSelectors(test.opts)
			assert.Len(suite.T(), sel.Includes, test.expectIncludeLen)
		})
	}
}
