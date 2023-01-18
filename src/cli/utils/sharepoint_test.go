package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
)

type SharePointUtilsSuite struct {
	suite.Suite
}

func TestSharePointUtilsSuite(t *testing.T) {
	suite.Run(t, new(SharePointUtilsSuite))
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
				LibraryItems: single,
				LibraryPaths: single,
				Sites:        single,
				WebURLs:      single,
			},
			expectIncludeLen: 4,
		},
		{
			name: "single extended",
			opts: utils.SharePointOpts{
				LibraryItems: single,
				LibraryPaths: single,
				ListItems:    single,
				ListPaths:    single,
				Sites:        single,
				WebURLs:      single,
			},
			expectIncludeLen: 5,
		},
		{
			name: "multi inputs",
			opts: utils.SharePointOpts{
				LibraryItems: multi,
				LibraryPaths: multi,
				Sites:        multi,
				WebURLs:      multi,
			},
			expectIncludeLen: 4,
		},
		{
			name: "library contains",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: containsOnly,
				Sites:        empty,
				WebURLs:      empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library prefixes",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: prefixOnly,
				Sites:        empty,
				WebURLs:      empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library prefixes and contains",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: containsAndPrefix,
				Sites:        empty,
				WebURLs:      empty,
			},
			expectIncludeLen: 2,
		},
		{
			name: "list contains",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: empty,
				ListItems:    empty,
				ListPaths:    containsOnly,
				Sites:        empty,
				WebURLs:      empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes",
			opts: utils.SharePointOpts{
				ListPaths: prefixOnly,
			},
			expectIncludeLen: 1,
		},
		{
			name: "list prefixes and contains",
			opts: utils.SharePointOpts{
				ListPaths: containsAndPrefix,
			},
			expectIncludeLen: 2,
		},
		{
			name: "weburl contains",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: empty,
				Sites:        empty,
				WebURLs:      containsOnly,
			},
			expectIncludeLen: 3,
		},
		{
			name: "library suffixes",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: empty,
				Sites:        empty,
				WebURLs:      prefixOnly, // prefix pattern matches suffix pattern
			},
			expectIncludeLen: 3,
		},
		{
			name: "library suffixes and contains",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: empty,
				Sites:        empty,
				WebURLs:      containsAndPrefix, // prefix pattern matches suffix pattern
			},
			expectIncludeLen: 6,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := utils.IncludeSharePointRestoreDataSelectors(test.opts)
			assert.Len(t, sel.Includes, test.expectIncludeLen)
		})
	}
}
