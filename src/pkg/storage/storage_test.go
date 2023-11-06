package storage

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type testConfig struct {
	expect string
	err    error
}

func (c testConfig) StringConfig() (map[string]string, error) {
	return map[string]string{"expect": c.expect}, c.err
}

type StorageUnitSuite struct {
	tester.Suite
}

func TestStorageUnitSuite(t *testing.T) {
	suite.Run(t, &StorageUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *StorageUnitSuite) TestNewStorage() {
	table := []struct {
		name     string
		p        ProviderType
		c        testConfig
		errCheck assert.ErrorAssertionFunc
	}{
		{"unknown no error", ProviderUnknown, testConfig{"configVal", nil}, assert.NoError},
		{"s3 no error", ProviderS3, testConfig{"configVal", nil}, assert.NoError},
		{"unknown w/ error", ProviderUnknown, testConfig{"configVal", assert.AnError}, assert.Error},
		{"s3 w/ error", ProviderS3, testConfig{"configVal", assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			s, err := NewStorage(test.p, test.c)
			test.errCheck(t, err, clues.ToCore(err))

			// remaining tests are dependent upon error-free state
			if test.c.err != nil {
				return
			}

			assert.Equalf(t,
				test.p,
				s.Provider,
				"expected storage provider [%s], got [%s]", test.p, s.Provider)
			assert.Equalf(t,
				test.c.expect,
				s.Config["expect"],
				"expected storage config [%s], got [%s]", test.c.expect, s.Config["expect"])
		})
	}
}

// Test ArePathsEquivalent
func (suite *StorageUnitSuite) TestArePathsEquivalent() {
	table := []struct {
		name        string
		path1       string
		path2       string
		expected    bool
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name:        "additional backslash in path",
			path1:       "/home/backups/",
			path2:       "/home/backups",
			expected:    true,
			expectedErr: assert.NoError,
		},
		{
			name:        "leading whitespace in path",
			path1:       " /home/backups",
			path2:       "/home/backups",
			expected:    true,
			expectedErr: assert.NoError,
		},
		{
			name:        "different paths",
			path1:       "/home/backups/1",
			path2:       "/home/backups/2",
			expected:    false,
			expectedErr: assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			actual, err := ArePathsEquivalent(test.path1, test.path2)
			assert.Equal(t, test.expected, actual)

			test.expectedErr(t, err)
		})
	}
}

// Test IsValidPath
func (suite *StorageUnitSuite) TestIsValidPath() {
	table := []struct {
		name     string
		path     string
		create   bool
		expected bool
	}{
		{
			name:     "valid directory",
			path:     "/tmp/backups/",
			expected: true,
		},
		{
			name:     "valid file path",
			path:     "/tmp/backups/a.txt",
			expected: true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			actual := IsValidPath(test.path)
			assert.Equal(t, test.expected, actual)
		})
	}
}
