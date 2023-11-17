package fault_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type ItemUnitSuite struct {
	tester.Suite
}

func TestItemUnitSuite(t *testing.T) {
	suite.Run(t, &ItemUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemUnitSuite) TestItem_Error() {
	var (
		t = suite.T()
		i *fault.Item
	)

	assert.Contains(t, i.Error(), "nil")

	i = &fault.Item{}
	assert.Contains(t, i.Error(), "unknown type")

	i = &fault.Item{Type: fault.FileType}
	assert.Contains(t, i.Error(), fault.FileType)
}

func (suite *ItemUnitSuite) TestContainerErr() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := fault.ContainerErr(clues.New("foo"), "ns", "id", "name", addtl)

	expect := fault.Item{
		Namespace:  "ns",
		ID:         "id",
		Name:       "name",
		Type:       fault.ContainerType,
		Cause:      "foo",
		Additional: addtl,
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestFileErr() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := fault.FileErr(clues.New("foo"), "ns", "id", "name", addtl)

	expect := fault.Item{
		Namespace:  "ns",
		ID:         "id",
		Name:       "name",
		Type:       fault.FileType,
		Cause:      "foo",
		Additional: addtl,
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestOwnerErr() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := fault.OwnerErr(clues.New("foo"), "ns", "id", "name", addtl)

	expect := fault.Item{
		Namespace:  "ns",
		ID:         "id",
		Name:       "name",
		Type:       fault.ResourceOwnerType,
		Cause:      "foo",
		Additional: addtl,
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestItemType_Printable() {
	table := []struct {
		t      fault.ItemType
		expect string
	}{
		{
			t:      fault.FileType,
			expect: "File",
		},
		{
			t:      fault.ContainerType,
			expect: "Container",
		},
		{
			t:      fault.ResourceOwnerType,
			expect: "Resource Owner",
		},
		{
			t:      fault.ItemType("foo"),
			expect: "Unknown",
		},
	}
	for _, test := range table {
		suite.Run(string(test.t), func() {
			assert.Equal(suite.T(), test.expect, test.t.Printable())
		})
	}
}

func (suite *ItemUnitSuite) TestItem_HeadersValues() {
	var (
		err   = assert.AnError
		cause = err.Error()
		addtl = map[string]any{
			fault.AddtlContainerID:   "cid",
			fault.AddtlContainerName: "cname",
		}
	)

	table := []struct {
		name   string
		item   *fault.Item
		expect []string
	}{
		{
			name:   "file",
			item:   fault.FileErr(assert.AnError, "ns", "id", "name", addtl),
			expect: []string{"Error", fault.FileType.Printable(), "name", "cname", cause},
		},
		{
			name:   "container",
			item:   fault.ContainerErr(assert.AnError, "ns", "id", "name", addtl),
			expect: []string{"Error", fault.ContainerType.Printable(), "name", "cname", cause},
		},
		{
			name:   "owner",
			item:   fault.OwnerErr(assert.AnError, "ns", "id", "name", nil),
			expect: []string{"Error", fault.ResourceOwnerType.Printable(), "name", "", cause},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, []string{"Action", "Type", "Name", "Container", "Cause"}, test.item.Headers(false))
			assert.Equal(t, test.expect, test.item.Values(false))

			assert.Equal(t, []string{"Action", "Type", "Name", "Container", "Cause"}, test.item.Headers(true))
			assert.Equal(t, test.expect, test.item.Values(true))
		})
	}
}
