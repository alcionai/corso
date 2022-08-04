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
