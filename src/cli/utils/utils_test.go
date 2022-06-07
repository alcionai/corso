package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli/utils"
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
			errCheck: assert.NoError,
		},
		{
			props:    map[string]string{"not-exists": ""},
			errCheck: assert.Error,
		},
	}
	for _, test := range table {
		test.errCheck(suite.T(), utils.RequireProps(test.props))
	}
}
