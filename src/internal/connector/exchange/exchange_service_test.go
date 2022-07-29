package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExchangeServiceSuite struct {
	suite.Suite
}

func TestExchangeServiceSuite(t *testing.T) {
	suite.Run(t, new(ExchangeServiceSuite))
}

// TestExchangeService_optionsForMessages checks to ensure approved query
// options are added to the type specific RequestBuildConfiguration. Expected
// will be +1 on all select parameters
func (suite *ExchangeServiceSuite) TestExchangeService_optionsForMessages() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Accepted",
			params:     []string{"subject"},
			checkError: assert.NoError,
		},
		{
			name:       "Multiple Accepted",
			params:     []string{"webLink", "parentFolderId"},
			checkError: assert.NoError,
		},
		{
			name:       "Incorrect param",
			params:     []string{"status"},
			checkError: assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			config, err := optionsForMessages(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(len(config.QueryParameters.Select), len(test.params)+1)
			}
		})
	}
}

// TestExchangeService_optionsForFolders ensures that approved query options
// are added to the RequestBuildConfiguration. Expected will always be +1
// on than the input as "id" are always included within the select parameters
func (suite *ExchangeServiceSuite) TestExchangeService_optionsForFolders() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
		expected   int
	}{
		{
			name:       "Accepted",
			params:     []string{"displayName"},
			checkError: assert.NoError,
			expected:   2,
		},
		{
			name:       "Multiple Accepted",
			params:     []string{"displayName", "parentFolderId"},
			checkError: assert.NoError,
			expected:   3,
		},
		{
			name:       "Incorrect param",
			params:     []string{"status"},
			checkError: assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			config, err := optionsForMailFolders(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(test.expected, len(config.QueryParameters.Select))
			}
		})

	}

}
