package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type SharePointUtilsSuite struct {
	suite.Suite
}

func TestSharePointUtilsSuite(t *testing.T) {
	suite.Run(t, new(SharePointUtilsSuite))
}

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
			name: "no inputs",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: empty,
				Sites:        empty,
				WebURLs:      empty,
			},
			expectIncludeLen: 2,
		},
		{
			name: "single inputs",
			opts: utils.SharePointOpts{
				LibraryItems: single,
				LibraryPaths: single,
				Sites:        single,
				WebURLs:      single,
			},
			expectIncludeLen: 3,
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
			expectIncludeLen: 4,
		},
		{
			name: "multi inputs",
			opts: utils.SharePointOpts{
				LibraryItems: multi,
				LibraryPaths: multi,
				Sites:        multi,
				WebURLs:      multi,
			},
			expectIncludeLen: 3,
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
			expectIncludeLen: 2,
		},
		{
			name: "library suffixes",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: empty,
				Sites:        empty,
				WebURLs:      prefixOnly, // prefix pattern matches suffix pattern
			},
			expectIncludeLen: 2,
		},
		{
			name: "library suffixes and contains",
			opts: utils.SharePointOpts{
				LibraryItems: empty,
				LibraryPaths: empty,
				Sites:        empty,
				WebURLs:      containsAndPrefix, // prefix pattern matches suffix pattern
			},
			expectIncludeLen: 4,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewSharePointRestore(nil)
			// no return, mutates sel as a side effect
			t.Logf("Options sent: %v\n", test.opts)
			utils.IncludeSharePointRestoreDataSelectors(sel, test.opts)
			assert.Len(t, sel.Includes, test.expectIncludeLen, sel)
		})
	}
}
