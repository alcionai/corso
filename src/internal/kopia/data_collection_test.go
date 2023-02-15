package kopia

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------
// unit tests
// ---------------
type KopiaDataCollectionUnitSuite struct {
	tester.Suite
}

func TestKopiaDataCollectionUnitSuite(t *testing.T) {
	s := &KopiaDataCollectionUnitSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

func (suite *KopiaDataCollectionUnitSuite) TestReturnsPath() {
	t := suite.T()

	b := path.Builder{}.Append("some", "path", "for", "data")
	pth, err := b.ToDataLayerExchangePathForCategory(
		"a-tenant",
		"a-user",
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	c := kopiaDataCollection{
		streams: []data.Stream{},
		path:    pth,
	}

	assert.Equal(t, pth, c.FullPath())
}

func (suite *KopiaDataCollectionUnitSuite) TestReturnsStreams() {
	testData := [][]byte{
		[]byte("abcdefghijklmnopqrstuvwxyz"),
		[]byte("zyxwvutsrqponmlkjihgfedcba"),
	}

	uuids := []string{
		"a-file",
		"another-file",
	}

	table := []struct {
		name    string
		streams []data.Stream
	}{
		{
			name: "SingleStream",
			streams: []data.Stream{
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[0])),
					uuid:   uuids[0],
					size:   int64(len(testData[0])),
				},
			},
		},
		{
			name: "MultipleStreams",
			streams: []data.Stream{
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[0])),
					uuid:   uuids[0],
					size:   int64(len(testData[0])),
				},
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[1])),
					uuid:   uuids[1],
					size:   int64(len(testData[1])),
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			c := kopiaDataCollection{
				streams: test.streams,
				path:    nil,
			}

			count := 0
			for returnedStream := range c.Items() {
				require.Less(t, count, len(test.streams))

				assert.Equal(t, returnedStream.UUID(), uuids[count])

				buf, err := io.ReadAll(returnedStream.ToReader())
				require.NoError(t, err)
				assert.Equal(t, buf, testData[count])
				require.Implements(t, (*data.StreamSize)(nil), returnedStream)
				ss := returnedStream.(data.StreamSize)
				assert.Equal(t, len(buf), int(ss.Size()))

				count++
			}

			assert.Equal(t, len(test.streams), count)
		})
	}
}

// These types are needed because we check that a fs.File was returned.
// Unfortunately fs.StreamingFile and fs.File have different interfaces so we
// have to fake things.
type mockSeeker struct{}

func (s mockSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, errors.New("not implemented")
}

type mockReader struct {
	io.ReadCloser
	mockSeeker
}

func (r mockReader) Entry() (fs.Entry, error) {
	return nil, errors.New("not implemented")
}

type mockFile struct {
	// Use for Entry interface.
	fs.StreamingFile
	r io.ReadCloser
}

func (f *mockFile) Open(ctx context.Context) (fs.Reader, error) {
	return mockReader{ReadCloser: f.r}, nil
}

func (suite *KopiaDataCollectionUnitSuite) TestFetch() {
	var (
		tenant   = "a-tenant"
		user     = "a-user"
		service  = path.ExchangeService.String()
		category = path.EmailCategory
		folder1  = "folder1"
		folder2  = "folder2"

		noErrFileName = "noError"
		errFileName   = "error"

		noErrFileData = "foo bar baz"

		errReader = &mockconnector.MockExchangeData{
			ReadErr: assert.AnError,
		}
	)

	// Needs to be a function so we can switch the serialization version as
	// needed.
	getLayout := func(serVersion uint32) fs.Entry {
		return virtualfs.NewStaticDirectory(encodeAsPath(tenant), []fs.Entry{
			virtualfs.NewStaticDirectory(encodeAsPath(service), []fs.Entry{
				virtualfs.NewStaticDirectory(encodeAsPath(user), []fs.Entry{
					virtualfs.NewStaticDirectory(encodeAsPath(category.String()), []fs.Entry{
						virtualfs.NewStaticDirectory(encodeAsPath(folder1), []fs.Entry{
							virtualfs.NewStaticDirectory(encodeAsPath(folder2), []fs.Entry{
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
							}),
						}),
					}),
				}),
			}),
		})
	}

	b := path.Builder{}.Append(folder1, folder2)
	pth, err := b.ToDataLayerExchangePathForCategory(
		tenant,
		user,
		category,
		false,
	)
	require.NoError(suite.T(), err)

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
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			root := getLayout(test.inputSerializationVersion)
			c := &i64counter{}

			col := &kopiaDataCollection{path: pth, snapshotRoot: root, counter: c}

			s, err := col.Fetch(ctx, test.inputName)

			test.lookupErr(t, err)

			if err != nil {
				if test.notFoundErr {
					assert.ErrorIs(t, err, data.ErrNotFound)
				}

				return
			}

			fileData, err := io.ReadAll(s.ToReader())

			test.readErr(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, test.expectedData, fileData)
		})
	}
}
