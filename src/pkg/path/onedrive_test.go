package path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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
			pathElements: []string{"drive", "driveID"},
			errCheck:     assert.Error,
		},
		{
			name:         "Root path",
			pathElements: []string{"drive", "driveID", "root:"},
			expected:     &path.DrivePath{DriveID: "driveID", Folders: []string{}},
			errCheck:     assert.NoError,
		},
		{
			name:         "Deeper path",
			pathElements: []string{"drive", "driveID", "root:", "folder1", "folder2"},
			expected:     &path.DrivePath{DriveID: "driveID", Folders: []string{"folder1", "folder2"}},
			errCheck:     assert.NoError,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()

			p, err := path.Build("tenant", "user", path.OneDriveService, path.FilesCategory, false, tt.pathElements...)
			require.NoError(suite.T(), err)

			got, err := path.ToOneDrivePath(p)
			tt.errCheck(t, err)
			if err != nil {
				return
			}
			assert.Equal(suite.T(), tt.expected, got)
		})
	}
}
