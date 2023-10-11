package fault_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type SkippedUnitSuite struct {
	tester.Suite
}

func TestSkippedUnitSuite(t *testing.T) {
	suite.Run(t, &SkippedUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SkippedUnitSuite) TestSkipped_String() {
	var (
		t = suite.T()
		i *fault.Skipped
	)

	assert.Contains(t, i.String(), "nil")

	i = &fault.Skipped{fault.Item{}}
	assert.Contains(t, i.String(), "unknown type")

	i = &fault.Skipped{
		fault.Item{
			Type: fault.FileType,
		},
	}
	assert.Contains(t, i.Item.Error(), fault.FileType)
}

func (suite *SkippedUnitSuite) TestContainerSkip() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := fault.ContainerSkip(fault.SkipMalware, "ns", "id", "name", addtl)

	expect := fault.Item{
		Namespace:  "ns",
		ID:         "id",
		Name:       "name",
		Type:       fault.ContainerType,
		Cause:      string(fault.SkipMalware),
		Additional: addtl,
	}

	assert.Equal(t, fault.Skipped{expect}, *i)
}

func (suite *SkippedUnitSuite) TestFileSkip() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := fault.FileSkip(fault.SkipMalware, "ns", "id", "name", addtl)

	expect := fault.Item{
		Namespace:  "ns",
		ID:         "id",
		Name:       "name",
		Type:       fault.FileType,
		Cause:      string(fault.SkipMalware),
		Additional: addtl,
	}

	assert.Equal(t, fault.Skipped{expect}, *i)
}

func (suite *SkippedUnitSuite) TestOwnerSkip() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := fault.OwnerSkip(fault.SkipMalware, "ns", "id", "name", addtl)

	expect := fault.Item{
		Namespace:  "ns",
		ID:         "id",
		Name:       "name",
		Type:       fault.ResourceOwnerType,
		Cause:      string(fault.SkipMalware),
		Additional: addtl,
	}

	assert.Equal(t, fault.Skipped{expect}, *i)
}

func (suite *SkippedUnitSuite) TestSkipped_HeadersValues() {
	addtl := map[string]any{
		fault.AddtlContainerID:   "cid",
		fault.AddtlContainerName: "cname",
	}

	table := []struct {
		name   string
		skip   *fault.Skipped
		expect []string
	}{
		{
			name:   "file",
			skip:   fault.FileSkip(fault.SkipMalware, "ns", "id", "name", addtl),
			expect: []string{"Skip", fault.FileType.Printable(), "name", "cname", string(fault.SkipMalware)},
		},
		{
			name:   "container",
			skip:   fault.ContainerSkip(fault.SkipMalware, "ns", "id", "name", addtl),
			expect: []string{"Skip", fault.ContainerType.Printable(), "name", "cname", string(fault.SkipMalware)},
		},
		{
			name:   "owner",
			skip:   fault.OwnerSkip(fault.SkipMalware, "ns", "id", "name", nil),
			expect: []string{"Skip", fault.ResourceOwnerType.Printable(), "name", "", string(fault.SkipMalware)},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, []string{"Action", "Type", "Name", "Container", "Cause"}, test.skip.Headers())
			assert.Equal(t, test.expect, test.skip.Values())
		})
	}
}

func (suite *SkippedUnitSuite) TestBus_AddSkip() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	n := fault.New(true)
	require.NotNil(t, n)

	n.Fail(assert.AnError)
	assert.Len(t, n.Skipped(), 0)

	n.AddRecoverable(ctx, assert.AnError)
	assert.Len(t, n.Skipped(), 0)

	n.AddSkip(ctx, fault.OwnerSkip(fault.SkipMalware, "ns", "id", "name", nil))
	assert.Len(t, n.Skipped(), 1)
}
