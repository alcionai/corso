package path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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
		suite.T().Run(tt.name, func(t *testing.T) {
			p, err := path.Builder{}.Append(tt.pathElements...).ToDataLayerOneDrivePath("tenant", "user", false)
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

func (suite *OneDrivePathSuite) TestOneDriveResourcePath_ShortRef() {
	path1 := "/tenant/" + path.OneDriveService.String() + "/user/" + path.FilesCategory.String() + "/drive-id/root:/item1"
	path2 := "/tenant/" + path.OneDriveService.String() + "/user/" + path.FilesCategory.String() + "/drive-id/root:/item2"
	itemName1 := "foo.txt"
	itemName2 := "bar.txt"

	table := []struct {
		name        string
		path1       string
		itemName1   string
		path2       string
		itemName2   string
		compareFunc assert.ComparisonAssertionFunc
	}{
		{
			name:        "SamePath_SameName",
			path1:       path1,
			itemName1:   itemName1,
			path2:       path1,
			itemName2:   itemName1,
			compareFunc: assert.Equal,
		},
		{
			name:        "DifferentPath_SameName",
			path1:       path1,
			itemName1:   itemName1,
			path2:       path2,
			itemName2:   itemName1,
			compareFunc: assert.NotEqual,
		},
		{
			name:        "SamePath_DifferentName",
			path1:       path1,
			itemName1:   itemName1,
			path2:       path1,
			itemName2:   itemName2,
			compareFunc: assert.NotEqual,
		},
		{
			name:        "DifferentPath_DifferentName",
			path1:       path1,
			itemName1:   itemName1,
			path2:       path2,
			itemName2:   itemName2,
			compareFunc: assert.NotEqual,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			p1, err := path.OneDriveResourcePath(test.itemName1, test.path1, true)
			require.NoError(t, err)

			p2, err := path.OneDriveResourcePath(test.itemName2, test.path2, true)
			require.NoError(t, err)

			test.compareFunc(t, p1.ShortRef(), p2.ShortRef())
		})
	}
}
