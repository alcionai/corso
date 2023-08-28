package path

import (
	"fmt"
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

func (suite *PathUnitSuite) TestFromDataLayerPathErrors() {
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

func (suite *PathUnitSuite) TestFromDataLayerPath() {
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

func (suite *PathUnitSuite) TestPrefixOrPathFromDataLayerPath() {
	t := suite.T()
	input := fmt.Sprintf(
		"%s/%s/%s/%s",
		"tenant",
		ExchangeService.String(),
		"user",
		EmailCategory.String())

	// Check that we can make a valid prefix path.
	p, err := PrefixOrPathFromDataLayerPath(input, false)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, input, p.String())

	// Check we can't make a regular path with the same input since it doesn't
	// have enough segments.
	_, err = FromDataLayerPath(input, false)
	assert.Error(t, err)
}

func (suite *PathUnitSuite) TestBuildPrefix() {
	table := []struct {
		name      string
		service   ServiceType
		category  CategoryType
		tenant    string
		owner     string
		expect    string
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "ok",
			service:   ExchangeService,
			category:  ContactsCategory,
			tenant:    "t",
			owner:     "ro",
			expect:    join([]string{"t", ExchangeService.String(), "ro", ContactsCategory.String()}),
			expectErr: require.NoError,
		},
		{
			name:      "bad category",
			service:   ExchangeService,
			category:  FilesCategory,
			tenant:    "t",
			owner:     "ro",
			expectErr: require.Error,
		},
		{
			name:      "bad tenant",
			service:   ExchangeService,
			category:  ContactsCategory,
			tenant:    "",
			owner:     "ro",
			expectErr: require.Error,
		},
		{
			name:      "bad owner",
			service:   ExchangeService,
			category:  ContactsCategory,
			tenant:    "t",
			owner:     "",
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			r, err := BuildPrefix(test.tenant, test.owner, test.service, test.category)
			test.expectErr(t, err, clues.ToCore(err))

			if r == nil {
				return
			}

			assert.Equal(t, test.expect, r.String())
			assert.NotPanics(t, func() {
				r.Folders()
				r.Item()
			}, "runs Folders() and Item()")
		})
	}
}
