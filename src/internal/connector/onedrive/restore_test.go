package onedrive

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
)

type RestoreUnitSuite struct {
	tester.Suite
}

func TestRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreUnitSuite) TestAugmentRestorePaths() {
	// Adding a simple test here so that we can be sure that this
	// function gets updated whenever we add a new version.
	require.LessOrEqual(suite.T(), version.Backup, version.All8MigrateUserPNToID, "unsupported backup version")

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
		{
			name:    "no change v6",
			version: version.OneDrive6NameInMeta,
			input: []string{
				"file.txt.data",
			},
			output: []string{
				"file.txt.data",
			},
		},
		{
			name:    "one folder v6",
			version: version.OneDrive6NameInMeta,
			input: []string{
				"folder/file.txt.data",
			},
			output: []string{
				"folder/.dirmeta",
				"folder/file.txt.data",
			},
		},
		{
			name:    "nested folders v6",
			version: version.OneDrive6NameInMeta,
			input: []string{
				"folder/file.txt.data",
				"folder/folder2/file.txt.data",
			},
			output: []string{
				"folder/.dirmeta",
				"folder/file.txt.data",
				"folder/folder2/.dirmeta",
				"folder/folder2/file.txt.data",
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			base := "id/onedrive/user/files/drives/driveID/root:/"

			inPaths := []path.RestorePaths{}
			for _, ps := range test.input {
				p, err := path.FromDataLayerPath(base+ps, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				pd, err := p.Dir()
				require.NoError(t, err, "creating collection path", clues.ToCore(err))

				inPaths = append(
					inPaths,
					path.RestorePaths{StoragePath: p, RestorePath: pd})
			}

			outPaths := []path.RestorePaths{}
			for _, ps := range test.output {
				p, err := path.FromDataLayerPath(base+ps, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				pd, err := p.Dir()
				require.NoError(t, err, "creating collection path", clues.ToCore(err))

				outPaths = append(
					outPaths,
					path.RestorePaths{StoragePath: p, RestorePath: pd})
			}

			actual, err := AugmentRestorePaths(test.version, inPaths)
			require.NoError(t, err, "augmenting paths", clues.ToCore(err))

			// Ordering of paths matter here as we need dirmeta files
			// to show up before file in dir
			assert.Equal(t, outPaths, actual, "augmented paths")
		})
	}
}

// TestAugmentRestorePaths_DifferentRestorePath tests that RestorePath
// substitution works properly. Since it's only possible for future backup
// versions to need restore path substitution (i.e. due to storing folders by
// ID instead of name) this is only tested against the most recent backup
// version at the moment.
func (suite *RestoreUnitSuite) TestAugmentRestorePaths_DifferentRestorePath() {
	// Adding a simple test here so that we can be sure that this
	// function gets updated whenever we add a new version.
	require.LessOrEqual(suite.T(), version.Backup, version.All8MigrateUserPNToID, "unsupported backup version")

	type pathPair struct {
		storage string
		restore string
	}

	table := []struct {
		name     string
		version  int
		input    []pathPair
		output   []pathPair
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name:    "nested folders",
			version: version.Backup,
			input: []pathPair{
				{storage: "folder-id/file.txt.data", restore: "folder"},
				{storage: "folder-id/folder2-id/file.txt.data", restore: "folder/folder2"},
			},
			output: []pathPair{
				{storage: "folder-id/.dirmeta", restore: "folder"},
				{storage: "folder-id/file.txt.data", restore: "folder"},
				{storage: "folder-id/folder2-id/.dirmeta", restore: "folder/folder2"},
				{storage: "folder-id/folder2-id/file.txt.data", restore: "folder/folder2"},
			},
			errCheck: assert.NoError,
		},
		{
			name:    "restore path longer one folder",
			version: version.Backup,
			input: []pathPair{
				{storage: "folder-id/file.txt.data", restore: "corso_restore/folder"},
			},
			output: []pathPair{
				{storage: "folder-id/.dirmeta", restore: "corso_restore/folder"},
				{storage: "folder-id/file.txt.data", restore: "corso_restore/folder"},
			},
			errCheck: assert.NoError,
		},
		{
			name:    "restore path shorter one folder",
			version: version.Backup,
			input: []pathPair{
				{storage: "folder-id/file.txt.data", restore: ""},
			},
			errCheck: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			base := "id/onedrive/user/files/drives/driveID/root:/"

			inPaths := []path.RestorePaths{}
			for _, ps := range test.input {
				p, err := path.FromDataLayerPath(base+ps.storage, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				r, err := path.FromDataLayerPath(base+ps.restore, false)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				inPaths = append(
					inPaths,
					path.RestorePaths{StoragePath: p, RestorePath: r})
			}

			outPaths := []path.RestorePaths{}
			for _, ps := range test.output {
				p, err := path.FromDataLayerPath(base+ps.storage, true)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				r, err := path.FromDataLayerPath(base+ps.restore, false)
				require.NoError(t, err, "creating path", clues.ToCore(err))

				outPaths = append(
					outPaths,
					path.RestorePaths{StoragePath: p, RestorePath: r})
			}

			actual, err := AugmentRestorePaths(test.version, inPaths)
			test.errCheck(t, err, "augmenting paths", clues.ToCore(err))

			if err != nil {
				return
			}

			// Ordering of paths matter here as we need dirmeta files
			// to show up before file in dir
			assert.Equal(t, outPaths, actual, "augmented paths")
		})
	}
}

func TestCopyBufferWithStallCheck_Valid(t *testing.T) {
	data := []byte("hello world")
	reader := bytes.NewReader(data)
	writer := bytes.Buffer{}

	copied, err := copyBufferWithStallCheck(&writer, reader, time.Second)
	require.NoError(t, err)
	assert.Equal(t, int64(len(data)), copied)
	assert.Equal(t, data, writer.Bytes()[:len(data)])
}

func TestCopyBufferWithStallCheck_Error(t *testing.T) {
	// Reader that always returns an error
	reader := &errorReader{}
	writer := bytes.Buffer{}

	copied, err := copyBufferWithStallCheck(&writer, reader, time.Second)
	assert.EqualError(t, err, "reading data: test error")
	assert.Equal(t, int64(0), copied)
}

type errorReader struct{}

func (r *errorReader) Read(p []byte) (int, error) {
	return 0, errors.New("test error")
}

func TestCopyBufferWithStallCheck_ReadStalls(t *testing.T) {
	// Reader that never returns any data
	reader := &stallReader{}
	writer := bytes.Buffer{}

	copied, err := copyBufferWithStallCheck(&writer, reader, time.Second)
	assert.EqualError(t, err, "copy stalled")
	assert.Equal(t, int64(0), copied)
}

type stallReader struct{}

func (r *stallReader) Read(p []byte) (int, error) {
	time.Sleep(time.Second * 2)

	p[0] = 'a' // just something to ensure test failure

	return 0, nil
}

func TestCopyBufferWithStallCheck_WriteStalls(t *testing.T) {
	// Writer that never returns any data
	reader := bytes.NewReader([]byte("hello world"))
	writer := &stallWriter{}

	copied, err := copyBufferWithStallCheck(writer, reader, time.Second)
	assert.EqualError(t, err, "copy stalled")
	assert.Equal(t, int64(0), copied)
}

type stallWriter struct{}

func (w *stallWriter) Write(p []byte) (int, error) {
	time.Sleep(time.Second * 2)
	return 10, nil
}
