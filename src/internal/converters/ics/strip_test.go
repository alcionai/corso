package ics

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

//go:embed testdata/utf8.html
var utf8In string

//go:embed testdata/utf8.txt
var utf8Out string

//go:embed testdata/everything.html
var everythingIn string

//go:embed testdata/everything.txt
var everythingOut string

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
		{
			name:      "empty",
			input:     "",
			expected:  "",
			errAssert: assert.NoError,
		},
		{
			name:      "teams meeting",
			input:     utf8In,
			expected:  utf8Out,
			errAssert: assert.NoError,
		},
		{
			name:      "everything",
			input:     everythingIn,
			expected:  everythingOut,
			errAssert: assert.NoError,
		},
	}

	for _, tt := range table {
		s.Run(tt.name, func() {
			actual, err := HTMLToText(tt.input)
			tt.errAssert(s.T(), err)
			assert.Equal(s.T(), tt.expected, actual)
		})
	}
}
