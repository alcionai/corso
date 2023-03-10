package onedrive

import (
	"testing"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RestoreUnitSuite struct {
	tester.Suite
}

func TestRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreUnitSuite) TestAugmentRestorePaths() {
	table := []struct {
		name    string
		version int
		input   []string
		output  []string
	}{
		{
			name:    "no change v0",
			version: 0,
			input: []string{
				"file.txt.data",
				"file.txt", // v0 does not have `.data`
			},
			output: []string{
				"file.txt", // ordering artifact of sorting
				"file.txt.data",
			},
		},
		{
			name:    "one folder v0",
			version: 0,
			input: []string{
				"folder/file.txt.data",
				"folder/file.txt",
			},
			output: []string{
				"folder/file.txt",
				"folder/file.txt.data",
			},
		},
		{
			name:    "no change v1",
			version: version.OneDrive1DataAndMetaFiles,
			input: []string{
				"file.txt.data",
			},
			output: []string{
				"file.txt.data",
			},
		},
		{
			name:    "one folder v1",
			version: version.OneDrive1DataAndMetaFiles,
			input: []string{
				"folder/file.txt.data",
			},
			output: []string{
				"folder.dirmeta",
				"folder/file.txt.data",
			},
		},
		{
			name:    "nested folders v1",
			version: version.OneDrive1DataAndMetaFiles,
			input: []string{
				"folder/file.txt.data",
				"folder/folder2/file.txt.data",
			},
			output: []string{
				"folder.dirmeta",
				"folder/file.txt.data",
				"folder/folder2.dirmeta",
				"folder/folder2/file.txt.data",
			},
		},
		{
			name:    "no change v4",
			version: version.OneDrive4DirIncludesPermissions,
			input: []string{
				"file.txt.data",
			},
			output: []string{
				"file.txt.data",
			},
		},
		{
			name:    "one folder v4",
			version: version.OneDrive4DirIncludesPermissions,
			input: []string{
				"folder/file.txt.data",
			},
			output: []string{
				"folder/file.txt.data",
				"folder/folder.dirmeta",
			},
		},
		{
			name:    "nested folders v4",
			version: version.OneDrive4DirIncludesPermissions,
			input: []string{
				"folder/file.txt.data",
				"folder/folder2/file.txt.data",
			},
			output: []string{
				"folder/file.txt.data",
				"folder/folder.dirmeta",
				"folder/folder2/file.txt.data",
				"folder/folder2/folder2.dirmeta",
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext()
			defer flush()

			base := "id/onedrive/user/files/drives/driveID/root:/"

			inPaths := []path.Path{}
			for _, ps := range test.input {
				p, err := path.FromDataLayerPath(base+ps, true)
				require.NoError(t, err, "creating path")

				inPaths = append(inPaths, p)
			}

			outPaths := []path.Path{}
			for _, ps := range test.output {
				p, err := path.FromDataLayerPath(base+ps, true)
				require.NoError(t, err, "creating path")

				outPaths = append(outPaths, p)
			}

			actual, err := AugmentRestorePaths(test.version, inPaths)
			require.NoError(t, err, "augmenting paths")

			// Ordering of paths matter here as we need dirmeta files
			// to show up before file in dir
			assert.Equal(t, outPaths, actual, "augmented paths")
		})
	}
}
