package pii_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/tester"
)

type PIIUnitSuite struct {
	tester.Suite
}

func TestPIIUnitSuite(t *testing.T) {
	suite.Run(t, &PIIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// set the clues hashing to mask for the span of this suite
func (suite *PIIUnitSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *PIIUnitSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
}

func (suite *PIIUnitSuite) TestMapWithPlurals() {
	t := suite.T()

	mwp := pii.MapWithPlurals()
	assert.Equal(t, map[string]struct{}{}, mwp)

	mwp = pii.MapWithPlurals("")
	assert.Equal(t, map[string]struct{}{"": {}, "s": {}}, mwp)

	mwp = pii.MapWithPlurals(" ")
	assert.Equal(t, map[string]struct{}{" ": {}, " s": {}}, mwp)

	mwp = pii.MapWithPlurals("foo", "bar")
	expect := map[string]struct{}{
		"foo":  {},
		"foos": {},
		"bar":  {},
		"bars": {},
	}
	assert.Equal(t, expect, mwp)
}

func (suite *PIIUnitSuite) TestConcealElements() {
	table := []struct {
		name   string
		in     []string
		safe   map[string]struct{}
		expect []string
	}{
		{
			name:   "nil",
			expect: []string{},
		},
		{
			name:   "no safe words",
			in:     []string{"fnords", "smarfs"},
			expect: []string{"***", "***"},
		},
		{
			name:   "safe words",
			in:     []string{"fnords", "smarfs"},
			safe:   map[string]struct{}{"fnords": {}},
			expect: []string{"fnords", "***"},
		},
		{
			name:   "non-matching safe words",
			in:     []string{"fnords", "smarfs"},
			safe:   map[string]struct{}{"beaux": {}},
			expect: []string{"***", "***"},
		},
		{
			name:   "case insensitivity",
			in:     []string{"FNORDS", "SMARFS"},
			safe:   map[string]struct{}{"fnords": {}},
			expect: []string{"FNORDS", "***"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := pii.ConcealElements(test.in, test.safe)
			assert.ElementsMatch(t, test.expect, result)
		})
	}
}
