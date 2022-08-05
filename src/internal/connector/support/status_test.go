package support

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		name   string
		params statusParams
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "Test: Status Success",
			params: statusParams{Backup, 12, 12, 3, nil},
			expect: assert.False,
		},
		{
			name:   "Test: Status Failed",
			params: statusParams{Restore, 12, 9, 8, WrapAndAppend("tres", errors.New("three"), WrapAndAppend("arc376", errors.New("one"), errors.New("two")))},
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := CreateStatus(
				context.Background(),
				test.params.operationType,
				test.params.objects,
				test.params.success,
				test.params.folders,
				test.params.err)
			test.expect(t, result.incomplete, "status is incomplete")
		})

	}
}

func (suite *GCStatusTestSuite) TestMergeStatus() {
	simpleContext := context.Background()
	table := []struct {
		name         string
		one          *ConnectorOperationStatus
		two          *ConnectorOperationStatus
		expected     statusParams
		isIncomplete assert.BoolAssertionFunc
	}{
		{
			name:         "Test:  Status + nil",
			one:          CreateStatus(simpleContext, Backup, 1, 1, 1, nil),
			two:          nil,
			expected:     statusParams{Backup, 1, 1, 1, nil},
			isIncomplete: assert.False,
		},
		{
			name:         "Test: nil + Status",
			one:          nil,
			two:          CreateStatus(simpleContext, Backup, 1, 1, 1, nil),
			expected:     statusParams{Backup, 1, 1, 1, nil},
			isIncomplete: assert.False,
		},
		{
			name:         "Test: Successful + Successful",
			one:          CreateStatus(simpleContext, Backup, 1, 1, 1, nil),
			two:          CreateStatus(simpleContext, Backup, 3, 3, 3, nil),
			expected:     statusParams{Backup, 4, 4, 4, nil},
			isIncomplete: assert.False,
		},
		{
			name: "Test: Successful + Unsuccessful",
			one:  CreateStatus(simpleContext, Backup, 17, 17, 13, nil),
			two: CreateStatus(
				simpleContext,
				Backup,
				12,
				9,
				8,
				WrapAndAppend("tres", errors.New("three"), WrapAndAppend("arc376", errors.New("one"), errors.New("two"))),
			),
			expected:     statusParams{Backup, 29, 26, 21, nil},
			isIncomplete: assert.True,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			returned := MergeStatus(test.one, test.two)
			suite.Equal(returned.folderCount, test.expected.folders)
			suite.Equal(returned.ObjectCount, test.expected.objects)
			suite.Equal(returned.LastOperation, test.expected.operationType)
			suite.Equal(returned.Successful, test.expected.success)
			test.isIncomplete(t, returned.incomplete)
		})
	}
}

func (suite *GCStatusTestSuite) TestCreateStatus_InvalidStatus() {
	t := suite.T()
	params := statusParams{Backup, 9, 3, 13, errors.New("invalidcl")}
	require.Panics(t, func() {
		CreateStatus(
			context.Background(),
			params.operationType,
			params.objects,
			params.success,
			params.folders,
			params.err,
		)
	})
}
