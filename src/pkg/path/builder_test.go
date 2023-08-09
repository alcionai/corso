package path

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type BuilderUnitSuite struct {
	tester.Suite
}

func TestBuilderUnitSuite(t *testing.T) {
	suite.Run(t, &BuilderUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// set the clues hashing to mask for the span of this suite
func (suite *BuilderUnitSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *BuilderUnitSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
}

func (suite *BuilderUnitSuite) TestAppend() {
	table := append(append([]testData{}, genericCases...), basicUnescapedInputs...)
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			p := Builder{}.Append(test.input...)
			assert.Equal(t, test.expectedString, p.String())
		})
	}
}

func (suite *BuilderUnitSuite) TestAppendItem() {
	t := suite.T()

	p, err := Build("t", "ro", ExchangeService, EmailCategory, false, "foo", "bar")
	require.NoError(t, err, clues.ToCore(err))

	pb := p.ToBuilder()
	assert.Equal(t, pb.String(), p.String())

	pb = pb.Append("qux")

	p, err = p.AppendItem("qux")

	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, pb.String(), p.String())

	_, err = p.AppendItem("fnords")
	require.Error(t, err, clues.ToCore(err))
}

func (suite *BuilderUnitSuite) TestUnescapeAndAppend() {
	table := append(append([]testData{}, genericCases...), basicEscapedInputs...)
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			p, err := Builder{}.UnescapeAndAppend(test.input...)
			require.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectedString, p.String())
		})
	}
}

func (suite *BuilderUnitSuite) TestEscapedFailure() {
	target := "i_s"

	for c := range charactersToEscape {
		suite.Run(fmt.Sprintf("Unescaped-%c", c), func() {
			tmp := strings.ReplaceAll(target, "_", string(c))

			_, err := Builder{}.UnescapeAndAppend("this", tmp, "path")
			assert.Errorf(suite.T(), err, "path with unescaped %s did not error", string(c))
		})
	}
}

func (suite *BuilderUnitSuite) TestBadEscapeSequenceErrors() {
	target := `i\_s/a`
	notEscapes := []rune{'a', 'b', '#', '%'}

	for _, c := range notEscapes {
		suite.Run(fmt.Sprintf("Escaped-%c", c), func() {
			tmp := strings.ReplaceAll(target, "_", string(c))

			_, err := Builder{}.UnescapeAndAppend("this", tmp, "path")
			assert.Errorf(
				suite.T(),
				err,
				"path with bad escape sequence %c%c did not error",
				escapeCharacter,
				c)
		})
	}
}

func (suite *BuilderUnitSuite) TestTrailingEscapeChar() {
	base := []string{"this", "is", "a", "path"}

	for i := 0; i < len(base); i++ {
		suite.Run(fmt.Sprintf("Element%v", i), func() {
			path := make([]string, len(base))
			copy(path, base)
			path[i] = path[i] + string(escapeCharacter)

			_, err := Builder{}.UnescapeAndAppend(path...)
			assert.Error(
				suite.T(),
				err,
				"path with trailing escape character did not error")
		})
	}
}

func (suite *BuilderUnitSuite) TestElements() {
	table := []struct {
		name     string
		input    []string
		output   []string
		pathFunc func(elements []string) (*Builder, error)
	}{
		{
			name:   "SimpleEscapedPath",
			input:  []string{"this", "is", "a", "path"},
			output: []string{"this", "is", "a", "path"},
			pathFunc: func(elements []string) (*Builder, error) {
				return Builder{}.UnescapeAndAppend(elements...)
			},
		},
		{
			name:   "SimpleUnescapedPath",
			input:  []string{"this", "is", "a", "path"},
			output: []string{"this", "is", "a", "path"},
			pathFunc: func(elements []string) (*Builder, error) {
				return Builder{}.Append(elements...), nil
			},
		},
		{
			name:   "EscapedPath",
			input:  []string{"this", `is\/`, "a", "path"},
			output: []string{"this", "is/", "a", "path"},
			pathFunc: func(elements []string) (*Builder, error) {
				return Builder{}.UnescapeAndAppend(elements...)
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			p, err := test.pathFunc(test.input)
			require.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, Elements(test.output), p.Elements())
		})
	}
}

func (suite *BuilderUnitSuite) TestPopFront() {
	table := []struct {
		name           string
		base           *Builder
		expectedString string
	}{
		{
			name:           "Empty",
			base:           &Builder{},
			expectedString: "",
		},
		{
			name:           "OneElement",
			base:           Builder{}.Append("something"),
			expectedString: "",
		},
		{
			name:           "TwoElements",
			base:           Builder{}.Append("something", "else"),
			expectedString: "else",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, test.expectedString, test.base.PopFront().String())
		})
	}
}

func (suite *BuilderUnitSuite) TestShortRef() {
	table := []struct {
		name          string
		inputElements []string
		expectedLen   int
	}{
		{
			name:          "PopulatedPath",
			inputElements: []string{"this", "is", "a", "path"},
			expectedLen:   shortRefCharacters,
		},
		{
			name:          "EmptyPath",
			inputElements: nil,
			expectedLen:   0,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			pb := Builder{}.Append(test.inputElements...)
			ref := pb.ShortRef()
			assert.Len(suite.T(), ref, test.expectedLen)
		})
	}
}

func (suite *BuilderUnitSuite) TestShortRefIsStable() {
	t := suite.T()
	pb := Builder{}.Append("this", "is", "a", "path")
	prevRef := pb.ShortRef()
	assert.Len(t, prevRef, shortRefCharacters)

	for i := 0; i < 5; i++ {
		ref := pb.ShortRef()
		assert.Len(t, ref, shortRefCharacters)
		assert.Equal(t, prevRef, ref, "ShortRef changed between calls")

		prevRef = ref
	}
}

func (suite *BuilderUnitSuite) TestShortRefIsUnique() {
	pb1 := Builder{}.Append("this", "is", "a", "path")
	pb2 := pb1.Append("also")

	require.NotEqual(suite.T(), pb1, pb2)
	assert.NotEqual(suite.T(), pb1.ShortRef(), pb2.ShortRef())
}

// TestShortRefUniqueWithEscaping tests that two paths that output the same
// unescaped string but different escaped strings have different shortrefs. This
// situation can occur when one path has embedded path separators while the
// other does not but contains the same characters.
func (suite *BuilderUnitSuite) TestShortRefUniqueWithEscaping() {
	pb1 := Builder{}.Append(`this`, `is`, `a`, `path`)
	pb2 := Builder{}.Append(`this`, `is/a`, `path`)

	require.NotEqual(suite.T(), pb1, pb2)
	assert.NotEqual(suite.T(), pb1.ShortRef(), pb2.ShortRef())
}

func (suite *BuilderUnitSuite) TestFolder() {
	table := []struct {
		name         string
		p            func(t *testing.T) Path
		escape       bool
		expectFolder string
		expectSplit  []string
	}{
		{
			name: "clean path",
			p: func(t *testing.T) Path {
				p, err := Builder{}.
					Append("a", "b", "c").
					ToDataLayerExchangePathForCategory("t", "u", EmailCategory, false)
				require.NoError(t, err, clues.ToCore(err))

				return p
			},
			expectFolder: "a/b/c",
			expectSplit:  []string{"a", "b", "c"},
		},
		{
			name: "clean path escaped",
			p: func(t *testing.T) Path {
				p, err := Builder{}.
					Append("a", "b", "c").
					ToDataLayerExchangePathForCategory("t", "u", EmailCategory, false)
				require.NoError(t, err, clues.ToCore(err))

				return p
			},
			escape:       true,
			expectFolder: "a/b/c",
			expectSplit:  []string{"a", "b", "c"},
		},
		{
			name: "escapable path",
			p: func(t *testing.T) Path {
				p, err := Builder{}.
					Append("a/", "b", "c").
					ToDataLayerExchangePathForCategory("t", "u", EmailCategory, false)
				require.NoError(t, err, clues.ToCore(err))

				return p
			},
			expectFolder: "a//b/c",
			expectSplit:  []string{"a", "b", "c"},
		},
		{
			name: "escapable path escaped",
			p: func(t *testing.T) Path {
				p, err := Builder{}.
					Append("a/", "b", "c").
					ToDataLayerExchangePathForCategory("t", "u", EmailCategory, false)
				require.NoError(t, err, clues.ToCore(err))

				return p
			},
			escape:       true,
			expectFolder: "a\\//b/c",
			expectSplit:  []string{"a\\/", "b", "c"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			p := test.p(t)
			result := p.Folder(test.escape)
			assert.Equal(t, test.expectFolder, result)
			assert.Equal(t, test.expectSplit, Split(result))
		})
	}
}

func (suite *BuilderUnitSuite) TestPIIHandling() {
	p, err := Build("t", "ro", ExchangeService, EventsCategory, true, "dir", "item")
	require.NoError(suite.T(), err)

	table := []struct {
		name        string
		p           Path
		expect      string
		expectPlain string
	}{
		{
			name:        "standard path",
			p:           p,
			expect:      "***/exchange/***/events/***/***",
			expectPlain: "t/exchange/ro/events/dir/item",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, test.expect, test.p.Conceal(), "conceal")
			assert.Equal(t, test.expectPlain, test.p.String(), "string")
			assert.Equal(t, test.expect, fmt.Sprintf("%s", test.p), "fmt %%s")
			assert.Equal(t, test.expect, fmt.Sprintf("%+v", test.p), "fmt %%+v")
			assert.Equal(t, test.expectPlain, test.p.PlainString(), "plain")
		})
	}
}
