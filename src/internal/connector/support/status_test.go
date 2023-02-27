package support

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type StatusUnitSuite struct {
	tester.Suite
}

func TestGraphConnectorStatus(t *testing.T) {
	suite.Run(t, &StatusUnitSuite{tester.NewUnitSuite(t)})
}

func metricsMatch(t *testing.T, expect, result CollectionMetrics) {
	assert.Equal(t, expect.Bytes, result.Bytes, "bytes")
	assert.Equal(t, expect.Objects, result.Objects, "objects")
	assert.Equal(t, expect.Successes, result.Successes, "successes")
}

func (suite *StatusUnitSuite) TestCreateStatus() {
	table := []struct {
		name    string
		op      Operation
		metrics CollectionMetrics
		folders int
	}{
		{
			name:    "Backup",
			op:      Backup,
			metrics: CollectionMetrics{12, 12, 3},
			folders: 1,
		},
		{
			name:    "Backup",
			op:      Restore,
			metrics: CollectionMetrics{12, 9, 8},
			folders: 2,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			result := CreateStatus(
				ctx,
				test.op,
				test.folders,
				test.metrics,
				"details")
			assert.Equal(t, test.op, result.op, "operation")
			assert.Equal(t, test.folders, result.Folders, "folders")
			metricsMatch(t, test.metrics, result.Metrics)
		})
	}
}

func (suite *StatusUnitSuite) TestMergeStatus() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name          string
		one           ConnectorOperationStatus
		two           ConnectorOperationStatus
		expectOp      Operation
		expectMetrics CollectionMetrics
		expectFolders int
	}{
		{
			name:          "Test:  Status + unknown",
			one:           *CreateStatus(ctx, Backup, 1, CollectionMetrics{1, 1, 0}, ""),
			two:           ConnectorOperationStatus{},
			expectOp:      Backup,
			expectMetrics: CollectionMetrics{1, 1, 0},
			expectFolders: 1,
		},
		{
			name:          "Test: unknown + Status",
			one:           ConnectorOperationStatus{},
			two:           *CreateStatus(ctx, Backup, 1, CollectionMetrics{1, 1, 0}, ""),
			expectOp:      Backup,
			expectMetrics: CollectionMetrics{1, 1, 0},
			expectFolders: 1,
		},
		{
			name:          "Test: Successful + Successful",
			one:           *CreateStatus(ctx, Backup, 1, CollectionMetrics{1, 1, 0}, ""),
			two:           *CreateStatus(ctx, Backup, 3, CollectionMetrics{3, 3, 0}, ""),
			expectOp:      Backup,
			expectMetrics: CollectionMetrics{4, 4, 0},
			expectFolders: 4,
		},
		{
			name:          "Test: Successful + Unsuccessful",
			one:           *CreateStatus(ctx, Backup, 13, CollectionMetrics{17, 17, 0}, ""),
			two:           *CreateStatus(ctx, Backup, 8, CollectionMetrics{12, 9, 0}, ""),
			expectOp:      Backup,
			expectMetrics: CollectionMetrics{29, 26, 0},
			expectFolders: 21,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := MergeStatus(test.one, test.two)
			assert.Equal(t, test.expectFolders, result.Folders, "folders")
			assert.Equal(t, test.expectOp, result.op, "operation")
			metricsMatch(t, test.expectMetrics, result.Metrics)
		})
	}
}
