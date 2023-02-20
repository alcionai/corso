package status

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

func (suite *StatusUnitSuite) TestNew() {
	ctx, flush := tester.NewContext()
	defer flush()

	result := New(
		ctx,
		Backup,
		Counts{1, 1, 1, 1},
		"details",
		true)
	assert.True(suite.T(), result.incomplete, "status is incomplete")
}

func (suite *StatusUnitSuite) TestMergeStatus() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name         string
		one          ConnectorStatus
		two          ConnectorStatus
		expectOP     Operation
		expected     Counts
		isIncomplete assert.BoolAssertionFunc
	}{
		{
			name:         "Test:  Status + unknown",
			one:          New(ctx, Backup, Counts{1, 1, 1, 1}, "details", false),
			two:          ConnectorStatus{},
			expectOP:     Backup,
			expected:     Counts{1, 1, 1, 1},
			isIncomplete: assert.False,
		},
		{
			name:         "Test: unknown + Status",
			one:          ConnectorStatus{},
			two:          New(ctx, Backup, Counts{1, 1, 1, 1}, "details", false),
			expectOP:     Backup,
			expected:     Counts{1, 1, 1, 1},
			isIncomplete: assert.False,
		},
		{
			name:         "Test: complete + complete",
			one:          New(ctx, Backup, Counts{1, 1, 3, 0}, "details", false),
			two:          New(ctx, Backup, Counts{3, 3, 3, 0}, "details", false),
			expectOP:     Backup,
			expected:     Counts{4, 4, 6, 0},
			isIncomplete: assert.False,
		},
		{
			name:         "Test: complete + incomplete",
			one:          New(ctx, Restore, Counts{17, 17, 13, 0}, "details", false),
			two:          New(ctx, Restore, Counts{12, 9, 8, 0}, "details", true),
			expectOP:     Restore,
			expected:     Counts{29, 26, 21, 0},
			isIncomplete: assert.True,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			result := Combine(test.one, test.two)
			assert.Equal(t, result.op, test.expectOP)
			assert.Equal(t, test.expected.Folders, result.Metrics.Folders)
			assert.Equal(t, test.expected.Objects, result.Metrics.Objects)
			assert.Equal(t, test.expected.Successes, result.Metrics.Successes)
			assert.Equal(t, test.expected.Bytes, result.Metrics.Bytes)
			test.isIncomplete(t, result.incomplete)
		})
	}
}
