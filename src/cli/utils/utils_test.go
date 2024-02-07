package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/pkg/selectors"
)

type CliUtilsSuite struct {
	tester.Suite
}

func TestCliUtilsSuite(t *testing.T) {
	suite.Run(t, &CliUtilsSuite{Suite: tester.NewUnitSuite(t)})
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
		suite.Run(test.name, func() {
			t := suite.T()

			c, p := splitFoldersIntoContainsAndPrefix(test.input)
			assert.ElementsMatch(t, test.expectC, c, "contains set")
			assert.ElementsMatch(t, test.expectP, p, "prefix set")
		})
	}
}

// Test MakeAbsoluteFilePath
func (suite *CliUtilsSuite) TestMakeAbsoluteFilePath() {
	currentDir, err := os.Getwd()
	assert.NoError(suite.T(), err)

	homeDir, err := os.UserHomeDir()
	assert.NoError(suite.T(), err)

	table := []struct {
		name        string
		input       string
		expected    string
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name:        "empty path",
			input:       "",
			expected:    "",
			expectedErr: assert.Error,
		},
		{
			name:        "absolute path",
			input:       "/tmp/dir",
			expected:    "/tmp/dir",
			expectedErr: assert.NoError,
		},
		{
			name:        "relative path",
			input:       "subdir/file.txt",
			expected:    filepath.Join(currentDir, "subdir/file.txt"),
			expectedErr: assert.NoError,
		},
		{
			name:        "relative path 2",
			input:       ".",
			expected:    currentDir,
			expectedErr: assert.NoError,
		},
		{
			name:        "home dir",
			input:       "~/file.txt",
			expected:    filepath.Join(homeDir, "file.txt"),
			expectedErr: assert.NoError,
		},
		{
			name:        "home dir 2",
			input:       "~",
			expected:    homeDir,
			expectedErr: assert.NoError,
		},
		{
			name:        "relative path with home dir",
			input:       "~/test/..",
			expected:    homeDir,
			expectedErr: assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			actual, err := MakeAbsoluteFilePath(test.input)
			assert.Equal(t, test.expected, actual)

			test.expectedErr(t, err)
		})
	}
}
