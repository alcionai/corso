package ics

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

var (
	//go:embed testdata/simple.html
	simpleIn string

	//go:embed testdata/simple.txt
	simpleOut string

	//go:embed testdata/utf8.html
	utf8In string

	//go:embed testdata/utf8.txt
	utf8Out string

	//go:embed testdata/everything.html
	everythingIn string

	//go:embed testdata/everything.txt
	everythingOut string
)

type StripUnitSuite struct {
	tester.Suite
}

func TestStripUnitSuite(t *testing.T) {
	suite.Run(t, &StripUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (s *StripUnitSuite) TestStrip() {
	table := []struct {
		name      string
		input     string
		expected  string
		errAssert assert.ErrorAssertionFunc
	}{
		// {
		// 	name:      "empty",
		// 	input:     "",
		// 	expected:  "",
		// 	errAssert: assert.NoError,
		// },
		{
			name:      "line with spans",
			input:     simpleIn,
			expected:  simpleOut,
			errAssert: assert.NoError,
		},
		// {
		// 	name:      "teams meeting",
		// 	input:     utf8In,
		// 	expected:  utf8Out,
		// 	errAssert: assert.NoError,
		// },
		// {
		// 	name:      "everything",
		// 	input:     everythingIn,
		// 	expected:  everythingOut,
		// 	errAssert: assert.NoError,
		// },
	}

	for _, tt := range table {
		s.Run(tt.name, func() {
			actual, err := HTMLToText(tt.input)
			tt.errAssert(s.T(), err)
			assert.Equal(s.T(), tt.expected, actual)
		})
	}
}
