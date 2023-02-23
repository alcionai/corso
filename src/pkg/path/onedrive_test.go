package path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester/aw"
	"github.com/alcionai/corso/src/pkg/path"
)

type OneDrivePathSuite struct {
	suite.Suite
}

func TestOneDrivePathSuite(t *testing.T) {
	suite.Run(t, new(OneDrivePathSuite))
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
			errCheck:     aw.Err,
		},
		{
			name:         "Root path",
			pathElements: []string{"drive", "driveID", "root:"},
			expected:     &path.DrivePath{DriveID: "driveID", Folders: []string{}},
			errCheck:     aw.NoErr,
		},
		{
			name:         "Deeper path",
			pathElements: []string{"drive", "driveID", "root:", "folder1", "folder2"},
			expected:     &path.DrivePath{DriveID: "driveID", Folders: []string{"folder1", "folder2"}},
			errCheck:     aw.NoErr,
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			p, err := path.Builder{}.Append(tt.pathElements...).ToDataLayerOneDrivePath("tenant", "user", false)
			aw.MustNoErr(suite.T(), err)

			got, err := path.ToOneDrivePath(p)
			tt.errCheck(t, err)
			if err != nil {
				return
			}
			assert.Equal(suite.T(), tt.expected, got)
		})
	}
}
