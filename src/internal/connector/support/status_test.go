package support

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
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
			name: "Test: Status Failed",
			params: statusParams{
				Restore,
				12, 9, 8,
				WrapAndAppend("tres", errors.New("three"), WrapAndAppend("arc376", errors.New("one"), errors.New("two"))),
			},
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			result := CreateStatus(
				ctx,
				test.params.operationType,
				test.params.folders,
				CollectionMetrics{test.params.objects, test.params.success, 0},
				test.params.err,
				"",
			)
			test.expect(t, result.incomplete, "status is incomplete")
		})
	}
}

func (suite *GCStatusTestSuite) TestCreateStatus_InvalidStatus() {
	t := suite.T()
	params := statusParams{Backup, 9, 3, 13, errors.New("invalidcl")}

	require.NotPanics(t, func() {
		ctx, flush := tester.NewContext()
		defer flush()

		CreateStatus(
			ctx,
			params.operationType,
			params.folders,
			CollectionMetrics{
				params.objects,
				params.success,
				0,
			},
			params.err,
			"",
		)
	})
}

func (suite *GCStatusTestSuite) TestMergeStatus() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name         string
		one          ConnectorOperationStatus
		two          ConnectorOperationStatus
		expected     statusParams
		isIncomplete assert.BoolAssertionFunc
	}{
		{
			name:         "Test:  Status + unknown",
			one:          *CreateStatus(ctx, Backup, 1, CollectionMetrics{1, 1, 0}, nil, ""),
			two:          ConnectorOperationStatus{},
			expected:     statusParams{Backup, 1, 1, 1, nil},
			isIncomplete: assert.False,
		},
		{
			name:         "Test: unknown + Status",
			one:          ConnectorOperationStatus{},
			two:          *CreateStatus(ctx, Backup, 1, CollectionMetrics{1, 1, 0}, nil, ""),
			expected:     statusParams{Backup, 1, 1, 1, nil},
			isIncomplete: assert.False,
		},
		{
			name:         "Test: Successful + Successful",
			one:          *CreateStatus(ctx, Backup, 1, CollectionMetrics{1, 1, 0}, nil, ""),
			two:          *CreateStatus(ctx, Backup, 3, CollectionMetrics{3, 3, 0}, nil, ""),
			expected:     statusParams{Backup, 4, 4, 4, nil},
			isIncomplete: assert.False,
		},
		{
			name: "Test: Successful + Unsuccessful",
			one:  *CreateStatus(ctx, Backup, 13, CollectionMetrics{17, 17, 0}, nil, ""),
			two: *CreateStatus(
				ctx,
				Backup,
				8,
				CollectionMetrics{
					12,
					9,
					0,
				},
				WrapAndAppend("tres", errors.New("three"), WrapAndAppend("arc376", errors.New("one"), errors.New("two"))),
				"",
			),
			expected:     statusParams{Backup, 29, 26, 21, nil},
			isIncomplete: assert.True,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			returned := MergeStatus(test.one, test.two)
			suite.Equal(returned.FolderCount, test.expected.folders)
			suite.Equal(returned.ObjectCount, test.expected.objects)
			suite.Equal(returned.lastOperation, test.expected.operationType)
			suite.Equal(returned.Successful, test.expected.success)
			test.isIncomplete(t, returned.incomplete)
		})
	}
}
