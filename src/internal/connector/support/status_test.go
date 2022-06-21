package support

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GCStatusTestSuite struct {
	suite.Suite
}

func TestGraphConnectorStatus(t *testing.T) {
	suite.Run(t, &GCStatusTestSuite{})
}

func (suite *GCStatusTestSuite) TestGetOperation() {
	table := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Backup Config",
			input:    0,
			expected: "Backup",
		},
		{
			name:     "Restore Config",
			input:    1,
			expected: "Restore",
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := GetOperation(operation(test.input))
			suite.Equal(result, test.expected)
		})
	}
}

// 			operationType, objects, success, folders, errCount int, errStatus string

type statusParams struct {
	operationType int
	objects       int
	success       int
	folders       int
	errCount      int
	errStatus     string
}

func (suite *GCStatusTestSuite) TestCreateStatus() {
	table := []struct {
		name       string
		params     statusParams
		expected   bool
		checkError assert.ValueAssertionFunc
	}{
		{
			name:       "Test: Status Success",
			params:     statusParams{0, 12, 12, 3, 0, ""},
			expected:   false,
			checkError: assert.Nil,
		},
		{
			name:       "Test: Status Failed",
			params:     statusParams{1, 12, 9, 8, 3, "Unable to convert Integer, network error, unexpected interruption"},
			expected:   true,
			checkError: assert.Nil,
		},
		{
			name:       "Invalid status",
			params:     statusParams{0, 9, 3, 12, 2, "We aren't getting here"},
			expected:   false,
			checkError: assert.NotNil,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateStatus(test.params.operationType, test.params.objects,
				test.params.success, test.params.folders, test.params.errCount, test.params.errStatus)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(result.incomplete, test.expected)
			}
		})

	}
}
