package uploadsession

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
)

type UploadSessionSuite struct {
	tester.Suite
}

func TestUploadSessionSuite(t *testing.T) {
	suite.Run(t, &UploadSessionSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *UploadSessionSuite) TestWriter() {
	t := suite.T()

	// Initialize a 100KB mockDataProvider
	td, writeSize := mockDataReader(int64(100 * 1024))

	// Expected Content-Range value format
	contentRangeRegex := regexp.MustCompile(`^bytes (?P<rangestart>\d+)-(?P<rangeend>\d+)/(?P<length>\d+)$`)
	nextOffset := -1

	// Initialize a test http server that validates expeected headers
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPut)

		// Validate the "Content-Range" header
		assert.True(t, contentRangeRegex.MatchString(r.Header[contentRangeHeaderKey][0]),
			"%s does not match expected value", r.Header[contentRangeHeaderKey][0])

		// Extract the Content-Range components
		matches := contentRangeRegex.FindStringSubmatch(r.Header[contentRangeHeaderKey][0])

		rangeStart, err := strconv.Atoi(matches[contentRangeRegex.SubexpIndex("rangestart")])
		assert.NoError(t, err, clues.ToCore(err))

		rangeEnd, err := strconv.Atoi(matches[contentRangeRegex.SubexpIndex("rangeend")])
		assert.NoError(t, err, clues.ToCore(err))

		length, err := strconv.Atoi(matches[contentRangeRegex.SubexpIndex("length")])
		assert.NoError(t, err, clues.ToCore(err))

		// Validate total size and range start/end
		assert.Equal(t, int(writeSize), length)
		assert.Equal(t, nextOffset+1, rangeStart)
		assert.Greater(t, rangeEnd, nextOffset)

		// Validate the "Content-Length" header
		assert.Equal(t, fmt.Sprintf("%d", (rangeEnd+1)-rangeStart), r.Header[contentLengthHeaderKey][0])

		nextOffset = rangeEnd
	}))

	defer ts.Close()

	writer := NewWriter("item", ts.URL, writeSize)

	// Using a 32 KB buffer for the copy allows us to validate the
	// multi-part upload. `io.CopyBuffer` will only write 32 KB at
	// a time
	copyBuffer := make([]byte, 32*1024)

	size, err := io.CopyBuffer(writer, td, copyBuffer)
	require.NoError(suite.T(), err, clues.ToCore(err))
	require.Equal(suite.T(), writeSize, size)
}

func mockDataReader(size int64) (io.Reader, int64) {
	data := bytes.Repeat([]byte("D"), int(size))
	return &mockReader{r: bytes.NewReader(data)}, size
}

// mockReader allows us to wrap a `bytes.NewReader` but *disable*
// ReaderFrom functionality. This forces io.CopyBuffer to do a
// buffered read (useful to validate that chunked writes are working)
type mockReader struct {
	r io.Reader
}

func (mr *mockReader) Read(b []byte) (n int, err error) {
	return mr.r.Read(b)
}
