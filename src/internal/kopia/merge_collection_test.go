package kopia

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type MergeCollectionUnitSuite struct {
	tester.Suite
}

func TestMergeCollectionUnitSuite(t *testing.T) {
	suite.Run(t, &MergeCollectionUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MergeCollectionUnitSuite) TestReturnsPath() {
	t := suite.T()

	pth, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"some", "path", "for", "data")
	require.NoError(t, err, clues.ToCore(err))

	c := mergeCollection{
		fullPath: pth,
	}

	assert.Equal(t, pth, c.FullPath())
}

func (suite *MergeCollectionUnitSuite) TestItems() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	storagePaths := []string{
		"tenant-id/exchange/user-id/mail/some/folder/path1",
		"tenant-id/exchange/user-id/mail/some/folder/path2",
	}

	expectedItemNames := []string{"1", "2"}

	pth, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"some", "path", "for", "data")
	require.NoError(t, err, clues.ToCore(err))

	c1 := mock.NewCollection(pth, nil, 1)
	c1.Names[0] = expectedItemNames[0]

	c2 := mock.NewCollection(pth, nil, 1)
	c2.Names[0] = expectedItemNames[1]

	// Not testing fetch here so safe to use this wrapper.
	cols := []data.RestoreCollection{
		data.NoFetchRestoreCollection{Collection: c1},
		data.NoFetchRestoreCollection{Collection: c2},
	}

	dc := &mergeCollection{fullPath: pth}

	for i, c := range cols {
		err := dc.addCollection(storagePaths[i], c)
		require.NoError(t, err, "adding collection", clues.ToCore(err))
	}

	gotItemNames := []string{}

	for item := range dc.Items(ctx, fault.New(true)) {
		gotItemNames = append(gotItemNames, item.UUID())
	}

	assert.ElementsMatch(t, expectedItemNames, gotItemNames)
}

func (suite *MergeCollectionUnitSuite) TestAddCollection_DifferentPathFails() {
	t := suite.T()

	pth1, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"some", "path", "for", "data")
	require.NoError(t, err, clues.ToCore(err))

	pth2, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"some", "path", "for", "data2")
	require.NoError(t, err, clues.ToCore(err))

	dc := mergeCollection{fullPath: pth1}

	err = dc.addCollection("some/path", &kopiaDataCollection{path: pth2})
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *MergeCollectionUnitSuite) TestFetchItemByName() {
	var (
		fileData1 = []byte("abcdefghijklmnopqrstuvwxyz")
		fileData2 = []byte("zyxwvutsrqponmlkjihgfedcba")
		fileData3 = []byte("foo bar baz")

		fileName1         = "file1"
		fileName2         = "file2"
		fileLookupErrName = "errLookup"
		fileOpenErrName   = "errOpen"

		colPaths = []string{
			"tenant-id/exchange/user-id/mail/some/data/directory1",
			"tenant-id/exchange/user-id/mail/some/data/directory2",
		}
	)

	pth, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"some", "path", "for", "data")
	require.NoError(suite.T(), err, clues.ToCore(err))

	// Needs to be a function so the readers get refreshed each time.
	layouts := []func() fs.Directory{
		// Has the following;
		//   - file1: data[0]
		//   - errOpen: (error opening file)
		func() fs.Directory {
			return virtualfs.NewStaticDirectory(encodeAsPath(colPaths[0]), []fs.Entry{
				&mockFile{
					StreamingFile: virtualfs.StreamingFileFromReader(
						encodeAsPath(fileName1),
						nil,
					),
					r: newBackupStreamReader(
						serializationVersion,
						io.NopCloser(bytes.NewReader(fileData1)),
					),
					size: int64(len(fileData1) + versionSize),
				},
				&mockFile{
					StreamingFile: virtualfs.StreamingFileFromReader(
						encodeAsPath(fileOpenErrName),
						nil,
					),
					openErr: assert.AnError,
				},
			})
		},

		// Has the following;
		//   - file1: data[1]
		//   - file2: data[0]
		//   - errOpen: data[2]
		func() fs.Directory {
			return virtualfs.NewStaticDirectory(encodeAsPath(colPaths[1]), []fs.Entry{
				&mockFile{
					StreamingFile: virtualfs.StreamingFileFromReader(
						encodeAsPath(fileName1),
						nil,
					),
					r: newBackupStreamReader(
						serializationVersion,
						io.NopCloser(bytes.NewReader(fileData2)),
					),
					size: int64(len(fileData2) + versionSize),
				},
				&mockFile{
					StreamingFile: virtualfs.StreamingFileFromReader(
						encodeAsPath(fileName2),
						nil,
					),
					r: newBackupStreamReader(
						serializationVersion,
						io.NopCloser(bytes.NewReader(fileData1)),
					),
					size: int64(len(fileData1) + versionSize),
				},
				&mockFile{
					StreamingFile: virtualfs.StreamingFileFromReader(
						encodeAsPath(fileOpenErrName),
						nil,
					),
					r: newBackupStreamReader(
						serializationVersion,
						io.NopCloser(bytes.NewReader(fileData3)),
					),
					size: int64(len(fileData3) + versionSize),
				},
			})
		},
	}

	table := []struct {
		name        string
		fileName    string
		expectError assert.ErrorAssertionFunc
		expectData  []byte
		notFoundErr bool
	}{
		{
			name:        "Duplicate File, first collection",
			fileName:    fileName1,
			expectError: assert.NoError,
			expectData:  fileData1,
		},
		{
			name:        "Distinct File, second collection",
			fileName:    fileName2,
			expectError: assert.NoError,
			expectData:  fileData1,
		},
		{
			name:        "Error opening file",
			fileName:    fileOpenErrName,
			expectError: assert.Error,
		},
		{
			name:        "File not found",
			fileName:    fileLookupErrName,
			expectError: assert.Error,
			notFoundErr: true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			c := &i64counter{}

			dc := mergeCollection{fullPath: pth}

			for i, layout := range layouts {
				col := &kopiaDataCollection{
					path:            pth,
					dir:             layout(),
					counter:         c,
					expectedVersion: serializationVersion,
				}

				err := dc.addCollection(colPaths[i], col)
				require.NoError(t, err, "adding collection", clues.ToCore(err))
			}

			s, err := dc.FetchItemByName(ctx, test.fileName)
			test.expectError(t, err, clues.ToCore(err))

			if err != nil {
				if test.notFoundErr {
					assert.ErrorIs(t, err, data.ErrNotFound, clues.ToCore(err))
				}

				return
			}

			fileData, err := io.ReadAll(s.ToReader())
			require.NoError(t, err, "reading file data", clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expectData, fileData)
		})
	}
}
