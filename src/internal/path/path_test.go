package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
