package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type CategoryTypeUnitSuite struct {
	tester.Suite
}

func TestCategoryTypeUnitSuite(t *testing.T) {
	suite.Run(t, &CategoryTypeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CategoryTypeUnitSuite) TestToCategoryType() {
	table := []struct {
		input  string
		expect CategoryType
	}{
		{input: "unknown", expect: 0},
		{input: "EMAIL", expect: 1},
		{input: "Email", expect: 1},
		{input: "email", expect: 1},
		{input: "contacts", expect: 2},
		{input: "events", expect: 3},
		{input: "files", expect: 4},
		{input: "lists", expect: 5},
		{input: "libraries", expect: 6},
		{input: "pages", expect: 7},
		{input: "details", expect: 8},
		{input: "channelmessages", expect: 9},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			assert.Equal(
				suite.T(),
				test.expect,
				ToCategoryType(test.input))
		})
	}
}

func (suite *CategoryTypeUnitSuite) TestHumanString() {
	table := []struct {
		input  CategoryType
		expect string
	}{
		{input: 0, expect: "Unknown Category"},
		{input: 1, expect: "Emails"},
		{input: 2, expect: "Contacts"},
		{input: 3, expect: "Events"},
		{input: 4, expect: "Files"},
		{input: 5, expect: "Lists"},
		{input: 6, expect: "Libraries"},
		{input: 7, expect: "Pages"},
		{input: 8, expect: "Details"},
		{input: 9, expect: "Messages"},
	}
	for _, test := range table {
		suite.Run(test.input.String(), func() {
			assert.Equal(
				suite.T(),
				test.expect,
				test.input.HumanString())
		})
	}
}
