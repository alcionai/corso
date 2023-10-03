package readers_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/tester"
)

type shortReader struct {
	maxReadLen int
	io.ReadCloser
}

func (s *shortReader) Read(p []byte) (int, error) {
	toRead := s.maxReadLen
	if len(p) < toRead {
		toRead = len(p)
	}

	return s.ReadCloser.Read(p[:toRead])
}

type SerializationReaderUnitSuite struct {
	tester.Suite
}

func TestSerializationReaderUnitSuite(t *testing.T) {
	suite.Run(t, &SerializationReaderUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SerializationReaderUnitSuite) TestBackupSerializationReader() {
	baseData := []byte("hello world")

	table := []struct {
		name         string
		format       readers.SerializationFormat
		inputReaders []io.ReadCloser

		expectErr  require.ErrorAssertionFunc
		expectData []byte
	}{
		{
			name: "DeletedInFlight NoVersion NoReaders",
			format: readers.SerializationFormat{
				DelInFlight: true,
			},
			expectErr:  require.NoError,
			expectData: []byte{0x80, 0x0, 0x0, 0x0},
		},
		{
			name: "DeletedInFlight NoReaders",
			format: readers.SerializationFormat{
				Version:     42,
				DelInFlight: true,
			},
			expectErr:  require.NoError,
			expectData: []byte{0x80, 0x0, 0x0, 42},
		},
		{
			name:       "NoVersion NoReaders",
			expectErr:  require.NoError,
			expectData: []byte{0x00, 0x0, 0x0, 0x0},
		},
		{
			name: "NoReaders",
			format: readers.SerializationFormat{
				Version: 42,
			},
			expectErr:  require.NoError,
			expectData: []byte{0x00, 0x0, 0x0, 42},
		},
		{
			name: "SingleReader",
			format: readers.SerializationFormat{
				Version: 42,
			},
			inputReaders: []io.ReadCloser{io.NopCloser(bytes.NewReader(baseData))},
			expectErr:    require.NoError,
			expectData:   append([]byte{0x00, 0x0, 0x0, 42}, baseData...),
		},
		{
			name: "MultipleReaders",
			format: readers.SerializationFormat{
				Version: 42,
			},
			inputReaders: []io.ReadCloser{
				io.NopCloser(bytes.NewReader(baseData)),
				io.NopCloser(bytes.NewReader(baseData)),
			},
			expectErr: require.NoError,
			expectData: append(
				append([]byte{0x00, 0x0, 0x0, 42}, baseData...),
				baseData...),
		},
		// Uncomment if we expand the version to 32 bits.
		//{
		//	name: "VersionWithHighBitSet NoReaders Errors",
		//	format: readers.SerializationFormat{
		//		Version: 0x80000000,
		//	},
		//	expectErr: require.Error,
		//},
		{
			name: "DeletedInFlight SingleReader Errors",
			format: readers.SerializationFormat{
				DelInFlight: true,
			},
			inputReaders: []io.ReadCloser{io.NopCloser(bytes.NewReader(baseData))},
			expectErr:    require.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			r, err := readers.NewVersionedBackupReader(
				test.format,
				test.inputReaders...)
			test.expectErr(t, err, "getting backup reader: %v", clues.ToCore(err))

			if err != nil {
				return
			}

			defer func() {
				err := r.Close()
				assert.NoError(t, err, "closing reader: %v", clues.ToCore(err))
			}()

			buf, err := io.ReadAll(r)
			require.NoError(
				t,
				err,
				"reading serialized data: %v",
				clues.ToCore(err))

			// Need to use equal because output is order-sensitive.
			assert.Equal(t, test.expectData, buf, "serialized data")
		})
	}
}

func (suite *SerializationReaderUnitSuite) TestBackupSerializationReader_ShortReads() {
	t := suite.T()

	baseData := []byte("hello world")
	expectData := append(
		append([]byte{0x00, 0x0, 0x0, 42}, baseData...),
		baseData...)

	r, err := readers.NewVersionedBackupReader(
		readers.SerializationFormat{Version: 42},
		io.NopCloser(bytes.NewReader(baseData)),
		io.NopCloser(bytes.NewReader(baseData)))
	require.NoError(t, err, "getting backup reader: %v", clues.ToCore(err))

	defer func() {
		err := r.Close()
		assert.NoError(t, err, "closing reader: %v", clues.ToCore(err))
	}()

	buf := make([]byte, len(expectData))
	r = &shortReader{
		maxReadLen: 3,
		ReadCloser: r,
	}

	for read := 0; ; {
		n, err := r.Read(buf[read:])

		read += n
		if read >= len(buf) {
			break
		}

		require.NoError(t, err, "reading data: %v", clues.ToCore(err))
	}

	// Need to use equal because output is order-sensitive.
	assert.Equal(t, expectData, buf, "serialized data")
}

// TestRestoreSerializationReader checks that we can read previously serialized
// data. For simplicity, it uses the versionedBackupReader to generate the
// input. This should be relatively safe because the tests for
// versionedBackupReader do compare directly against serialized data.
func (suite *SerializationReaderUnitSuite) TestRestoreSerializationReader() {
	baseData := []byte("hello world")

	table := []struct {
		name        string
		inputReader func(*testing.T) io.ReadCloser

		expectErr         require.ErrorAssertionFunc
		expectVersion     readers.SerializationVersion
		expectDelInFlight bool
		expectData        []byte
	}{
		{
			name: "NoVersion NoReaders",
			inputReader: func(t *testing.T) io.ReadCloser {
				r, err := readers.NewVersionedBackupReader(readers.SerializationFormat{})
				require.NoError(t, err, "making reader: %v", clues.ToCore(err))

				return r
			},
			expectErr:  require.NoError,
			expectData: []byte{},
		},
		{
			name: "DeletedInFlight NoReaders",
			inputReader: func(t *testing.T) io.ReadCloser {
				r, err := readers.NewVersionedBackupReader(
					readers.SerializationFormat{
						Version:     42,
						DelInFlight: true,
					})
				require.NoError(t, err, "making reader: %v", clues.ToCore(err))

				return r
			},
			expectErr:         require.NoError,
			expectVersion:     42,
			expectDelInFlight: true,
			expectData:        []byte{},
		},
		{
			name: "DeletedInFlight SingleReader",
			inputReader: func(t *testing.T) io.ReadCloser {
				// Need to specify the bytes manually because the backup reader won't
				// allow creating something with the deleted flag and data.
				return io.NopCloser(bytes.NewReader(append(
					[]byte{0x80, 0x0, 0x0, 42},
					baseData...)))
			},
			expectErr:         require.NoError,
			expectVersion:     42,
			expectDelInFlight: true,
			expectData:        baseData,
		},
		{
			name: "NoVersion SingleReader",
			inputReader: func(t *testing.T) io.ReadCloser {
				r, err := readers.NewVersionedBackupReader(
					readers.SerializationFormat{},
					io.NopCloser(bytes.NewReader(baseData)))
				require.NoError(t, err, "making reader: %v", clues.ToCore(err))

				return r
			},
			expectErr:  require.NoError,
			expectData: baseData,
		},
		{
			name: "SingleReader",
			inputReader: func(t *testing.T) io.ReadCloser {
				r, err := readers.NewVersionedBackupReader(
					readers.SerializationFormat{Version: 42},
					io.NopCloser(bytes.NewReader(baseData)))
				require.NoError(t, err, "making reader: %v", clues.ToCore(err))

				return r
			},
			expectErr:     require.NoError,
			expectVersion: 42,
			expectData:    baseData,
		},
		{
			name: "ShortReads SingleReader",
			inputReader: func(t *testing.T) io.ReadCloser {
				r, err := readers.NewVersionedBackupReader(
					readers.SerializationFormat{Version: 42},
					io.NopCloser(bytes.NewReader(baseData)))
				require.NoError(t, err, "making reader: %v", clues.ToCore(err))

				r = &shortReader{
					maxReadLen: 3,
					ReadCloser: r,
				}

				return r
			},
			expectErr:     require.NoError,
			expectVersion: 42,
			expectData:    baseData,
		},
		{
			name: "MultipleReaders",
			inputReader: func(t *testing.T) io.ReadCloser {
				r, err := readers.NewVersionedBackupReader(
					readers.SerializationFormat{Version: 42},
					io.NopCloser(bytes.NewReader(baseData)),
					io.NopCloser(bytes.NewReader(baseData)))
				require.NoError(t, err, "making reader: %v", clues.ToCore(err))

				return r
			},
			expectErr:     require.NoError,
			expectVersion: 42,
			expectData:    append(slices.Clone(baseData), baseData...),
		},
		{
			name: "EmptyReader Errors",
			inputReader: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(bytes.NewReader([]byte{}))
			},
			expectErr: require.Error,
		},
		{
			name: "TruncatedVersion Errors",
			inputReader: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(bytes.NewReader([]byte{0x80, 0x0}))
			},
			expectErr: require.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			r, err := readers.NewVersionedRestoreReader(test.inputReader(t))
			test.expectErr(t, err, "getting restore reader: %v", clues.ToCore(err))

			if err != nil {
				return
			}

			defer func() {
				err := r.Close()
				assert.NoError(t, err, "closing reader: %v", clues.ToCore(err))
			}()

			assert.Equal(
				t,
				test.expectVersion,
				r.Format().Version,
				"version")
			assert.Equal(
				t,
				test.expectDelInFlight,
				r.Format().DelInFlight,
				"deleted in flight")

			buf, err := io.ReadAll(r)
			require.NoError(t, err, "reading serialized data: %v", clues.ToCore(err))

			// Need to use equal because output is order-sensitive.
			assert.Equal(t, test.expectData, buf, "serialized data")
		})
	}
}
