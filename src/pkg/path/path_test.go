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
		name: "MultipleTrailingElementSeparator",
		input: []string{
			`this`,
			`is`,
			`a///`,
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
		name: "MultipleTrailingSeparatorAtEnd",
		input: []string{
			`this`,
			`is`,
			`a`,
			`path///`,
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
	tester.Suite
}

func TestPathUnitSuite(t *testing.T) {
	suite.Run(t, &PathUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// set the clues hashing to mask for the span of this suite
func (suite *PathUnitSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *PathUnitSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
}

func (suite *PathUnitSuite) TestAppend() {
	table := append(append([]testData{}, genericCases...), basicUnescapedInputs...)
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			p := Builder{}.Append(test.input...)
			assert.Equal(t, test.expectedString, p.String())
		})
	}
}

func (suite *PathUnitSuite) TestUnescapeAndAppend() {
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

func (suite *PathUnitSuite) TestEscapedFailure() {
	target := "i_s"

	for c := range charactersToEscape {
		suite.Run(fmt.Sprintf("Unescaped-%c", c), func() {
			tmp := strings.ReplaceAll(target, "_", string(c))

			_, err := Builder{}.UnescapeAndAppend("this", tmp, "path")
			assert.Errorf(suite.T(), err, "path with unescaped %s did not error", string(c))
		})
	}
}

func (suite *PathUnitSuite) TestBadEscapeSequenceErrors() {
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

func (suite *PathUnitSuite) TestTrailingEscapeChar() {
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

func (suite *PathUnitSuite) TestElements() {
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

func (suite *PathUnitSuite) TestPopFront() {
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

func (suite *PathUnitSuite) TestShortRef() {
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

func (suite *PathUnitSuite) TestShortRefIsStable() {
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

func (suite *PathUnitSuite) TestShortRefIsUnique() {
	pb1 := Builder{}.Append("this", "is", "a", "path")
	pb2 := pb1.Append("also")

	require.NotEqual(suite.T(), pb1, pb2)
	assert.NotEqual(suite.T(), pb1.ShortRef(), pb2.ShortRef())
}

// TestShortRefUniqueWithEscaping tests that two paths that output the same
// unescaped string but different escaped strings have different shortrefs. This
// situation can occur when one path has embedded path separators while the
// other does not but contains the same characters.
func (suite *PathUnitSuite) TestShortRefUniqueWithEscaping() {
	pb1 := Builder{}.Append(`this`, `is`, `a`, `path`)
	pb2 := Builder{}.Append(`this`, `is/a`, `path`)

	require.NotEqual(suite.T(), pb1, pb2)
	assert.NotEqual(suite.T(), pb1.ShortRef(), pb2.ShortRef())
}

func (suite *PathUnitSuite) TestFromStringErrors() {
	table := []struct {
		name        string
		escapedPath string
	}{
		{
			name:        "TooFewElements",
			escapedPath: `some/short/path`,
		},
		{
			name:        "TooFewElementsEmptyElement",
			escapedPath: `tenant/exchange//email/folder`,
		},
		{
			name:        "BadEscapeSequence",
			escapedPath: `tenant/exchange/user/email/folder\a`,
		},
		{
			name:        "TrailingEscapeCharacter",
			escapedPath: `tenant/exchange/user/email/folder\`,
		},
		{
			name:        "UnknownService",
			escapedPath: `tenant/badService/user/email/folder`,
		},
		{
			name:        "UnknownCategory",
			escapedPath: `tenant/exchange/user/badCategory/folder`,
		},
		{
			name:        "NoFolderOrItem",
			escapedPath: `tenant/exchange/user/email`,
		},
		{
			name:        "EmptyPath",
			escapedPath: ``,
		},
		{
			name:        "JustPathSeparator",
			escapedPath: `/`,
		},
		{
			name:        "JustMultiplePathSeparators",
			escapedPath: `//`,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, err := FromDataLayerPath(test.escapedPath, false)
			assert.Error(t, err)
		})
	}
}

func (suite *PathUnitSuite) TestFolder() {
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

func (suite *PathUnitSuite) TestFromString() {
	const (
		testTenant         = "tenant"
		testUser           = "user"
		testElement1       = "folder/"
		testElementTrimmed = "folder"
		testElement2       = "folder2"
		testElement3       = "other"
	)

	isItem := []struct {
		name   string
		isItem bool
	}{
		{
			name:   "Folder",
			isItem: false,
		},
		{
			name:   "Item",
			isItem: true,
		},
	}
	table := []struct {
		name string
		// Should have placeholders of '%s' for service and category.
		unescapedPath string
		// Expected result for Folder() if path is marked as a folder.
		expectedFolder string
		// Expected result for Item() if path is marked as an item.
		// Expected result for Split(Folder()) if path is marked as a folder.
		expectedSplit []string
		expectedItem  string
		// Expected result for Folder() if path is marked as an item.
		expectedItemFolder string
		// Expected result for Split(Folder()) if path is marked as an item.
		expectedItemSplit []string
	}{
		{
			name: "BasicPath",
			unescapedPath: fmt.Sprintf(
				"%s/%%s/%s/%%s/%s/%s/%s",
				testTenant,
				testUser,
				testElement1,
				testElement2,
				testElement3,
			),
			expectedFolder: fmt.Sprintf(
				"%s/%s/%s",
				testElementTrimmed,
				testElement2,
				testElement3,
			),
			expectedSplit: []string{
				testElementTrimmed,
				testElement2,
				testElement3,
			},
			expectedItem: testElement3,
			expectedItemFolder: fmt.Sprintf(
				"%s/%s",
				testElementTrimmed,
				testElement2,
			),
			expectedItemSplit: []string{
				testElementTrimmed,
				testElement2,
			},
		},
		{
			name: "PathWithEmptyElements",
			unescapedPath: fmt.Sprintf(
				"/%s//%%s//%s//%%s//%s///%s//%s//",
				testTenant,
				testUser,
				testElementTrimmed,
				testElement2,
				testElement3,
			),
			expectedFolder: fmt.Sprintf(
				"%s/%s/%s",
				testElementTrimmed,
				testElement2,
				testElement3,
			),
			expectedSplit: []string{
				testElementTrimmed,
				testElement2,
				testElement3,
			},
			expectedItem: testElement3,
			expectedItemFolder: fmt.Sprintf(
				"%s/%s",
				testElementTrimmed,
				testElement2,
			),
			expectedItemSplit: []string{
				testElementTrimmed,
				testElement2,
			},
		},
	}

	for service, cats := range serviceCategories {
		for cat := range cats {
			for _, item := range isItem {
				suite.Run(fmt.Sprintf("%s-%s-%s", service, cat, item.name), func() {
					for _, test := range table {
						suite.Run(test.name, func() {
							t := suite.T()
							testPath := fmt.Sprintf(test.unescapedPath, service, cat)

							p, err := FromDataLayerPath(testPath, item.isItem)
							require.NoError(t, err, clues.ToCore(err))

							assert.Equal(t, service, p.Service(), "service")
							assert.Equal(t, cat, p.Category(), "category")
							assert.Equal(t, testTenant, p.Tenant(), "tenant")
							assert.Equal(t, testUser, p.ResourceOwner(), "resource owner")

							fld := p.Folder(false)
							escfld := p.Folder(true)

							if item.isItem {
								assert.Equal(t, test.expectedItemFolder, fld, "item folder")
								assert.Equal(t, test.expectedItemSplit, Split(fld), "item split")
								assert.Equal(t, test.expectedItemFolder, escfld, "escaped item folder")
								assert.Equal(t, test.expectedItemSplit, Split(escfld), "escaped item split")
								assert.Equal(t, test.expectedItem, p.Item(), "item")
							} else {
								assert.Equal(t, test.expectedFolder, fld, "dir folder")
								assert.Equal(t, test.expectedSplit, Split(fld), "dir split")
								assert.Equal(t, test.expectedFolder, escfld, "escaped dir folder")
								assert.Equal(t, test.expectedSplit, Split(escfld), "escaped dir split")
							}
						})
					}
				})
			}
		}
	}
}

func (suite *PathUnitSuite) TestPath_piiHandling() {
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
