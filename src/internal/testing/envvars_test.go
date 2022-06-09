package testing

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
	os.Setenv("Foo", "1")
	suite.Run(t, new(EnvvarsTestSuite))
}

func (suite EnvvarsTestSuite) TestRunOnAny() {

	table := []struct {
		name     string
		param    string
		function assert.ErrorAssertionFunc
	}{
		{
			name:     "Valid Environment",
			param:    "Foo",
			function: assert.NoError,
		},
		{
			name:     "Invalid Environment",
			param:    "bar",
			function: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := RunOnAny(test.param)
			test.function(suite.T(), result)
		})
	}
}
