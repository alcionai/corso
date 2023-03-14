package kopia

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
)

type PathEncoderSuite struct {
	tester.Suite
}

func TestPathEncoderSuite(t *testing.T) {
	suite.Run(t, &PathEncoderSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PathEncoderSuite) TestEncodeDecode() {
	t := suite.T()
	elements := []string{"these", "are", "some", "path", "elements"}
	encoded := encodeElements(elements...)
	decoded := make([]string, 0, len(elements))

	for _, e := range encoded {
		dec, err := decodeElement(e)
		require.NoError(t, err, clues.ToCore(err))

		decoded = append(decoded, dec)
	}

	assert.Equal(t, elements, decoded)
}

func (suite *PathEncoderSuite) TestEncodeAsPathDecode() {
	table := []struct {
		name     string
		elements []string
		expected []string
	}{
		{
			name:     "MultipleElements",
			elements: []string{"these", "are", "some", "path", "elements"},
			expected: []string{"these", "are", "some", "path", "elements"},
		},
		{
			name:     "SingleElement",
			elements: []string{"elements"},
			expected: []string{"elements"},
		},
		{
			name:     "EmptyPath",
			elements: []string{""},
			expected: []string{""},
		},
		{
			name:     "NilPath",
			elements: nil,
			// Gets "" back because individual elements are decoded and "" is the 0
			// value for the decoder.
			expected: []string{""},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			encoded := encodeAsPath(test.elements...)

			// Sanity check, first and last character should not be '/'.
			assert.Equal(t, strings.Trim(encoded, "/"), encoded)

			decoded := make([]string, 0, len(test.elements))

			for _, e := range strings.Split(encoded, "/") {
				dec, err := decodeElement(e)
				require.NoError(t, err, clues.ToCore(err))

				decoded = append(decoded, dec)
			}

			assert.Equal(t, test.expected, decoded)
		})
	}
}

func FuzzEncodeDecodeSingleString(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		encoded := encodeElements(in)
		assert.Len(t, encoded, 1)
		assert.False(t, strings.ContainsRune(encoded[0], '/'))

		decoded, err := decodeElement(encoded[0])
		require.NoError(t, err, clues.ToCore(err))
		assert.Equal(t, in, decoded)
	})
}
