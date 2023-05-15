package path_test

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type OneDrivePathSuite struct {
	tester.Suite
}

func TestOneDrivePathSuite(t *testing.T) {
	suite.Run(t, &OneDrivePathSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDrivePathSuite) Test_ToOneDrivePath() {
	tests := []struct {
		name         string
		pathElements []string
		expected     *path.DrivePath
		errCheck     assert.ErrorAssertionFunc
	}{
		{
			name:         "Not enough path elements",
			pathElements: []string{odConsts.DrivesPathDir, "driveID"},
			errCheck:     assert.Error,
		},
		{
			name:         "Root path",
			pathElements: []string{odConsts.DrivesPathDir, "driveID", odConsts.RootPathDir},
			expected: &path.DrivePath{
				DriveID: "driveID",
				Root:    odConsts.RootPathDir,
				Folders: []string{},
			},
			errCheck: assert.NoError,
		},
		{
			name:         "Deeper path",
			pathElements: []string{odConsts.DrivesPathDir, "driveID", odConsts.RootPathDir, "folder1", "folder2"},
			expected: &path.DrivePath{
				DriveID: "driveID",
				Root:    odConsts.RootPathDir,
				Folders: []string{"folder1", "folder2"},
			},
			errCheck: assert.NoError,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()

			p, err := path.Build("tenant", "user", path.OneDriveService, path.FilesCategory, false, tt.pathElements...)
			require.NoError(suite.T(), err, clues.ToCore(err))

			got, err := path.ToDrivePath(p)
			tt.errCheck(t, err)
			if err != nil {
				return
			}
			assert.Equal(suite.T(), tt.expected, got)
		})
	}
}

func (suite *OneDrivePathSuite) TestFormatDriveFolders() {
	const (
		driveID     = "some-drive-id"
		drivePrefix = "drives/" + driveID
	)

	table := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name: "normal",
			input: []string{
				"root:",
				"foo",
				"bar",
			},
			expected: strings.Join(
				append([]string{drivePrefix}, "root:", "foo", "bar"),
				"/"),
		},
		{
			name: "has character that would be escaped",
			input: []string{
				"root:",
				"foo/",
				"bar",
			},
			// Element "foo/" should end up escaped in the string output.
			expected: strings.Join(
				append([]string{drivePrefix}, "root:", `foo\/`, "bar"),
				"/"),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			assert.Equal(
				suite.T(),
				test.expected,
				path.BuildDriveLocation(driveID, test.input...).String())
		})
	}
}
