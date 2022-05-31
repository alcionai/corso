package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CliRepoSuite struct {
	suite.Suite
}

func TestCliRepoSuite(t *testing.T) {
	suite.Run(t, new(CliRepoSuite))
}

func (suite *CliRepoSuite) TestRequireProps() {
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
		test.errCheck(suite.T(), requireProps(test.props))
	}
}
