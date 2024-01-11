package ics

import (
	_ "embed"
	"testing"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//go:embed testdata/in.html
var in string

//go:embed testdata/out.txt
var out string

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
			input:     in,
			expected:  out,
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
