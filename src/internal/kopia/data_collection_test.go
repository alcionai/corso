package kopia

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------
// Wrappers to match required interfaces.
// ---------------

// These types are needed because we check that a fs.File was returned.
// Unfortunately fs.StreamingFile and fs.File have different interfaces so we
// have to fake things.
type mockSeeker struct{}

func (s mockSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, clues.New("not implemented")
}

type mockReader struct {
	io.ReadCloser
	mockSeeker
}

func (r mockReader) Entry() (fs.Entry, error) {
	return nil, clues.New("not implemented")
}

type mockFile struct {
	// Use for Entry interface.
	fs.StreamingFile
	r       io.ReadCloser
	openErr error
	size    int64
}

func (f *mockFile) Open(ctx context.Context) (fs.Reader, error) {
	if f.openErr != nil {
		return nil, f.openErr
	}

	return mockReader{ReadCloser: f.r}, nil
}

func (f *mockFile) Size() int64 {
	return f.size
}

// ---------------
// unit tests
// ---------------
type KopiaDataCollectionUnitSuite struct {
	tester.Suite
}

func TestKopiaDataCollectionUnitSuite(t *testing.T) {
	suite.Run(t, &KopiaDataCollectionUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *KopiaDataCollectionUnitSuite) TestReturnsPath() {
	t := suite.T()

	pth, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"some", "path", "for", "data")
	require.NoError(t, err, clues.ToCore(err))

	c := kopiaDataCollection{
		path: pth,
	}

	assert.Equal(t, pth, c.FullPath())
}

func (suite *KopiaDataCollectionUnitSuite) TestReturnsStreams() {
	type loadedData struct {
		uuid string
		data []byte
		size int64
	}

	var (
		fileData = [][]byte{
			[]byte("abcdefghijklmnopqrstuvwxyz"),
			[]byte("zyxwvutsrqponmlkjihgfedcba"),
		}

		uuids = []string{
			"a-file",
			"another-file",
		}

		files = []loadedData{
			{uuid: uuids[0], data: fileData[0], size: int64(len(fileData[0]))},
			{uuid: uuids[1], data: fileData[1], size: int64(len(fileData[1]))},
		}

		fileLookupErrName = "errLookup"
		fileOpenErrName   = "errOpen"
		notFileErrName    = "errNotFile"
	)

	// Needs to be a function so the readers get refreshed each time.
	getLayout := func() fs.Directory {
		return virtualfs.NewStaticDirectory(encodeAsPath("foo"), []fs.Entry{
			&mockFile{
				StreamingFile: virtualfs.StreamingFileFromReader(
					encodeAsPath(files[0].uuid),
					nil,
				),
				r: newBackupStreamReader(
					serializationVersion,
					io.NopCloser(bytes.NewReader(files[0].data)),
				),
				size: int64(len(files[0].data) + versionSize),
			},
			&mockFile{
				StreamingFile: virtualfs.StreamingFileFromReader(
					encodeAsPath(files[1].uuid),
					nil,
				),
				r: newBackupStreamReader(
					serializationVersion,
					io.NopCloser(bytes.NewReader(files[1].data)),
				),
				size: int64(len(files[1].data) + versionSize),
			},
			&mockFile{
				StreamingFile: virtualfs.StreamingFileFromReader(
					encodeAsPath(fileOpenErrName),
					nil,
				),
				openErr: assert.AnError,
			},
			virtualfs.NewStaticDirectory(encodeAsPath(notFileErrName), []fs.Entry{}),
		})
	}

	table := []struct {
		name           string
		uuidsAndErrors map[string]assert.ErrorAssertionFunc
		// Data and stuff about the loaded data.
		expectedLoaded []loadedData
	}{
		{
			name: "SingleStream",
			uuidsAndErrors: map[string]assert.ErrorAssertionFunc{
				uuids[0]: nil,
			},
			expectedLoaded: []loadedData{files[0]},
		},
		{
			name: "MultipleStreams",
			uuidsAndErrors: map[string]assert.ErrorAssertionFunc{
				uuids[0]: nil,
				uuids[1]: nil,
			},
			expectedLoaded: files,
		},
		{
			name: "Some Not Found Errors",
			uuidsAndErrors: map[string]assert.ErrorAssertionFunc{
				fileLookupErrName: assert.Error,
				uuids[0]:          nil,
			},
			expectedLoaded: []loadedData{files[0]},
		},
		{
			name: "Some Not A File Errors",
			uuidsAndErrors: map[string]assert.ErrorAssertionFunc{
				notFileErrName: assert.Error,
				uuids[0]:       nil,
			},
			expectedLoaded: []loadedData{files[0]},
		},
		{
			name: "Some Open Errors",
			uuidsAndErrors: map[string]assert.ErrorAssertionFunc{
				fileOpenErrName: assert.Error,
				uuids[0]:        nil,
			},
			expectedLoaded: []loadedData{files[0]},
		},
		{
			name: "Empty Name Errors",
			uuidsAndErrors: map[string]assert.ErrorAssertionFunc{
				"": assert.Error,
			},
			expectedLoaded: []loadedData{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			items := []string{}
			errs := []assert.ErrorAssertionFunc{}

			for uuid, err := range test.uuidsAndErrors {
				if err != nil {
					errs = append(errs, err)
				}

				items = append(items, uuid)
			}

			c := kopiaDataCollection{
				dir:             getLayout(),
				path:            nil,
				items:           items,
				expectedVersion: serializationVersion,
			}

			var (
				found []loadedData
				bus   = fault.New(false)
			)

			for item := range c.Items(ctx, bus) {
				require.Less(t, len(found), len(test.expectedLoaded), "items read safety")

				found = append(found, loadedData{})
				f := &found[len(found)-1]
				f.uuid = item.ID()

				buf, err := io.ReadAll(item.ToReader())
				if !assert.NoError(t, err, clues.ToCore(err)) {
					continue
				}

				f.data = buf

				if !assert.Implements(t, (*data.ItemSize)(nil), item) {
					continue
				}

				ss := item.(data.ItemSize)

				f.size = ss.Size()
			}

			// We expect the items to be fetched in the order they are
			// in the struct or the errors will not line up
			for i, err := range bus.Recovered() {
				assert.True(t, errs[i](t, err), "expected error", clues.ToCore(err))
			}

			assert.NoError(t, bus.Failure(), "expected no hard failures")

			assert.ElementsMatch(t, test.expectedLoaded, found, "loaded items")
		})
	}
}

func (suite *KopiaDataCollectionUnitSuite) TestFetchItemByName() {
	var (
		tenant   = "a-tenant"
		user     = "a-user"
		category = path.EmailCategory
		folder1  = "folder1"
		folder2  = "folder2"

		noErrFileName = "noError"
		errFileName   = "error"
		errFileName2  = "error2"

		noErrFileData = "foo bar baz"
		errReader     = &dataMock.Item{
			ReadErr: assert.AnError,
		}
	)

	// Needs to be a function so we can switch the serialization version as
	// needed.
	getLayout := func(serVersion uint32) fs.Directory {
		return virtualfs.NewStaticDirectory(encodeAsPath(folder2), []fs.Entry{
			&mockFile{
				StreamingFile: virtualfs.StreamingFileFromReader(
					encodeAsPath(noErrFileName),
					nil,
				),
				r: newBackupStreamReader(
					serVersion,
					io.NopCloser(bytes.NewReader([]byte(noErrFileData))),
				),
			},
			&mockFile{
				StreamingFile: virtualfs.StreamingFileFromReader(
					encodeAsPath(errFileName),
					nil,
				),
				r: newBackupStreamReader(
					serVersion,
					errReader.ToReader(),
				),
			},
			&mockFile{
				StreamingFile: virtualfs.StreamingFileFromReader(
					encodeAsPath(errFileName2),
					nil,
				),
				openErr: assert.AnError,
			},
		})
	}

	pth, err := path.Build(
		tenant,
		user,
		path.ExchangeService,
		category,
		false,
		folder1, folder2)
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name                      string
		inputName                 string
		inputSerializationVersion uint32
		expectedData              []byte
		lookupErr                 assert.ErrorAssertionFunc
		readErr                   assert.ErrorAssertionFunc
		notFoundErr               bool
	}{
		{
			name:                      "FileFound_NoError",
			inputName:                 noErrFileName,
			inputSerializationVersion: serializationVersion,
			expectedData:              []byte(noErrFileData),
			lookupErr:                 assert.NoError,
			readErr:                   assert.NoError,
		},
		{
			name:                      "FileFound_ReadError",
			inputName:                 errFileName,
			inputSerializationVersion: serializationVersion,
			lookupErr:                 assert.NoError,
			readErr:                   assert.Error,
		},
		{
			name:                      "FileFound_VersionError",
			inputName:                 noErrFileName,
			inputSerializationVersion: serializationVersion + 1,
			lookupErr:                 assert.NoError,
			readErr:                   assert.Error,
		},
		{
			name:                      "FileNotFound",
			inputName:                 "foo",
			inputSerializationVersion: serializationVersion + 1,
			lookupErr:                 assert.Error,
			notFoundErr:               true,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			root := getLayout(test.inputSerializationVersion)
			c := &i64counter{}

			col := &kopiaDataCollection{
				path:            pth,
				dir:             root,
				counter:         c,
				expectedVersion: serializationVersion,
			}

			s, err := col.FetchItemByName(ctx, test.inputName)

			test.lookupErr(t, err)

			if err != nil {
				if test.notFoundErr {
					assert.ErrorIs(t, err, data.ErrNotFound, clues.ToCore(err))
				}

				return
			}

			fileData, err := io.ReadAll(s.ToReader())
			test.readErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expectedData, fileData)
		})
	}
}
