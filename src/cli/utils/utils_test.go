package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester/aw"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type CliUtilsSuite struct {
	suite.Suite
}

func TestCliUtilsSuite(t *testing.T) {
	suite.Run(t, new(CliUtilsSuite))
}

func (suite *CliUtilsSuite) TestRequireProps() {
	table := []struct {
		name     string
		props    map[string]string
		errCheck assert.ErrorAssertionFunc
	}{
		{
			props:    map[string]string{"exists": "I have seen the fnords!"},
			errCheck: aw.NoErr,
		},
		{
			props:    map[string]string{"not-exists": ""},
			errCheck: aw.Err,
		},
	}
	for _, test := range table {
		test.errCheck(suite.T(), RequireProps(test.props))
	}
}

func (suite *CliUtilsSuite) TestSplitFoldersIntoContainsAndPrefix() {
	table := []struct {
		name    string
		input   []string
		expectC []string
		expectP []string
	}{
		{
			name:    "empty",
			expectC: selectors.Any(),
			expectP: nil,
		},
		{
			name:    "only contains",
			input:   []string{"a", "b", "c"},
			expectC: []string{"a", "b", "c"},
			expectP: []string{},
		},
		{
			name:    "only leading slash counts as contains",
			input:   []string{"a/////", "\\/b", "\\//c\\/"},
			expectC: []string{"a/////", "\\/b", "\\//c\\/"},
			expectP: []string{},
		},
		{
			name:    "only prefix",
			input:   []string{"/a", "/b", "/\\/c"},
			expectC: []string{},
			expectP: []string{"/a", "/b", "/\\/c"},
		},
		{
			name:    "mixed",
			input:   []string{"/a", "b", "/c"},
			expectC: []string{"b"},
			expectP: []string{"/a", "/c"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			c, p := splitFoldersIntoContainsAndPrefix(test.input)
			assert.ElementsMatch(t, test.expectC, c, "contains set")
			assert.ElementsMatch(t, test.expectP, p, "prefix set")
		})
	}
}
