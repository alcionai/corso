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
	env_variable := "TEST_ENVVARS_SUITE"
	os.Setenv(env_variable, "1")
	table := []struct {
		name     string
		param    string
		function assert.ErrorAssertionFunc
	}{
		{
			name:     "Valid Environment",
			param:    env_variable,
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
	os.Unsetenv(env_variable)
}
