package ics

import (
	_ "embed"
	"testing"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//go:embed testdata/utf8.html
var utf8_in string

//go:embed testdata/utf8.txt
var utf8_out string

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
			input:     utf8_in,
			expected:  utf8_out,
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
