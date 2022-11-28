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

func (suite *ExchangeUtilsSuite) TestIncludeSharePointRestoreDataSelectors() {
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
				Sites:        empty,
				Libraries:    empty,
				LibraryItems: empty,
			},
			expectIncludeLen: 0,
		},
		{
			name: "single inputs",
			opts: utils.SharePointOpts{
				Sites:        single,
				Libraries:    single,
				LibraryItems: single,
			},
			expectIncludeLen: 1,
		},
		{
			name: "multi inputs",
			opts: utils.SharePointOpts{
				Sites:        multi,
				Libraries:    multi,
				LibraryItems: multi,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library contains",
			opts: utils.SharePointOpts{
				Sites:        empty,
				Libraries:    containsOnly,
				LibraryItems: empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library prefixes",
			opts: utils.SharePointOpts{
				Sites:        empty,
				Libraries:    prefixOnly,
				LibraryItems: empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "library prefixes and contains",
			opts: utils.SharePointOpts{
				Sites:        empty,
				Libraries:    containsAndPrefix,
				LibraryItems: empty,
			},
			expectIncludeLen: 2,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewSharePointRestore()
			// no return, mutates sel as a side effect
			utils.IncludeSharePointRestoreDataSelectors(sel, test.opts)
			assert.Len(t, sel.Includes, test.expectIncludeLen)
		})
	}
}
