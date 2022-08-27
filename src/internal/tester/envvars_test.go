package tester

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvvarsTestSuite struct {
	suite.Suite
}

func TestEnvvarsSuite(t *testing.T) {
	suite.Run(t, new(EnvvarsTestSuite))
}

func (suite *EnvvarsTestSuite) TestRunOnAny() {
	envVariable := "TEST_ENVVARS_SUITE"
	os.Setenv(envVariable, "1")

	table := []struct {
		name     string
		param    string
		function assert.ErrorAssertionFunc
	}{
		{
			name:     "Valid Environment",
			param:    envVariable,
			function: assert.NoError,
		},
		{
			name:     "Invalid Environment",
			param:    "TEST_ENVVARS_SUITE_INVALID",
			function: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := RunOnAny(test.param)
			test.function(suite.T(), result)
		})
	}

	os.Unsetenv(envVariable)
}
