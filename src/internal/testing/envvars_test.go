package testing

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnvvarsTestSuite struct {
	suite.Suite
}

func TestEnvvarsSuite(t *testing.T) {
	os.Setenv("Foo", "1")
	suite.Run(t, new(EnvvarsTestSuite))
}

func (suite EnvvarsTestSuite) TestRunOnAny() {

	table := []struct {
		name     string
		param    string
		expected bool
	}{
		{
			name:     "Valid Environment",
			param:    "Foo",
			expected: true,
		},
		{
			name:     "Invalid Environment",
			param:    "bar",
			expected: false,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := RunOnAny(test.param)
			suite.Equal((result == nil), test.expected)
		})
	}
}
