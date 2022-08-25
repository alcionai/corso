package path

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testData struct {
	name           string
	input          []string
	expectedString string
}

// Test cases that are the same with and without escaping by the
// system-under-test.
var genericCases = []testData{
	{
		name: "SimplePath",
		input: []string{
			`this`,
			`is`,
			`a`,
			`path`,
		},
		expectedString: "this/is/a/path",
	},
	{
		name: "EmptyElement",
		input: []string{
			`this`,
			`is`,
			``,
			`a`,
			`path`,
		},
		expectedString: `this/is/a/path`,
	},
	{
		name:           "EmptyInput",
		expectedString: "",
	},
}

// Inputs that should be escaped.
var basicUnescapedInputs = []testData{
	{
		name: "EscapeSeparator",
		input: []string{
			`this`,
			`is/a`,
			`path`,
		},
		expectedString: `this/is\/a/path`,
	},
	{
		name: "EscapeEscapeChar",
		input: []string{
			`this`,
			`is\`,
			`a`,
			`path`,
		},
		expectedString: `this/is\\/a/path`,
	},
	{
		name: "EscapeEscapeAndSeparator",
		input: []string{
			`this`,
			`is\/a`,
			`path`,
		},
		expectedString: `this/is\\\/a/path`,
	},
	{
		name: "SeparatorAtEndOfElement",
		input: []string{
			`this`,
			`is/`,
			`a`,
			`path`,
		},
		expectedString: `this/is\//a/path`,
	},
	{
		name: "SeparatorAtEndOfPath",
		input: []string{
			`this`,
			`is`,
			`a`,
			`path/`,
		},
		expectedString: `this/is/a/path\/`,
	},
}

// Inputs that are already escaped.
var basicEscapedInputs = []testData{
	{
		name: "EscapedSeparator",
		input: []string{
			`this`,
			`is\/a`,
			`path`,
		},
		expectedString: `this/is\/a/path`,
	},
	{
		name: "EscapedEscapeChar",
		input: []string{
			`this`,
			`is\\`,
			`a`,
			`path`,
		},
		expectedString: `this/is\\/a/path`,
	},
	{
		name: "EscapedEscapeAndSeparator",
		input: []string{
			`this`,
			`is\\\/a`,
			`path`,
		},
		expectedString: `this/is\\\/a/path`,
	},
	{
		name: "EscapedSeparatorAtEndOfElement",
		input: []string{
			`this`,
			`is\/`,
			`a`,
			`path`,
		},
		expectedString: `this/is\//a/path`,
	},
	{
		name: "EscapedSeparatorAtEndOfPath",
		input: []string{
			`this`,
			`is`,
			`a`,
			`path\/`,
		},
		expectedString: `this/is/a/path\/`,
	},
	{
		name: "ElementOfSeparator",
		input: []string{
			`this`,
			`is`,
			`/`,
			`a`,
			`path`,
		},
		expectedString: `this/is/a/path`,
	},
	{
		name: "TrailingElementSeparator",
		input: []string{
			`this`,
			`is`,
			`a/`,
			`path`,
		},
		expectedString: `this/is/a/path`,
	},
	{
		name: "TrailingSeparatorAtEnd",
		input: []string{
			`this`,
			`is`,
			`a`,
			`path/`,
		},
		expectedString: `this/is/a/path`,
	},
	{
		name: "TrailingSeparatorWithEmptyElementAtEnd",
		input: []string{
			`this`,
			`is`,
			`a`,
			`path/`,
			``,
		},
		expectedString: `this/is/a/path`,
	},
}

type PathUnitSuite struct {
	suite.Suite
}

func TestPathUnitSuite(t *testing.T) {
	suite.Run(t, new(PathUnitSuite))
}

func (suite *PathUnitSuite) TestAppend() {
	table := append(append([]testData{}, genericCases...), basicUnescapedInputs...)
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			p := Builder{}.Append(test.input...)
			assert.Equal(t, test.expectedString, p.String())
		})
	}
}

func (suite *PathUnitSuite) TestUnescapeAndAppend() {
	table := append(append([]testData{}, genericCases...), basicEscapedInputs...)
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			p, err := Builder{}.UnescapeAndAppend(test.input...)
			require.NoError(t, err)

			assert.Equal(t, test.expectedString, p.String())
		})
	}
}

func (suite *PathUnitSuite) TestEscapedFailure() {
	target := "i_s"

	for c := range charactersToEscape {
		suite.T().Run(fmt.Sprintf("Unescaped-%c", c), func(t *testing.T) {
			tmp := strings.ReplaceAll(target, "_", string(c))

			_, err := Builder{}.UnescapeAndAppend("this", tmp, "path")
			assert.Error(t, err, "path with unescaped %s did not error", string(c))
		})
	}
}

func (suite *PathUnitSuite) TestBadEscapeSequenceErrors() {
	target := `i\_s/a`
	notEscapes := []rune{'a', 'b', '#', '%'}

	for _, c := range notEscapes {
		suite.T().Run(fmt.Sprintf("Escaped-%c", c), func(t *testing.T) {
			tmp := strings.ReplaceAll(target, "_", string(c))

			_, err := Builder{}.UnescapeAndAppend("this", tmp, "path")
			assert.Error(
				t,
				err,
				"path with bad escape sequence %c%c did not error",
				escapeCharacter,
				c,
			)
		})
	}
}

func (suite *PathUnitSuite) TestTrailingEscapeChar() {
	base := []string{"this", "is", "a", "path"}

	for i := 0; i < len(base); i++ {
		suite.T().Run(fmt.Sprintf("Element%v", i), func(t *testing.T) {
			path := make([]string, len(base))
			copy(path, base)
			path[i] = path[i] + string(escapeCharacter)

			_, err := Builder{}.UnescapeAndAppend(path...)
			assert.Error(
				t,
				err,
				"path with trailing escape character did not error",
			)
		})
	}
}
