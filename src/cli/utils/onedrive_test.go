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
				Users:      empty,
				FileName:   empty,
				FolderPath: empty,
			},
			expectIncludeLen: 1,
		},
		{
			name: "single inputs",
			opts: utils.OneDriveOpts{
				Users:      single,
				FileName:   single,
				FolderPath: single,
			},
			expectIncludeLen: 1,
		},
		{
			name: "multi inputs",
			opts: utils.OneDriveOpts{
				Users:      multi,
				FileName:   multi,
				FolderPath: multi,
			},
			expectIncludeLen: 1,
		},
		{
			name: "folder contains",
			opts: utils.OneDriveOpts{
				Users:      empty,
				FileName:   empty,
				FolderPath: containsOnly,
			},
			expectIncludeLen: 1,
		},
		{
			name: "folder prefixes",
			opts: utils.OneDriveOpts{
				Users:      empty,
				FileName:   empty,
				FolderPath: prefixOnly,
			},
			expectIncludeLen: 1,
		},
		{
			name: "folder prefixes and contains",
			opts: utils.OneDriveOpts{
				Users:      empty,
				FileName:   empty,
				FolderPath: containsAndPrefix,
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
