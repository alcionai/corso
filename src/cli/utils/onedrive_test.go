package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/tester"
)

type OneDriveUtilsSuite struct {
	tester.Suite
}

func TestOneDriveUtilsSuite(t *testing.T) {
	suite.Run(t, &OneDriveUtilsSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDriveUtilsSuite) TestIncludeOneDriveRestoreDataSelectors() {
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
		opts             utils.OneDriveOpts
		expectIncludeLen int
	}{
		{
			name: "no inputs",
			opts: utils.OneDriveOpts{
				Users: empty,
				Paths: empty,
				Names: empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single inputs",
			opts: utils.OneDriveOpts{
				Users: single,
				Paths: single,
				Names: single,
			},
			expectIncludeLen: 1,
		},
		{
			name: "multi inputs",
			opts: utils.OneDriveOpts{
				Users: multi,
				Paths: multi,
				Names: multi,
			},
			expectIncludeLen: 1,
		},
		{
			name: "folder contains",
			opts: utils.OneDriveOpts{
				Users: empty,
				Paths: containsOnly,
				Names: empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "folder prefixes",
			opts: utils.OneDriveOpts{
				Users: empty,
				Paths: prefixOnly,
				Names: empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "folder prefixes and contains",
			opts: utils.OneDriveOpts{
				Users: empty,
				Paths: containsAndPrefix,
				Names: empty,
			},
			expectIncludeLen: 2,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			sel := utils.IncludeOneDriveRestoreDataSelectors(test.opts)
			assert.Len(suite.T(), sel.Includes, test.expectIncludeLen)
		})
	}
}
