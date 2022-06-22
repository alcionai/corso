package support

import (
	"errors"
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

// 			operationType, objects, success, folders, errCount int, errStatus string

type statusParams struct {
	operationType Operation
	objects       int
	success       int
	folders       int
	err           error
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
			params:     statusParams{Backup, 12, 12, 3, nil},
			expected:   false,
			checkError: assert.Nil,
		},
		{
			name:       "Test: Status Failed",
			params:     statusParams{Restore, 12, 9, 8, WrapAndAppend("tres", errors.New("three"), WrapAndAppend("arc376", errors.New("one"), errors.New("two")))},
			expected:   true,
			checkError: assert.Nil,
		},
		{
			name:       "Invalid status",
			params:     statusParams{Backup, 9, 3, 12, errors.New("invalidcl")},
			expected:   false,
			checkError: assert.NotNil,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateStatus(test.params.operationType, test.params.objects,
				test.params.success, test.params.folders, test.params.err)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(result.incomplete, test.expected)
			}
		})

	}
}
