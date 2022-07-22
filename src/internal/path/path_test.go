package path

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var basicInputs = []struct {
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

type PathUnitSuite struct {
	suite.Suite
}

func TestPathUnitSuite(t *testing.T) {
	suite.Run(t, new(PathUnitSuite))
}

func (suite *PathUnitSuite) TestPathEscapingAndSegments() {
	for _, test := range basicInputs {
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

func (suite *PathUnitSuite) TestPathEscapingAndSegments_EmpytElements() {
	table := []struct {
		name     string
		input    [][]string
		expected string
	}{
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
			name: "EmptyInternalElement2",
			input: [][]string{
				{`this`},
				{`is`},
				{"", "", ""},
				{`a`},
				{`path`},
			},
			expected: "this/is/a/path",
		},
		{
			name: "EmptyInternalElement3",
			input: [][]string{
				{`this`},
				{`is`},
				{},
				{`a`},
				{`path`},
			},
			expected: "this/is/a/path",
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			p := newPath(test.input)

			idx := 0
			for i := 0; i < len(test.input); i++ {
				if i == 2 {
					continue
				}

				assert.NotPanics(t, func() {
					_ = p.segment(idx)
				})
				idx++
			}

			assert.Panics(t, func() {
				_ = p.segment(len(test.input))
			})
		})
	}
}

func (suite *PathUnitSuite) TestElementUnescaping() {
	for _, test := range basicInputs {
		suite.T().Run(test.name, func(t *testing.T) {
			p := newPath(test.input)

			for i, s := range test.input {
				elements := []string{}
				require.NotPanics(t, func() {
					elements = p.unescapedSegmentElements(i)
				})

				assert.True(t, reflect.DeepEqual(s, elements))
			}

			assert.Panics(t, func() {
				_ = p.unescapedSegmentElements(len(test.input))
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
		{
			name:             "ManyEscapesAndSeparator",
			input:            []string{`this`, `is\\\/a`, `path`},
			expected:         `this/is\\\/a/path`,
			expectedSegments: []string{`this`, `is\\\/a`, `path`},
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

func (suite *PathUnitSuite) TestBadEscapeSequenceErrors() {
	target := `i\_s/a`
	notEscapes := []rune{'a', 'b', '#', '%'}

	for _, c := range notEscapes {
		tmp := strings.ReplaceAll(target, "_", string(c))
		basePath := []string{"this", tmp, "path"}
		_, err := newPathFromEscapedSegments(basePath)
		assert.Error(
			suite.T(),
			err,
			"path with bad escape sequence %c%c did not error",
			escapeCharacter,
			c,
		)
	}
}

func (suite *PathUnitSuite) TestTrailingEscapeChar() {
	base := []string{"this", "is", "a", "path"}

	for i := 0; i < len(base); i++ {
		suite.T().Run(fmt.Sprintf("Segment%v", i), func(t *testing.T) {
			path := make([]string, len(base))
			copy(path, base)
			path[i] = path[i] + string(escapeCharacter)

			_, err := newPathFromEscapedSegments(path)
			assert.Error(suite.T(), err)
		})
	}
}
