package kopia

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PathEncoderSuite struct {
	suite.Suite
}

func TestPathEncoderSuite(t *testing.T) {
	suite.Run(t, new(PathEncoderSuite))
}

func (suite *PathEncoderSuite) TestEncodeDecode() {
	t := suite.T()
	elements := []string{"these", "are", "some", "path", "elements"}

	encoded := encodeElements(elements...)

	decoded := make([]string, 0, len(elements))

	for _, e := range encoded {
		// Encoded values should not contain '/'.
		assert.False(t, strings.ContainsRune(e, '/'))
		dec, err := decodeElement(e)
		require.NoError(t, err)

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
		suite.T().Run(test.name, func(t *testing.T) {
			encoded := encodeAsPath(test.elements...)

			// Sanity check, first and last character should not be '/'.
			assert.Equal(t, strings.Trim(encoded, "/"), encoded)

			decoded := make([]string, 0, len(test.elements))

			for _, e := range strings.Split(encoded, "/") {
				dec, err := decodeElement(e)
				require.NoError(t, err)

				decoded = append(decoded, dec)
			}

			assert.Equal(t, test.expected, decoded)
		})
	}
}
