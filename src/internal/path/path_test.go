package path

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PathUnitSuite struct {
	suite.Suite
}

func TestPathUnitSuite(t *testing.T) {
	suite.Run(t, new(PathUnitSuite))
}

func (suite *PathUnitSuite) TestPathEscapingAndSegments() {
	table := []struct {
		name     string
		input    [][]string
		expected string
	}{
		{
			name: "SimplePath",
			input: [][]string{
				{`this`},
				{`is`},
				{`a`},
				{`path`},
			},
			expected: "this/is/a/path",
		},
		{
			name: "EscapeSeparator",
			input: [][]string{
				{`this`},
				{`is/a`},
				{`path`},
			},
			expected: `this/is\/a/path`,
		},
		{
			name: "EscapeEscapeChar",
			input: [][]string{
				{`this`},
				{`is\`},
				{`a`},
				{`path`},
			},
			expected: `this/is\\/a/path`,
		},
		{
			name: "EscapeEscapeAndSeparator",
			input: [][]string{
				{`this`},
				{`is\/a`},
				{`path`},
			},
			expected: `this/is\\\/a/path`,
		},
		{
			name: "EmptyInternalElement",
			input: [][]string{
				{`this`},
				{`is`},
				{""},
				{`a`},
				{`path`},
			},
			expected: "this/is/a/path",
		},
		{
			name: "SeparatorAtEndOfElement",
			input: [][]string{
				{`this`},
				{`is/`},
				{`a`},
				{`path`},
			},
			expected: `this/is\//a/path`,
		},
		{
			name: "SeparatorAtEndOfPath",
			input: [][]string{
				{`this`},
				{`is`},
				{`a`},
				{`path/`},
			},
			expected: `this/is/a/path\/`,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			p := newPath(test.input)
			assert.Equal(t, test.expected, p.String())

			for i := 0; i < len(test.input); i++ {
				assert.NotPanics(t, func() {
					_ = p.segment(i)
				})
			}

			assert.Panics(t, func() {
				_ = p.segment(len(test.input))
			})
		})
	}
}

func (suite *PathUnitSuite) TestPathSplitsEscapedPath() {
	table := []struct {
		name             string
		input            []string
		expected         string
		expectedSegments []string
	}{
		{
			name:             "SimplePath",
			input:            []string{`this`, `is/a`, `path`},
			expected:         "this/is/a/path",
			expectedSegments: []string{`this`, `is/a`, `path`},
		},
		{
			name:             "EscapeSeparator",
			input:            []string{`this`, `is\/a`, `path`},
			expected:         `this/is\/a/path`,
			expectedSegments: []string{`this`, `is\/a`, `path`},
		},
		{
			name:             "EscapeEscapeChar",
			input:            []string{`this`, `is\\/a`, `path`},
			expected:         `this/is\\/a/path`,
			expectedSegments: []string{`this`, `is\\/a`, `path`},
		},
		{
			name:             "EscapeEscapeAndSeparator",
			input:            []string{`this`, `is\\\/a`, `path`},
			expected:         `this/is\\\/a/path`,
			expectedSegments: []string{`this`, `is\\\/a`, `path`},
		},
		{
			name:             "EmptyInternalElement",
			input:            []string{`this`, `is//a`, `path`},
			expected:         "this/is/a/path",
			expectedSegments: []string{`this`, `is/a`, `path`},
		},
		{
			name:             "SeparatorAtEndOfElement",
			input:            []string{`this`, `is\//a`, `path`},
			expected:         `this/is\//a/path`,
			expectedSegments: []string{`this`, `is\//a`, `path`},
		},
		{
			name:             "SeparatorAtEndOfPath",
			input:            []string{`this`, `is/a`, `path\/`},
			expected:         `this/is/a/path\/`,
			expectedSegments: []string{`this`, `is/a`, `path\/`},
		},
		{
			name:             "TrailingSeparator",
			input:            []string{`this`, `is/a`, `path/`},
			expected:         `this/is/a/path`,
			expectedSegments: []string{`this`, `is/a`, `path`},
		},
		{
			name:             "TrailingSeparator2",
			input:            []string{`this`, `is/a`, `path\\\\/`},
			expected:         `this/is/a/path\\\\`,
			expectedSegments: []string{`this`, `is/a`, `path\\\\`},
		},
		{
			name:             "ManyEscapesNotSeparator",
			input:            []string{`this`, `is\\\\/a`, `path/`},
			expected:         `this/is\\\\/a/path`,
			expectedSegments: []string{`this`, `is\\\\/a`, `path`},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			p, err := newPathFromEscapedSegments(test.input)
			require.NoError(t, err)
			assert.Equal(t, test.expected, p.String())

			for i, s := range test.expectedSegments {
				segment := ""
				require.NotPanics(t, func() {
					segment = p.segment(i)
				})

				assert.Equal(t, s, segment)
			}
		})
	}
}

func (suite *PathUnitSuite) TestEscapedFailure() {
	target := "i_s/a"

	for c := range charactersToEscape {
		if c == pathSeparator {
			// Extra path separators in the path will just lead to more segments, not
			// a validation error.
			continue
		}

		tmp := strings.ReplaceAll(target, "_", string(c))
		basePath := []string{"this", tmp, "path"}
		_, err := newPathFromEscapedSegments(basePath)
		assert.Error(suite.T(), err, "path with unescaped %s did not error", string(c))
	}
}

func (suite *PathUnitSuite) TestTrailingEscapeChar() {
	path := []string{"this", "is", "a", `path\`}
	_, err := newPathFromEscapedSegments(path)
	assert.Error(suite.T(), err)
}
