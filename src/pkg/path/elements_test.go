package path

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ElementsUnitSuite struct {
	tester.Suite
}

func TestElementsUnitSuite(t *testing.T) {
	suite.Run(t, &ElementsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// set the clues hashing to mask for the span of this suite
func (suite *ElementsUnitSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *ElementsUnitSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
}

func (suite *ElementsUnitSuite) TestNewElements() {
	t := suite.T()

	result := NewElements("")
	assert.Equal(t, Elements{""}, result)

	result = NewElements("fnords")
	assert.Equal(t, Elements{"fnords"}, result)

	result = NewElements("fnords/smarf")
	assert.Equal(t, Elements{"fnords", "smarf"}, result)
}

func (suite *ElementsUnitSuite) TestElements_piiHandling() {
	table := []struct {
		name         string
		elems        Elements
		expect       string
		expectString string
		expectPlain  string
	}{
		{
			name:         "all concealed",
			elems:        Elements{"foo", "bar/", "baz"},
			expect:       "***/***/***",
			expectString: `foo/bar\//baz`,
			expectPlain:  `foo/bar//baz`,
		},
		{
			name:         "all safe",
			elems:        Elements{UnknownService.String(), UnknownCategory.String(), ExchangeMetadataService.String()},
			expect:       "UnknownService/UnknownCategory/exchangeMetadata",
			expectString: "UnknownService/UnknownCategory/exchangeMetadata",
			expectPlain:  "UnknownService/UnknownCategory/exchangeMetadata",
		},
		{
			name:         "mixed",
			elems:        Elements{UnknownService.String(), "smarf", ExchangeMetadataService.String()},
			expect:       "UnknownService/***/exchangeMetadata",
			expectString: "UnknownService/smarf/exchangeMetadata",
			expectPlain:  "UnknownService/smarf/exchangeMetadata",
		},
		{
			name:         "empty elements",
			elems:        Elements{},
			expect:       "",
			expectString: "",
			expectPlain:  "",
		},
		{
			name:         "empty string",
			elems:        Elements{""},
			expect:       "",
			expectString: "",
			expectPlain:  "",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, test.expect, test.elems.Conceal(), "conceal")
			assert.Equal(t, test.expectString, test.elems.String(), "string")
			assert.Equal(t, test.expect, fmt.Sprintf("%s", test.elems), "fmt %%s")
			assert.Equal(t, test.expect, fmt.Sprintf("%+v", test.elems), "fmt %%+v")
			assert.Equal(t, test.expectPlain, join(test.elems), "plain")
		})
	}
}

func (suite *ElementsUnitSuite) TestLoggableDir() {
	table := []struct {
		inpt   string
		expect string
	}{
		{
			inpt:   "archive/clutter",
			expect: "archive/clutter",
		},
		{
			inpt:   "foo/bar",
			expect: "***/***",
		},
		{
			inpt:   "inbox/foo",
			expect: "inbox/***",
		},
		{
			inpt:   "foo/",
			expect: "***",
		},
		{
			inpt:   "foo//",
			expect: "***",
		},
		{
			inpt:   "foo///",
			expect: "***",
		},
	}
	for _, test := range table {
		suite.Run(test.inpt, func() {
			assert.Equal(suite.T(), test.expect, LoggableDir(test.inpt))
		})
	}
}
