package uploadsession

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UploadSessionSuite struct {
	suite.Suite
}

func TestUploadSessionSuite(t *testing.T) {
	suite.Run(t, new(UploadSessionSuite))
}

func (suite *UploadSessionSuite) TestWriter() {
	t := suite.T()

	// Initialize a 100KB mockDataProvider
	td, writeSize := mockDataReader(int64(100 * 1024))

	// Expected Content-Range value format
	contentRangeRegex := regexp.MustCompile(`^bytes \d+-\d+/\d+$`)

	// Initialize a test http server that validates expeected headers
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPut)
		// Validate the "Content-Length" header
		assert.Equal(t, r.Header[contentLengthHeaderKey][0], fmt.Sprintf("%d", writeSize))
		// Validate the "Content-Range" header
		assert.True(t, contentRangeRegex.MatchString(r.Header[contentRangeHeaderKey][0]),
			"%s does not match expected value", r.Header[contentRangeHeaderKey][0])
	}))
	defer ts.Close()

	writer := NewWriter("item", ts.URL, writeSize)

	// Using a 32 KB buffer for the copy allows us to validate the
	// multi-part upload. `io.CopyBuffer` will only write 32 KB at
	// a time
	copyBuffer := make([]byte, 32*1024)

	size, err := io.CopyBuffer(writer, td, copyBuffer)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), writeSize, size)
}

func mockDataReader(size int64) (io.Reader, int64) {
	data := bytes.Repeat([]byte("D"), int(size))
	return bytes.NewReader(data), size
}
